// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Displays a list of all entitlements that have been granted to this account.
//
// This request returns 20 results per page.
func (c *Client) ListEntitlements(ctx context.Context, params *ListEntitlementsInput, optFns ...func(*Options)) (*ListEntitlementsOutput, error) {
	if params == nil {
		params = &ListEntitlementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitlements", params, optFns, c.addOperationListEntitlementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitlementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitlementsInput struct {

	//  The maximum number of results to return per API request.
	//
	// For example, you submit a ListEntitlements request with set at 5. Although 20
	// items match your request, the service returns no more than the first 5 items.
	// (The service also returns a NextToken value that you can use to fetch the next
	// batch of results.)
	//
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 20 results per page.
	MaxResults *int32

	//  The token that identifies the batch of results that you want to see.
	//
	// For example, you submit a ListEntitlements request with MaxResults set at 5.
	// The service returns the first batch of results (up to 5) and a NextToken value.
	// To see the next batch of results, you can submit the ListEntitlements request a
	// second time and specify the NextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntitlementsOutput struct {

	// A list of entitlements that have been granted to you from other Amazon Web
	// Services accounts.
	Entitlements []types.ListedEntitlement

	// The token that identifies the batch of results that you want to see.
	//
	// For example, you submit a ListEntitlements request with MaxResults set at 5.
	// The service returns the first batch of results (up to 5) and a NextToken value.
	// To see the next batch of results, you can submit the ListEntitlements request a
	// second time and specify the NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntitlementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEntitlements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEntitlements{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEntitlements"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitlements(options.Region), middleware.Before); err != nil {
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

// ListEntitlementsPaginatorOptions is the paginator options for ListEntitlements
type ListEntitlementsPaginatorOptions struct {
	//  The maximum number of results to return per API request.
	//
	// For example, you submit a ListEntitlements request with set at 5. Although 20
	// items match your request, the service returns no more than the first 5 items.
	// (The service also returns a NextToken value that you can use to fetch the next
	// batch of results.)
	//
	// The service might return fewer results than the MaxResults value. If MaxResults
	// is not included in the request, the service defaults to pagination with a
	// maximum of 20 results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntitlementsPaginator is a paginator for ListEntitlements
type ListEntitlementsPaginator struct {
	options   ListEntitlementsPaginatorOptions
	client    ListEntitlementsAPIClient
	params    *ListEntitlementsInput
	nextToken *string
	firstPage bool
}

// NewListEntitlementsPaginator returns a new ListEntitlementsPaginator
func NewListEntitlementsPaginator(client ListEntitlementsAPIClient, params *ListEntitlementsInput, optFns ...func(*ListEntitlementsPaginatorOptions)) *ListEntitlementsPaginator {
	if params == nil {
		params = &ListEntitlementsInput{}
	}

	options := ListEntitlementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEntitlementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntitlementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEntitlements page.
func (p *ListEntitlementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntitlementsOutput, error) {
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
	result, err := p.client.ListEntitlements(ctx, &params, optFns...)
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

// ListEntitlementsAPIClient is a client that implements the ListEntitlements
// operation.
type ListEntitlementsAPIClient interface {
	ListEntitlements(context.Context, *ListEntitlementsInput, ...func(*Options)) (*ListEntitlementsOutput, error)
}

var _ ListEntitlementsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEntitlements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEntitlements",
	}
}
