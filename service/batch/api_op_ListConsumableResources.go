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

// Returns a list of Batch consumable resources.
func (c *Client) ListConsumableResources(ctx context.Context, params *ListConsumableResourcesInput, optFns ...func(*Options)) (*ListConsumableResourcesOutput, error) {
	if params == nil {
		params = &ListConsumableResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConsumableResources", params, optFns, c.addOperationListConsumableResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConsumableResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConsumableResourcesInput struct {

	// The filters to apply to the consumable resource list query. If used, only those
	// consumable resources that match the filter are listed. Filter names and values
	// can be:
	//
	//   - name: CONSUMABLE_RESOURCE_NAME
	//
	// values: case-insensitive matches for the consumable resource name. If a filter
	//   value ends with an asterisk (*), it matches any consumable resource name that
	//   begins with the string before the '*'.
	Filters []types.KeyValuesPair

	// The maximum number of results returned by ListConsumableResources in paginated
	// output. When this parameter is used, ListConsumableResources only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListConsumableResources request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, then
	// ListConsumableResources returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListConsumableResources
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	//
	// Treat this token as an opaque identifier that's only used to retrieve the next
	// items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConsumableResourcesOutput struct {

	// A list of consumable resources that match the request.
	//
	// This member is required.
	ConsumableResources []types.ConsumableResourceSummary

	// The nextToken value to include in a future ListConsumableResources request.
	// When the results of a ListConsumableResources request exceed maxResults , this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConsumableResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConsumableResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConsumableResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConsumableResources"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConsumableResources(options.Region), middleware.Before); err != nil {
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

// ListConsumableResourcesPaginatorOptions is the paginator options for
// ListConsumableResources
type ListConsumableResourcesPaginatorOptions struct {
	// The maximum number of results returned by ListConsumableResources in paginated
	// output. When this parameter is used, ListConsumableResources only returns
	// maxResults results in a single page and a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListConsumableResources request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter isn't used, then
	// ListConsumableResources returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConsumableResourcesPaginator is a paginator for ListConsumableResources
type ListConsumableResourcesPaginator struct {
	options   ListConsumableResourcesPaginatorOptions
	client    ListConsumableResourcesAPIClient
	params    *ListConsumableResourcesInput
	nextToken *string
	firstPage bool
}

// NewListConsumableResourcesPaginator returns a new
// ListConsumableResourcesPaginator
func NewListConsumableResourcesPaginator(client ListConsumableResourcesAPIClient, params *ListConsumableResourcesInput, optFns ...func(*ListConsumableResourcesPaginatorOptions)) *ListConsumableResourcesPaginator {
	if params == nil {
		params = &ListConsumableResourcesInput{}
	}

	options := ListConsumableResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConsumableResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConsumableResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConsumableResources page.
func (p *ListConsumableResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConsumableResourcesOutput, error) {
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
	result, err := p.client.ListConsumableResources(ctx, &params, optFns...)
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

// ListConsumableResourcesAPIClient is a client that implements the
// ListConsumableResources operation.
type ListConsumableResourcesAPIClient interface {
	ListConsumableResources(context.Context, *ListConsumableResourcesInput, ...func(*Options)) (*ListConsumableResourcesOutput, error)
}

var _ ListConsumableResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListConsumableResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConsumableResources",
	}
}
