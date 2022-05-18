// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: grpc/deck_service.proto

package grpc

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

// CreateDeckRequest holds the composition and optional codes needed to create a deck.
type CreateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comp  *string  `protobuf:"bytes,1,opt,name=comp,proto3,oneof" json:"comp,omitempty"`
	Codes []string `protobuf:"bytes,2,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *CreateDeckRequest) Reset() {
	*x = CreateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckRequest) ProtoMessage() {}

func (x *CreateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckRequest.ProtoReflect.Descriptor instead.
func (*CreateDeckRequest) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDeckRequest) GetComp() string {
	if x != nil && x.Comp != nil {
		return *x.Comp
	}
	return ""
}

func (x *CreateDeckRequest) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

// CreateDeckResponse holds the deck created.
type CreateDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *CreateDeckResponse) Reset() {
	*x = CreateDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckResponse) ProtoMessage() {}

func (x *CreateDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckResponse.ProtoReflect.Descriptor instead.
func (*CreateDeckResponse) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

// OpenDeckRequest holds the deck id needed to open a deck.
type OpenDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpenDeckRequest) Reset() {
	*x = OpenDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDeckRequest) ProtoMessage() {}

func (x *OpenDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDeckRequest.ProtoReflect.Descriptor instead.
func (*OpenDeckRequest) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{2}
}

func (x *OpenDeckRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// OpenDeckResponse holds the deck opened.
type OpenDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *OpenDeckResponse) Reset() {
	*x = OpenDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDeckResponse) ProtoMessage() {}

func (x *OpenDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDeckResponse.ProtoReflect.Descriptor instead.
func (*OpenDeckResponse) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{3}
}

func (x *OpenDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

// DrawCardsRequest holds the id of the deck and number of cards to draw from the deck.
type DrawCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int32  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *DrawCardsRequest) Reset() {
	*x = DrawCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawCardsRequest) ProtoMessage() {}

func (x *DrawCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawCardsRequest.ProtoReflect.Descriptor instead.
func (*DrawCardsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{4}
}

func (x *DrawCardsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DrawCardsRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

// DrawCardsResponse holds the cards drawn.
type DrawCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *DrawCardsResponse) Reset() {
	*x = DrawCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawCardsResponse) ProtoMessage() {}

func (x *DrawCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawCardsResponse.ProtoReflect.Descriptor instead.
func (*DrawCardsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{5}
}

func (x *DrawCardsResponse) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

// ShuffleDeckRequest holds the id of the deck to shuffle.
type ShuffleDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShuffleDeckRequest) Reset() {
	*x = ShuffleDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShuffleDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShuffleDeckRequest) ProtoMessage() {}

func (x *ShuffleDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShuffleDeckRequest.ProtoReflect.Descriptor instead.
func (*ShuffleDeckRequest) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{6}
}

func (x *ShuffleDeckRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ShuffleDeckResponse holds the message after shuffling a deck.
type ShuffleDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ShuffleDeckResponse) Reset() {
	*x = ShuffleDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_deck_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShuffleDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShuffleDeckResponse) ProtoMessage() {}

func (x *ShuffleDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_deck_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShuffleDeckResponse.ProtoReflect.Descriptor instead.
func (*ShuffleDeckResponse) Descriptor() ([]byte, []int) {
	return file_grpc_deck_service_proto_rawDescGZIP(), []int{7}
}

func (x *ShuffleDeckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_deck_service_proto protoreflect.FileDescriptor

var file_grpc_deck_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x22, 0x34, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64,
	0x65, 0x63, 0x6b, 0x22, 0x21, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x65,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x22, 0x34, 0x0a, 0x10, 0x44, 0x72,
	0x61, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x35, 0x0a, 0x11, 0x44, 0x72, 0x61, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x68, 0x75, 0x66, 0x66,
	0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a,
	0x13, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x93,
	0x02, 0x0a, 0x0b, 0x44, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x09, 0x44, 0x72, 0x61, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0b, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x68, 0x75, 0x66, 0x66, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x67, 0x6a, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_deck_service_proto_rawDescOnce sync.Once
	file_grpc_deck_service_proto_rawDescData = file_grpc_deck_service_proto_rawDesc
)

func file_grpc_deck_service_proto_rawDescGZIP() []byte {
	file_grpc_deck_service_proto_rawDescOnce.Do(func() {
		file_grpc_deck_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_deck_service_proto_rawDescData)
	})
	return file_grpc_deck_service_proto_rawDescData
}

var file_grpc_deck_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_deck_service_proto_goTypes = []interface{}{
	(*CreateDeckRequest)(nil),   // 0: grpc.CreateDeckRequest
	(*CreateDeckResponse)(nil),  // 1: grpc.CreateDeckResponse
	(*OpenDeckRequest)(nil),     // 2: grpc.OpenDeckRequest
	(*OpenDeckResponse)(nil),    // 3: grpc.OpenDeckResponse
	(*DrawCardsRequest)(nil),    // 4: grpc.DrawCardsRequest
	(*DrawCardsResponse)(nil),   // 5: grpc.DrawCardsResponse
	(*ShuffleDeckRequest)(nil),  // 6: grpc.ShuffleDeckRequest
	(*ShuffleDeckResponse)(nil), // 7: grpc.ShuffleDeckResponse
	(*Deck)(nil),                // 8: grpc.Deck
	(*Card)(nil),                // 9: grpc.Card
}
var file_grpc_deck_service_proto_depIdxs = []int32{
	8, // 0: grpc.CreateDeckResponse.deck:type_name -> grpc.Deck
	8, // 1: grpc.OpenDeckResponse.deck:type_name -> grpc.Deck
	9, // 2: grpc.DrawCardsResponse.cards:type_name -> grpc.Card
	0, // 3: grpc.DeckService.CreateDeck:input_type -> grpc.CreateDeckRequest
	2, // 4: grpc.DeckService.OpenDeck:input_type -> grpc.OpenDeckRequest
	4, // 5: grpc.DeckService.DrawCards:input_type -> grpc.DrawCardsRequest
	6, // 6: grpc.DeckService.ShuffleDeck:input_type -> grpc.ShuffleDeckRequest
	1, // 7: grpc.DeckService.CreateDeck:output_type -> grpc.CreateDeckResponse
	3, // 8: grpc.DeckService.OpenDeck:output_type -> grpc.OpenDeckResponse
	5, // 9: grpc.DeckService.DrawCards:output_type -> grpc.DrawCardsResponse
	7, // 10: grpc.DeckService.ShuffleDeck:output_type -> grpc.ShuffleDeckResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_deck_service_proto_init() }
func file_grpc_deck_service_proto_init() {
	if File_grpc_deck_service_proto != nil {
		return
	}
	file_grpc_deck_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_deck_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckRequest); i {
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
		file_grpc_deck_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckResponse); i {
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
		file_grpc_deck_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDeckRequest); i {
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
		file_grpc_deck_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDeckResponse); i {
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
		file_grpc_deck_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawCardsRequest); i {
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
		file_grpc_deck_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawCardsResponse); i {
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
		file_grpc_deck_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShuffleDeckRequest); i {
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
		file_grpc_deck_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShuffleDeckResponse); i {
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
	file_grpc_deck_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_deck_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_deck_service_proto_goTypes,
		DependencyIndexes: file_grpc_deck_service_proto_depIdxs,
		MessageInfos:      file_grpc_deck_service_proto_msgTypes,
	}.Build()
	File_grpc_deck_service_proto = out.File
	file_grpc_deck_service_proto_rawDesc = nil
	file_grpc_deck_service_proto_goTypes = nil
	file_grpc_deck_service_proto_depIdxs = nil
}
