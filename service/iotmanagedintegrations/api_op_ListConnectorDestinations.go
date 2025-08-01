// Code generated by smithy-go-codegen DO NOT EDIT.

package iotmanagedintegrations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all connector destinations, with optional filtering by cloud connector ID.
func (c *Client) ListConnectorDestinations(ctx context.Context, params *ListConnectorDestinationsInput, optFns ...func(*Options)) (*ListConnectorDestinationsOutput, error) {
	if params == nil {
		params = &ListConnectorDestinationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectorDestinations", params, optFns, c.addOperationListConnectorDestinationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorDestinationsInput struct {

	// The identifier of the cloud connector to filter connector destinations by.
	CloudConnectorId *string

	// The maximum number of connector destinations to return in a single response.
	MaxResults *int32

	// A token used for pagination of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConnectorDestinationsOutput struct {

	// The list of connector destinations that match the specified criteria.
	ConnectorDestinationList []types.ConnectorDestinationSummary

	// A token used for pagination of results when there are more connector
	// destinations than can be returned in a single response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConnectorDestinationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConnectorDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectorDestinations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConnectorDestinations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectorDestinations(options.Region), middleware.Before); err != nil {
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

// ListConnectorDestinationsPaginatorOptions is the paginator options for
// ListConnectorDestinations
type ListConnectorDestinationsPaginatorOptions struct {
	// The maximum number of connector destinations to return in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectorDestinationsPaginator is a paginator for ListConnectorDestinations
type ListConnectorDestinationsPaginator struct {
	options   ListConnectorDestinationsPaginatorOptions
	client    ListConnectorDestinationsAPIClient
	params    *ListConnectorDestinationsInput
	nextToken *string
	firstPage bool
}

// NewListConnectorDestinationsPaginator returns a new
// ListConnectorDestinationsPaginator
func NewListConnectorDestinationsPaginator(client ListConnectorDestinationsAPIClient, params *ListConnectorDestinationsInput, optFns ...func(*ListConnectorDestinationsPaginatorOptions)) *ListConnectorDestinationsPaginator {
	if params == nil {
		params = &ListConnectorDestinationsInput{}
	}

	options := ListConnectorDestinationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConnectorDestinationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectorDestinationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConnectorDestinations page.
func (p *ListConnectorDestinationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectorDestinationsOutput, error) {
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
	result, err := p.client.ListConnectorDestinations(ctx, &params, optFns...)
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

// ListConnectorDestinationsAPIClient is a client that implements the
// ListConnectorDestinations operation.
type ListConnectorDestinationsAPIClient interface {
	ListConnectorDestinations(context.Context, *ListConnectorDestinationsInput, ...func(*Options)) (*ListConnectorDestinationsOutput, error)
}

var _ ListConnectorDestinationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListConnectorDestinations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConnectorDestinations",
	}
}
