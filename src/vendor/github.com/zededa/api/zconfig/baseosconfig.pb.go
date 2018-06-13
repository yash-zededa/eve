// Code generated by protoc-gen-go. DO NOT EDIT.
// source: baseosconfig.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// OS version key and value pair
type OSKeyTags struct {
	OSVerKey   string `protobuf:"bytes,1,opt,name=OSVerKey" json:"OSVerKey,omitempty"`
	OSVerValue string `protobuf:"bytes,2,opt,name=OSVerValue" json:"OSVerValue,omitempty"`
}

func (m *OSKeyTags) Reset()                    { *m = OSKeyTags{} }
func (m *OSKeyTags) String() string            { return proto.CompactTextString(m) }
func (*OSKeyTags) ProtoMessage()               {}
func (*OSKeyTags) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *OSKeyTags) GetOSVerKey() string {
	if m != nil {
		return m.OSVerKey
	}
	return ""
}

func (m *OSKeyTags) GetOSVerValue() string {
	if m != nil {
		return m.OSVerValue
	}
	return ""
}

// repeated key value tags compromising
type OSVerDetails struct {
	BaseOSParams []*OSKeyTags `protobuf:"bytes,12,rep,name=baseOSParams" json:"baseOSParams,omitempty"`
}

func (m *OSVerDetails) Reset()                    { *m = OSVerDetails{} }
func (m *OSVerDetails) String() string            { return proto.CompactTextString(m) }
func (*OSVerDetails) ProtoMessage()               {}
func (*OSVerDetails) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *OSVerDetails) GetBaseOSParams() []*OSKeyTags {
	if m != nil {
		return m.BaseOSParams
	}
	return nil
}

type BaseOSConfig struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	Drives         []*Drive        `protobuf:"bytes,3,rep,name=drives" json:"drives,omitempty"`
	Activate       bool            `protobuf:"varint,4,opt,name=activate" json:"activate,omitempty"`
	BaseOSVersion  string          `protobuf:"bytes,10,opt,name=baseOSVersion" json:"baseOSVersion,omitempty"`
	BaseOSDetails  *OSVerDetails   `protobuf:"bytes,11,opt,name=baseOSDetails" json:"baseOSDetails,omitempty"`
}

func (m *BaseOSConfig) Reset()                    { *m = BaseOSConfig{} }
func (m *BaseOSConfig) String() string            { return proto.CompactTextString(m) }
func (*BaseOSConfig) ProtoMessage()               {}
func (*BaseOSConfig) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *BaseOSConfig) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *BaseOSConfig) GetDrives() []*Drive {
	if m != nil {
		return m.Drives
	}
	return nil
}

func (m *BaseOSConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *BaseOSConfig) GetBaseOSVersion() string {
	if m != nil {
		return m.BaseOSVersion
	}
	return ""
}

func (m *BaseOSConfig) GetBaseOSDetails() *OSVerDetails {
	if m != nil {
		return m.BaseOSDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*OSKeyTags)(nil), "OSKeyTags")
	proto.RegisterType((*OSVerDetails)(nil), "OSVerDetails")
	proto.RegisterType((*BaseOSConfig)(nil), "BaseOSConfig")
}

func init() { proto.RegisterFile("baseosconfig.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0xfa, 0x40,
	0x10, 0xc5, 0xc9, 0xdf, 0x3f, 0xa2, 0x63, 0xac, 0xb0, 0xa7, 0x45, 0xa8, 0x15, 0xe9, 0xc1, 0xd3,
	0x0a, 0x7a, 0xe8, 0xad, 0x14, 0x2b, 0x94, 0xe2, 0xc1, 0x12, 0xab, 0x87, 0xde, 0xc6, 0xdd, 0x69,
	0xba, 0x60, 0xb2, 0xb2, 0x9b, 0x04, 0xf4, 0xc3, 0xf6, 0xb3, 0x14, 0x37, 0xa9, 0x98, 0xde, 0xf2,
	0x7e, 0xf3, 0xc8, 0xbc, 0xb7, 0x03, 0x6c, 0x87, 0x8e, 0x8c, 0x93, 0x26, 0xfd, 0xd4, 0xb1, 0x38,
	0x58, 0x93, 0x99, 0x7e, 0x4f, 0x51, 0x21, 0x4d, 0x92, 0x98, 0xb4, 0x02, 0x5d, 0x97, 0x19, 0x8b,
	0x31, 0x95, 0x72, 0xf4, 0x02, 0xed, 0xd5, 0x7a, 0x49, 0xc7, 0x77, 0x8c, 0x1d, 0xeb, 0x43, 0x6b,
	0xb5, 0xde, 0x92, 0x5d, 0xd2, 0x91, 0x07, 0xc3, 0x60, 0xdc, 0x8e, 0x2e, 0x9a, 0x0d, 0x00, 0xfc,
	0xf7, 0x16, 0xf7, 0x39, 0xf1, 0x7f, 0x7e, 0x7a, 0x45, 0x46, 0x8f, 0x10, 0x7a, 0xb5, 0xa0, 0x0c,
	0xf5, 0xde, 0x31, 0x01, 0xe1, 0x39, 0xce, 0x6a, 0xfd, 0x86, 0x16, 0x13, 0xc7, 0xc3, 0x61, 0x63,
	0xdc, 0x99, 0x82, 0xb8, 0x6c, 0x8b, 0x6a, 0xf3, 0xd1, 0x77, 0x00, 0xe1, 0xdc, 0x83, 0x67, 0x9f,
	0x9f, 0x3d, 0xc0, 0x4d, 0x9e, 0x6b, 0x85, 0xa9, 0x2a, 0xc8, 0x3a, 0x6d, 0x52, 0x1f, 0xa9, 0x33,
	0xed, 0x89, 0xcd, 0xe6, 0x75, 0x81, 0xa9, 0xda, 0x96, 0x38, 0xfa, 0x63, 0x63, 0x03, 0x68, 0x2a,
	0xab, 0x0b, 0x72, 0xbc, 0xe1, 0x77, 0x36, 0xc5, 0xe2, 0x2c, 0xa3, 0x8a, 0x9e, 0x5b, 0xa2, 0xcc,
	0x74, 0x81, 0x19, 0xf1, 0xff, 0xc3, 0x60, 0xdc, 0x8a, 0x2e, 0x9a, 0xdd, 0x43, 0xb7, 0x4c, 0x55,
	0xfd, 0x9c, 0x83, 0x2f, 0x5a, 0x87, 0x6c, 0xf6, 0xeb, 0xaa, 0xca, 0xf2, 0x8e, 0x4f, 0xd6, 0x15,
	0xd7, 0x2f, 0x10, 0xd5, 0x3d, 0xf3, 0x27, 0xb8, 0x93, 0x26, 0x11, 0x27, 0x52, 0xa4, 0x50, 0xc8,
	0xbd, 0xc9, 0x95, 0xc8, 0x1d, 0xd9, 0x42, 0xcb, 0xea, 0x18, 0x1f, 0xb7, 0xb1, 0xce, 0xbe, 0xf2,
	0x9d, 0x90, 0x26, 0x99, 0x94, 0xbe, 0x09, 0x1e, 0xf4, 0xe4, 0x54, 0x5e, 0x74, 0xd7, 0xf4, 0xae,
	0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xed, 0x96, 0x7e, 0xe8, 0x01, 0x00, 0x00,
}
