// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: shorten.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type DeleteURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URLs []string `protobuf:"bytes,1,rep,name=URLs,proto3" json:"URLs,omitempty"`
}

func (x *DeleteURLsRequest) Reset() {
	*x = DeleteURLsRequest{}
	mi := &file_shorten_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteURLsRequest) ProtoMessage() {}

func (x *DeleteURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteURLsRequest.ProtoReflect.Descriptor instead.
func (*DeleteURLsRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteURLsRequest) GetURLs() []string {
	if x != nil {
		return x.URLs
	}
	return nil
}

type GetOriginalURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetOriginalURLRequest) Reset() {
	*x = GetOriginalURLRequest{}
	mi := &file_shorten_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginalURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLRequest) ProtoMessage() {}

func (x *GetOriginalURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalURLRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{1}
}

func (x *GetOriginalURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetAnOriginalURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetAnOriginalURLResponse) Reset() {
	*x = GetAnOriginalURLResponse{}
	mi := &file_shorten_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAnOriginalURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnOriginalURLResponse) ProtoMessage() {}

func (x *GetAnOriginalURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnOriginalURLResponse.ProtoReflect.Descriptor instead.
func (*GetAnOriginalURLResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{2}
}

func (x *GetAnOriginalURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	mi := &file_shorten_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	mi := &file_shorten_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{4}
}

func (x *ShortenResponse) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

type ShortenBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*ShortenBatchRequest_URL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ShortenBatchRequest) Reset() {
	*x = ShortenBatchRequest{}
	mi := &file_shorten_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchRequest) ProtoMessage() {}

func (x *ShortenBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchRequest.ProtoReflect.Descriptor instead.
func (*ShortenBatchRequest) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{5}
}

func (x *ShortenBatchRequest) GetUrls() []*ShortenBatchRequest_URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ShortenBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*ShortenBatchResponse_URL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ShortenBatchResponse) Reset() {
	*x = ShortenBatchResponse{}
	mi := &file_shorten_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchResponse) ProtoMessage() {}

func (x *ShortenBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchResponse.ProtoReflect.Descriptor instead.
func (*ShortenBatchResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{6}
}

func (x *ShortenBatchResponse) GetUrls() []*ShortenBatchResponse_URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersAmount uint32 `protobuf:"varint,1,opt,name=users_amount,json=usersAmount,proto3" json:"users_amount,omitempty"`
	UrlsAmount  uint64 `protobuf:"varint,2,opt,name=urls_amount,json=urlsAmount,proto3" json:"urls_amount,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	mi := &file_shorten_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{7}
}

func (x *StatsResponse) GetUsersAmount() uint32 {
	if x != nil {
		return x.UsersAmount
	}
	return 0
}

func (x *StatsResponse) GetUrlsAmount() uint64 {
	if x != nil {
		return x.UrlsAmount
	}
	return 0
}

type UsersURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*UsersURLsResponse_URL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *UsersURLsResponse) Reset() {
	*x = UsersURLsResponse{}
	mi := &file_shorten_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersURLsResponse) ProtoMessage() {}

func (x *UsersURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersURLsResponse.ProtoReflect.Descriptor instead.
func (*UsersURLsResponse) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{8}
}

func (x *UsersURLsResponse) GetUrls() []*UsersURLsResponse_URL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ShortenBatchRequest_URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ShortenBatchRequest_URL) Reset() {
	*x = ShortenBatchRequest_URL{}
	mi := &file_shorten_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenBatchRequest_URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchRequest_URL) ProtoMessage() {}

func (x *ShortenBatchRequest_URL) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchRequest_URL.ProtoReflect.Descriptor instead.
func (*ShortenBatchRequest_URL) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ShortenBatchRequest_URL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenBatchRequest_URL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenBatchResponse_URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortenUrl    string `protobuf:"bytes,2,opt,name=shorten_url,json=shortenUrl,proto3" json:"shorten_url,omitempty"`
}

func (x *ShortenBatchResponse_URL) Reset() {
	*x = ShortenBatchResponse_URL{}
	mi := &file_shorten_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenBatchResponse_URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenBatchResponse_URL) ProtoMessage() {}

func (x *ShortenBatchResponse_URL) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenBatchResponse_URL.ProtoReflect.Descriptor instead.
func (*ShortenBatchResponse_URL) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ShortenBatchResponse_URL) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenBatchResponse_URL) GetShortenUrl() string {
	if x != nil {
		return x.ShortenUrl
	}
	return ""
}

type UsersURLsResponse_URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short    string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *UsersURLsResponse_URL) Reset() {
	*x = UsersURLsResponse_URL{}
	mi := &file_shorten_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersURLsResponse_URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersURLsResponse_URL) ProtoMessage() {}

func (x *UsersURLsResponse_URL) ProtoReflect() protoreflect.Message {
	mi := &file_shorten_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersURLsResponse_URL.ProtoReflect.Descriptor instead.
func (*UsersURLsResponse_URL) Descriptor() ([]byte, []int) {
	return file_shorten_proto_rawDescGZIP(), []int{8, 0}
}

func (x *UsersURLsResponse_URL) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *UsersURLsResponse_URL) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

var File_shorten_proto protoreflect.FileDescriptor

var file_shorten_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52,
	0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x52, 0x4c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x22, 0x34, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x22, 0x2c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x33, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x1a, 0x4f, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x22, 0x9b, 0x01, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x1a, 0x4d, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x22,
	0x53, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x72, 0x6c, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x72, 0x6c, 0x73, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7f, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x1a, 0x37, 0x0a, 0x03,
	0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x32, 0xe1, 0x03, 0x0a, 0x13, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x42, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shorten_proto_rawDescOnce sync.Once
	file_shorten_proto_rawDescData = file_shorten_proto_rawDesc
)

func file_shorten_proto_rawDescGZIP() []byte {
	file_shorten_proto_rawDescOnce.Do(func() {
		file_shorten_proto_rawDescData = protoimpl.X.CompressGZIP(file_shorten_proto_rawDescData)
	})
	return file_shorten_proto_rawDescData
}

var file_shorten_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_shorten_proto_goTypes = []any{
	(*DeleteURLsRequest)(nil),        // 0: server.DeleteURLsRequest
	(*GetOriginalURLRequest)(nil),    // 1: server.GetOriginalURLRequest
	(*GetAnOriginalURLResponse)(nil), // 2: server.GetAnOriginalURLResponse
	(*ShortenRequest)(nil),           // 3: server.ShortenRequest
	(*ShortenResponse)(nil),          // 4: server.ShortenResponse
	(*ShortenBatchRequest)(nil),      // 5: server.ShortenBatchRequest
	(*ShortenBatchResponse)(nil),     // 6: server.ShortenBatchResponse
	(*StatsResponse)(nil),            // 7: server.StatsResponse
	(*UsersURLsResponse)(nil),        // 8: server.UsersURLsResponse
	(*ShortenBatchRequest_URL)(nil),  // 9: server.ShortenBatchRequest.URL
	(*ShortenBatchResponse_URL)(nil), // 10: server.ShortenBatchResponse.URL
	(*UsersURLsResponse_URL)(nil),    // 11: server.UsersURLsResponse.URL
	(*empty.Empty)(nil),              // 12: google.protobuf.Empty
}
var file_shorten_proto_depIdxs = []int32{
	9,  // 0: server.ShortenBatchRequest.urls:type_name -> server.ShortenBatchRequest.URL
	10, // 1: server.ShortenBatchResponse.urls:type_name -> server.ShortenBatchResponse.URL
	11, // 2: server.UsersURLsResponse.urls:type_name -> server.UsersURLsResponse.URL
	0,  // 3: server.URLShortenerService.DeleteURLs:input_type -> server.DeleteURLsRequest
	1,  // 4: server.URLShortenerService.GetOriginalURL:input_type -> server.GetOriginalURLRequest
	12, // 5: server.URLShortenerService.PingDB:input_type -> google.protobuf.Empty
	3,  // 6: server.URLShortenerService.Shorten:input_type -> server.ShortenRequest
	5,  // 7: server.URLShortenerService.ShortenBatch:input_type -> server.ShortenBatchRequest
	12, // 8: server.URLShortenerService.Stats:input_type -> google.protobuf.Empty
	12, // 9: server.URLShortenerService.UserURLs:input_type -> google.protobuf.Empty
	12, // 10: server.URLShortenerService.DeleteURLs:output_type -> google.protobuf.Empty
	2,  // 11: server.URLShortenerService.GetOriginalURL:output_type -> server.GetAnOriginalURLResponse
	12, // 12: server.URLShortenerService.PingDB:output_type -> google.protobuf.Empty
	4,  // 13: server.URLShortenerService.Shorten:output_type -> server.ShortenResponse
	6,  // 14: server.URLShortenerService.ShortenBatch:output_type -> server.ShortenBatchResponse
	7,  // 15: server.URLShortenerService.Stats:output_type -> server.StatsResponse
	8,  // 16: server.URLShortenerService.UserURLs:output_type -> server.UsersURLsResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_shorten_proto_init() }
func file_shorten_proto_init() {
	if File_shorten_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shorten_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorten_proto_goTypes,
		DependencyIndexes: file_shorten_proto_depIdxs,
		MessageInfos:      file_shorten_proto_msgTypes,
	}.Build()
	File_shorten_proto = out.File
	file_shorten_proto_rawDesc = nil
	file_shorten_proto_goTypes = nil
	file_shorten_proto_depIdxs = nil
}
