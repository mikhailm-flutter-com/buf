// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: buf/alpha/module/v1alpha1/module.proto

package modulev1alpha1

import (
	v1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/breaking/v1"
	v11 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/lint/v1"
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

type DigestType int32

const (
	DigestType_DIGEST_TYPE_UNSPECIFIED DigestType = 0
	DigestType_DIGEST_TYPE_SHAKE256    DigestType = 1
)

// Enum value maps for DigestType.
var (
	DigestType_name = map[int32]string{
		0: "DIGEST_TYPE_UNSPECIFIED",
		1: "DIGEST_TYPE_SHAKE256",
	}
	DigestType_value = map[string]int32{
		"DIGEST_TYPE_UNSPECIFIED": 0,
		"DIGEST_TYPE_SHAKE256":    1,
	}
)

func (x DigestType) Enum() *DigestType {
	p := new(DigestType)
	*p = x
	return p
}

func (x DigestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DigestType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_module_v1alpha1_module_proto_enumTypes[0].Descriptor()
}

func (DigestType) Type() protoreflect.EnumType {
	return &file_buf_alpha_module_v1alpha1_module_proto_enumTypes[0]
}

func (x DigestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DigestType.Descriptor instead.
func (DigestType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{0}
}

// Digest represents a hash function's value.
type Digest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// digest_type describes the hash algorithm. e.g. "SHAKE256"
	DigestType DigestType `protobuf:"varint,1,opt,name=digest_type,json=digestType,proto3,enum=buf.alpha.module.v1alpha1.DigestType" json:"digest_type,omitempty"`
	// digest is the hash's output without encoding.
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *Digest) Reset() {
	*x = Digest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Digest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Digest) ProtoMessage() {}

func (x *Digest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Digest.ProtoReflect.Descriptor instead.
func (*Digest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{0}
}

func (x *Digest) GetDigestType() DigestType {
	if x != nil {
		return x.DigestType
	}
	return DigestType_DIGEST_TYPE_UNSPECIFIED
}

func (x *Digest) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Blob represents some module content with an associated digest.
type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Digest of the content.
	Digest *Digest `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	// Content of the blob.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{1}
}

func (x *Blob) GetDigest() *Digest {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *Blob) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// Module is a module.
type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// files are the files that make up the set.
	//
	// Sorted by path.
	// Path must be unique.
	// Only the target files. No imports.
	//
	// Maximum total size of all content: 32MB.
	Files []*ModuleFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// dependencies are the dependencies.
	Dependencies []*ModulePin `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// documentation is the string representation of the contents of the file at documentation_path.
	//
	// string is used to enforce UTF-8 encoding or 7-bit ASCII text.
	Documentation string `protobuf:"bytes,3,opt,name=documentation,proto3" json:"documentation,omitempty"`
	// breaking_config is the breaking change detection configuration set for the module.
	BreakingConfig *v1.Config `protobuf:"bytes,4,opt,name=breaking_config,json=breakingConfig,proto3" json:"breaking_config,omitempty"`
	// lint_config is the lint configuration set for the module.
	LintConfig *v11.Config `protobuf:"bytes,5,opt,name=lint_config,json=lintConfig,proto3" json:"lint_config,omitempty"`
	// license is the string representation of the contents of the `LICENSE` file.
	//
	// string is used to enforce UTF-8 encoding or 7-bit ASCII text.
	License string `protobuf:"bytes,6,opt,name=license,proto3" json:"license,omitempty"`
	// documentation_path is the path of the file which contains the module documentation.
	//
	// either `buf.md`, `README.md` or `README.markdown`.
	// if empty, assumes buf.md.
	DocumentationPath string `protobuf:"bytes,7,opt,name=documentation_path,json=documentationPath,proto3" json:"documentation_path,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetFiles() []*ModuleFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Module) GetDependencies() []*ModulePin {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Module) GetDocumentation() string {
	if x != nil {
		return x.Documentation
	}
	return ""
}

func (x *Module) GetBreakingConfig() *v1.Config {
	if x != nil {
		return x.BreakingConfig
	}
	return nil
}

func (x *Module) GetLintConfig() *v11.Config {
	if x != nil {
		return x.LintConfig
	}
	return nil
}

func (x *Module) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *Module) GetDocumentationPath() string {
	if x != nil {
		return x.DocumentationPath
	}
	return ""
}

// ModuleFile is a file within a FileSet.
type ModuleFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path is the relative path of the file.
	// Path can only use '/' as the separator character, and includes no ".." components.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// content is the content of the file.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ModuleFile) Reset() {
	*x = ModuleFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleFile) ProtoMessage() {}

func (x *ModuleFile) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleFile.ProtoReflect.Descriptor instead.
func (*ModuleFile) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{3}
}

func (x *ModuleFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ModuleFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// ModuleReference is a module reference.
type ModuleReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remote     string `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	// either tag, or commit
	Reference string `protobuf:"bytes,4,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *ModuleReference) Reset() {
	*x = ModuleReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleReference) ProtoMessage() {}

func (x *ModuleReference) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleReference.ProtoReflect.Descriptor instead.
func (*ModuleReference) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{4}
}

func (x *ModuleReference) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *ModuleReference) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ModuleReference) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ModuleReference) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

// ModulePin is a module pin.
type ModulePin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remote     string                 `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	Owner      string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string                 `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch     string                 `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	Commit     string                 `protobuf:"bytes,5,opt,name=commit,proto3" json:"commit,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Module's manifest digest. Replacement for previous b1/b3 digests.
	// This is in the format '<digest_type>:<digest>`, where '<digest_type>' is the lowercase digest name ('shake256'),
	// and '<digest>' is the lowercase hex-encoded digest.
	// This value is persisted directly to the buf.lock file (https://buf.build/docs/configuration/v1/buf-lock) as the 'digest' key.
	ManifestDigest string `protobuf:"bytes,8,opt,name=manifest_digest,json=manifestDigest,proto3" json:"manifest_digest,omitempty"`
}

func (x *ModulePin) Reset() {
	*x = ModulePin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModulePin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModulePin) ProtoMessage() {}

func (x *ModulePin) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_module_v1alpha1_module_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModulePin.ProtoReflect.Descriptor instead.
func (*ModulePin) Descriptor() ([]byte, []int) {
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP(), []int{5}
}

func (x *ModulePin) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *ModulePin) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ModulePin) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ModulePin) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *ModulePin) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *ModulePin) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ModulePin) GetManifestDigest() string {
	if x != nil {
		return x.ManifestDigest
	}
	return ""
}

var File_buf_alpha_module_v1alpha1_module_proto protoreflect.FileDescriptor

var file_buf_alpha_module_v1alpha1_module_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x22, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62,
	0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x22, 0x5b, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x82, 0x03, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x50, 0x69, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0f, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3a, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x6c, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x3a, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x7d, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0xfd, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x2a,
	0x43, 0x0a, 0x0a, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x44, 0x49, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x49,
	0x47, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x32,
	0x35, 0x36, 0x10, 0x01, 0x42, 0x8a, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42,
	0x41, 0x4d, 0xaa, 0x02, 0x19, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x19, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x42, 0x75, 0x66,
	0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1c, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a,
	0x3a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_module_v1alpha1_module_proto_rawDescOnce sync.Once
	file_buf_alpha_module_v1alpha1_module_proto_rawDescData = file_buf_alpha_module_v1alpha1_module_proto_rawDesc
)

func file_buf_alpha_module_v1alpha1_module_proto_rawDescGZIP() []byte {
	file_buf_alpha_module_v1alpha1_module_proto_rawDescOnce.Do(func() {
		file_buf_alpha_module_v1alpha1_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_module_v1alpha1_module_proto_rawDescData)
	})
	return file_buf_alpha_module_v1alpha1_module_proto_rawDescData
}

var file_buf_alpha_module_v1alpha1_module_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_module_v1alpha1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_buf_alpha_module_v1alpha1_module_proto_goTypes = []interface{}{
	(DigestType)(0),               // 0: buf.alpha.module.v1alpha1.DigestType
	(*Digest)(nil),                // 1: buf.alpha.module.v1alpha1.Digest
	(*Blob)(nil),                  // 2: buf.alpha.module.v1alpha1.Blob
	(*Module)(nil),                // 3: buf.alpha.module.v1alpha1.Module
	(*ModuleFile)(nil),            // 4: buf.alpha.module.v1alpha1.ModuleFile
	(*ModuleReference)(nil),       // 5: buf.alpha.module.v1alpha1.ModuleReference
	(*ModulePin)(nil),             // 6: buf.alpha.module.v1alpha1.ModulePin
	(*v1.Config)(nil),             // 7: buf.alpha.breaking.v1.Config
	(*v11.Config)(nil),            // 8: buf.alpha.lint.v1.Config
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_buf_alpha_module_v1alpha1_module_proto_depIdxs = []int32{
	0, // 0: buf.alpha.module.v1alpha1.Digest.digest_type:type_name -> buf.alpha.module.v1alpha1.DigestType
	1, // 1: buf.alpha.module.v1alpha1.Blob.digest:type_name -> buf.alpha.module.v1alpha1.Digest
	4, // 2: buf.alpha.module.v1alpha1.Module.files:type_name -> buf.alpha.module.v1alpha1.ModuleFile
	6, // 3: buf.alpha.module.v1alpha1.Module.dependencies:type_name -> buf.alpha.module.v1alpha1.ModulePin
	7, // 4: buf.alpha.module.v1alpha1.Module.breaking_config:type_name -> buf.alpha.breaking.v1.Config
	8, // 5: buf.alpha.module.v1alpha1.Module.lint_config:type_name -> buf.alpha.lint.v1.Config
	9, // 6: buf.alpha.module.v1alpha1.ModulePin.create_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_buf_alpha_module_v1alpha1_module_proto_init() }
func file_buf_alpha_module_v1alpha1_module_proto_init() {
	if File_buf_alpha_module_v1alpha1_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Digest); i {
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
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blob); i {
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
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleFile); i {
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
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleReference); i {
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
		file_buf_alpha_module_v1alpha1_module_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModulePin); i {
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
			RawDescriptor: file_buf_alpha_module_v1alpha1_module_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_module_v1alpha1_module_proto_goTypes,
		DependencyIndexes: file_buf_alpha_module_v1alpha1_module_proto_depIdxs,
		EnumInfos:         file_buf_alpha_module_v1alpha1_module_proto_enumTypes,
		MessageInfos:      file_buf_alpha_module_v1alpha1_module_proto_msgTypes,
	}.Build()
	File_buf_alpha_module_v1alpha1_module_proto = out.File
	file_buf_alpha_module_v1alpha1_module_proto_rawDesc = nil
	file_buf_alpha_module_v1alpha1_module_proto_goTypes = nil
	file_buf_alpha_module_v1alpha1_module_proto_depIdxs = nil
}
