// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of built-in intents provided by Amazon Lex that you can use in your
// bot.
//
// To use a built-in intent as a the base for your own intent, include the
// built-in intent signature in the parentIntentSignature parameter when you call
// the CreateIntent operation. For more information, see [CreateIntent].
//
// [CreateIntent]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_CreateIntent.html
func (c *Client) ListBuiltInIntents(ctx context.Context, params *ListBuiltInIntentsInput, optFns ...func(*Options)) (*ListBuiltInIntentsOutput, error) {
	if params == nil {
		params = &ListBuiltInIntentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuiltInIntents", params, optFns, c.addOperationListBuiltInIntentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuiltInIntentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuiltInIntentsInput struct {

	// The identifier of the language and locale of the intents to list. The string
	// must match one of the supported locales. For more information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	//
	// This member is required.
	LocaleId *string

	// The maximum number of built-in intents to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the ListBuiltInIntents operation contains more results
	// than specified in the maxResults parameter, a token is returned in the
	// response. Use that token in the nextToken parameter to return the next page of
	// results.
	NextToken *string

	// Specifies sorting parameters for the list of built-in intents. You can specify
	// that the list be sorted by the built-in intent signature in either ascending or
	// descending order.
	SortBy *types.BuiltInIntentSortBy

	noSmithyDocumentSerde
}

type ListBuiltInIntentsOutput struct {

	// Summary information for the built-in intents that meet the filter criteria
	// specified in the request. The length of the list is specified in the maxResults
	// parameter of the request. If there are more intents available, the nextToken
	// field contains a token to get the next page of results.
	BuiltInIntentSummaries []types.BuiltInIntentSummary

	// The language and locale of the intents in the list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListBuiltInIntents operation. If the nextToken field is present, you
	// send the contents as the nextToken parameter of a ListBotAliases operation
	// request to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBuiltInIntentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBuiltInIntents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBuiltInIntents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBuiltInIntents"); err != nil {
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
	if err = addOpListBuiltInIntentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuiltInIntents(options.Region), middleware.Before); err != nil {
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

// ListBuiltInIntentsPaginatorOptions is the paginator options for
// ListBuiltInIntents
type ListBuiltInIntentsPaginatorOptions struct {
	// The maximum number of built-in intents to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuiltInIntentsPaginator is a paginator for ListBuiltInIntents
type ListBuiltInIntentsPaginator struct {
	options   ListBuiltInIntentsPaginatorOptions
	client    ListBuiltInIntentsAPIClient
	params    *ListBuiltInIntentsInput
	nextToken *string
	firstPage bool
}

// NewListBuiltInIntentsPaginator returns a new ListBuiltInIntentsPaginator
func NewListBuiltInIntentsPaginator(client ListBuiltInIntentsAPIClient, params *ListBuiltInIntentsInput, optFns ...func(*ListBuiltInIntentsPaginatorOptions)) *ListBuiltInIntentsPaginator {
	if params == nil {
		params = &ListBuiltInIntentsInput{}
	}

	options := ListBuiltInIntentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuiltInIntentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuiltInIntentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBuiltInIntents page.
func (p *ListBuiltInIntentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuiltInIntentsOutput, error) {
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
	result, err := p.client.ListBuiltInIntents(ctx, &params, optFns...)
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

// ListBuiltInIntentsAPIClient is a client that implements the ListBuiltInIntents
// operation.
type ListBuiltInIntentsAPIClient interface {
	ListBuiltInIntents(context.Context, *ListBuiltInIntentsInput, ...func(*Options)) (*ListBuiltInIntentsOutput, error)
}

var _ ListBuiltInIntentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBuiltInIntents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBuiltInIntents",
	}
}
