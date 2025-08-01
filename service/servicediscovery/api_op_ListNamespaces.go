// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists summary information about the namespaces that were created by the current
// Amazon Web Services account.
func (c *Client) ListNamespaces(ctx context.Context, params *ListNamespacesInput, optFns ...func(*Options)) (*ListNamespacesOutput, error) {
	if params == nil {
		params = &ListNamespacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNamespaces", params, optFns, c.addOperationListNamespacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNamespacesInput struct {

	// A complex type that contains specifications for the namespaces that you want to
	// list.
	//
	// If you specify more than one filter, a namespace must match all filters to be
	// returned by ListNamespaces .
	Filters []types.NamespaceFilter

	// The maximum number of namespaces that you want Cloud Map to return in the
	// response to a ListNamespaces request. If you don't specify a value for
	// MaxResults , Cloud Map returns up to 100 namespaces.
	MaxResults *int32

	// For the first ListNamespaces request, omit this value.
	//
	// If the response contains NextToken , submit another ListNamespaces request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// Cloud Map gets MaxResults namespaces and then filters them based on the
	// specified criteria. It's possible that no namespaces in the first MaxResults
	// namespaces matched the specified criteria but that subsequent groups of
	// MaxResults namespaces do contain namespaces that match the criteria.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNamespacesOutput struct {

	// An array that contains one NamespaceSummary object for each namespace that
	// matches the specified filter criteria.
	Namespaces []types.NamespaceSummary

	// If the response contains NextToken , submit another ListNamespaces request to
	// get the next group of results. Specify the value of NextToken from the previous
	// response in the next request.
	//
	// Cloud Map gets MaxResults namespaces and then filters them based on the
	// specified criteria. It's possible that no namespaces in the first MaxResults
	// namespaces matched the specified criteria but that subsequent groups of
	// MaxResults namespaces do contain namespaces that match the criteria.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNamespacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNamespaces"); err != nil {
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
	if err = addOpListNamespacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNamespaces(options.Region), middleware.Before); err != nil {
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

// ListNamespacesPaginatorOptions is the paginator options for ListNamespaces
type ListNamespacesPaginatorOptions struct {
	// The maximum number of namespaces that you want Cloud Map to return in the
	// response to a ListNamespaces request. If you don't specify a value for
	// MaxResults , Cloud Map returns up to 100 namespaces.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNamespacesPaginator is a paginator for ListNamespaces
type ListNamespacesPaginator struct {
	options   ListNamespacesPaginatorOptions
	client    ListNamespacesAPIClient
	params    *ListNamespacesInput
	nextToken *string
	firstPage bool
}

// NewListNamespacesPaginator returns a new ListNamespacesPaginator
func NewListNamespacesPaginator(client ListNamespacesAPIClient, params *ListNamespacesInput, optFns ...func(*ListNamespacesPaginatorOptions)) *ListNamespacesPaginator {
	if params == nil {
		params = &ListNamespacesInput{}
	}

	options := ListNamespacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNamespacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNamespacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNamespaces page.
func (p *ListNamespacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNamespacesOutput, error) {
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
	result, err := p.client.ListNamespaces(ctx, &params, optFns...)
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

// ListNamespacesAPIClient is a client that implements the ListNamespaces
// operation.
type ListNamespacesAPIClient interface {
	ListNamespaces(context.Context, *ListNamespacesInput, ...func(*Options)) (*ListNamespacesOutput, error)
}

var _ ListNamespacesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListNamespaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNamespaces",
	}
}
