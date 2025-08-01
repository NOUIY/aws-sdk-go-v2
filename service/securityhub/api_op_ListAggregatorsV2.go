// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of V2 aggregators. This API is in private preview and subject
// to change.
func (c *Client) ListAggregatorsV2(ctx context.Context, params *ListAggregatorsV2Input, optFns ...func(*Options)) (*ListAggregatorsV2Output, error) {
	if params == nil {
		params = &ListAggregatorsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAggregatorsV2", params, optFns, c.addOperationListAggregatorsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAggregatorsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAggregatorsV2Input struct {

	// The maximum number of results to return.
	MaxResults *int32

	// The token required for pagination. On your first call, set the value of this
	// parameter to NULL . For subsequent calls, to continue listing data, set the
	// value of this parameter to the value returned in the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAggregatorsV2Output struct {

	// An array of aggregators.
	AggregatorsV2 []types.AggregatorV2

	// The pagination token to use to request the next page of results. Otherwise,
	// this parameter is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAggregatorsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAggregatorsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAggregatorsV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAggregatorsV2"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAggregatorsV2(options.Region), middleware.Before); err != nil {
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

// ListAggregatorsV2PaginatorOptions is the paginator options for ListAggregatorsV2
type ListAggregatorsV2PaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAggregatorsV2Paginator is a paginator for ListAggregatorsV2
type ListAggregatorsV2Paginator struct {
	options   ListAggregatorsV2PaginatorOptions
	client    ListAggregatorsV2APIClient
	params    *ListAggregatorsV2Input
	nextToken *string
	firstPage bool
}

// NewListAggregatorsV2Paginator returns a new ListAggregatorsV2Paginator
func NewListAggregatorsV2Paginator(client ListAggregatorsV2APIClient, params *ListAggregatorsV2Input, optFns ...func(*ListAggregatorsV2PaginatorOptions)) *ListAggregatorsV2Paginator {
	if params == nil {
		params = &ListAggregatorsV2Input{}
	}

	options := ListAggregatorsV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAggregatorsV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAggregatorsV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAggregatorsV2 page.
func (p *ListAggregatorsV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAggregatorsV2Output, error) {
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
	result, err := p.client.ListAggregatorsV2(ctx, &params, optFns...)
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

// ListAggregatorsV2APIClient is a client that implements the ListAggregatorsV2
// operation.
type ListAggregatorsV2APIClient interface {
	ListAggregatorsV2(context.Context, *ListAggregatorsV2Input, ...func(*Options)) (*ListAggregatorsV2Output, error)
}

var _ ListAggregatorsV2APIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAggregatorsV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAggregatorsV2",
	}
}
