// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.19.4
// source: rpc/cache/service.proto

package cache

import (
	common "github.com/khulnasoft/tunnel/rpc/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ArtifactInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaVersion   int32                  `protobuf:"varint,1,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	Architecture    string                 `protobuf:"bytes,2,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Created         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	DockerVersion   string                 `protobuf:"bytes,4,opt,name=docker_version,json=dockerVersion,proto3" json:"docker_version,omitempty"`
	Os              string                 `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	HistoryPackages []*common.Package      `protobuf:"bytes,6,rep,name=history_packages,json=historyPackages,proto3" json:"history_packages,omitempty"`
}

func (x *ArtifactInfo) Reset() {
	*x = ArtifactInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactInfo) ProtoMessage() {}

func (x *ArtifactInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactInfo.ProtoReflect.Descriptor instead.
func (*ArtifactInfo) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactInfo) GetSchemaVersion() int32 {
	if x != nil {
		return x.SchemaVersion
	}
	return 0
}

func (x *ArtifactInfo) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *ArtifactInfo) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ArtifactInfo) GetDockerVersion() string {
	if x != nil {
		return x.DockerVersion
	}
	return ""
}

func (x *ArtifactInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ArtifactInfo) GetHistoryPackages() []*common.Package {
	if x != nil {
		return x.HistoryPackages
	}
	return nil
}

type PutArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactId   string        `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	ArtifactInfo *ArtifactInfo `protobuf:"bytes,2,opt,name=artifact_info,json=artifactInfo,proto3" json:"artifact_info,omitempty"`
}

func (x *PutArtifactRequest) Reset() {
	*x = PutArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutArtifactRequest) ProtoMessage() {}

func (x *PutArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutArtifactRequest.ProtoReflect.Descriptor instead.
func (*PutArtifactRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{1}
}

func (x *PutArtifactRequest) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *PutArtifactRequest) GetArtifactInfo() *ArtifactInfo {
	if x != nil {
		return x.ArtifactInfo
	}
	return nil
}

type BlobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaVersion     int32                      `protobuf:"varint,1,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	Os                *common.OS                 `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Repository        *common.Repository         `protobuf:"bytes,11,opt,name=repository,proto3" json:"repository,omitempty"`
	PackageInfos      []*common.PackageInfo      `protobuf:"bytes,3,rep,name=package_infos,json=packageInfos,proto3" json:"package_infos,omitempty"`
	Applications      []*common.Application      `protobuf:"bytes,4,rep,name=applications,proto3" json:"applications,omitempty"`
	Misconfigurations []*common.Misconfiguration `protobuf:"bytes,9,rep,name=misconfigurations,proto3" json:"misconfigurations,omitempty"`
	OpaqueDirs        []string                   `protobuf:"bytes,5,rep,name=opaque_dirs,json=opaqueDirs,proto3" json:"opaque_dirs,omitempty"`
	WhiteoutFiles     []string                   `protobuf:"bytes,6,rep,name=whiteout_files,json=whiteoutFiles,proto3" json:"whiteout_files,omitempty"`
	Digest            string                     `protobuf:"bytes,7,opt,name=digest,proto3" json:"digest,omitempty"`
	DiffId            string                     `protobuf:"bytes,8,opt,name=diff_id,json=diffId,proto3" json:"diff_id,omitempty"`
	CustomResources   []*common.CustomResource   `protobuf:"bytes,10,rep,name=custom_resources,json=customResources,proto3" json:"custom_resources,omitempty"`
	Secrets           []*common.Secret           `protobuf:"bytes,12,rep,name=secrets,proto3" json:"secrets,omitempty"`
	Licenses          []*common.LicenseFile      `protobuf:"bytes,13,rep,name=licenses,proto3" json:"licenses,omitempty"`
}

func (x *BlobInfo) Reset() {
	*x = BlobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobInfo) ProtoMessage() {}

func (x *BlobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobInfo.ProtoReflect.Descriptor instead.
func (*BlobInfo) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{2}
}

func (x *BlobInfo) GetSchemaVersion() int32 {
	if x != nil {
		return x.SchemaVersion
	}
	return 0
}

func (x *BlobInfo) GetOs() *common.OS {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *BlobInfo) GetRepository() *common.Repository {
	if x != nil {
		return x.Repository
	}
	return nil
}

func (x *BlobInfo) GetPackageInfos() []*common.PackageInfo {
	if x != nil {
		return x.PackageInfos
	}
	return nil
}

func (x *BlobInfo) GetApplications() []*common.Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *BlobInfo) GetMisconfigurations() []*common.Misconfiguration {
	if x != nil {
		return x.Misconfigurations
	}
	return nil
}

func (x *BlobInfo) GetOpaqueDirs() []string {
	if x != nil {
		return x.OpaqueDirs
	}
	return nil
}

func (x *BlobInfo) GetWhiteoutFiles() []string {
	if x != nil {
		return x.WhiteoutFiles
	}
	return nil
}

func (x *BlobInfo) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *BlobInfo) GetDiffId() string {
	if x != nil {
		return x.DiffId
	}
	return ""
}

func (x *BlobInfo) GetCustomResources() []*common.CustomResource {
	if x != nil {
		return x.CustomResources
	}
	return nil
}

func (x *BlobInfo) GetSecrets() []*common.Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *BlobInfo) GetLicenses() []*common.LicenseFile {
	if x != nil {
		return x.Licenses
	}
	return nil
}

type PutBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiffId   string    `protobuf:"bytes,1,opt,name=diff_id,json=diffId,proto3" json:"diff_id,omitempty"`
	BlobInfo *BlobInfo `protobuf:"bytes,3,opt,name=blob_info,json=blobInfo,proto3" json:"blob_info,omitempty"`
}

func (x *PutBlobRequest) Reset() {
	*x = PutBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBlobRequest) ProtoMessage() {}

func (x *PutBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBlobRequest.ProtoReflect.Descriptor instead.
func (*PutBlobRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{3}
}

func (x *PutBlobRequest) GetDiffId() string {
	if x != nil {
		return x.DiffId
	}
	return ""
}

func (x *PutBlobRequest) GetBlobInfo() *BlobInfo {
	if x != nil {
		return x.BlobInfo
	}
	return nil
}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os   *common.OS `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Eosl bool       `protobuf:"varint,2,opt,name=eosl,proto3" json:"eosl,omitempty"`
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{4}
}

func (x *PutResponse) GetOs() *common.OS {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *PutResponse) GetEosl() bool {
	if x != nil {
		return x.Eosl
	}
	return false
}

type MissingBlobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactId string   `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	BlobIds    []string `protobuf:"bytes,2,rep,name=blob_ids,json=blobIds,proto3" json:"blob_ids,omitempty"`
}

func (x *MissingBlobsRequest) Reset() {
	*x = MissingBlobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissingBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingBlobsRequest) ProtoMessage() {}

func (x *MissingBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingBlobsRequest.ProtoReflect.Descriptor instead.
func (*MissingBlobsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{5}
}

func (x *MissingBlobsRequest) GetArtifactId() string {
	if x != nil {
		return x.ArtifactId
	}
	return ""
}

func (x *MissingBlobsRequest) GetBlobIds() []string {
	if x != nil {
		return x.BlobIds
	}
	return nil
}

type MissingBlobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissingArtifact bool     `protobuf:"varint,1,opt,name=missing_artifact,json=missingArtifact,proto3" json:"missing_artifact,omitempty"`
	MissingBlobIds  []string `protobuf:"bytes,2,rep,name=missing_blob_ids,json=missingBlobIds,proto3" json:"missing_blob_ids,omitempty"`
}

func (x *MissingBlobsResponse) Reset() {
	*x = MissingBlobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissingBlobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingBlobsResponse) ProtoMessage() {}

func (x *MissingBlobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingBlobsResponse.ProtoReflect.Descriptor instead.
func (*MissingBlobsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{6}
}

func (x *MissingBlobsResponse) GetMissingArtifact() bool {
	if x != nil {
		return x.MissingArtifact
	}
	return false
}

func (x *MissingBlobsResponse) GetMissingBlobIds() []string {
	if x != nil {
		return x.MissingBlobIds
	}
	return nil
}

type DeleteBlobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlobIds []string `protobuf:"bytes,1,rep,name=blob_ids,json=blobIds,proto3" json:"blob_ids,omitempty"`
}

func (x *DeleteBlobsRequest) Reset() {
	*x = DeleteBlobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cache_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlobsRequest) ProtoMessage() {}

func (x *DeleteBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cache_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlobsRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlobsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cache_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBlobsRequest) GetBlobIds() []string {
	if x != nil {
		return x.BlobIds
	}
	return nil
}

var File_rpc_cache_service_proto protoreflect.FileDescriptor

var file_rpc_cache_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x72, 0x69, 0x76, 0x79,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x40, 0x0a, 0x10, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x78, 0x0a, 0x12,
	0x50, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x69,
	0x76, 0x79, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x83, 0x05, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x62, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x38, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x11, 0x6d, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x6d, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x5f, 0x64, 0x69,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x44, 0x69, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6f, 0x75, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x68,
	0x69, 0x74, 0x65, 0x6f, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x10,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x0e,
	0x50, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x62, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x69,
	0x76, 0x79, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x43,
	0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x69, 0x76,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x6f, 0x73, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x65,
	0x6f, 0x73, 0x6c, 0x22, 0x51, 0x0a, 0x13, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x6c,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6c, 0x6f, 0x62, 0x49, 0x64, 0x73, 0x22, 0x6b, 0x0a, 0x14, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x62,
	0x49, 0x64, 0x73, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x62, 0x49, 0x64, 0x73, 0x32, 0xbb, 0x02, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x49,
	0x0a, 0x0b, 0x50, 0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x22, 0x2e,
	0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x07, 0x50, 0x75, 0x74,
	0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x59, 0x0a, 0x0c,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x23, 0x2e, 0x74,
	0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x69, 0x76, 0x79, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x71, 0x75, 0x61, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72,
	0x69, 0x76, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x3b, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_cache_service_proto_rawDescOnce sync.Once
	file_rpc_cache_service_proto_rawDescData = file_rpc_cache_service_proto_rawDesc
)

func file_rpc_cache_service_proto_rawDescGZIP() []byte {
	file_rpc_cache_service_proto_rawDescOnce.Do(func() {
		file_rpc_cache_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_cache_service_proto_rawDescData)
	})
	return file_rpc_cache_service_proto_rawDescData
}

var file_rpc_cache_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_cache_service_proto_goTypes = []interface{}{
	(*ArtifactInfo)(nil),            // 0: tunnel.cache.v1.ArtifactInfo
	(*PutArtifactRequest)(nil),      // 1: tunnel.cache.v1.PutArtifactRequest
	(*BlobInfo)(nil),                // 2: tunnel.cache.v1.BlobInfo
	(*PutBlobRequest)(nil),          // 3: tunnel.cache.v1.PutBlobRequest
	(*PutResponse)(nil),             // 4: tunnel.cache.v1.PutResponse
	(*MissingBlobsRequest)(nil),     // 5: tunnel.cache.v1.MissingBlobsRequest
	(*MissingBlobsResponse)(nil),    // 6: tunnel.cache.v1.MissingBlobsResponse
	(*DeleteBlobsRequest)(nil),      // 7: tunnel.cache.v1.DeleteBlobsRequest
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
	(*common.Package)(nil),          // 9: tunnel.common.Package
	(*common.OS)(nil),               // 10: tunnel.common.OS
	(*common.Repository)(nil),       // 11: tunnel.common.Repository
	(*common.PackageInfo)(nil),      // 12: tunnel.common.PackageInfo
	(*common.Application)(nil),      // 13: tunnel.common.Application
	(*common.Misconfiguration)(nil), // 14: tunnel.common.Misconfiguration
	(*common.CustomResource)(nil),   // 15: tunnel.common.CustomResource
	(*common.Secret)(nil),           // 16: tunnel.common.Secret
	(*common.LicenseFile)(nil),      // 17: tunnel.common.LicenseFile
	(*emptypb.Empty)(nil),           // 18: google.protobuf.Empty
}
var file_rpc_cache_service_proto_depIdxs = []int32{
	8,  // 0: tunnel.cache.v1.ArtifactInfo.created:type_name -> google.protobuf.Timestamp
	9,  // 1: tunnel.cache.v1.ArtifactInfo.history_packages:type_name -> tunnel.common.Package
	0,  // 2: tunnel.cache.v1.PutArtifactRequest.artifact_info:type_name -> tunnel.cache.v1.ArtifactInfo
	10, // 3: tunnel.cache.v1.BlobInfo.os:type_name -> tunnel.common.OS
	11, // 4: tunnel.cache.v1.BlobInfo.repository:type_name -> tunnel.common.Repository
	12, // 5: tunnel.cache.v1.BlobInfo.package_infos:type_name -> tunnel.common.PackageInfo
	13, // 6: tunnel.cache.v1.BlobInfo.applications:type_name -> tunnel.common.Application
	14, // 7: tunnel.cache.v1.BlobInfo.misconfigurations:type_name -> tunnel.common.Misconfiguration
	15, // 8: tunnel.cache.v1.BlobInfo.custom_resources:type_name -> tunnel.common.CustomResource
	16, // 9: tunnel.cache.v1.BlobInfo.secrets:type_name -> tunnel.common.Secret
	17, // 10: tunnel.cache.v1.BlobInfo.licenses:type_name -> tunnel.common.LicenseFile
	2,  // 11: tunnel.cache.v1.PutBlobRequest.blob_info:type_name -> tunnel.cache.v1.BlobInfo
	10, // 12: tunnel.cache.v1.PutResponse.os:type_name -> tunnel.common.OS
	1,  // 13: tunnel.cache.v1.Cache.PutArtifact:input_type -> tunnel.cache.v1.PutArtifactRequest
	3,  // 14: tunnel.cache.v1.Cache.PutBlob:input_type -> tunnel.cache.v1.PutBlobRequest
	5,  // 15: tunnel.cache.v1.Cache.MissingBlobs:input_type -> tunnel.cache.v1.MissingBlobsRequest
	7,  // 16: tunnel.cache.v1.Cache.DeleteBlobs:input_type -> tunnel.cache.v1.DeleteBlobsRequest
	18, // 17: tunnel.cache.v1.Cache.PutArtifact:output_type -> google.protobuf.Empty
	18, // 18: tunnel.cache.v1.Cache.PutBlob:output_type -> google.protobuf.Empty
	6,  // 19: tunnel.cache.v1.Cache.MissingBlobs:output_type -> tunnel.cache.v1.MissingBlobsResponse
	18, // 20: tunnel.cache.v1.Cache.DeleteBlobs:output_type -> google.protobuf.Empty
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_rpc_cache_service_proto_init() }
func file_rpc_cache_service_proto_init() {
	if File_rpc_cache_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_cache_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactInfo); i {
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
		file_rpc_cache_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutArtifactRequest); i {
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
		file_rpc_cache_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobInfo); i {
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
		file_rpc_cache_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutBlobRequest); i {
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
		file_rpc_cache_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
		file_rpc_cache_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissingBlobsRequest); i {
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
		file_rpc_cache_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissingBlobsResponse); i {
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
		file_rpc_cache_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlobsRequest); i {
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
			RawDescriptor: file_rpc_cache_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_cache_service_proto_goTypes,
		DependencyIndexes: file_rpc_cache_service_proto_depIdxs,
		MessageInfos:      file_rpc_cache_service_proto_msgTypes,
	}.Build()
	File_rpc_cache_service_proto = out.File
	file_rpc_cache_service_proto_rawDesc = nil
	file_rpc_cache_service_proto_goTypes = nil
	file_rpc_cache_service_proto_depIdxs = nil
}