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

// List FeatureGroup s based on given filter and order.
func (c *Client) ListFeatureGroups(ctx context.Context, params *ListFeatureGroupsInput, optFns ...func(*Options)) (*ListFeatureGroupsOutput, error) {
	if params == nil {
		params = &ListFeatureGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFeatureGroups", params, optFns, c.addOperationListFeatureGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFeatureGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFeatureGroupsInput struct {

	// Use this parameter to search for FeatureGroups s created after a specific date
	// and time.
	CreationTimeAfter *time.Time

	// Use this parameter to search for FeatureGroups s created before a specific date
	// and time.
	CreationTimeBefore *time.Time

	// A FeatureGroup status. Filters by FeatureGroup status.
	FeatureGroupStatusEquals types.FeatureGroupStatus

	// The maximum number of results returned by ListFeatureGroups .
	MaxResults *int32

	// A string that partially matches one or more FeatureGroup s names. Filters
	// FeatureGroup s by name.
	NameContains *string

	// A token to resume pagination of ListFeatureGroups results.
	NextToken *string

	// An OfflineStore status. Filters by OfflineStore status.
	OfflineStoreStatusEquals types.OfflineStoreStatusValue

	// The value on which the feature group list is sorted.
	SortBy types.FeatureGroupSortBy

	// The order in which feature groups are listed.
	SortOrder types.FeatureGroupSortOrder

	noSmithyDocumentSerde
}

type ListFeatureGroupsOutput struct {

	// A summary of feature groups.
	//
	// This member is required.
	FeatureGroupSummaries []types.FeatureGroupSummary

	// A token to resume pagination of ListFeatureGroups results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFeatureGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFeatureGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFeatureGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFeatureGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFeatureGroups(options.Region), middleware.Before); err != nil {
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

// ListFeatureGroupsPaginatorOptions is the paginator options for ListFeatureGroups
type ListFeatureGroupsPaginatorOptions struct {
	// The maximum number of results returned by ListFeatureGroups .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFeatureGroupsPaginator is a paginator for ListFeatureGroups
type ListFeatureGroupsPaginator struct {
	options   ListFeatureGroupsPaginatorOptions
	client    ListFeatureGroupsAPIClient
	params    *ListFeatureGroupsInput
	nextToken *string
	firstPage bool
}

// NewListFeatureGroupsPaginator returns a new ListFeatureGroupsPaginator
func NewListFeatureGroupsPaginator(client ListFeatureGroupsAPIClient, params *ListFeatureGroupsInput, optFns ...func(*ListFeatureGroupsPaginatorOptions)) *ListFeatureGroupsPaginator {
	if params == nil {
		params = &ListFeatureGroupsInput{}
	}

	options := ListFeatureGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFeatureGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFeatureGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFeatureGroups page.
func (p *ListFeatureGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFeatureGroupsOutput, error) {
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
	result, err := p.client.ListFeatureGroups(ctx, &params, optFns...)
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

// ListFeatureGroupsAPIClient is a client that implements the ListFeatureGroups
// operation.
type ListFeatureGroupsAPIClient interface {
	ListFeatureGroups(context.Context, *ListFeatureGroupsInput, ...func(*Options)) (*ListFeatureGroupsOutput, error)
}

var _ ListFeatureGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListFeatureGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFeatureGroups",
	}
}
