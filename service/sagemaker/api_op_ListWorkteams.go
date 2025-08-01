// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of private work teams that you have defined in a region. The list
// may be empty if no work team satisfies the filter specified in the NameContains
// parameter.
func (c *Client) ListWorkteams(ctx context.Context, params *ListWorkteamsInput, optFns ...func(*Options)) (*ListWorkteamsOutput, error) {
	if params == nil {
		params = &ListWorkteamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkteams", params, optFns, c.addOperationListWorkteamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkteamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkteamsInput struct {

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int32

	// A string in the work team's name. This filter returns only work teams whose
	// name contains the specified string.
	NameContains *string

	// If the result of the previous ListWorkteams request was truncated, the response
	// includes a NextToken . To retrieve the next set of labeling jobs, use the token
	// in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime .
	SortBy types.ListWorkteamsSortByOptions

	// The sort order for results. The default is Ascending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListWorkteamsOutput struct {

	// An array of Workteam objects, each describing a work team.
	//
	// This member is required.
	Workteams []types.Workteam

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkteamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkteams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkteams{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkteams"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkteams(options.Region), middleware.Before); err != nil {
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

// ListWorkteamsPaginatorOptions is the paginator options for ListWorkteams
type ListWorkteamsPaginatorOptions struct {
	// The maximum number of work teams to return in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkteamsPaginator is a paginator for ListWorkteams
type ListWorkteamsPaginator struct {
	options   ListWorkteamsPaginatorOptions
	client    ListWorkteamsAPIClient
	params    *ListWorkteamsInput
	nextToken *string
	firstPage bool
}

// NewListWorkteamsPaginator returns a new ListWorkteamsPaginator
func NewListWorkteamsPaginator(client ListWorkteamsAPIClient, params *ListWorkteamsInput, optFns ...func(*ListWorkteamsPaginatorOptions)) *ListWorkteamsPaginator {
	if params == nil {
		params = &ListWorkteamsInput{}
	}

	options := ListWorkteamsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkteamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkteamsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkteams page.
func (p *ListWorkteamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkteamsOutput, error) {
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
	result, err := p.client.ListWorkteams(ctx, &params, optFns...)
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

// ListWorkteamsAPIClient is a client that implements the ListWorkteams operation.
type ListWorkteamsAPIClient interface {
	ListWorkteams(context.Context, *ListWorkteamsInput, ...func(*Options)) (*ListWorkteamsOutput, error)
}

var _ ListWorkteamsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkteams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkteams",
	}
}
