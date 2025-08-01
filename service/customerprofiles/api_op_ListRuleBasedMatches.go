// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of MatchIds that belong to the given domain.
func (c *Client) ListRuleBasedMatches(ctx context.Context, params *ListRuleBasedMatchesInput, optFns ...func(*Options)) (*ListRuleBasedMatchesOutput, error) {
	if params == nil {
		params = &ListRuleBasedMatchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRuleBasedMatches", params, optFns, c.addOperationListRuleBasedMatchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRuleBasedMatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRuleBasedMatchesInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The maximum number of MatchIds returned per page.
	MaxResults *int32

	// The pagination token from the previous ListRuleBasedMatches API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRuleBasedMatchesOutput struct {

	// The list of MatchIds for the given domain.
	MatchIds []string

	// The pagination token from the previous ListRuleBasedMatches API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRuleBasedMatchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRuleBasedMatches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRuleBasedMatches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRuleBasedMatches"); err != nil {
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
	if err = addOpListRuleBasedMatchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRuleBasedMatches(options.Region), middleware.Before); err != nil {
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

// ListRuleBasedMatchesPaginatorOptions is the paginator options for
// ListRuleBasedMatches
type ListRuleBasedMatchesPaginatorOptions struct {
	// The maximum number of MatchIds returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRuleBasedMatchesPaginator is a paginator for ListRuleBasedMatches
type ListRuleBasedMatchesPaginator struct {
	options   ListRuleBasedMatchesPaginatorOptions
	client    ListRuleBasedMatchesAPIClient
	params    *ListRuleBasedMatchesInput
	nextToken *string
	firstPage bool
}

// NewListRuleBasedMatchesPaginator returns a new ListRuleBasedMatchesPaginator
func NewListRuleBasedMatchesPaginator(client ListRuleBasedMatchesAPIClient, params *ListRuleBasedMatchesInput, optFns ...func(*ListRuleBasedMatchesPaginatorOptions)) *ListRuleBasedMatchesPaginator {
	if params == nil {
		params = &ListRuleBasedMatchesInput{}
	}

	options := ListRuleBasedMatchesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRuleBasedMatchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRuleBasedMatchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRuleBasedMatches page.
func (p *ListRuleBasedMatchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRuleBasedMatchesOutput, error) {
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
	result, err := p.client.ListRuleBasedMatches(ctx, &params, optFns...)
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

// ListRuleBasedMatchesAPIClient is a client that implements the
// ListRuleBasedMatches operation.
type ListRuleBasedMatchesAPIClient interface {
	ListRuleBasedMatches(context.Context, *ListRuleBasedMatchesInput, ...func(*Options)) (*ListRuleBasedMatchesOutput, error)
}

var _ ListRuleBasedMatchesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRuleBasedMatches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRuleBasedMatches",
	}
}
