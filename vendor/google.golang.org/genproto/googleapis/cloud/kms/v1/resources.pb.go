// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/kms/v1/resources.proto

package kms // import "google.golang.org/genproto/googleapis/cloud/kms/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [CryptoKeyPurpose][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose] describes the capabilities of a [CryptoKey][google.cloud.kms.v1.CryptoKey]. Two
// keys with the same purpose may use different underlying algorithms, but
// must support the same set of operations.
type CryptoKey_CryptoKeyPurpose int32

const (
	// Not specified.
	CryptoKey_CRYPTO_KEY_PURPOSE_UNSPECIFIED CryptoKey_CryptoKeyPurpose = 0
	// [CryptoKeys][google.cloud.kms.v1.CryptoKey] with this purpose may be used with
	// [Encrypt][google.cloud.kms.v1.KeyManagementService.Encrypt] and
	// [Decrypt][google.cloud.kms.v1.KeyManagementService.Decrypt].
	CryptoKey_ENCRYPT_DECRYPT CryptoKey_CryptoKeyPurpose = 1
)

var CryptoKey_CryptoKeyPurpose_name = map[int32]string{
	0: "CRYPTO_KEY_PURPOSE_UNSPECIFIED",
	1: "ENCRYPT_DECRYPT",
}
var CryptoKey_CryptoKeyPurpose_value = map[string]int32{
	"CRYPTO_KEY_PURPOSE_UNSPECIFIED": 0,
	"ENCRYPT_DECRYPT":                1,
}

func (x CryptoKey_CryptoKeyPurpose) String() string {
	return proto.EnumName(CryptoKey_CryptoKeyPurpose_name, int32(x))
}
func (CryptoKey_CryptoKeyPurpose) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resources_3fd296b2ee28c3bd, []int{1, 0}
}

// The state of a [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion], indicating if it can be used.
type CryptoKeyVersion_CryptoKeyVersionState int32

const (
	// Not specified.
	CryptoKeyVersion_CRYPTO_KEY_VERSION_STATE_UNSPECIFIED CryptoKeyVersion_CryptoKeyVersionState = 0
	// This version may be used in [Encrypt][google.cloud.kms.v1.KeyManagementService.Encrypt] and
	// [Decrypt][google.cloud.kms.v1.KeyManagementService.Decrypt] requests.
	CryptoKeyVersion_ENABLED CryptoKeyVersion_CryptoKeyVersionState = 1
	// This version may not be used, but the key material is still available,
	// and the version can be placed back into the [ENABLED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.ENABLED] state.
	CryptoKeyVersion_DISABLED CryptoKeyVersion_CryptoKeyVersionState = 2
	// This version is destroyed, and the key material is no longer stored.
	// A version may not leave this state once entered.
	CryptoKeyVersion_DESTROYED CryptoKeyVersion_CryptoKeyVersionState = 3
	// This version is scheduled for destruction, and will be destroyed soon.
	// Call
	// [RestoreCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.RestoreCryptoKeyVersion]
	// to put it back into the [DISABLED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DISABLED] state.
	CryptoKeyVersion_DESTROY_SCHEDULED CryptoKeyVersion_CryptoKeyVersionState = 4
)

var CryptoKeyVersion_CryptoKeyVersionState_name = map[int32]string{
	0: "CRYPTO_KEY_VERSION_STATE_UNSPECIFIED",
	1: "ENABLED",
	2: "DISABLED",
	3: "DESTROYED",
	4: "DESTROY_SCHEDULED",
}
var CryptoKeyVersion_CryptoKeyVersionState_value = map[string]int32{
	"CRYPTO_KEY_VERSION_STATE_UNSPECIFIED": 0,
	"ENABLED":           1,
	"DISABLED":          2,
	"DESTROYED":         3,
	"DESTROY_SCHEDULED": 4,
}

func (x CryptoKeyVersion_CryptoKeyVersionState) String() string {
	return proto.EnumName(CryptoKeyVersion_CryptoKeyVersionState_name, int32(x))
}
func (CryptoKeyVersion_CryptoKeyVersionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_resources_3fd296b2ee28c3bd, []int{2, 0}
}

// A [KeyRing][google.cloud.kms.v1.KeyRing] is a toplevel logical grouping of [CryptoKeys][google.cloud.kms.v1.CryptoKey].
type KeyRing struct {
	// Output only. The resource name for the [KeyRing][google.cloud.kms.v1.KeyRing] in the format
	// `projects/*/locations/*/keyRings/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The time at which this [KeyRing][google.cloud.kms.v1.KeyRing] was created.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeyRing) Reset()         { *m = KeyRing{} }
func (m *KeyRing) String() string { return proto.CompactTextString(m) }
func (*KeyRing) ProtoMessage()    {}
func (*KeyRing) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_3fd296b2ee28c3bd, []int{0}
}
func (m *KeyRing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRing.Unmarshal(m, b)
}
func (m *KeyRing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRing.Marshal(b, m, deterministic)
}
func (dst *KeyRing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRing.Merge(dst, src)
}
func (m *KeyRing) XXX_Size() int {
	return xxx_messageInfo_KeyRing.Size(m)
}
func (m *KeyRing) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRing.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRing proto.InternalMessageInfo

func (m *KeyRing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyRing) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// A [CryptoKey][google.cloud.kms.v1.CryptoKey] represents a logical key that can be used for cryptographic
// operations.
//
// A [CryptoKey][google.cloud.kms.v1.CryptoKey] is made up of one or more [versions][google.cloud.kms.v1.CryptoKeyVersion], which
// represent the actual key material used in cryptographic operations.
type CryptoKey struct {
	// Output only. The resource name for this [CryptoKey][google.cloud.kms.v1.CryptoKey] in the format
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. A copy of the "primary" [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] that will be used
	// by [Encrypt][google.cloud.kms.v1.KeyManagementService.Encrypt] when this [CryptoKey][google.cloud.kms.v1.CryptoKey] is given
	// in [EncryptRequest.name][google.cloud.kms.v1.EncryptRequest.name].
	//
	// The [CryptoKey][google.cloud.kms.v1.CryptoKey]'s primary version can be updated via
	// [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion].
	Primary *CryptoKeyVersion `protobuf:"bytes,2,opt,name=primary,proto3" json:"primary,omitempty"`
	// The immutable purpose of this [CryptoKey][google.cloud.kms.v1.CryptoKey]. Currently, the only acceptable
	// purpose is [ENCRYPT_DECRYPT][google.cloud.kms.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	Purpose CryptoKey_CryptoKeyPurpose `protobuf:"varint,3,opt,name=purpose,proto3,enum=google.cloud.kms.v1.CryptoKey_CryptoKeyPurpose" json:"purpose,omitempty"`
	// Output only. The time at which this [CryptoKey][google.cloud.kms.v1.CryptoKey] was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// At [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time], the Key Management Service will automatically:
	//
	// 1. Create a new version of this [CryptoKey][google.cloud.kms.v1.CryptoKey].
	// 2. Mark the new version as primary.
	//
	// Key rotations performed manually via
	// [CreateCryptoKeyVersion][google.cloud.kms.v1.KeyManagementService.CreateCryptoKeyVersion] and
	// [UpdateCryptoKeyPrimaryVersion][google.cloud.kms.v1.KeyManagementService.UpdateCryptoKeyPrimaryVersion]
	// do not affect [next_rotation_time][google.cloud.kms.v1.CryptoKey.next_rotation_time].
	NextRotationTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=next_rotation_time,json=nextRotationTime,proto3" json:"next_rotation_time,omitempty"`
	// Controls the rate of automatic rotation.
	//
	// Types that are valid to be assigned to RotationSchedule:
	//	*CryptoKey_RotationPeriod
	RotationSchedule isCryptoKey_RotationSchedule `protobuf_oneof:"rotation_schedule"`
	// Labels with user-defined metadata. For more information, see
	// [Labeling Keys](/kms/docs/labeling-keys).
	Labels               map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CryptoKey) Reset()         { *m = CryptoKey{} }
func (m *CryptoKey) String() string { return proto.CompactTextString(m) }
func (*CryptoKey) ProtoMessage()    {}
func (*CryptoKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_3fd296b2ee28c3bd, []int{1}
}
func (m *CryptoKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoKey.Unmarshal(m, b)
}
func (m *CryptoKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoKey.Marshal(b, m, deterministic)
}
func (dst *CryptoKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoKey.Merge(dst, src)
}
func (m *CryptoKey) XXX_Size() int {
	return xxx_messageInfo_CryptoKey.Size(m)
}
func (m *CryptoKey) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoKey.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoKey proto.InternalMessageInfo

type isCryptoKey_RotationSchedule interface {
	isCryptoKey_RotationSchedule()
}

type CryptoKey_RotationPeriod struct {
	RotationPeriod *duration.Duration `protobuf:"bytes,8,opt,name=rotation_period,json=rotationPeriod,proto3,oneof"`
}

func (*CryptoKey_RotationPeriod) isCryptoKey_RotationSchedule() {}

func (m *CryptoKey) GetRotationSchedule() isCryptoKey_RotationSchedule {
	if m != nil {
		return m.RotationSchedule
	}
	return nil
}

func (m *CryptoKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CryptoKey) GetPrimary() *CryptoKeyVersion {
	if m != nil {
		return m.Primary
	}
	return nil
}

func (m *CryptoKey) GetPurpose() CryptoKey_CryptoKeyPurpose {
	if m != nil {
		return m.Purpose
	}
	return CryptoKey_CRYPTO_KEY_PURPOSE_UNSPECIFIED
}

func (m *CryptoKey) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *CryptoKey) GetNextRotationTime() *timestamp.Timestamp {
	if m != nil {
		return m.NextRotationTime
	}
	return nil
}

func (m *CryptoKey) GetRotationPeriod() *duration.Duration {
	if x, ok := m.GetRotationSchedule().(*CryptoKey_RotationPeriod); ok {
		return x.RotationPeriod
	}
	return nil
}

func (m *CryptoKey) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CryptoKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CryptoKey_OneofMarshaler, _CryptoKey_OneofUnmarshaler, _CryptoKey_OneofSizer, []interface{}{
		(*CryptoKey_RotationPeriod)(nil),
	}
}

func _CryptoKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CryptoKey)
	// rotation_schedule
	switch x := m.RotationSchedule.(type) {
	case *CryptoKey_RotationPeriod:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RotationPeriod); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CryptoKey.RotationSchedule has unexpected type %T", x)
	}
	return nil
}

func _CryptoKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CryptoKey)
	switch tag {
	case 8: // rotation_schedule.rotation_period
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(duration.Duration)
		err := b.DecodeMessage(msg)
		m.RotationSchedule = &CryptoKey_RotationPeriod{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CryptoKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CryptoKey)
	// rotation_schedule
	switch x := m.RotationSchedule.(type) {
	case *CryptoKey_RotationPeriod:
		s := proto.Size(x.RotationPeriod)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] represents an individual cryptographic key, and the
// associated key material.
//
// It can be used for cryptographic operations either directly, or via its
// parent [CryptoKey][google.cloud.kms.v1.CryptoKey], in which case the server will choose the appropriate
// version for the operation.
//
// For security reasons, the raw cryptographic key material represented by a
// [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] can never be viewed or exported. It can only be used to
// encrypt or decrypt data when an authorized user or application invokes Cloud
// KMS.
type CryptoKeyVersion struct {
	// Output only. The resource name for this [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] in the format
	// `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The current state of the [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion].
	State CryptoKeyVersion_CryptoKeyVersionState `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.kms.v1.CryptoKeyVersion_CryptoKeyVersionState" json:"state,omitempty"`
	// Output only. The time at which this [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time this [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion]'s key material is scheduled
	// for destruction. Only present if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	// [DESTROY_SCHEDULED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROY_SCHEDULED].
	DestroyTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=destroy_time,json=destroyTime,proto3" json:"destroy_time,omitempty"`
	// Output only. The time this CryptoKeyVersion's key material was
	// destroyed. Only present if [state][google.cloud.kms.v1.CryptoKeyVersion.state] is
	// [DESTROYED][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionState.DESTROYED].
	DestroyEventTime     *timestamp.Timestamp `protobuf:"bytes,6,opt,name=destroy_event_time,json=destroyEventTime,proto3" json:"destroy_event_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CryptoKeyVersion) Reset()         { *m = CryptoKeyVersion{} }
func (m *CryptoKeyVersion) String() string { return proto.CompactTextString(m) }
func (*CryptoKeyVersion) ProtoMessage()    {}
func (*CryptoKeyVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_resources_3fd296b2ee28c3bd, []int{2}
}
func (m *CryptoKeyVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoKeyVersion.Unmarshal(m, b)
}
func (m *CryptoKeyVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoKeyVersion.Marshal(b, m, deterministic)
}
func (dst *CryptoKeyVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoKeyVersion.Merge(dst, src)
}
func (m *CryptoKeyVersion) XXX_Size() int {
	return xxx_messageInfo_CryptoKeyVersion.Size(m)
}
func (m *CryptoKeyVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoKeyVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoKeyVersion proto.InternalMessageInfo

func (m *CryptoKeyVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CryptoKeyVersion) GetState() CryptoKeyVersion_CryptoKeyVersionState {
	if m != nil {
		return m.State
	}
	return CryptoKeyVersion_CRYPTO_KEY_VERSION_STATE_UNSPECIFIED
}

func (m *CryptoKeyVersion) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *CryptoKeyVersion) GetDestroyTime() *timestamp.Timestamp {
	if m != nil {
		return m.DestroyTime
	}
	return nil
}

func (m *CryptoKeyVersion) GetDestroyEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.DestroyEventTime
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRing)(nil), "google.cloud.kms.v1.KeyRing")
	proto.RegisterType((*CryptoKey)(nil), "google.cloud.kms.v1.CryptoKey")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.kms.v1.CryptoKey.LabelsEntry")
	proto.RegisterType((*CryptoKeyVersion)(nil), "google.cloud.kms.v1.CryptoKeyVersion")
	proto.RegisterEnum("google.cloud.kms.v1.CryptoKey_CryptoKeyPurpose", CryptoKey_CryptoKeyPurpose_name, CryptoKey_CryptoKeyPurpose_value)
	proto.RegisterEnum("google.cloud.kms.v1.CryptoKeyVersion_CryptoKeyVersionState", CryptoKeyVersion_CryptoKeyVersionState_name, CryptoKeyVersion_CryptoKeyVersionState_value)
}

func init() {
	proto.RegisterFile("google/cloud/kms/v1/resources.proto", fileDescriptor_resources_3fd296b2ee28c3bd)
}

var fileDescriptor_resources_3fd296b2ee28c3bd = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x6e, 0xda, 0x4a,
	0x14, 0x86, 0x63, 0x20, 0x21, 0x1c, 0x72, 0x13, 0x32, 0xdc, 0xe8, 0x72, 0x51, 0x95, 0x22, 0xda,
	0x4a, 0xa8, 0x0f, 0xb6, 0x48, 0xa5, 0xaa, 0x6d, 0x54, 0x55, 0x01, 0xbb, 0x0d, 0x22, 0x02, 0x77,
	0x4c, 0x90, 0x12, 0x45, 0xb2, 0x1c, 0x98, 0x52, 0x0b, 0xec, 0xb1, 0x66, 0x0c, 0xaa, 0x5f, 0xbb,
	0x87, 0x6e, 0xa2, 0xeb, 0xe8, 0x53, 0x77, 0xd0, 0x9d, 0xf4, 0xb1, 0xf2, 0x78, 0x9c, 0xa6, 0x04,
	0x35, 0xc9, 0x13, 0x33, 0x67, 0xfe, 0xff, 0x33, 0x67, 0xce, 0x6f, 0xc3, 0xa3, 0x09, 0xa5, 0x93,
	0x19, 0xd1, 0x46, 0x33, 0x3a, 0x1f, 0x6b, 0x53, 0x8f, 0x6b, 0x8b, 0xa6, 0xc6, 0x08, 0xa7, 0x73,
	0x36, 0x22, 0x5c, 0x0d, 0x18, 0x0d, 0x29, 0x2a, 0x27, 0x22, 0x55, 0x88, 0xd4, 0xa9, 0xc7, 0xd5,
	0x45, 0xb3, 0xfa, 0x40, 0x3a, 0x9d, 0xc0, 0xd5, 0x1c, 0xdf, 0xa7, 0xa1, 0x13, 0xba, 0xd4, 0x97,
	0x96, 0xea, 0xbe, 0x3c, 0x15, 0xbb, 0xcb, 0xf9, 0x07, 0x6d, 0x3c, 0x67, 0x42, 0x20, 0xcf, 0x1f,
	0x2e, 0x9f, 0x87, 0xae, 0x47, 0x78, 0xe8, 0x78, 0x41, 0x22, 0xa8, 0x9f, 0x43, 0xbe, 0x4b, 0x22,
	0xec, 0xfa, 0x13, 0x84, 0x20, 0xe7, 0x3b, 0x1e, 0xa9, 0x28, 0x35, 0xa5, 0x51, 0xc0, 0x62, 0x8d,
	0x0e, 0xa1, 0x38, 0x62, 0xc4, 0x09, 0x89, 0x1d, 0x1b, 0x2b, 0x99, 0x9a, 0xd2, 0x28, 0x1e, 0x54,
	0x55, 0xf9, 0x47, 0x53, 0xaa, 0x3a, 0x48, 0xa9, 0x18, 0x12, 0x79, 0x5c, 0xa8, 0xff, 0xc8, 0x41,
	0xa1, 0xcd, 0xa2, 0x20, 0xa4, 0x5d, 0x12, 0xad, 0xc4, 0xbf, 0x81, 0x7c, 0xc0, 0x5c, 0xcf, 0x61,
	0x91, 0x44, 0x3f, 0x51, 0x57, 0xdc, 0x81, 0x7a, 0x05, 0x19, 0x12, 0xc6, 0x5d, 0xea, 0xe3, 0xd4,
	0x85, 0x3a, 0x90, 0x0f, 0xe6, 0x2c, 0xa0, 0x9c, 0x54, 0xb2, 0x35, 0xa5, 0xb1, 0x7d, 0xa0, 0xfd,
	0x1d, 0xf0, 0x7b, 0x65, 0x26, 0x36, 0x9c, 0xfa, 0x97, 0x5b, 0x5d, 0xbf, 0x4f, 0xab, 0xe8, 0x18,
	0x90, 0x4f, 0x3e, 0x85, 0x36, 0x93, 0xf3, 0x49, 0x18, 0xf9, 0x5b, 0x19, 0xa5, 0xd8, 0x85, 0xa5,
	0x49, 0x90, 0x74, 0xd8, 0xb9, 0x82, 0x04, 0x84, 0xb9, 0x74, 0x5c, 0xd9, 0x14, 0x98, 0xff, 0x6f,
	0x60, 0x74, 0x39, 0xeb, 0xe3, 0x35, 0xbc, 0x9d, 0x7a, 0x4c, 0x61, 0x41, 0x2d, 0xd8, 0x98, 0x39,
	0x97, 0x64, 0xc6, 0x2b, 0x50, 0xcb, 0x36, 0x8a, 0x07, 0x4f, 0x6f, 0xb9, 0x96, 0x13, 0x21, 0x36,
	0xfc, 0x90, 0x45, 0x58, 0x3a, 0xab, 0x2f, 0xa1, 0x78, 0xad, 0x8c, 0x4a, 0x90, 0x9d, 0x92, 0x48,
	0x8e, 0x2f, 0x5e, 0xa2, 0x7f, 0x61, 0x7d, 0xe1, 0xcc, 0xe6, 0x49, 0x2c, 0x0a, 0x38, 0xd9, 0xbc,
	0xca, 0xbc, 0x50, 0xea, 0x5d, 0x28, 0x2d, 0x5f, 0x34, 0xaa, 0xc3, 0x7e, 0x1b, 0x9f, 0x99, 0x83,
	0xbe, 0xdd, 0x35, 0xce, 0x6c, 0xf3, 0x14, 0x9b, 0x7d, 0xcb, 0xb0, 0x4f, 0x7b, 0x96, 0x69, 0xb4,
	0x3b, 0x6f, 0x3b, 0x86, 0x5e, 0x5a, 0x43, 0x65, 0xd8, 0x31, 0x7a, 0x42, 0x65, 0xeb, 0x86, 0xf8,
	0x2d, 0x29, 0xad, 0x32, 0xec, 0x5e, 0xdd, 0x08, 0x1f, 0x7d, 0x24, 0xe3, 0xf9, 0x8c, 0xd4, 0xbf,
	0x65, 0xaf, 0x3d, 0x42, 0xc6, 0x62, 0x65, 0xc4, 0xde, 0xc3, 0x3a, 0x0f, 0x9d, 0x30, 0xcd, 0xc7,
	0xe1, 0x9d, 0x02, 0x76, 0xa3, 0x60, 0xc5, 0x08, 0x9c, 0x90, 0x96, 0x93, 0x92, 0xbb, 0x57, 0x52,
	0x5e, 0xc3, 0xd6, 0x98, 0xf0, 0x90, 0xd1, 0xe8, 0xae, 0x39, 0x2b, 0x4a, 0x7d, 0x1a, 0xb4, 0xd4,
	0x4e, 0x16, 0xc4, 0x0f, 0x13, 0xc8, 0xc6, 0xed, 0x41, 0x93, 0x2e, 0x23, 0x36, 0x89, 0xb7, 0xf3,
	0xb3, 0x02, 0x7b, 0x2b, 0xdb, 0x44, 0x0d, 0x78, 0x7c, 0x6d, 0x52, 0x43, 0x03, 0x5b, 0x9d, 0x7e,
	0xcf, 0xb6, 0x06, 0x47, 0x83, 0xe5, 0x79, 0x15, 0x21, 0x6f, 0xf4, 0x8e, 0x5a, 0x27, 0x86, 0x5e,
	0x52, 0xd0, 0x16, 0x6c, 0xea, 0x1d, 0x2b, 0xd9, 0x65, 0xd0, 0x3f, 0x50, 0xd0, 0x0d, 0x6b, 0x80,
	0xfb, 0x67, 0x86, 0x5e, 0xca, 0xa2, 0x3d, 0xd8, 0x95, 0x5b, 0xdb, 0x6a, 0x1f, 0x1b, 0xfa, 0x69,
	0xac, 0xca, 0xb5, 0xbe, 0x28, 0xf0, 0xdf, 0x88, 0x7a, 0xab, 0x86, 0xd2, 0xda, 0xed, 0x7a, 0x1c,
	0xa7, 0x9f, 0x48, 0x33, 0x6e, 0xc9, 0x54, 0xce, 0x9f, 0x4b, 0xe5, 0x84, 0xce, 0x1c, 0x7f, 0xa2,
	0x52, 0x36, 0xd1, 0x26, 0xc4, 0x17, 0x0d, 0x6b, 0xc9, 0x91, 0x13, 0xb8, 0xfc, 0x8f, 0xef, 0xec,
	0xe1, 0xd4, 0xe3, 0x3f, 0x15, 0xe5, 0x6b, 0xa6, 0xfc, 0x2e, 0xf1, 0xb6, 0xc5, 0x53, 0xba, 0x1e,
	0x57, 0x87, 0xcd, 0xef, 0x69, 0xf5, 0x42, 0x54, 0x2f, 0xba, 0x1e, 0xbf, 0x18, 0x36, 0x2f, 0x37,
	0x04, 0xf1, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xc0, 0x0a, 0x0e, 0xb8, 0x05, 0x00,
	0x00,
}
