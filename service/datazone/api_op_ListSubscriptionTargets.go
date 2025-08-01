// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists subscription targets in Amazon DataZone.
func (c *Client) ListSubscriptionTargets(ctx context.Context, params *ListSubscriptionTargetsInput, optFns ...func(*Options)) (*ListSubscriptionTargetsOutput, error) {
	if params == nil {
		params = &ListSubscriptionTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptionTargets", params, optFns, c.addOperationListSubscriptionTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscriptionTargetsInput struct {

	// The identifier of the Amazon DataZone domain where you want to list
	// subscription targets.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the environment where you want to list subscription targets.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The maximum number of subscription targets to return in a single call to
	// ListSubscriptionTargets . When the number of subscription targets to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListSubscriptionTargets to list
	// the next set of subscription targets.
	MaxResults *int32

	// When the number of subscription targets is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription targets, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionTargets to list the next set of subscription
	// targets.
	NextToken *string

	// Specifies the way in which the results of this action are to be sorted.
	SortBy types.SortKey

	// Specifies the sort order for the results of this action.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListSubscriptionTargetsOutput struct {

	// The results of the ListSubscriptionTargets action.
	//
	// This member is required.
	Items []types.SubscriptionTargetSummary

	// When the number of subscription targets is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription targets, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionTargets to list the next set of subscription
	// targets.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscriptionTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubscriptionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubscriptionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubscriptionTargets"); err != nil {
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
	if err = addOpListSubscriptionTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionTargets(options.Region), middleware.Before); err != nil {
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

// ListSubscriptionTargetsPaginatorOptions is the paginator options for
// ListSubscriptionTargets
type ListSubscriptionTargetsPaginatorOptions struct {
	// The maximum number of subscription targets to return in a single call to
	// ListSubscriptionTargets . When the number of subscription targets to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListSubscriptionTargets to list
	// the next set of subscription targets.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscriptionTargetsPaginator is a paginator for ListSubscriptionTargets
type ListSubscriptionTargetsPaginator struct {
	options   ListSubscriptionTargetsPaginatorOptions
	client    ListSubscriptionTargetsAPIClient
	params    *ListSubscriptionTargetsInput
	nextToken *string
	firstPage bool
}

// NewListSubscriptionTargetsPaginator returns a new
// ListSubscriptionTargetsPaginator
func NewListSubscriptionTargetsPaginator(client ListSubscriptionTargetsAPIClient, params *ListSubscriptionTargetsInput, optFns ...func(*ListSubscriptionTargetsPaginatorOptions)) *ListSubscriptionTargetsPaginator {
	if params == nil {
		params = &ListSubscriptionTargetsInput{}
	}

	options := ListSubscriptionTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscriptionTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscriptionTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubscriptionTargets page.
func (p *ListSubscriptionTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscriptionTargetsOutput, error) {
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
	result, err := p.client.ListSubscriptionTargets(ctx, &params, optFns...)
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

// ListSubscriptionTargetsAPIClient is a client that implements the
// ListSubscriptionTargets operation.
type ListSubscriptionTargetsAPIClient interface {
	ListSubscriptionTargets(context.Context, *ListSubscriptionTargetsInput, ...func(*Options)) (*ListSubscriptionTargetsOutput, error)
}

var _ ListSubscriptionTargetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSubscriptionTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubscriptionTargets",
	}
}
