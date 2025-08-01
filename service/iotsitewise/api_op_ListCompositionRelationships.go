// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of composition relationships for an asset model of
// type COMPONENT_MODEL .
func (c *Client) ListCompositionRelationships(ctx context.Context, params *ListCompositionRelationshipsInput, optFns ...func(*Options)) (*ListCompositionRelationshipsOutput, error) {
	if params == nil {
		params = &ListCompositionRelationshipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompositionRelationships", params, optFns, c.addOperationListCompositionRelationshipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCompositionRelationshipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCompositionRelationshipsInput struct {

	// The ID of the asset model. This can be either the actual ID in UUID format, or
	// else externalId: followed by the external ID, if it has one. For more
	// information, see [Referencing objects with external IDs]in the IoT SiteWise User Guide.
	//
	// [Referencing objects with external IDs]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references
	//
	// This member is required.
	AssetModelId *string

	// The maximum number of results to return for each paginated request.
	//
	// Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCompositionRelationshipsOutput struct {

	// A list that summarizes each composition relationship.
	//
	// This member is required.
	CompositionRelationshipSummaries []types.CompositionRelationshipSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCompositionRelationshipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCompositionRelationships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCompositionRelationships{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCompositionRelationships"); err != nil {
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
	if err = addEndpointPrefix_opListCompositionRelationshipsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListCompositionRelationshipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCompositionRelationships(options.Region), middleware.Before); err != nil {
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

// ListCompositionRelationshipsPaginatorOptions is the paginator options for
// ListCompositionRelationships
type ListCompositionRelationshipsPaginatorOptions struct {
	// The maximum number of results to return for each paginated request.
	//
	// Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCompositionRelationshipsPaginator is a paginator for
// ListCompositionRelationships
type ListCompositionRelationshipsPaginator struct {
	options   ListCompositionRelationshipsPaginatorOptions
	client    ListCompositionRelationshipsAPIClient
	params    *ListCompositionRelationshipsInput
	nextToken *string
	firstPage bool
}

// NewListCompositionRelationshipsPaginator returns a new
// ListCompositionRelationshipsPaginator
func NewListCompositionRelationshipsPaginator(client ListCompositionRelationshipsAPIClient, params *ListCompositionRelationshipsInput, optFns ...func(*ListCompositionRelationshipsPaginatorOptions)) *ListCompositionRelationshipsPaginator {
	if params == nil {
		params = &ListCompositionRelationshipsInput{}
	}

	options := ListCompositionRelationshipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCompositionRelationshipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCompositionRelationshipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCompositionRelationships page.
func (p *ListCompositionRelationshipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCompositionRelationshipsOutput, error) {
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
	result, err := p.client.ListCompositionRelationships(ctx, &params, optFns...)
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

type endpointPrefix_opListCompositionRelationshipsMiddleware struct {
}

func (*endpointPrefix_opListCompositionRelationshipsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListCompositionRelationshipsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListCompositionRelationshipsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListCompositionRelationshipsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListCompositionRelationshipsAPIClient is a client that implements the
// ListCompositionRelationships operation.
type ListCompositionRelationshipsAPIClient interface {
	ListCompositionRelationships(context.Context, *ListCompositionRelationshipsInput, ...func(*Options)) (*ListCompositionRelationshipsOutput, error)
}

var _ ListCompositionRelationshipsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCompositionRelationships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCompositionRelationships",
	}
}
