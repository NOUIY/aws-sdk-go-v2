// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available session logger resources.
func (c *Client) ListSessionLoggers(ctx context.Context, params *ListSessionLoggersInput, optFns ...func(*Options)) (*ListSessionLoggersOutput, error) {
	if params == nil {
		params = &ListSessionLoggersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSessionLoggers", params, optFns, c.addOperationListSessionLoggersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSessionLoggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSessionLoggersInput struct {

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSessionLoggersOutput struct {

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// The list of session loggers, including summaries of their details.
	SessionLoggers []types.SessionLoggerSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSessionLoggersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSessionLoggers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSessionLoggers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSessionLoggers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSessionLoggers(options.Region), middleware.Before); err != nil {
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

// ListSessionLoggersPaginatorOptions is the paginator options for
// ListSessionLoggers
type ListSessionLoggersPaginatorOptions struct {
	// The maximum number of results to be included in the next page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSessionLoggersPaginator is a paginator for ListSessionLoggers
type ListSessionLoggersPaginator struct {
	options   ListSessionLoggersPaginatorOptions
	client    ListSessionLoggersAPIClient
	params    *ListSessionLoggersInput
	nextToken *string
	firstPage bool
}

// NewListSessionLoggersPaginator returns a new ListSessionLoggersPaginator
func NewListSessionLoggersPaginator(client ListSessionLoggersAPIClient, params *ListSessionLoggersInput, optFns ...func(*ListSessionLoggersPaginatorOptions)) *ListSessionLoggersPaginator {
	if params == nil {
		params = &ListSessionLoggersInput{}
	}

	options := ListSessionLoggersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSessionLoggersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSessionLoggersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSessionLoggers page.
func (p *ListSessionLoggersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSessionLoggersOutput, error) {
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
	result, err := p.client.ListSessionLoggers(ctx, &params, optFns...)
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

// ListSessionLoggersAPIClient is a client that implements the ListSessionLoggers
// operation.
type ListSessionLoggersAPIClient interface {
	ListSessionLoggers(context.Context, *ListSessionLoggersInput, ...func(*Options)) (*ListSessionLoggersOutput, error)
}

var _ ListSessionLoggersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSessionLoggers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSessionLoggers",
	}
}
