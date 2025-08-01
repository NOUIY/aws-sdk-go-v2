// Code generated by smithy-go-codegen DO NOT EDIT.

package backupsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backupsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation exports search results of a search job to a specified
// destination S3 bucket.
func (c *Client) ListSearchResultExportJobs(ctx context.Context, params *ListSearchResultExportJobsInput, optFns ...func(*Options)) (*ListSearchResultExportJobsOutput, error) {
	if params == nil {
		params = &ListSearchResultExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSearchResultExportJobs", params, optFns, c.addOperationListSearchResultExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSearchResultExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSearchResultExportJobsInput struct {

	// The maximum number of resource list items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned backups included in a search
	// job.
	//
	// For example, if a request is made to return MaxResults number of backups,
	// NextToken allows you to return more items in your list starting at the location
	// pointed to by the next token.
	NextToken *string

	// The unique string that specifies the search job.
	SearchJobIdentifier *string

	// The search jobs to be included in the export job can be filtered by including
	// this parameter.
	Status types.ExportJobStatus

	noSmithyDocumentSerde
}

type ListSearchResultExportJobsOutput struct {

	// The operation returns the included export jobs.
	//
	// This member is required.
	ExportJobs []types.ExportJobSummary

	// The next item following a partial list of returned backups included in a search
	// job.
	//
	// For example, if a request is made to return MaxResults number of backups,
	// NextToken allows you to return more items in your list starting at the location
	// pointed to by the next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSearchResultExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSearchResultExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSearchResultExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSearchResultExportJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSearchResultExportJobs(options.Region), middleware.Before); err != nil {
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

// ListSearchResultExportJobsPaginatorOptions is the paginator options for
// ListSearchResultExportJobs
type ListSearchResultExportJobsPaginatorOptions struct {
	// The maximum number of resource list items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSearchResultExportJobsPaginator is a paginator for
// ListSearchResultExportJobs
type ListSearchResultExportJobsPaginator struct {
	options   ListSearchResultExportJobsPaginatorOptions
	client    ListSearchResultExportJobsAPIClient
	params    *ListSearchResultExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListSearchResultExportJobsPaginator returns a new
// ListSearchResultExportJobsPaginator
func NewListSearchResultExportJobsPaginator(client ListSearchResultExportJobsAPIClient, params *ListSearchResultExportJobsInput, optFns ...func(*ListSearchResultExportJobsPaginatorOptions)) *ListSearchResultExportJobsPaginator {
	if params == nil {
		params = &ListSearchResultExportJobsInput{}
	}

	options := ListSearchResultExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSearchResultExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSearchResultExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSearchResultExportJobs page.
func (p *ListSearchResultExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSearchResultExportJobsOutput, error) {
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
	result, err := p.client.ListSearchResultExportJobs(ctx, &params, optFns...)
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

// ListSearchResultExportJobsAPIClient is a client that implements the
// ListSearchResultExportJobs operation.
type ListSearchResultExportJobsAPIClient interface {
	ListSearchResultExportJobs(context.Context, *ListSearchResultExportJobsInput, ...func(*Options)) (*ListSearchResultExportJobsOutput, error)
}

var _ ListSearchResultExportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSearchResultExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSearchResultExportJobs",
	}
}
