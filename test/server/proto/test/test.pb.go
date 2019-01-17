// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test/test.proto

package test // import "github.com/johanbrandhorst/protobuf/test/server/proto/test"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/johanbrandhorst/protobuf/proto"

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

type PingRequest_FailureType int32

const (
	PingRequest_NONE PingRequest_FailureType = 0
	PingRequest_CODE PingRequest_FailureType = 1
	PingRequest_DROP PingRequest_FailureType = 2
)

var PingRequest_FailureType_name = map[int32]string{
	0: "NONE",
	1: "CODE",
	2: "DROP",
}
var PingRequest_FailureType_value = map[string]int32{
	"NONE": 0,
	"CODE": 1,
	"DROP": 2,
}

func (x PingRequest_FailureType) String() string {
	return proto.EnumName(PingRequest_FailureType_name, int32(x))
}
func (PingRequest_FailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_c175bf3ee4cc4a47, []int{1, 0}
}

type ExtraStuff struct {
	Addresses map[int32]string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Title:
	//	*ExtraStuff_FirstName
	//	*ExtraStuff_IdNumber
	Title                isExtraStuff_Title `protobuf_oneof:"title"`
	CardNumbers          []uint32           `protobuf:"varint,4,rep,packed,name=card_numbers,json=cardNumbers" json:"card_numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExtraStuff) Reset()         { *m = ExtraStuff{} }
func (m *ExtraStuff) String() string { return proto.CompactTextString(m) }
func (*ExtraStuff) ProtoMessage()    {}
func (*ExtraStuff) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_c175bf3ee4cc4a47, []int{0}
}
func (m *ExtraStuff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraStuff.Unmarshal(m, b)
}
func (m *ExtraStuff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraStuff.Marshal(b, m, deterministic)
}
func (dst *ExtraStuff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraStuff.Merge(dst, src)
}
func (m *ExtraStuff) XXX_Size() int {
	return xxx_messageInfo_ExtraStuff.Size(m)
}
func (m *ExtraStuff) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraStuff.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraStuff proto.InternalMessageInfo

type isExtraStuff_Title interface {
	isExtraStuff_Title()
}

type ExtraStuff_FirstName struct {
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,oneof"`
}
type ExtraStuff_IdNumber struct {
	IdNumber int32 `protobuf:"varint,3,opt,name=id_number,json=idNumber,oneof"`
}

func (*ExtraStuff_FirstName) isExtraStuff_Title() {}
func (*ExtraStuff_IdNumber) isExtraStuff_Title()  {}

func (m *ExtraStuff) GetTitle() isExtraStuff_Title {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *ExtraStuff) GetAddresses() map[int32]string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ExtraStuff) GetFirstName() string {
	if x, ok := m.GetTitle().(*ExtraStuff_FirstName); ok {
		return x.FirstName
	}
	return ""
}

func (m *ExtraStuff) GetIdNumber() int32 {
	if x, ok := m.GetTitle().(*ExtraStuff_IdNumber); ok {
		return x.IdNumber
	}
	return 0
}

func (m *ExtraStuff) GetCardNumbers() []uint32 {
	if m != nil {
		return m.CardNumbers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExtraStuff) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExtraStuff_OneofMarshaler, _ExtraStuff_OneofUnmarshaler, _ExtraStuff_OneofSizer, []interface{}{
		(*ExtraStuff_FirstName)(nil),
		(*ExtraStuff_IdNumber)(nil),
	}
}

func _ExtraStuff_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FirstName)
	case *ExtraStuff_IdNumber:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IdNumber))
	case nil:
	default:
		return fmt.Errorf("ExtraStuff.Title has unexpected type %T", x)
	}
	return nil
}

func _ExtraStuff_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExtraStuff)
	switch tag {
	case 2: // title.first_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Title = &ExtraStuff_FirstName{x}
		return true, err
	case 3: // title.id_number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Title = &ExtraStuff_IdNumber{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ExtraStuff_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExtraStuff)
	// title
	switch x := m.Title.(type) {
	case *ExtraStuff_FirstName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.FirstName)))
		n += len(x.FirstName)
	case *ExtraStuff_IdNumber:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IdNumber))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PingRequest struct {
	Value                string                  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	ResponseCount        int32                   `protobuf:"varint,2,opt,name=response_count,json=responseCount" json:"response_count,omitempty"`
	ErrorCodeReturned    uint32                  `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
	FailureType          PingRequest_FailureType `protobuf:"varint,4,opt,name=failure_type,json=failureType,enum=test.PingRequest_FailureType" json:"failure_type,omitempty"`
	CheckMetadata        bool                    `protobuf:"varint,5,opt,name=check_metadata,json=checkMetadata" json:"check_metadata,omitempty"`
	SendHeaders          bool                    `protobuf:"varint,6,opt,name=send_headers,json=sendHeaders" json:"send_headers,omitempty"`
	SendTrailers         bool                    `protobuf:"varint,7,opt,name=send_trailers,json=sendTrailers" json:"send_trailers,omitempty"`
	MessageLatencyMs     int32                   `protobuf:"varint,8,opt,name=message_latency_ms,json=messageLatencyMs" json:"message_latency_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_c175bf3ee4cc4a47, []int{1}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

func (m *PingRequest) GetFailureType() PingRequest_FailureType {
	if m != nil {
		return m.FailureType
	}
	return PingRequest_NONE
}

func (m *PingRequest) GetCheckMetadata() bool {
	if m != nil {
		return m.CheckMetadata
	}
	return false
}

func (m *PingRequest) GetSendHeaders() bool {
	if m != nil {
		return m.SendHeaders
	}
	return false
}

func (m *PingRequest) GetSendTrailers() bool {
	if m != nil {
		return m.SendTrailers
	}
	return false
}

func (m *PingRequest) GetMessageLatencyMs() int32 {
	if m != nil {
		return m.MessageLatencyMs
	}
	return 0
}

type PingResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
	Counter              int32    `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_c175bf3ee4cc4a47, []int{2}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtraStuff)(nil), "test.ExtraStuff")
	proto.RegisterMapType((map[int32]string)(nil), "test.ExtraStuff.AddressesEntry")
	proto.RegisterType((*PingRequest)(nil), "test.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "test.PingResponse")
	proto.RegisterEnum("test.PingRequest_FailureType", PingRequest_FailureType_name, PingRequest_FailureType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
	PingClientStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingClientStreamClient, error)
	PingClientStreamError(ctx context.Context, opts ...grpc.CallOption) (TestService_PingClientStreamErrorClient, error)
	PingBidiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingBidiStreamClient, error)
	PingBidiStreamError(ctx context.Context, opts ...grpc.CallOption) (TestService_PingBidiStreamErrorClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/test.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/test.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/test.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingClientStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/test.TestService/PingClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingClientStreamClient{stream}
	return x, nil
}

type TestService_PingClientStreamClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingClientStreamClient struct {
	grpc.ClientStream
}

func (x *testServicePingClientStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingClientStreamClient) CloseAndRecv() (*PingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingClientStreamError(ctx context.Context, opts ...grpc.CallOption) (TestService_PingClientStreamErrorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/test.TestService/PingClientStreamError", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingClientStreamErrorClient{stream}
	return x, nil
}

type TestService_PingClientStreamErrorClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingClientStreamErrorClient struct {
	grpc.ClientStream
}

func (x *testServicePingClientStreamErrorClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingClientStreamErrorClient) CloseAndRecv() (*PingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingBidiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_PingBidiStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[3], c.cc, "/test.TestService/PingBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingBidiStreamClient{stream}
	return x, nil
}

type TestService_PingBidiStreamClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingBidiStreamClient struct {
	grpc.ClientStream
}

func (x *testServicePingBidiStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingBidiStreamClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) PingBidiStreamError(ctx context.Context, opts ...grpc.CallOption) (TestService_PingBidiStreamErrorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[4], c.cc, "/test.TestService/PingBidiStreamError", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingBidiStreamErrorClient{stream}
	return x, nil
}

type TestService_PingBidiStreamErrorClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingBidiStreamErrorClient struct {
	grpc.ClientStream
}

func (x *testServicePingBidiStreamErrorClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServicePingBidiStreamErrorClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *empty.Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*empty.Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
	PingClientStream(TestService_PingClientStreamServer) error
	PingClientStreamError(TestService_PingClientStreamErrorServer) error
	PingBidiStream(TestService_PingBidiStreamServer) error
	PingBidiStreamError(TestService_PingBidiStreamErrorServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_PingClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingClientStream(&testServicePingClientStreamServer{stream})
}

type TestService_PingClientStreamServer interface {
	SendAndClose(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingClientStreamServer struct {
	grpc.ServerStream
}

func (x *testServicePingClientStreamServer) SendAndClose(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingClientStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_PingClientStreamError_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingClientStreamError(&testServicePingClientStreamErrorServer{stream})
}

type TestService_PingClientStreamErrorServer interface {
	SendAndClose(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingClientStreamErrorServer struct {
	grpc.ServerStream
}

func (x *testServicePingClientStreamErrorServer) SendAndClose(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingClientStreamErrorServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_PingBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingBidiStream(&testServicePingBidiStreamServer{stream})
}

type TestService_PingBidiStreamServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingBidiStreamServer struct {
	grpc.ServerStream
}

func (x *testServicePingBidiStreamServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingBidiStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_PingBidiStreamError_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).PingBidiStreamError(&testServicePingBidiStreamErrorServer{stream})
}

type TestService_PingBidiStreamErrorServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testServicePingBidiStreamErrorServer struct {
	grpc.ServerStream
}

func (x *testServicePingBidiStreamErrorServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServicePingBidiStreamErrorServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingClientStream",
			Handler:       _TestService_PingClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PingClientStreamError",
			Handler:       _TestService_PingClientStreamError_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PingBidiStream",
			Handler:       _TestService_PingBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PingBidiStreamError",
			Handler:       _TestService_PingBidiStreamError_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/test/test.proto",
}

func init() { proto.RegisterFile("proto/test/test.proto", fileDescriptor_test_c175bf3ee4cc4a47) }

var fileDescriptor_test_c175bf3ee4cc4a47 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0xc5, 0x01, 0x12, 0x18, 0x07, 0x44, 0x36, 0xbf, 0xfc, 0x64, 0x51, 0x45, 0xa1, 0x54, 0x95,
	0x2c, 0xb5, 0x32, 0x51, 0xaa, 0x4a, 0x69, 0xd4, 0x54, 0x29, 0x84, 0x2a, 0x87, 0x84, 0x44, 0x4e,
	0xd4, 0x43, 0x2f, 0xd6, 0x62, 0x0f, 0xe0, 0xc4, 0x7f, 0xe8, 0xee, 0x1a, 0xd5, 0x5f, 0xa0, 0x9f,
	0xb7, 0x97, 0x5e, 0x7a, 0xaa, 0x76, 0x6d, 0x04, 0x69, 0x8b, 0x44, 0x2e, 0xd6, 0xcc, 0xdb, 0xf7,
	0x66, 0x9e, 0x67, 0xb4, 0x0b, 0x7b, 0x53, 0x16, 0x8b, 0xb8, 0x23, 0x90, 0x0b, 0xf5, 0xb1, 0x54,
	0x4e, 0x4a, 0x32, 0x6e, 0x3e, 0x1b, 0xc7, 0xf1, 0x38, 0xc0, 0x8e, 0xc2, 0x86, 0xc9, 0xa8, 0x83,
	0xe1, 0x54, 0xa4, 0x19, 0xa5, 0x79, 0x3c, 0xf6, 0xc5, 0x24, 0x19, 0x5a, 0x6e, 0x1c, 0x76, 0xee,
	0xe3, 0x09, 0x8d, 0x86, 0x8c, 0x46, 0xde, 0x24, 0x66, 0x5c, 0x2c, 0x04, 0x59, 0xf5, 0x71, 0x3c,
	0x9d, 0x20, 0xbb, 0xe7, 0x99, 0xb2, 0xfd, 0x53, 0x03, 0xe8, 0x7f, 0x13, 0x8c, 0xde, 0x8a, 0x64,
	0x34, 0x22, 0xa7, 0x50, 0xa5, 0x9e, 0xc7, 0x90, 0x73, 0xe4, 0x86, 0xd6, 0x2a, 0x9a, 0xfa, 0xd1,
	0x81, 0xa5, 0xbc, 0x2c, 0x48, 0xd6, 0xc7, 0x39, 0xa3, 0x1f, 0x09, 0x96, 0xda, 0x0b, 0x05, 0x39,
	0x00, 0x18, 0xf9, 0x8c, 0x0b, 0x27, 0xa2, 0x21, 0x1a, 0x1b, 0x2d, 0xcd, 0xac, 0x5e, 0x14, 0xec,
	0xaa, 0xc2, 0x06, 0x34, 0x44, 0xb2, 0x0f, 0x55, 0xdf, 0x73, 0xa2, 0x24, 0x1c, 0x22, 0x33, 0x8a,
	0x2d, 0xcd, 0x2c, 0x5f, 0x14, 0xec, 0x8a, 0xef, 0x0d, 0x14, 0x42, 0x9e, 0xc3, 0xb6, 0x4b, 0xd9,
	0x9c, 0xc0, 0x8d, 0x52, 0xab, 0x68, 0xd6, 0x6c, 0x5d, 0x62, 0x19, 0x83, 0x37, 0xdf, 0x43, 0xfd,
	0x71, 0x7f, 0xd2, 0x80, 0xe2, 0x03, 0xa6, 0x86, 0x26, 0xab, 0xd9, 0x32, 0x24, 0xff, 0x41, 0x79,
	0x46, 0x83, 0x24, 0x77, 0x60, 0x67, 0xc9, 0xc9, 0xc6, 0xb1, 0xd6, 0xdd, 0x82, 0xb2, 0xf0, 0x45,
	0x80, 0xed, 0xef, 0x45, 0xd0, 0x6f, 0xfc, 0x68, 0x6c, 0xe3, 0xd7, 0x04, 0xb9, 0x58, 0x48, 0xb4,
	0x25, 0x09, 0x79, 0x09, 0x75, 0x86, 0x7c, 0x1a, 0x47, 0x1c, 0x1d, 0x37, 0x4e, 0x22, 0xa1, 0x2a,
	0x96, 0xed, 0xda, 0x1c, 0xed, 0x49, 0x90, 0x58, 0xb0, 0x8b, 0x8c, 0xc5, 0xcc, 0x71, 0x63, 0x0f,
	0x1d, 0x86, 0x22, 0x61, 0x11, 0x7a, 0xea, 0xff, 0x6a, 0xf6, 0x8e, 0x3a, 0xea, 0xc5, 0x1e, 0xda,
	0xf9, 0x01, 0x39, 0x83, 0xed, 0x11, 0xf5, 0x83, 0x84, 0xa1, 0x23, 0xd2, 0x29, 0x1a, 0xa5, 0x96,
	0x66, 0xd6, 0x8f, 0xf6, 0xb3, 0x41, 0x2f, 0xb9, 0xb2, 0x3e, 0x65, 0xac, 0xbb, 0x74, 0x8a, 0xb6,
	0x3e, 0x5a, 0x24, 0xd2, 0x98, 0x3b, 0x41, 0xf7, 0xc1, 0x09, 0x51, 0x50, 0x8f, 0x0a, 0x6a, 0x94,
	0x5b, 0x9a, 0x59, 0xb1, 0x6b, 0x0a, 0xbd, 0xca, 0x41, 0x39, 0x4f, 0x8e, 0x91, 0xe7, 0x4c, 0x90,
	0x7a, 0x72, 0x9e, 0x9b, 0x8a, 0xa4, 0x4b, 0xec, 0x22, 0x83, 0xc8, 0x0b, 0xa8, 0x29, 0x8a, 0x60,
	0xd4, 0x0f, 0x24, 0x67, 0x4b, 0x71, 0x94, 0xee, 0x2e, 0xc7, 0xc8, 0x6b, 0x20, 0x21, 0x72, 0x4e,
	0xc7, 0xe8, 0x04, 0x54, 0x60, 0xe4, 0xa6, 0x4e, 0xc8, 0x8d, 0x8a, 0x9a, 0x45, 0x23, 0x3f, 0xb9,
	0xcc, 0x0e, 0xae, 0x78, 0xfb, 0x15, 0xe8, 0x4b, 0xc6, 0x49, 0x05, 0x4a, 0x83, 0xeb, 0x41, 0xbf,
	0x51, 0x90, 0x51, 0xef, 0xfa, 0xbc, 0xdf, 0xd0, 0x64, 0x74, 0x6e, 0x5f, 0xdf, 0x34, 0x36, 0xda,
	0x1f, 0x60, 0x3b, 0xfb, 0xe3, 0x6c, 0xa0, 0x72, 0x11, 0x9f, 0x97, 0x17, 0xa1, 0x12, 0x62, 0xc0,
	0x96, 0x9a, 0x3f, 0xb2, 0x7c, 0x03, 0xf3, 0xf4, 0xe8, 0x47, 0x11, 0xf4, 0x3b, 0xe4, 0xe2, 0x16,
	0xd9, 0xcc, 0x77, 0x91, 0xbc, 0x83, 0xaa, 0xac, 0xd7, 0x97, 0xb7, 0x83, 0xfc, 0x6f, 0x65, 0xb7,
	0xc6, 0x9a, 0x5f, 0x02, 0x4b, 0xe1, 0x4d, 0xb2, 0x3c, 0xea, 0xac, 0x71, 0xbb, 0x40, 0x3a, 0x50,
	0x92, 0x08, 0xd9, 0xf9, 0x6b, 0x11, 0x2b, 0x04, 0xc7, 0x79, 0x2f, 0xb9, 0xe0, 0x7f, 0xa9, 0x56,
	0xb4, 0x6f, 0x17, 0xc8, 0x5b, 0xa8, 0x48, 0xe2, 0xa5, 0xcf, 0xc5, 0xda, 0xed, 0x0e, 0x35, 0x72,
	0x0a, 0x0d, 0x89, 0xf5, 0x02, 0x1f, 0x23, 0x71, 0x2b, 0x18, 0xd2, 0x70, 0x6d, 0xb9, 0xa9, 0x91,
	0x2e, 0xec, 0xfd, 0x29, 0x5f, 0xe9, 0x7d, 0x55, 0x8d, 0x53, 0xa8, 0x4b, 0xac, 0xeb, 0x7b, 0xfe,
	0x93, 0x0d, 0x1c, 0x4a, 0x0b, 0xbb, 0x8f, 0xe5, 0x4f, 0x35, 0x70, 0xa8, 0x75, 0xd3, 0x5f, 0x67,
	0x27, 0xeb, 0x3c, 0x78, 0xea, 0x21, 0x75, 0xd5, 0xcf, 0x76, 0x16, 0x4f, 0xeb, 0x97, 0xf5, 0xb5,
	0x1c, 0xd9, 0x0c, 0xd9, 0x92, 0x76, 0xb8, 0xa9, 0xe2, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x18, 0xdc, 0x7e, 0xab, 0x05, 0x00, 0x00,
}
