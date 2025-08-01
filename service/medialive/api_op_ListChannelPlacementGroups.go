// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve the list of ChannelPlacementGroups in the specified Cluster.
func (c *Client) ListChannelPlacementGroups(ctx context.Context, params *ListChannelPlacementGroupsInput, optFns ...func(*Options)) (*ListChannelPlacementGroupsOutput, error) {
	if params == nil {
		params = &ListChannelPlacementGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelPlacementGroups", params, optFns, c.addOperationListChannelPlacementGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelPlacementGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListChannelPlacementGroupsRequest
type ListChannelPlacementGroupsInput struct {

	// The ID of the cluster
	//
	// This member is required.
	ClusterId *string

	// The maximum number of items to return.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListChannelPlacementGroupsResponse
type ListChannelPlacementGroupsOutput struct {

	// An array of ChannelPlacementGroups that exist in the Cluster.
	ChannelPlacementGroups []types.DescribeChannelPlacementGroupSummary

	// Token for the next result.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelPlacementGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelPlacementGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelPlacementGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListChannelPlacementGroups"); err != nil {
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
	if err = addOpListChannelPlacementGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelPlacementGroups(options.Region), middleware.Before); err != nil {
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

// ListChannelPlacementGroupsPaginatorOptions is the paginator options for
// ListChannelPlacementGroups
type ListChannelPlacementGroupsPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelPlacementGroupsPaginator is a paginator for
// ListChannelPlacementGroups
type ListChannelPlacementGroupsPaginator struct {
	options   ListChannelPlacementGroupsPaginatorOptions
	client    ListChannelPlacementGroupsAPIClient
	params    *ListChannelPlacementGroupsInput
	nextToken *string
	firstPage bool
}

// NewListChannelPlacementGroupsPaginator returns a new
// ListChannelPlacementGroupsPaginator
func NewListChannelPlacementGroupsPaginator(client ListChannelPlacementGroupsAPIClient, params *ListChannelPlacementGroupsInput, optFns ...func(*ListChannelPlacementGroupsPaginatorOptions)) *ListChannelPlacementGroupsPaginator {
	if params == nil {
		params = &ListChannelPlacementGroupsInput{}
	}

	options := ListChannelPlacementGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelPlacementGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelPlacementGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelPlacementGroups page.
func (p *ListChannelPlacementGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelPlacementGroupsOutput, error) {
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
	result, err := p.client.ListChannelPlacementGroups(ctx, &params, optFns...)
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

// ListChannelPlacementGroupsAPIClient is a client that implements the
// ListChannelPlacementGroups operation.
type ListChannelPlacementGroupsAPIClient interface {
	ListChannelPlacementGroups(context.Context, *ListChannelPlacementGroupsInput, ...func(*Options)) (*ListChannelPlacementGroupsOutput, error)
}

var _ ListChannelPlacementGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListChannelPlacementGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListChannelPlacementGroups",
	}
}
