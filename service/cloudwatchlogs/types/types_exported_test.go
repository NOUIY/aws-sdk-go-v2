// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func ExampleGetLogObjectResponseStream_outputUsage() {
	var union types.GetLogObjectResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.GetLogObjectResponseStreamMemberFields:
		_ = v.Value // Value is types.FieldsData

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FieldsData

func ExampleIntegrationDetails_outputUsage() {
	var union types.IntegrationDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.IntegrationDetailsMemberOpenSearchIntegrationDetails:
		_ = v.Value // Value is types.OpenSearchIntegrationDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OpenSearchIntegrationDetails

func ExampleResourceConfig_outputUsage() {
	var union types.ResourceConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ResourceConfigMemberOpenSearchResourceConfig:
		_ = v.Value // Value is types.OpenSearchResourceConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OpenSearchResourceConfig

func ExampleStartLiveTailResponseStream_outputUsage() {
	var union types.StartLiveTailResponseStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.StartLiveTailResponseStreamMemberSessionStart:
		_ = v.Value // Value is types.LiveTailSessionStart

	case *types.StartLiveTailResponseStreamMemberSessionUpdate:
		_ = v.Value // Value is types.LiveTailSessionUpdate

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LiveTailSessionUpdate
var _ *types.LiveTailSessionStart
