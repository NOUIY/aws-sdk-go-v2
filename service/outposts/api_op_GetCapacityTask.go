// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details of the specified capacity task.
func (c *Client) GetCapacityTask(ctx context.Context, params *GetCapacityTaskInput, optFns ...func(*Options)) (*GetCapacityTaskOutput, error) {
	if params == nil {
		params = &GetCapacityTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCapacityTask", params, optFns, c.addOperationGetCapacityTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCapacityTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCapacityTaskInput struct {

	// ID of the capacity task.
	//
	// This member is required.
	CapacityTaskId *string

	// ID or ARN of the Outpost associated with the specified capacity task.
	//
	// This member is required.
	OutpostIdentifier *string

	noSmithyDocumentSerde
}

type GetCapacityTaskOutput struct {

	// The ID of the Outpost asset. An Outpost asset can be a single server within an
	// Outposts rack or an Outposts server configuration.
	AssetId *string

	// ID of the capacity task.
	CapacityTaskId *string

	// Status of the capacity task.
	//
	// A capacity task can have one of the following statuses:
	//
	//   - REQUESTED - The capacity task was created and is awaiting the next step by
	//   Amazon Web Services Outposts.
	//
	//   - IN_PROGRESS - The capacity task is running and cannot be cancelled.
	//
	//   - FAILED - The capacity task could not be completed.
	//
	//   - COMPLETED - The capacity task has completed successfully.
	//
	//   - WAITING_FOR_EVACUATION - The capacity task requires capacity to run. You
	//   must stop the recommended EC2 running instances to free up capacity for the task
	//   to run.
	//
	//   - CANCELLATION_IN_PROGRESS - The capacity task has been cancelled and is in
	//   the process of cleaning up resources.
	//
	//   - CANCELLED - The capacity task is cancelled.
	CapacityTaskStatus types.CapacityTaskStatus

	// The date the capacity task ran successfully.
	CompletionDate *time.Time

	// The date the capacity task was created.
	CreationDate *time.Time

	// Performs a dry run to determine if you are above or below instance capacity.
	DryRun bool

	// Reason why the capacity task failed.
	Failed *types.CapacityTaskFailure

	// Instances that the user specified they cannot stop in order to free up the
	// capacity needed to run the capacity task.
	InstancesToExclude *types.InstancesToExclude

	// The date the capacity task was last modified.
	LastModifiedDate *time.Time

	// ID of the Amazon Web Services Outposts order associated with the specified
	// capacity task.
	OrderId *string

	// ID of the Outpost associated with the specified capacity task.
	OutpostId *string

	// List of instance pools requested in the capacity task.
	RequestedInstancePools []types.InstanceTypeCapacity

	// User-specified option in case an instance is blocking the capacity task from
	// running. Shows one of the following options:
	//
	//   - WAIT_FOR_EVACUATION - Checks every 10 minutes over 48 hours to determine if
	//   instances have stopped and capacity is available to complete the task.
	//
	//   - FAIL_TASK - The capacity task fails.
	TaskActionOnBlockingInstances types.TaskActionOnBlockingInstances

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCapacityTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCapacityTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCapacityTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCapacityTask"); err != nil {
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
	if err = addOpGetCapacityTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCapacityTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCapacityTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCapacityTask",
	}
}
