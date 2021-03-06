// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.2
// source: pb_get_book.proto

package book

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   int64   `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BookName string  `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	Stock    int32   `protobuf:"varint,3,opt,name=stock,proto3" json:"stock,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_get_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_pb_get_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_pb_get_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *Book) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *Book) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Book) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId int64 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_get_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_get_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_pb_get_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookRequest) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_get_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_get_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_pb_get_book_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_pb_get_book_proto protoreflect.FileDescriptor

var file_pb_get_book_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x67,
	0x65, 0x6e, 0x22, 0x68, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x0e,
	0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x67, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_get_book_proto_rawDescOnce sync.Once
	file_pb_get_book_proto_rawDescData = file_pb_get_book_proto_rawDesc
)

func file_pb_get_book_proto_rawDescGZIP() []byte {
	file_pb_get_book_proto_rawDescOnce.Do(func() {
		file_pb_get_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_get_book_proto_rawDescData)
	})
	return file_pb_get_book_proto_rawDescData
}

var file_pb_get_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_get_book_proto_goTypes = []interface{}{
	(*Book)(nil),            // 0: com.bookstore.pb_gen.book_gen.Book
	(*GetBookRequest)(nil),  // 1: com.bookstore.pb_gen.book_gen.getBookRequest
	(*GetBookResponse)(nil), // 2: com.bookstore.pb_gen.book_gen.getBookResponse
}
var file_pb_get_book_proto_depIdxs = []int32{
	0, // 0: com.bookstore.pb_gen.book_gen.getBookResponse.book:type_name -> com.bookstore.pb_gen.book_gen.Book
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_get_book_proto_init() }
func file_pb_get_book_proto_init() {
	if File_pb_get_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_get_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_pb_get_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
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
		file_pb_get_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResponse); i {
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
			RawDescriptor: file_pb_get_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_get_book_proto_goTypes,
		DependencyIndexes: file_pb_get_book_proto_depIdxs,
		MessageInfos:      file_pb_get_book_proto_msgTypes,
	}.Build()
	File_pb_get_book_proto = out.File
	file_pb_get_book_proto_rawDesc = nil
	file_pb_get_book_proto_goTypes = nil
	file_pb_get_book_proto_depIdxs = nil
}
