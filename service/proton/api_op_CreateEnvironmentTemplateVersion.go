// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new major or minor version of an environment template. A major version
// of an environment template is a version that isn't backwards compatible. A minor
// version of an environment template is a version that's backwards compatible
// within its major version.
func (c *Client) CreateEnvironmentTemplateVersion(ctx context.Context, params *CreateEnvironmentTemplateVersionInput, optFns ...func(*Options)) (*CreateEnvironmentTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateEnvironmentTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentTemplateVersion", params, optFns, c.addOperationCreateEnvironmentTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentTemplateVersionInput struct {

	// An object that includes the template bundle S3 bucket path and name for the new
	// version of an template.
	//
	// This member is required.
	Source types.TemplateVersionSourceInput

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	// When included, if two identical requests are made with the same client token,
	// Proton returns the environment template version that the first request created.
	ClientToken *string

	// A description of the new version of an environment template.
	Description *string

	// To create a new minor version of the environment template, include major Version
	// .
	//
	// To create a new major and minor version of the environment template, exclude
	// major Version .
	MajorVersion *string

	// An optional list of metadata items that you can associate with the Proton
	// environment template version. A tag is a key-value pair.
	//
	// For more information, see [Proton resources and tagging] in the Proton User Guide.
	//
	// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEnvironmentTemplateVersionOutput struct {

	// The environment template detail data that's returned by Proton.
	//
	// This member is required.
	EnvironmentTemplateVersion *types.EnvironmentTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentTemplateVersion"); err != nil {
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
	if err = addIdempotencyToken_opCreateEnvironmentTemplateVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentTemplateVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEnvironmentTemplateVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEnvironmentTemplateVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEnvironmentTemplateVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEnvironmentTemplateVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEnvironmentTemplateVersionInput ")
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
func addIdempotencyToken_opCreateEnvironmentTemplateVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEnvironmentTemplateVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEnvironmentTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentTemplateVersion",
	}
}
