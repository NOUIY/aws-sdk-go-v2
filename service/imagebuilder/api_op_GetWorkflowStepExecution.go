// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the runtime information that was logged for a specific runtime instance of
// the workflow step.
func (c *Client) GetWorkflowStepExecution(ctx context.Context, params *GetWorkflowStepExecutionInput, optFns ...func(*Options)) (*GetWorkflowStepExecutionOutput, error) {
	if params == nil {
		params = &GetWorkflowStepExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowStepExecution", params, optFns, c.addOperationGetWorkflowStepExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowStepExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowStepExecutionInput struct {

	// Use the unique identifier for a specific runtime instance of the workflow step
	// to get runtime details for that step.
	//
	// This member is required.
	StepExecutionId *string

	noSmithyDocumentSerde
}

type GetWorkflowStepExecutionOutput struct {

	// The name of the action that the specified step performs.
	Action *string

	// Describes the specified workflow step.
	Description *string

	// The timestamp when the specified runtime instance of the workflow step finished.
	EndTime *string

	// The Amazon Resource Name (ARN) of the image resource build version that the
	// specified runtime instance of the workflow step creates.
	ImageBuildVersionArn *string

	// Input parameters that Image Builder provided for the specified runtime instance
	// of the workflow step.
	Inputs *string

	// The output message from the specified runtime instance of the workflow step, if
	// applicable.
	Message *string

	// The name of the specified runtime instance of the workflow step.
	Name *string

	// The action to perform if the workflow step fails.
	OnFailure *string

	// The file names that the specified runtime version of the workflow step created
	// as output.
	Outputs *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Reports on the rollback status of the specified runtime version of the workflow
	// step, if applicable.
	RollbackStatus types.WorkflowStepExecutionRollbackStatus

	// The timestamp when the specified runtime version of the workflow step started.
	StartTime *string

	// The current status for the specified runtime version of the workflow step.
	Status types.WorkflowStepExecutionStatus

	// The unique identifier for the runtime version of the workflow step that you
	// specified in the request.
	StepExecutionId *string

	// The maximum duration in seconds for this step to complete its action.
	TimeoutSeconds *int32

	// The Amazon Resource Name (ARN) of the build version for the Image Builder
	// workflow resource that defines this workflow step.
	WorkflowBuildVersionArn *string

	// The unique identifier that Image Builder assigned to keep track of runtime
	// details when it ran the workflow.
	WorkflowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowStepExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflowStepExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflowStepExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetWorkflowStepExecution"); err != nil {
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
	if err = addOpGetWorkflowStepExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowStepExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflowStepExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetWorkflowStepExecution",
	}
}
