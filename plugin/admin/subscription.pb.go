// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: subscription.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubFilterType int32

const (
	SubFilterType_SUB_FILTER_TYPE_SYS_UNSPECIFIED SubFilterType = 0
	SubFilterType_SUB_FILTER_TYPE_SYS             SubFilterType = 1
	SubFilterType_SUB_FILTER_TYPE_SHARED          SubFilterType = 2
	SubFilterType_SUB_FILTER_TYPE_NON_SHARED      SubFilterType = 3
)

// Enum value maps for SubFilterType.
var (
	SubFilterType_name = map[int32]string{
		0: "SUB_FILTER_TYPE_SYS_UNSPECIFIED",
		1: "SUB_FILTER_TYPE_SYS",
		2: "SUB_FILTER_TYPE_SHARED",
		3: "SUB_FILTER_TYPE_NON_SHARED",
	}
	SubFilterType_value = map[string]int32{
		"SUB_FILTER_TYPE_SYS_UNSPECIFIED": 0,
		"SUB_FILTER_TYPE_SYS":             1,
		"SUB_FILTER_TYPE_SHARED":          2,
		"SUB_FILTER_TYPE_NON_SHARED":      3,
	}
)

func (x SubFilterType) Enum() *SubFilterType {
	p := new(SubFilterType)
	*p = x
	return p
}

func (x SubFilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubFilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[0].Descriptor()
}

func (SubFilterType) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[0]
}

func (x SubFilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubFilterType.Descriptor instead.
func (SubFilterType) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

type SubMatchType int32

const (
	SubMatchType_SUB_MATCH_TYPE_MATCH_UNSPECIFIED SubMatchType = 0
	SubMatchType_SUB_MATCH_TYPE_MATCH_NAME        SubMatchType = 1
	SubMatchType_SUB_MATCH_TYPE_MATCH_FILTER      SubMatchType = 2
)

// Enum value maps for SubMatchType.
var (
	SubMatchType_name = map[int32]string{
		0: "SUB_MATCH_TYPE_MATCH_UNSPECIFIED",
		1: "SUB_MATCH_TYPE_MATCH_NAME",
		2: "SUB_MATCH_TYPE_MATCH_FILTER",
	}
	SubMatchType_value = map[string]int32{
		"SUB_MATCH_TYPE_MATCH_UNSPECIFIED": 0,
		"SUB_MATCH_TYPE_MATCH_NAME":        1,
		"SUB_MATCH_TYPE_MATCH_FILTER":      2,
	}
)

func (x SubMatchType) Enum() *SubMatchType {
	p := new(SubMatchType)
	*p = x
	return p
}

func (x SubMatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubMatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_subscription_proto_enumTypes[1].Descriptor()
}

func (SubMatchType) Type() protoreflect.EnumType {
	return &file_subscription_proto_enumTypes[1]
}

func (x SubMatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubMatchType.Descriptor instead.
func (SubMatchType) EnumDescriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

type ListSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page     uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListSubscriptionRequest) Reset() {
	*x = ListSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionRequest) ProtoMessage() {}

func (x *ListSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *ListSubscriptionRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSubscriptionRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subIndexer,proto3" json:"subIndexer,omitempty"`
	TotalCount    uint32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListSubscriptionResponse) Reset() {
	*x = ListSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionResponse) ProtoMessage() {}

func (x *ListSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubscriptionResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *ListSubscriptionResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type FilterSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If Set, only filter the subIndexer that belongs to the client.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// filter_type indicates what kinds of topics are going to filter.
	// If there are multiple types, use ',' to separate. e.g : 1,2
	// There are 3 kinds of topic can be filtered, defined by SubFilterType:
	// 1 = System Topic(begin with '$')
	// 2 = Shared Topic
	// 3 = NonShared Topic
	FilterType string `protobuf:"bytes,2,opt,name=filter_type,json=filterType,proto3" json:"filter_type,omitempty"`
	// If 1 (SUB_MATCH_TYPE_MATCH_NAME), the server will return subIndexer which has the same topic name with request topic_name.
	// If 2 (SUB_MATCH_TYPE_MATCH_FILTER),the server will return subIndexer which match the request topic_name .
	// match_type must be Set when filter_type is not empty.
	MatchType SubMatchType `protobuf:"varint,3,opt,name=match_type,json=matchType,proto3,enum=gmqtt.admin.api.SubMatchType" json:"match_type,omitempty"`
	// topic_name must be Set when match_type is not zero.
	TopicName string `protobuf:"bytes,4,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// The maximum subIndexer can be returned.
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FilterSubscriptionRequest) Reset() {
	*x = FilterSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSubscriptionRequest) ProtoMessage() {}

func (x *FilterSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*FilterSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *FilterSubscriptionRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *FilterSubscriptionRequest) GetFilterType() string {
	if x != nil {
		return x.FilterType
	}
	return ""
}

func (x *FilterSubscriptionRequest) GetMatchType() SubMatchType {
	if x != nil {
		return x.MatchType
	}
	return SubMatchType_SUB_MATCH_TYPE_MATCH_UNSPECIFIED
}

func (x *FilterSubscriptionRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *FilterSubscriptionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FilterSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subIndexer,proto3" json:"subIndexer,omitempty"`
}

func (x *FilterSubscriptionResponse) Reset() {
	*x = FilterSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSubscriptionResponse) ProtoMessage() {}

func (x *FilterSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*FilterSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *FilterSubscriptionResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId      string          `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Subscriptions []*Subscription `protobuf:"bytes,2,rep,name=subIndexer,proto3" json:"subIndexer,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SubscribeRequest) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// indicates whether it is a new subscription or the subscription is already existed.
	New []bool `protobuf:"varint,1,rep,packed,name=new,proto3" json:"new,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeResponse) GetNew() []bool {
	if x != nil {
		return x.New
	}
	return nil
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Topics   []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UnsubscribeRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName         string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	Id                uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Qos               uint32 `protobuf:"varint,3,opt,name=qos,proto3" json:"qos,omitempty"`
	NoLocal           bool   `protobuf:"varint,4,opt,name=no_local,json=noLocal,proto3" json:"no_local,omitempty"`
	RetainAsPublished bool   `protobuf:"varint,5,opt,name=retain_as_published,json=retainAsPublished,proto3" json:"retain_as_published,omitempty"`
	RetainHandling    uint32 `protobuf:"varint,6,opt,name=retain_handling,json=retainHandling,proto3" json:"retain_handling,omitempty"`
	ClientId          string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *Subscription) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *Subscription) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subscription) GetQos() uint32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *Subscription) GetNoLocal() bool {
	if x != nil {
		return x.NoLocal
	}
	return false
}

func (x *Subscription) GetRetainAsPublished() bool {
	if x != nil {
		return x.RetainAsPublished
	}
	return false
}

func (x *Subscription) GetRetainHandling() uint32 {
	if x != nil {
		return x.RetainHandling
	}
	return 0
}

func (x *Subscription) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xcc, 0x01, 0x0a, 0x19, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x61,
	0x0a, 0x1a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x74, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x43, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6d, 0x71, 0x74,
	0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x65, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x22, 0x49,
	0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e,
	0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e,
	0x5f, 0x61, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x41, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e,
	0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0x89, 0x01, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x1f, 0x53, 0x55, 0x42, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x55, 0x42, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x55, 0x42, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x55, 0x42, 0x5f,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x74, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x55, 0x42, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d,
	0x0a, 0x19, 0x53, 0x55, 0x42, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x55, 0x42, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x02, 0x32, 0xe9,
	0x03, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28,
	0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x83,
	0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x67, 0x6d, 0x71, 0x74,
	0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x21, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x66, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x23, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_subscription_proto_goTypes = []interface{}{
	(SubFilterType)(0),                 // 0: gmqtt.admin.api.SubFilterType
	(SubMatchType)(0),                  // 1: gmqtt.admin.api.SubMatchType
	(*ListSubscriptionRequest)(nil),    // 2: gmqtt.admin.api.ListSubscriptionRequest
	(*ListSubscriptionResponse)(nil),   // 3: gmqtt.admin.api.ListSubscriptionResponse
	(*FilterSubscriptionRequest)(nil),  // 4: gmqtt.admin.api.FilterSubscriptionRequest
	(*FilterSubscriptionResponse)(nil), // 5: gmqtt.admin.api.FilterSubscriptionResponse
	(*SubscribeRequest)(nil),           // 6: gmqtt.admin.api.SubscribeRequest
	(*SubscribeResponse)(nil),          // 7: gmqtt.admin.api.SubscribeResponse
	(*UnsubscribeRequest)(nil),         // 8: gmqtt.admin.api.UnsubscribeRequest
	(*Subscription)(nil),               // 9: gmqtt.admin.api.Subscription
	(*empty.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_subscription_proto_depIdxs = []int32{
	9,  // 0: gmqtt.admin.api.ListSubscriptionResponse.subIndexer:type_name -> gmqtt.admin.api.Subscription
	1,  // 1: gmqtt.admin.api.FilterSubscriptionRequest.match_type:type_name -> gmqtt.admin.api.SubMatchType
	9,  // 2: gmqtt.admin.api.FilterSubscriptionResponse.subIndexer:type_name -> gmqtt.admin.api.Subscription
	9,  // 3: gmqtt.admin.api.SubscribeRequest.subIndexer:type_name -> gmqtt.admin.api.Subscription
	2,  // 4: gmqtt.admin.api.SubscriptionService.List:input_type -> gmqtt.admin.api.ListSubscriptionRequest
	4,  // 5: gmqtt.admin.api.SubscriptionService.Filter:input_type -> gmqtt.admin.api.FilterSubscriptionRequest
	6,  // 6: gmqtt.admin.api.SubscriptionService.Subscribe:input_type -> gmqtt.admin.api.SubscribeRequest
	8,  // 7: gmqtt.admin.api.SubscriptionService.Unsubscribe:input_type -> gmqtt.admin.api.UnsubscribeRequest
	3,  // 8: gmqtt.admin.api.SubscriptionService.List:output_type -> gmqtt.admin.api.ListSubscriptionResponse
	5,  // 9: gmqtt.admin.api.SubscriptionService.Filter:output_type -> gmqtt.admin.api.FilterSubscriptionResponse
	7,  // 10: gmqtt.admin.api.SubscriptionService.Subscribe:output_type -> gmqtt.admin.api.SubscribeResponse
	10, // 11: gmqtt.admin.api.SubscriptionService.Unsubscribe:output_type -> google.protobuf.Empty
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterSubscriptionRequest); i {
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
		file_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterSubscriptionResponse); i {
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
		file_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_subscription_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		EnumInfos:         file_subscription_proto_enumTypes,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}
