// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: billing_service.proto

package qreeket

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_billing_service_proto protoreflect.FileDescriptor

var file_billing_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x02,
	0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x71, 0x72, 0x65,
	0x65, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x14, 0x2e, 0x71,
	0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x21, 0x2e, 0x71, 0x72, 0x65, 0x65, 0x6b, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x6c, 0x6c, 0x63, 0x2f, 0x71, 0x72, 0x65,
	0x65, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_billing_service_proto_goTypes = []interface{}{
	(*ChargeRequest)(nil),            // 0: qreeket.ChargeRequest
	(*wrapperspb.StringValue)(nil),   // 1: google.protobuf.StringValue
	(*durationpb.Duration)(nil),      // 2: google.protobuf.Duration
	(*ChargeResponse)(nil),           // 3: qreeket.ChargeResponse
	(*Transaction)(nil),              // 4: qreeket.Transaction
	(*ListTransactionsResponse)(nil), // 5: qreeket.ListTransactionsResponse
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
}
var file_billing_service_proto_depIdxs = []int32{
	0, // 0: qreeket.BillingService.charge:input_type -> qreeket.ChargeRequest
	1, // 1: qreeket.BillingService.get_transaction:input_type -> google.protobuf.StringValue
	2, // 2: qreeket.BillingService.list_transactions:input_type -> google.protobuf.Duration
	1, // 3: qreeket.BillingService.delete_transaction:input_type -> google.protobuf.StringValue
	3, // 4: qreeket.BillingService.charge:output_type -> qreeket.ChargeResponse
	4, // 5: qreeket.BillingService.get_transaction:output_type -> qreeket.Transaction
	5, // 6: qreeket.BillingService.list_transactions:output_type -> qreeket.ListTransactionsResponse
	6, // 7: qreeket.BillingService.delete_transaction:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_billing_service_proto_init() }
func file_billing_service_proto_init() {
	if File_billing_service_proto != nil {
		return
	}
	file_billing_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_billing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_billing_service_proto_goTypes,
		DependencyIndexes: file_billing_service_proto_depIdxs,
	}.Build()
	File_billing_service_proto = out.File
	file_billing_service_proto_rawDesc = nil
	file_billing_service_proto_goTypes = nil
	file_billing_service_proto_depIdxs = nil
}
