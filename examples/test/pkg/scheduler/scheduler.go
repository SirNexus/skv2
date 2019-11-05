package scheduler

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/solo-io/autopilot/pkg/metrics"
	"github.com/solo-io/autopilot/pkg/utils"

	v1 "github.com/solo-io/autopilot/examples/test/pkg/apis/tests/v1"

	config "github.com/solo-io/autopilot/examples/test/pkg/config"
	finalizer "github.com/solo-io/autopilot/examples/test/pkg/finalizer"
	initializing "github.com/solo-io/autopilot/examples/test/pkg/workers/initializing"
	processing "github.com/solo-io/autopilot/examples/test/pkg/workers/processing"
	aliases "github.com/solo-io/autopilot/pkg/aliases"
)

var log = logf.Log.WithName("scheduler")

func AddToManager(ctx context.Context, mgr manager.Manager, namespace string) error {
	scheduler, err := NewScheduler(ctx, mgr, namespace)
	if err != nil {
		return err
	}
	// Create a new controller
	c, err := controller.New("test-controller", mgr, controller.Options{Reconciler: scheduler})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Test
	log.Info("Registering watch for primary resource Test")
	err = c.Watch(&source.Kind{Type: &v1.Test{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource VirtualServices and requeue the owner Test
	log.Info("Registering watch for primary resource secondary resource VirtualServices")
	err = c.Watch(&source.Kind{Type: &aliases.VirtualService{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &v1.Test{},
	})
	if err != nil {
		return err
	}

	return nil

}

var WorkInterval = config.WorkInterval
var FinalizerName = "test-finalizer"

type Scheduler struct {
	ctx       context.Context
	mgr       manager.Manager
	Metrics   metrics.Metrics
	namespace string
}

func NewScheduler(ctx context.Context, mgr manager.Manager, namespace string) (*Scheduler, error) {
	metricsFactory, err := metrics.NewFactory(config.MetricsServer, config.MeshProvider, time.Second*30)
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		ctx:       ctx,
		mgr:       mgr,
		Metrics:   metricsFactory.Observer(),
		namespace: namespace,
	}, nil
}

func (s *Scheduler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	result := reconcile.Result{RequeueAfter: WorkInterval}

	test := &v1.Test{}
	test.Namespace = request.Namespace
	test.Name = request.Name

	kube := utils.NewEzKube(test, s.mgr)

	if err := kube.Get(s.ctx, test); err != nil {
		// garbage collection and finalizers should handle cleaning up after deletion
		if errors.IsNotFound(err) {
			return result, nil
		}
		return result, err
	}
	// examine DeletionTimestamp to determine if object is under deletion
	if test.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !utils.ContainsString(test.Finalizers, FinalizerName) {
			test.Finalizers = append(test.Finalizers, FinalizerName)
			if err := kube.Ensure(s.ctx, test); err != nil {
				return result, err
			}
		}
	} else {
		// The object is being deleted
		if utils.ContainsString(test.Finalizers, FinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			if err := (&finalizer.Finalizer{Kube: kube}).Finalize(s.ctx, test); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return result, err
			}

			// remove our finalizer from the list and update it.
			test.Finalizers = utils.RemoveString(test.Finalizers, FinalizerName)
			if err := kube.Ensure(s.ctx, test); err != nil {
				return result, err
			}
		}

		return result, nil
	}

	switch test.Status.Phase {
	case "", v1.TestPhaseInitializing:
		log.Info("Syncing Test %v in phase Initializing", test.Name)
		inputs, err := s.makeInitializingInputs(kube)
		if err != nil {
			return result, err
		}
		outputs, nextPhase, statusInfo, err := (&initializing.Worker{Kube: kube}).Sync(s.ctx, test, inputs)
		if err != nil {
			return result, err
		}
		for _, out := range outputs.VirtualServices {
			if err := kube.Ensure(s.ctx, out); err != nil {
				return result, err
			}
		}

		test.Status.Phase = nextPhase
		if statusInfo != nil {
			test.Status.TestStatusInfo = *statusInfo
		}
		if err := kube.UpdateStatus(s.ctx, test); err != nil {
			return result, err
		}

		return result, err
	case v1.TestPhaseProcessing:
		log.Info("Syncing Test %v in phase Processing", test.Name)
		inputs, err := s.makeProcessingInputs(kube)
		if err != nil {
			return result, err
		}
		nextPhase, statusInfo, err := (&processing.Worker{Kube: kube}).Sync(s.ctx, test, inputs)
		if err != nil {
			return result, err
		}

		test.Status.Phase = nextPhase
		if statusInfo != nil {
			test.Status.TestStatusInfo = *statusInfo
		}
		if err := kube.UpdateStatus(s.ctx, test); err != nil {
			return result, err
		}

		return result, err
	case v1.TestPhaseFinished:
		log.Info("Syncing Test %v in phase Finished", test.Name)
		// end state, do not requeue
		return reconcile.Result{}, nil
	case v1.TestPhaseFailed:
		log.Info("Syncing Test %v in phase Failed", test.Name)
		// end state, do not requeue
		return reconcile.Result{}, nil
	}
	return result, fmt.Errorf("cannot process Test in unknown phase: %v", test.Status.Phase)
}
func (s *Scheduler) makeInitializingInputs(kube utils.EzKube) (initializing.Inputs, error) {
	var (
		inputs initializing.Inputs
		err    error
	)
	inputs.Services, err = kube.ListServices(s.ctx, s.namespace)
	if err != nil {
		return inputs, err
	}

	return inputs, err
}
func (s *Scheduler) makeProcessingInputs(kube utils.EzKube) (processing.Inputs, error) {
	var (
		inputs processing.Inputs
		err    error
	)
	inputs.Metrics = s.Metrics

	return inputs, err
}