// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists coverage details for your environment.
func (c *Client) ListCoverage(ctx context.Context, params *ListCoverageInput, optFns ...func(*Options)) (*ListCoverageOutput, error) {
	if params == nil {
		params = &ListCoverageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCoverage", params, optFns, c.addOperationListCoverageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCoverageInput struct {

	// An object that contains details on the filters to apply to the coverage data
	// for your environment.
	FilterCriteria *types.CoverageFilterCriteria

	// The maximum number of results the response can return. If your request would
	// return more than the maximum the response will return a nextToken value, use
	// this value when you call the action again to get the remaining results.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. If
	// your response returns more than the maxResults maximum value it will also
	// return a nextToken value. For subsequent calls, use the nextToken value
	// returned from the previous request to continue listing results after the first
	// page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCoverageOutput struct {

	// An object that contains details on the covered resources in your environment.
	CoveredResources []types.CoveredResource

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCoverageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCoverage"); err != nil {
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
	if err = addOpListCoverageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCoverage(options.Region), middleware.Before); err != nil {
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

// ListCoveragePaginatorOptions is the paginator options for ListCoverage
type ListCoveragePaginatorOptions struct {
	// The maximum number of results the response can return. If your request would
	// return more than the maximum the response will return a nextToken value, use
	// this value when you call the action again to get the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCoveragePaginator is a paginator for ListCoverage
type ListCoveragePaginator struct {
	options   ListCoveragePaginatorOptions
	client    ListCoverageAPIClient
	params    *ListCoverageInput
	nextToken *string
	firstPage bool
}

// NewListCoveragePaginator returns a new ListCoveragePaginator
func NewListCoveragePaginator(client ListCoverageAPIClient, params *ListCoverageInput, optFns ...func(*ListCoveragePaginatorOptions)) *ListCoveragePaginator {
	if params == nil {
		params = &ListCoverageInput{}
	}

	options := ListCoveragePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCoveragePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCoveragePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCoverage page.
func (p *ListCoveragePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCoverageOutput, error) {
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
	result, err := p.client.ListCoverage(ctx, &params, optFns...)
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

// ListCoverageAPIClient is a client that implements the ListCoverage operation.
type ListCoverageAPIClient interface {
	ListCoverage(context.Context, *ListCoverageInput, ...func(*Options)) (*ListCoverageOutput, error)
}

var _ ListCoverageAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCoverage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCoverage",
	}
}
