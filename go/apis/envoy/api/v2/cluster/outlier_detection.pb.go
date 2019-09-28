// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package envoy_api_v2_cluster

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

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
type OutlierDetection struct {
	// The number of consecutive 5xx responses or local origin errors that are mapped
	// to 5xx error codes before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *types.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *types.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *types.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *types.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *types.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *types.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *types.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *types.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *types.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes)
	// before a consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *types.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *types.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	// Determines whether to distinguish local origin failures from external errors. If set to true
	// the following configuration parameters are taken into account:
	// :ref:`consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_cluster.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingLocalOriginSuccessRate *types.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}           `json:"-"`
	XXX_unrecognized                []byte             `json:"-"`
	XXX_sizecache                   int32              `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_56cd87362a3f00c9, []int{0}
}
func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(m, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return m.Size()
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *types.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *types.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *types.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *types.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *types.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *types.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if m != nil {
		return m.SplitExternalLocalOriginErrors
	}
	return false
}

func (m *OutlierDetection) GetConsecutiveLocalOriginFailure() *types.UInt32Value {
	if m != nil {
		return m.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *types.UInt32Value {
	if m != nil {
		return m.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor_56cd87362a3f00c9)
}

var fileDescriptor_56cd87362a3f00c9 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcb, 0x6e, 0xd3, 0x4e,
	0x18, 0xc5, 0xff, 0x4e, 0x2f, 0xff, 0x76, 0x0a, 0x6d, 0x35, 0x2a, 0xad, 0xd3, 0x42, 0xa8, 0x82,
	0x40, 0x55, 0x85, 0x6c, 0x29, 0x55, 0x1f, 0xa0, 0x69, 0x53, 0x2e, 0x02, 0x5a, 0x25, 0x50, 0x84,
	0x00, 0x8d, 0xa6, 0xce, 0x97, 0x30, 0xc8, 0xf6, 0x98, 0xb9, 0xb8, 0x2e, 0x2c, 0x79, 0x9b, 0xf2,
	0x06, 0xac, 0x58, 0xb2, 0xe4, 0x11, 0x50, 0x77, 0xf0, 0x14, 0xc8, 0x76, 0x2e, 0x4e, 0x6a, 0x44,
	0xbc, 0x8b, 0xf4, 0x9d, 0xf3, 0x3b, 0x67, 0x3e, 0x8f, 0x63, 0x74, 0x1f, 0xfc, 0x90, 0x9f, 0xdb,
	0x34, 0x60, 0x76, 0x58, 0xb3, 0x1d, 0x57, 0x4b, 0x05, 0xc2, 0xe6, 0x5a, 0xb9, 0x0c, 0x04, 0x69,
	0x83, 0x02, 0x47, 0x31, 0xee, 0x5b, 0x81, 0xe0, 0x8a, 0xe3, 0x95, 0x44, 0x6d, 0xd1, 0x80, 0x59,
	0x61, 0xcd, 0xea, 0xa9, 0xd7, 0x2b, 0x5d, 0xce, 0xbb, 0x2e, 0xd8, 0x89, 0xe6, 0x54, 0x77, 0xec,
	0xb6, 0x16, 0x74, 0xe8, 0xba, 0x3a, 0x3f, 0x13, 0x34, 0x08, 0x40, 0xc8, 0xde, 0x7c, 0x2d, 0xa4,
	0x2e, 0x6b, 0x53, 0x05, 0x76, 0xff, 0x47, 0x3a, 0xa8, 0x7e, 0x59, 0x40, 0xcb, 0x47, 0x69, 0x95,
	0x83, 0x7e, 0x13, 0xdc, 0x40, 0x4b, 0x0e, 0xf7, 0x25, 0x38, 0x5a, 0xb1, 0x10, 0xc8, 0x6e, 0x14,
	0x99, 0xc6, 0xa6, 0xb1, 0xb5, 0x50, 0xbb, 0x69, 0xa5, 0x39, 0x56, 0x3f, 0xc7, 0x7a, 0xf1, 0xc8,
	0x57, 0x3b, 0xb5, 0x13, 0xea, 0x6a, 0x68, 0x2e, 0x66, 0x4c, 0xbb, 0x51, 0x84, 0xf7, 0xd0, 0x1c,
	0xf3, 0x15, 0x88, 0x90, 0xba, 0x66, 0x29, 0xf1, 0x97, 0xaf, 0xf8, 0x0f, 0x7a, 0xe7, 0xa8, 0xa3,
	0xaf, 0xbf, 0xbe, 0x4d, 0xcd, 0x5c, 0x18, 0xa5, 0xed, 0xff, 0x9a, 0x03, 0x1b, 0x6e, 0x21, 0x7c,
	0x4a, 0x25, 0x10, 0x78, 0x9f, 0x56, 0x23, 0x8a, 0x79, 0x60, 0x4e, 0x15, 0x81, 0x2d, 0xc7, 0x80,
	0x46, 0xcf, 0xff, 0x9c, 0x79, 0x80, 0x5f, 0xa1, 0x15, 0x8f, 0x46, 0x43, 0x66, 0x00, 0xc2, 0x01,
	0x5f, 0x99, 0xd3, 0xff, 0x3e, 0x63, 0x7d, 0x3e, 0x26, 0x4f, 0x6f, 0x97, 0xcc, 0x76, 0x13, 0x7b,
	0x34, 0xea, 0x73, 0x8f, 0x53, 0x04, 0x76, 0x50, 0x19, 0xfc, 0x0e, 0x17, 0x0e, 0xf3, 0xbb, 0x64,
	0x7c, 0x87, 0x33, 0xc5, 0xf8, 0x6b, 0x03, 0xd2, 0xfe, 0xe8, 0x5e, 0xdf, 0xa2, 0xd5, 0x61, 0x88,
	0xd4, 0x8e, 0x03, 0x52, 0x12, 0x41, 0x15, 0x98, 0xb3, 0xc5, 0x12, 0x56, 0x06, 0x98, 0x56, 0x4a,
	0x69, 0x52, 0x15, 0xaf, 0x67, 0x3d, 0x0b, 0x25, 0x1e, 0xf3, 0x99, 0xa7, 0x3d, 0xf2, 0x8e, 0x4b,
	0x25, 0xcd, 0xff, 0x27, 0xb8, 0x08, 0x6b, 0x72, 0x88, 0x7b, 0x9a, 0xba, 0x1f, 0xc6, 0x66, 0xfc,
	0x1a, 0x6d, 0x8c, 0xa0, 0x05, 0x7c, 0xd0, 0x20, 0x15, 0x09, 0xb9, 0xab, 0x3d, 0x30, 0xe7, 0x26,
	0x60, 0x9b, 0x19, 0x76, 0x33, 0xb5, 0x9f, 0x24, 0x6e, 0xfc, 0x12, 0x95, 0x47, 0xe0, 0x52, 0xb5,
	0x21, 0x24, 0x1d, 0xea, 0x28, 0x2e, 0xcc, 0xf9, 0x09, 0xd0, 0xab, 0x19, 0x74, 0x2b, 0x36, 0x1f,
	0x26, 0x5e, 0xfc, 0x06, 0x6d, 0x64, 0x1f, 0x65, 0x97, 0x2a, 0x38, 0xa3, 0xe7, 0xa4, 0x43, 0x99,
	0xab, 0x05, 0x98, 0x68, 0x02, 0x74, 0x39, 0x03, 0x78, 0x90, 0xfa, 0x0f, 0x53, 0x3b, 0xfe, 0x88,
	0xee, 0xe6, 0x5f, 0x99, 0xf1, 0x9c, 0x85, 0x62, 0x0f, 0xb7, 0x9a, 0x77, 0x7d, 0xc6, 0xb2, 0x1f,
	0xa3, 0xaa, 0x0c, 0x5c, 0xa6, 0x08, 0x44, 0x0a, 0x84, 0x4f, 0x5d, 0xe2, 0x72, 0x87, 0xba, 0x84,
	0x0b, 0xd6, 0x65, 0x3e, 0x01, 0x21, 0xb8, 0x90, 0xe6, 0xb5, 0x4d, 0x63, 0x6b, 0xae, 0x59, 0x49,
	0x94, 0x8d, 0x9e, 0xf0, 0x49, 0xac, 0x3b, 0x4a, 0x64, 0x8d, 0x44, 0x85, 0x01, 0x6d, 0x66, 0xdb,
	0x8f, 0x80, 0xfa, 0x47, 0xb8, 0x3e, 0xc1, 0xaa, 0x6e, 0x65, 0x28, 0x99, 0x94, 0x7e, 0xe5, 0xcf,
	0x06, 0xda, 0xce, 0xdf, 0x57, 0x6e, 0xe2, 0x62, 0xb1, 0xa5, 0xdd, 0xcb, 0x5b, 0x5a, 0x4e, 0x0b,
	0x8d, 0xee, 0x0c, 0x4b, 0x8c, 0x04, 0x8f, 0xbc, 0x8f, 0x4b, 0xc5, 0xd2, 0x6f, 0x0f, 0x98, 0x99,
	0xc8, 0xcc, 0xab, 0x59, 0xff, 0xf4, 0xfd, 0xb2, 0x62, 0xfc, 0xb8, 0xac, 0x18, 0x3f, 0x2f, 0x2b,
	0x06, 0xaa, 0x32, 0x6e, 0x25, 0x5f, 0x8b, 0x40, 0xf0, 0xe8, 0xdc, 0xca, 0xfb, 0x70, 0xd4, 0x6f,
	0x8c, 0xff, 0xb9, 0x1f, 0xc7, 0x1d, 0x8e, 0x8d, 0x8b, 0xd2, 0x6a, 0x23, 0xd1, 0xef, 0x05, 0xcc,
	0x3a, 0xa9, 0x59, 0xfb, 0xa9, 0xfe, 0x59, 0xeb, 0xf7, 0xdf, 0x06, 0xa7, 0xb3, 0x49, 0xfd, 0x9d,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x43, 0x31, 0x7b, 0xd0, 0x06, 0x00, 0x00,
}

func (m *OutlierDetection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutlierDetection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutlierDetection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EnforcingLocalOriginSuccessRate != nil {
		{
			size, err := m.EnforcingLocalOriginSuccessRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	if m.EnforcingConsecutiveLocalOriginFailure != nil {
		{
			size, err := m.EnforcingConsecutiveLocalOriginFailure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.ConsecutiveLocalOriginFailure != nil {
		{
			size, err := m.ConsecutiveLocalOriginFailure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.SplitExternalLocalOriginErrors {
		i--
		if m.SplitExternalLocalOriginErrors {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if m.EnforcingConsecutiveGatewayFailure != nil {
		{
			size, err := m.EnforcingConsecutiveGatewayFailure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.ConsecutiveGatewayFailure != nil {
		{
			size, err := m.ConsecutiveGatewayFailure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.SuccessRateStdevFactor != nil {
		{
			size, err := m.SuccessRateStdevFactor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.SuccessRateRequestVolume != nil {
		{
			size, err := m.SuccessRateRequestVolume.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.SuccessRateMinimumHosts != nil {
		{
			size, err := m.SuccessRateMinimumHosts.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.EnforcingSuccessRate != nil {
		{
			size, err := m.EnforcingSuccessRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.EnforcingConsecutive_5Xx != nil {
		{
			size, err := m.EnforcingConsecutive_5Xx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.MaxEjectionPercent != nil {
		{
			size, err := m.MaxEjectionPercent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.BaseEjectionTime != nil {
		{
			size, err := m.BaseEjectionTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Interval != nil {
		{
			size, err := m.Interval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Consecutive_5Xx != nil {
		{
			size, err := m.Consecutive_5Xx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOutlierDetection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutlierDetection(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutlierDetection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutlierDetection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Consecutive_5Xx != nil {
		l = m.Consecutive_5Xx.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.Interval != nil {
		l = m.Interval.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.BaseEjectionTime != nil {
		l = m.BaseEjectionTime.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.MaxEjectionPercent != nil {
		l = m.MaxEjectionPercent.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingConsecutive_5Xx != nil {
		l = m.EnforcingConsecutive_5Xx.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingSuccessRate != nil {
		l = m.EnforcingSuccessRate.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateMinimumHosts != nil {
		l = m.SuccessRateMinimumHosts.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateRequestVolume != nil {
		l = m.SuccessRateRequestVolume.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateStdevFactor != nil {
		l = m.SuccessRateStdevFactor.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.ConsecutiveGatewayFailure != nil {
		l = m.ConsecutiveGatewayFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingConsecutiveGatewayFailure != nil {
		l = m.EnforcingConsecutiveGatewayFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SplitExternalLocalOriginErrors {
		n += 2
	}
	if m.ConsecutiveLocalOriginFailure != nil {
		l = m.ConsecutiveLocalOriginFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingConsecutiveLocalOriginFailure != nil {
		l = m.EnforcingConsecutiveLocalOriginFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingLocalOriginSuccessRate != nil {
		l = m.EnforcingLocalOriginSuccessRate.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOutlierDetection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutlierDetection(x uint64) (n int) {
	return sovOutlierDetection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutlierDetection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutlierDetection
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
			return fmt.Errorf("proto: OutlierDetection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutlierDetection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consecutive_5Xx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Consecutive_5Xx == nil {
				m.Consecutive_5Xx = &types.UInt32Value{}
			}
			if err := m.Consecutive_5Xx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Interval == nil {
				m.Interval = &types.Duration{}
			}
			if err := m.Interval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseEjectionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseEjectionTime == nil {
				m.BaseEjectionTime = &types.Duration{}
			}
			if err := m.BaseEjectionTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEjectionPercent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxEjectionPercent == nil {
				m.MaxEjectionPercent = &types.UInt32Value{}
			}
			if err := m.MaxEjectionPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingConsecutive_5Xx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingConsecutive_5Xx == nil {
				m.EnforcingConsecutive_5Xx = &types.UInt32Value{}
			}
			if err := m.EnforcingConsecutive_5Xx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingSuccessRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingSuccessRate == nil {
				m.EnforcingSuccessRate = &types.UInt32Value{}
			}
			if err := m.EnforcingSuccessRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateMinimumHosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateMinimumHosts == nil {
				m.SuccessRateMinimumHosts = &types.UInt32Value{}
			}
			if err := m.SuccessRateMinimumHosts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateRequestVolume", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateRequestVolume == nil {
				m.SuccessRateRequestVolume = &types.UInt32Value{}
			}
			if err := m.SuccessRateRequestVolume.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateStdevFactor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateStdevFactor == nil {
				m.SuccessRateStdevFactor = &types.UInt32Value{}
			}
			if err := m.SuccessRateStdevFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsecutiveGatewayFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsecutiveGatewayFailure == nil {
				m.ConsecutiveGatewayFailure = &types.UInt32Value{}
			}
			if err := m.ConsecutiveGatewayFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingConsecutiveGatewayFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingConsecutiveGatewayFailure == nil {
				m.EnforcingConsecutiveGatewayFailure = &types.UInt32Value{}
			}
			if err := m.EnforcingConsecutiveGatewayFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitExternalLocalOriginErrors", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SplitExternalLocalOriginErrors = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsecutiveLocalOriginFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsecutiveLocalOriginFailure == nil {
				m.ConsecutiveLocalOriginFailure = &types.UInt32Value{}
			}
			if err := m.ConsecutiveLocalOriginFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingConsecutiveLocalOriginFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingConsecutiveLocalOriginFailure == nil {
				m.EnforcingConsecutiveLocalOriginFailure = &types.UInt32Value{}
			}
			if err := m.EnforcingConsecutiveLocalOriginFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingLocalOriginSuccessRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
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
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingLocalOriginSuccessRate == nil {
				m.EnforcingLocalOriginSuccessRate = &types.UInt32Value{}
			}
			if err := m.EnforcingLocalOriginSuccessRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutlierDetection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOutlierDetection
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
func skipOutlierDetection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutlierDetection
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
					return 0, ErrIntOverflowOutlierDetection
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
					return 0, ErrIntOverflowOutlierDetection
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
				return 0, ErrInvalidLengthOutlierDetection
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOutlierDetection
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOutlierDetection
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
				next, err := skipOutlierDetection(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOutlierDetection
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
	ErrInvalidLengthOutlierDetection = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutlierDetection   = fmt.Errorf("proto: integer overflow")
)
