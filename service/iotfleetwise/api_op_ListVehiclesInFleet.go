// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Retrieves a list of summaries of all vehicles associated with a fleet.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListVehiclesInFleet(ctx context.Context, params *ListVehiclesInFleetInput, optFns ...func(*Options)) (*ListVehiclesInFleetOutput, error) {
	if params == nil {
		params = &ListVehiclesInFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVehiclesInFleet", params, optFns, c.addOperationListVehiclesInFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVehiclesInFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVehiclesInFleetInput struct {

	//  The ID of a fleet.
	//
	// This member is required.
	FleetId *string

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results.
	//
	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next set of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVehiclesInFleetOutput struct {

	//  The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	//  A list of vehicles associated with the fleet.
	Vehicles []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVehiclesInFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVehiclesInFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVehiclesInFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVehiclesInFleet"); err != nil {
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
	if err = addOpListVehiclesInFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVehiclesInFleet(options.Region), middleware.Before); err != nil {
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

// ListVehiclesInFleetPaginatorOptions is the paginator options for
// ListVehiclesInFleet
type ListVehiclesInFleetPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVehiclesInFleetPaginator is a paginator for ListVehiclesInFleet
type ListVehiclesInFleetPaginator struct {
	options   ListVehiclesInFleetPaginatorOptions
	client    ListVehiclesInFleetAPIClient
	params    *ListVehiclesInFleetInput
	nextToken *string
	firstPage bool
}

// NewListVehiclesInFleetPaginator returns a new ListVehiclesInFleetPaginator
func NewListVehiclesInFleetPaginator(client ListVehiclesInFleetAPIClient, params *ListVehiclesInFleetInput, optFns ...func(*ListVehiclesInFleetPaginatorOptions)) *ListVehiclesInFleetPaginator {
	if params == nil {
		params = &ListVehiclesInFleetInput{}
	}

	options := ListVehiclesInFleetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVehiclesInFleetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVehiclesInFleetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVehiclesInFleet page.
func (p *ListVehiclesInFleetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVehiclesInFleetOutput, error) {
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
	result, err := p.client.ListVehiclesInFleet(ctx, &params, optFns...)
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

// ListVehiclesInFleetAPIClient is a client that implements the
// ListVehiclesInFleet operation.
type ListVehiclesInFleetAPIClient interface {
	ListVehiclesInFleet(context.Context, *ListVehiclesInFleetInput, ...func(*Options)) (*ListVehiclesInFleetOutput, error)
}

var _ ListVehiclesInFleetAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListVehiclesInFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVehiclesInFleet",
	}
}
