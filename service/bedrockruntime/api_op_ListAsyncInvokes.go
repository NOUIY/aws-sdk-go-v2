// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists asynchronous invocations.
func (c *Client) ListAsyncInvokes(ctx context.Context, params *ListAsyncInvokesInput, optFns ...func(*Options)) (*ListAsyncInvokesOutput, error) {
	if params == nil {
		params = &ListAsyncInvokesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAsyncInvokes", params, optFns, c.addOperationListAsyncInvokesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAsyncInvokesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAsyncInvokesInput struct {

	// The maximum number of invocations to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// How to sort the response.
	SortBy types.SortAsyncInvocationBy

	// The sorting order for the response.
	SortOrder types.SortOrder

	// Filter invocations by status.
	StatusEquals types.AsyncInvokeStatus

	// Include invocations submitted after this time.
	SubmitTimeAfter *time.Time

	// Include invocations submitted before this time.
	SubmitTimeBefore *time.Time

	noSmithyDocumentSerde
}

type ListAsyncInvokesOutput struct {

	// A list of invocation summaries.
	AsyncInvokeSummaries []types.AsyncInvokeSummary

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAsyncInvokesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAsyncInvokes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAsyncInvokes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAsyncInvokes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAsyncInvokes(options.Region), middleware.Before); err != nil {
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

// ListAsyncInvokesPaginatorOptions is the paginator options for ListAsyncInvokes
type ListAsyncInvokesPaginatorOptions struct {
	// The maximum number of invocations to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAsyncInvokesPaginator is a paginator for ListAsyncInvokes
type ListAsyncInvokesPaginator struct {
	options   ListAsyncInvokesPaginatorOptions
	client    ListAsyncInvokesAPIClient
	params    *ListAsyncInvokesInput
	nextToken *string
	firstPage bool
}

// NewListAsyncInvokesPaginator returns a new ListAsyncInvokesPaginator
func NewListAsyncInvokesPaginator(client ListAsyncInvokesAPIClient, params *ListAsyncInvokesInput, optFns ...func(*ListAsyncInvokesPaginatorOptions)) *ListAsyncInvokesPaginator {
	if params == nil {
		params = &ListAsyncInvokesInput{}
	}

	options := ListAsyncInvokesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAsyncInvokesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAsyncInvokesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAsyncInvokes page.
func (p *ListAsyncInvokesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAsyncInvokesOutput, error) {
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
	result, err := p.client.ListAsyncInvokes(ctx, &params, optFns...)
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

// ListAsyncInvokesAPIClient is a client that implements the ListAsyncInvokes
// operation.
type ListAsyncInvokesAPIClient interface {
	ListAsyncInvokes(context.Context, *ListAsyncInvokesInput, ...func(*Options)) (*ListAsyncInvokesOutput, error)
}

var _ ListAsyncInvokesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAsyncInvokes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAsyncInvokes",
	}
}
