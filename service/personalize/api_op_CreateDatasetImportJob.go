// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a job that imports training data from your data source (an Amazon S3
// bucket) to an Amazon Personalize dataset. To allow Amazon Personalize to import
// the training data, you must specify an IAM service role that has permission to
// read from the data source, as Amazon Personalize makes a copy of your data and
// processes it internally. For information on granting access to your Amazon S3
// bucket, see [Giving Amazon Personalize Access to Amazon S3 Resources].
//
// If you already created a recommender or deployed a custom solution version with
// a campaign, how new bulk records influence recommendations depends on the domain
// use case or recipe that you use. For more information, see [How new data influences real-time recommendations].
//
// By default, a dataset import job replaces any existing data in the dataset that
// you imported in bulk. To add new records without replacing existing data,
// specify INCREMENTAL for the import mode in the CreateDatasetImportJob operation.
//
// # Status
//
// A dataset import job can be in one of the following states:
//
//   - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the import job, call [DescribeDatasetImportJob], providing the Amazon Resource Name
// (ARN) of the dataset import job. The dataset import is complete when the status
// shows as ACTIVE. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed.
//
// Importing takes time. You must wait until the status shows as ACTIVE before
// training a model using the dataset.
//
// # Related APIs
//
// [ListDatasetImportJobs]
//
// [DescribeDatasetImportJob]
//
// [ListDatasetImportJobs]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasetImportJobs.html
// [Giving Amazon Personalize Access to Amazon S3 Resources]: https://docs.aws.amazon.com/personalize/latest/dg/granting-personalize-s3-access.html
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetImportJob.html
// [How new data influences real-time recommendations]: https://docs.aws.amazon.com/personalize/latest/dg/how-new-data-influences-recommendations.html
func (c *Client) CreateDatasetImportJob(ctx context.Context, params *CreateDatasetImportJobInput, optFns ...func(*Options)) (*CreateDatasetImportJobOutput, error) {
	if params == nil {
		params = &CreateDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatasetImportJob", params, optFns, c.addOperationCreateDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetImportJobInput struct {

	// The Amazon S3 bucket that contains the training data to import.
	//
	// This member is required.
	DataSource *types.DataSource

	// The ARN of the dataset that receives the imported data.
	//
	// This member is required.
	DatasetArn *string

	// The name for the dataset import job.
	//
	// This member is required.
	JobName *string

	// The ARN of the IAM role that has permissions to read from the Amazon S3 data
	// source.
	//
	// This member is required.
	RoleArn *string

	// Specify how to add the new records to an existing dataset. The default import
	// mode is FULL . If you haven't imported bulk records into the dataset previously,
	// you can only specify FULL .
	//
	//   - Specify FULL to overwrite all existing bulk data in your dataset. Data you
	//   imported individually is not replaced.
	//
	//   - Specify INCREMENTAL to append the new records to the existing data in your
	//   dataset. Amazon Personalize replaces any record with the same ID with the new
	//   one.
	ImportMode types.ImportMode

	// If you created a metric attribution, specify whether to publish metrics for
	// this import job to Amazon S3
	PublishAttributionMetricsToS3 *bool

	// A list of [tags] to apply to the dataset import job.
	//
	// [tags]: https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDatasetImportJobOutput struct {

	// The ARN of the dataset import job.
	DatasetImportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDatasetImportJob"); err != nil {
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
	if err = addOpCreateDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDatasetImportJob",
	}
}
