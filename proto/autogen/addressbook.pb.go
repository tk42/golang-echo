// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package __addresspb is a generated protocol buffer package.

package は、プロジェクト間での名前衝突を防ぐためのパッケージ名。
このパッケージ名は各言語に応じた解釈が行われる。
(例)
  C++    : C++の名前空間になる
  C#     : パスカル形式に変換後にC#の名前空間になる(csharp_namespaceの指定が無い場合)
  Go     : Goのパッケージ名になる(go_packageの指定が無い場合)
  Python : 無視される

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package __addresspb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 列挙型の定義
type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// メッセージ(Person)の定義
type Person struct {
	// メッセージのフィールド。
	// 各フィールドの識別子として 1, 2... というフィールド番号(タグ)が必要。
	// → シリアライズされたデータでは、フィールド番号でフィールドを識別するため。
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	// repeatedは配列(要素の数は任意。0個でもよい。)
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	// Import した *.proto ファイルで定義された型
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// メッセージの中に別のメッセージの定義を含められる(定義のネスト)
type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=addresspb.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// メッセージ(AddressBook)の定義
// AddressBook は複数の Person を含む。
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "addresspb.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "addresspb.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "addresspb.AddressBook")
	proto.RegisterEnum("addresspb.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x9a, 0x06, 0x3b, 0xd1, 0xd2, 0x0e, 0x22, 0xa1, 0x22, 0x86, 0x9e, 0x22, 0xc2,
	0x16, 0x2a, 0x82, 0x20, 0x1e, 0x2c, 0x14, 0x14, 0xad, 0x2d, 0x4b, 0x55, 0xf0, 0x22, 0x09, 0x19,
	0x6b, 0x68, 0x92, 0x5d, 0x92, 0xf4, 0xd0, 0x1f, 0xe7, 0x7f, 0x93, 0xec, 0xa6, 0x41, 0x10, 0x6f,
	0xb3, 0xf3, 0x3e, 0xde, 0xbe, 0x37, 0xd0, 0x0f, 0xa2, 0x28, 0xa7, 0xa2, 0x08, 0x85, 0x58, 0x33,
	0x99, 0x8b, 0x52, 0x60, 0xa7, 0x5e, 0xc9, 0x70, 0x70, 0xb6, 0x12, 0x62, 0x95, 0xd0, 0x48, 0x09,
	0xe1, 0xe6, 0x73, 0x54, 0xc6, 0x29, 0x15, 0x65, 0x90, 0x4a, 0xcd, 0x0e, 0xbf, 0x4d, 0xb0, 0x17,
	0x94, 0x17, 0x22, 0x43, 0x04, 0x2b, 0x0b, 0x52, 0x72, 0x0d, 0xcf, 0xf0, 0x3b, 0x5c, 0xcd, 0xd8,
	0x05, 0x33, 0x8e, 0x5c, 0xd3, 0x33, 0xfc, 0x36, 0x37, 0xe3, 0x08, 0x8f, 0xa0, 0x4d, 0x69, 0x10,
	0x27, 0x6e, 0x4b, 0x41, 0xfa, 0x81, 0x57, 0x60, 0xcb, 0x2f, 0x91, 0x51, 0xe1, 0x5a, 0x5e, 0xcb,
	0x77, 0xc6, 0xa7, 0xac, 0x49, 0xc0, 0xb4, 0x39, 0x5b, 0x54, 0xfa, 0xf3, 0x26, 0x0d, 0x29, 0xe7,
	0x35, 0x8c, 0xb7, 0x70, 0x90, 0x04, 0x45, 0xf9, 0xb1, 0x91, 0x51, 0x50, 0x52, 0xe4, 0xb6, 0x3d,
	0xc3, 0x77, 0xc6, 0x03, 0xa6, 0x33, 0xb3, 0x5d, 0x66, 0xb6, 0xdc, 0x65, 0xe6, 0x4e, 0xc5, 0xbf,
	0x68, 0x7c, 0xf0, 0x0a, 0xce, 0x2f, 0x57, 0x3c, 0x06, 0x3b, 0x53, 0x53, 0x5d, 0xa0, 0x7e, 0xe1,
	0x08, 0xac, 0x72, 0x2b, 0x49, 0x95, 0xe8, 0x8e, 0x4f, 0xfe, 0x89, 0xb6, 0xdc, 0x4a, 0xe2, 0x0a,
	0x1c, 0x5e, 0x40, 0xa7, 0x59, 0x21, 0x80, 0x3d, 0x9b, 0x4f, 0x1e, 0x9e, 0xa6, 0xbd, 0x3d, 0xdc,
	0x07, 0xeb, 0x7e, 0x3e, 0x9b, 0xf6, 0x8c, 0x6a, 0x7a, 0x9b, 0xf3, 0xc7, 0x9e, 0x39, 0xbc, 0x06,
	0xe7, 0x4e, 0x1b, 0x4e, 0x84, 0x58, 0xe3, 0x39, 0xd8, 0x92, 0x84, 0x4c, 0xaa, 0x2b, 0x56, 0x97,
	0xe8, 0xff, 0xf9, 0x8e, 0xd7, 0xc0, 0xe4, 0xf0, 0xdd, 0x61, 0x37, 0x8d, 0x1a, 0xda, 0xaa, 0xee,
	0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x04, 0xde, 0x4c, 0xd0, 0x01, 0x00, 0x00,
}