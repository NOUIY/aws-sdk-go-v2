// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about playback restriction policies.
func (c *Client) ListPlaybackRestrictionPolicies(ctx context.Context, params *ListPlaybackRestrictionPoliciesInput, optFns ...func(*Options)) (*ListPlaybackRestrictionPoliciesOutput, error) {
	if params == nil {
		params = &ListPlaybackRestrictionPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlaybackRestrictionPolicies", params, optFns, c.addOperationListPlaybackRestrictionPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlaybackRestrictionPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaybackRestrictionPoliciesInput struct {

	// Maximum number of policies to return. Default: 1.
	MaxResults *int32

	// The first policy to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlaybackRestrictionPoliciesOutput struct {

	// List of the matching policies.
	//
	// This member is required.
	PlaybackRestrictionPolicies []types.PlaybackRestrictionPolicySummary

	// If there are more channels than maxResults , use nextToken in the request to
	// get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlaybackRestrictionPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlaybackRestrictionPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaybackRestrictionPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPlaybackRestrictionPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlaybackRestrictionPolicies(options.Region), middleware.Before); err != nil {
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

// ListPlaybackRestrictionPoliciesPaginatorOptions is the paginator options for
// ListPlaybackRestrictionPolicies
type ListPlaybackRestrictionPoliciesPaginatorOptions struct {
	// Maximum number of policies to return. Default: 1.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlaybackRestrictionPoliciesPaginator is a paginator for
// ListPlaybackRestrictionPolicies
type ListPlaybackRestrictionPoliciesPaginator struct {
	options   ListPlaybackRestrictionPoliciesPaginatorOptions
	client    ListPlaybackRestrictionPoliciesAPIClient
	params    *ListPlaybackRestrictionPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListPlaybackRestrictionPoliciesPaginator returns a new
// ListPlaybackRestrictionPoliciesPaginator
func NewListPlaybackRestrictionPoliciesPaginator(client ListPlaybackRestrictionPoliciesAPIClient, params *ListPlaybackRestrictionPoliciesInput, optFns ...func(*ListPlaybackRestrictionPoliciesPaginatorOptions)) *ListPlaybackRestrictionPoliciesPaginator {
	if params == nil {
		params = &ListPlaybackRestrictionPoliciesInput{}
	}

	options := ListPlaybackRestrictionPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlaybackRestrictionPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlaybackRestrictionPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlaybackRestrictionPolicies page.
func (p *ListPlaybackRestrictionPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlaybackRestrictionPoliciesOutput, error) {
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
	result, err := p.client.ListPlaybackRestrictionPolicies(ctx, &params, optFns...)
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

// ListPlaybackRestrictionPoliciesAPIClient is a client that implements the
// ListPlaybackRestrictionPolicies operation.
type ListPlaybackRestrictionPoliciesAPIClient interface {
	ListPlaybackRestrictionPolicies(context.Context, *ListPlaybackRestrictionPoliciesInput, ...func(*Options)) (*ListPlaybackRestrictionPoliciesOutput, error)
}

var _ ListPlaybackRestrictionPoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPlaybackRestrictionPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPlaybackRestrictionPolicies",
	}
}
