// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: op_account.proto

package protoc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OpAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PayPassword string  `protobuf:"bytes,2,opt,name=pay_password,json=payPassword,proto3" json:"pay_password,omitempty"`
	Prices      float64 `protobuf:"fixed64,3,opt,name=prices,proto3" json:"prices,omitempty"`
	PaymentId   int64   `protobuf:"varint,4,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	OrderNum    string  `protobuf:"bytes,5,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	Type        int64   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *OpAccountRequest) Reset() {
	*x = OpAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpAccountRequest) ProtoMessage() {}

func (x *OpAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_op_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpAccountRequest.ProtoReflect.Descriptor instead.
func (*OpAccountRequest) Descriptor() ([]byte, []int) {
	return file_op_account_proto_rawDescGZIP(), []int{0}
}

func (x *OpAccountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OpAccountRequest) GetPayPassword() string {
	if x != nil {
		return x.PayPassword
	}
	return ""
}

func (x *OpAccountRequest) GetPrices() float64 {
	if x != nil {
		return x.Prices
	}
	return 0
}

func (x *OpAccountRequest) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *OpAccountRequest) GetOrderNum() string {
	if x != nil {
		return x.OrderNum
	}
	return ""
}

func (x *OpAccountRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type OpAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States int64 `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
}

func (x *OpAccountResponse) Reset() {
	*x = OpAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpAccountResponse) ProtoMessage() {}

func (x *OpAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_op_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpAccountResponse.ProtoReflect.Descriptor instead.
func (*OpAccountResponse) Descriptor() ([]byte, []int) {
	return file_op_account_proto_rawDescGZIP(), []int{1}
}

func (x *OpAccountResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_op_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_op_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States  int64   `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	UserId  int64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance float64 `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_op_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_op_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_op_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *GetAccountResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetAccountResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_op_account_proto protoreflect.FileDescriptor

var file_op_account_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x10,
	0x4f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x4f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x32, 0x94, 0x01, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x09,
	0x4f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x4f, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4f,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_op_account_proto_rawDescOnce sync.Once
	file_op_account_proto_rawDescData = file_op_account_proto_rawDesc
)

func file_op_account_proto_rawDescGZIP() []byte {
	file_op_account_proto_rawDescOnce.Do(func() {
		file_op_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_op_account_proto_rawDescData)
	})
	return file_op_account_proto_rawDescData
}

var file_op_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_op_account_proto_goTypes = []interface{}{
	(*OpAccountRequest)(nil),   // 0: account.OpAccountRequest
	(*OpAccountResponse)(nil),  // 1: account.OpAccountResponse
	(*GetAccountRequest)(nil),  // 2: account.GetAccountRequest
	(*GetAccountResponse)(nil), // 3: account.GetAccountResponse
}
var file_op_account_proto_depIdxs = []int32{
	0, // 0: account.account.OpAccount:input_type -> account.OpAccountRequest
	2, // 1: account.account.GetAccount:input_type -> account.GetAccountRequest
	1, // 2: account.account.OpAccount:output_type -> account.OpAccountResponse
	3, // 3: account.account.GetAccount:output_type -> account.GetAccountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_op_account_proto_init() }
func file_op_account_proto_init() {
	if File_op_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_op_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_op_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_op_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_op_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_op_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_op_account_proto_goTypes,
		DependencyIndexes: file_op_account_proto_depIdxs,
		MessageInfos:      file_op_account_proto_msgTypes,
	}.Build()
	File_op_account_proto = out.File
	file_op_account_proto_rawDesc = nil
	file_op_account_proto_goTypes = nil
	file_op_account_proto_depIdxs = nil
}
