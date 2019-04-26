// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tele.proto

package tele

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PaymentMethod int32

const (
	PaymentMethod_Nothing  PaymentMethod = 0
	PaymentMethod_Cash     PaymentMethod = 1
	PaymentMethod_Cashless PaymentMethod = 2
)

var PaymentMethod_name = map[int32]string{
	0: "Nothing",
	1: "Cash",
	2: "Cashless",
}
var PaymentMethod_value = map[string]int32{
	"Nothing":  0,
	"Cash":     1,
	"Cashless": 2,
}

func (x PaymentMethod) String() string {
	return proto.EnumName(PaymentMethod_name, int32(x))
}
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0}
}

type Command_Task int32

const (
	Command_Invalid       Command_Task = 0
	Command_Report        Command_Task = 1
	Command_Abort         Command_Task = 2
	Command_MoneyDispense Command_Task = 3
)

var Command_Task_name = map[int32]string{
	0: "Invalid",
	1: "Report",
	2: "Abort",
	3: "MoneyDispense",
}
var Command_Task_value = map[string]int32{
	"Invalid":       0,
	"Report":        1,
	"Abort":         2,
	"MoneyDispense": 3,
}

func (x Command_Task) String() string {
	return proto.EnumName(Command_Task_name, int32(x))
}
func (Command_Task) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{1, 0}
}

// Optimising for rare, bulk delivery on cell network.
// "Touching network" is expensive, while 10 or 900 bytes is about same cost.
type Telemetry struct {
	VmId                 int32                  `protobuf:"varint,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	Error                *Telemetry_Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Inventory            *Telemetry_Inventory   `protobuf:"bytes,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Money                *Telemetry_Money       `protobuf:"bytes,4,opt,name=money,proto3" json:"money,omitempty"`
	Transaction          *Telemetry_Transaction `protobuf:"bytes,5,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0}
}
func (m *Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry.Unmarshal(m, b)
}
func (m *Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry.Marshal(b, m, deterministic)
}
func (dst *Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry.Merge(dst, src)
}
func (m *Telemetry) XXX_Size() int {
	return xxx_messageInfo_Telemetry.Size(m)
}
func (m *Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry proto.InternalMessageInfo

func (m *Telemetry) GetVmId() int32 {
	if m != nil {
		return m.VmId
	}
	return 0
}

func (m *Telemetry) GetError() *Telemetry_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Telemetry) GetInventory() *Telemetry_Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

func (m *Telemetry) GetMoney() *Telemetry_Money {
	if m != nil {
		return m.Money
	}
	return nil
}

func (m *Telemetry) GetTransaction() *Telemetry_Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type Telemetry_Error struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Error) Reset()         { *m = Telemetry_Error{} }
func (m *Telemetry_Error) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Error) ProtoMessage()    {}
func (*Telemetry_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0, 0}
}
func (m *Telemetry_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Error.Unmarshal(m, b)
}
func (m *Telemetry_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Error.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Error.Merge(dst, src)
}
func (m *Telemetry_Error) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Error.Size(m)
}
func (m *Telemetry_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Error proto.InternalMessageInfo

func (m *Telemetry_Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Telemetry_Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Telemetry_Error) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Telemetry_Inventory struct {
	Water                int32    `protobuf:"varint,1,opt,name=water,proto3" json:"water,omitempty"`
	Coffee               int32    `protobuf:"varint,2,opt,name=coffee,proto3" json:"coffee,omitempty"`
	Cup                  int32    `protobuf:"varint,3,opt,name=cup,proto3" json:"cup,omitempty"`
	Hoppers              []int32  `protobuf:"varint,4,rep,packed,name=hoppers,proto3" json:"hoppers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Inventory) Reset()         { *m = Telemetry_Inventory{} }
func (m *Telemetry_Inventory) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Inventory) ProtoMessage()    {}
func (*Telemetry_Inventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0, 1}
}
func (m *Telemetry_Inventory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Inventory.Unmarshal(m, b)
}
func (m *Telemetry_Inventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Inventory.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Inventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Inventory.Merge(dst, src)
}
func (m *Telemetry_Inventory) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Inventory.Size(m)
}
func (m *Telemetry_Inventory) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Inventory.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Inventory proto.InternalMessageInfo

func (m *Telemetry_Inventory) GetWater() int32 {
	if m != nil {
		return m.Water
	}
	return 0
}

func (m *Telemetry_Inventory) GetCoffee() int32 {
	if m != nil {
		return m.Coffee
	}
	return 0
}

func (m *Telemetry_Inventory) GetCup() int32 {
	if m != nil {
		return m.Cup
	}
	return 0
}

func (m *Telemetry_Inventory) GetHoppers() []int32 {
	if m != nil {
		return m.Hoppers
	}
	return nil
}

type Telemetry_Money struct {
	TotalBills           uint32   `protobuf:"varint,1,opt,name=total_bills,json=totalBills,proto3" json:"total_bills,omitempty"`
	TotalCoins           uint32   `protobuf:"varint,2,opt,name=total_coins,json=totalCoins,proto3" json:"total_coins,omitempty"`
	Bills                []uint32 `protobuf:"varint,3,rep,packed,name=bills,proto3" json:"bills,omitempty"`
	Coins                []uint32 `protobuf:"varint,4,rep,packed,name=coins,proto3" json:"coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Money) Reset()         { *m = Telemetry_Money{} }
func (m *Telemetry_Money) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Money) ProtoMessage()    {}
func (*Telemetry_Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0, 2}
}
func (m *Telemetry_Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Money.Unmarshal(m, b)
}
func (m *Telemetry_Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Money.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Money.Merge(dst, src)
}
func (m *Telemetry_Money) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Money.Size(m)
}
func (m *Telemetry_Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Money proto.InternalMessageInfo

func (m *Telemetry_Money) GetTotalBills() uint32 {
	if m != nil {
		return m.TotalBills
	}
	return 0
}

func (m *Telemetry_Money) GetTotalCoins() uint32 {
	if m != nil {
		return m.TotalCoins
	}
	return 0
}

func (m *Telemetry_Money) GetBills() []uint32 {
	if m != nil {
		return m.Bills
	}
	return nil
}

func (m *Telemetry_Money) GetCoins() []uint32 {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Telemetry_Transaction struct {
	Code                 int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Options              []int32       `protobuf:"varint,2,rep,packed,name=options,proto3" json:"options,omitempty"`
	Price                uint32        `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	PaymentMethod        PaymentMethod `protobuf:"varint,4,opt,name=payment_method,json=paymentMethod,proto3,enum=tele.PaymentMethod" json:"payment_method,omitempty"`
	CreditBills          uint32        `protobuf:"varint,5,opt,name=credit_bills,json=creditBills,proto3" json:"credit_bills,omitempty"`
	CreditCoins          uint32        `protobuf:"varint,6,opt,name=credit_coins,json=creditCoins,proto3" json:"credit_coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Telemetry_Transaction) Reset()         { *m = Telemetry_Transaction{} }
func (m *Telemetry_Transaction) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Transaction) ProtoMessage()    {}
func (*Telemetry_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{0, 3}
}
func (m *Telemetry_Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Transaction.Unmarshal(m, b)
}
func (m *Telemetry_Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Transaction.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Transaction.Merge(dst, src)
}
func (m *Telemetry_Transaction) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Transaction.Size(m)
}
func (m *Telemetry_Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Transaction proto.InternalMessageInfo

func (m *Telemetry_Transaction) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Telemetry_Transaction) GetOptions() []int32 {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Telemetry_Transaction) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Telemetry_Transaction) GetPaymentMethod() PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return PaymentMethod_Nothing
}

func (m *Telemetry_Transaction) GetCreditBills() uint32 {
	if m != nil {
		return m.CreditBills
	}
	return 0
}

func (m *Telemetry_Transaction) GetCreditCoins() uint32 {
	if m != nil {
		return m.CreditCoins
	}
	return 0
}

type Command struct {
	Id                   uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Task                 Command_Task `protobuf:"varint,2,opt,name=task,proto3,enum=tele.Command_Task" json:"task,omitempty"`
	Args                 []string     `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	ReplyTopic           string       `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic,proto3" json:"reply_topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{1}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Command) GetTask() Command_Task {
	if m != nil {
		return m.Task
	}
	return Command_Invalid
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetReplyTopic() string {
	if m != nil {
		return m.ReplyTopic
	}
	return ""
}

type Response struct {
	CommandId            uint32   `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_c093fa7185778f84, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "tele.Telemetry")
	proto.RegisterType((*Telemetry_Error)(nil), "tele.Telemetry.Error")
	proto.RegisterType((*Telemetry_Inventory)(nil), "tele.Telemetry.Inventory")
	proto.RegisterType((*Telemetry_Money)(nil), "tele.Telemetry.Money")
	proto.RegisterType((*Telemetry_Transaction)(nil), "tele.Telemetry.Transaction")
	proto.RegisterType((*Command)(nil), "tele.Command")
	proto.RegisterType((*Response)(nil), "tele.Response")
	proto.RegisterEnum("tele.PaymentMethod", PaymentMethod_name, PaymentMethod_value)
	proto.RegisterEnum("tele.Command_Task", Command_Task_name, Command_Task_value)
}

func init() { proto.RegisterFile("tele.proto", fileDescriptor_tele_c093fa7185778f84) }

var fileDescriptor_tele_c093fa7185778f84 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x13, 0x6f, 0x5a, 0x4f, 0x9a, 0xc8, 0x6c, 0x01, 0x99, 0x20, 0x44, 0xc8, 0x01, 0x55,
	0x20, 0xe5, 0x50, 0x90, 0x90, 0x90, 0x7a, 0x80, 0xc2, 0x21, 0x42, 0x45, 0x68, 0xc9, 0x3d, 0xda,
	0xda, 0xd3, 0xc6, 0xaa, 0xed, 0x5d, 0xed, 0x6e, 0x83, 0xf2, 0x60, 0x3c, 0x09, 0x8f, 0xc2, 0x0b,
	0xa0, 0x9d, 0x75, 0x7e, 0x5a, 0x71, 0x9b, 0x6f, 0xe6, 0x9b, 0x9f, 0xef, 0xdb, 0x38, 0x00, 0x0e,
	0x2b, 0x9c, 0x6a, 0xa3, 0x9c, 0xe2, 0xb1, 0x8f, 0x27, 0x7f, 0x19, 0x24, 0x73, 0xac, 0xb0, 0x46,
	0x67, 0xd6, 0xfc, 0x04, 0xd8, 0xaa, 0x5e, 0x94, 0x45, 0x16, 0x8d, 0xa3, 0x53, 0x26, 0xe2, 0x55,
	0x3d, 0x2b, 0xf8, 0x5b, 0x60, 0x68, 0x8c, 0x32, 0x59, 0x67, 0x1c, 0x9d, 0xf6, 0xcf, 0x9e, 0x4c,
	0x69, 0xc8, 0xb6, 0x69, 0xfa, 0xd5, 0x17, 0x45, 0xe0, 0xf0, 0x0f, 0x90, 0x94, 0xcd, 0x0a, 0x1b,
	0xa7, 0xcc, 0x3a, 0xeb, 0x52, 0xc3, 0xb3, 0x87, 0x0d, 0xb3, 0x0d, 0x41, 0xec, 0xb8, 0x7e, 0x4b,
	0xad, 0x1a, 0x5c, 0x67, 0xf1, 0xff, 0xb7, 0x5c, 0xfa, 0xa2, 0x08, 0x1c, 0x7e, 0x0e, 0x7d, 0x67,
	0x64, 0x63, 0x65, 0xee, 0x4a, 0xd5, 0x64, 0x8c, 0x5a, 0x9e, 0x3f, 0x6c, 0x99, 0xef, 0x28, 0x62,
	0x9f, 0x3f, 0xfa, 0x06, 0x8c, 0x8e, 0xe6, 0x1c, 0xe2, 0x5c, 0x15, 0x48, 0x72, 0x07, 0x82, 0x62,
	0x9e, 0xc1, 0x61, 0x8d, 0xd6, 0xca, 0x1b, 0x24, 0xc1, 0x89, 0xd8, 0x40, 0xfe, 0x18, 0x58, 0xae,
	0xee, 0x1a, 0x47, 0xba, 0x06, 0x22, 0x80, 0x11, 0x42, 0xb2, 0x15, 0xe4, 0x29, 0xbf, 0xa4, 0x43,
	0xd3, 0x1a, 0x18, 0x00, 0x7f, 0x0a, 0xbd, 0x5c, 0x5d, 0x5f, 0x63, 0x98, 0xc8, 0x44, 0x8b, 0x78,
	0x0a, 0xdd, 0xfc, 0x4e, 0xd3, 0x38, 0x26, 0x7c, 0xe8, 0x97, 0x2f, 0x95, 0xd6, 0x68, 0x6c, 0x16,
	0x8f, 0xbb, 0xa7, 0x4c, 0x6c, 0xe0, 0xe8, 0x0e, 0x18, 0x59, 0xc0, 0x5f, 0x42, 0xdf, 0x29, 0x27,
	0xab, 0xc5, 0x55, 0x59, 0x55, 0xb6, 0x3d, 0x1d, 0x28, 0xf5, 0xd9, 0x67, 0x76, 0x84, 0x5c, 0x95,
	0x8d, 0xa5, 0x95, 0x1b, 0xc2, 0x85, 0xcf, 0xf8, 0x23, 0x43, 0x6f, 0x77, 0xdc, 0xf5, 0x3a, 0x08,
	0x04, 0x75, 0xbe, 0x21, 0x0e, 0x59, 0x02, 0xa3, 0x3f, 0x11, 0xf4, 0xf7, 0x7c, 0xbc, 0xe7, 0x18,
	0xdb, 0x39, 0xa6, 0xb4, 0xaf, 0xfa, 0x65, 0x74, 0x74, 0x0b, 0xfd, 0x4c, 0x6d, 0xca, 0x1c, 0x37,
	0x8e, 0x11, 0xe0, 0x1f, 0x61, 0xa8, 0xe5, 0xba, 0xc6, 0xc6, 0x2d, 0x6a, 0x74, 0x4b, 0x55, 0xd0,
	0x9b, 0x0f, 0xcf, 0x4e, 0xc2, 0x03, 0xfe, 0x08, 0xb5, 0x4b, 0x2a, 0x89, 0x81, 0xde, 0x87, 0xfc,
	0x15, 0x1c, 0xe7, 0x06, 0x8b, 0xd2, 0xb5, 0xf2, 0x19, 0x0d, 0xee, 0x87, 0x5c, 0xd0, 0xbf, 0xa3,
	0x04, 0x3d, 0xbd, 0x7d, 0x0a, 0x39, 0x30, 0xf9, 0x1d, 0xc1, 0xe1, 0x85, 0xaa, 0x6b, 0xd9, 0x14,
	0x7c, 0x08, 0x9d, 0xf6, 0x07, 0x3f, 0x10, 0x9d, 0xb2, 0xe0, 0xaf, 0x21, 0x76, 0xd2, 0xde, 0x92,
	0x6f, 0xc3, 0x33, 0x1e, 0x6e, 0x6a, 0xc9, 0xd3, 0xb9, 0xb4, 0xb7, 0x82, 0xea, 0xde, 0x09, 0x69,
	0x6e, 0x82, 0x89, 0x89, 0xa0, 0xd8, 0x5b, 0x6f, 0x50, 0x57, 0xeb, 0x85, 0x53, 0xba, 0xcc, 0x49,
	0x56, 0x22, 0x80, 0x52, 0x73, 0x9f, 0x99, 0x9c, 0x43, 0xec, 0x47, 0xf0, 0x3e, 0x1c, 0xce, 0x9a,
	0x95, 0xac, 0xca, 0x22, 0x3d, 0xe0, 0x00, 0x3d, 0x81, 0x5a, 0x19, 0x97, 0x46, 0x3c, 0x01, 0xf6,
	0xe9, 0xca, 0x87, 0x1d, 0xfe, 0x08, 0x06, 0xf4, 0xe2, 0x5f, 0x4a, 0xab, 0xb1, 0xb1, 0x98, 0x76,
	0x27, 0x3f, 0xe1, 0x48, 0xa0, 0xd5, 0xaa, 0xb1, 0xc8, 0x5f, 0x00, 0xe4, 0xe1, 0xaa, 0xc5, 0xf6,
	0xfe, 0xa4, 0xcd, 0xcc, 0x0a, 0x6f, 0xfd, 0xee, 0xab, 0x4d, 0x36, 0x9f, 0x27, 0x87, 0xb8, 0x90,
	0x4e, 0xd2, 0x7b, 0x24, 0x82, 0xe2, 0x37, 0xef, 0x61, 0x70, 0xcf, 0x72, 0x7f, 0xdc, 0x77, 0xe5,
	0x96, 0x65, 0x73, 0x93, 0x1e, 0xf0, 0x23, 0x88, 0x2f, 0xa4, 0x5d, 0xa6, 0x11, 0x3f, 0x86, 0x23,
	0x1f, 0x55, 0x68, 0x6d, 0xda, 0xb9, 0xea, 0xd1, 0xbf, 0xc8, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x48, 0xa9, 0x2f, 0x90, 0x53, 0x04, 0x00, 0x00,
}
