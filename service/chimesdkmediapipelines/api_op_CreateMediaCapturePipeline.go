// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a media pipeline.
func (c *Client) CreateMediaCapturePipeline(ctx context.Context, params *CreateMediaCapturePipelineInput, optFns ...func(*Options)) (*CreateMediaCapturePipelineOutput, error) {
	if params == nil {
		params = &CreateMediaCapturePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMediaCapturePipeline", params, optFns, c.addOperationCreateMediaCapturePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMediaCapturePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMediaCapturePipelineInput struct {

	// The ARN of the sink type.
	//
	// This member is required.
	SinkArn *string

	// Destination type to which the media artifacts are saved. You must use an S3
	// bucket.
	//
	// This member is required.
	SinkType types.MediaPipelineSinkType

	// ARN of the source from which the media artifacts are captured.
	//
	// This member is required.
	SourceArn *string

	// Source type from which the media artifacts are captured. A Chime SDK Meeting is
	// the only supported source.
	//
	// This member is required.
	SourceType types.MediaPipelineSourceType

	// The configuration for a specified media pipeline. SourceType must be
	// ChimeSdkMeeting .
	ChimeSdkMeetingConfiguration *types.ChimeSdkMeetingConfiguration

	// The unique identifier for the client request. The token makes the API request
	// idempotent. Use a unique token for each media pipeline request.
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the sink role to be used with AwsKmsKeyId in
	// SseAwsKeyManagementParams . Can only interact with S3Bucket sink type. The role
	// must belong to the caller’s account and be able to act on behalf of the caller
	// during the API call. All minimum policy permissions requirements for the caller
	// to perform sink-related actions are the same for SinkIamRoleArn .
	//
	// Additionally, the role must have permission to kms:GenerateDataKey using KMS
	// key supplied as AwsKmsKeyId in SseAwsKeyManagementParams . If media
	// concatenation will be required later, the role must also have permission to
	// kms:Decrypt for the same KMS key.
	SinkIamRoleArn *string

	// An object that contains server side encryption parameters to be used by media
	// capture pipeline. The parameters can also be used by media concatenation
	// pipeline taking media capture pipeline as a media source.
	SseAwsKeyManagementParams *types.SseAwsKeyManagementParams

	// The tag key-value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMediaCapturePipelineOutput struct {

	// A media pipeline object, the ID, source type, source ARN, sink type, and sink
	// ARN of a media pipeline object.
	MediaCapturePipeline *types.MediaCapturePipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMediaCapturePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMediaCapturePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMediaCapturePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMediaCapturePipeline"); err != nil {
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
	if err = addIdempotencyToken_opCreateMediaCapturePipelineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMediaCapturePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMediaCapturePipeline(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMediaCapturePipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMediaCapturePipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMediaCapturePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMediaCapturePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMediaCapturePipelineInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMediaCapturePipelineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMediaCapturePipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMediaCapturePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMediaCapturePipeline",
	}
}
