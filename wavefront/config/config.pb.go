// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/wavefront/config/config.proto

/*
	Package config is a generated protocol buffer package.

	config for wavefront

	It is generated from these files:
		mixer/adapter/wavefront/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// represents metric types as in https://docs.wavefront.com/metric_types.html
type Params_MetricInfo_Type int32

const (
	// reserved for unknown metric types
	UNKNOWN Params_MetricInfo_Type = 0
	// represents a gauge metric type
	GAUGE Params_MetricInfo_Type = 1
	// represents a counter metric type
	COUNTER Params_MetricInfo_Type = 2
	// represents a delta counter metric type
	DELTA_COUNTER Params_MetricInfo_Type = 3
	// represents a histogram metric type
	HISTOGRAM Params_MetricInfo_Type = 4
)

var Params_MetricInfo_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "GAUGE",
	2: "COUNTER",
	3: "DELTA_COUNTER",
	4: "HISTOGRAM",
}
var Params_MetricInfo_Type_value = map[string]int32{
	"UNKNOWN":       0,
	"GAUGE":         1,
	"COUNTER":       2,
	"DELTA_COUNTER": 3,
	"HISTOGRAM":     4,
}

func (Params_MetricInfo_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{0, 2, 0}
}

// config for wavefront
type Params struct {
	// the wavefront service credentials
	//
	// Types that are valid to be assigned to Credentials:
	//	*Params_Direct
	//	*Params_Proxy
	Credentials isParams_Credentials `protobuf_oneof:"credentials"`
	// the metrics flush interval
	FlushInterval time.Duration `protobuf:"bytes,3,opt,name=flush_interval,json=flushInterval,stdduration" json:"flush_interval"`
	// the prefix to prepend metrics with
	Prefix string `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// the metrics
	Metrics []*Params_MetricInfo `protobuf:"bytes,5,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

type isParams_Credentials interface {
	isParams_Credentials()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Params_Direct struct {
	Direct *Params_WavefrontDirect `protobuf:"bytes,1,opt,name=direct,oneof"`
}
type Params_Proxy struct {
	Proxy *Params_WavefrontProxy `protobuf:"bytes,2,opt,name=proxy,oneof"`
}

func (*Params_Direct) isParams_Credentials() {}
func (*Params_Proxy) isParams_Credentials()  {}

func (m *Params) GetCredentials() isParams_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *Params) GetDirect() *Params_WavefrontDirect {
	if x, ok := m.GetCredentials().(*Params_Direct); ok {
		return x.Direct
	}
	return nil
}

func (m *Params) GetProxy() *Params_WavefrontProxy {
	if x, ok := m.GetCredentials().(*Params_Proxy); ok {
		return x.Proxy
	}
	return nil
}

func (m *Params) GetFlushInterval() time.Duration {
	if m != nil {
		return m.FlushInterval
	}
	return 0
}

func (m *Params) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Params) GetMetrics() []*Params_MetricInfo {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Params) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Params_OneofMarshaler, _Params_OneofUnmarshaler, _Params_OneofSizer, []interface{}{
		(*Params_Direct)(nil),
		(*Params_Proxy)(nil),
	}
}

func _Params_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Params)
	// credentials
	switch x := m.Credentials.(type) {
	case *Params_Direct:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Direct); err != nil {
			return err
		}
	case *Params_Proxy:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Proxy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Params.Credentials has unexpected type %T", x)
	}
	return nil
}

func _Params_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Params)
	switch tag {
	case 1: // credentials.direct
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Params_WavefrontDirect)
		err := b.DecodeMessage(msg)
		m.Credentials = &Params_Direct{msg}
		return true, err
	case 2: // credentials.proxy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Params_WavefrontProxy)
		err := b.DecodeMessage(msg)
		m.Credentials = &Params_Proxy{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Params_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Params)
	// credentials
	switch x := m.Credentials.(type) {
	case *Params_Direct:
		s := proto.Size(x.Direct)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Params_Proxy:
		s := proto.Size(x.Proxy)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// represents wavefront server credentials
type Params_WavefrontDirect struct {
	// the wavefront server
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// the wavefront token
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *Params_WavefrontDirect) Reset()                    { *m = Params_WavefrontDirect{} }
func (*Params_WavefrontDirect) ProtoMessage()               {}
func (*Params_WavefrontDirect) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

func (m *Params_WavefrontDirect) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Params_WavefrontDirect) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// represents wavefront proxy credentials
type Params_WavefrontProxy struct {
	// the wavefront proxy address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Params_WavefrontProxy) Reset()                    { *m = Params_WavefrontProxy{} }
func (*Params_WavefrontProxy) ProtoMessage()               {}
func (*Params_WavefrontProxy) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

func (m *Params_WavefrontProxy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// represents a metric
type Params_MetricInfo struct {
	// the metric name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the metric type
	Type Params_MetricInfo_Type `protobuf:"varint,2,opt,name=type,proto3,enum=wavefront.config.Params_MetricInfo_Type" json:"type,omitempty"`
}

func (m *Params_MetricInfo) Reset()                    { *m = Params_MetricInfo{} }
func (*Params_MetricInfo) ProtoMessage()               {}
func (*Params_MetricInfo) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 2} }

func (m *Params_MetricInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params_MetricInfo) GetType() Params_MetricInfo_Type {
	if m != nil {
		return m.Type
	}
	return UNKNOWN
}

func init() {
	proto.RegisterType((*Params)(nil), "wavefront.config.Params")
	proto.RegisterType((*Params_WavefrontDirect)(nil), "wavefront.config.Params.WavefrontDirect")
	proto.RegisterType((*Params_WavefrontProxy)(nil), "wavefront.config.Params.WavefrontProxy")
	proto.RegisterType((*Params_MetricInfo)(nil), "wavefront.config.Params.MetricInfo")
	proto.RegisterEnum("wavefront.config.Params_MetricInfo_Type", Params_MetricInfo_Type_name, Params_MetricInfo_Type_value)
}
func (x Params_MetricInfo_Type) String() string {
	s, ok := Params_MetricInfo_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if that1.Credentials == nil {
		if this.Credentials != nil {
			return false
		}
	} else if this.Credentials == nil {
		return false
	} else if !this.Credentials.Equal(that1.Credentials) {
		return false
	}
	if this.FlushInterval != that1.FlushInterval {
		return false
	}
	if this.Prefix != that1.Prefix {
		return false
	}
	if len(this.Metrics) != len(that1.Metrics) {
		return false
	}
	for i := range this.Metrics {
		if !this.Metrics[i].Equal(that1.Metrics[i]) {
			return false
		}
	}
	return true
}
func (this *Params_Direct) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params_Direct)
	if !ok {
		that2, ok := that.(Params_Direct)
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
	if !this.Direct.Equal(that1.Direct) {
		return false
	}
	return true
}
func (this *Params_Proxy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params_Proxy)
	if !ok {
		that2, ok := that.(Params_Proxy)
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
	if !this.Proxy.Equal(that1.Proxy) {
		return false
	}
	return true
}
func (this *Params_WavefrontDirect) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params_WavefrontDirect)
	if !ok {
		that2, ok := that.(Params_WavefrontDirect)
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
	if this.Server != that1.Server {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	return true
}
func (this *Params_WavefrontProxy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params_WavefrontProxy)
	if !ok {
		that2, ok := that.(Params_WavefrontProxy)
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
	if this.Address != that1.Address {
		return false
	}
	return true
}
func (this *Params_MetricInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params_MetricInfo)
	if !ok {
		that2, ok := that.(Params_MetricInfo)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&config.Params{")
	if this.Credentials != nil {
		s = append(s, "Credentials: "+fmt.Sprintf("%#v", this.Credentials)+",\n")
	}
	s = append(s, "FlushInterval: "+fmt.Sprintf("%#v", this.FlushInterval)+",\n")
	s = append(s, "Prefix: "+fmt.Sprintf("%#v", this.Prefix)+",\n")
	if this.Metrics != nil {
		s = append(s, "Metrics: "+fmt.Sprintf("%#v", this.Metrics)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Params_Direct) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&config.Params_Direct{` +
		`Direct:` + fmt.Sprintf("%#v", this.Direct) + `}`}, ", ")
	return s
}
func (this *Params_Proxy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&config.Params_Proxy{` +
		`Proxy:` + fmt.Sprintf("%#v", this.Proxy) + `}`}, ", ")
	return s
}
func (this *Params_WavefrontDirect) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&config.Params_WavefrontDirect{")
	s = append(s, "Server: "+fmt.Sprintf("%#v", this.Server)+",\n")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Params_WavefrontProxy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&config.Params_WavefrontProxy{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Params_MetricInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&config.Params_MetricInfo{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Credentials != nil {
		nn1, err := m.Credentials.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(types.SizeOfStdDuration(m.FlushInterval)))
	n2, err := types.StdDurationMarshalTo(m.FlushInterval, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Prefix) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Metrics) > 0 {
		for _, msg := range m.Metrics {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Params_Direct) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Direct != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Direct.Size()))
		n3, err := m.Direct.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Params_Proxy) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Proxy != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Proxy.Size()))
		n4, err := m.Proxy.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Params_WavefrontDirect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_WavefrontDirect) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Server) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Server)))
		i += copy(dAtA[i:], m.Server)
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *Params_WavefrontProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_WavefrontProxy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	return i, nil
}

func (m *Params_MetricInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_MetricInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	if m.Credentials != nil {
		n += m.Credentials.Size()
	}
	l = types.SizeOfStdDuration(m.FlushInterval)
	n += 1 + l + sovConfig(uint64(l))
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *Params_Direct) Size() (n int) {
	var l int
	_ = l
	if m.Direct != nil {
		l = m.Direct.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}
func (m *Params_Proxy) Size() (n int) {
	var l int
	_ = l
	if m.Proxy != nil {
		l = m.Proxy.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}
func (m *Params_WavefrontDirect) Size() (n int) {
	var l int
	_ = l
	l = len(m.Server)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Params_WavefrontProxy) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Params_MetricInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovConfig(uint64(m.Type))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Credentials:` + fmt.Sprintf("%v", this.Credentials) + `,`,
		`FlushInterval:` + strings.Replace(strings.Replace(this.FlushInterval.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`Prefix:` + fmt.Sprintf("%v", this.Prefix) + `,`,
		`Metrics:` + strings.Replace(fmt.Sprintf("%v", this.Metrics), "Params_MetricInfo", "Params_MetricInfo", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Direct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_Direct{`,
		`Direct:` + strings.Replace(fmt.Sprintf("%v", this.Direct), "Params_WavefrontDirect", "Params_WavefrontDirect", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_Proxy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_Proxy{`,
		`Proxy:` + strings.Replace(fmt.Sprintf("%v", this.Proxy), "Params_WavefrontProxy", "Params_WavefrontProxy", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_WavefrontDirect) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_WavefrontDirect{`,
		`Server:` + fmt.Sprintf("%v", this.Server) + `,`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_WavefrontProxy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_WavefrontProxy{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params_MetricInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params_MetricInfo{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direct", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Params_WavefrontDirect{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Credentials = &Params_Direct{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proxy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Params_WavefrontProxy{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Credentials = &Params_Proxy{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.FlushInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, &Params_MetricInfo{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_WavefrontDirect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WavefrontDirect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WavefrontDirect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Server = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_WavefrontProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WavefrontProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WavefrontProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params_MetricInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MetricInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Params_MetricInfo_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/wavefront/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x6f, 0x12, 0x41,
	0x18, 0xde, 0x69, 0x17, 0x90, 0x97, 0x80, 0x38, 0x69, 0xcc, 0xca, 0x61, 0x4a, 0xea, 0x41, 0xa2,
	0xc9, 0x6e, 0x52, 0x3d, 0x6a, 0x1a, 0x10, 0x02, 0xa8, 0x40, 0x33, 0x42, 0x9a, 0x78, 0x69, 0xa6,
	0xec, 0x2c, 0x6e, 0x84, 0x9d, 0xcd, 0xec, 0x80, 0x70, 0xf3, 0x27, 0x78, 0xf4, 0x27, 0x78, 0xf1,
	0x7f, 0xf4, 0xd8, 0xa3, 0x07, 0xa3, 0xb2, 0x5e, 0x3c, 0xf6, 0x27, 0x18, 0x66, 0x17, 0x1b, 0x9b,
	0xa8, 0xa7, 0x99, 0xe7, 0x7d, 0x9f, 0xe7, 0x79, 0xbf, 0xe0, 0xc1, 0xcc, 0x5f, 0x72, 0xe9, 0x30,
	0x97, 0x85, 0x8a, 0x4b, 0xe7, 0x2d, 0x5b, 0x70, 0x4f, 0x8a, 0x40, 0x39, 0x63, 0x11, 0x78, 0xfe,
	0x24, 0x7d, 0xec, 0x50, 0x0a, 0x25, 0x70, 0xf9, 0x77, 0xda, 0x4e, 0xe2, 0x95, 0xbd, 0x89, 0x98,
	0x08, 0x9d, 0x74, 0x36, 0xbf, 0x84, 0x57, 0x21, 0x13, 0x21, 0x26, 0x53, 0xee, 0x68, 0x74, 0x36,
	0xf7, 0x1c, 0x77, 0x2e, 0x99, 0xf2, 0x45, 0x90, 0xe4, 0x0f, 0xbe, 0x98, 0x90, 0x3d, 0x66, 0x92,
	0xcd, 0x22, 0xdc, 0x80, 0xac, 0xeb, 0x4b, 0x3e, 0x56, 0x16, 0xaa, 0xa2, 0x5a, 0xe1, 0xb0, 0x66,
	0x5f, 0xaf, 0x61, 0x27, 0x4c, 0xfb, 0x64, 0x1b, 0x6f, 0x6a, 0x7e, 0xc7, 0xa0, 0xa9, 0x12, 0x1f,
	0x41, 0x26, 0x94, 0x62, 0xb9, 0xb2, 0x76, 0xb4, 0xc5, 0xbd, 0xff, 0x5b, 0x1c, 0x6f, 0xe8, 0x1d,
	0x83, 0x26, 0x3a, 0xfc, 0x0c, 0x4a, 0xde, 0x74, 0x1e, 0xbd, 0x3e, 0xf5, 0x03, 0xc5, 0xe5, 0x82,
	0x4d, 0xad, 0x5d, 0xed, 0x74, 0xc7, 0x4e, 0x06, 0xb1, 0xb7, 0x83, 0xd8, 0xcd, 0x74, 0x90, 0xc6,
	0x8d, 0xf3, 0xaf, 0xfb, 0xc6, 0x87, 0x6f, 0xfb, 0x88, 0x16, 0xb5, 0xb4, 0x9b, 0x2a, 0xf1, 0x6d,
	0xc8, 0x86, 0x92, 0x7b, 0xfe, 0xd2, 0x32, 0xab, 0xa8, 0x96, 0xa7, 0x29, 0xc2, 0x4f, 0x20, 0x37,
	0xe3, 0x4a, 0xfa, 0xe3, 0xc8, 0xca, 0x54, 0x77, 0x6b, 0x85, 0xc3, 0xbb, 0x7f, 0x6d, 0xb3, 0xa7,
	0x79, 0xdd, 0xc0, 0x13, 0x74, 0xab, 0xa9, 0x1c, 0xc1, 0xcd, 0x6b, 0x0b, 0xd8, 0x54, 0x8a, 0xb8,
	0x5c, 0x70, 0xa9, 0x57, 0x97, 0xa7, 0x29, 0xc2, 0x7b, 0x90, 0x51, 0xe2, 0x0d, 0x0f, 0xf4, 0x3a,
	0xf2, 0x34, 0x01, 0x95, 0xfb, 0x50, 0xfa, 0x73, 0x7c, 0x6c, 0x41, 0x8e, 0xb9, 0xae, 0xe4, 0x51,
	0x94, 0x1a, 0x6c, 0x61, 0xe5, 0x13, 0x02, 0xb8, 0x6a, 0x02, 0x63, 0x30, 0x03, 0x36, 0xe3, 0x29,
	0x4b, 0xff, 0xf1, 0x63, 0x30, 0xd5, 0x2a, 0xe4, 0xba, 0x46, 0xe9, 0x1f, 0x57, 0xbb, 0xb2, 0xb1,
	0x87, 0xab, 0x90, 0x53, 0xad, 0x3a, 0xe8, 0x81, 0xb9, 0x41, 0xb8, 0x00, 0xb9, 0x51, 0xff, 0x79,
	0x7f, 0x70, 0xd2, 0x2f, 0x1b, 0x38, 0x0f, 0x99, 0x76, 0x7d, 0xd4, 0x6e, 0x95, 0xd1, 0x26, 0xfe,
	0x74, 0x30, 0xea, 0x0f, 0x5b, 0xb4, 0xbc, 0x83, 0x6f, 0x41, 0xb1, 0xd9, 0x7a, 0x31, 0xac, 0x9f,
	0x6e, 0x43, 0xbb, 0xb8, 0x08, 0xf9, 0x4e, 0xf7, 0xe5, 0x70, 0xd0, 0xa6, 0xf5, 0x5e, 0xd9, 0x6c,
	0x14, 0xa1, 0x30, 0x96, 0xdc, 0xe5, 0x81, 0xf2, 0xd9, 0x34, 0x6a, 0x3c, 0xba, 0x58, 0x13, 0xe3,
	0xf3, 0x9a, 0x18, 0x97, 0x6b, 0x82, 0xde, 0xc5, 0x04, 0x7d, 0x8c, 0x09, 0x3a, 0x8f, 0x09, 0xba,
	0x88, 0x09, 0xfa, 0x1e, 0x13, 0xf4, 0x33, 0x26, 0xc6, 0x65, 0x4c, 0xd0, 0xfb, 0x1f, 0xc4, 0x78,
	0x95, 0x4d, 0x1a, 0x3e, 0xcb, 0xea, 0x23, 0x3f, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xfa,
	0x1e, 0x66, 0x12, 0x03, 0x00, 0x00,
}
