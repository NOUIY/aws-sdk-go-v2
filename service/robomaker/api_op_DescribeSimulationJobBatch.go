// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// End of support notice: On September 10, 2025, Amazon Web Services will
// discontinue support for Amazon Web Services RoboMaker. After September 10, 2025,
// you will no longer be able to access the Amazon Web Services RoboMaker console
// or Amazon Web Services RoboMaker resources. For more information on
// transitioning to Batch to help run containerized simulations, visit [https://aws.amazon.com/blogs/hpc/run-simulations-using-multiple-containers-in-a-single-aws-batch-job/].
//
// Describes a simulation job batch.
//
// [https://aws.amazon.com/blogs/hpc/run-simulations-using-multiple-containers-in-a-single-aws-batch-job/]: https://aws.amazon.com/blogs/hpc/run-simulations-using-multiple-containers-in-a-single-aws-batch-job/
func (c *Client) DescribeSimulationJobBatch(ctx context.Context, params *DescribeSimulationJobBatchInput, optFns ...func(*Options)) (*DescribeSimulationJobBatchOutput, error) {
	if params == nil {
		params = &DescribeSimulationJobBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSimulationJobBatch", params, optFns, c.addOperationDescribeSimulationJobBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSimulationJobBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSimulationJobBatchInput struct {

	// The id of the batch to describe.
	//
	// This member is required.
	Batch *string

	noSmithyDocumentSerde
}

type DescribeSimulationJobBatchOutput struct {

	// The Amazon Resource Name (ARN) of the batch.
	Arn *string

	// The batch policy.
	BatchPolicy *types.BatchPolicy

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// created.
	CreatedAt *time.Time

	// A list of created simulation job summaries.
	CreatedRequests []types.SimulationJobSummary

	// A list of failed create simulation job requests. The request failed to be
	// created into a simulation job. Failed requests do not have a simulation job ID.
	FailedRequests []types.FailedCreateSimulationJobRequest

	// The failure code of the simulation job batch.
	FailureCode types.SimulationJobBatchErrorCode

	// The reason the simulation job batch failed.
	FailureReason *string

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// last updated.
	LastUpdatedAt *time.Time

	// A list of pending simulation job requests. These requests have not yet been
	// created into simulation jobs.
	PendingRequests []types.SimulationJobRequest

	// The status of the batch.
	//
	// Pending The simulation job batch request is pending.
	//
	// InProgress The simulation job batch is in progress.
	//
	// Failed The simulation job batch failed. One or more simulation job requests
	// could not be completed due to an internal failure (like InternalServiceError ).
	// See failureCode and failureReason for more information.
	//
	// Completed The simulation batch job completed. A batch is complete when (1)
	// there are no pending simulation job requests in the batch and none of the failed
	// simulation job requests are due to InternalServiceError and (2) when all
	// created simulation jobs have reached a terminal state (for example, Completed
	// or Failed ).
	//
	// Canceled The simulation batch job was cancelled.
	//
	// Canceling The simulation batch job is being cancelled.
	//
	// Completing The simulation batch job is completing.
	//
	// TimingOut The simulation job batch is timing out.
	//
	// If a batch timing out, and there are pending requests that were failing due to
	// an internal failure (like InternalServiceError ), the batch status will be
	// Failed . If there are no such failing request, the batch status will be TimedOut
	// .
	//
	// TimedOut The simulation batch job timed out.
	Status types.SimulationJobBatchStatus

	// A map that contains tag keys and tag values that are attached to the simulation
	// job batch.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSimulationJobBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSimulationJobBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSimulationJobBatch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSimulationJobBatch"); err != nil {
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
	if err = addOpDescribeSimulationJobBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSimulationJobBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSimulationJobBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSimulationJobBatch",
	}
}
