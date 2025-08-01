// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches the flow modules in an Amazon Connect instance, with optional
// filtering.
func (c *Client) SearchContactFlowModules(ctx context.Context, params *SearchContactFlowModulesInput, optFns ...func(*Options)) (*SearchContactFlowModulesOutput, error) {
	if params == nil {
		params = &SearchContactFlowModulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchContactFlowModules", params, optFns, c.addOperationSearchContactFlowModulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchContactFlowModulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchContactFlowModulesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID in
	// the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return flow modules.
	//
	// The name and description fields support "contains" queries with a minimum of 2
	// characters and a maximum of 25 characters. Any queries with character lengths
	// outside of this range will result in invalid results.
	SearchCriteria *types.ContactFlowModuleSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.ContactFlowModuleSearchFilter

	noSmithyDocumentSerde
}

type SearchContactFlowModulesOutput struct {

	// The total number of flows which matched your search query.
	ApproximateTotalCount *int64

	// The search criteria to be used to return flow modules.
	ContactFlowModules []types.ContactFlowModule

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchContactFlowModulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchContactFlowModules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchContactFlowModules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchContactFlowModules"); err != nil {
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
	if err = addOpSearchContactFlowModulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchContactFlowModules(options.Region), middleware.Before); err != nil {
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

// SearchContactFlowModulesPaginatorOptions is the paginator options for
// SearchContactFlowModules
type SearchContactFlowModulesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchContactFlowModulesPaginator is a paginator for SearchContactFlowModules
type SearchContactFlowModulesPaginator struct {
	options   SearchContactFlowModulesPaginatorOptions
	client    SearchContactFlowModulesAPIClient
	params    *SearchContactFlowModulesInput
	nextToken *string
	firstPage bool
}

// NewSearchContactFlowModulesPaginator returns a new
// SearchContactFlowModulesPaginator
func NewSearchContactFlowModulesPaginator(client SearchContactFlowModulesAPIClient, params *SearchContactFlowModulesInput, optFns ...func(*SearchContactFlowModulesPaginatorOptions)) *SearchContactFlowModulesPaginator {
	if params == nil {
		params = &SearchContactFlowModulesInput{}
	}

	options := SearchContactFlowModulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchContactFlowModulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchContactFlowModulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchContactFlowModules page.
func (p *SearchContactFlowModulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchContactFlowModulesOutput, error) {
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
	result, err := p.client.SearchContactFlowModules(ctx, &params, optFns...)
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

// SearchContactFlowModulesAPIClient is a client that implements the
// SearchContactFlowModules operation.
type SearchContactFlowModulesAPIClient interface {
	SearchContactFlowModules(context.Context, *SearchContactFlowModulesInput, ...func(*Options)) (*SearchContactFlowModulesOutput, error)
}

var _ SearchContactFlowModulesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchContactFlowModules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchContactFlowModules",
	}
}
