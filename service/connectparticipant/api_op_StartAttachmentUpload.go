// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a pre-signed Amazon S3 URL in response for uploading the file directly
// to S3.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func (c *Client) StartAttachmentUpload(ctx context.Context, params *StartAttachmentUploadInput, optFns ...func(*Options)) (*StartAttachmentUploadOutput, error) {
	if params == nil {
		params = &StartAttachmentUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAttachmentUpload", params, optFns, c.addOperationStartAttachmentUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAttachmentUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAttachmentUploadInput struct {

	// A case-sensitive name of the attachment being uploaded.
	//
	// This member is required.
	AttachmentName *string

	// The size of the attachment in bytes.
	//
	// This member is required.
	AttachmentSizeInBytes int64

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	//
	// This member is required.
	ClientToken *string

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	// Describes the MIME file type of the attachment. For a list of supported file
	// types, see [Feature specifications]in the Amazon Connect Administrator Guide.
	//
	// [Feature specifications]: https://docs.aws.amazon.com/connect/latest/adminguide/feature-limits.html
	//
	// This member is required.
	ContentType *string

	noSmithyDocumentSerde
}

type StartAttachmentUploadOutput struct {

	// A unique identifier for the attachment.
	AttachmentId *string

	// The headers to be provided while uploading the file to the URL.
	UploadMetadata *types.UploadMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAttachmentUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAttachmentUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAttachmentUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAttachmentUpload"); err != nil {
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
	if err = addIdempotencyToken_opStartAttachmentUploadMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartAttachmentUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAttachmentUpload(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartAttachmentUpload struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartAttachmentUpload) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartAttachmentUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartAttachmentUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartAttachmentUploadInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartAttachmentUploadMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartAttachmentUpload{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartAttachmentUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAttachmentUpload",
	}
}
