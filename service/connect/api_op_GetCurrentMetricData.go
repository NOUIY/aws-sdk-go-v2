// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the real-time metric data from the specified Amazon Connect instance.
//
// For a description of each metric, see [Metrics definitions] in the Amazon Connect Administrator
// Guide.
//
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
func (c *Client) GetCurrentMetricData(ctx context.Context, params *GetCurrentMetricDataInput, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) {
	if params == nil {
		params = &GetCurrentMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCurrentMetricData", params, optFns, c.addOperationGetCurrentMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCurrentMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCurrentMetricDataInput struct {

	// The metrics to retrieve. Specify the name and unit for each metric. The
	// following metrics are available. For a description of all the metrics, see [Metrics definitions]in
	// the Amazon Connect Administrator Guide.
	//
	// AGENTS_AFTER_CONTACT_WORK Unit: COUNT
	//
	// Name in real-time metrics report: [ACW]
	//
	// AGENTS_AVAILABLE Unit: COUNT
	//
	// Name in real-time metrics report: [Available]
	//
	// AGENTS_ERROR Unit: COUNT
	//
	// Name in real-time metrics report: [Error]
	//
	// AGENTS_NON_PRODUCTIVE Unit: COUNT
	//
	// Name in real-time metrics report: [NPT (Non-Productive Time)]
	//
	// AGENTS_ON_CALL Unit: COUNT
	//
	// Name in real-time metrics report: [On contact]
	//
	// AGENTS_ON_CONTACT Unit: COUNT
	//
	// Name in real-time metrics report: [On contact]
	//
	// AGENTS_ONLINE Unit: COUNT
	//
	// Name in real-time metrics report: [Online]
	//
	// AGENTS_STAFFED Unit: COUNT
	//
	// Name in real-time metrics report: [Staffed]
	//
	// CONTACTS_IN_QUEUE Unit: COUNT
	//
	// Name in real-time metrics report: [In queue]
	//
	// CONTACTS_SCHEDULED Unit: COUNT
	//
	// Name in real-time metrics report: [Scheduled]
	//
	// OLDEST_CONTACT_AGE Unit: SECONDS
	//
	// When you use groupings, Unit says SECONDS and the Value is returned in SECONDS.
	//
	// When you do not use groupings, Unit says SECONDS but the Value is returned in
	// MILLISECONDS. For example, if you get a response like this:
	//
	// { "Metric": { "Name": "OLDEST_CONTACT_AGE", "Unit": "SECONDS" }, "Value":
	// 24113.0 }
	//
	// The actual OLDEST_CONTACT_AGE is 24 seconds.
	//
	// When the filter RoutingStepExpression is used, this metric is still calculated
	// from enqueue time. For example, if a contact that has been queued under for 10
	// seconds has expired and becomes active, then OLDEST_CONTACT_AGE for this queue
	// will be counted starting from 10, not 0.
	//
	// Name in real-time metrics report: [Oldest]
	//
	// SLOTS_ACTIVE Unit: COUNT
	//
	// Name in real-time metrics report: [Active]
	//
	// SLOTS_AVAILABLE Unit: COUNT
	//
	// Name in real-time metrics report: [Availability]
	//
	// [Availability]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#availability-real-time
	// [Staffed]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#staffed-real-time
	// [NPT (Non-Productive Time)]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#non-productive-time-real-time
	// [Error]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#error-real-time
	// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
	// [Scheduled]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#scheduled-real-time
	// [Oldest]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#oldest-real-time
	// [On contact]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#on-call-real-time
	// [Active]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#active-real-time
	// [ACW]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#aftercallwork-real-time
	// [Available]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#available-real-time
	// [In queue]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#in-queue-real-time
	// [Online]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html#online-real-time
	//
	// This member is required.
	CurrentMetrics []types.CurrentMetric

	// The filters to apply to returned metrics. You can filter up to the following
	// limits:
	//
	//   - Queues: 100
	//
	//   - Routing profiles: 100
	//
	//   - Channels: 3 (VOICE, CHAT, and TASK channels are supported.)
	//
	//   - RoutingStepExpressions: 50
	//
	// Metric data is retrieved only for the resources associated with the queues or
	// routing profiles, and by any channels included in the filter. (You cannot filter
	// by both queue AND routing profile.) You can include both resource IDs and
	// resource ARNs in the same request.
	//
	// When using the RoutingStepExpression filter, you need to pass exactly one
	// QueueId . The filter is also case sensitive so when using the
	// RoutingStepExpression filter, grouping by ROUTING_STEP_EXPRESSION is required.
	//
	// Currently tagging is only supported on the resources that are passed in the
	// filter.
	//
	// This member is required.
	Filters *types.Filters

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The grouping applied to the metrics returned. For example, when grouped by QUEUE
	// , the metrics returned apply to each queue rather than aggregated for all
	// queues.
	//
	//   - If you group by CHANNEL , you should include a Channels filter. VOICE, CHAT,
	//   and TASK channels are supported.
	//
	//   - If you group by ROUTING_PROFILE , you must include either a queue or routing
	//   profile filter. In addition, a routing profile filter is required for metrics
	//   CONTACTS_SCHEDULED , CONTACTS_IN_QUEUE , and OLDEST_CONTACT_AGE .
	//
	//   - If no Grouping is included in the request, a summary of metrics is returned.
	//
	//   - When using the RoutingStepExpression filter, group by
	//   ROUTING_STEP_EXPRESSION is required.
	Groupings []types.Grouping

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string

	// The way to sort the resulting response based on metrics. You can enter one sort
	// criteria. By default resources are sorted based on AGENTS_ONLINE , DESCENDING .
	// The metric collection is sorted based on the input metrics.
	//
	// Note the following:
	//
	//   - Sorting on SLOTS_ACTIVE and SLOTS_AVAILABLE is not supported.
	SortCriteria []types.CurrentMetricSortCriteria

	noSmithyDocumentSerde
}

type GetCurrentMetricDataOutput struct {

	// The total count of the result, regardless of the current page size.
	ApproximateTotalCount *int64

	// The time at which the metrics were retrieved and cached for pagination.
	DataSnapshotTime *time.Time

	// Information about the real-time metrics.
	MetricResults []types.CurrentMetricResult

	// If there are additional results, this is the token for the next set of results.
	//
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCurrentMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCurrentMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCurrentMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCurrentMetricData"); err != nil {
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
	if err = addOpGetCurrentMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCurrentMetricData(options.Region), middleware.Before); err != nil {
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

// GetCurrentMetricDataPaginatorOptions is the paginator options for
// GetCurrentMetricData
type GetCurrentMetricDataPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCurrentMetricDataPaginator is a paginator for GetCurrentMetricData
type GetCurrentMetricDataPaginator struct {
	options   GetCurrentMetricDataPaginatorOptions
	client    GetCurrentMetricDataAPIClient
	params    *GetCurrentMetricDataInput
	nextToken *string
	firstPage bool
}

// NewGetCurrentMetricDataPaginator returns a new GetCurrentMetricDataPaginator
func NewGetCurrentMetricDataPaginator(client GetCurrentMetricDataAPIClient, params *GetCurrentMetricDataInput, optFns ...func(*GetCurrentMetricDataPaginatorOptions)) *GetCurrentMetricDataPaginator {
	if params == nil {
		params = &GetCurrentMetricDataInput{}
	}

	options := GetCurrentMetricDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCurrentMetricDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCurrentMetricDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCurrentMetricData page.
func (p *GetCurrentMetricDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetCurrentMetricData(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetCurrentMetricDataAPIClient is a client that implements the
// GetCurrentMetricData operation.
type GetCurrentMetricDataAPIClient interface {
	GetCurrentMetricData(context.Context, *GetCurrentMetricDataInput, ...func(*Options)) (*GetCurrentMetricDataOutput, error)
}

var _ GetCurrentMetricDataAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetCurrentMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCurrentMetricData",
	}
}
