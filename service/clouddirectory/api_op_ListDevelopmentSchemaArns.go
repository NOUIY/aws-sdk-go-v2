// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves each Amazon Resource Name (ARN) of schemas in the development state.
func (c *Client) ListDevelopmentSchemaArns(ctx context.Context, params *ListDevelopmentSchemaArnsInput, optFns ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error) {
	if params == nil {
		params = &ListDevelopmentSchemaArnsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevelopmentSchemaArns", params, optFns, c.addOperationListDevelopmentSchemaArnsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevelopmentSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevelopmentSchemaArnsInput struct {

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDevelopmentSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of retrieved development schemas.
	SchemaArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevelopmentSchemaArnsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevelopmentSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevelopmentSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevelopmentSchemaArns"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevelopmentSchemaArns(options.Region), middleware.Before); err != nil {
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

// ListDevelopmentSchemaArnsPaginatorOptions is the paginator options for
// ListDevelopmentSchemaArns
type ListDevelopmentSchemaArnsPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevelopmentSchemaArnsPaginator is a paginator for ListDevelopmentSchemaArns
type ListDevelopmentSchemaArnsPaginator struct {
	options   ListDevelopmentSchemaArnsPaginatorOptions
	client    ListDevelopmentSchemaArnsAPIClient
	params    *ListDevelopmentSchemaArnsInput
	nextToken *string
	firstPage bool
}

// NewListDevelopmentSchemaArnsPaginator returns a new
// ListDevelopmentSchemaArnsPaginator
func NewListDevelopmentSchemaArnsPaginator(client ListDevelopmentSchemaArnsAPIClient, params *ListDevelopmentSchemaArnsInput, optFns ...func(*ListDevelopmentSchemaArnsPaginatorOptions)) *ListDevelopmentSchemaArnsPaginator {
	if params == nil {
		params = &ListDevelopmentSchemaArnsInput{}
	}

	options := ListDevelopmentSchemaArnsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevelopmentSchemaArnsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevelopmentSchemaArnsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevelopmentSchemaArns page.
func (p *ListDevelopmentSchemaArnsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error) {
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
	result, err := p.client.ListDevelopmentSchemaArns(ctx, &params, optFns...)
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

// ListDevelopmentSchemaArnsAPIClient is a client that implements the
// ListDevelopmentSchemaArns operation.
type ListDevelopmentSchemaArnsAPIClient interface {
	ListDevelopmentSchemaArns(context.Context, *ListDevelopmentSchemaArnsInput, ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error)
}

var _ ListDevelopmentSchemaArnsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDevelopmentSchemaArns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevelopmentSchemaArns",
	}
}
