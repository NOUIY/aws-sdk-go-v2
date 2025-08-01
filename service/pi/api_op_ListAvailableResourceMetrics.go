// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve metrics of the specified types that can be queried for a specified DB
// instance.
func (c *Client) ListAvailableResourceMetrics(ctx context.Context, params *ListAvailableResourceMetricsInput, optFns ...func(*Options)) (*ListAvailableResourceMetricsOutput, error) {
	if params == nil {
		params = &ListAvailableResourceMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableResourceMetrics", params, optFns, c.addOperationListAvailableResourceMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableResourceMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableResourceMetricsInput struct {

	// An immutable identifier for a data source that is unique within an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. To
	// use an Amazon RDS DB instance as a data source, specify its DbiResourceId
	// value. For example, specify db-ABCDEFGHIJKLMNOPQRSTU1VWZ .
	//
	// This member is required.
	Identifier *string

	// The types of metrics to return in the response. Valid values in the array
	// include the following:
	//
	//   - os (OS counter metrics) - All engines
	//
	//   - db (DB load metrics) - All engines except for Amazon DocumentDB
	//
	//   - db.sql.stats (per-SQL metrics) - All engines except for Amazon DocumentDB
	//
	//   - db.sql_tokenized.stats (per-SQL digest metrics) - All engines except for
	//   Amazon DocumentDB
	//
	// This member is required.
	MetricTypes []string

	// The Amazon Web Services service for which Performance Insights returns metrics.
	//
	// This member is required.
	ServiceType types.ServiceType

	// The maximum number of items to return. If the MaxRecords value is less than the
	// number of existing items, the response includes a pagination token.
	MaxResults *int32

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	noSmithyDocumentSerde
}

type ListAvailableResourceMetricsOutput struct {

	// An array of metrics available to query. Each array element contains the full
	// name, description, and unit of the metric.
	Metrics []types.ResponseResourceMetric

	// A pagination token that indicates the response didn’t return all available
	// records because MaxRecords was specified in the previous request. To get the
	// remaining records, specify NextToken in a separate request with this value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAvailableResourceMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAvailableResourceMetrics"); err != nil {
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
	if err = addOpListAvailableResourceMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableResourceMetrics(options.Region), middleware.Before); err != nil {
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

// ListAvailableResourceMetricsPaginatorOptions is the paginator options for
// ListAvailableResourceMetrics
type ListAvailableResourceMetricsPaginatorOptions struct {
	// The maximum number of items to return. If the MaxRecords value is less than the
	// number of existing items, the response includes a pagination token.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAvailableResourceMetricsPaginator is a paginator for
// ListAvailableResourceMetrics
type ListAvailableResourceMetricsPaginator struct {
	options   ListAvailableResourceMetricsPaginatorOptions
	client    ListAvailableResourceMetricsAPIClient
	params    *ListAvailableResourceMetricsInput
	nextToken *string
	firstPage bool
}

// NewListAvailableResourceMetricsPaginator returns a new
// ListAvailableResourceMetricsPaginator
func NewListAvailableResourceMetricsPaginator(client ListAvailableResourceMetricsAPIClient, params *ListAvailableResourceMetricsInput, optFns ...func(*ListAvailableResourceMetricsPaginatorOptions)) *ListAvailableResourceMetricsPaginator {
	if params == nil {
		params = &ListAvailableResourceMetricsInput{}
	}

	options := ListAvailableResourceMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAvailableResourceMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAvailableResourceMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAvailableResourceMetrics page.
func (p *ListAvailableResourceMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAvailableResourceMetricsOutput, error) {
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
	result, err := p.client.ListAvailableResourceMetrics(ctx, &params, optFns...)
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

// ListAvailableResourceMetricsAPIClient is a client that implements the
// ListAvailableResourceMetrics operation.
type ListAvailableResourceMetricsAPIClient interface {
	ListAvailableResourceMetrics(context.Context, *ListAvailableResourceMetricsInput, ...func(*Options)) (*ListAvailableResourceMetricsOutput, error)
}

var _ ListAvailableResourceMetricsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAvailableResourceMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAvailableResourceMetrics",
	}
}
