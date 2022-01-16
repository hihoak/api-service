// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hihoak/music_api/v1/music_api.proto

package api_service

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpotifySearchArtistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Genre []string `protobuf:"bytes,2,rep,name=genre,proto3" json:"genre,omitempty"`
}

func (x *SpotifySearchArtistRequest) Reset() {
	*x = SpotifySearchArtistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotifySearchArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotifySearchArtistRequest) ProtoMessage() {}

func (x *SpotifySearchArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotifySearchArtistRequest.ProtoReflect.Descriptor instead.
func (*SpotifySearchArtistRequest) Descriptor() ([]byte, []int) {
	return file_hihoak_music_api_v1_music_api_proto_rawDescGZIP(), []int{0}
}

func (x *SpotifySearchArtistRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpotifySearchArtistRequest) GetGenre() []string {
	if x != nil {
		return x.Genre
	}
	return nil
}

type SpotifySearchArtistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artist *SpotifyArtist `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
}

func (x *SpotifySearchArtistResponse) Reset() {
	*x = SpotifySearchArtistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotifySearchArtistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotifySearchArtistResponse) ProtoMessage() {}

func (x *SpotifySearchArtistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotifySearchArtistResponse.ProtoReflect.Descriptor instead.
func (*SpotifySearchArtistResponse) Descriptor() ([]byte, []int) {
	return file_hihoak_music_api_v1_music_api_proto_rawDescGZIP(), []int{1}
}

func (x *SpotifySearchArtistResponse) GetArtist() *SpotifyArtist {
	if x != nil {
		return x.Artist
	}
	return nil
}

type SpotifySearchTopArtistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genres        []string `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	ExcludeGenres []string `protobuf:"bytes,2,rep,name=exclude_genres,json=excludeGenres,proto3" json:"exclude_genres,omitempty"`
	MinFollowers  int64    `protobuf:"varint,3,opt,name=min_followers,json=minFollowers,proto3" json:"min_followers,omitempty"`
	MinPopularity int64    `protobuf:"varint,4,opt,name=min_popularity,json=minPopularity,proto3" json:"min_popularity,omitempty"`
	Count         int64    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SpotifySearchTopArtistsRequest) Reset() {
	*x = SpotifySearchTopArtistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotifySearchTopArtistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotifySearchTopArtistsRequest) ProtoMessage() {}

func (x *SpotifySearchTopArtistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotifySearchTopArtistsRequest.ProtoReflect.Descriptor instead.
func (*SpotifySearchTopArtistsRequest) Descriptor() ([]byte, []int) {
	return file_hihoak_music_api_v1_music_api_proto_rawDescGZIP(), []int{2}
}

func (x *SpotifySearchTopArtistsRequest) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *SpotifySearchTopArtistsRequest) GetExcludeGenres() []string {
	if x != nil {
		return x.ExcludeGenres
	}
	return nil
}

func (x *SpotifySearchTopArtistsRequest) GetMinFollowers() int64 {
	if x != nil {
		return x.MinFollowers
	}
	return 0
}

func (x *SpotifySearchTopArtistsRequest) GetMinPopularity() int64 {
	if x != nil {
		return x.MinPopularity
	}
	return 0
}

func (x *SpotifySearchTopArtistsRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SpotifySearchTopArtistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artists []*SpotifyArtist `protobuf:"bytes,1,rep,name=artists,proto3" json:"artists,omitempty"`
}

func (x *SpotifySearchTopArtistsResponse) Reset() {
	*x = SpotifySearchTopArtistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotifySearchTopArtistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotifySearchTopArtistsResponse) ProtoMessage() {}

func (x *SpotifySearchTopArtistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotifySearchTopArtistsResponse.ProtoReflect.Descriptor instead.
func (*SpotifySearchTopArtistsResponse) Descriptor() ([]byte, []int) {
	return file_hihoak_music_api_v1_music_api_proto_rawDescGZIP(), []int{3}
}

func (x *SpotifySearchTopArtistsResponse) GetArtists() []*SpotifyArtist {
	if x != nil {
		return x.Artists
	}
	return nil
}

type SpotifyArtist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Followers int64    `protobuf:"varint,2,opt,name=followers,proto3" json:"followers,omitempty"`
	Genres    []string `protobuf:"bytes,3,rep,name=genres,proto3" json:"genres,omitempty"`
	Uri       string   `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *SpotifyArtist) Reset() {
	*x = SpotifyArtist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotifyArtist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotifyArtist) ProtoMessage() {}

func (x *SpotifyArtist) ProtoReflect() protoreflect.Message {
	mi := &file_hihoak_music_api_v1_music_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotifyArtist.ProtoReflect.Descriptor instead.
func (*SpotifyArtist) Descriptor() ([]byte, []int) {
	return file_hihoak_music_api_v1_music_api_proto_rawDescGZIP(), []int{4}
}

func (x *SpotifyArtist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpotifyArtist) GetFollowers() int64 {
	if x != nil {
		return x.Followers
	}
	return 0
}

func (x *SpotifyArtist) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *SpotifyArtist) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

var File_hihoak_music_api_v1_music_api_proto protoreflect.FileDescriptor

var file_hihoak_music_api_v1_music_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x1a, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x1b, 0x53, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x69, 0x68, 0x6f,
	0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x06, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x1e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6d, 0x69, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x1f, 0x53, 0x70, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x22, 0x6b, 0x0a, 0x0d, 0x53, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x32, 0xd0, 0x02, 0x0a, 0x08, 0x4d, 0x75, 0x73, 0x69,
	0x63, 0x41, 0x70, 0x69, 0x12, 0x9a, 0x01, 0x0a, 0x13, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x68,
	0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0xa6, 0x01, 0x0a, 0x17, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x33, 0x2e,
	0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x68, 0x69, 0x68, 0x6f, 0x61, 0x6b, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x22, 0x15, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x2f, 0x74, 0x6f, 0x70, 0x3a, 0x01, 0x2a, 0x42, 0x14, 0x5a, 0x12, 0x68, 0x69,
	0x68, 0x6f, 0x61, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hihoak_music_api_v1_music_api_proto_rawDescOnce sync.Once
	file_hihoak_music_api_v1_music_api_proto_rawDescData = file_hihoak_music_api_v1_music_api_proto_rawDesc
)

func file_hihoak_music_api_v1_music_api_proto_rawDescGZIP() []byte {
	file_hihoak_music_api_v1_music_api_proto_rawDescOnce.Do(func() {
		file_hihoak_music_api_v1_music_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_hihoak_music_api_v1_music_api_proto_rawDescData)
	})
	return file_hihoak_music_api_v1_music_api_proto_rawDescData
}

var file_hihoak_music_api_v1_music_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hihoak_music_api_v1_music_api_proto_goTypes = []interface{}{
	(*SpotifySearchArtistRequest)(nil),      // 0: hihoak.music_api.v1.SpotifySearchArtistRequest
	(*SpotifySearchArtistResponse)(nil),     // 1: hihoak.music_api.v1.SpotifySearchArtistResponse
	(*SpotifySearchTopArtistsRequest)(nil),  // 2: hihoak.music_api.v1.SpotifySearchTopArtistsRequest
	(*SpotifySearchTopArtistsResponse)(nil), // 3: hihoak.music_api.v1.SpotifySearchTopArtistsResponse
	(*SpotifyArtist)(nil),                   // 4: hihoak.music_api.v1.SpotifyArtist
}
var file_hihoak_music_api_v1_music_api_proto_depIdxs = []int32{
	4, // 0: hihoak.music_api.v1.SpotifySearchArtistResponse.artist:type_name -> hihoak.music_api.v1.SpotifyArtist
	4, // 1: hihoak.music_api.v1.SpotifySearchTopArtistsResponse.artists:type_name -> hihoak.music_api.v1.SpotifyArtist
	0, // 2: hihoak.music_api.v1.MusicApi.SpotifySearchArtist:input_type -> hihoak.music_api.v1.SpotifySearchArtistRequest
	2, // 3: hihoak.music_api.v1.MusicApi.SpotifySearchTopArtists:input_type -> hihoak.music_api.v1.SpotifySearchTopArtistsRequest
	1, // 4: hihoak.music_api.v1.MusicApi.SpotifySearchArtist:output_type -> hihoak.music_api.v1.SpotifySearchArtistResponse
	3, // 5: hihoak.music_api.v1.MusicApi.SpotifySearchTopArtists:output_type -> hihoak.music_api.v1.SpotifySearchTopArtistsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hihoak_music_api_v1_music_api_proto_init() }
func file_hihoak_music_api_v1_music_api_proto_init() {
	if File_hihoak_music_api_v1_music_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hihoak_music_api_v1_music_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotifySearchArtistRequest); i {
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
		file_hihoak_music_api_v1_music_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotifySearchArtistResponse); i {
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
		file_hihoak_music_api_v1_music_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotifySearchTopArtistsRequest); i {
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
		file_hihoak_music_api_v1_music_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotifySearchTopArtistsResponse); i {
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
		file_hihoak_music_api_v1_music_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotifyArtist); i {
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
			RawDescriptor: file_hihoak_music_api_v1_music_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hihoak_music_api_v1_music_api_proto_goTypes,
		DependencyIndexes: file_hihoak_music_api_v1_music_api_proto_depIdxs,
		MessageInfos:      file_hihoak_music_api_v1_music_api_proto_msgTypes,
	}.Build()
	File_hihoak_music_api_v1_music_api_proto = out.File
	file_hihoak_music_api_v1_music_api_proto_rawDesc = nil
	file_hihoak_music_api_v1_music_api_proto_goTypes = nil
	file_hihoak_music_api_v1_music_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MusicApiClient is the client API for MusicApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MusicApiClient interface {
	SpotifySearchArtist(ctx context.Context, in *SpotifySearchArtistRequest, opts ...grpc.CallOption) (*SpotifySearchArtistResponse, error)
	SpotifySearchTopArtists(ctx context.Context, in *SpotifySearchTopArtistsRequest, opts ...grpc.CallOption) (*SpotifySearchTopArtistsResponse, error)
}

type musicApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicApiClient(cc grpc.ClientConnInterface) MusicApiClient {
	return &musicApiClient{cc}
}

func (c *musicApiClient) SpotifySearchArtist(ctx context.Context, in *SpotifySearchArtistRequest, opts ...grpc.CallOption) (*SpotifySearchArtistResponse, error) {
	out := new(SpotifySearchArtistResponse)
	err := c.cc.Invoke(ctx, "/hihoak.music_api.v1.MusicApi/SpotifySearchArtist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicApiClient) SpotifySearchTopArtists(ctx context.Context, in *SpotifySearchTopArtistsRequest, opts ...grpc.CallOption) (*SpotifySearchTopArtistsResponse, error) {
	out := new(SpotifySearchTopArtistsResponse)
	err := c.cc.Invoke(ctx, "/hihoak.music_api.v1.MusicApi/SpotifySearchTopArtists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicApiServer is the server API for MusicApi service.
type MusicApiServer interface {
	SpotifySearchArtist(context.Context, *SpotifySearchArtistRequest) (*SpotifySearchArtistResponse, error)
	SpotifySearchTopArtists(context.Context, *SpotifySearchTopArtistsRequest) (*SpotifySearchTopArtistsResponse, error)
}

// UnimplementedMusicApiServer can be embedded to have forward compatible implementations.
type UnimplementedMusicApiServer struct {
}

func (*UnimplementedMusicApiServer) SpotifySearchArtist(context.Context, *SpotifySearchArtistRequest) (*SpotifySearchArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpotifySearchArtist not implemented")
}
func (*UnimplementedMusicApiServer) SpotifySearchTopArtists(context.Context, *SpotifySearchTopArtistsRequest) (*SpotifySearchTopArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpotifySearchTopArtists not implemented")
}

func RegisterMusicApiServer(s *grpc.Server, srv MusicApiServer) {
	s.RegisterService(&_MusicApi_serviceDesc, srv)
}

func _MusicApi_SpotifySearchArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotifySearchArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicApiServer).SpotifySearchArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hihoak.music_api.v1.MusicApi/SpotifySearchArtist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicApiServer).SpotifySearchArtist(ctx, req.(*SpotifySearchArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicApi_SpotifySearchTopArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpotifySearchTopArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicApiServer).SpotifySearchTopArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hihoak.music_api.v1.MusicApi/SpotifySearchTopArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicApiServer).SpotifySearchTopArtists(ctx, req.(*SpotifySearchTopArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MusicApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hihoak.music_api.v1.MusicApi",
	HandlerType: (*MusicApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SpotifySearchArtist",
			Handler:    _MusicApi_SpotifySearchArtist_Handler,
		},
		{
			MethodName: "SpotifySearchTopArtists",
			Handler:    _MusicApi_SpotifySearchTopArtists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hihoak/music_api/v1/music_api.proto",
}