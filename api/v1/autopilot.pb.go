// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilot.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The AutoPilotProject file is the root configuration file for the project itself.
//
// This file will be used to build and deploy the autopilot operator.
// It is loaded automatically by the autopilot CLI. Its
// default location is 'autopilot.yaml'
type AutoPilotProject struct {
	// the name (kubernetes Kind) of the top-level
	// CRD for the operator
	// Specified via the `ap init <Kind>` command
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// the ApiVersion of the top-level
	// CRD for the operator
	ApiVersion string `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	// the name of the Operator
	// this is used to name and label loggers, k8s resources, and metrics exposed
	// by the operator. Should be [valid Kube resource names](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	OperatorName string `protobuf:"bytes,3,opt,name=operatorName,proto3" json:"operatorName,omitempty"`
	// Each phase represents a different
	// stage in the lifecycle of the CRD (e.g. Pending/Succeeded/Failed).
	//
	// Each phase specifies a unique name
	// and its own set of inputs and outputs.
	Phases []*Phase `protobuf:"bytes,4,rep,name=phases,proto3" json:"phases,omitempty"`
	// enable use of a Finalizer to handle object deletion
	EnableFinalizer bool `protobuf:"varint,5,opt,name=enableFinalizer,proto3" json:"enableFinalizer,omitempty"`
	// custom Parameters which extend AutoPilot's builtin types
	CustomParameters []*Parameter `protobuf:"bytes,6,rep,name=customParameters,proto3" json:"customParameters,omitempty"`
	// custom Queries which extend AutoPilot's metrics queries
	Queries              []*MetricsQuery `protobuf:"bytes,7,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AutoPilotProject) Reset()         { *m = AutoPilotProject{} }
func (m *AutoPilotProject) String() string { return proto.CompactTextString(m) }
func (*AutoPilotProject) ProtoMessage()    {}
func (*AutoPilotProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{0}
}

func (m *AutoPilotProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoPilotProject.Unmarshal(m, b)
}
func (m *AutoPilotProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoPilotProject.Marshal(b, m, deterministic)
}
func (m *AutoPilotProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoPilotProject.Merge(m, src)
}
func (m *AutoPilotProject) XXX_Size() int {
	return xxx_messageInfo_AutoPilotProject.Size(m)
}
func (m *AutoPilotProject) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoPilotProject.DiscardUnknown(m)
}

var xxx_messageInfo_AutoPilotProject proto.InternalMessageInfo

func (m *AutoPilotProject) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *AutoPilotProject) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *AutoPilotProject) GetOperatorName() string {
	if m != nil {
		return m.OperatorName
	}
	return ""
}

func (m *AutoPilotProject) GetPhases() []*Phase {
	if m != nil {
		return m.Phases
	}
	return nil
}

func (m *AutoPilotProject) GetEnableFinalizer() bool {
	if m != nil {
		return m.EnableFinalizer
	}
	return false
}

func (m *AutoPilotProject) GetCustomParameters() []*Parameter {
	if m != nil {
		return m.CustomParameters
	}
	return nil
}

func (m *AutoPilotProject) GetQueries() []*MetricsQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// MeshProviders provide an interface to monitoring and managing a specific
// mesh.
//
// AutoPilot does not abstract the mesh API - AutoPilot developers must
// still reason able about Provider-specific CRDs. AutoPilot's job is to
// abstract operational concerns such as discovering control plane configuration
// and monitoring metrics.
type Phase struct {
	// name of the phase. must be unique
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// description of the phase. used for comments and docs
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// indicates whether this is the initial phase of the system.
	// exactly one phase must be the initial phase
	Initial bool `protobuf:"varint,3,opt,name=initial,proto3" json:"initial,omitempty"`
	// indicates whether this is a "final" or "resting" phase of the system.
	// when the CRD is in the final phase, no more processing will be done on it
	Final bool `protobuf:"varint,4,opt,name=final,proto3" json:"final,omitempty"`
	// the set of inputs for this phase
	// the inputs will be retrieved by the scheduler
	// and passed to the worker as input parameters
	//
	// custom inputs can be defined in the
	// autopilot.yaml
	Inputs []string `protobuf:"bytes,5,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// the set of outputs for this phase
	// the inputs will be propagated to k8s storage (etcd) by the scheduler.
	//
	// custom outputs can be defined in the
	// autopilot.yaml
	Outputs              []string `protobuf:"bytes,6,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phase) Reset()         { *m = Phase{} }
func (m *Phase) String() string { return proto.CompactTextString(m) }
func (*Phase) ProtoMessage()    {}
func (*Phase) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{1}
}

func (m *Phase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phase.Unmarshal(m, b)
}
func (m *Phase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phase.Marshal(b, m, deterministic)
}
func (m *Phase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phase.Merge(m, src)
}
func (m *Phase) XXX_Size() int {
	return xxx_messageInfo_Phase.Size(m)
}
func (m *Phase) XXX_DiscardUnknown() {
	xxx_messageInfo_Phase.DiscardUnknown(m)
}

var xxx_messageInfo_Phase proto.InternalMessageInfo

func (m *Phase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Phase) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Phase) GetInitial() bool {
	if m != nil {
		return m.Initial
	}
	return false
}

func (m *Phase) GetFinal() bool {
	if m != nil {
		return m.Final
	}
	return false
}

func (m *Phase) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Phase) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Custom Parameters allow code to be generated
// for inputs/outputs that are not built-in to AutoPilot.
// These types must be Kubernetes-compatible Go structs.
type Parameter struct {
	// the fully lower-case name of this resource
	// e.g. "pods", "services", "replicasets", "configmaps"
	LowerName string `protobuf:"bytes,1,opt,name=lowerName,proto3" json:"lowerName,omitempty"`
	// the singular CamelCased name of the resource
	// equivalent to Kind
	SingleName string `protobuf:"bytes,2,opt,name=singleName,proto3" json:"singleName,omitempty"`
	// the plural CamelCased name of the resource
	// equivalent to the pluralized form of Kind
	PluralName string `protobuf:"bytes,3,opt,name=pluralName,proto3" json:"pluralName,omitempty"`
	// import prefix used by generated code
	ImportPrefix string `protobuf:"bytes,4,opt,name=importPrefix,proto3" json:"importPrefix,omitempty"`
	// go package (import path) to the go struct for the resource
	Package string `protobuf:"bytes,5,opt,name=package,proto3" json:"package,omitempty"`
	// Kubernetes API group for the resource
	// e.g. "networking.istio.io"
	ApiGroup string `protobuf:"bytes,6,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	// indicates whether the resource is a CRD
	// if true, the Resource will be added to the operator's runtime.Scheme
	IsCrd                bool     `protobuf:"varint,7,opt,name=isCrd,proto3" json:"isCrd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{2}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetLowerName() string {
	if m != nil {
		return m.LowerName
	}
	return ""
}

func (m *Parameter) GetSingleName() string {
	if m != nil {
		return m.SingleName
	}
	return ""
}

func (m *Parameter) GetPluralName() string {
	if m != nil {
		return m.PluralName
	}
	return ""
}

func (m *Parameter) GetImportPrefix() string {
	if m != nil {
		return m.ImportPrefix
	}
	return ""
}

func (m *Parameter) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Parameter) GetApiGroup() string {
	if m != nil {
		return m.ApiGroup
	}
	return ""
}

func (m *Parameter) GetIsCrd() bool {
	if m != nil {
		return m.IsCrd
	}
	return false
}

type MetricsQuery struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	QueryTemplate        string   `protobuf:"bytes,2,opt,name=queryTemplate,proto3" json:"queryTemplate,omitempty"`
	Parameters           []string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsQuery) Reset()         { *m = MetricsQuery{} }
func (m *MetricsQuery) String() string { return proto.CompactTextString(m) }
func (*MetricsQuery) ProtoMessage()    {}
func (*MetricsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c7e86e2b87635e, []int{3}
}

func (m *MetricsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsQuery.Unmarshal(m, b)
}
func (m *MetricsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsQuery.Marshal(b, m, deterministic)
}
func (m *MetricsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsQuery.Merge(m, src)
}
func (m *MetricsQuery) XXX_Size() int {
	return xxx_messageInfo_MetricsQuery.Size(m)
}
func (m *MetricsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsQuery proto.InternalMessageInfo

func (m *MetricsQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsQuery) GetQueryTemplate() string {
	if m != nil {
		return m.QueryTemplate
	}
	return ""
}

func (m *MetricsQuery) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*AutoPilotProject)(nil), "autopilot.AutoPilotProject")
	proto.RegisterType((*Phase)(nil), "autopilot.Phase")
	proto.RegisterType((*Parameter)(nil), "autopilot.Parameter")
	proto.RegisterType((*MetricsQuery)(nil), "autopilot.MetricsQuery")
}

func init() { proto.RegisterFile("autopilot.proto", fileDescriptor_f7c7e86e2b87635e) }

var fileDescriptor_f7c7e86e2b87635e = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xdb, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xb7, 0x6d, 0xda, 0xcc, 0x2e, 0xda, 0xca, 0x5a, 0x81, 0x85, 0x10, 0xaa, 0x02, 0x48,
	0x79, 0xa1, 0xd5, 0xc2, 0x0f, 0x70, 0x91, 0xe0, 0x09, 0x14, 0x22, 0xc4, 0x03, 0x6f, 0x6e, 0x3a,
	0xdb, 0x0e, 0xeb, 0xc4, 0xc6, 0x17, 0x60, 0xf9, 0x19, 0x3e, 0x82, 0xaf, 0xe1, 0x6f, 0x90, 0xdd,
	0xa6, 0x49, 0x61, 0xdf, 0x7c, 0xce, 0x19, 0x8f, 0x66, 0xce, 0xb1, 0xe1, 0x5c, 0x78, 0xa7, 0x34,
	0x49, 0xe5, 0x16, 0xda, 0x28, 0xa7, 0x58, 0x7a, 0x20, 0xb2, 0xdf, 0x27, 0x30, 0x7b, 0xe9, 0x9d,
	0x2a, 0x02, 0x2a, 0x8c, 0xfa, 0x82, 0x95, 0x63, 0x0c, 0x46, 0xd7, 0xd4, 0xac, 0xf9, 0x60, 0x3e,
	0xc8, 0xd3, 0x32, 0x9e, 0xd9, 0x43, 0x00, 0xa1, 0xe9, 0x13, 0x1a, 0x4b, 0xaa, 0xe1, 0x27, 0x51,
	0xe9, 0x31, 0x2c, 0x83, 0x33, 0xa5, 0xd1, 0x08, 0xa7, 0xcc, 0x7b, 0x51, 0x23, 0x1f, 0xc6, 0x8a,
	0x23, 0x8e, 0xe5, 0x90, 0xe8, 0xad, 0xb0, 0x68, 0xf9, 0x68, 0x3e, 0xcc, 0x4f, 0x9f, 0xcd, 0x16,
	0xdd, 0x64, 0x45, 0x10, 0xca, 0xbd, 0xce, 0x72, 0x38, 0xc7, 0x46, 0xac, 0x24, 0xbe, 0xa1, 0x46,
	0x48, 0xfa, 0x89, 0x86, 0x8f, 0xe7, 0x83, 0x7c, 0x5a, 0xfe, 0x4b, 0xb3, 0x17, 0x30, 0xab, 0xbc,
	0x75, 0xaa, 0x2e, 0x84, 0x11, 0x35, 0x3a, 0x34, 0x96, 0x27, 0xb1, 0xfb, 0x45, 0xbf, 0x7b, 0x2b,
	0x96, 0xff, 0x55, 0xb3, 0x4b, 0x98, 0x7c, 0xf5, 0x68, 0x08, 0x2d, 0x9f, 0xc4, 0x8b, 0xf7, 0x7a,
	0x17, 0xdf, 0xa1, 0x33, 0x54, 0xd9, 0x0f, 0x1e, 0xcd, 0x4d, 0xd9, 0xd6, 0x65, 0xbf, 0x06, 0x30,
	0x8e, 0x03, 0x07, 0xab, 0x9a, 0xb0, 0xee, 0xde, 0xaa, 0x70, 0x66, 0x73, 0x38, 0x5d, 0xa3, 0xad,
	0x0c, 0x69, 0xd7, 0x79, 0xd5, 0xa7, 0x18, 0x87, 0x09, 0x35, 0xe4, 0x48, 0xc8, 0xe8, 0xd3, 0xb4,
	0x6c, 0x21, 0xbb, 0x80, 0xf1, 0x55, 0xd8, 0x8d, 0x8f, 0x22, 0xbf, 0x03, 0xec, 0x2e, 0x24, 0xd4,
	0x68, 0xef, 0x2c, 0x1f, 0xcf, 0x87, 0x79, 0x5a, 0xee, 0x51, 0xe8, 0xa3, 0xbc, 0x8b, 0x42, 0x12,
	0x85, 0x16, 0x66, 0x7f, 0x06, 0x90, 0x1e, 0x76, 0x64, 0x0f, 0x20, 0x95, 0xea, 0x3b, 0xee, 0x92,
	0xd9, 0x8d, 0xda, 0x11, 0x21, 0x5a, 0x4b, 0xcd, 0x46, 0x62, 0x94, 0xf7, 0xd1, 0x76, 0x4c, 0xd0,
	0xb5, 0xf4, 0x46, 0xc8, 0x5e, 0xb0, 0x3d, 0x26, 0x44, 0x4f, 0xb5, 0x56, 0xc6, 0x15, 0x06, 0xaf,
	0xe8, 0x47, 0x1c, 0x3d, 0x2d, 0x8f, 0xb8, 0x30, 0xa9, 0x16, 0xd5, 0xb5, 0xd8, 0x60, 0x0c, 0x32,
	0x2d, 0x5b, 0xc8, 0xee, 0xc3, 0x54, 0x68, 0x7a, 0x6b, 0x94, 0xd7, 0x3c, 0x89, 0xd2, 0x01, 0x07,
	0x37, 0xc8, 0xbe, 0x36, 0x6b, 0x3e, 0xd9, 0xb9, 0x11, 0x41, 0xb6, 0x85, 0xb3, 0x7e, 0x2c, 0xb7,
	0x66, 0xf0, 0x18, 0xee, 0x84, 0xb0, 0x6e, 0x3e, 0x62, 0xad, 0xa5, 0x70, 0xed, 0x5a, 0xc7, 0x64,
	0xdc, 0xac, 0x7b, 0x36, 0xc3, 0x68, 0x61, 0x8f, 0x79, 0xf5, 0xe4, 0xf3, 0xa3, 0x0d, 0xb9, 0xad,
	0x5f, 0x2d, 0x2a, 0x55, 0x2f, 0xad, 0x92, 0xea, 0x29, 0xa9, 0xe5, 0xe1, 0x75, 0x2c, 0x85, 0xa6,
	0xe5, 0xb7, 0xcb, 0x55, 0x12, 0xbf, 0xd5, 0xf3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0x50,
	0x88, 0x10, 0x69, 0x03, 0x00, 0x00,
}
