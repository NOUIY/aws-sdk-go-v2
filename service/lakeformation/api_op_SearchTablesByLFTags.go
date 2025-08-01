// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation allows a search on TABLE resources by LFTag s. This will be used
// by admins who want to grant user permissions on certain LF-tags. Before making a
// grant, the admin can use SearchTablesByLFTags to find all resources where the
// given LFTag s are valid to verify whether the returned resources can be shared.
func (c *Client) SearchTablesByLFTags(ctx context.Context, params *SearchTablesByLFTagsInput, optFns ...func(*Options)) (*SearchTablesByLFTagsOutput, error) {
	if params == nil {
		params = &SearchTablesByLFTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTablesByLFTags", params, optFns, c.addOperationSearchTablesByLFTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTablesByLFTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTablesByLFTagsInput struct {

	// A list of conditions ( LFTag structures) to search for in table resources.
	//
	// This member is required.
	Expression []types.LFTag

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchTablesByLFTagsOutput struct {

	// A continuation token, present if the current list segment is not the last. On
	// the first run, if you include a not null (a value) token you can get empty
	// pages.
	NextToken *string

	// A list of tables that meet the LF-tag conditions.
	TableList []types.TaggedTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTablesByLFTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchTablesByLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchTablesByLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchTablesByLFTags"); err != nil {
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
	if err = addOpSearchTablesByLFTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTablesByLFTags(options.Region), middleware.Before); err != nil {
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

// SearchTablesByLFTagsPaginatorOptions is the paginator options for
// SearchTablesByLFTags
type SearchTablesByLFTagsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchTablesByLFTagsPaginator is a paginator for SearchTablesByLFTags
type SearchTablesByLFTagsPaginator struct {
	options   SearchTablesByLFTagsPaginatorOptions
	client    SearchTablesByLFTagsAPIClient
	params    *SearchTablesByLFTagsInput
	nextToken *string
	firstPage bool
}

// NewSearchTablesByLFTagsPaginator returns a new SearchTablesByLFTagsPaginator
func NewSearchTablesByLFTagsPaginator(client SearchTablesByLFTagsAPIClient, params *SearchTablesByLFTagsInput, optFns ...func(*SearchTablesByLFTagsPaginatorOptions)) *SearchTablesByLFTagsPaginator {
	if params == nil {
		params = &SearchTablesByLFTagsInput{}
	}

	options := SearchTablesByLFTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchTablesByLFTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchTablesByLFTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchTablesByLFTags page.
func (p *SearchTablesByLFTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchTablesByLFTagsOutput, error) {
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
	result, err := p.client.SearchTablesByLFTags(ctx, &params, optFns...)
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

// SearchTablesByLFTagsAPIClient is a client that implements the
// SearchTablesByLFTags operation.
type SearchTablesByLFTagsAPIClient interface {
	SearchTablesByLFTags(context.Context, *SearchTablesByLFTagsInput, ...func(*Options)) (*SearchTablesByLFTagsOutput, error)
}

var _ SearchTablesByLFTagsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchTablesByLFTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchTablesByLFTags",
	}
}
