// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch template.
//
// A launch template contains the parameters to launch an instance. When you
// launch an instance using RunInstances, you can specify a launch template instead of
// providing the launch parameters in the request. For more information, see [Store instance launch parameters in Amazon EC2 launch templates]in
// the Amazon EC2 User Guide.
//
// To clone an existing launch template as the basis for a new launch template,
// use the Amazon EC2 console. The API, SDKs, and CLI do not support cloning a
// template. For more information, see [Create a launch template from an existing launch template]in the Amazon EC2 User Guide.
//
// [Create a launch template from an existing launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-launch-template.html#create-launch-template-from-existing-launch-template
// [Store instance launch parameters in Amazon EC2 launch templates]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
func (c *Client) CreateLaunchTemplate(ctx context.Context, params *CreateLaunchTemplateInput, optFns ...func(*Options)) (*CreateLaunchTemplateOutput, error) {
	if params == nil {
		params = &CreateLaunchTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchTemplate", params, optFns, c.addOperationCreateLaunchTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchTemplateInput struct {

	// The information for the launch template.
	//
	// This member is required.
	LaunchTemplateData *types.RequestLaunchTemplateData

	// A name for the launch template.
	//
	// This member is required.
	LaunchTemplateName *string

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request. If a client token isn't specified, a randomly generated token is used
	// in the request to ensure idempotency.
	//
	// For more information, see [Ensuring idempotency].
	//
	// Constraint: Maximum 128 ASCII characters.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Reserved for internal use.
	Operator *types.OperatorRequest

	// The tags to apply to the launch template on creation. To tag the launch
	// template, the resource type must be launch-template .
	//
	// To specify the tags for the resources that are created when an instance is
	// launched, you must use the TagSpecifications parameter in the [launch template data] structure.
	//
	// [launch template data]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestLaunchTemplateData.html
	TagSpecifications []types.TagSpecification

	// A description for the first version of the launch template.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateLaunchTemplateOutput struct {

	// Information about the launch template.
	LaunchTemplate *types.LaunchTemplate

	// If the launch template contains parameters or parameter combinations that are
	// not valid, an error code and an error message are returned for each issue that's
	// found.
	Warning *types.ValidationWarning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateLaunchTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLaunchTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLaunchTemplate"); err != nil {
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
	if err = addIdempotencyToken_opCreateLaunchTemplateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLaunchTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchTemplate(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateLaunchTemplate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateLaunchTemplate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateLaunchTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateLaunchTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateLaunchTemplateInput ")
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
func addIdempotencyToken_opCreateLaunchTemplateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateLaunchTemplate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateLaunchTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLaunchTemplate",
	}
}
