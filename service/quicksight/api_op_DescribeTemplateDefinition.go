// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a detailed description of the definition of a template.
//
// If you do not need to know details about the content of a template, for
// instance if you are trying to check the status of a recently created or updated
// template, use the [DescribeTemplate]DescribeTemplate instead.
//
// [DescribeTemplate]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeTemplate.html
func (c *Client) DescribeTemplateDefinition(ctx context.Context, params *DescribeTemplateDefinitionInput, optFns ...func(*Options)) (*DescribeTemplateDefinitionOutput, error) {
	if params == nil {
		params = &DescribeTemplateDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTemplateDefinition", params, optFns, c.addOperationDescribeTemplateDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTemplateDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTemplateDefinitionInput struct {

	// The ID of the Amazon Web Services account that contains the template. You must
	// be using the Amazon Web Services account that the template is in.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the template that you're describing.
	//
	// This member is required.
	TemplateId *string

	// The alias of the template that you want to describe. If you name a specific
	// alias, you describe the version that the alias points to. You can specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter. The keyword $PUBLISHED doesn't apply to templates.
	AliasName *string

	// The version number of the template.
	VersionNumber *int64

	noSmithyDocumentSerde
}

type DescribeTemplateDefinitionOutput struct {

	// The definition of the template.
	//
	// A definition is the data model of all features in a Dashboard, Template, or
	// Analysis.
	Definition *types.TemplateVersionDefinition

	// Errors associated with the template version.
	Errors []types.TemplateError

	// The descriptive name of the template.
	Name *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// Status associated with the template.
	//
	//   - CREATION_IN_PROGRESS
	//
	//   - CREATION_SUCCESSFUL
	//
	//   - CREATION_FAILED
	//
	//   - UPDATE_IN_PROGRESS
	//
	//   - UPDATE_SUCCESSFUL
	//
	//   - UPDATE_FAILED
	//
	//   - DELETED
	ResourceStatus types.ResourceStatus

	// The HTTP status of the request.
	Status int32

	// The ID of the template described.
	TemplateId *string

	// The ARN of the theme of the template.
	ThemeArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTemplateDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTemplateDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTemplateDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTemplateDefinition"); err != nil {
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
	if err = addOpDescribeTemplateDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTemplateDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTemplateDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTemplateDefinition",
	}
}
