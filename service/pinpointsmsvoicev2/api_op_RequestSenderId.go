// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Request a new sender ID that doesn't require registration.
func (c *Client) RequestSenderId(ctx context.Context, params *RequestSenderIdInput, optFns ...func(*Options)) (*RequestSenderIdOutput, error) {
	if params == nil {
		params = &RequestSenderIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestSenderId", params, optFns, c.addOperationRequestSenderIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestSenderIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestSenderIdInput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The sender ID string to request.
	//
	// This member is required.
	SenderId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// By default this is set to false. When set to true the sender ID can't be
	// deleted.
	DeletionProtectionEnabled *bool

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	MessageTypes []types.MessageType

	// An array of tags (key and value pairs) to associate with the sender ID.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type RequestSenderIdOutput struct {

	// By default this is set to false. When set to true the sender ID can't be
	// deleted.
	//
	// This member is required.
	DeletionProtectionEnabled bool

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageTypes []types.MessageType

	// The monthly price, in US dollars, to lease the sender ID.
	//
	// This member is required.
	MonthlyLeasingPrice *string

	// True if the sender ID is registered.
	//
	// This member is required.
	Registered bool

	// The sender ID that was requested.
	//
	// This member is required.
	SenderId *string

	// The Amazon Resource Name (ARN) associated with the SenderId.
	//
	// This member is required.
	SenderIdArn *string

	// An array of tags (key and value pairs) to associate with the sender ID.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestSenderIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRequestSenderId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRequestSenderId{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RequestSenderId"); err != nil {
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
	if err = addIdempotencyToken_opRequestSenderIdMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRequestSenderIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestSenderId(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRequestSenderId struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRequestSenderId) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRequestSenderId) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RequestSenderIdInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RequestSenderIdInput ")
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
func addIdempotencyToken_opRequestSenderIdMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRequestSenderId{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRequestSenderId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RequestSenderId",
	}
}
