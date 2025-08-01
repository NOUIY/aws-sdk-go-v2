// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a metric stream. Metric streams can automatically stream
// CloudWatch metrics to Amazon Web Services destinations, including Amazon S3, and
// to many third-party solutions.
//
// For more information, see [Using Metric Streams].
//
// To create a metric stream, you must be signed in to an account that has the
// iam:PassRole permission and either the CloudWatchFullAccess policy or the
// cloudwatch:PutMetricStream permission.
//
// When you create or update a metric stream, you choose one of the following:
//
//   - Stream metrics from all metric namespaces in the account.
//
//   - Stream metrics from all metric namespaces in the account, except for the
//     namespaces that you list in ExcludeFilters .
//
//   - Stream metrics from only the metric namespaces that you list in
//     IncludeFilters .
//
// By default, a metric stream always sends the MAX , MIN , SUM , and SAMPLECOUNT
// statistics for each metric that is streamed. You can use the
// StatisticsConfigurations parameter to have the metric stream send additional
// statistics in the stream. Streaming additional statistics incurs additional
// costs. For more information, see [Amazon CloudWatch Pricing].
//
// When you use PutMetricStream to create a new metric stream, the stream is
// created in the running state. If you use it to update an existing stream, the
// state of the stream is not changed.
//
// If you are using CloudWatch cross-account observability and you create a metric
// stream in a monitoring account, you can choose whether to include metrics from
// source accounts in the stream. For more information, see [CloudWatch cross-account observability].
//
// [Using Metric Streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metric-Streams.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
func (c *Client) PutMetricStream(ctx context.Context, params *PutMetricStreamInput, optFns ...func(*Options)) (*PutMetricStreamOutput, error) {
	if params == nil {
		params = &PutMetricStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricStream", params, optFns, c.addOperationPutMetricStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricStreamInput struct {

	// The ARN of the Amazon Kinesis Data Firehose delivery stream to use for this
	// metric stream. This Amazon Kinesis Data Firehose delivery stream must already
	// exist and must be in the same account as the metric stream.
	//
	// This member is required.
	FirehoseArn *string

	// If you are creating a new metric stream, this is the name for the new stream.
	// The name must be different than the names of other metric streams in this
	// account and Region.
	//
	// If you are updating a metric stream, specify the name of that stream here.
	//
	// Valid characters are A-Z, a-z, 0-9, "-" and "_".
	//
	// This member is required.
	Name *string

	// The output format for the stream. Valid values are json , opentelemetry1.0 , and
	// opentelemetry0.7 . For more information about metric stream output formats, see [Metric streams output formats]
	// .
	//
	// [Metric streams output formats]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html
	//
	// This member is required.
	OutputFormat types.MetricStreamOutputFormat

	// The ARN of an IAM role that this metric stream will use to access Amazon
	// Kinesis Data Firehose resources. This IAM role must already exist and must be in
	// the same account as the metric stream. This IAM role must include the following
	// permissions:
	//
	//   - firehose:PutRecord
	//
	//   - firehose:PutRecordBatch
	//
	// This member is required.
	RoleArn *string

	// If you specify this parameter, the stream sends metrics from all metric
	// namespaces except for the namespaces that you specify here.
	//
	// You cannot include ExcludeFilters and IncludeFilters in the same operation.
	ExcludeFilters []types.MetricStreamFilter

	// If you specify this parameter, the stream sends only the metrics from the
	// metric namespaces that you specify here.
	//
	// You cannot include IncludeFilters and ExcludeFilters in the same operation.
	IncludeFilters []types.MetricStreamFilter

	// If you are creating a metric stream in a monitoring account, specify true to
	// include metrics from source accounts in the metric stream.
	IncludeLinkedAccountsMetrics *bool

	// By default, a metric stream always sends the MAX , MIN , SUM , and SAMPLECOUNT
	// statistics for each metric that is streamed. You can use this parameter to have
	// the metric stream also send additional statistics in the stream. This array can
	// have up to 100 members.
	//
	// For each entry in this array, you specify one or more metrics and the list of
	// additional statistics to stream for those metrics. The additional statistics
	// that you can stream depend on the stream's OutputFormat . If the OutputFormat
	// is json , you can stream any additional statistic that is supported by
	// CloudWatch, listed in [CloudWatch statistics definitions]. If the OutputFormat is opentelemetry1.0 or
	// opentelemetry0.7 , you can stream percentile statistics such as p95, p99.9, and
	// so on.
	//
	// [CloudWatch statistics definitions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html
	StatisticsConfigurations []types.MetricStreamStatisticsConfiguration

	// A list of key-value pairs to associate with the metric stream. You can
	// associate as many as 50 tags with a metric stream.
	//
	// Tags can help you organize and categorize your resources. You can also use them
	// to scope user permissions by granting a user permission to access or change only
	// resources with certain tag values.
	//
	// You can use this parameter only when you are creating a new metric stream. If
	// you are using this operation to update an existing metric stream, any tags you
	// specify in this parameter are ignored. To change the tags of an existing metric
	// stream, use [TagResource]or [UntagResource].
	//
	// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html
	// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_UntagResource.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutMetricStreamOutput struct {

	// The ARN of the metric stream.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMetricStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutMetricStream"); err != nil {
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
	if err = addOpPutMetricStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetricStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutMetricStream",
	}
}
