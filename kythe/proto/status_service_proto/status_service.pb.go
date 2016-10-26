// Code generated by protoc-gen-gogo.
// source: kythe/proto/status_service.proto
// DO NOT EDIT!

/*
	Package status_service_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/status_service.proto

	It has these top-level messages:
		StatusRequest
		StatusReply
*/
package status_service_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type StatusReply_Language_Support int32

const (
	StatusReply_Language_UNSUPPORTED  StatusReply_Language_Support = 0
	StatusReply_Language_EXPERIMENTAL StatusReply_Language_Support = 1
	StatusReply_Language_SUPPORTED    StatusReply_Language_Support = 2
)

var StatusReply_Language_Support_name = map[int32]string{
	0: "UNSUPPORTED",
	1: "EXPERIMENTAL",
	2: "SUPPORTED",
}
var StatusReply_Language_Support_value = map[string]int32{
	"UNSUPPORTED":  0,
	"EXPERIMENTAL": 1,
	"SUPPORTED":    2,
}

func (x StatusReply_Language_Support) String() string {
	return proto.EnumName(StatusReply_Language_Support_name, int32(x))
}
func (StatusReply_Language_Support) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorStatusService, []int{1, 1, 0}
}

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptorStatusService, []int{0} }

type StatusReply struct {
	// The origins of the data served by this endpoint.
	Origins []*StatusReply_Origin `protobuf:"bytes,1,rep,name=origins" json:"origins,omitempty"`
	// The languages supported by this endpoint.
	Languages []*StatusReply_Language `protobuf:"bytes,2,rep,name=languages" json:"languages,omitempty"`
}

func (m *StatusReply) Reset()                    { *m = StatusReply{} }
func (m *StatusReply) String() string            { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()               {}
func (*StatusReply) Descriptor() ([]byte, []int) { return fileDescriptorStatusService, []int{1} }

func (m *StatusReply) GetOrigins() []*StatusReply_Origin {
	if m != nil {
		return m.Origins
	}
	return nil
}

func (m *StatusReply) GetLanguages() []*StatusReply_Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

type StatusReply_Origin struct {
	Corpus   string `protobuf:"bytes,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (m *StatusReply_Origin) Reset()         { *m = StatusReply_Origin{} }
func (m *StatusReply_Origin) String() string { return proto.CompactTextString(m) }
func (*StatusReply_Origin) ProtoMessage()    {}
func (*StatusReply_Origin) Descriptor() ([]byte, []int) {
	return fileDescriptorStatusService, []int{1, 0}
}

type StatusReply_Language struct {
	Name    string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Support StatusReply_Language_Support `protobuf:"varint,2,opt,name=support,proto3,enum=kythe.proto.StatusReply_Language_Support" json:"support,omitempty"`
}

func (m *StatusReply_Language) Reset()         { *m = StatusReply_Language{} }
func (m *StatusReply_Language) String() string { return proto.CompactTextString(m) }
func (*StatusReply_Language) ProtoMessage()    {}
func (*StatusReply_Language) Descriptor() ([]byte, []int) {
	return fileDescriptorStatusService, []int{1, 1}
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "kythe.proto.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "kythe.proto.StatusReply")
	proto.RegisterType((*StatusReply_Origin)(nil), "kythe.proto.StatusReply.Origin")
	proto.RegisterType((*StatusReply_Language)(nil), "kythe.proto.StatusReply.Language")
	proto.RegisterEnum("kythe.proto.StatusReply_Language_Support", StatusReply_Language_Support_name, StatusReply_Language_Support_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for StatusService service

type StatusServiceClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type statusServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatusServiceClient(cc *grpc.ClientConn) StatusServiceClient {
	return &statusServiceClient{cc}
}

func (c *statusServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := grpc.Invoke(ctx, "/kythe.proto.StatusService/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StatusService service

type StatusServiceServer interface {
	Status(context.Context, *StatusRequest) (*StatusReply, error)
}

func RegisterStatusServiceServer(s *grpc.Server, srv StatusServiceServer) {
	s.RegisterService(&_StatusService_serviceDesc, srv)
}

func _StatusService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.StatusService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.StatusService",
	HandlerType: (*StatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _StatusService_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *StatusRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *StatusReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusReply) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Origins) > 0 {
		for _, msg := range m.Origins {
			data[i] = 0xa
			i++
			i = encodeVarintStatusService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Languages) > 0 {
		for _, msg := range m.Languages {
			data[i] = 0x12
			i++
			i = encodeVarintStatusService(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StatusReply_Origin) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusReply_Origin) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Corpus) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStatusService(data, i, uint64(len(m.Corpus)))
		i += copy(data[i:], m.Corpus)
	}
	if len(m.Revision) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintStatusService(data, i, uint64(len(m.Revision)))
		i += copy(data[i:], m.Revision)
	}
	return i, nil
}

func (m *StatusReply_Language) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusReply_Language) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStatusService(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if m.Support != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintStatusService(data, i, uint64(m.Support))
	}
	return i, nil
}

func encodeFixed64StatusService(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32StatusService(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStatusService(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StatusRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *StatusReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Origins) > 0 {
		for _, e := range m.Origins {
			l = e.Size()
			n += 1 + l + sovStatusService(uint64(l))
		}
	}
	if len(m.Languages) > 0 {
		for _, e := range m.Languages {
			l = e.Size()
			n += 1 + l + sovStatusService(uint64(l))
		}
	}
	return n
}

func (m *StatusReply_Origin) Size() (n int) {
	var l int
	_ = l
	l = len(m.Corpus)
	if l > 0 {
		n += 1 + l + sovStatusService(uint64(l))
	}
	l = len(m.Revision)
	if l > 0 {
		n += 1 + l + sovStatusService(uint64(l))
	}
	return n
}

func (m *StatusReply_Language) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStatusService(uint64(l))
	}
	if m.Support != 0 {
		n += 1 + sovStatusService(uint64(m.Support))
	}
	return n
}

func sovStatusService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatusService(x uint64) (n int) {
	return sovStatusService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatusRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatusService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStatusService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatusService
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
func (m *StatusReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatusService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatusService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origins = append(m.Origins, &StatusReply_Origin{})
			if err := m.Origins[len(m.Origins)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Languages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatusService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Languages = append(m.Languages, &StatusReply_Language{})
			if err := m.Languages[len(m.Languages)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatusService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatusService
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
func (m *StatusReply_Origin) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatusService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Origin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Origin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Corpus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatusService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Corpus = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatusService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Revision = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatusService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatusService
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
func (m *StatusReply_Language) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatusService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Language: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Language: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatusService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Support", wireType)
			}
			m.Support = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Support |= (StatusReply_Language_Support(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStatusService(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatusService
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
func skipStatusService(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatusService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowStatusService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStatusService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatusService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipStatusService(data[start:])
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
	ErrInvalidLengthStatusService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatusService   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorStatusService = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc2, 0x50,
	0x14, 0xc6, 0xb9, 0xd5, 0x14, 0x38, 0x15, 0x69, 0xce, 0x60, 0x9a, 0x0e, 0x15, 0x3b, 0xe1, 0x52,
	0x12, 0x9c, 0x8c, 0x26, 0xfe, 0xed, 0x60, 0x82, 0x80, 0xb7, 0x90, 0xb8, 0x99, 0x4a, 0x6e, 0xb0,
	0x11, 0xdb, 0xda, 0x7b, 0x4b, 0xc2, 0x9b, 0x38, 0xfa, 0x38, 0x8e, 0x3e, 0x82, 0xa9, 0x9b, 0x4f,
	0x61, 0xbc, 0x6d, 0x85, 0x41, 0xe2, 0xd6, 0xf3, 0xf5, 0xf7, 0x3b, 0xf9, 0x72, 0x0f, 0xb4, 0x1e,
	0x17, 0xe2, 0x81, 0x75, 0xe2, 0x24, 0x12, 0x51, 0x87, 0x0b, 0x5f, 0xa4, 0xfc, 0x8e, 0xb3, 0x64,
	0x1e, 0x4c, 0x98, 0x23, 0x43, 0xd4, 0x24, 0x91, 0x0f, 0x76, 0x13, 0x1a, 0x9e, 0x84, 0x28, 0x7b,
	0x4e, 0x19, 0x17, 0xf6, 0x97, 0x02, 0x5a, 0x99, 0xc4, 0xb3, 0x05, 0x1e, 0x42, 0x35, 0x4a, 0x82,
	0x69, 0x10, 0x72, 0x83, 0xb4, 0x36, 0xda, 0x5a, 0x77, 0xd7, 0x59, 0xf1, 0x9d, 0x15, 0xd4, 0x19,
	0x48, 0x8e, 0x96, 0x3c, 0x9e, 0x40, 0x7d, 0xe6, 0x87, 0xd3, 0xd4, 0x9f, 0x32, 0x6e, 0x28, 0x52,
	0xde, 0x5b, 0x2b, 0xf7, 0x0a, 0x92, 0x2e, 0x1d, 0xf3, 0x18, 0xd4, 0x7c, 0x27, 0xee, 0x80, 0x3a,
	0x89, 0x92, 0x38, 0xfd, 0x29, 0x41, 0xda, 0x75, 0x5a, 0x4c, 0x68, 0x42, 0x2d, 0x61, 0xf3, 0x80,
	0x07, 0x51, 0x68, 0x28, 0xf2, 0xcf, 0xef, 0x6c, 0xbe, 0x12, 0xa8, 0x95, 0x5b, 0x11, 0x61, 0x33,
	0xf4, 0x9f, 0x58, 0xa1, 0xcb, 0x6f, 0xbc, 0x80, 0x2a, 0x4f, 0xe3, 0x38, 0x4a, 0x84, 0x74, 0xb7,
	0xbb, 0xfb, 0xff, 0xb6, 0x73, 0xbc, 0x5c, 0xa0, 0xa5, 0x69, 0x1f, 0x41, 0xb5, 0xc8, 0xb0, 0x09,
	0xda, 0xb8, 0xef, 0x8d, 0x87, 0xc3, 0x01, 0x1d, 0xb9, 0x97, 0x7a, 0x05, 0x75, 0xd8, 0x72, 0x6f,
	0x87, 0x2e, 0xbd, 0xba, 0x76, 0xfb, 0xa3, 0xb3, 0x9e, 0x4e, 0xb0, 0x01, 0xf5, 0x25, 0xa0, 0x74,
	0x6f, 0xca, 0xd7, 0xf7, 0xf2, 0x0b, 0xe1, 0x29, 0xa8, 0x79, 0x80, 0xe6, 0x9f, 0x5d, 0xe4, 0x8d,
	0x4c, 0x63, 0x5d, 0x4f, 0xbb, 0x72, 0xae, 0xbf, 0x65, 0x16, 0x79, 0xcf, 0x2c, 0xf2, 0x91, 0x59,
	0xe4, 0xe5, 0xd3, 0xaa, 0xdc, 0xab, 0x12, 0x3b, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x27, 0x71,
	0xa0, 0xcf, 0x1a, 0x02, 0x00, 0x00,
}