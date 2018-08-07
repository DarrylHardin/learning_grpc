// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc/calcpb/calc.proto

package calcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Calc struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	TotalNumber          int32    `protobuf:"varint,3,opt,name=total_number,json=totalNumber,proto3" json:"total_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calc) Reset()         { *m = Calc{} }
func (m *Calc) String() string { return proto.CompactTextString(m) }
func (*Calc) ProtoMessage()    {}
func (*Calc) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_5ace64bde9aac670, []int{0}
}
func (m *Calc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calc.Unmarshal(m, b)
}
func (m *Calc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calc.Marshal(b, m, deterministic)
}
func (dst *Calc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calc.Merge(dst, src)
}
func (m *Calc) XXX_Size() int {
	return xxx_messageInfo_Calc.Size(m)
}
func (m *Calc) XXX_DiscardUnknown() {
	xxx_messageInfo_Calc.DiscardUnknown(m)
}

var xxx_messageInfo_Calc proto.InternalMessageInfo

func (m *Calc) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Calc) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

func (m *Calc) GetTotalNumber() int32 {
	if m != nil {
		return m.TotalNumber
	}
	return 0
}

type CalcRequest struct {
	Calc                 *Calc    `protobuf:"bytes,1,opt,name=calc,proto3" json:"calc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_5ace64bde9aac670, []int{1}
}
func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (dst *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(dst, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetCalc() *Calc {
	if m != nil {
		return m.Calc
	}
	return nil
}

type CalcResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Answer               bool     `protobuf:"varint,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_5ace64bde9aac670, []int{2}
}
func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (dst *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(dst, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CalcResponse) GetAnswer() bool {
	if m != nil {
		return m.Answer
	}
	return false
}

type PrimeNumbersDecompRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumbersDecompRequest) Reset()         { *m = PrimeNumbersDecompRequest{} }
func (m *PrimeNumbersDecompRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumbersDecompRequest) ProtoMessage()    {}
func (*PrimeNumbersDecompRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_5ace64bde9aac670, []int{3}
}
func (m *PrimeNumbersDecompRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumbersDecompRequest.Unmarshal(m, b)
}
func (m *PrimeNumbersDecompRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumbersDecompRequest.Marshal(b, m, deterministic)
}
func (dst *PrimeNumbersDecompRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumbersDecompRequest.Merge(dst, src)
}
func (m *PrimeNumbersDecompRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumbersDecompRequest.Size(m)
}
func (m *PrimeNumbersDecompRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumbersDecompRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumbersDecompRequest proto.InternalMessageInfo

func (m *PrimeNumbersDecompRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumbersDecompResponse struct {
	Answer               int32    `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumbersDecompResponse) Reset()         { *m = PrimeNumbersDecompResponse{} }
func (m *PrimeNumbersDecompResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumbersDecompResponse) ProtoMessage()    {}
func (*PrimeNumbersDecompResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calc_5ace64bde9aac670, []int{4}
}
func (m *PrimeNumbersDecompResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumbersDecompResponse.Unmarshal(m, b)
}
func (m *PrimeNumbersDecompResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumbersDecompResponse.Marshal(b, m, deterministic)
}
func (dst *PrimeNumbersDecompResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumbersDecompResponse.Merge(dst, src)
}
func (m *PrimeNumbersDecompResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumbersDecompResponse.Size(m)
}
func (m *PrimeNumbersDecompResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumbersDecompResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumbersDecompResponse proto.InternalMessageInfo

func (m *PrimeNumbersDecompResponse) GetAnswer() int32 {
	if m != nil {
		return m.Answer
	}
	return 0
}

func init() {
	proto.RegisterType((*Calc)(nil), "calc.Calc")
	proto.RegisterType((*CalcRequest)(nil), "calc.CalcRequest")
	proto.RegisterType((*CalcResponse)(nil), "calc.CalcResponse")
	proto.RegisterType((*PrimeNumbersDecompRequest)(nil), "calc.PrimeNumbersDecompRequest")
	proto.RegisterType((*PrimeNumbersDecompResponse)(nil), "calc.PrimeNumbersDecompResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	PrimeNumbersDecomp(ctx context.Context, in *PrimeNumbersDecompRequest, opts ...grpc.CallOption) (CalcService_PrimeNumbersDecompClient, error)
}

type calcServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalcServiceClient(cc *grpc.ClientConn) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeNumbersDecomp(ctx context.Context, in *PrimeNumbersDecompRequest, opts ...grpc.CallOption) (CalcService_PrimeNumbersDecompClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calc.CalcService/PrimeNumbersDecomp", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeNumbersDecompClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeNumbersDecompClient interface {
	Recv() (*PrimeNumbersDecompResponse, error)
	grpc.ClientStream
}

type calcServicePrimeNumbersDecompClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeNumbersDecompClient) Recv() (*PrimeNumbersDecompResponse, error) {
	m := new(PrimeNumbersDecompResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
	PrimeNumbersDecomp(*PrimeNumbersDecompRequest, CalcService_PrimeNumbersDecompServer) error
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeNumbersDecomp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumbersDecompRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeNumbersDecomp(m, &calcServicePrimeNumbersDecompServer{stream})
}

type CalcService_PrimeNumbersDecompServer interface {
	Send(*PrimeNumbersDecompResponse) error
	grpc.ServerStream
}

type calcServicePrimeNumbersDecompServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeNumbersDecompServer) Send(m *PrimeNumbersDecompResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumbersDecomp",
			Handler:       _CalcService_PrimeNumbersDecomp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calc/calcpb/calc.proto",
}

func init() { proto.RegisterFile("calc/calcpb/calc.proto", fileDescriptor_calc_5ace64bde9aac670) }

var fileDescriptor_calc_5ace64bde9aac670 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x4f, 0xbc, 0x40,
	0x10, 0xc5, 0xff, 0xfc, 0x3d, 0xc9, 0x65, 0xc0, 0xc2, 0x29, 0x88, 0x52, 0xe8, 0x1d, 0x36, 0x36,
	0xde, 0x99, 0x3b, 0x6b, 0x0b, 0xb5, 0x36, 0x06, 0x3b, 0x2d, 0x0c, 0xac, 0x63, 0x42, 0x02, 0x2c,
	0xb7, 0xbb, 0xe8, 0x87, 0xf1, 0xcb, 0x1a, 0x66, 0x96, 0x78, 0x46, 0xaf, 0x01, 0xe6, 0xe5, 0xb7,
	0xf3, 0xde, 0x23, 0x0b, 0x89, 0x2a, 0x6a, 0xb5, 0x1c, 0x1e, 0x5d, 0xc9, 0xaf, 0x45, 0x67, 0xb4,
	0xd3, 0x38, 0x19, 0xbe, 0xb3, 0x0d, 0x4c, 0x6e, 0x8b, 0x5a, 0xe1, 0x1c, 0xe2, 0xb7, 0xca, 0x58,
	0xf7, 0xd2, 0xf6, 0x4d, 0x49, 0xe6, 0x28, 0x98, 0x05, 0xe7, 0xfb, 0x79, 0xc4, 0xda, 0x3d, 0x4b,
	0x78, 0x06, 0x07, 0x96, 0x94, 0x6e, 0x5f, 0x47, 0xe6, 0x3f, 0x33, 0xb1, 0x88, 0x1e, 0x9a, 0x43,
	0xec, 0xb4, 0x2b, 0xea, 0x91, 0xd9, 0x93, 0x3d, 0xac, 0x09, 0x92, 0x5d, 0x40, 0x34, 0x58, 0xe6,
	0xb4, 0xe9, 0xc9, 0x3a, 0x3c, 0x01, 0x4e, 0xc2, 0x8e, 0xd1, 0x0a, 0x16, 0x1c, 0x91, 0x01, 0x49,
	0x78, 0x0d, 0xb1, 0xe0, 0xb6, 0xd3, 0xad, 0x25, 0x4c, 0x20, 0x34, 0x64, 0xfb, 0xda, 0xf9, 0x8c,
	0x7e, 0x1a, 0xf4, 0xa2, 0xb5, 0x1f, 0x3e, 0xd7, 0x34, 0xf7, 0x53, 0xb6, 0x86, 0xe3, 0x07, 0x53,
	0x35, 0x24, 0xee, 0xf6, 0x8e, 0x94, 0x6e, 0xba, 0xd1, 0x3c, 0x81, 0xf0, 0x47, 0x61, 0x3f, 0x65,
	0x57, 0x90, 0xfe, 0x75, 0xe8, 0x3b, 0x82, 0xb7, 0xf2, 0xa7, 0x64, 0x5a, 0x7d, 0x06, 0x52, 0xed,
	0x91, 0xcc, 0x7b, 0xa5, 0x08, 0x97, 0xfe, 0xe7, 0x1e, 0x6e, 0x95, 0x12, 0xe3, 0x14, 0xb7, 0x25,
	0x59, 0x9b, 0xfd, 0xc3, 0x67, 0xc0, 0xdf, 0xb6, 0x78, 0x2a, 0xec, 0xce, 0x16, 0xe9, 0x6c, 0x37,
	0x30, 0xae, 0xbe, 0x0c, 0x6e, 0xa6, 0x4f, 0xa1, 0xdc, 0x82, 0x32, 0xe4, 0x1b, 0xb0, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x0a, 0x34, 0x72, 0x1b, 0x02, 0x00, 0x00,
}
