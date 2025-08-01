// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List existing model cards.
func (c *Client) ListModelCards(ctx context.Context, params *ListModelCardsInput, optFns ...func(*Options)) (*ListModelCardsOutput, error) {
	if params == nil {
		params = &ListModelCardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelCards", params, optFns, c.addOperationListModelCardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelCardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelCardsInput struct {

	// Only list model cards that were created after the time specified.
	CreationTimeAfter *time.Time

	// Only list model cards that were created before the time specified.
	CreationTimeBefore *time.Time

	// The maximum number of model cards to list.
	MaxResults *int32

	// Only list model cards with the specified approval status.
	ModelCardStatus types.ModelCardStatus

	// Only list model cards with names that contain the specified string.
	NameContains *string

	// If the response to a previous ListModelCards request was truncated, the
	// response includes a NextToken . To retrieve the next set of model cards, use the
	// token in the next request.
	NextToken *string

	// Sort model cards by either name or creation time. Sorts by creation time by
	// default.
	SortBy types.ModelCardSortBy

	// Sort model cards by ascending or descending order.
	SortOrder types.ModelCardSortOrder

	noSmithyDocumentSerde
}

type ListModelCardsOutput struct {

	// The summaries of the listed model cards.
	//
	// This member is required.
	ModelCardSummaries []types.ModelCardSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of model cards, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelCardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelCards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelCards{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListModelCards"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelCards(options.Region), middleware.Before); err != nil {
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

// ListModelCardsPaginatorOptions is the paginator options for ListModelCards
type ListModelCardsPaginatorOptions struct {
	// The maximum number of model cards to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelCardsPaginator is a paginator for ListModelCards
type ListModelCardsPaginator struct {
	options   ListModelCardsPaginatorOptions
	client    ListModelCardsAPIClient
	params    *ListModelCardsInput
	nextToken *string
	firstPage bool
}

// NewListModelCardsPaginator returns a new ListModelCardsPaginator
func NewListModelCardsPaginator(client ListModelCardsAPIClient, params *ListModelCardsInput, optFns ...func(*ListModelCardsPaginatorOptions)) *ListModelCardsPaginator {
	if params == nil {
		params = &ListModelCardsInput{}
	}

	options := ListModelCardsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelCardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelCardsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelCards page.
func (p *ListModelCardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelCardsOutput, error) {
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
	result, err := p.client.ListModelCards(ctx, &params, optFns...)
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

// ListModelCardsAPIClient is a client that implements the ListModelCards
// operation.
type ListModelCardsAPIClient interface {
	ListModelCards(context.Context, *ListModelCardsInput, ...func(*Options)) (*ListModelCardsOutput, error)
}

var _ ListModelCardsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListModelCards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListModelCards",
	}
}
