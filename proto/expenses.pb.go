// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/expenses.proto

package expensespb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type Payment struct {
	Date                 int64    `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Amount               float32  `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2084bc508b6b2bcc, []int{0}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Payment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Payment) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Payment) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Payment) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Payment) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Expenses struct {
	Payments             []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Expenses) Reset()         { *m = Expenses{} }
func (m *Expenses) String() string { return proto.CompactTextString(m) }
func (*Expenses) ProtoMessage()    {}
func (*Expenses) Descriptor() ([]byte, []int) {
	return fileDescriptor_2084bc508b6b2bcc, []int{1}
}

func (m *Expenses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expenses.Unmarshal(m, b)
}
func (m *Expenses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expenses.Marshal(b, m, deterministic)
}
func (m *Expenses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expenses.Merge(m, src)
}
func (m *Expenses) XXX_Size() int {
	return xxx_messageInfo_Expenses.Size(m)
}
func (m *Expenses) XXX_DiscardUnknown() {
	xxx_messageInfo_Expenses.DiscardUnknown(m)
}

var xxx_messageInfo_Expenses proto.InternalMessageInfo

func (m *Expenses) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

type ExpensesRequest struct {
	FromDate             int64    `protobuf:"varint,1,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate               int64    `protobuf:"varint,2,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpensesRequest) Reset()         { *m = ExpensesRequest{} }
func (m *ExpensesRequest) String() string { return proto.CompactTextString(m) }
func (*ExpensesRequest) ProtoMessage()    {}
func (*ExpensesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2084bc508b6b2bcc, []int{2}
}

func (m *ExpensesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpensesRequest.Unmarshal(m, b)
}
func (m *ExpensesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpensesRequest.Marshal(b, m, deterministic)
}
func (m *ExpensesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpensesRequest.Merge(m, src)
}
func (m *ExpensesRequest) XXX_Size() int {
	return xxx_messageInfo_ExpensesRequest.Size(m)
}
func (m *ExpensesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpensesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExpensesRequest proto.InternalMessageInfo

func (m *ExpensesRequest) GetFromDate() int64 {
	if m != nil {
		return m.FromDate
	}
	return 0
}

func (m *ExpensesRequest) GetToDate() int64 {
	if m != nil {
		return m.ToDate
	}
	return 0
}

func (m *ExpensesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ExpensesRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*Payment)(nil), "expenses.Payment")
	proto.RegisterType((*Expenses)(nil), "expenses.Expenses")
	proto.RegisterType((*ExpensesRequest)(nil), "expenses.ExpensesRequest")
}

func init() { proto.RegisterFile("proto/expenses.proto", fileDescriptor_2084bc508b6b2bcc) }

var fileDescriptor_2084bc508b6b2bcc = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbb, 0x4f, 0xc3, 0x30,
	0x10, 0xc6, 0x71, 0xd2, 0xa4, 0xe9, 0x15, 0xf1, 0xb0, 0x2a, 0x30, 0xb0, 0x44, 0x99, 0xb2, 0x50,
	0xa4, 0x32, 0xb1, 0x30, 0x20, 0x10, 0x2b, 0x98, 0x8d, 0x05, 0xb9, 0x70, 0x44, 0x91, 0x48, 0x6c,
	0xe2, 0x2b, 0xa2, 0x23, 0xff, 0x39, 0xb2, 0xf3, 0x28, 0xa2, 0xdb, 0xf7, 0xc8, 0x29, 0xbf, 0x4f,
	0x86, 0x99, 0x69, 0x34, 0xe9, 0x0b, 0xfc, 0x36, 0x58, 0x5b, 0xb4, 0x73, 0x6f, 0x79, 0xd2, 0xfb,
	0xec, 0x87, 0xc1, 0xf8, 0x41, 0xad, 0x2b, 0xac, 0x89, 0x73, 0x18, 0xbd, 0x29, 0x42, 0xc1, 0x52,
	0x96, 0x87, 0xd2, 0x6b, 0x97, 0xd1, 0xda, 0xa0, 0x08, 0x52, 0x96, 0x4f, 0xa4, 0xd7, 0x2e, 0x7b,
	0x6f, 0x74, 0x25, 0xc2, 0x36, 0x73, 0x9a, 0xef, 0x41, 0x40, 0x5a, 0x8c, 0x7c, 0x12, 0x90, 0xf6,
	0x77, 0xaa, 0xb0, 0x22, 0x4a, 0x43, 0x7f, 0xa7, 0x0a, 0xcb, 0x8f, 0x20, 0x56, 0x95, 0x5e, 0xd5,
	0x24, 0xe2, 0x94, 0xe5, 0x81, 0xec, 0x5c, 0x76, 0x05, 0xc9, 0x5d, 0xc7, 0xc3, 0xcf, 0x21, 0x31,
	0x2d, 0x8e, 0x15, 0x2c, 0x0d, 0xf3, 0xe9, 0xe2, 0x70, 0x3e, 0xc0, 0x77, 0xa0, 0x72, 0xf8, 0x24,
	0xd3, 0xb0, 0xdf, 0x9f, 0x4a, 0xfc, 0x5c, 0xa1, 0x25, 0x7e, 0x06, 0x13, 0x47, 0xf4, 0xf2, 0x67,
	0x4a, 0xe2, 0x82, 0x5b, 0x37, 0xe7, 0x18, 0xc6, 0xa4, 0xdb, 0x2a, 0xf0, 0x55, 0x4c, 0xda, 0x17,
	0x33, 0x88, 0x3e, 0xca, 0xaa, 0x24, 0x3f, 0x2a, 0x92, 0xad, 0xe1, 0x07, 0x10, 0x92, 0x2a, 0xba,
	0x59, 0x4e, 0x2e, 0x1e, 0x37, 0x3f, 0x7c, 0xc2, 0xe6, 0xab, 0x7c, 0x45, 0x7e, 0x0d, 0xd3, 0x7b,
	0xa4, 0x61, 0xc1, 0xc9, 0x86, 0xf7, 0x1f, 0xda, 0x29, 0xdf, 0xae, 0xb2, 0x9d, 0x9b, 0xdd, 0x67,
	0xe8, 0x63, 0xb3, 0x5c, 0xc6, 0xfe, 0x85, 0x2e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0xdd,
	0x80, 0x75, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExpensesServiceClient is the client API for ExpensesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExpensesServiceClient interface {
	GetExpenses(ctx context.Context, in *ExpensesRequest, opts ...grpc.CallOption) (*Expenses, error)
}

type expensesServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpensesServiceClient(cc *grpc.ClientConn) ExpensesServiceClient {
	return &expensesServiceClient{cc}
}

func (c *expensesServiceClient) GetExpenses(ctx context.Context, in *ExpensesRequest, opts ...grpc.CallOption) (*Expenses, error) {
	out := new(Expenses)
	err := c.cc.Invoke(ctx, "/expenses.ExpensesService/GetExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpensesServiceServer is the server API for ExpensesService service.
type ExpensesServiceServer interface {
	GetExpenses(context.Context, *ExpensesRequest) (*Expenses, error)
}

func RegisterExpensesServiceServer(s *grpc.Server, srv ExpensesServiceServer) {
	s.RegisterService(&_ExpensesService_serviceDesc, srv)
}

func _ExpensesService_GetExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpensesServiceServer).GetExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/expenses.ExpensesService/GetExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpensesServiceServer).GetExpenses(ctx, req.(*ExpensesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpensesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "expenses.ExpensesService",
	HandlerType: (*ExpensesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpenses",
			Handler:    _ExpensesService_GetExpenses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/expenses.proto",
}
