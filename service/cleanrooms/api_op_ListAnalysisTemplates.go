// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists analysis templates that the caller owns.
func (c *Client) ListAnalysisTemplates(ctx context.Context, params *ListAnalysisTemplatesInput, optFns ...func(*Options)) (*ListAnalysisTemplatesOutput, error) {
	if params == nil {
		params = &ListAnalysisTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalysisTemplates", params, optFns, c.addOperationListAnalysisTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalysisTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnalysisTemplatesInput struct {

	// The identifier for a membership resource.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum number of results that are returned for an API request call. The
	// service chooses a default number if you don't set one. The service might return
	// a `nextToken` even if the `maxResults` value has not been met.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnalysisTemplatesOutput struct {

	// Lists analysis template metadata.
	//
	// This member is required.
	AnalysisTemplateSummaries []types.AnalysisTemplateSummary

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnalysisTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalysisTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalysisTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnalysisTemplates"); err != nil {
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
	if err = addOpListAnalysisTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnalysisTemplates(options.Region), middleware.Before); err != nil {
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

// ListAnalysisTemplatesPaginatorOptions is the paginator options for
// ListAnalysisTemplates
type ListAnalysisTemplatesPaginatorOptions struct {
	// The maximum number of results that are returned for an API request call. The
	// service chooses a default number if you don't set one. The service might return
	// a `nextToken` even if the `maxResults` value has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnalysisTemplatesPaginator is a paginator for ListAnalysisTemplates
type ListAnalysisTemplatesPaginator struct {
	options   ListAnalysisTemplatesPaginatorOptions
	client    ListAnalysisTemplatesAPIClient
	params    *ListAnalysisTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListAnalysisTemplatesPaginator returns a new ListAnalysisTemplatesPaginator
func NewListAnalysisTemplatesPaginator(client ListAnalysisTemplatesAPIClient, params *ListAnalysisTemplatesInput, optFns ...func(*ListAnalysisTemplatesPaginatorOptions)) *ListAnalysisTemplatesPaginator {
	if params == nil {
		params = &ListAnalysisTemplatesInput{}
	}

	options := ListAnalysisTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnalysisTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnalysisTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnalysisTemplates page.
func (p *ListAnalysisTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnalysisTemplatesOutput, error) {
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
	result, err := p.client.ListAnalysisTemplates(ctx, &params, optFns...)
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

// ListAnalysisTemplatesAPIClient is a client that implements the
// ListAnalysisTemplates operation.
type ListAnalysisTemplatesAPIClient interface {
	ListAnalysisTemplates(context.Context, *ListAnalysisTemplatesInput, ...func(*Options)) (*ListAnalysisTemplatesOutput, error)
}

var _ ListAnalysisTemplatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAnalysisTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnalysisTemplates",
	}
}
