// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a specific task running on a specific target.
func (c *Client) GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInvocationInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionTaskInvocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecutionTaskInvocation", params, optFns, c.addOperationGetMaintenanceWindowExecutionTaskInvocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionTaskInvocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionTaskInvocationInput struct {

	// The invocation ID to retrieve.
	//
	// This member is required.
	InvocationId *string

	// The ID of the specific task in the maintenance window task that should be
	// retrieved.
	//
	// This member is required.
	TaskId *string

	// The ID of the maintenance window execution for which the task is a part.
	//
	// This member is required.
	WindowExecutionId *string

	noSmithyDocumentSerde
}

type GetMaintenanceWindowExecutionTaskInvocationOutput struct {

	// The time that the task finished running on the target.
	EndTime *time.Time

	// The execution ID.
	ExecutionId *string

	// The invocation ID.
	InvocationId *string

	// User-provided value to be included in any Amazon CloudWatch Events or Amazon
	// EventBridge events raised while running tasks for these targets in this
	// maintenance window.
	OwnerInformation *string

	// The parameters used at the time that the task ran.
	Parameters *string

	// The time that the task started running on the target.
	StartTime *time.Time

	// The task status for an invocation.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the status. Details are only available for certain
	// status values.
	StatusDetails *string

	// The task execution ID.
	TaskExecutionId *string

	// Retrieves the task type for a maintenance window.
	TaskType types.MaintenanceWindowTaskType

	// The maintenance window execution ID.
	WindowExecutionId *string

	// The maintenance window target ID.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMaintenanceWindowExecutionTaskInvocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecutionTaskInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecutionTaskInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMaintenanceWindowExecutionTaskInvocation"); err != nil {
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
	if err = addOpGetMaintenanceWindowExecutionTaskInvocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTaskInvocation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTaskInvocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMaintenanceWindowExecutionTaskInvocation",
	}
}
