// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: commands.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_commands_proto protoreflect.FileDescriptor

var file_commands_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6c, 0x6f, 0x67, 0x73, 0x5f,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb2, 0x05, 0x0a,
	0x09, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x4c, 0x49, 0x12, 0x43, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x73, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c,
	0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_commands_proto_goTypes = []interface{}{
	(*EventsResendRequest)(nil),   // 0: rpc.EventsResendRequest
	(*ListenRequest)(nil),         // 1: rpc.ListenRequest
	(*LoginRequest)(nil),          // 2: rpc.LoginRequest
	(*LoginStatusRequest)(nil),    // 3: rpc.LoginStatusRequest
	(*LogsTailRequest)(nil),       // 4: rpc.LogsTailRequest
	(*SampleConfigsRequest)(nil),  // 5: rpc.SampleConfigsRequest
	(*SampleCreateRequest)(nil),   // 6: rpc.SampleCreateRequest
	(*SamplesListRequest)(nil),    // 7: rpc.SamplesListRequest
	(*TriggerRequest)(nil),        // 8: rpc.TriggerRequest
	(*TriggersListRequest)(nil),   // 9: rpc.TriggersListRequest
	(*VersionRequest)(nil),        // 10: rpc.VersionRequest
	(*EventsResendResponse)(nil),  // 11: rpc.EventsResendResponse
	(*ListenResponse)(nil),        // 12: rpc.ListenResponse
	(*LoginResponse)(nil),         // 13: rpc.LoginResponse
	(*LoginStatusResponse)(nil),   // 14: rpc.LoginStatusResponse
	(*LogsTailResponse)(nil),      // 15: rpc.LogsTailResponse
	(*SampleConfigsResponse)(nil), // 16: rpc.SampleConfigsResponse
	(*SampleCreateResponse)(nil),  // 17: rpc.SampleCreateResponse
	(*SamplesListResponse)(nil),   // 18: rpc.SamplesListResponse
	(*TriggerResponse)(nil),       // 19: rpc.TriggerResponse
	(*TriggersListResponse)(nil),  // 20: rpc.TriggersListResponse
	(*VersionResponse)(nil),       // 21: rpc.VersionResponse
}
var file_commands_proto_depIdxs = []int32{
	0,  // 0: rpc.StripeCLI.EventsResend:input_type -> rpc.EventsResendRequest
	1,  // 1: rpc.StripeCLI.Listen:input_type -> rpc.ListenRequest
	2,  // 2: rpc.StripeCLI.Login:input_type -> rpc.LoginRequest
	3,  // 3: rpc.StripeCLI.LoginStatus:input_type -> rpc.LoginStatusRequest
	4,  // 4: rpc.StripeCLI.LogsTail:input_type -> rpc.LogsTailRequest
	5,  // 5: rpc.StripeCLI.SampleConfigs:input_type -> rpc.SampleConfigsRequest
	6,  // 6: rpc.StripeCLI.SampleCreate:input_type -> rpc.SampleCreateRequest
	7,  // 7: rpc.StripeCLI.SamplesList:input_type -> rpc.SamplesListRequest
	8,  // 8: rpc.StripeCLI.Trigger:input_type -> rpc.TriggerRequest
	9,  // 9: rpc.StripeCLI.TriggersList:input_type -> rpc.TriggersListRequest
	10, // 10: rpc.StripeCLI.Version:input_type -> rpc.VersionRequest
	11, // 11: rpc.StripeCLI.EventsResend:output_type -> rpc.EventsResendResponse
	12, // 12: rpc.StripeCLI.Listen:output_type -> rpc.ListenResponse
	13, // 13: rpc.StripeCLI.Login:output_type -> rpc.LoginResponse
	14, // 14: rpc.StripeCLI.LoginStatus:output_type -> rpc.LoginStatusResponse
	15, // 15: rpc.StripeCLI.LogsTail:output_type -> rpc.LogsTailResponse
	16, // 16: rpc.StripeCLI.SampleConfigs:output_type -> rpc.SampleConfigsResponse
	17, // 17: rpc.StripeCLI.SampleCreate:output_type -> rpc.SampleCreateResponse
	18, // 18: rpc.StripeCLI.SamplesList:output_type -> rpc.SamplesListResponse
	19, // 19: rpc.StripeCLI.Trigger:output_type -> rpc.TriggerResponse
	20, // 20: rpc.StripeCLI.TriggersList:output_type -> rpc.TriggersListResponse
	21, // 21: rpc.StripeCLI.Version:output_type -> rpc.VersionResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_commands_proto_init() }
func file_commands_proto_init() {
	if File_commands_proto != nil {
		return
	}
	file_events_resend_proto_init()
	file_listen_proto_init()
	file_login_proto_init()
	file_login_status_proto_init()
	file_logs_tail_proto_init()
	file_sample_configs_proto_init()
	file_sample_create_proto_init()
	file_samples_list_proto_init()
	file_trigger_proto_init()
	file_triggers_list_proto_init()
	file_version_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commands_proto_goTypes,
		DependencyIndexes: file_commands_proto_depIdxs,
	}.Build()
	File_commands_proto = out.File
	file_commands_proto_rawDesc = nil
	file_commands_proto_goTypes = nil
	file_commands_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StripeCLIClient is the client API for StripeCLI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StripeCLIClient interface {
	// Resend an event given an event ID. Like `stripe events resend`.
	EventsResend(ctx context.Context, in *EventsResendRequest, opts ...grpc.CallOption) (*EventsResendResponse, error)
	// Receive webhook events from the Stripe API to your local machine. Like `stripe listen`.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (StripeCLI_ListenClient, error)
	// Get a link to log in to the Stripe CLI. The client will have to open the browser to complete
	// the login. Use `LoginStatus` after this method to wait for success. Like `stripe login`.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Successfully returns when login has succeeded, or returns an error if login has failed or timed
	// out. Use this method after `Login` to check for success.
	LoginStatus(ctx context.Context, in *LoginStatusRequest, opts ...grpc.CallOption) (*LoginStatusResponse, error)
	// Get a realtime stream of API logs. Like `stripe logs tail`.
	LogsTail(ctx context.Context, in *LogsTailRequest, opts ...grpc.CallOption) (StripeCLI_LogsTailClient, error)
	// Get a list of available configs for a given Stripe sample.
	SampleConfigs(ctx context.Context, in *SampleConfigsRequest, opts ...grpc.CallOption) (*SampleConfigsResponse, error)
	// Clone a Stripe sample. Like `stripe samples create`.
	SampleCreate(ctx context.Context, in *SampleCreateRequest, opts ...grpc.CallOption) (*SampleCreateResponse, error)
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error)
	// Trigger a webhook event. Like `stripe trigger`.
	Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// Get a list of supported events for `Trigger`.
	TriggersList(ctx context.Context, in *TriggersListRequest, opts ...grpc.CallOption) (*TriggersListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type stripeCLIClient struct {
	cc grpc.ClientConnInterface
}

func NewStripeCLIClient(cc grpc.ClientConnInterface) StripeCLIClient {
	return &stripeCLIClient{cc}
}

func (c *stripeCLIClient) EventsResend(ctx context.Context, in *EventsResendRequest, opts ...grpc.CallOption) (*EventsResendResponse, error) {
	out := new(EventsResendResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/EventsResend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (StripeCLI_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StripeCLI_serviceDesc.Streams[0], "/rpc.StripeCLI/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &stripeCLIListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StripeCLI_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type stripeCLIListenClient struct {
	grpc.ClientStream
}

func (x *stripeCLIListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stripeCLIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) LoginStatus(ctx context.Context, in *LoginStatusRequest, opts ...grpc.CallOption) (*LoginStatusResponse, error) {
	out := new(LoginStatusResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/LoginStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) LogsTail(ctx context.Context, in *LogsTailRequest, opts ...grpc.CallOption) (StripeCLI_LogsTailClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StripeCLI_serviceDesc.Streams[1], "/rpc.StripeCLI/LogsTail", opts...)
	if err != nil {
		return nil, err
	}
	x := &stripeCLILogsTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StripeCLI_LogsTailClient interface {
	Recv() (*LogsTailResponse, error)
	grpc.ClientStream
}

type stripeCLILogsTailClient struct {
	grpc.ClientStream
}

func (x *stripeCLILogsTailClient) Recv() (*LogsTailResponse, error) {
	m := new(LogsTailResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stripeCLIClient) SampleConfigs(ctx context.Context, in *SampleConfigsRequest, opts ...grpc.CallOption) (*SampleConfigsResponse, error) {
	out := new(SampleConfigsResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SampleConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) SampleCreate(ctx context.Context, in *SampleCreateRequest, opts ...grpc.CallOption) (*SampleCreateResponse, error) {
	out := new(SampleCreateResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SampleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error) {
	out := new(SamplesListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SamplesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) TriggersList(ctx context.Context, in *TriggersListRequest, opts ...grpc.CallOption) (*TriggersListResponse, error) {
	out := new(TriggersListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/TriggersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StripeCLIServer is the server API for StripeCLI service.
type StripeCLIServer interface {
	// Resend an event given an event ID. Like `stripe events resend`.
	EventsResend(context.Context, *EventsResendRequest) (*EventsResendResponse, error)
	// Receive webhook events from the Stripe API to your local machine. Like `stripe listen`.
	Listen(*ListenRequest, StripeCLI_ListenServer) error
	// Get a link to log in to the Stripe CLI. The client will have to open the browser to complete
	// the login. Use `LoginStatus` after this method to wait for success. Like `stripe login`.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Successfully returns when login has succeeded, or returns an error if login has failed or timed
	// out. Use this method after `Login` to check for success.
	LoginStatus(context.Context, *LoginStatusRequest) (*LoginStatusResponse, error)
	// Get a realtime stream of API logs. Like `stripe logs tail`.
	LogsTail(*LogsTailRequest, StripeCLI_LogsTailServer) error
	// Get a list of available configs for a given Stripe sample.
	SampleConfigs(context.Context, *SampleConfigsRequest) (*SampleConfigsResponse, error)
	// Clone a Stripe sample. Like `stripe samples create`.
	SampleCreate(context.Context, *SampleCreateRequest) (*SampleCreateResponse, error)
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error)
	// Trigger a webhook event. Like `stripe trigger`.
	Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error)
	// Get a list of supported events for `Trigger`.
	TriggersList(context.Context, *TriggersListRequest) (*TriggersListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedStripeCLIServer can be embedded to have forward compatible implementations.
type UnimplementedStripeCLIServer struct {
}

func (*UnimplementedStripeCLIServer) EventsResend(context.Context, *EventsResendRequest) (*EventsResendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventsResend not implemented")
}
func (*UnimplementedStripeCLIServer) Listen(*ListenRequest, StripeCLI_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (*UnimplementedStripeCLIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedStripeCLIServer) LoginStatus(context.Context, *LoginStatusRequest) (*LoginStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginStatus not implemented")
}
func (*UnimplementedStripeCLIServer) LogsTail(*LogsTailRequest, StripeCLI_LogsTailServer) error {
	return status.Errorf(codes.Unimplemented, "method LogsTail not implemented")
}
func (*UnimplementedStripeCLIServer) SampleConfigs(context.Context, *SampleConfigsRequest) (*SampleConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleConfigs not implemented")
}
func (*UnimplementedStripeCLIServer) SampleCreate(context.Context, *SampleCreateRequest) (*SampleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleCreate not implemented")
}
func (*UnimplementedStripeCLIServer) SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SamplesList not implemented")
}
func (*UnimplementedStripeCLIServer) Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}
func (*UnimplementedStripeCLIServer) TriggersList(context.Context, *TriggersListRequest) (*TriggersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggersList not implemented")
}
func (*UnimplementedStripeCLIServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterStripeCLIServer(s *grpc.Server, srv StripeCLIServer) {
	s.RegisterService(&_StripeCLI_serviceDesc, srv)
}

func _StripeCLI_EventsResend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsResendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).EventsResend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/EventsResend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).EventsResend(ctx, req.(*EventsResendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StripeCLIServer).Listen(m, &stripeCLIListenServer{stream})
}

type StripeCLI_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type stripeCLIListenServer struct {
	grpc.ServerStream
}

func (x *stripeCLIListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StripeCLI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_LoginStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).LoginStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/LoginStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).LoginStatus(ctx, req.(*LoginStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_LogsTail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsTailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StripeCLIServer).LogsTail(m, &stripeCLILogsTailServer{stream})
}

type StripeCLI_LogsTailServer interface {
	Send(*LogsTailResponse) error
	grpc.ServerStream
}

type stripeCLILogsTailServer struct {
	grpc.ServerStream
}

func (x *stripeCLILogsTailServer) Send(m *LogsTailResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StripeCLI_SampleConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SampleConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SampleConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SampleConfigs(ctx, req.(*SampleConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_SampleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SampleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SampleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SampleCreate(ctx, req.(*SampleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_SamplesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SamplesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SamplesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SamplesList(ctx, req.(*SamplesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Trigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_TriggersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).TriggersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/TriggersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).TriggersList(ctx, req.(*TriggersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StripeCLI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.StripeCLI",
	HandlerType: (*StripeCLIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventsResend",
			Handler:    _StripeCLI_EventsResend_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _StripeCLI_Login_Handler,
		},
		{
			MethodName: "LoginStatus",
			Handler:    _StripeCLI_LoginStatus_Handler,
		},
		{
			MethodName: "SampleConfigs",
			Handler:    _StripeCLI_SampleConfigs_Handler,
		},
		{
			MethodName: "SampleCreate",
			Handler:    _StripeCLI_SampleCreate_Handler,
		},
		{
			MethodName: "SamplesList",
			Handler:    _StripeCLI_SamplesList_Handler,
		},
		{
			MethodName: "Trigger",
			Handler:    _StripeCLI_Trigger_Handler,
		},
		{
			MethodName: "TriggersList",
			Handler:    _StripeCLI_TriggersList_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _StripeCLI_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _StripeCLI_Listen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogsTail",
			Handler:       _StripeCLI_LogsTail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "commands.proto",
}
