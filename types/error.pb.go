// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: error.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Error describes an error captured while processing an event in Sensu's
// pipeline.
type Error struct {
	// Name is the unique identifier for an asset
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Environment indicates to which env the error is associated with
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// Organization indicates to which org the error is associated with
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	// Message is the details associated with the error
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Component refers to the component in-which the error occurred.
	Component string `protobuf:"bytes,5,opt,name=component,proto3" json:"component,omitempty"`
	// Event refers to the event that was being processed when the error occurred.
	Event Event `protobuf:"bytes,6,opt,name=event" json:"event"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorError, []int{0} }

func (m *Error) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Error) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Error) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *Error) GetEvent() Event {
	if m != nil {
		return m.Event
	}
	return Event{}
}

func init() {
	proto.RegisterType((*Error)(nil), "sensu.types.Error")
}
func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Environment != that1.Environment {
		return false
	}
	if this.Organization != that1.Organization {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.Component != that1.Component {
		return false
	}
	if !this.Event.Equal(&that1.Event) {
		return false
	}
	return true
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Environment) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Environment)))
		i += copy(dAtA[i:], m.Environment)
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Component) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintError(dAtA, i, uint64(len(m.Component)))
		i += copy(dAtA[i:], m.Component)
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintError(dAtA, i, uint64(m.Event.Size()))
	n1, err := m.Event.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedError(r randyError, easy bool) *Error {
	this := &Error{}
	this.Name = string(randStringError(r))
	this.Environment = string(randStringError(r))
	this.Organization = string(randStringError(r))
	this.Message = string(randStringError(r))
	this.Component = string(randStringError(r))
	v1 := NewPopulatedEvent(r, easy)
	this.Event = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyError interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneError(r randyError) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringError(r randyError) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneError(r)
	}
	return string(tmps)
}
func randUnrecognizedError(r randyError, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldError(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldError(dAtA []byte, r randyError, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateError(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateError(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateError(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateError(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Error) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Environment)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = len(m.Component)
	if l > 0 {
		n += 1 + l + sovError(uint64(l))
	}
	l = m.Event.Size()
	n += 1 + l + sovError(uint64(l))
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Environment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Component = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
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
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
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
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowError
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
				next, err := skipError(dAtA[start:])
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
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("error.proto", fileDescriptorError) }

var fileDescriptorError = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x3b, 0xff, 0x9f, 0x14, 0xd5, 0x61, 0xe5, 0x95, 0x55, 0x21, 0x13, 0x95, 0x4d, 0x37,
	0xb8, 0x12, 0xdc, 0xa0, 0x52, 0x2f, 0xd0, 0x25, 0xbb, 0xa4, 0x1a, 0x42, 0x16, 0xf6, 0x44, 0xb6,
	0x53, 0x09, 0x4e, 0xc2, 0x11, 0x38, 0x02, 0x47, 0x28, 0x3b, 0x4e, 0x80, 0xc0, 0x5c, 0x82, 0x25,
	0xca, 0x44, 0x88, 0xb2, 0x9b, 0xf7, 0xbe, 0x37, 0xa3, 0x79, 0xa2, 0x40, 0xef, 0xc9, 0x9b, 0xce,
	0x53, 0x24, 0x59, 0x04, 0x74, 0xa1, 0x37, 0xf1, 0xbe, 0xc3, 0x30, 0xbf, 0x6c, 0xda, 0x78, 0xd7,
	0xd7, 0x66, 0x47, 0x76, 0xd5, 0x50, 0x43, 0x2b, 0xce, 0xd4, 0xfd, 0x2d, 0x2b, 0x16, 0x3c, 0x8d,
	0xbb, 0xf3, 0x02, 0xf7, 0xe8, 0xe2, 0x28, 0x16, 0x2f, 0x20, 0xf2, 0xcd, 0x70, 0x58, 0x4a, 0x91,
	0xb9, 0xca, 0xa2, 0x82, 0x12, 0x96, 0xb3, 0x2d, 0xcf, 0xb2, 0x14, 0x05, 0xba, 0x7d, 0xeb, 0xc9,
	0x59, 0x74, 0x51, 0xfd, 0x63, 0x74, 0x6c, 0xc9, 0x85, 0x38, 0x25, 0xdf, 0x54, 0xae, 0x7d, 0xa8,
	0x62, 0x4b, 0x4e, 0xfd, 0xe7, 0xc8, 0x1f, 0x4f, 0x2a, 0x71, 0x62, 0x31, 0x84, 0xaa, 0x41, 0x95,
	0x31, 0xfe, 0x91, 0xf2, 0x4c, 0xcc, 0x76, 0x64, 0x3b, 0x72, 0xc3, 0xf5, 0x9c, 0xd9, 0xaf, 0x21,
	0x8d, 0xc8, 0xf9, 0x55, 0x35, 0x2d, 0x61, 0x59, 0x5c, 0x49, 0x73, 0x54, 0xda, 0x6c, 0x06, 0xb2,
	0xce, 0x0e, 0x6f, 0xe7, 0x93, 0xed, 0x18, 0x5b, 0x5f, 0x7c, 0x7d, 0x68, 0x78, 0x4a, 0x1a, 0x9e,
	0x93, 0x86, 0x43, 0xd2, 0xf0, 0x9a, 0x34, 0xbc, 0x27, 0x0d, 0x8f, 0x9f, 0x7a, 0x72, 0x93, 0xf3,
	0x5e, 0x3d, 0xe5, 0xde, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0xe6, 0x99, 0xf8, 0x4f,
	0x01, 0x00, 0x00,
}