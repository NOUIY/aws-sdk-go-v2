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

// This operation returns a list of a specified search job.
func (c *Client) ListSearchJobResults(ctx context.Context, params *ListSearchJobResultsInput, optFns ...func(*Options)) (*ListSearchJobResultsOutput, error) {
	if params == nil {
		params = &ListSearchJobResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSearchJobResults", params, optFns, c.addOperationListSearchJobResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSearchJobResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSearchJobResultsInput struct {

	// The unique string that specifies the search job.
	//
	// This member is required.
	SearchJobIdentifier *string

	// The maximum number of resource list items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned search job results.
	//
	// For example, if a request is made to return MaxResults number of search job
	// results, NextToken allows you to return more items in your list starting at the
	// location pointed to by the next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSearchJobResultsOutput struct {

	// The results consist of either EBSResultItem or S3ResultItem.
	//
	// This member is required.
	Results []types.ResultItem

	// The next item following a partial list of search job results.
	//
	// For example, if a request is made to return MaxResults number of backups,
	// NextToken allows you to return more items in your list starting at the location
	// pointed to by the next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSearchJobResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSearchJobResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSearchJobResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSearchJobResults"); err != nil {
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
	if err = addOpListSearchJobResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSearchJobResults(options.Region), middleware.Before); err != nil {
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

// ListSearchJobResultsPaginatorOptions is the paginator options for
// ListSearchJobResults
type ListSearchJobResultsPaginatorOptions struct {
	// The maximum number of resource list items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSearchJobResultsPaginator is a paginator for ListSearchJobResults
type ListSearchJobResultsPaginator struct {
	options   ListSearchJobResultsPaginatorOptions
	client    ListSearchJobResultsAPIClient
	params    *ListSearchJobResultsInput
	nextToken *string
	firstPage bool
}

// NewListSearchJobResultsPaginator returns a new ListSearchJobResultsPaginator
func NewListSearchJobResultsPaginator(client ListSearchJobResultsAPIClient, params *ListSearchJobResultsInput, optFns ...func(*ListSearchJobResultsPaginatorOptions)) *ListSearchJobResultsPaginator {
	if params == nil {
		params = &ListSearchJobResultsInput{}
	}

	options := ListSearchJobResultsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSearchJobResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSearchJobResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSearchJobResults page.
func (p *ListSearchJobResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSearchJobResultsOutput, error) {
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
	result, err := p.client.ListSearchJobResults(ctx, &params, optFns...)
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

// ListSearchJobResultsAPIClient is a client that implements the
// ListSearchJobResults operation.
type ListSearchJobResultsAPIClient interface {
	ListSearchJobResults(context.Context, *ListSearchJobResultsInput, ...func(*Options)) (*ListSearchJobResultsOutput, error)
}

var _ ListSearchJobResultsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSearchJobResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSearchJobResults",
	}
}
