// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the definition for a model explainability job.
func (c *Client) CreateModelExplainabilityJobDefinition(ctx context.Context, params *CreateModelExplainabilityJobDefinitionInput, optFns ...func(*Options)) (*CreateModelExplainabilityJobDefinitionOutput, error) {
	if params == nil {
		params = &CreateModelExplainabilityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelExplainabilityJobDefinition", params, optFns, c.addOperationCreateModelExplainabilityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelExplainabilityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelExplainabilityJobDefinitionInput struct {

	//  The name of the model explainability job definition. The name must be unique
	// within an Amazon Web Services Region in the Amazon Web Services account.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// Configures the model explainability job to run a specified Docker container
	// image.
	//
	// This member is required.
	ModelExplainabilityAppSpecification *types.ModelExplainabilityAppSpecification

	// Inputs for the model explainability job.
	//
	// This member is required.
	ModelExplainabilityJobInput *types.ModelExplainabilityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	ModelExplainabilityJobOutputConfig *types.MonitoringOutputConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can
	// assume to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig *types.ModelExplainabilityBaselineConfig

	// Networking options for a model explainability job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// (Optional) An array of key-value pairs. For more information, see [Using Cost Allocation Tags] in the
	// Amazon Web Services Billing and Cost Management User Guide.
	//
	// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateModelExplainabilityJobDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the model explainability job.
	//
	// This member is required.
	JobDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelExplainabilityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelExplainabilityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelExplainabilityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModelExplainabilityJobDefinition"); err != nil {
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
	if err = addOpCreateModelExplainabilityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelExplainabilityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModelExplainabilityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModelExplainabilityJobDefinition",
	}
}
