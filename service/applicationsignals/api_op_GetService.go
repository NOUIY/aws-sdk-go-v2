// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationsignals

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a service discovered by Application Signals.
func (c *Client) GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) {
	if params == nil {
		params = &GetServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetService", params, optFns, c.addOperationGetServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceInput struct {

	// The end of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// Your requested start time will be rounded to the nearest hour.
	//
	// This member is required.
	EndTime *time.Time

	// Use this field to specify which service you want to retrieve information for.
	// You must specify at least the Type , Name , and Environment attributes.
	//
	// This is a string-to-string map. It can include the following fields.
	//
	//   - Type designates the type of object this is.
	//
	//   - ResourceType specifies the type of the resource. This field is used only
	//   when the value of the Type field is Resource or AWS::Resource .
	//
	//   - Name specifies the name of the object. This is used only if the value of the
	//   Type field is Service , RemoteService , or AWS::Service .
	//
	//   - Identifier identifies the resource objects of this resource. This is used
	//   only if the value of the Type field is Resource or AWS::Resource .
	//
	//   - Environment specifies the location where this object is hosted, or what it
	//   belongs to.
	//
	// This member is required.
	KeyAttributes map[string]string

	// The start of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// Your requested start time will be rounded to the nearest hour.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetServiceOutput struct {

	// The end time of the data included in the response. In a raw HTTP Query API, it
	// is formatted as be epoch time in seconds. For example: 1698778057 .
	//
	// This displays the time that Application Signals used for the request. It might
	// not match your request exactly, because it was rounded to the nearest hour.
	//
	// This member is required.
	EndTime *time.Time

	// A structure containing information about the service.
	//
	// This member is required.
	Service *types.Service

	// The start time of the data included in the response. In a raw HTTP Query API,
	// it is formatted as be epoch time in seconds. For example: 1698778057 .
	//
	// This displays the time that Application Signals used for the request. It might
	// not match your request exactly, because it was rounded to the nearest hour.
	//
	// This member is required.
	StartTime *time.Time

	// An array of string-to-string maps that each contain information about one log
	// group associated with this service. Each string-to-string map includes the
	// following fields:
	//
	//   - "Type": "AWS::Resource"
	//
	//   - "ResourceType": "AWS::Logs::LogGroup"
	//
	//   - "Identifier": "name-of-log-group"
	LogGroupReferences []map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetService"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpGetServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetService(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addInterceptAttempt(stack, options); err != nil {
		return err
	}
	if err = addInterceptExecution(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptTransmit(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeDeserialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterDeserialization(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetService",
	}
}
