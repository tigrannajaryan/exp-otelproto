// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: collector/trace/v1/trace_service_stream.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_collector_trace_v1_trace_service_stream_proto protoreflect.FileDescriptor

var file_collector_trace_v1_trace_service_stream_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x22, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x33, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0xae, 0x01, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x67, 0x72,
	0x61, 0x6e, 0x6e, 0x61, 0x6a, 0x61, 0x72, 0x79, 0x61, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x2d, 0x6f,
	0x74, 0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0xaa, 0x02, 0x1f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_collector_trace_v1_trace_service_stream_proto_goTypes = []interface{}{
	(*TraceExportRequest)(nil), // 0: experimental.collector.trace.v1.TraceExportRequest
	(*ExportResponse)(nil),     // 1: experimental.collector.trace.v1.ExportResponse
}
var file_collector_trace_v1_trace_service_stream_proto_depIdxs = []int32{
	0, // 0: experimental.collector.trace.v1.StreamExporter.ExportTraces:input_type -> experimental.collector.trace.v1.TraceExportRequest
	1, // 1: experimental.collector.trace.v1.StreamExporter.ExportTraces:output_type -> experimental.collector.trace.v1.ExportResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_collector_trace_v1_trace_service_stream_proto_init() }
func file_collector_trace_v1_trace_service_stream_proto_init() {
	if File_collector_trace_v1_trace_service_stream_proto != nil {
		return
	}
	file_collector_trace_v1_exchangee_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_collector_trace_v1_trace_service_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collector_trace_v1_trace_service_stream_proto_goTypes,
		DependencyIndexes: file_collector_trace_v1_trace_service_stream_proto_depIdxs,
	}.Build()
	File_collector_trace_v1_trace_service_stream_proto = out.File
	file_collector_trace_v1_trace_service_stream_proto_rawDesc = nil
	file_collector_trace_v1_trace_service_stream_proto_goTypes = nil
	file_collector_trace_v1_trace_service_stream_proto_depIdxs = nil
}