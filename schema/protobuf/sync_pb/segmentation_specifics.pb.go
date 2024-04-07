// Copyright 2022 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for segmentations.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: segmentation_specifics.proto

package sync_pb

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

type SegmentationSpecifics_DeviceMetadata_PlatformType int32

const (
	SegmentationSpecifics_DeviceMetadata_PLATFORM_TYPE_UNSPECIFIED SegmentationSpecifics_DeviceMetadata_PlatformType = 0
	SegmentationSpecifics_DeviceMetadata_PLATFORM_WINDOWS          SegmentationSpecifics_DeviceMetadata_PlatformType = 1
	SegmentationSpecifics_DeviceMetadata_PLATFORM_MAC              SegmentationSpecifics_DeviceMetadata_PlatformType = 2
	SegmentationSpecifics_DeviceMetadata_PLATFORM_LINUX            SegmentationSpecifics_DeviceMetadata_PlatformType = 3
	SegmentationSpecifics_DeviceMetadata_PLATFORM_CHROMEOS_ASH     SegmentationSpecifics_DeviceMetadata_PlatformType = 4
	SegmentationSpecifics_DeviceMetadata_PLATFORM_ANDROID          SegmentationSpecifics_DeviceMetadata_PlatformType = 5
	SegmentationSpecifics_DeviceMetadata_PLATFORM_IOS              SegmentationSpecifics_DeviceMetadata_PlatformType = 6
	SegmentationSpecifics_DeviceMetadata_PLATFORM_CHROMEOS_LACROS  SegmentationSpecifics_DeviceMetadata_PlatformType = 7
)

// Enum value maps for SegmentationSpecifics_DeviceMetadata_PlatformType.
var (
	SegmentationSpecifics_DeviceMetadata_PlatformType_name = map[int32]string{
		0: "PLATFORM_TYPE_UNSPECIFIED",
		1: "PLATFORM_WINDOWS",
		2: "PLATFORM_MAC",
		3: "PLATFORM_LINUX",
		4: "PLATFORM_CHROMEOS_ASH",
		5: "PLATFORM_ANDROID",
		6: "PLATFORM_IOS",
		7: "PLATFORM_CHROMEOS_LACROS",
	}
	SegmentationSpecifics_DeviceMetadata_PlatformType_value = map[string]int32{
		"PLATFORM_TYPE_UNSPECIFIED": 0,
		"PLATFORM_WINDOWS":          1,
		"PLATFORM_MAC":              2,
		"PLATFORM_LINUX":            3,
		"PLATFORM_CHROMEOS_ASH":     4,
		"PLATFORM_ANDROID":          5,
		"PLATFORM_IOS":              6,
		"PLATFORM_CHROMEOS_LACROS":  7,
	}
)

func (x SegmentationSpecifics_DeviceMetadata_PlatformType) Enum() *SegmentationSpecifics_DeviceMetadata_PlatformType {
	p := new(SegmentationSpecifics_DeviceMetadata_PlatformType)
	*p = x
	return p
}

func (x SegmentationSpecifics_DeviceMetadata_PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SegmentationSpecifics_DeviceMetadata_PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_segmentation_specifics_proto_enumTypes[0].Descriptor()
}

func (SegmentationSpecifics_DeviceMetadata_PlatformType) Type() protoreflect.EnumType {
	return &file_segmentation_specifics_proto_enumTypes[0]
}

func (x SegmentationSpecifics_DeviceMetadata_PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SegmentationSpecifics_DeviceMetadata_PlatformType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SegmentationSpecifics_DeviceMetadata_PlatformType(num)
	return nil
}

// Deprecated: Use SegmentationSpecifics_DeviceMetadata_PlatformType.Descriptor instead.
func (SegmentationSpecifics_DeviceMetadata_PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Sync data proto for storing segmentation data. Keyed by the combination of
// cache_guid (a Sync client id) and segmentation key.
//
// The segmentation platform is a platform that uses intelligence and machine
// learning to guide developers for building purpose-built user experience for
// specific segments of users. See go/chrome-segmentation for more details.
type SegmentationSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key used to identify the type of segmentation associated with a
	// feature, e.g. 'user_engagement'.
	SegmentationKey        *string                                       `protobuf:"bytes,1,opt,name=segmentation_key,json=segmentationKey" json:"segmentation_key,omitempty"`
	SegmentSelectionResult *SegmentationSpecifics_SegmentSelectionResult `protobuf:"bytes,2,opt,name=segment_selection_result,json=segmentSelectionResult" json:"segment_selection_result,omitempty"`
	DeviceMetadata         *SegmentationSpecifics_DeviceMetadata         `protobuf:"bytes,3,opt,name=device_metadata,json=deviceMetadata" json:"device_metadata,omitempty"`
	// One or more model execution data associated with each segment for the
	// segmentation key.
	ModelExecutionData []*SegmentationSpecifics_ModelExecutionData `protobuf:"bytes,4,rep,name=model_execution_data,json=modelExecutionData" json:"model_execution_data,omitempty"`
}

func (x *SegmentationSpecifics) Reset() {
	*x = SegmentationSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationSpecifics) ProtoMessage() {}

func (x *SegmentationSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationSpecifics.ProtoReflect.Descriptor instead.
func (*SegmentationSpecifics) Descriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *SegmentationSpecifics) GetSegmentationKey() string {
	if x != nil && x.SegmentationKey != nil {
		return *x.SegmentationKey
	}
	return ""
}

func (x *SegmentationSpecifics) GetSegmentSelectionResult() *SegmentationSpecifics_SegmentSelectionResult {
	if x != nil {
		return x.SegmentSelectionResult
	}
	return nil
}

func (x *SegmentationSpecifics) GetDeviceMetadata() *SegmentationSpecifics_DeviceMetadata {
	if x != nil {
		return x.DeviceMetadata
	}
	return nil
}

func (x *SegmentationSpecifics) GetModelExecutionData() []*SegmentationSpecifics_ModelExecutionData {
	if x != nil {
		return x.ModelExecutionData
	}
	return nil
}

// The selected segment by the segmentation scheme.
type SegmentationSpecifics_SegmentSelectionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string ID that refers to the segment the user was assigned to, e.g.,
	// 'highly_engaged_user'.
	SelectedSegment *string `protobuf:"bytes,1,opt,name=selected_segment,json=selectedSegment" json:"selected_segment,omitempty"`
	// Expiry time of selection result. Represents time from windows epoch in
	// seconds. Expired results are ignored by clients.
	ExpiryTimeWindowsEpochSeconds *int64 `protobuf:"varint,2,opt,name=expiry_time_windows_epoch_seconds,json=expiryTimeWindowsEpochSeconds" json:"expiry_time_windows_epoch_seconds,omitempty"`
	// Time when the segmentation data is updated. Used to weigh results by
	// recency. Represents time from windows epoch in seconds.
	LastUpdatedTimeWindowsEpochSeconds *int64 `protobuf:"varint,3,opt,name=last_updated_time_windows_epoch_seconds,json=lastUpdatedTimeWindowsEpochSeconds" json:"last_updated_time_windows_epoch_seconds,omitempty"`
}

func (x *SegmentationSpecifics_SegmentSelectionResult) Reset() {
	*x = SegmentationSpecifics_SegmentSelectionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationSpecifics_SegmentSelectionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationSpecifics_SegmentSelectionResult) ProtoMessage() {}

func (x *SegmentationSpecifics_SegmentSelectionResult) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationSpecifics_SegmentSelectionResult.ProtoReflect.Descriptor instead.
func (*SegmentationSpecifics_SegmentSelectionResult) Descriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SegmentationSpecifics_SegmentSelectionResult) GetSelectedSegment() string {
	if x != nil && x.SelectedSegment != nil {
		return *x.SelectedSegment
	}
	return ""
}

func (x *SegmentationSpecifics_SegmentSelectionResult) GetExpiryTimeWindowsEpochSeconds() int64 {
	if x != nil && x.ExpiryTimeWindowsEpochSeconds != nil {
		return *x.ExpiryTimeWindowsEpochSeconds
	}
	return 0
}

func (x *SegmentationSpecifics_SegmentSelectionResult) GetLastUpdatedTimeWindowsEpochSeconds() int64 {
	if x != nil && x.LastUpdatedTimeWindowsEpochSeconds != nil {
		return *x.LastUpdatedTimeWindowsEpochSeconds
	}
	return 0
}

// Metadata about the client device as used by the segmentation platform.
type SegmentationSpecifics_DeviceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cache_guid created to identify a sync client on this device.
	// Reuses the same Sync guid.
	CacheGuid *string `protobuf:"bytes,1,opt,name=cache_guid,json=cacheGuid" json:"cache_guid,omitempty"`
	// The OS platform of the device.
	PlatformType *SegmentationSpecifics_DeviceMetadata_PlatformType `protobuf:"varint,2,opt,name=platform_type,json=platformType,enum=sync_pb.SegmentationSpecifics_DeviceMetadata_PlatformType" json:"platform_type,omitempty"`
}

func (x *SegmentationSpecifics_DeviceMetadata) Reset() {
	*x = SegmentationSpecifics_DeviceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_specifics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationSpecifics_DeviceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationSpecifics_DeviceMetadata) ProtoMessage() {}

func (x *SegmentationSpecifics_DeviceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_specifics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationSpecifics_DeviceMetadata.ProtoReflect.Descriptor instead.
func (*SegmentationSpecifics_DeviceMetadata) Descriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SegmentationSpecifics_DeviceMetadata) GetCacheGuid() string {
	if x != nil && x.CacheGuid != nil {
		return *x.CacheGuid
	}
	return ""
}

func (x *SegmentationSpecifics_DeviceMetadata) GetPlatformType() SegmentationSpecifics_DeviceMetadata_PlatformType {
	if x != nil && x.PlatformType != nil {
		return *x.PlatformType
	}
	return SegmentationSpecifics_DeviceMetadata_PLATFORM_TYPE_UNSPECIFIED
}

// Model execution data including segment scores and related metadata, e.g.,
// model version.
type SegmentationSpecifics_ModelExecutionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string ID that identifies a model.
	ModelId *string `protobuf:"bytes,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	// A model may output one or multiple scores, one score for each segment
	// label.
	ModelOutputs []*SegmentationSpecifics_ModelExecutionData_ModelOutput `protobuf:"bytes,2,rep,name=model_outputs,json=modelOutputs" json:"model_outputs,omitempty"`
	// Timestamp when the ML model was executed.
	// Represents time from windows epoch in seconds.
	ExecutionTimeWindowsEpochSeconds *int64 `protobuf:"varint,3,opt,name=execution_time_windows_epoch_seconds,json=executionTimeWindowsEpochSeconds" json:"execution_time_windows_epoch_seconds,omitempty"`
	// Expiry timestamp for the model scores.
	// Represents time from windows epoch in seconds.
	ScoreExpiryTimeWindowsEpochSeconds *int64 `protobuf:"varint,4,opt,name=score_expiry_time_windows_epoch_seconds,json=scoreExpiryTimeWindowsEpochSeconds" json:"score_expiry_time_windows_epoch_seconds,omitempty"`
	// The version of the ML model that was run.
	ModelVersion *int32 `protobuf:"varint,5,opt,name=model_version,json=modelVersion" json:"model_version,omitempty"`
}

func (x *SegmentationSpecifics_ModelExecutionData) Reset() {
	*x = SegmentationSpecifics_ModelExecutionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_specifics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationSpecifics_ModelExecutionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationSpecifics_ModelExecutionData) ProtoMessage() {}

func (x *SegmentationSpecifics_ModelExecutionData) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_specifics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationSpecifics_ModelExecutionData.ProtoReflect.Descriptor instead.
func (*SegmentationSpecifics_ModelExecutionData) Descriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SegmentationSpecifics_ModelExecutionData) GetModelId() string {
	if x != nil && x.ModelId != nil {
		return *x.ModelId
	}
	return ""
}

func (x *SegmentationSpecifics_ModelExecutionData) GetModelOutputs() []*SegmentationSpecifics_ModelExecutionData_ModelOutput {
	if x != nil {
		return x.ModelOutputs
	}
	return nil
}

func (x *SegmentationSpecifics_ModelExecutionData) GetExecutionTimeWindowsEpochSeconds() int64 {
	if x != nil && x.ExecutionTimeWindowsEpochSeconds != nil {
		return *x.ExecutionTimeWindowsEpochSeconds
	}
	return 0
}

func (x *SegmentationSpecifics_ModelExecutionData) GetScoreExpiryTimeWindowsEpochSeconds() int64 {
	if x != nil && x.ScoreExpiryTimeWindowsEpochSeconds != nil {
		return *x.ScoreExpiryTimeWindowsEpochSeconds
	}
	return 0
}

func (x *SegmentationSpecifics_ModelExecutionData) GetModelVersion() int32 {
	if x != nil && x.ModelVersion != nil {
		return *x.ModelVersion
	}
	return 0
}

// Output of model.
type SegmentationSpecifics_ModelExecutionData_ModelOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When outputting multiple scores from a single model, this is the
	// segment label for each output.
	Label *string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// Raw segment scores provided by the ML model or the heuristic model.
	// The score is derived based on a combination of UMA histograms,
	// user actions or UKM and URLs visited. The score can be treated as a
	// probability of the user liking a feature, like feed, or NTP.
	Score *float32 `protobuf:"fixed32,2,opt,name=score" json:"score,omitempty"`
	// A rank defined by the segmentation scheme.
	Rank *int32 `protobuf:"varint,3,opt,name=rank" json:"rank,omitempty"`
}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) Reset() {
	*x = SegmentationSpecifics_ModelExecutionData_ModelOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segmentation_specifics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmentationSpecifics_ModelExecutionData_ModelOutput) ProtoMessage() {}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) ProtoReflect() protoreflect.Message {
	mi := &file_segmentation_specifics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmentationSpecifics_ModelExecutionData_ModelOutput.ProtoReflect.Descriptor instead.
func (*SegmentationSpecifics_ModelExecutionData_ModelOutput) Descriptor() ([]byte, []int) {
	return file_segmentation_specifics_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) GetScore() float32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *SegmentationSpecifics_ModelExecutionData_ModelOutput) GetRank() int32 {
	if x != nil && x.Rank != nil {
		return *x.Rank
	}
	return 0
}

var File_segmentation_specifics_proto protoreflect.FileDescriptor

var file_segmentation_specifics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xe4, 0x0a, 0x0a, 0x15, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x6f, 0x0a, 0x18,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x16, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a,
	0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x63, 0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0xe2, 0x01, 0x0a, 0x16, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x48, 0x0a, 0x21, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x53, 0x0a, 0x27, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x22, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x73, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x1a,
	0xdd, 0x02, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x12, 0x5f, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x03, 0x12, 0x19,
	0x0a, 0x15, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x48, 0x52, 0x4f, 0x4d,
	0x45, 0x4f, 0x53, 0x5f, 0x41, 0x53, 0x48, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41,
	0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x05, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4f, 0x53, 0x10,
	0x06, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x48,
	0x52, 0x4f, 0x4d, 0x45, 0x4f, 0x53, 0x5f, 0x4c, 0x41, 0x43, 0x52, 0x4f, 0x53, 0x10, 0x07, 0x1a,
	0xac, 0x03, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x62, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x4e, 0x0a, 0x24, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x53, 0x0a, 0x27, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x73, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x22, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0x4d, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x42, 0x36,
	0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_segmentation_specifics_proto_rawDescOnce sync.Once
	file_segmentation_specifics_proto_rawDescData = file_segmentation_specifics_proto_rawDesc
)

func file_segmentation_specifics_proto_rawDescGZIP() []byte {
	file_segmentation_specifics_proto_rawDescOnce.Do(func() {
		file_segmentation_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_segmentation_specifics_proto_rawDescData)
	})
	return file_segmentation_specifics_proto_rawDescData
}

var file_segmentation_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_segmentation_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_segmentation_specifics_proto_goTypes = []interface{}{
	(SegmentationSpecifics_DeviceMetadata_PlatformType)(0),       // 0: sync_pb.SegmentationSpecifics.DeviceMetadata.PlatformType
	(*SegmentationSpecifics)(nil),                                // 1: sync_pb.SegmentationSpecifics
	(*SegmentationSpecifics_SegmentSelectionResult)(nil),         // 2: sync_pb.SegmentationSpecifics.SegmentSelectionResult
	(*SegmentationSpecifics_DeviceMetadata)(nil),                 // 3: sync_pb.SegmentationSpecifics.DeviceMetadata
	(*SegmentationSpecifics_ModelExecutionData)(nil),             // 4: sync_pb.SegmentationSpecifics.ModelExecutionData
	(*SegmentationSpecifics_ModelExecutionData_ModelOutput)(nil), // 5: sync_pb.SegmentationSpecifics.ModelExecutionData.ModelOutput
}
var file_segmentation_specifics_proto_depIdxs = []int32{
	2, // 0: sync_pb.SegmentationSpecifics.segment_selection_result:type_name -> sync_pb.SegmentationSpecifics.SegmentSelectionResult
	3, // 1: sync_pb.SegmentationSpecifics.device_metadata:type_name -> sync_pb.SegmentationSpecifics.DeviceMetadata
	4, // 2: sync_pb.SegmentationSpecifics.model_execution_data:type_name -> sync_pb.SegmentationSpecifics.ModelExecutionData
	0, // 3: sync_pb.SegmentationSpecifics.DeviceMetadata.platform_type:type_name -> sync_pb.SegmentationSpecifics.DeviceMetadata.PlatformType
	5, // 4: sync_pb.SegmentationSpecifics.ModelExecutionData.model_outputs:type_name -> sync_pb.SegmentationSpecifics.ModelExecutionData.ModelOutput
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_segmentation_specifics_proto_init() }
func file_segmentation_specifics_proto_init() {
	if File_segmentation_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_segmentation_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationSpecifics); i {
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
		file_segmentation_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationSpecifics_SegmentSelectionResult); i {
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
		file_segmentation_specifics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationSpecifics_DeviceMetadata); i {
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
		file_segmentation_specifics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationSpecifics_ModelExecutionData); i {
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
		file_segmentation_specifics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmentationSpecifics_ModelExecutionData_ModelOutput); i {
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
			RawDescriptor: file_segmentation_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_segmentation_specifics_proto_goTypes,
		DependencyIndexes: file_segmentation_specifics_proto_depIdxs,
		EnumInfos:         file_segmentation_specifics_proto_enumTypes,
		MessageInfos:      file_segmentation_specifics_proto_msgTypes,
	}.Build()
	File_segmentation_specifics_proto = out.File
	file_segmentation_specifics_proto_rawDesc = nil
	file_segmentation_specifics_proto_goTypes = nil
	file_segmentation_specifics_proto_depIdxs = nil
}
