// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of StreamInfo objects. Each object describes a stream. To
// retrieve only streams that satisfy a specific condition, you can specify a
// StreamNameCondition .
func (c *Client) ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) {
	if params == nil {
		params = &ListStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreams", params, optFns, c.addOperationListStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamsInput struct {

	// The maximum number of streams to return in the response. The default is 10,000.
	MaxResults *int32

	// If you specify this parameter, when the result of a ListStreams operation is
	// truncated, the call returns the NextToken in the response. To get another batch
	// of streams, provide this token in your next request.
	NextToken *string

	// Optional: Returns only streams that satisfy a specific condition. Currently,
	// you can specify only the prefix of a stream name as a condition.
	StreamNameCondition *types.StreamNameCondition

	noSmithyDocumentSerde
}

type ListStreamsOutput struct {

	// If the response is truncated, the call returns this element with a token. To
	// get the next batch of streams, use this token in your next request.
	NextToken *string

	// An array of StreamInfo objects.
	StreamInfoList []types.StreamInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreams{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStreams"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreams(options.Region), middleware.Before); err != nil {
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

// ListStreamsPaginatorOptions is the paginator options for ListStreams
type ListStreamsPaginatorOptions struct {
	// The maximum number of streams to return in the response. The default is 10,000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamsPaginator is a paginator for ListStreams
type ListStreamsPaginator struct {
	options   ListStreamsPaginatorOptions
	client    ListStreamsAPIClient
	params    *ListStreamsInput
	nextToken *string
	firstPage bool
}

// NewListStreamsPaginator returns a new ListStreamsPaginator
func NewListStreamsPaginator(client ListStreamsAPIClient, params *ListStreamsInput, optFns ...func(*ListStreamsPaginatorOptions)) *ListStreamsPaginator {
	if params == nil {
		params = &ListStreamsInput{}
	}

	options := ListStreamsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreams page.
func (p *ListStreamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamsOutput, error) {
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
	result, err := p.client.ListStreams(ctx, &params, optFns...)
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

// ListStreamsAPIClient is a client that implements the ListStreams operation.
type ListStreamsAPIClient interface {
	ListStreams(context.Context, *ListStreamsInput, ...func(*Options)) (*ListStreamsOutput, error)
}

var _ ListStreamsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStreams",
	}
}
