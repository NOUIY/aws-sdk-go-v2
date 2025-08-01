// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists [Lambda layers] and shows information about the latest version of each. Specify a [runtime identifier] to
// list only layers that indicate that they're compatible with that runtime.
// Specify a compatible architecture to include only layers that are compatible
// with that [instruction set architecture].
//
// [instruction set architecture]: https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html
// [runtime identifier]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
// [Lambda layers]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-layers.html
func (c *Client) ListLayers(ctx context.Context, params *ListLayersInput, optFns ...func(*Options)) (*ListLayersOutput, error) {
	if params == nil {
		params = &ListLayersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLayers", params, optFns, c.addOperationListLayersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLayersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLayersInput struct {

	// The compatible [instruction set architecture].
	//
	// [instruction set architecture]: https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html
	CompatibleArchitecture types.Architecture

	// A runtime identifier.
	//
	// The following list includes deprecated runtimes. For more information, see [Runtime use after deprecation].
	//
	// For a list of all currently supported runtimes, see [Supported runtimes].
	//
	// [Runtime use after deprecation]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-deprecation-levels
	// [Supported runtimes]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtimes-supported
	CompatibleRuntime types.Runtime

	// A pagination token returned by a previous call.
	Marker *string

	// The maximum number of layers to return.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListLayersOutput struct {

	// A list of function layers.
	Layers []types.LayersListItem

	// A pagination token returned when the response doesn't contain all layers.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLayersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLayers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLayers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLayers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLayers(options.Region), middleware.Before); err != nil {
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

// ListLayersPaginatorOptions is the paginator options for ListLayers
type ListLayersPaginatorOptions struct {
	// The maximum number of layers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLayersPaginator is a paginator for ListLayers
type ListLayersPaginator struct {
	options   ListLayersPaginatorOptions
	client    ListLayersAPIClient
	params    *ListLayersInput
	nextToken *string
	firstPage bool
}

// NewListLayersPaginator returns a new ListLayersPaginator
func NewListLayersPaginator(client ListLayersAPIClient, params *ListLayersInput, optFns ...func(*ListLayersPaginatorOptions)) *ListLayersPaginator {
	if params == nil {
		params = &ListLayersInput{}
	}

	options := ListLayersPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLayersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLayersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLayers page.
func (p *ListLayersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLayersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListLayers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListLayersAPIClient is a client that implements the ListLayers operation.
type ListLayersAPIClient interface {
	ListLayers(context.Context, *ListLayersInput, ...func(*Options)) (*ListLayersOutput, error)
}

var _ ListLayersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLayers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLayers",
	}
}
