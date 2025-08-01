// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the available transformers. A transformer can take an EDI file as input
// and transform it into a JSON-or XML-formatted document. Alternatively, a
// transformer can take a JSON-or XML-formatted document as input and transform it
// into an EDI file.
func (c *Client) ListTransformers(ctx context.Context, params *ListTransformersInput, optFns ...func(*Options)) (*ListTransformersOutput, error) {
	if params == nil {
		params = &ListTransformersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransformers", params, optFns, c.addOperationListTransformersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransformersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransformersInput struct {

	// Specifies the number of items to return for the API response.
	MaxResults *int32

	// When additional results are obtained from the command, a NextToken parameter is
	// returned in the output. You can then pass the NextToken parameter in a
	// subsequent command to continue listing additional resources.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTransformersOutput struct {

	// Returns an array of one or more transformer objects.
	//
	// For each transformer, a TransformerSummary object is returned. The
	// TransformerSummary contains all the details for a specific transformer.
	//
	// This member is required.
	Transformers []types.TransformerSummary

	// When additional results are obtained from the command, a NextToken parameter is
	// returned in the output. You can then pass the NextToken parameter in a
	// subsequent command to continue listing additional resources.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTransformersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListTransformers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTransformers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTransformers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTransformers(options.Region), middleware.Before); err != nil {
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

// ListTransformersPaginatorOptions is the paginator options for ListTransformers
type ListTransformersPaginatorOptions struct {
	// Specifies the number of items to return for the API response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTransformersPaginator is a paginator for ListTransformers
type ListTransformersPaginator struct {
	options   ListTransformersPaginatorOptions
	client    ListTransformersAPIClient
	params    *ListTransformersInput
	nextToken *string
	firstPage bool
}

// NewListTransformersPaginator returns a new ListTransformersPaginator
func NewListTransformersPaginator(client ListTransformersAPIClient, params *ListTransformersInput, optFns ...func(*ListTransformersPaginatorOptions)) *ListTransformersPaginator {
	if params == nil {
		params = &ListTransformersInput{}
	}

	options := ListTransformersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTransformersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTransformersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTransformers page.
func (p *ListTransformersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTransformersOutput, error) {
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
	result, err := p.client.ListTransformers(ctx, &params, optFns...)
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

// ListTransformersAPIClient is a client that implements the ListTransformers
// operation.
type ListTransformersAPIClient interface {
	ListTransformers(context.Context, *ListTransformersInput, ...func(*Options)) (*ListTransformersOutput, error)
}

var _ ListTransformersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTransformers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTransformers",
	}
}
