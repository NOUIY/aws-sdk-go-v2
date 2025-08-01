// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the transaction events for a transaction
//
// This action will return transaction details for all transactions that are
// confirmed on the blockchain, even if they have not reached [finality].
//
// [finality]: https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality
func (c *Client) ListTransactionEvents(ctx context.Context, params *ListTransactionEventsInput, optFns ...func(*Options)) (*ListTransactionEventsOutput, error) {
	if params == nil {
		params = &ListTransactionEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransactionEvents", params, optFns, c.addOperationListTransactionEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransactionEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransactionEventsInput struct {

	// The blockchain network where the transaction events occurred.
	//
	// This member is required.
	Network types.QueryNetwork

	// The maximum number of transaction events to list.
	//
	// Default: 100
	//
	// Even if additional results can be retrieved, the request can return less
	// results than maxResults or an empty array of results.
	//
	// To retrieve the next set of results, make another request with the returned
	// nextToken value. The value of nextToken is null when there are no more results
	// to return
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The hash of a transaction. It is generated when a transaction is created.
	TransactionHash *string

	// The identifier of a Bitcoin transaction. It is generated when a transaction is
	// created.
	//
	// transactionId is only supported on the Bitcoin networks.
	TransactionId *string

	noSmithyDocumentSerde
}

type ListTransactionEventsOutput struct {

	// An array of TransactionEvent objects. Each object contains details about the
	// transaction events.
	//
	// This member is required.
	Events []types.TransactionEvent

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTransactionEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTransactionEvents"); err != nil {
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
	if err = addOpListTransactionEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTransactionEvents(options.Region), middleware.Before); err != nil {
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

// ListTransactionEventsPaginatorOptions is the paginator options for
// ListTransactionEvents
type ListTransactionEventsPaginatorOptions struct {
	// The maximum number of transaction events to list.
	//
	// Default: 100
	//
	// Even if additional results can be retrieved, the request can return less
	// results than maxResults or an empty array of results.
	//
	// To retrieve the next set of results, make another request with the returned
	// nextToken value. The value of nextToken is null when there are no more results
	// to return
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTransactionEventsPaginator is a paginator for ListTransactionEvents
type ListTransactionEventsPaginator struct {
	options   ListTransactionEventsPaginatorOptions
	client    ListTransactionEventsAPIClient
	params    *ListTransactionEventsInput
	nextToken *string
	firstPage bool
}

// NewListTransactionEventsPaginator returns a new ListTransactionEventsPaginator
func NewListTransactionEventsPaginator(client ListTransactionEventsAPIClient, params *ListTransactionEventsInput, optFns ...func(*ListTransactionEventsPaginatorOptions)) *ListTransactionEventsPaginator {
	if params == nil {
		params = &ListTransactionEventsInput{}
	}

	options := ListTransactionEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTransactionEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTransactionEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTransactionEvents page.
func (p *ListTransactionEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTransactionEventsOutput, error) {
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
	result, err := p.client.ListTransactionEvents(ctx, &params, optFns...)
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

// ListTransactionEventsAPIClient is a client that implements the
// ListTransactionEvents operation.
type ListTransactionEventsAPIClient interface {
	ListTransactionEvents(context.Context, *ListTransactionEventsInput, ...func(*Options)) (*ListTransactionEventsOutput, error)
}

var _ ListTransactionEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTransactionEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTransactionEvents",
	}
}
