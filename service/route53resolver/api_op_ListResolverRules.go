// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Resolver rules that were created using the current Amazon Web
// Services account.
func (c *Client) ListResolverRules(ctx context.Context, params *ListResolverRulesInput, optFns ...func(*Options)) (*ListResolverRulesOutput, error) {
	if params == nil {
		params = &ListResolverRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverRules", params, optFns, c.addOperationListResolverRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverRulesInput struct {

	// An optional specification to return a subset of Resolver rules, such as all
	// Resolver rules that are associated with the same Resolver endpoint.
	//
	// If you submit a second or subsequent ListResolverRules request and specify the
	// NextToken parameter, you must use the same values for Filters , if any, as in
	// the previous request.
	Filters []types.Filter

	// The maximum number of Resolver rules that you want to return in the response to
	// a ListResolverRules request. If you don't specify a value for MaxResults ,
	// Resolver returns up to 100 Resolver rules.
	MaxResults *int32

	// For the first ListResolverRules request, omit this value.
	//
	// If you have more than MaxResults Resolver rules, you can submit another
	// ListResolverRules request to get the next group of Resolver rules. In the next
	// request, specify the value of NextToken from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverRulesOutput struct {

	// The value that you specified for MaxResults in the request.
	MaxResults *int32

	// If more than MaxResults Resolver rules match the specified criteria, you can
	// submit another ListResolverRules request to get the next group of results. In
	// the next request, specify the value of NextToken from the previous response.
	NextToken *string

	// The Resolver rules that were created using the current Amazon Web Services
	// account and that match the specified filters, if any.
	ResolverRules []types.ResolverRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResolverRules"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverRules(options.Region), middleware.Before); err != nil {
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

// ListResolverRulesPaginatorOptions is the paginator options for ListResolverRules
type ListResolverRulesPaginatorOptions struct {
	// The maximum number of Resolver rules that you want to return in the response to
	// a ListResolverRules request. If you don't specify a value for MaxResults ,
	// Resolver returns up to 100 Resolver rules.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverRulesPaginator is a paginator for ListResolverRules
type ListResolverRulesPaginator struct {
	options   ListResolverRulesPaginatorOptions
	client    ListResolverRulesAPIClient
	params    *ListResolverRulesInput
	nextToken *string
	firstPage bool
}

// NewListResolverRulesPaginator returns a new ListResolverRulesPaginator
func NewListResolverRulesPaginator(client ListResolverRulesAPIClient, params *ListResolverRulesInput, optFns ...func(*ListResolverRulesPaginatorOptions)) *ListResolverRulesPaginator {
	if params == nil {
		params = &ListResolverRulesInput{}
	}

	options := ListResolverRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverRules page.
func (p *ListResolverRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverRulesOutput, error) {
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
	result, err := p.client.ListResolverRules(ctx, &params, optFns...)
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

// ListResolverRulesAPIClient is a client that implements the ListResolverRules
// operation.
type ListResolverRulesAPIClient interface {
	ListResolverRules(context.Context, *ListResolverRulesInput, ...func(*Options)) (*ListResolverRulesOutput, error)
}

var _ ListResolverRulesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListResolverRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResolverRules",
	}
}
