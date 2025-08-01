// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the major version families of each published schema. If a major version
// ARN is provided as SchemaArn , the minor version revisions in that family are
// listed instead.
func (c *Client) ListPublishedSchemaArns(ctx context.Context, params *ListPublishedSchemaArnsInput, optFns ...func(*Options)) (*ListPublishedSchemaArnsOutput, error) {
	if params == nil {
		params = &ListPublishedSchemaArnsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublishedSchemaArns", params, optFns, c.addOperationListPublishedSchemaArnsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublishedSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublishedSchemaArnsInput struct {

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The response for ListPublishedSchemaArns when this parameter is used will list
	// all minor version ARNs for a major version.
	SchemaArn *string

	noSmithyDocumentSerde
}

type ListPublishedSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of published schemas.
	SchemaArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPublishedSchemaArnsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPublishedSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublishedSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPublishedSchemaArns"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublishedSchemaArns(options.Region), middleware.Before); err != nil {
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

// ListPublishedSchemaArnsPaginatorOptions is the paginator options for
// ListPublishedSchemaArns
type ListPublishedSchemaArnsPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPublishedSchemaArnsPaginator is a paginator for ListPublishedSchemaArns
type ListPublishedSchemaArnsPaginator struct {
	options   ListPublishedSchemaArnsPaginatorOptions
	client    ListPublishedSchemaArnsAPIClient
	params    *ListPublishedSchemaArnsInput
	nextToken *string
	firstPage bool
}

// NewListPublishedSchemaArnsPaginator returns a new
// ListPublishedSchemaArnsPaginator
func NewListPublishedSchemaArnsPaginator(client ListPublishedSchemaArnsAPIClient, params *ListPublishedSchemaArnsInput, optFns ...func(*ListPublishedSchemaArnsPaginatorOptions)) *ListPublishedSchemaArnsPaginator {
	if params == nil {
		params = &ListPublishedSchemaArnsInput{}
	}

	options := ListPublishedSchemaArnsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPublishedSchemaArnsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPublishedSchemaArnsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPublishedSchemaArns page.
func (p *ListPublishedSchemaArnsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPublishedSchemaArnsOutput, error) {
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
	result, err := p.client.ListPublishedSchemaArns(ctx, &params, optFns...)
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

// ListPublishedSchemaArnsAPIClient is a client that implements the
// ListPublishedSchemaArns operation.
type ListPublishedSchemaArnsAPIClient interface {
	ListPublishedSchemaArns(context.Context, *ListPublishedSchemaArnsInput, ...func(*Options)) (*ListPublishedSchemaArnsOutput, error)
}

var _ ListPublishedSchemaArnsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPublishedSchemaArns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPublishedSchemaArns",
	}
}
