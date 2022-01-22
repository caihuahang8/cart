// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: cart.proto

package go_micro_service_cart

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

type CartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *CartInfo) Reset() {
	*x = CartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartInfo) ProtoMessage() {}

func (x *CartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartInfo.ProtoReflect.Descriptor instead.
func (*CartInfo) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartInfo) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartInfo) GetSizeId() int64 {
	if x != nil {
		return x.SizeId
	}
	return 0
}

func (x *CartInfo) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ResponseAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId int64  `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseAdd) Reset() {
	*x = ResponseAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAdd) ProtoMessage() {}

func (x *ResponseAdd) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAdd.ProtoReflect.Descriptor instead.
func (*ResponseAdd) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseAdd) GetCartId() int64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *ResponseAdd) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Clean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Clean) Reset() {
	*x = Clean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clean) ProtoMessage() {}

func (x *Clean) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clean.ProtoReflect.Descriptor instead.
func (*Clean) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *Clean) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meg string `protobuf:"bytes,1,opt,name=meg,proto3" json:"meg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMeg() string {
	if x != nil {
		return x.Meg
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChangeNum int64 `protobuf:"varint,2,opt,name=change_num,json=changeNum,proto3" json:"change_num,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetChangeNum() int64 {
	if x != nil {
		return x.ChangeNum
	}
	return 0
}

type CartID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CartID) Reset() {
	*x = CartID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartID) ProtoMessage() {}

func (x *CartID) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartID.ProtoReflect.Descriptor instead.
func (*CartID) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *CartID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CartFindAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CartFindAll) Reset() {
	*x = CartFindAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartFindAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartFindAll) ProtoMessage() {}

func (x *CartFindAll) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartFindAll.ProtoReflect.Descriptor instead.
func (*CartFindAll) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CartFindAll) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CartAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartInfo []*CartInfo `protobuf:"bytes,1,rep,name=cart_info,json=cartInfo,proto3" json:"cart_info,omitempty"`
}

func (x *CartAll) Reset() {
	*x = CartAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartAll) ProtoMessage() {}

func (x *CartAll) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartAll.ProtoReflect.Descriptor instead.
func (*CartAll) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

func (x *CartAll) GetCartInfo() []*CartInfo {
	if x != nil {
		return x.CartInfo
	}
	return nil
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x22, 0x7d, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x38, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x05,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x65, 0x67, 0x22, 0x35, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x22, 0x18, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x3c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xda,
	0x03, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x64, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x6e, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x04, 0x49, 0x6e, 0x63, 0x72, 0x12,
	0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x1f, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cart_proto_goTypes = []interface{}{
	(*CartInfo)(nil),    // 0: go.micro.service.cart.CartInfo
	(*ResponseAdd)(nil), // 1: go.micro.service.cart.ResponseAdd
	(*Clean)(nil),       // 2: go.micro.service.cart.Clean
	(*Response)(nil),    // 3: go.micro.service.cart.Response
	(*Item)(nil),        // 4: go.micro.service.cart.Item
	(*CartID)(nil),      // 5: go.micro.service.cart.CartID
	(*CartFindAll)(nil), // 6: go.micro.service.cart.CartFindAll
	(*CartAll)(nil),     // 7: go.micro.service.cart.CartAll
}
var file_cart_proto_depIdxs = []int32{
	0, // 0: go.micro.service.cart.CartAll.cart_info:type_name -> go.micro.service.cart.CartInfo
	0, // 1: go.micro.service.cart.Cart.AddCart:input_type -> go.micro.service.cart.CartInfo
	2, // 2: go.micro.service.cart.Cart.CleanCart:input_type -> go.micro.service.cart.Clean
	4, // 3: go.micro.service.cart.Cart.Incr:input_type -> go.micro.service.cart.Item
	4, // 4: go.micro.service.cart.Cart.Decr:input_type -> go.micro.service.cart.Item
	5, // 5: go.micro.service.cart.Cart.DeleteItemByID:input_type -> go.micro.service.cart.CartID
	6, // 6: go.micro.service.cart.Cart.GetAll:input_type -> go.micro.service.cart.CartFindAll
	1, // 7: go.micro.service.cart.Cart.AddCart:output_type -> go.micro.service.cart.ResponseAdd
	3, // 8: go.micro.service.cart.Cart.CleanCart:output_type -> go.micro.service.cart.Response
	3, // 9: go.micro.service.cart.Cart.Incr:output_type -> go.micro.service.cart.Response
	3, // 10: go.micro.service.cart.Cart.Decr:output_type -> go.micro.service.cart.Response
	3, // 11: go.micro.service.cart.Cart.DeleteItemByID:output_type -> go.micro.service.cart.Response
	7, // 12: go.micro.service.cart.Cart.GetAll:output_type -> go.micro.service.cart.CartAll
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartInfo); i {
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
		file_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAdd); i {
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
		file_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clean); i {
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
		file_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartID); i {
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
		file_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartFindAll); i {
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
		file_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartAll); i {
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
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
