// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies a prompt in your prompt library. Include both fields that you want to
// keep and fields that you want to replace. For more information, see [Prompt management in Amazon Bedrock]and [Edit prompts in your prompt library] in the
// Amazon Bedrock User Guide.
//
// [Edit prompts in your prompt library]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-manage.html#prompt-management-edit
// [Prompt management in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html
func (c *Client) UpdatePrompt(ctx context.Context, params *UpdatePromptInput, optFns ...func(*Options)) (*UpdatePromptOutput, error) {
	if params == nil {
		params = &UpdatePromptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrompt", params, optFns, c.addOperationUpdatePromptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePromptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePromptInput struct {

	// A name for the prompt.
	//
	// This member is required.
	Name *string

	// The unique identifier of the prompt.
	//
	// This member is required.
	PromptIdentifier *string

	// The Amazon Resource Name (ARN) of the KMS key to encrypt the prompt.
	CustomerEncryptionKeyArn *string

	// The name of the default variant for the prompt. This value must match the name
	// field in the relevant [PromptVariant]object.
	//
	// [PromptVariant]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PromptVariant.html
	DefaultVariant *string

	// A description for the prompt.
	Description *string

	// A list of objects, each containing details about a variant of the prompt.
	Variants []types.PromptVariant

	noSmithyDocumentSerde
}

type UpdatePromptOutput struct {

	// The Amazon Resource Name (ARN) of the prompt.
	//
	// This member is required.
	Arn *string

	// The time at which the prompt was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique identifier of the prompt.
	//
	// This member is required.
	Id *string

	// The name of the prompt.
	//
	// This member is required.
	Name *string

	// The time at which the prompt was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The version of the prompt. When you update a prompt, the version updated is the
	// DRAFT version.
	//
	// This member is required.
	Version *string

	// The Amazon Resource Name (ARN) of the KMS key to encrypt the prompt.
	CustomerEncryptionKeyArn *string

	// The name of the default variant for the prompt. This value must match the name
	// field in the relevant [PromptVariant]object.
	//
	// [PromptVariant]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PromptVariant.html
	DefaultVariant *string

	// The description of the prompt.
	Description *string

	// A list of objects, each containing details about a variant of the prompt.
	Variants []types.PromptVariant

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePromptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePrompt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePrompt{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePrompt"); err != nil {
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
	if err = addOpUpdatePromptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrompt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePrompt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePrompt",
	}
}
