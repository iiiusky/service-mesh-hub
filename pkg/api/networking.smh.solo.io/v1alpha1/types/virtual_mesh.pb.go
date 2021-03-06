// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/virtual_mesh.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//If ENABLED, by default disallow traffic to all Services in the VirtualMesh unless explicitly allowed through AccessControlPolicies.
//If DISABLED, by default allow traffic to all Services in the VirtualMesh.
//If MESH_DEFAULT, the default value depends on the type service mesh:
//Istio: false
//Appmesh: true
type VirtualMeshSpec_EnforcementPolicy int32

const (
	VirtualMeshSpec_MESH_DEFAULT VirtualMeshSpec_EnforcementPolicy = 0
	VirtualMeshSpec_ENABLED      VirtualMeshSpec_EnforcementPolicy = 1
	VirtualMeshSpec_DISABLED     VirtualMeshSpec_EnforcementPolicy = 2
)

var VirtualMeshSpec_EnforcementPolicy_name = map[int32]string{
	0: "MESH_DEFAULT",
	1: "ENABLED",
	2: "DISABLED",
}

var VirtualMeshSpec_EnforcementPolicy_value = map[string]int32{
	"MESH_DEFAULT": 0,
	"ENABLED":      1,
	"DISABLED":     2,
}

func (x VirtualMeshSpec_EnforcementPolicy) String() string {
	return proto.EnumName(VirtualMeshSpec_EnforcementPolicy_name, int32(x))
}

func (VirtualMeshSpec_EnforcementPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0}
}

type VirtualMeshSpec_Federation_Mode int32

const (
	// All services in a VirtualMesh will be federated to all workloads in that Virtual Mesh.
	VirtualMeshSpec_Federation_PERMISSIVE VirtualMeshSpec_Federation_Mode = 0
)

var VirtualMeshSpec_Federation_Mode_name = map[int32]string{
	0: "PERMISSIVE",
}

var VirtualMeshSpec_Federation_Mode_value = map[string]int32{
	"PERMISSIVE": 0,
}

func (x VirtualMeshSpec_Federation_Mode) String() string {
	return proto.EnumName(VirtualMeshSpec_Federation_Mode_name, int32(x))
}

func (VirtualMeshSpec_Federation_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 1, 0}
}

type VirtualMeshSpec struct {
	// User-provided display name for the virtual mesh.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The meshes contained in this virtual mesh.
	Meshes               []*types.ResourceRef                  `protobuf:"bytes,2,rep,name=meshes,proto3" json:"meshes,omitempty"`
	CertificateAuthority *VirtualMeshSpec_CertificateAuthority `protobuf:"bytes,3,opt,name=certificate_authority,json=certificateAuthority,proto3" json:"certificate_authority,omitempty"`
	Federation           *VirtualMeshSpec_Federation           `protobuf:"bytes,4,opt,name=federation,proto3" json:"federation,omitempty"`
	// Types that are valid to be assigned to TrustModel:
	//	*VirtualMeshSpec_Shared
	//	*VirtualMeshSpec_Limited
	TrustModel           isVirtualMeshSpec_TrustModel      `protobuf_oneof:"trust_model"`
	EnforceAccessControl VirtualMeshSpec_EnforcementPolicy `protobuf:"varint,7,opt,name=enforce_access_control,json=enforceAccessControl,proto3,enum=networking.smh.solo.io.VirtualMeshSpec_EnforcementPolicy" json:"enforce_access_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *VirtualMeshSpec) Reset()         { *m = VirtualMeshSpec{} }
func (m *VirtualMeshSpec) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec) ProtoMessage()    {}
func (*VirtualMeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0}
}
func (m *VirtualMeshSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec.Unmarshal(m, b)
}
func (m *VirtualMeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec.Merge(m, src)
}
func (m *VirtualMeshSpec) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec.Size(m)
}
func (m *VirtualMeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec proto.InternalMessageInfo

type isVirtualMeshSpec_TrustModel interface {
	isVirtualMeshSpec_TrustModel()
	Equal(interface{}) bool
}

type VirtualMeshSpec_Shared struct {
	Shared *VirtualMeshSpec_SharedTrust `protobuf:"bytes,5,opt,name=shared,proto3,oneof" json:"shared,omitempty"`
}
type VirtualMeshSpec_Limited struct {
	Limited *VirtualMeshSpec_LimitedTrust `protobuf:"bytes,6,opt,name=limited,proto3,oneof" json:"limited,omitempty"`
}

func (*VirtualMeshSpec_Shared) isVirtualMeshSpec_TrustModel()  {}
func (*VirtualMeshSpec_Limited) isVirtualMeshSpec_TrustModel() {}

func (m *VirtualMeshSpec) GetTrustModel() isVirtualMeshSpec_TrustModel {
	if m != nil {
		return m.TrustModel
	}
	return nil
}

func (m *VirtualMeshSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualMeshSpec) GetMeshes() []*types.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *VirtualMeshSpec) GetCertificateAuthority() *VirtualMeshSpec_CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func (m *VirtualMeshSpec) GetFederation() *VirtualMeshSpec_Federation {
	if m != nil {
		return m.Federation
	}
	return nil
}

func (m *VirtualMeshSpec) GetShared() *VirtualMeshSpec_SharedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Shared); ok {
		return x.Shared
	}
	return nil
}

func (m *VirtualMeshSpec) GetLimited() *VirtualMeshSpec_LimitedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Limited); ok {
		return x.Limited
	}
	return nil
}

func (m *VirtualMeshSpec) GetEnforceAccessControl() VirtualMeshSpec_EnforcementPolicy {
	if m != nil {
		return m.EnforceAccessControl
	}
	return VirtualMeshSpec_MESH_DEFAULT
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_Shared)(nil),
		(*VirtualMeshSpec_Limited)(nil),
	}
}

type VirtualMeshSpec_CertificateAuthority struct {
	// If omitted, defaults to builtin.
	//
	// Types that are valid to be assigned to Type:
	//	*VirtualMeshSpec_CertificateAuthority_Builtin_
	//	*VirtualMeshSpec_CertificateAuthority_Provided_
	Type                 isVirtualMeshSpec_CertificateAuthority_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority) Reset()         { *m = VirtualMeshSpec_CertificateAuthority{} }
func (m *VirtualMeshSpec_CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_CertificateAuthority) ProtoMessage()    {}
func (*VirtualMeshSpec_CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority proto.InternalMessageInfo

type isVirtualMeshSpec_CertificateAuthority_Type interface {
	isVirtualMeshSpec_CertificateAuthority_Type()
	Equal(interface{}) bool
}

type VirtualMeshSpec_CertificateAuthority_Builtin_ struct {
	Builtin *VirtualMeshSpec_CertificateAuthority_Builtin `protobuf:"bytes,1,opt,name=builtin,proto3,oneof" json:"builtin,omitempty"`
}
type VirtualMeshSpec_CertificateAuthority_Provided_ struct {
	Provided *VirtualMeshSpec_CertificateAuthority_Provided `protobuf:"bytes,2,opt,name=provided,proto3,oneof" json:"provided,omitempty"`
}

func (*VirtualMeshSpec_CertificateAuthority_Builtin_) isVirtualMeshSpec_CertificateAuthority_Type() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided_) isVirtualMeshSpec_CertificateAuthority_Type() {
}

func (m *VirtualMeshSpec_CertificateAuthority) GetType() isVirtualMeshSpec_CertificateAuthority_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetBuiltin() *VirtualMeshSpec_CertificateAuthority_Builtin {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Builtin_); ok {
		return x.Builtin
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetProvided() *VirtualMeshSpec_CertificateAuthority_Provided {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Provided_); ok {
		return x.Provided
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec_CertificateAuthority) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_CertificateAuthority_Builtin_)(nil),
		(*VirtualMeshSpec_CertificateAuthority_Provided_)(nil),
	}
}

//
//Configuration for auto-generated root certificate unique to the VirtualMesh
//Uses the X.509 format, RFC5280
type VirtualMeshSpec_CertificateAuthority_Builtin struct {
	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "service-mesh-hub"
	OrgName              string   `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Builtin{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetTtlDays() uint32 {
	if m != nil {
		return m.TtlDays
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetRsaKeySizeBytes() uint32 {
	if m != nil {
		return m.RsaKeySizeBytes
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

// Configuration for user-provided root certificate.
type VirtualMeshSpec_CertificateAuthority_Provided struct {
	// Reference to a Secret object containing the root certificate.
	Certificate          *types.ResourceRef `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Provided) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Provided{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Provided) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 0, 1}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Provided) GetCertificate() *types.ResourceRef {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type VirtualMeshSpec_Federation struct {
	Mode                 VirtualMeshSpec_Federation_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=networking.smh.solo.io.VirtualMeshSpec_Federation_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *VirtualMeshSpec_Federation) Reset()         { *m = VirtualMeshSpec_Federation{} }
func (m *VirtualMeshSpec_Federation) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_Federation) ProtoMessage()    {}
func (*VirtualMeshSpec_Federation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 1}
}
func (m *VirtualMeshSpec_Federation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_Federation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_Federation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_Federation.Merge(m, src)
}
func (m *VirtualMeshSpec_Federation) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Size(m)
}
func (m *VirtualMeshSpec_Federation) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_Federation.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_Federation proto.InternalMessageInfo

func (m *VirtualMeshSpec_Federation) GetMode() VirtualMeshSpec_Federation_Mode {
	if m != nil {
		return m.Mode
	}
	return VirtualMeshSpec_Federation_PERMISSIVE
}

//
//Shared trust is a virtual mesh trust model requiring a shared root certificate, as well as shared identity
//between all entities which wish to communicate within the virtual mesh.
//
//The best current example of this would be the replicated control planes example from Istio:
//https://preliminary.istio.io/docs/setup/install/multicluster/gateways/
type VirtualMeshSpec_SharedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_SharedTrust) Reset()         { *m = VirtualMeshSpec_SharedTrust{} }
func (m *VirtualMeshSpec_SharedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_SharedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_SharedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 2}
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Size(m)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_SharedTrust proto.InternalMessageInfo

//
//Limited trust is a virtual mesh trust model which does not require all meshes sharing the same root certificate
//or identity model. But rather, the limited trust creates trust between meshes running on different clusters
//by connecting their ingress/egress gateways with a common cert/identity. In this model all requests
//between different have the following request path when communicating between clusters
//
//cluster 1 MTLS               shared MTLS                  cluster 2 MTLS
//client/workload <-----------> egress gateway <----------> ingress gateway <--------------> server
//
//This approach has the downside of not maintaining identity from client to server, but allows for ad-hoc
//addition of additional clusters into a virtual mesh.
type VirtualMeshSpec_LimitedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_LimitedTrust) Reset()         { *m = VirtualMeshSpec_LimitedTrust{} }
func (m *VirtualMeshSpec_LimitedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_LimitedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_LimitedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{0, 3}
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Size(m)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_LimitedTrust proto.InternalMessageInfo

type VirtualMeshStatus struct {
	// Status of the process writing federation decision metadata onto MeshServices.
	FederationStatus *types.Status `protobuf:"bytes,1,opt,name=federation_status,json=federationStatus,proto3" json:"federation_status,omitempty"`
	// Status of the process signing CSRs.
	CertificateStatus *types.Status `protobuf:"bytes,2,opt,name=certificate_status,json=certificateStatus,proto3" json:"certificate_status,omitempty"`
	// Overall validation status of this VirtualMesh.
	ConfigStatus *types.Status `protobuf:"bytes,3,opt,name=config_status,json=configStatus,proto3" json:"config_status,omitempty"`
	// Status of ensuring that access control is enforced within this VirtualMesh.
	AccessControlEnforcementStatus *types.Status `protobuf:"bytes,4,opt,name=access_control_enforcement_status,json=accessControlEnforcementStatus,proto3" json:"access_control_enforcement_status,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}      `json:"-"`
	XXX_unrecognized               []byte        `json:"-"`
	XXX_sizecache                  int32         `json:"-"`
}

func (m *VirtualMeshStatus) Reset()         { *m = VirtualMeshStatus{} }
func (m *VirtualMeshStatus) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshStatus) ProtoMessage()    {}
func (*VirtualMeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c28e03fd4cc9e166, []int{1}
}
func (m *VirtualMeshStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshStatus.Unmarshal(m, b)
}
func (m *VirtualMeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshStatus.Marshal(b, m, deterministic)
}
func (m *VirtualMeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshStatus.Merge(m, src)
}
func (m *VirtualMeshStatus) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshStatus.Size(m)
}
func (m *VirtualMeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshStatus proto.InternalMessageInfo

func (m *VirtualMeshStatus) GetFederationStatus() *types.Status {
	if m != nil {
		return m.FederationStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetCertificateStatus() *types.Status {
	if m != nil {
		return m.CertificateStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetConfigStatus() *types.Status {
	if m != nil {
		return m.ConfigStatus
	}
	return nil
}

func (m *VirtualMeshStatus) GetAccessControlEnforcementStatus() *types.Status {
	if m != nil {
		return m.AccessControlEnforcementStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("networking.smh.solo.io.VirtualMeshSpec_EnforcementPolicy", VirtualMeshSpec_EnforcementPolicy_name, VirtualMeshSpec_EnforcementPolicy_value)
	proto.RegisterEnum("networking.smh.solo.io.VirtualMeshSpec_Federation_Mode", VirtualMeshSpec_Federation_Mode_name, VirtualMeshSpec_Federation_Mode_value)
	proto.RegisterType((*VirtualMeshSpec)(nil), "networking.smh.solo.io.VirtualMeshSpec")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Builtin)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority.Builtin")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Provided)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority.Provided")
	proto.RegisterType((*VirtualMeshSpec_Federation)(nil), "networking.smh.solo.io.VirtualMeshSpec.Federation")
	proto.RegisterType((*VirtualMeshSpec_SharedTrust)(nil), "networking.smh.solo.io.VirtualMeshSpec.SharedTrust")
	proto.RegisterType((*VirtualMeshSpec_LimitedTrust)(nil), "networking.smh.solo.io.VirtualMeshSpec.LimitedTrust")
	proto.RegisterType((*VirtualMeshStatus)(nil), "networking.smh.solo.io.VirtualMeshStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/virtual_mesh.proto", fileDescriptor_c28e03fd4cc9e166)
}

var fileDescriptor_c28e03fd4cc9e166 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0xe3, 0x8f, 0xd9, 0xe9, 0xb1, 0x93, 0x3a, 0x44, 0x16, 0x68, 0x06, 0x16, 0xa4, 0xb9,
	0x0a, 0xb0, 0x45, 0x42, 0xdd, 0x0d, 0xc3, 0x80, 0x0d, 0x5b, 0x5c, 0xab, 0x4b, 0x90, 0x38, 0xc8,
	0xe4, 0xae, 0x03, 0x76, 0xa3, 0xd1, 0xd4, 0xb1, 0x44, 0x44, 0x12, 0x55, 0x92, 0x4a, 0xa1, 0x3e,
	0xcd, 0x2e, 0xf7, 0x26, 0x03, 0xf6, 0x18, 0x7b, 0x92, 0x41, 0x1f, 0xae, 0xb4, 0x36, 0x40, 0x3c,
	0xf4, 0x4e, 0xa4, 0xce, 0xff, 0x77, 0x0e, 0xc9, 0xf3, 0x27, 0xe1, 0xda, 0xe7, 0x3a, 0x48, 0x97,
	0x26, 0x13, 0x91, 0xa5, 0x44, 0x28, 0x4e, 0xb9, 0xb0, 0x14, 0xca, 0x3b, 0xce, 0xf0, 0x34, 0x42,
	0x15, 0x9c, 0x06, 0xe9, 0xd2, 0xa2, 0x09, 0xb7, 0x62, 0xd4, 0x6f, 0x84, 0xbc, 0xe5, 0xb1, 0x6f,
	0xdd, 0x3d, 0xa5, 0x61, 0x12, 0xd0, 0xa7, 0xd6, 0x1d, 0x97, 0x3a, 0xa5, 0xa1, 0x9b, 0x07, 0x9a,
	0x89, 0x14, 0x5a, 0x90, 0x83, 0x3a, 0xce, 0x54, 0x51, 0x60, 0xe6, 0x4c, 0x93, 0x8b, 0xf1, 0x97,
	0xf7, 0x42, 0x99, 0x90, 0x58, 0xe3, 0x24, 0xae, 0x4a, 0xca, 0xd8, 0xda, 0x20, 0x5a, 0x69, 0xaa,
	0x53, 0x55, 0x09, 0x0e, 0x7d, 0x21, 0xfc, 0x10, 0xad, 0x62, 0xb4, 0x4c, 0x57, 0xd6, 0x1b, 0x49,
	0x93, 0x04, 0xe5, 0xfa, 0xff, 0xbe, 0x2f, 0x7c, 0x51, 0x7c, 0x5a, 0xf9, 0x57, 0x39, 0x7b, 0xfc,
	0xf7, 0x23, 0x78, 0xfc, 0xaa, 0x5c, 0xc3, 0x1c, 0x55, 0xb0, 0x48, 0x90, 0x91, 0x27, 0x30, 0xf4,
	0xb8, 0x4a, 0x42, 0x9a, 0xb9, 0x31, 0x8d, 0xd0, 0x68, 0x1d, 0xb5, 0x4e, 0x1e, 0x39, 0x83, 0x6a,
	0xee, 0x9a, 0x46, 0x48, 0xbe, 0x86, 0x5e, 0x5e, 0x17, 0x2a, 0xa3, 0x7d, 0xd4, 0x39, 0x19, 0x4c,
	0x3e, 0x37, 0xf3, 0xca, 0x9a, 0xcb, 0x35, 0x1d, 0x54, 0x22, 0x95, 0x0c, 0x1d, 0x5c, 0x39, 0x55,
	0x30, 0x79, 0x0d, 0x9f, 0x32, 0x94, 0x9a, 0xaf, 0x38, 0xa3, 0x1a, 0x5d, 0x9a, 0xea, 0x40, 0x48,
	0xae, 0x33, 0xa3, 0x73, 0xd4, 0x3a, 0x19, 0x4c, 0xbe, 0x33, 0xef, 0xdf, 0x3a, 0xf3, 0xbd, 0x0a,
	0xcd, 0xe7, 0x35, 0xe4, 0x6c, 0xcd, 0x70, 0xf6, 0xd9, 0x3d, 0xb3, 0xc4, 0x01, 0x58, 0xa1, 0x87,
	0x92, 0x6a, 0x2e, 0x62, 0xa3, 0x5b, 0xe4, 0x99, 0x6c, 0x9a, 0xe7, 0xc5, 0x3b, 0xa5, 0xd3, 0xa0,
	0x90, 0x39, 0xf4, 0x54, 0x40, 0x25, 0x7a, 0xc6, 0x27, 0x05, 0xef, 0xd9, 0xa6, 0xbc, 0x45, 0xa1,
	0x7a, 0x29, 0x53, 0xa5, 0xcf, 0xb7, 0x9c, 0x0a, 0x42, 0x6e, 0xa0, 0x1f, 0xf2, 0x88, 0x6b, 0xf4,
	0x8c, 0x5e, 0xc1, 0xfb, 0x6a, 0x53, 0xde, 0x55, 0x29, 0x5b, 0x03, 0xd7, 0x18, 0x22, 0xe0, 0x00,
	0xe3, 0x95, 0x90, 0x0c, 0x5d, 0xca, 0x18, 0x2a, 0xe5, 0x32, 0x11, 0x6b, 0x29, 0x42, 0xa3, 0x7f,
	0xd4, 0x3a, 0xd9, 0x9d, 0x7c, 0xbb, 0x69, 0x02, 0xbb, 0xa4, 0x44, 0x18, 0xeb, 0x1b, 0x11, 0x72,
	0x96, 0x39, 0xfb, 0x15, 0xf8, 0xac, 0xe0, 0x3e, 0x2f, 0xb1, 0xe3, 0x3f, 0x3a, 0xb0, 0x7f, 0xdf,
	0xa1, 0x90, 0xdf, 0xa1, 0xbf, 0x4c, 0x79, 0xa8, 0x79, 0x5c, 0xb4, 0xd1, 0x60, 0x32, 0xfb, 0x98,
	0x33, 0x36, 0xa7, 0x25, 0x2b, 0x5f, 0x6b, 0x85, 0x25, 0x0c, 0xb6, 0x13, 0x29, 0xee, 0xb8, 0x87,
	0x9e, 0xd1, 0x2e, 0x52, 0xd8, 0x1f, 0x95, 0xe2, 0xa6, 0x82, 0x9d, 0x6f, 0x39, 0xef, 0xc0, 0xe3,
	0x10, 0xfa, 0x55, 0x6a, 0xf2, 0x19, 0x6c, 0x6b, 0x1d, 0xba, 0x1e, 0xcd, 0x54, 0xb1, 0xa4, 0x1d,
	0xa7, 0xaf, 0x75, 0x38, 0xa3, 0x99, 0x22, 0x5f, 0x00, 0x91, 0x8a, 0xba, 0xb7, 0x98, 0xb9, 0x8a,
	0xbf, 0x45, 0x77, 0x99, 0xe9, 0xc2, 0x21, 0x79, 0xd0, 0x63, 0xa9, 0xe8, 0x25, 0x66, 0x0b, 0xfe,
	0x16, 0xa7, 0xf9, 0x74, 0xce, 0x11, 0xd2, 0x2f, 0x1d, 0xd6, 0x29, 0x1c, 0xd6, 0x17, 0xd2, 0xcf,
	0xdd, 0x35, 0xbe, 0x84, 0xed, 0x75, 0x15, 0xe4, 0x07, 0x18, 0x34, 0xfa, 0xba, 0x32, 0xca, 0x03,
	0x76, 0x6b, 0x2a, 0xa6, 0x3d, 0xe8, 0xea, 0x2c, 0xc1, 0xf1, 0x6b, 0x80, 0xba, 0x9d, 0xc9, 0x25,
	0x74, 0x23, 0xe1, 0x95, 0xde, 0xde, 0x9d, 0x7c, 0xf3, 0xff, 0x0d, 0x61, 0xce, 0x85, 0x87, 0x4e,
	0x01, 0x39, 0x3e, 0x80, 0x6e, 0x3e, 0x22, 0xbb, 0x00, 0x37, 0xb6, 0x33, 0xbf, 0x58, 0x2c, 0x2e,
	0x5e, 0xd9, 0xa3, 0xad, 0xf1, 0x0e, 0x0c, 0x1a, 0x1d, 0x3f, 0xde, 0x85, 0x61, 0xb3, 0x61, 0x8f,
	0x7f, 0x84, 0xbd, 0x0f, 0xfa, 0x8b, 0x8c, 0x60, 0x38, 0xb7, 0x17, 0xe7, 0xee, 0xcc, 0x7e, 0x71,
	0xf6, 0xcb, 0xd5, 0xcb, 0xd1, 0x16, 0x19, 0x40, 0xdf, 0xbe, 0x3e, 0x9b, 0x5e, 0xd9, 0xb3, 0x51,
	0x8b, 0x0c, 0x61, 0x7b, 0x76, 0xb1, 0x28, 0x47, 0xed, 0xe9, 0x0e, 0x0c, 0x74, 0x8e, 0x72, 0xf3,
	0x32, 0xc2, 0xe3, 0xbf, 0xda, 0xb0, 0xd7, 0xac, 0xb8, 0xb8, 0x1e, 0x89, 0x0d, 0x7b, 0xb5, 0x77,
	0xdd, 0xf2, 0xce, 0xac, 0x9a, 0xd1, 0xf8, 0x70, 0x1f, 0x4b, 0x91, 0x33, 0xaa, 0x25, 0x15, 0xe6,
	0x27, 0x20, 0xcd, 0xbb, 0xab, 0xe2, 0xb4, 0x1f, 0xe0, 0xec, 0x35, 0x34, 0x15, 0xe8, 0x7b, 0xd8,
	0x61, 0x22, 0x5e, 0x71, 0x7f, 0xcd, 0xe8, 0x3c, 0xc0, 0x18, 0x96, 0xe1, 0x95, 0x9c, 0xc1, 0x93,
	0xff, 0x7a, 0xda, 0xc5, 0x7a, 0x13, 0xd7, 0xc8, 0xee, 0x03, 0xc8, 0x43, 0xda, 0xb4, 0x6f, 0xe3,
	0x14, 0xca, 0xff, 0xd3, 0x5f, 0xff, 0xfc, 0xe7, 0xb0, 0xf5, 0xdb, 0xcf, 0x9b, 0xbc, 0x8c, 0xc9,
	0xad, 0xff, 0xde, 0xeb, 0xd8, 0x4c, 0x58, 0x3f, 0x56, 0x79, 0x13, 0xaa, 0x65, 0xaf, 0x78, 0x76,
	0x9e, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xce, 0x97, 0x1a, 0x4f, 0x75, 0x07, 0x00, 0x00,
}

func (this *VirtualMeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec)
	if !ok {
		that2, ok := that.(VirtualMeshSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !this.CertificateAuthority.Equal(that1.CertificateAuthority) {
		return false
	}
	if !this.Federation.Equal(that1.Federation) {
		return false
	}
	if that1.TrustModel == nil {
		if this.TrustModel != nil {
			return false
		}
	} else if this.TrustModel == nil {
		return false
	} else if !this.TrustModel.Equal(that1.TrustModel) {
		return false
	}
	if this.EnforceAccessControl != that1.EnforceAccessControl {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Shared) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Shared)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Shared)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Shared.Equal(that1.Shared) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Limited) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Limited)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Limited)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Limited.Equal(that1.Limited) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Builtin.Equal(that1.Builtin) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Provided.Equal(that1.Provided) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TtlDays != that1.TtlDays {
		return false
	}
	if this.RsaKeySizeBytes != that1.RsaKeySizeBytes {
		return false
	}
	if this.OrgName != that1.OrgName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Certificate.Equal(that1.Certificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Federation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Federation)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Federation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Mode != that1.Mode {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_SharedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_SharedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_SharedTrust)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_LimitedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_LimitedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_LimitedTrust)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshStatus)
	if !ok {
		that2, ok := that.(VirtualMeshStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.FederationStatus.Equal(that1.FederationStatus) {
		return false
	}
	if !this.CertificateStatus.Equal(that1.CertificateStatus) {
		return false
	}
	if !this.ConfigStatus.Equal(that1.ConfigStatus) {
		return false
	}
	if !this.AccessControlEnforcementStatus.Equal(that1.AccessControlEnforcementStatus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
