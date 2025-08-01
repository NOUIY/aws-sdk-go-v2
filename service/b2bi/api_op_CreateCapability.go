// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Instantiates a capability based on the specified parameters. A trading
// capability contains the information required to transform incoming EDI documents
// into JSON or XML outputs.
func (c *Client) CreateCapability(ctx context.Context, params *CreateCapabilityInput, optFns ...func(*Options)) (*CreateCapabilityOutput, error) {
	if params == nil {
		params = &CreateCapabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCapability", params, optFns, c.addOperationCreateCapabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCapabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCapabilityInput struct {

	// Specifies a structure that contains the details for a capability.
	//
	// This member is required.
	Configuration types.CapabilityConfiguration

	// Specifies the name of the capability, used to identify it.
	//
	// This member is required.
	Name *string

	// Specifies the type of the capability. Currently, only edi is supported.
	//
	// This member is required.
	Type types.CapabilityType

	// Reserved for future use.
	ClientToken *string

	// Specifies one or more locations in Amazon S3, each specifying an EDI document
	// that can be used with this capability. Each item contains the name of the bucket
	// and the key, to identify the document's location.
	InstructionsDocuments []types.S3Location

	// Specifies the key-value pairs assigned to ARNs that you can use to group and
	// search for resources by type. You can attach this metadata to resources
	// (capabilities, partnerships, and so on) for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCapabilityOutput struct {

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	CapabilityArn *string

	// Returns a system-assigned unique identifier for the capability.
	//
	// This member is required.
	CapabilityId *string

	// Returns a structure that contains the details for a capability.
	//
	// This member is required.
	Configuration types.CapabilityConfiguration

	// Returns a timestamp for creation date and time of the capability.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns the name of the capability used to identify it.
	//
	// This member is required.
	Name *string

	// Returns the type of the capability. Currently, only edi is supported.
	//
	// This member is required.
	Type types.CapabilityType

	// Returns one or more locations in Amazon S3, each specifying an EDI document
	// that can be used with this capability. Each item contains the name of the bucket
	// and the key, to identify the document's location.
	InstructionsDocuments []types.S3Location

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCapabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateCapability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateCapability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCapability"); err != nil {
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
	if err = addIdempotencyToken_opCreateCapabilityMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCapabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCapability(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCapability struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCapability) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCapability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCapabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCapabilityInput ")
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
func addIdempotencyToken_opCreateCapabilityMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCapability{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCapability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCapability",
	}
}
