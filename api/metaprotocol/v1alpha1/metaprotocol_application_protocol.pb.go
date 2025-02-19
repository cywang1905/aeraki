// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/metaprotocol/v1alpha1/metaprotocol_application_protocol.proto

// $schema: metaprotocol.aeraki.io.v1alpha1.ApplicationProtocol
// $title: Application Protocol
// $description: ApplicationProtocol defines an application protocol built on top of MetaProtocol.
//
// ApplicationProtocol defines an application protocol built on top of MetaProtocol.
//
// ```yaml
// apiVersion: metaprotocol.aeraki.io/v1alpha1
// kind: ApplicationProtocol
// metadata:
//   name: dubbo
//   namespace: istio-system
// spec:
//   protocol: dubbo
//   codec: aeraki.meta_protocol.codec.dubbo
// ```

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
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

// ApplicationProtocol defines an application protocol built on top of MetaProtocol.
//
// <!-- crd generation tags
// +cue-gen:ApplicationProtocol:groupName:metaprotocol.aeraki.io
// +cue-gen:ApplicationProtocol:version:v1alpha1
// +cue-gen:ApplicationProtocol:storageVersion
// +cue-gen:ApplicationProtocol:annotations:helm.sh/resource-policy=keep
// +cue-gen:ApplicationProtocol:labels:app=aeraki,chart=aeraki,heritage=Tiller,release=aeraki
// +cue-gen:ApplicationProtocol:subresource:status
// +cue-gen:ApplicationProtocol:scope:Namespaced
// +cue-gen:ApplicationProtocol:resource:categories=aeraki-io,metaprotocol-aeraki-io
// +cue-gen:ApplicationProtocol:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=metaprotocol.aeraki.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ApplicationProtocol struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Codec                string   `protobuf:"bytes,2,opt,name=codec,proto3" json:"codec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationProtocol) Reset()         { *m = ApplicationProtocol{} }
func (m *ApplicationProtocol) String() string { return proto.CompactTextString(m) }
func (*ApplicationProtocol) ProtoMessage()    {}
func (*ApplicationProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_54bc1cd743033a01, []int{0}
}
func (m *ApplicationProtocol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationProtocol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplicationProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationProtocol.Merge(m, src)
}
func (m *ApplicationProtocol) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationProtocol proto.InternalMessageInfo

func (m *ApplicationProtocol) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *ApplicationProtocol) GetCodec() string {
	if m != nil {
		return m.Codec
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationProtocol)(nil), "metaprotocol.aeraki.io.v1alpha1.ApplicationProtocol")
}

func init() {
	proto.RegisterFile("api/metaprotocol/v1alpha1/metaprotocol_application_protocol.proto", fileDescriptor_54bc1cd743033a01)
}

var fileDescriptor_54bc1cd743033a01 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0x2c, 0xc8, 0xd4,
	0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x33, 0x4c,
	0xcc, 0x29, 0xc8, 0x48, 0x34, 0x44, 0x11, 0x8d, 0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c,
	0xc9, 0xcc, 0xcf, 0x8b, 0x87, 0x09, 0xea, 0x81, 0x19, 0x42, 0xf2, 0xc8, 0x0a, 0xf5, 0x12, 0x53,
	0x8b, 0x12, 0xb3, 0x33, 0xf5, 0x32, 0xf3, 0xf5, 0x60, 0x06, 0x49, 0xc9, 0xa7, 0xe7, 0xe7, 0xa7,
	0xe7, 0xa4, 0xea, 0x83, 0xac, 0x4a, 0xcb, 0x4c, 0xcd, 0x49, 0x89, 0x4f, 0x4a, 0xcd, 0x48, 0x2c,
	0xcb, 0xcc, 0x2f, 0x82, 0x98, 0xa0, 0xe4, 0xce, 0x25, 0xec, 0x88, 0x30, 0x3f, 0x00, 0x6a, 0x94,
	0x90, 0x14, 0x17, 0x07, 0xcc, 0x58, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48,
	0x84, 0x8b, 0x35, 0x39, 0x3f, 0x25, 0x35, 0x59, 0x82, 0x09, 0x2c, 0x01, 0xe1, 0x38, 0x79, 0x9d,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x51, 0x36, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0x67, 0xe9, 0xa6, 0x15, 0x25, 0xe6,
	0xa6, 0x96, 0xe7, 0x17, 0x65, 0x43, 0x05, 0xf4, 0x71, 0x7a, 0x3f, 0x89, 0x0d, 0x2c, 0x64, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x18, 0x9c, 0xda, 0x22, 0x01, 0x00, 0x00,
}

func (m *ApplicationProtocol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationProtocol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplicationProtocol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codec) > 0 {
		i -= len(m.Codec)
		copy(dAtA[i:], m.Codec)
		i = encodeVarintMetaprotocolApplicationProtocol(dAtA, i, uint64(len(m.Codec)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarintMetaprotocolApplicationProtocol(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetaprotocolApplicationProtocol(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetaprotocolApplicationProtocol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApplicationProtocol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovMetaprotocolApplicationProtocol(uint64(l))
	}
	l = len(m.Codec)
	if l > 0 {
		n += 1 + l + sovMetaprotocolApplicationProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetaprotocolApplicationProtocol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetaprotocolApplicationProtocol(x uint64) (n int) {
	return sovMetaprotocolApplicationProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApplicationProtocol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetaprotocolApplicationProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApplicationProtocol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationProtocol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaprotocolApplicationProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetaprotocolApplicationProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetaprotocolApplicationProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetaprotocolApplicationProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetaprotocolApplicationProtocol
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetaprotocolApplicationProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetaprotocolApplicationProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMetaprotocolApplicationProtocol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetaprotocolApplicationProtocol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetaprotocolApplicationProtocol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetaprotocolApplicationProtocol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetaprotocolApplicationProtocol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetaprotocolApplicationProtocol = fmt.Errorf("proto: unexpected end of group")
)
