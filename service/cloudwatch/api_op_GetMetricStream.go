// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the metric stream that you specify.
func (c *Client) GetMetricStream(ctx context.Context, params *GetMetricStreamInput, optFns ...func(*Options)) (*GetMetricStreamOutput, error) {
	if params == nil {
		params = &GetMetricStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricStream", params, optFns, c.addOperationGetMetricStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricStreamInput struct {

	// The name of the metric stream to retrieve information about.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetMetricStreamOutput struct {

	// The ARN of the metric stream.
	Arn *string

	// The date that the metric stream was created.
	CreationDate *time.Time

	// If this array of metric namespaces is present, then these namespaces are the
	// only metric namespaces that are not streamed by this metric stream. In this
	// case, all other metric namespaces in the account are streamed by this metric
	// stream.
	ExcludeFilters []types.MetricStreamFilter

	// The ARN of the Amazon Kinesis Data Firehose delivery stream that is used by
	// this metric stream.
	FirehoseArn *string

	// If this array of metric namespaces is present, then these namespaces are the
	// only metric namespaces that are streamed by this metric stream.
	IncludeFilters []types.MetricStreamFilter

	// If this is true and this metric stream is in a monitoring account, then the
	// stream includes metrics from source accounts that the monitoring account is
	// linked to.
	IncludeLinkedAccountsMetrics *bool

	// The date of the most recent update to the metric stream's configuration.
	LastUpdateDate *time.Time

	// The name of the metric stream.
	Name *string

	// The output format for the stream. Valid values are json , opentelemetry1.0 , and
	// opentelemetry0.7 . For more information about metric stream output formats, see [Metric streams output formats]
	// .
	//
	// [Metric streams output formats]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html
	OutputFormat types.MetricStreamOutputFormat

	// The ARN of the IAM role that is used by this metric stream.
	RoleArn *string

	// The state of the metric stream. The possible values are running and stopped .
	State *string

	// Each entry in this array displays information about one or more metrics that
	// include additional statistics in the metric stream. For more information about
	// the additional statistics, see [CloudWatch statistics definitions].
	//
	// [CloudWatch statistics definitions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.html
	StatisticsConfigurations []types.MetricStreamStatisticsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMetricStream"); err != nil {
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
	if err = addOpGetMetricStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMetricStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMetricStream",
	}
}
