// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of dataset import jobs created using the [CreateDatasetImportJob] operation. For each
// import job, this operation returns a summary of its properties, including its
// Amazon Resource Name (ARN). You can retrieve the complete set of properties by
// using the ARN with the [DescribeDatasetImportJob]operation. You can filter the list by providing an array
// of [Filter]objects.
//
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetImportJob.html
// [Filter]: https://docs.aws.amazon.com/forecast/latest/dg/API_Filter.html
func (c *Client) ListDatasetImportJobs(ctx context.Context, params *ListDatasetImportJobsInput, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) {
	if params == nil {
		params = &ListDatasetImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetImportJobs", params, optFns, c.addOperationListDatasetImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetImportJobsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the datasets that match the statement from the list,
	// respectively. The match statement consists of a key and a value.
	//
	// Filter properties
	//
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT . To
	//   include the datasets that match the statement, specify IS . To exclude
	//   matching datasets, specify IS_NOT .
	//
	//   - Key - The name of the parameter to filter on. Valid values are DatasetArn
	//   and Status .
	//
	//   - Value - The value to match.
	//
	// For example, to list all dataset import jobs whose status is ACTIVE, you
	// specify the following filter:
	//
	//     "Filters": [ { "Condition": "IS", "Key": "Status", "Value": "ACTIVE" } ]
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken . To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatasetImportJobsOutput struct {

	// An array of objects that summarize each dataset import job's properties.
	DatasetImportJobs []types.DatasetImportJobSummary

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDatasetImportJobs"); err != nil {
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
	if err = addOpListDatasetImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetImportJobs(options.Region), middleware.Before); err != nil {
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

// ListDatasetImportJobsPaginatorOptions is the paginator options for
// ListDatasetImportJobs
type ListDatasetImportJobsPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasetImportJobsPaginator is a paginator for ListDatasetImportJobs
type ListDatasetImportJobsPaginator struct {
	options   ListDatasetImportJobsPaginatorOptions
	client    ListDatasetImportJobsAPIClient
	params    *ListDatasetImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListDatasetImportJobsPaginator returns a new ListDatasetImportJobsPaginator
func NewListDatasetImportJobsPaginator(client ListDatasetImportJobsAPIClient, params *ListDatasetImportJobsInput, optFns ...func(*ListDatasetImportJobsPaginatorOptions)) *ListDatasetImportJobsPaginator {
	if params == nil {
		params = &ListDatasetImportJobsInput{}
	}

	options := ListDatasetImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasetImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasetImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatasetImportJobs page.
func (p *ListDatasetImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) {
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
	result, err := p.client.ListDatasetImportJobs(ctx, &params, optFns...)
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

// ListDatasetImportJobsAPIClient is a client that implements the
// ListDatasetImportJobs operation.
type ListDatasetImportJobsAPIClient interface {
	ListDatasetImportJobs(context.Context, *ListDatasetImportJobsInput, ...func(*Options)) (*ListDatasetImportJobsOutput, error)
}

var _ ListDatasetImportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDatasetImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDatasetImportJobs",
	}
}
