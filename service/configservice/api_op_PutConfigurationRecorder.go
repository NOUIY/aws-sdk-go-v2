// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the customer managed configuration recorder.
//
// You can use this operation to create a new customer managed configuration
// recorder or to update the roleARN and the recordingGroup for an existing
// customer managed configuration recorder.
//
// To start the customer managed configuration recorder and begin recording
// configuration changes for the resource types you specify, use the [StartConfigurationRecorder]operation.
//
// For more information, see [Working with the Configuration Recorder] in the Config Developer Guide.
//
// # One customer managed configuration recorder per account per Region
//
// You can create only one customer managed configuration recorder for each
// account for each Amazon Web Services Region.
//
// Default is to record all supported resource types, excluding the global IAM
// resource types
//
// If you have not specified values for the recordingGroup field, the default for
// the customer managed configuration recorder is to record all supported resource
// types, excluding the global IAM resource types: AWS::IAM::Group ,
// AWS::IAM::Policy , AWS::IAM::Role , and AWS::IAM::User .
//
// # Tags are added at creation and cannot be updated
//
// PutConfigurationRecorder is an idempotent API. Subsequent requests won’t create
// a duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated, even
// if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Working with the Configuration Recorder]: https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [StartConfigurationRecorder]: https://docs.aws.amazon.com/config/latest/APIReference/API_StartConfigurationRecorder.html
func (c *Client) PutConfigurationRecorder(ctx context.Context, params *PutConfigurationRecorderInput, optFns ...func(*Options)) (*PutConfigurationRecorderOutput, error) {
	if params == nil {
		params = &PutConfigurationRecorderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationRecorder", params, optFns, c.addOperationPutConfigurationRecorderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationRecorderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the PutConfigurationRecorder action.
type PutConfigurationRecorderInput struct {

	// An object for the configuration recorder. A configuration recorder records
	// configuration changes for the resource types in scope.
	//
	// This member is required.
	ConfigurationRecorder *types.ConfigurationRecorder

	// The tags for the customer managed configuration recorder. Each tag consists of
	// a key and an optional value, both of which you define.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutConfigurationRecorderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConfigurationRecorderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutConfigurationRecorder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConfigurationRecorder{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutConfigurationRecorder"); err != nil {
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
	if err = addOpPutConfigurationRecorderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationRecorder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationRecorder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutConfigurationRecorder",
	}
}
