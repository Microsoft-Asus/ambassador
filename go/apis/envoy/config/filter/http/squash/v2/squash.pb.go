// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v2/squash.proto

package envoy_config_filter_http_squash_v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *types.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *types.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *types.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *types.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fc8434388b1e13, []int{0}
}
func (m *Squash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return m.Size()
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *types.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *types.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *types.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *types.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v2/squash.proto", fileDescriptor_63fc8434388b1e13)
}

var fileDescriptor_63fc8434388b1e13 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4f, 0x32, 0x31,
	0x18, 0xc7, 0x53, 0xe0, 0xe5, 0x0d, 0x25, 0xd1, 0x78, 0x1a, 0x39, 0x89, 0xb9, 0x10, 0x5c, 0x98,
	0x5a, 0x83, 0xdf, 0xe0, 0xe2, 0xc0, 0x62, 0xbc, 0x00, 0x3b, 0x29, 0x47, 0x81, 0x26, 0xe5, 0x7a,
	0xf4, 0x9e, 0x5e, 0xe4, 0xab, 0x39, 0x39, 0x3a, 0x3a, 0x39, 0x1b, 0x36, 0xbf, 0x85, 0xa1, 0x2d,
	0x81, 0xe8, 0xc0, 0xf6, 0x24, 0xff, 0xfe, 0x7e, 0xcf, 0xbf, 0x79, 0x30, 0xe5, 0x59, 0xa9, 0x36,
	0x34, 0x55, 0xd9, 0x5c, 0x2c, 0xe8, 0x5c, 0x48, 0xe0, 0x9a, 0x2e, 0x01, 0x72, 0x5a, 0xac, 0x0d,
	0x2b, 0x96, 0xb4, 0xec, 0xfb, 0x89, 0xe4, 0x5a, 0x81, 0x0a, 0xba, 0x16, 0x20, 0x0e, 0x20, 0x0e,
	0x20, 0x3b, 0x80, 0xf8, 0x67, 0x65, 0xbf, 0x1d, 0x2d, 0x94, 0x5a, 0x48, 0x4e, 0x2d, 0x31, 0x35,
	0x73, 0x3a, 0x33, 0x9a, 0x81, 0x50, 0x99, 0x73, 0xb4, 0x6f, 0x7f, 0xe7, 0x05, 0x68, 0x93, 0x82,
	0x4f, 0x5b, 0x25, 0x93, 0x62, 0xc6, 0x80, 0xd3, 0xfd, 0xe0, 0x82, 0xee, 0x67, 0x05, 0xd7, 0x47,
	0x76, 0x49, 0x70, 0x87, 0xff, 0xa7, 0xd2, 0x14, 0xc0, 0x75, 0x88, 0x3a, 0xa8, 0xd7, 0x88, 0x1b,
	0xaf, 0xdf, 0x6f, 0xd5, 0x9a, 0xae, 0x74, 0xd0, 0x70, 0x9f, 0x04, 0x03, 0x7c, 0xc9, 0x00, 0x58,
	0xba, 0x5c, 0xf1, 0x0c, 0x26, 0xc0, 0x57, 0xb9, 0x64, 0xc0, 0xc3, 0x4a, 0x07, 0xf5, 0x9a, 0xfd,
	0x16, 0x71, 0x25, 0xc8, 0xbe, 0x04, 0x19, 0xd9, 0x12, 0xc3, 0xe0, 0xc0, 0x8c, 0x3d, 0x12, 0xc4,
	0xf8, 0x5c, 0xf3, 0xb5, 0xe1, 0x05, 0x4c, 0x40, 0xac, 0xb8, 0x32, 0x10, 0x56, 0xad, 0xe5, 0xe6,
	0x8f, 0xe5, 0xd1, 0x7f, 0x75, 0x78, 0xe6, 0x89, 0xb1, 0x03, 0x82, 0x01, 0x0e, 0x8e, 0xdb, 0x78,
	0x4d, 0xed, 0x94, 0xe6, 0xe2, 0xa8, 0x8e, 0x37, 0x3d, 0xe3, 0xeb, 0x23, 0x53, 0xae, 0xa4, 0x9c,
	0xe4, 0x5c, 0x0b, 0x35, 0x0b, 0xff, 0x9d, 0xb2, 0x5d, 0x1d, 0xc0, 0x44, 0x49, 0x99, 0x58, 0x2c,
	0x7e, 0x7a, 0xdf, 0x46, 0xe8, 0x63, 0x1b, 0xa1, 0xaf, 0x6d, 0x84, 0xf0, 0xbd, 0x50, 0xc4, 0x1e,
	0x39, 0xd7, 0xea, 0x65, 0x43, 0x4e, 0xdf, 0x3b, 0x6e, 0xba, 0xab, 0x24, 0xbb, 0x75, 0x09, 0x9a,
	0xd6, 0xed, 0xde, 0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x31, 0xc3, 0x9d, 0x5c, 0x02,
	0x00, 0x00,
}

func (m *Squash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Squash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Squash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AttachmentPollPeriod != nil {
		{
			size, err := m.AttachmentPollPeriod.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.AttachmentTimeout != nil {
		{
			size, err := m.AttachmentTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.RequestTimeout != nil {
		{
			size, err := m.RequestTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AttachmentTemplate != nil {
		{
			size, err := m.AttachmentTemplate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSquash(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintSquash(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSquash(dAtA []byte, offset int, v uint64) int {
	offset -= sovSquash(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Squash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTemplate != nil {
		l = m.AttachmentTemplate.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.RequestTimeout != nil {
		l = m.RequestTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTimeout != nil {
		l = m.AttachmentTimeout.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentPollPeriod != nil {
		l = m.AttachmentPollPeriod.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSquash(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSquash(x uint64) (n int) {
	return sovSquash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Squash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSquash
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
			return fmt.Errorf("proto: Squash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Squash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTemplate == nil {
				m.AttachmentTemplate = &types.Struct{}
			}
			if err := m.AttachmentTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = &types.Duration{}
			}
			if err := m.RequestTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTimeout == nil {
				m.AttachmentTimeout = &types.Duration{}
			}
			if err := m.AttachmentTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentPollPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSquash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentPollPeriod == nil {
				m.AttachmentPollPeriod = &types.Duration{}
			}
			if err := m.AttachmentPollPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSquash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSquash
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSquash
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
func skipSquash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSquash
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
				return 0, ErrInvalidLengthSquash
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSquash
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSquash
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSquash(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSquash
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSquash = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSquash   = fmt.Errorf("proto: integer overflow")
)
