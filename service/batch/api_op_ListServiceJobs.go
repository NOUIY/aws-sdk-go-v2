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

// Returns a list of service jobs for a specified job queue.
func (c *Client) ListServiceJobs(ctx context.Context, params *ListServiceJobsInput, optFns ...func(*Options)) (*ListServiceJobsOutput, error) {
	if params == nil {
		params = &ListServiceJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceJobs", params, optFns, c.addOperationListServiceJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceJobsInput struct {

	// The filters to apply to the service job list query. The filter names and values
	// can be:
	//
	//   - name: JOB_STATUS
	//
	// values: SUBMITTED | PENDING | RUNNABLE | STARTING | RUNNING | SUCCEEDED |
	//   FAILED | SCHEDULED
	//
	//   - name: JOB_NAME
	//
	// values: case-insensitive matches for the job name. If a filter value ends with
	//   an asterisk (*), it matches any job name that begins with the string before the
	//   '*'.
	Filters []types.KeyValuesPair

	// The name or ARN of the job queue with which to list service jobs.
	JobQueue *string

	// The job status with which to filter service jobs.
	JobStatus types.ServiceJobStatus

	// The maximum number of results returned by ListServiceJobs in paginated output.
	// When this parameter is used, ListServiceJobs only returns maxResults results in
	// a single page and a nextToken response element. The remaining results of the
	// initial request can be seen by sending another ListServiceJobs request with the
	// returned nextToken value. This value can be between 1 and 100. If this
	// parameter isn't used, then ListServiceJobs returns up to 100 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListServiceJobs request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	//
	// Treat this token as an opaque identifier that's only used to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceJobsOutput struct {

	// A list of service job summaries.
	//
	// This member is required.
	JobSummaryList []types.ServiceJobSummary

	// The nextToken value to include in a future ListServiceJobs request. When the
	// results of a ListServiceJobs request exceed maxResults , this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServiceJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServiceJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceJobs(options.Region), middleware.Before); err != nil {
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

// ListServiceJobsPaginatorOptions is the paginator options for ListServiceJobs
type ListServiceJobsPaginatorOptions struct {
	// The maximum number of results returned by ListServiceJobs in paginated output.
	// When this parameter is used, ListServiceJobs only returns maxResults results in
	// a single page and a nextToken response element. The remaining results of the
	// initial request can be seen by sending another ListServiceJobs request with the
	// returned nextToken value. This value can be between 1 and 100. If this
	// parameter isn't used, then ListServiceJobs returns up to 100 results and a
	// nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceJobsPaginator is a paginator for ListServiceJobs
type ListServiceJobsPaginator struct {
	options   ListServiceJobsPaginatorOptions
	client    ListServiceJobsAPIClient
	params    *ListServiceJobsInput
	nextToken *string
	firstPage bool
}

// NewListServiceJobsPaginator returns a new ListServiceJobsPaginator
func NewListServiceJobsPaginator(client ListServiceJobsAPIClient, params *ListServiceJobsInput, optFns ...func(*ListServiceJobsPaginatorOptions)) *ListServiceJobsPaginator {
	if params == nil {
		params = &ListServiceJobsInput{}
	}

	options := ListServiceJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceJobs page.
func (p *ListServiceJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceJobsOutput, error) {
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
	result, err := p.client.ListServiceJobs(ctx, &params, optFns...)
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

// ListServiceJobsAPIClient is a client that implements the ListServiceJobs
// operation.
type ListServiceJobsAPIClient interface {
	ListServiceJobs(context.Context, *ListServiceJobsInput, ...func(*Options)) (*ListServiceJobsOutput, error)
}

var _ ListServiceJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListServiceJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceJobs",
	}
}
