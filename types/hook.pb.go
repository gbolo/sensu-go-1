// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hook.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HookConfig is the specification of a hook
type HookConfig struct {
	// Name is the unique identifier for a hook
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Command is the command to be executed
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the timeout, in seconds, at which the hook has to run
	Timeout uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout"`
	// Stdin indicates if hook requests have stdin enabled
	Stdin bool `protobuf:"varint,4,opt,name=stdin,proto3" json:"stdin"`
	// Environment indicates to which env a hook belongs to
	Environment string `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	// Organization indicates to which org a hook belongs to
	Organization         string   `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookConfig) Reset()         { *m = HookConfig{} }
func (m *HookConfig) String() string { return proto.CompactTextString(m) }
func (*HookConfig) ProtoMessage()    {}
func (*HookConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_95eb01caedc29e8f, []int{0}
}
func (m *HookConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HookConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HookConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HookConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookConfig.Merge(dst, src)
}
func (m *HookConfig) XXX_Size() int {
	return m.Size()
}
func (m *HookConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HookConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HookConfig proto.InternalMessageInfo

func (m *HookConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HookConfig) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *HookConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *HookConfig) GetStdin() bool {
	if m != nil {
		return m.Stdin
	}
	return false
}

func (m *HookConfig) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *HookConfig) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

// A Hook is a hook specification and optionally the results of the hook's
// execution.
type Hook struct {
	// Config is the specification of a hook
	HookConfig `protobuf:"bytes,1,opt,name=config,embedded=config" json:""`
	// Duration of execution
	Duration float64 `protobuf:"fixed64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Executed describes the time in which the hook request was executed
	Executed int64 `protobuf:"varint,3,opt,name=executed,proto3" json:"executed"`
	// Issued describes the time in which the hook request was issued
	Issued int64 `protobuf:"varint,4,opt,name=issued,proto3" json:"issued"`
	// Output from the execution of Command
	Output string `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	// Status is the exit status code produced by the hook
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hook) Reset()         { *m = Hook{} }
func (m *Hook) String() string { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()    {}
func (*Hook) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_95eb01caedc29e8f, []int{1}
}
func (m *Hook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Hook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Hook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hook.Merge(dst, src)
}
func (m *Hook) XXX_Size() int {
	return m.Size()
}
func (m *Hook) XXX_DiscardUnknown() {
	xxx_messageInfo_Hook.DiscardUnknown(m)
}

var xxx_messageInfo_Hook proto.InternalMessageInfo

func (m *Hook) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Hook) GetExecuted() int64 {
	if m != nil {
		return m.Executed
	}
	return 0
}

func (m *Hook) GetIssued() int64 {
	if m != nil {
		return m.Issued
	}
	return 0
}

func (m *Hook) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Hook) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type HookList struct {
	// Hooks is the list of hooks for the check hook
	Hooks []string `protobuf:"bytes,1,rep,name=hooks" json:"hooks"`
	// Type indicates the type or response code for the check hook
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookList) Reset()         { *m = HookList{} }
func (m *HookList) String() string { return proto.CompactTextString(m) }
func (*HookList) ProtoMessage()    {}
func (*HookList) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_95eb01caedc29e8f, []int{2}
}
func (m *HookList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HookList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HookList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HookList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookList.Merge(dst, src)
}
func (m *HookList) XXX_Size() int {
	return m.Size()
}
func (m *HookList) XXX_DiscardUnknown() {
	xxx_messageInfo_HookList.DiscardUnknown(m)
}

var xxx_messageInfo_HookList proto.InternalMessageInfo

func (m *HookList) GetHooks() []string {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func (m *HookList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*HookConfig)(nil), "sensu.types.HookConfig")
	proto.RegisterType((*Hook)(nil), "sensu.types.Hook")
	proto.RegisterType((*HookList)(nil), "sensu.types.HookList")
}
func (this *HookConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HookConfig)
	if !ok {
		that2, ok := that.(HookConfig)
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
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.Stdin != that1.Stdin {
		return false
	}
	if this.Environment != that1.Environment {
		return false
	}
	if this.Organization != that1.Organization {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Hook) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Hook)
	if !ok {
		that2, ok := that.(Hook)
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
	if !this.HookConfig.Equal(&that1.HookConfig) {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.Executed != that1.Executed {
		return false
	}
	if this.Issued != that1.Issued {
		return false
	}
	if this.Output != that1.Output {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HookList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HookList)
	if !ok {
		that2, ok := that.(HookList)
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
	if len(this.Hooks) != len(that1.Hooks) {
		return false
	}
	for i := range this.Hooks {
		if this.Hooks[i] != that1.Hooks[i] {
			return false
		}
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *HookConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Command) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Command)))
		i += copy(dAtA[i:], m.Command)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Timeout))
	}
	if m.Stdin {
		dAtA[i] = 0x20
		i++
		if m.Stdin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Environment) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Environment)))
		i += copy(dAtA[i:], m.Environment)
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Hook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hook) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHook(dAtA, i, uint64(m.HookConfig.Size()))
	n1, err := m.HookConfig.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Duration != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Duration))))
		i += 8
	}
	if m.Executed != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Executed))
	}
	if m.Issued != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Issued))
	}
	if len(m.Output) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Output)))
		i += copy(dAtA[i:], m.Output)
	}
	if m.Status != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HookList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hooks) > 0 {
		for _, s := range m.Hooks {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHook(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHookConfig(r randyHook, easy bool) *HookConfig {
	this := &HookConfig{}
	this.Name = string(randStringHook(r))
	this.Command = string(randStringHook(r))
	this.Timeout = uint32(r.Uint32())
	this.Stdin = bool(bool(r.Intn(2) == 0))
	this.Environment = string(randStringHook(r))
	this.Organization = string(randStringHook(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 7)
	}
	return this
}

func NewPopulatedHook(r randyHook, easy bool) *Hook {
	this := &Hook{}
	v1 := NewPopulatedHookConfig(r, easy)
	this.HookConfig = *v1
	this.Duration = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Duration *= -1
	}
	this.Executed = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Executed *= -1
	}
	this.Issued = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Issued *= -1
	}
	this.Output = string(randStringHook(r))
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 7)
	}
	return this
}

func NewPopulatedHookList(r randyHook, easy bool) *HookList {
	this := &HookList{}
	v2 := r.Intn(10)
	this.Hooks = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Hooks[i] = string(randStringHook(r))
	}
	this.Type = string(randStringHook(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 3)
	}
	return this
}

type randyHook interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHook(r randyHook) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHook(r randyHook) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneHook(r)
	}
	return string(tmps)
}
func randUnrecognizedHook(r randyHook, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHook(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHook(dAtA []byte, r randyHook, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateHook(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHook(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHook(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HookConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHook(uint64(m.Timeout))
	}
	if m.Stdin {
		n += 2
	}
	l = len(m.Environment)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Hook) Size() (n int) {
	var l int
	_ = l
	l = m.HookConfig.Size()
	n += 1 + l + sovHook(uint64(l))
	if m.Duration != 0 {
		n += 9
	}
	if m.Executed != 0 {
		n += 1 + sovHook(uint64(m.Executed))
	}
	if m.Issued != 0 {
		n += 1 + sovHook(uint64(m.Issued))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovHook(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HookList) Size() (n int) {
	var l int
	_ = l
	if len(m.Hooks) > 0 {
		for _, s := range m.Hooks {
			l = len(s)
			n += 1 + l + sovHook(uint64(l))
		}
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHook(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHook(x uint64) (n int) {
	return sovHook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HookConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: HookConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stdin = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Environment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func (m *Hook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: Hook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HookConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HookConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Duration = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executed", wireType)
			}
			m.Executed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Executed |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			m.Issued = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Issued |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func (m *HookList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: HookList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hooks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hooks = append(m.Hooks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func skipHook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHook
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
					return 0, ErrIntOverflowHook
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
					return 0, ErrIntOverflowHook
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
				return 0, ErrInvalidLengthHook
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHook
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
				next, err := skipHook(dAtA[start:])
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
	ErrInvalidLengthHook = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHook   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hook.proto", fileDescriptor_hook_95eb01caedc29e8f) }

var fileDescriptor_hook_95eb01caedc29e8f = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x6b, 0xda, 0xe6, 0xb6, 0xa7, 0x65, 0xf1, 0x00, 0xd1, 0x1d, 0x92, 0x28, 0x08, 0x29,
	0x0b, 0xb9, 0x12, 0xcc, 0x08, 0x29, 0x2c, 0x0c, 0x4c, 0x1e, 0xd9, 0xd2, 0xc6, 0x37, 0xd7, 0xba,
	0x8a, 0x4f, 0x55, 0xdb, 0x08, 0x78, 0x12, 0x1e, 0x81, 0x47, 0xe0, 0x11, 0x3a, 0xde, 0x27, 0x88,
	0x20, 0x6c, 0xe1, 0x05, 0x18, 0x51, 0x8e, 0x93, 0x52, 0xa6, 0x73, 0xfe, 0x4f, 0x3e, 0xb2, 0xbe,
	0x63, 0x03, 0xdc, 0x21, 0xde, 0xe7, 0x87, 0x23, 0x5a, 0xe4, 0x1b, 0x23, 0xb5, 0x71, 0xb9, 0xfd,
	0x7c, 0x90, 0xe6, 0xfa, 0x45, 0xad, 0xec, 0x9d, 0xdb, 0xe5, 0x7b, 0x6c, 0x6e, 0x6a, 0xac, 0xf1,
	0x86, 0xce, 0xec, 0xdc, 0x2d, 0x25, 0x0a, 0xd4, 0xf9, 0xd9, 0xf4, 0xc4, 0x00, 0xde, 0x21, 0xde,
	0xbf, 0x45, 0x7d, 0xab, 0x6a, 0xce, 0x61, 0xa1, 0xcb, 0x46, 0x86, 0x2c, 0x61, 0xd9, 0x5a, 0x50,
	0xcf, 0x43, 0xb8, 0xda, 0x63, 0xd3, 0x94, 0xba, 0x0a, 0x1f, 0x11, 0x9e, 0x22, 0x7f, 0x0e, 0x57,
	0x56, 0x35, 0x12, 0x9d, 0x0d, 0xe7, 0x09, 0xcb, 0x1e, 0x17, 0x9b, 0xbe, 0x8d, 0x27, 0x24, 0xa6,
	0x86, 0xc7, 0xb0, 0x34, 0xb6, 0x52, 0x3a, 0x5c, 0x24, 0x2c, 0x5b, 0x15, 0xeb, 0xbe, 0x8d, 0x3d,
	0x10, 0xbe, 0xf0, 0x04, 0x36, 0x52, 0x7f, 0x54, 0x47, 0xd4, 0x8d, 0xd4, 0x36, 0x5c, 0xd2, 0x2d,
	0x97, 0x88, 0xa7, 0xb0, 0xc5, 0x63, 0x5d, 0x6a, 0xf5, 0xa5, 0xb4, 0x0a, 0x75, 0x18, 0xd0, 0x91,
	0xff, 0x58, 0xfa, 0x9b, 0xc1, 0x62, 0x50, 0xe1, 0xaf, 0x21, 0xd8, 0x93, 0x0e, 0x69, 0x6c, 0x5e,
	0x3e, 0xcd, 0x2f, 0x16, 0x94, 0xff, 0xb3, 0x2d, 0xb6, 0xa7, 0x36, 0x9e, 0x3d, 0xb4, 0x31, 0xeb,
	0xdb, 0x78, 0x26, 0xc6, 0x21, 0x7e, 0x0d, 0xab, 0xca, 0x1d, 0xfd, 0x3d, 0x83, 0x30, 0x13, 0xe7,
	0xcc, 0x33, 0x58, 0xc9, 0x4f, 0x72, 0xef, 0xac, 0xac, 0x48, 0x79, 0x5e, 0x6c, 0xfb, 0x36, 0x3e,
	0x33, 0x71, 0xee, 0x78, 0x0a, 0x81, 0x32, 0xc6, 0xc9, 0x8a, 0xac, 0xe7, 0x05, 0xf4, 0x6d, 0x3c,
	0x12, 0x31, 0x56, 0xfe, 0x04, 0x02, 0x74, 0xf6, 0xe0, 0x26, 0xe5, 0x31, 0x0d, 0xb3, 0xc6, 0x96,
	0xd6, 0x19, 0xf2, 0x5c, 0xfa, 0x59, 0x4f, 0xc4, 0x58, 0xd3, 0x37, 0xb0, 0x1a, 0x4c, 0xde, 0x2b,
	0x43, 0x0b, 0x1e, 0xbe, 0x83, 0x09, 0x59, 0x32, 0xcf, 0xd6, 0x7e, 0xc1, 0x04, 0x84, 0x2f, 0xc3,
	0xb3, 0x0e, 0xf2, 0xe3, 0xfb, 0x51, 0x5f, 0x3c, 0xfb, 0xf3, 0x33, 0x62, 0xdf, 0xba, 0x88, 0x7d,
	0xef, 0x22, 0x76, 0xea, 0x22, 0xf6, 0xd0, 0x45, 0xec, 0x47, 0x17, 0xb1, 0xaf, 0xbf, 0xa2, 0xd9,
	0x87, 0x25, 0x2d, 0x6b, 0x17, 0xd0, 0x2f, 0x79, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x56, 0x2f,
	0x75, 0xdd, 0x6f, 0x02, 0x00, 0x00,
}
