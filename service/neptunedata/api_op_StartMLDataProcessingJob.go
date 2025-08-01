// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Neptune ML data processing job for processing the graph data
// exported from Neptune for training. See [The dataprocessing command]dataprocessing .
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelDataProcessingJob]IAM action in that cluster.
//
// [neptune-db:StartMLModelDataProcessingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeldataprocessingjob
// [The dataprocessing command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-dataprocessing.html
func (c *Client) StartMLDataProcessingJob(ctx context.Context, params *StartMLDataProcessingJobInput, optFns ...func(*Options)) (*StartMLDataProcessingJobOutput, error) {
	if params == nil {
		params = &StartMLDataProcessingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLDataProcessingJob", params, optFns, c.addOperationStartMLDataProcessingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLDataProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLDataProcessingJobInput struct {

	// The URI of the Amazon S3 location where you want SageMaker to download the data
	// needed to run the data processing job.
	//
	// This member is required.
	InputDataS3Location *string

	// The URI of the Amazon S3 location where you want SageMaker to save the results
	// of a data processing job.
	//
	// This member is required.
	ProcessedDataS3Location *string

	// A data specification file that describes how to load the exported graph data
	// for training. The file is automatically generated by the Neptune export toolkit.
	// The default is training-data-configuration.json .
	ConfigFileName *string

	// A unique identifier for the new job. The default is an autogenerated UUID.
	Id *string

	// One of the two model types that Neptune ML currently supports: heterogeneous
	// graph models ( heterogeneous ), and knowledge graph ( kge ). The default is
	// none. If not specified, Neptune ML chooses the model type automatically based on
	// the data.
	ModelType *string

	// The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to
	// perform tasks on your behalf. This must be listed in your DB cluster parameter
	// group or an error will occur.
	NeptuneIamRoleArn *string

	// The job ID of a completed data processing job run on an earlier version of the
	// data.
	PreviousDataProcessingJobId *string

	// The type of ML instance used during data processing. Its memory should be large
	// enough to hold the processed dataset. The default is the smallest ml.r5 type
	// whose memory is ten times larger than the size of the exported graph data on
	// disk.
	ProcessingInstanceType *string

	// The disk volume size of the processing instance. Both input data and processed
	// data are stored on disk, so the volume size must be large enough to hold both
	// data sets. The default is 0. If not specified or 0, Neptune ML chooses the
	// volume size automatically based on the data size.
	ProcessingInstanceVolumeSizeInGB *int32

	// Timeout in seconds for the data processing job. The default is 86,400 (1 day).
	ProcessingTimeOutInSeconds *int32

	// The Amazon Key Management Service (Amazon KMS) key that SageMaker uses to
	// encrypt the output of the processing job. The default is none.
	S3OutputEncryptionKMSKey *string

	// The ARN of an IAM role for SageMaker execution. This must be listed in your DB
	// cluster parameter group or an error will occur.
	SagemakerIamRoleArn *string

	// The VPC security group IDs. The default is None.
	SecurityGroupIds []string

	// The IDs of the subnets in the Neptune VPC. The default is None.
	Subnets []string

	// The Amazon Key Management Service (Amazon KMS) key that SageMaker uses to
	// encrypt data on the storage volume attached to the ML compute instances that run
	// the training job. The default is None.
	VolumeEncryptionKMSKey *string

	noSmithyDocumentSerde
}

type StartMLDataProcessingJobOutput struct {

	// The ARN of the data processing job.
	Arn *string

	// The time it took to create the new processing job, in milliseconds.
	CreationTimeInMillis *int64

	// The unique ID of the new data processing job.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLDataProcessingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMLDataProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMLDataProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMLDataProcessingJob"); err != nil {
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
	if err = addOpStartMLDataProcessingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLDataProcessingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLDataProcessingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMLDataProcessingJob",
	}
}
