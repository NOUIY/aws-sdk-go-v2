// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details about a batch inference job. For more information, see [Monitor batch inference jobs]
//
// [Monitor batch inference jobs]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-monitor
func (c *Client) GetModelInvocationJob(ctx context.Context, params *GetModelInvocationJobInput, optFns ...func(*Options)) (*GetModelInvocationJobOutput, error) {
	if params == nil {
		params = &GetModelInvocationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModelInvocationJob", params, optFns, c.addOperationGetModelInvocationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelInvocationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelInvocationJobInput struct {

	// The Amazon Resource Name (ARN) of the batch inference job.
	//
	// This member is required.
	JobIdentifier *string

	noSmithyDocumentSerde
}

type GetModelInvocationJobOutput struct {

	// Details about the location of the input to the batch inference job.
	//
	// This member is required.
	InputDataConfig types.ModelInvocationJobInputDataConfig

	// The Amazon Resource Name (ARN) of the batch inference job.
	//
	// This member is required.
	JobArn *string

	// The unique identifier of the foundation model used for model inference.
	//
	// This member is required.
	ModelId *string

	// Details about the location of the output of the batch inference job.
	//
	// This member is required.
	OutputDataConfig types.ModelInvocationJobOutputDataConfig

	// The Amazon Resource Name (ARN) of the service role with permissions to carry
	// out and manage batch inference. You can use the console to create a default
	// service role or follow the steps at [Create a service role for batch inference].
	//
	// [Create a service role for batch inference]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-iam-sr.html
	//
	// This member is required.
	RoleArn *string

	// The time at which the batch inference job was submitted.
	//
	// This member is required.
	SubmitTime *time.Time

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientRequestToken *string

	// The time at which the batch inference job ended.
	EndTime *time.Time

	// The time at which the batch inference job times or timed out.
	JobExpirationTime *time.Time

	// The name of the batch inference job.
	JobName *string

	// The time at which the batch inference job was last modified.
	LastModifiedTime *time.Time

	// If the batch inference job failed, this field contains a message describing why
	// the job failed.
	Message *string

	// The status of the batch inference job.
	//
	// The following statuses are possible:
	//
	//   - Submitted – This job has been submitted to a queue for validation.
	//
	//   - Validating – This job is being validated for the requirements described in [Format and upload your batch inference data]
	//   . The criteria include the following:
	//
	//   - Your IAM service role has access to the Amazon S3 buckets containing your
	//   files.
	//
	//   - Your files are .jsonl files and each individual record is a JSON object in
	//   the correct format. Note that validation doesn't check if the modelInput value
	//   matches the request body for the model.
	//
	//   - Your files fulfill the requirements for file size and number of records.
	//   For more information, see [Quotas for Amazon Bedrock].
	//
	//   - Scheduled – This job has been validated and is now in a queue. The job will
	//   automatically start when it reaches its turn.
	//
	//   - Expired – This job timed out because it was scheduled but didn't begin
	//   before the set timeout duration. Submit a new job request.
	//
	//   - InProgress – This job has begun. You can start viewing the results in the
	//   output S3 location.
	//
	//   - Completed – This job has successfully completed. View the output files in
	//   the output S3 location.
	//
	//   - PartiallyCompleted – This job has partially completed. Not all of your
	//   records could be processed in time. View the output files in the output S3
	//   location.
	//
	//   - Failed – This job has failed. Check the failure message for any further
	//   details. For further assistance, reach out to the [Amazon Web ServicesSupport Center].
	//
	//   - Stopped – This job was stopped by a user.
	//
	//   - Stopping – This job is being stopped by a user.
	//
	// [Format and upload your batch inference data]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-data.html
	// [Amazon Web ServicesSupport Center]: https://console.aws.amazon.com/support/home/
	// [Quotas for Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/quotas.html
	Status types.ModelInvocationJobStatus

	// The number of hours after which batch inference job was set to time out.
	TimeoutDurationInHours *int32

	// The configuration of the Virtual Private Cloud (VPC) for the data in the batch
	// inference job. For more information, see [Protect batch inference jobs using a VPC].
	//
	// [Protect batch inference jobs using a VPC]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-vpc
	VpcConfig *types.VpcConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetModelInvocationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetModelInvocationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModelInvocationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetModelInvocationJob"); err != nil {
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
	if err = addOpGetModelInvocationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelInvocationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModelInvocationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetModelInvocationJob",
	}
}
