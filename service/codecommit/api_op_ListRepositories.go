// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about one or more repositories.
func (c *Client) ListRepositories(ctx context.Context, params *ListRepositoriesInput, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	if params == nil {
		params = &ListRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositories", params, optFns, c.addOperationListRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a list repositories operation.
type ListRepositoriesInput struct {

	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the client
	// sends the token back to CodeCommit, another page of 1,000 records is retrieved.
	NextToken *string

	// The order in which to sort the results of a list repositories operation.
	Order types.OrderEnum

	// The criteria used to sort the results of a list repositories operation.
	SortBy types.SortByEnum

	noSmithyDocumentSerde
}

// Represents the output of a list repositories operation.
type ListRepositoriesOutput struct {

	// An enumeration token that allows the operation to batch the results of the
	// operation. Batch sizes are 1,000 for list repository operations. When the client
	// sends the token back to CodeCommit, another page of 1,000 records is retrieved.
	NextToken *string

	// Lists the repositories called by the list repositories operation.
	Repositories []types.RepositoryNameIdPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRepositories"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositories(options.Region), middleware.Before); err != nil {
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

// ListRepositoriesPaginatorOptions is the paginator options for ListRepositories
type ListRepositoriesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositoriesPaginator is a paginator for ListRepositories
type ListRepositoriesPaginator struct {
	options   ListRepositoriesPaginatorOptions
	client    ListRepositoriesAPIClient
	params    *ListRepositoriesInput
	nextToken *string
	firstPage bool
}

// NewListRepositoriesPaginator returns a new ListRepositoriesPaginator
func NewListRepositoriesPaginator(client ListRepositoriesAPIClient, params *ListRepositoriesInput, optFns ...func(*ListRepositoriesPaginatorOptions)) *ListRepositoriesPaginator {
	if params == nil {
		params = &ListRepositoriesInput{}
	}

	options := ListRepositoriesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRepositoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRepositories page.
func (p *ListRepositoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListRepositories(ctx, &params, optFns...)
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

// ListRepositoriesAPIClient is a client that implements the ListRepositories
// operation.
type ListRepositoriesAPIClient interface {
	ListRepositories(context.Context, *ListRepositoriesInput, ...func(*Options)) (*ListRepositoriesOutput, error)
}

var _ ListRepositoriesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRepositories",
	}
}
