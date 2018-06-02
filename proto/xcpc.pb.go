// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xcpc.proto

package xcpctx

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

type XCPCTransaction struct {
	// Types that are valid to be assigned to Msgtype:
	//	*XCPCTransaction_Send_
	//	*XCPCTransaction_Broadcast_
	Msgtype              isXCPCTransaction_Msgtype `protobuf_oneof:"msgtype"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *XCPCTransaction) Reset()         { *m = XCPCTransaction{} }
func (m *XCPCTransaction) String() string { return proto.CompactTextString(m) }
func (*XCPCTransaction) ProtoMessage()    {}
func (*XCPCTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_xcpc_9c1bed54658c9423, []int{0}
}
func (m *XCPCTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XCPCTransaction.Unmarshal(m, b)
}
func (m *XCPCTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XCPCTransaction.Marshal(b, m, deterministic)
}
func (dst *XCPCTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XCPCTransaction.Merge(dst, src)
}
func (m *XCPCTransaction) XXX_Size() int {
	return xxx_messageInfo_XCPCTransaction.Size(m)
}
func (m *XCPCTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_XCPCTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_XCPCTransaction proto.InternalMessageInfo

type isXCPCTransaction_Msgtype interface {
	isXCPCTransaction_Msgtype()
}

type XCPCTransaction_Send_ struct {
	Send *XCPCTransaction_Send `protobuf:"bytes,1,opt,name=send,oneof"`
}
type XCPCTransaction_Broadcast_ struct {
	Broadcast *XCPCTransaction_Broadcast `protobuf:"bytes,2,opt,name=broadcast,oneof"`
}

func (*XCPCTransaction_Send_) isXCPCTransaction_Msgtype()      {}
func (*XCPCTransaction_Broadcast_) isXCPCTransaction_Msgtype() {}

func (m *XCPCTransaction) GetMsgtype() isXCPCTransaction_Msgtype {
	if m != nil {
		return m.Msgtype
	}
	return nil
}

func (m *XCPCTransaction) GetSend() *XCPCTransaction_Send {
	if x, ok := m.GetMsgtype().(*XCPCTransaction_Send_); ok {
		return x.Send
	}
	return nil
}

func (m *XCPCTransaction) GetBroadcast() *XCPCTransaction_Broadcast {
	if x, ok := m.GetMsgtype().(*XCPCTransaction_Broadcast_); ok {
		return x.Broadcast
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*XCPCTransaction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _XCPCTransaction_OneofMarshaler, _XCPCTransaction_OneofUnmarshaler, _XCPCTransaction_OneofSizer, []interface{}{
		(*XCPCTransaction_Send_)(nil),
		(*XCPCTransaction_Broadcast_)(nil),
	}
}

func _XCPCTransaction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*XCPCTransaction)
	// msgtype
	switch x := m.Msgtype.(type) {
	case *XCPCTransaction_Send_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Send); err != nil {
			return err
		}
	case *XCPCTransaction_Broadcast_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Broadcast); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("XCPCTransaction.Msgtype has unexpected type %T", x)
	}
	return nil
}

func _XCPCTransaction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*XCPCTransaction)
	switch tag {
	case 1: // msgtype.send
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(XCPCTransaction_Send)
		err := b.DecodeMessage(msg)
		m.Msgtype = &XCPCTransaction_Send_{msg}
		return true, err
	case 2: // msgtype.broadcast
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(XCPCTransaction_Broadcast)
		err := b.DecodeMessage(msg)
		m.Msgtype = &XCPCTransaction_Broadcast_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _XCPCTransaction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*XCPCTransaction)
	// msgtype
	switch x := m.Msgtype.(type) {
	case *XCPCTransaction_Send_:
		s := proto.Size(x.Send)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *XCPCTransaction_Broadcast_:
		s := proto.Size(x.Broadcast)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type XCPCTransaction_Send struct {
	Asset                string   `protobuf:"bytes,1,opt,name=asset" json:"asset,omitempty"`
	Quantity             uint64   `protobuf:"varint,2,opt,name=quantity" json:"quantity,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Memo                 string   `protobuf:"bytes,4,opt,name=memo" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XCPCTransaction_Send) Reset()         { *m = XCPCTransaction_Send{} }
func (m *XCPCTransaction_Send) String() string { return proto.CompactTextString(m) }
func (*XCPCTransaction_Send) ProtoMessage()    {}
func (*XCPCTransaction_Send) Descriptor() ([]byte, []int) {
	return fileDescriptor_xcpc_9c1bed54658c9423, []int{0, 0}
}
func (m *XCPCTransaction_Send) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XCPCTransaction_Send.Unmarshal(m, b)
}
func (m *XCPCTransaction_Send) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XCPCTransaction_Send.Marshal(b, m, deterministic)
}
func (dst *XCPCTransaction_Send) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XCPCTransaction_Send.Merge(dst, src)
}
func (m *XCPCTransaction_Send) XXX_Size() int {
	return xxx_messageInfo_XCPCTransaction_Send.Size(m)
}
func (m *XCPCTransaction_Send) XXX_DiscardUnknown() {
	xxx_messageInfo_XCPCTransaction_Send.DiscardUnknown(m)
}

var xxx_messageInfo_XCPCTransaction_Send proto.InternalMessageInfo

func (m *XCPCTransaction_Send) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *XCPCTransaction_Send) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *XCPCTransaction_Send) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *XCPCTransaction_Send) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

type XCPCTransaction_Broadcast struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Feefraction          int32    `protobuf:"varint,3,opt,name=feefraction" json:"feefraction,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XCPCTransaction_Broadcast) Reset()         { *m = XCPCTransaction_Broadcast{} }
func (m *XCPCTransaction_Broadcast) String() string { return proto.CompactTextString(m) }
func (*XCPCTransaction_Broadcast) ProtoMessage()    {}
func (*XCPCTransaction_Broadcast) Descriptor() ([]byte, []int) {
	return fileDescriptor_xcpc_9c1bed54658c9423, []int{0, 1}
}
func (m *XCPCTransaction_Broadcast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XCPCTransaction_Broadcast.Unmarshal(m, b)
}
func (m *XCPCTransaction_Broadcast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XCPCTransaction_Broadcast.Marshal(b, m, deterministic)
}
func (dst *XCPCTransaction_Broadcast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XCPCTransaction_Broadcast.Merge(dst, src)
}
func (m *XCPCTransaction_Broadcast) XXX_Size() int {
	return xxx_messageInfo_XCPCTransaction_Broadcast.Size(m)
}
func (m *XCPCTransaction_Broadcast) XXX_DiscardUnknown() {
	xxx_messageInfo_XCPCTransaction_Broadcast.DiscardUnknown(m)
}

var xxx_messageInfo_XCPCTransaction_Broadcast proto.InternalMessageInfo

func (m *XCPCTransaction_Broadcast) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *XCPCTransaction_Broadcast) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *XCPCTransaction_Broadcast) GetFeefraction() int32 {
	if m != nil {
		return m.Feefraction
	}
	return 0
}

func (m *XCPCTransaction_Broadcast) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*XCPCTransaction)(nil), "xcpctx.XCPCTransaction")
	proto.RegisterType((*XCPCTransaction_Send)(nil), "xcpctx.XCPCTransaction.Send")
	proto.RegisterType((*XCPCTransaction_Broadcast)(nil), "xcpctx.XCPCTransaction.Broadcast")
}

func init() { proto.RegisterFile("xcpc.proto", fileDescriptor_xcpc_9c1bed54658c9423) }

var fileDescriptor_xcpc_9c1bed54658c9423 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xfb, 0x27, 0x6d, 0xc9, 0x75, 0x40, 0x3a, 0x31, 0x44, 0x55, 0x87, 0xc2, 0xc4, 0x94,
	0xa1, 0x7c, 0x02, 0xda, 0xa5, 0x23, 0x32, 0x0c, 0xac, 0x6e, 0x7c, 0x41, 0x91, 0xb0, 0x1d, 0xec,
	0x0b, 0x4a, 0xbf, 0x36, 0x9f, 0x00, 0xe5, 0x42, 0x5a, 0x84, 0xc4, 0x76, 0xef, 0xe9, 0xbd, 0xfb,
	0x9d, 0x65, 0x80, 0xb6, 0xa8, 0x8b, 0xbc, 0x0e, 0x9e, 0x3d, 0xce, 0xbb, 0x99, 0xdb, 0xbb, 0xaf,
	0x09, 0x5c, 0xbf, 0xee, 0x9f, 0xf6, 0x2f, 0x41, 0xbb, 0xa8, 0x0b, 0xae, 0xbc, 0xc3, 0x2d, 0x24,
	0x91, 0x9c, 0xc9, 0xc6, 0x9b, 0xf1, 0xfd, 0x72, 0xbb, 0xce, 0xfb, 0x68, 0xfe, 0x27, 0x96, 0x3f,
	0x93, 0x33, 0x87, 0x91, 0x92, 0x2c, 0x3e, 0x42, 0x7a, 0x0c, 0x5e, 0x9b, 0x42, 0x47, 0xce, 0x26,
	0x52, 0xbc, 0xfd, 0xaf, 0xb8, 0x1b, 0x82, 0x87, 0x91, 0xba, 0xb4, 0x56, 0x25, 0x24, 0xdd, 0x4a,
	0xbc, 0x81, 0x99, 0x8e, 0x91, 0x58, 0xf8, 0xa9, 0xea, 0x05, 0xae, 0xe0, 0xea, 0xa3, 0xd1, 0x8e,
	0x2b, 0x3e, 0xc9, 0xfe, 0x44, 0x9d, 0x35, 0x66, 0xb0, 0xd0, 0xc6, 0x04, 0x8a, 0x31, 0x9b, 0x4a,
	0x67, 0x90, 0x88, 0x90, 0x58, 0xb2, 0x3e, 0x4b, 0xc4, 0x96, 0x79, 0xd5, 0x40, 0x7a, 0xbe, 0xa0,
	0x0b, 0x30, 0xb5, 0x03, 0x4b, 0xe6, 0xee, 0x80, 0x4f, 0xfd, 0xde, 0x90, 0x70, 0xa6, 0xaa, 0x17,
	0xb8, 0x81, 0x65, 0x49, 0x54, 0x86, 0xfe, 0x11, 0x02, 0x9a, 0xa9, 0xdf, 0x16, 0xae, 0x21, 0xe5,
	0xca, 0x52, 0x64, 0x6d, 0xeb, 0x1f, 0xe2, 0xc5, 0xd8, 0xa5, 0xb0, 0xb0, 0xf1, 0x8d, 0x4f, 0x35,
	0x1d, 0xe7, 0xf2, 0x07, 0x0f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x30, 0x09, 0x5c, 0x91,
	0x01, 0x00, 0x00,
}
