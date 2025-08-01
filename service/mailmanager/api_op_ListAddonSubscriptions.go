// Code generated by smithy-go-codegen DO NOT EDIT.

package mailmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Add On subscriptions in your account.
func (c *Client) ListAddonSubscriptions(ctx context.Context, params *ListAddonSubscriptionsInput, optFns ...func(*Options)) (*ListAddonSubscriptionsOutput, error) {
	if params == nil {
		params = &ListAddonSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAddonSubscriptions", params, optFns, c.addOperationListAddonSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAddonSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAddonSubscriptionsInput struct {

	// If you received a pagination token from a previous call to this API, you can
	// provide it here to continue paginating through the next page of results.
	NextToken *string

	// The maximum number of ingress endpoint resources that are returned per call.
	// You can use NextToken to obtain further ingress endpoints.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListAddonSubscriptionsOutput struct {

	// The list of ingress endpoints.
	AddonSubscriptions []types.AddonSubscription

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAddonSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListAddonSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListAddonSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAddonSubscriptions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAddonSubscriptions(options.Region), middleware.Before); err != nil {
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

// ListAddonSubscriptionsPaginatorOptions is the paginator options for
// ListAddonSubscriptions
type ListAddonSubscriptionsPaginatorOptions struct {
	// The maximum number of ingress endpoint resources that are returned per call.
	// You can use NextToken to obtain further ingress endpoints.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAddonSubscriptionsPaginator is a paginator for ListAddonSubscriptions
type ListAddonSubscriptionsPaginator struct {
	options   ListAddonSubscriptionsPaginatorOptions
	client    ListAddonSubscriptionsAPIClient
	params    *ListAddonSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListAddonSubscriptionsPaginator returns a new ListAddonSubscriptionsPaginator
func NewListAddonSubscriptionsPaginator(client ListAddonSubscriptionsAPIClient, params *ListAddonSubscriptionsInput, optFns ...func(*ListAddonSubscriptionsPaginatorOptions)) *ListAddonSubscriptionsPaginator {
	if params == nil {
		params = &ListAddonSubscriptionsInput{}
	}

	options := ListAddonSubscriptionsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAddonSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAddonSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAddonSubscriptions page.
func (p *ListAddonSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAddonSubscriptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAddonSubscriptions(ctx, &params, optFns...)
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

// ListAddonSubscriptionsAPIClient is a client that implements the
// ListAddonSubscriptions operation.
type ListAddonSubscriptionsAPIClient interface {
	ListAddonSubscriptions(context.Context, *ListAddonSubscriptionsInput, ...func(*Options)) (*ListAddonSubscriptionsOutput, error)
}

var _ ListAddonSubscriptionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAddonSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAddonSubscriptions",
	}
}
