// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: articles/articles.proto

package articles

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListArticlesRequest) Reset() {
	*x = ListArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesRequest) ProtoMessage() {}

func (x *ListArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_articles_articles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesRequest.ProtoReflect.Descriptor instead.
func (*ListArticlesRequest) Descriptor() ([]byte, []int) {
	return file_articles_articles_proto_rawDescGZIP(), []int{0}
}

type ListArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ListArticlesResponse) Reset() {
	*x = ListArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticlesResponse) ProtoMessage() {}

func (x *ListArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_articles_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticlesResponse.ProtoReflect.Descriptor instead.
func (*ListArticlesResponse) Descriptor() ([]byte, []int) {
	return file_articles_articles_proto_rawDescGZIP(), []int{1}
}

func (x *ListArticlesResponse) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

// ????????? ????????? ???
type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ???????????????
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// ?????? ??????
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// ?????? ??????
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// ?????? ??????
	Summary string `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	// ?????? ??????
	AuthorName string `protobuf:"bytes,5,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	// ????????? ???????????? ??????
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// ??????????????? ????????? ??????
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// ?????? ??? ??? ?????? ???????????? ??????
	Url string `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_articles_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_articles_articles_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Article) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Article) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Article) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Article) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_articles_articles_proto protoreflect.FileDescriptor

var file_articles_articles_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22,
	0x94, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x47, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1d, 0x5a, 0x1b, 0x63, 0x68, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_articles_articles_proto_rawDescOnce sync.Once
	file_articles_articles_proto_rawDescData = file_articles_articles_proto_rawDesc
)

func file_articles_articles_proto_rawDescGZIP() []byte {
	file_articles_articles_proto_rawDescOnce.Do(func() {
		file_articles_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_articles_articles_proto_rawDescData)
	})
	return file_articles_articles_proto_rawDescData
}

var file_articles_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_articles_articles_proto_goTypes = []interface{}{
	(*ListArticlesRequest)(nil),   // 0: ListArticlesRequest
	(*ListArticlesResponse)(nil),  // 1: ListArticlesResponse
	(*Article)(nil),               // 2: Article
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_articles_articles_proto_depIdxs = []int32{
	2, // 0: ListArticlesResponse.articles:type_name -> Article
	3, // 1: Article.created_time:type_name -> google.protobuf.Timestamp
	3, // 2: Article.updated_time:type_name -> google.protobuf.Timestamp
	0, // 3: Articles.ListArticles:input_type -> ListArticlesRequest
	1, // 4: Articles.ListArticles:output_type -> ListArticlesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_articles_articles_proto_init() }
func file_articles_articles_proto_init() {
	if File_articles_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_articles_articles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticlesRequest); i {
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
		file_articles_articles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticlesResponse); i {
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
		file_articles_articles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
			RawDescriptor: file_articles_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_articles_articles_proto_goTypes,
		DependencyIndexes: file_articles_articles_proto_depIdxs,
		MessageInfos:      file_articles_articles_proto_msgTypes,
	}.Build()
	File_articles_articles_proto = out.File
	file_articles_articles_proto_rawDesc = nil
	file_articles_articles_proto_goTypes = nil
	file_articles_articles_proto_depIdxs = nil
}
