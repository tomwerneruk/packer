// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iot/devices/v1/registry.proto

package devices

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	math "math"
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

type Registry_Status int32

const (
	Registry_STATUS_UNSPECIFIED Registry_Status = 0
	Registry_CREATING           Registry_Status = 1
	Registry_ACTIVE             Registry_Status = 2
	Registry_DELETING           Registry_Status = 3
)

var Registry_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "ACTIVE",
	3: "DELETING",
}

var Registry_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"ACTIVE":             2,
	"DELETING":           3,
}

func (x Registry_Status) String() string {
	return proto.EnumName(Registry_Status_name, int32(x))
}

func (Registry_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{0, 0}
}

type Registry struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId             string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status               Registry_Status      `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.iot.devices.v1.Registry_Status" json:"status,omitempty"`
	LogGroupId           string               `protobuf:"bytes,8,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{0}
}

func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Registry) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Registry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Registry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Registry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Registry) GetStatus() Registry_Status {
	if m != nil {
		return m.Status
	}
	return Registry_STATUS_UNSPECIFIED
}

func (m *Registry) GetLogGroupId() string {
	if m != nil {
		return m.LogGroupId
	}
	return ""
}

type RegistryCertificate struct {
	RegistryId           string               `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	Fingerprint          string               `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	CertificateData      string               `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistryCertificate) Reset()         { *m = RegistryCertificate{} }
func (m *RegistryCertificate) String() string { return proto.CompactTextString(m) }
func (*RegistryCertificate) ProtoMessage()    {}
func (*RegistryCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{1}
}

func (m *RegistryCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryCertificate.Unmarshal(m, b)
}
func (m *RegistryCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryCertificate.Marshal(b, m, deterministic)
}
func (m *RegistryCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryCertificate.Merge(m, src)
}
func (m *RegistryCertificate) XXX_Size() int {
	return xxx_messageInfo_RegistryCertificate.Size(m)
}
func (m *RegistryCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryCertificate proto.InternalMessageInfo

func (m *RegistryCertificate) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *RegistryCertificate) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *RegistryCertificate) GetCertificateData() string {
	if m != nil {
		return m.CertificateData
	}
	return ""
}

func (m *RegistryCertificate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type DeviceAlias struct {
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// prefix of canonical topic name to be aliased, e.g. $devices/abcdef
	TopicPrefix          string   `protobuf:"bytes,2,opt,name=topic_prefix,json=topicPrefix,proto3" json:"topic_prefix,omitempty"`
	Alias                string   `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceAlias) Reset()         { *m = DeviceAlias{} }
func (m *DeviceAlias) String() string { return proto.CompactTextString(m) }
func (*DeviceAlias) ProtoMessage()    {}
func (*DeviceAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{2}
}

func (m *DeviceAlias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAlias.Unmarshal(m, b)
}
func (m *DeviceAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAlias.Marshal(b, m, deterministic)
}
func (m *DeviceAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAlias.Merge(m, src)
}
func (m *DeviceAlias) XXX_Size() int {
	return xxx_messageInfo_DeviceAlias.Size(m)
}
func (m *DeviceAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAlias.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAlias proto.InternalMessageInfo

func (m *DeviceAlias) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceAlias) GetTopicPrefix() string {
	if m != nil {
		return m.TopicPrefix
	}
	return ""
}

func (m *DeviceAlias) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func init() {
	proto.RegisterEnum("yandex.cloud.iot.devices.v1.Registry_Status", Registry_Status_name, Registry_Status_value)
	proto.RegisterType((*Registry)(nil), "yandex.cloud.iot.devices.v1.Registry")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.iot.devices.v1.Registry.LabelsEntry")
	proto.RegisterType((*RegistryCertificate)(nil), "yandex.cloud.iot.devices.v1.RegistryCertificate")
	proto.RegisterType((*DeviceAlias)(nil), "yandex.cloud.iot.devices.v1.DeviceAlias")
}

func init() {
	proto.RegisterFile("yandex/cloud/iot/devices/v1/registry.proto", fileDescriptor_39c05472a87f1ea4)
}

var fileDescriptor_39c05472a87f1ea4 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x26, 0xed, 0x16, 0xda, 0x97, 0x69, 0x54, 0x06, 0xa1, 0xa8, 0x13, 0x5a, 0xe8, 0xa9, 0x20,
	0xe6, 0x68, 0xe3, 0xc2, 0xe0, 0x54, 0xda, 0x50, 0x22, 0x4d, 0xd3, 0x94, 0x76, 0x1c, 0xb8, 0x44,
	0x6e, 0xec, 0x06, 0x8b, 0x34, 0x8e, 0x1c, 0xa7, 0x5a, 0xff, 0x1c, 0xbf, 0x85, 0x9f, 0x82, 0x62,
	0x27, 0xac, 0x70, 0x98, 0xe0, 0x66, 0x7f, 0xef, 0x7b, 0xef, 0x7d, 0x7e, 0xef, 0x33, 0xbc, 0xde,
	0x91, 0x9c, 0xb2, 0x3b, 0x3f, 0xc9, 0x44, 0x45, 0x7d, 0x2e, 0x94, 0x4f, 0xd9, 0x96, 0x27, 0xac,
	0xf4, 0xb7, 0xe7, 0xbe, 0x64, 0x29, 0x2f, 0x95, 0xdc, 0xe1, 0x42, 0x0a, 0x25, 0xd0, 0x89, 0xe1,
	0x62, 0xcd, 0xc5, 0x5c, 0x28, 0xdc, 0x70, 0xf1, 0xf6, 0x7c, 0x78, 0x9a, 0x0a, 0x91, 0x66, 0xcc,
	0xd7, 0xd4, 0x55, 0xb5, 0xf6, 0x15, 0xdf, 0xb0, 0x52, 0x91, 0x4d, 0x61, 0xb2, 0x87, 0x2f, 0xfe,
	0xe8, 0xb4, 0x25, 0x19, 0xa7, 0x44, 0x71, 0x91, 0x9b, 0xf0, 0xe8, 0x67, 0x17, 0x7a, 0x51, 0xd3,
	0x0f, 0x1d, 0x43, 0x87, 0x53, 0xd7, 0xf2, 0xac, 0x71, 0x3f, 0xea, 0x70, 0x8a, 0x4e, 0xa0, 0xbf,
	0x16, 0x19, 0x65, 0x32, 0xe6, 0xd4, 0xed, 0x68, 0xb8, 0x67, 0x80, 0x90, 0xa2, 0x4b, 0x80, 0x44,
	0x32, 0xa2, 0x18, 0x8d, 0x89, 0x72, 0xbb, 0x9e, 0x35, 0x76, 0x2e, 0x86, 0xd8, 0xc8, 0xc1, 0xad,
	0x1c, 0xbc, 0x6c, 0xe5, 0x44, 0xfd, 0x86, 0x3d, 0x51, 0x08, 0xc1, 0x41, 0x4e, 0x36, 0xcc, 0x3d,
	0xd0, 0x25, 0xf5, 0x19, 0x79, 0xe0, 0x50, 0x56, 0x26, 0x92, 0x17, 0xb5, 0x3a, 0xf7, 0x50, 0x87,
	0xf6, 0x21, 0x14, 0x82, 0x9d, 0x91, 0x15, 0xcb, 0x4a, 0xd7, 0xf6, 0xba, 0x63, 0xe7, 0xe2, 0x1c,
	0x3f, 0x30, 0x18, 0xdc, 0x3e, 0x0a, 0x5f, 0xe9, 0x9c, 0x20, 0x57, 0x72, 0x17, 0x35, 0x05, 0xd0,
	0x0c, 0xec, 0x52, 0x11, 0x55, 0x95, 0xee, 0x63, 0xcf, 0x1a, 0x1f, 0x5f, 0xbc, 0xf9, 0xb7, 0x52,
	0x0b, 0x9d, 0x13, 0x35, 0xb9, 0xc8, 0x83, 0xa3, 0x4c, 0xa4, 0x71, 0x2a, 0x45, 0x55, 0xd4, 0x13,
	0xea, 0x69, 0xcd, 0x90, 0x89, 0x74, 0x5e, 0x43, 0x21, 0x1d, 0x5e, 0x82, 0xb3, 0xd7, 0x1e, 0x0d,
	0xa0, 0xfb, 0x9d, 0xed, 0x9a, 0x01, 0xd7, 0x47, 0xf4, 0x0c, 0x0e, 0xb7, 0x24, 0xab, 0x58, 0x33,
	0x5d, 0x73, 0x79, 0xdf, 0x79, 0x67, 0x8d, 0x3e, 0x83, 0x6d, 0xda, 0xa1, 0xe7, 0x80, 0x16, 0xcb,
	0xc9, 0xf2, 0x76, 0x11, 0xdf, 0x5e, 0x2f, 0x6e, 0x82, 0x69, 0xf8, 0x29, 0x0c, 0x66, 0x83, 0x47,
	0xe8, 0x08, 0x7a, 0xd3, 0x28, 0x98, 0x2c, 0xc3, 0xeb, 0xf9, 0xc0, 0x42, 0x00, 0xf6, 0x64, 0xba,
	0x0c, 0xbf, 0x04, 0x83, 0x4e, 0x1d, 0x99, 0x05, 0x57, 0x81, 0x8e, 0x74, 0x47, 0x3f, 0x2c, 0x78,
	0xda, 0x3e, 0x61, 0xca, 0xa4, 0xe2, 0x6b, 0x9e, 0x10, 0xc5, 0xd0, 0x29, 0x38, 0xad, 0xd3, 0xe2,
	0xdf, 0x6b, 0x87, 0x16, 0x0a, 0x69, 0xbd, 0x92, 0x35, 0xcf, 0x53, 0x26, 0x0b, 0xc9, 0x73, 0xd5,
	0x48, 0xdc, 0x87, 0xd0, 0x2b, 0x18, 0x24, 0xf7, 0x15, 0x63, 0x4a, 0x14, 0xd1, 0x4e, 0xe8, 0x47,
	0x4f, 0xf6, 0xf0, 0x19, 0x51, 0xe4, 0x2f, 0xbb, 0x1c, 0xfc, 0x87, 0x5d, 0x46, 0x09, 0x38, 0x33,
	0xbd, 0x8d, 0x49, 0xc6, 0x49, 0x59, 0xbb, 0xd2, 0x2c, 0xe7, 0x5e, 0x75, 0xcf, 0x00, 0x21, 0x45,
	0x2f, 0xe1, 0x48, 0x89, 0x82, 0x27, 0x71, 0x21, 0xd9, 0x9a, 0xdf, 0xb5, 0xa2, 0x35, 0x76, 0xa3,
	0xa1, 0x7a, 0xe6, 0xa4, 0x2e, 0xd4, 0x28, 0x35, 0x97, 0x8f, 0xe1, 0xd7, 0x79, 0xca, 0xd5, 0xb7,
	0x6a, 0x85, 0x13, 0xb1, 0xf1, 0x8d, 0x1d, 0xce, 0xcc, 0xa7, 0x49, 0xc5, 0x59, 0xca, 0x72, 0xad,
	0xd1, 0x7f, 0xe0, 0xdf, 0x7e, 0x68, 0x8e, 0x2b, 0x5b, 0x53, 0xdf, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xea, 0xea, 0x45, 0x1e, 0xe5, 0x03, 0x00, 0x00,
}
