// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	rbac_authorization_k8s_io_v1 "k8s.io/api/rbac/v1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type RoleSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*rbac_authorization_k8s_io_v1.Role) bool) []*rbac_authorization_k8s_io_v1.Role
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.Role) bool) []*rbac_authorization_k8s_io_v1.Role
	// Return the Set as a map of key to resource.
	Map() map[string]*rbac_authorization_k8s_io_v1.Role
	// Insert a resource into the set.
	Insert(role ...*rbac_authorization_k8s_io_v1.Role)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(roleSet RoleSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(role ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(role ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RoleSet) RoleSet
	// Return the difference with the provided set
	Difference(set RoleSet) RoleSet
	// Return the intersection with the provided set
	Intersection(set RoleSet) RoleSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.Role, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RoleSet
	Delta(newSet RoleSet) sksets.ResourceDelta
	// Create a deep copy of the current RoleSet
	Clone() RoleSet
}

func makeGenericRoleSet(roleList []*rbac_authorization_k8s_io_v1.Role) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range roleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type roleSet struct {
	set sksets.ResourceSet
}

func NewRoleSet(roleList ...*rbac_authorization_k8s_io_v1.Role) RoleSet {
	return &roleSet{set: makeGenericRoleSet(roleList)}
}

func NewRoleSetFromList(roleList *rbac_authorization_k8s_io_v1.RoleList) RoleSet {
	list := make([]*rbac_authorization_k8s_io_v1.Role, 0, len(roleList.Items))
	for idx := range roleList.Items {
		list = append(list, &roleList.Items[idx])
	}
	return &roleSet{set: makeGenericRoleSet(list)}
}

func (s *roleSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *roleSet) List(filterResource ...func(*rbac_authorization_k8s_io_v1.Role) bool) []*rbac_authorization_k8s_io_v1.Role {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.Role))
		})
	}

	objs := s.Generic().List(genericFilters...)
	roleList := make([]*rbac_authorization_k8s_io_v1.Role, 0, len(objs))
	for _, obj := range objs {
		roleList = append(roleList, obj.(*rbac_authorization_k8s_io_v1.Role))
	}
	return roleList
}

func (s *roleSet) UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.Role) bool) []*rbac_authorization_k8s_io_v1.Role {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.Role))
		})
	}

	var roleList []*rbac_authorization_k8s_io_v1.Role
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		roleList = append(roleList, obj.(*rbac_authorization_k8s_io_v1.Role))
	}
	return roleList
}

func (s *roleSet) Map() map[string]*rbac_authorization_k8s_io_v1.Role {
	if s == nil {
		return nil
	}

	newMap := map[string]*rbac_authorization_k8s_io_v1.Role{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*rbac_authorization_k8s_io_v1.Role)
	}
	return newMap
}

func (s *roleSet) Insert(
	roleList ...*rbac_authorization_k8s_io_v1.Role,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range roleList {
		s.Generic().Insert(obj)
	}
}

func (s *roleSet) Has(role ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(role)
}

func (s *roleSet) Equal(
	roleSet RoleSet,
) bool {
	if s == nil {
		return roleSet == nil
	}
	return s.Generic().Equal(roleSet.Generic())
}

func (s *roleSet) Delete(Role ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(Role)
}

func (s *roleSet) Union(set RoleSet) RoleSet {
	if s == nil {
		return set
	}
	return NewRoleSet(append(s.List(), set.List()...)...)
}

func (s *roleSet) Difference(set RoleSet) RoleSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &roleSet{set: newSet}
}

func (s *roleSet) Intersection(set RoleSet) RoleSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var roleList []*rbac_authorization_k8s_io_v1.Role
	for _, obj := range newSet.List() {
		roleList = append(roleList, obj.(*rbac_authorization_k8s_io_v1.Role))
	}
	return NewRoleSet(roleList...)
}

func (s *roleSet) Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.Role, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find Role %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&rbac_authorization_k8s_io_v1.Role{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*rbac_authorization_k8s_io_v1.Role), nil
}

func (s *roleSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *roleSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *roleSet) Delta(newSet RoleSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *roleSet) Clone() RoleSet {
	if s == nil {
		return nil
	}
	return &roleSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type RoleBindingSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*rbac_authorization_k8s_io_v1.RoleBinding) bool) []*rbac_authorization_k8s_io_v1.RoleBinding
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.RoleBinding) bool) []*rbac_authorization_k8s_io_v1.RoleBinding
	// Return the Set as a map of key to resource.
	Map() map[string]*rbac_authorization_k8s_io_v1.RoleBinding
	// Insert a resource into the set.
	Insert(roleBinding ...*rbac_authorization_k8s_io_v1.RoleBinding)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(roleBindingSet RoleBindingSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(roleBinding ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(roleBinding ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RoleBindingSet) RoleBindingSet
	// Return the difference with the provided set
	Difference(set RoleBindingSet) RoleBindingSet
	// Return the intersection with the provided set
	Intersection(set RoleBindingSet) RoleBindingSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.RoleBinding, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RoleBindingSet
	Delta(newSet RoleBindingSet) sksets.ResourceDelta
	// Create a deep copy of the current RoleBindingSet
	Clone() RoleBindingSet
}

func makeGenericRoleBindingSet(roleBindingList []*rbac_authorization_k8s_io_v1.RoleBinding) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range roleBindingList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type roleBindingSet struct {
	set sksets.ResourceSet
}

func NewRoleBindingSet(roleBindingList ...*rbac_authorization_k8s_io_v1.RoleBinding) RoleBindingSet {
	return &roleBindingSet{set: makeGenericRoleBindingSet(roleBindingList)}
}

func NewRoleBindingSetFromList(roleBindingList *rbac_authorization_k8s_io_v1.RoleBindingList) RoleBindingSet {
	list := make([]*rbac_authorization_k8s_io_v1.RoleBinding, 0, len(roleBindingList.Items))
	for idx := range roleBindingList.Items {
		list = append(list, &roleBindingList.Items[idx])
	}
	return &roleBindingSet{set: makeGenericRoleBindingSet(list)}
}

func (s *roleBindingSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *roleBindingSet) List(filterResource ...func(*rbac_authorization_k8s_io_v1.RoleBinding) bool) []*rbac_authorization_k8s_io_v1.RoleBinding {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.RoleBinding))
		})
	}

	objs := s.Generic().List(genericFilters...)
	roleBindingList := make([]*rbac_authorization_k8s_io_v1.RoleBinding, 0, len(objs))
	for _, obj := range objs {
		roleBindingList = append(roleBindingList, obj.(*rbac_authorization_k8s_io_v1.RoleBinding))
	}
	return roleBindingList
}

func (s *roleBindingSet) UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.RoleBinding) bool) []*rbac_authorization_k8s_io_v1.RoleBinding {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.RoleBinding))
		})
	}

	var roleBindingList []*rbac_authorization_k8s_io_v1.RoleBinding
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		roleBindingList = append(roleBindingList, obj.(*rbac_authorization_k8s_io_v1.RoleBinding))
	}
	return roleBindingList
}

func (s *roleBindingSet) Map() map[string]*rbac_authorization_k8s_io_v1.RoleBinding {
	if s == nil {
		return nil
	}

	newMap := map[string]*rbac_authorization_k8s_io_v1.RoleBinding{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*rbac_authorization_k8s_io_v1.RoleBinding)
	}
	return newMap
}

func (s *roleBindingSet) Insert(
	roleBindingList ...*rbac_authorization_k8s_io_v1.RoleBinding,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range roleBindingList {
		s.Generic().Insert(obj)
	}
}

func (s *roleBindingSet) Has(roleBinding ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(roleBinding)
}

func (s *roleBindingSet) Equal(
	roleBindingSet RoleBindingSet,
) bool {
	if s == nil {
		return roleBindingSet == nil
	}
	return s.Generic().Equal(roleBindingSet.Generic())
}

func (s *roleBindingSet) Delete(RoleBinding ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RoleBinding)
}

func (s *roleBindingSet) Union(set RoleBindingSet) RoleBindingSet {
	if s == nil {
		return set
	}
	return NewRoleBindingSet(append(s.List(), set.List()...)...)
}

func (s *roleBindingSet) Difference(set RoleBindingSet) RoleBindingSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &roleBindingSet{set: newSet}
}

func (s *roleBindingSet) Intersection(set RoleBindingSet) RoleBindingSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var roleBindingList []*rbac_authorization_k8s_io_v1.RoleBinding
	for _, obj := range newSet.List() {
		roleBindingList = append(roleBindingList, obj.(*rbac_authorization_k8s_io_v1.RoleBinding))
	}
	return NewRoleBindingSet(roleBindingList...)
}

func (s *roleBindingSet) Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.RoleBinding, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RoleBinding %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&rbac_authorization_k8s_io_v1.RoleBinding{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*rbac_authorization_k8s_io_v1.RoleBinding), nil
}

func (s *roleBindingSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *roleBindingSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *roleBindingSet) Delta(newSet RoleBindingSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *roleBindingSet) Clone() RoleBindingSet {
	if s == nil {
		return nil
	}
	return &roleBindingSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type ClusterRoleSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRole) bool) []*rbac_authorization_k8s_io_v1.ClusterRole
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRole) bool) []*rbac_authorization_k8s_io_v1.ClusterRole
	// Return the Set as a map of key to resource.
	Map() map[string]*rbac_authorization_k8s_io_v1.ClusterRole
	// Insert a resource into the set.
	Insert(clusterRole ...*rbac_authorization_k8s_io_v1.ClusterRole)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(clusterRoleSet ClusterRoleSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(clusterRole ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(clusterRole ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ClusterRoleSet) ClusterRoleSet
	// Return the difference with the provided set
	Difference(set ClusterRoleSet) ClusterRoleSet
	// Return the intersection with the provided set
	Intersection(set ClusterRoleSet) ClusterRoleSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.ClusterRole, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ClusterRoleSet
	Delta(newSet ClusterRoleSet) sksets.ResourceDelta
	// Create a deep copy of the current ClusterRoleSet
	Clone() ClusterRoleSet
}

func makeGenericClusterRoleSet(clusterRoleList []*rbac_authorization_k8s_io_v1.ClusterRole) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range clusterRoleList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterRoleSet struct {
	set sksets.ResourceSet
}

func NewClusterRoleSet(clusterRoleList ...*rbac_authorization_k8s_io_v1.ClusterRole) ClusterRoleSet {
	return &clusterRoleSet{set: makeGenericClusterRoleSet(clusterRoleList)}
}

func NewClusterRoleSetFromList(clusterRoleList *rbac_authorization_k8s_io_v1.ClusterRoleList) ClusterRoleSet {
	list := make([]*rbac_authorization_k8s_io_v1.ClusterRole, 0, len(clusterRoleList.Items))
	for idx := range clusterRoleList.Items {
		list = append(list, &clusterRoleList.Items[idx])
	}
	return &clusterRoleSet{set: makeGenericClusterRoleSet(list)}
}

func (s *clusterRoleSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *clusterRoleSet) List(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRole) bool) []*rbac_authorization_k8s_io_v1.ClusterRole {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.ClusterRole))
		})
	}

	objs := s.Generic().List(genericFilters...)
	clusterRoleList := make([]*rbac_authorization_k8s_io_v1.ClusterRole, 0, len(objs))
	for _, obj := range objs {
		clusterRoleList = append(clusterRoleList, obj.(*rbac_authorization_k8s_io_v1.ClusterRole))
	}
	return clusterRoleList
}

func (s *clusterRoleSet) UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRole) bool) []*rbac_authorization_k8s_io_v1.ClusterRole {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.ClusterRole))
		})
	}

	var clusterRoleList []*rbac_authorization_k8s_io_v1.ClusterRole
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		clusterRoleList = append(clusterRoleList, obj.(*rbac_authorization_k8s_io_v1.ClusterRole))
	}
	return clusterRoleList
}

func (s *clusterRoleSet) Map() map[string]*rbac_authorization_k8s_io_v1.ClusterRole {
	if s == nil {
		return nil
	}

	newMap := map[string]*rbac_authorization_k8s_io_v1.ClusterRole{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*rbac_authorization_k8s_io_v1.ClusterRole)
	}
	return newMap
}

func (s *clusterRoleSet) Insert(
	clusterRoleList ...*rbac_authorization_k8s_io_v1.ClusterRole,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range clusterRoleList {
		s.Generic().Insert(obj)
	}
}

func (s *clusterRoleSet) Has(clusterRole ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(clusterRole)
}

func (s *clusterRoleSet) Equal(
	clusterRoleSet ClusterRoleSet,
) bool {
	if s == nil {
		return clusterRoleSet == nil
	}
	return s.Generic().Equal(clusterRoleSet.Generic())
}

func (s *clusterRoleSet) Delete(ClusterRole ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ClusterRole)
}

func (s *clusterRoleSet) Union(set ClusterRoleSet) ClusterRoleSet {
	if s == nil {
		return set
	}
	return NewClusterRoleSet(append(s.List(), set.List()...)...)
}

func (s *clusterRoleSet) Difference(set ClusterRoleSet) ClusterRoleSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &clusterRoleSet{set: newSet}
}

func (s *clusterRoleSet) Intersection(set ClusterRoleSet) ClusterRoleSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var clusterRoleList []*rbac_authorization_k8s_io_v1.ClusterRole
	for _, obj := range newSet.List() {
		clusterRoleList = append(clusterRoleList, obj.(*rbac_authorization_k8s_io_v1.ClusterRole))
	}
	return NewClusterRoleSet(clusterRoleList...)
}

func (s *clusterRoleSet) Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.ClusterRole, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ClusterRole %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&rbac_authorization_k8s_io_v1.ClusterRole{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*rbac_authorization_k8s_io_v1.ClusterRole), nil
}

func (s *clusterRoleSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *clusterRoleSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *clusterRoleSet) Delta(newSet ClusterRoleSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *clusterRoleSet) Clone() ClusterRoleSet {
	if s == nil {
		return nil
	}
	return &clusterRoleSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}

type ClusterRoleBindingSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRoleBinding) bool) []*rbac_authorization_k8s_io_v1.ClusterRoleBinding
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRoleBinding) bool) []*rbac_authorization_k8s_io_v1.ClusterRoleBinding
	// Return the Set as a map of key to resource.
	Map() map[string]*rbac_authorization_k8s_io_v1.ClusterRoleBinding
	// Insert a resource into the set.
	Insert(clusterRoleBinding ...*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(clusterRoleBindingSet ClusterRoleBindingSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(clusterRoleBinding ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(clusterRoleBinding ezkube.ResourceId)
	// Return the union with the provided set
	Union(set ClusterRoleBindingSet) ClusterRoleBindingSet
	// Return the difference with the provided set
	Difference(set ClusterRoleBindingSet) ClusterRoleBindingSet
	// Return the intersection with the provided set
	Intersection(set ClusterRoleBindingSet) ClusterRoleBindingSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.ClusterRoleBinding, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another ClusterRoleBindingSet
	Delta(newSet ClusterRoleBindingSet) sksets.ResourceDelta
	// Create a deep copy of the current ClusterRoleBindingSet
	Clone() ClusterRoleBindingSet
}

func makeGenericClusterRoleBindingSet(clusterRoleBindingList []*rbac_authorization_k8s_io_v1.ClusterRoleBinding) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range clusterRoleBindingList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type clusterRoleBindingSet struct {
	set sksets.ResourceSet
}

func NewClusterRoleBindingSet(clusterRoleBindingList ...*rbac_authorization_k8s_io_v1.ClusterRoleBinding) ClusterRoleBindingSet {
	return &clusterRoleBindingSet{set: makeGenericClusterRoleBindingSet(clusterRoleBindingList)}
}

func NewClusterRoleBindingSetFromList(clusterRoleBindingList *rbac_authorization_k8s_io_v1.ClusterRoleBindingList) ClusterRoleBindingSet {
	list := make([]*rbac_authorization_k8s_io_v1.ClusterRoleBinding, 0, len(clusterRoleBindingList.Items))
	for idx := range clusterRoleBindingList.Items {
		list = append(list, &clusterRoleBindingList.Items[idx])
	}
	return &clusterRoleBindingSet{set: makeGenericClusterRoleBindingSet(list)}
}

func (s *clusterRoleBindingSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *clusterRoleBindingSet) List(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRoleBinding) bool) []*rbac_authorization_k8s_io_v1.ClusterRoleBinding {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding))
		})
	}

	objs := s.Generic().List(genericFilters...)
	clusterRoleBindingList := make([]*rbac_authorization_k8s_io_v1.ClusterRoleBinding, 0, len(objs))
	for _, obj := range objs {
		clusterRoleBindingList = append(clusterRoleBindingList, obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding))
	}
	return clusterRoleBindingList
}

func (s *clusterRoleBindingSet) UnsortedList(filterResource ...func(*rbac_authorization_k8s_io_v1.ClusterRoleBinding) bool) []*rbac_authorization_k8s_io_v1.ClusterRoleBinding {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		filter := filter
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding))
		})
	}

	var clusterRoleBindingList []*rbac_authorization_k8s_io_v1.ClusterRoleBinding
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		clusterRoleBindingList = append(clusterRoleBindingList, obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding))
	}
	return clusterRoleBindingList
}

func (s *clusterRoleBindingSet) Map() map[string]*rbac_authorization_k8s_io_v1.ClusterRoleBinding {
	if s == nil {
		return nil
	}

	newMap := map[string]*rbac_authorization_k8s_io_v1.ClusterRoleBinding{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)
	}
	return newMap
}

func (s *clusterRoleBindingSet) Insert(
	clusterRoleBindingList ...*rbac_authorization_k8s_io_v1.ClusterRoleBinding,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range clusterRoleBindingList {
		s.Generic().Insert(obj)
	}
}

func (s *clusterRoleBindingSet) Has(clusterRoleBinding ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(clusterRoleBinding)
}

func (s *clusterRoleBindingSet) Equal(
	clusterRoleBindingSet ClusterRoleBindingSet,
) bool {
	if s == nil {
		return clusterRoleBindingSet == nil
	}
	return s.Generic().Equal(clusterRoleBindingSet.Generic())
}

func (s *clusterRoleBindingSet) Delete(ClusterRoleBinding ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(ClusterRoleBinding)
}

func (s *clusterRoleBindingSet) Union(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	if s == nil {
		return set
	}
	return NewClusterRoleBindingSet(append(s.List(), set.List()...)...)
}

func (s *clusterRoleBindingSet) Difference(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &clusterRoleBindingSet{set: newSet}
}

func (s *clusterRoleBindingSet) Intersection(set ClusterRoleBindingSet) ClusterRoleBindingSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var clusterRoleBindingList []*rbac_authorization_k8s_io_v1.ClusterRoleBinding
	for _, obj := range newSet.List() {
		clusterRoleBindingList = append(clusterRoleBindingList, obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding))
	}
	return NewClusterRoleBindingSet(clusterRoleBindingList...)
}

func (s *clusterRoleBindingSet) Find(id ezkube.ResourceId) (*rbac_authorization_k8s_io_v1.ClusterRoleBinding, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find ClusterRoleBinding %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&rbac_authorization_k8s_io_v1.ClusterRoleBinding{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding), nil
}

func (s *clusterRoleBindingSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *clusterRoleBindingSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *clusterRoleBindingSet) Delta(newSet ClusterRoleBindingSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *clusterRoleBindingSet) Clone() ClusterRoleBindingSet {
	if s == nil {
		return nil
	}
	return &clusterRoleBindingSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
