// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The details of a service job.
func (c *Client) DescribeServiceJob(ctx context.Context, params *DescribeServiceJobInput, optFns ...func(*Options)) (*DescribeServiceJobOutput, error) {
	if params == nil {
		params = &DescribeServiceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceJob", params, optFns, c.addOperationDescribeServiceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceJobInput struct {

	// The job ID for the service job to describe.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeServiceJobOutput struct {

	// The job ID for the service job.
	//
	// This member is required.
	JobId *string

	// The name of the service job.
	//
	// This member is required.
	JobName *string

	// The ARN of the job queue that the service job is associated with.
	//
	// This member is required.
	JobQueue *string

	// The type of service job. For SageMaker Training jobs, this value is
	// SAGEMAKER_TRAINING .
	//
	// This member is required.
	ServiceJobType types.ServiceJobType

	// The Unix timestamp (in milliseconds) for when the service job was started.
	//
	// This member is required.
	StartedAt *int64

	// The current status of the service job.
	//
	// This member is required.
	Status types.ServiceJobStatus

	// A list of job attempts associated with the service job.
	Attempts []types.ServiceJobAttemptDetail

	// The Unix timestamp (in milliseconds) for when the service job was created.
	CreatedAt *int64

	// Indicates whether the service job has been terminated.
	IsTerminated *bool

	// The Amazon Resource Name (ARN) of the service job.
	JobArn *string

	// The latest attempt associated with the service job.
	LatestAttempt *types.LatestServiceJobAttempt

	// The retry strategy to use for failed service jobs that are submitted with this
	// service job.
	RetryStrategy *types.ServiceJobRetryStrategy

	// The scheduling priority of the service job.
	SchedulingPriority *int32

	// The request, in JSON, for the service that the SubmitServiceJob operation is
	// queueing.
	ServiceRequestPayload *string

	// The share identifier for the service job. This is used for fair-share
	// scheduling.
	ShareIdentifier *string

	// A short, human-readable string to provide more details for the current status
	// of the service job.
	StatusReason *string

	// The Unix timestamp (in milliseconds) for when the service job stopped running.
	StoppedAt *int64

	// The tags that are associated with the service job. Each tag consists of a key
	// and an optional value. For more information, see [Tagging your Batch resources].
	//
	// [Tagging your Batch resources]: https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html
	Tags map[string]string

	// The timeout configuration for the service job.
	TimeoutConfig *types.ServiceJobTimeout

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServiceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeServiceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeServiceJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeServiceJob"); err != nil {
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
	if err = addOpDescribeServiceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeServiceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeServiceJob",
	}
}
