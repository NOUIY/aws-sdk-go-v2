// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerlinuxsubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List Bring Your Own License (BYOL) subscription registration resources for your
// account.
func (c *Client) ListRegisteredSubscriptionProviders(ctx context.Context, params *ListRegisteredSubscriptionProvidersInput, optFns ...func(*Options)) (*ListRegisteredSubscriptionProvidersOutput, error) {
	if params == nil {
		params = &ListRegisteredSubscriptionProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegisteredSubscriptionProviders", params, optFns, c.addOperationListRegisteredSubscriptionProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegisteredSubscriptionProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegisteredSubscriptionProvidersInput struct {

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	// To filter your results, specify which subscription providers to return in the
	// list.
	SubscriptionProviderSources []types.SubscriptionProviderSource

	noSmithyDocumentSerde
}

type ListRegisteredSubscriptionProvidersOutput struct {

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The list of BYOL registration resources that fit the criteria you specified in
	// the request.
	RegisteredSubscriptionProviders []types.RegisteredSubscriptionProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRegisteredSubscriptionProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRegisteredSubscriptionProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRegisteredSubscriptionProviders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRegisteredSubscriptionProviders"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegisteredSubscriptionProviders(options.Region), middleware.Before); err != nil {
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

// ListRegisteredSubscriptionProvidersPaginatorOptions is the paginator options
// for ListRegisteredSubscriptionProviders
type ListRegisteredSubscriptionProvidersPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegisteredSubscriptionProvidersPaginator is a paginator for
// ListRegisteredSubscriptionProviders
type ListRegisteredSubscriptionProvidersPaginator struct {
	options   ListRegisteredSubscriptionProvidersPaginatorOptions
	client    ListRegisteredSubscriptionProvidersAPIClient
	params    *ListRegisteredSubscriptionProvidersInput
	nextToken *string
	firstPage bool
}

// NewListRegisteredSubscriptionProvidersPaginator returns a new
// ListRegisteredSubscriptionProvidersPaginator
func NewListRegisteredSubscriptionProvidersPaginator(client ListRegisteredSubscriptionProvidersAPIClient, params *ListRegisteredSubscriptionProvidersInput, optFns ...func(*ListRegisteredSubscriptionProvidersPaginatorOptions)) *ListRegisteredSubscriptionProvidersPaginator {
	if params == nil {
		params = &ListRegisteredSubscriptionProvidersInput{}
	}

	options := ListRegisteredSubscriptionProvidersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRegisteredSubscriptionProvidersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegisteredSubscriptionProvidersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRegisteredSubscriptionProviders page.
func (p *ListRegisteredSubscriptionProvidersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegisteredSubscriptionProvidersOutput, error) {
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
	result, err := p.client.ListRegisteredSubscriptionProviders(ctx, &params, optFns...)
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

// ListRegisteredSubscriptionProvidersAPIClient is a client that implements the
// ListRegisteredSubscriptionProviders operation.
type ListRegisteredSubscriptionProvidersAPIClient interface {
	ListRegisteredSubscriptionProviders(context.Context, *ListRegisteredSubscriptionProvidersInput, ...func(*Options)) (*ListRegisteredSubscriptionProvidersOutput, error)
}

var _ ListRegisteredSubscriptionProvidersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRegisteredSubscriptionProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRegisteredSubscriptionProviders",
	}
}
