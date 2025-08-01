// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of models you've imported. You can filter the results to return
// based on one or more criteria. For more information, see [Import a customized model]in the [Amazon Bedrock User Guide].
//
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func (c *Client) ListImportedModels(ctx context.Context, params *ListImportedModelsInput, optFns ...func(*Options)) (*ListImportedModelsOutput, error) {
	if params == nil {
		params = &ListImportedModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportedModels", params, optFns, c.addOperationListImportedModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportedModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportedModelsInput struct {

	// Return imported models that were created after the specified time.
	CreationTimeAfter *time.Time

	// Return imported models that created before the specified time.
	CreationTimeBefore *time.Time

	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxResults *int32

	// Return imported models only if the model name contains these characters.
	NameContains *string

	// If the total number of results is greater than the maxResults value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	// The field to sort by in the returned list of imported models.
	SortBy types.SortModelsBy

	// Specifies whetehr to sort the results in ascending or descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListImportedModelsOutput struct {

	// Model summaries.
	ModelSummaries []types.ImportedModelSummary

	// If the total number of results is greater than the maxResults value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportedModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImportedModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImportedModels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImportedModels"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportedModels(options.Region), middleware.Before); err != nil {
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

// ListImportedModelsPaginatorOptions is the paginator options for
// ListImportedModels
type ListImportedModelsPaginatorOptions struct {
	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportedModelsPaginator is a paginator for ListImportedModels
type ListImportedModelsPaginator struct {
	options   ListImportedModelsPaginatorOptions
	client    ListImportedModelsAPIClient
	params    *ListImportedModelsInput
	nextToken *string
	firstPage bool
}

// NewListImportedModelsPaginator returns a new ListImportedModelsPaginator
func NewListImportedModelsPaginator(client ListImportedModelsAPIClient, params *ListImportedModelsInput, optFns ...func(*ListImportedModelsPaginatorOptions)) *ListImportedModelsPaginator {
	if params == nil {
		params = &ListImportedModelsInput{}
	}

	options := ListImportedModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportedModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportedModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImportedModels page.
func (p *ListImportedModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportedModelsOutput, error) {
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
	result, err := p.client.ListImportedModels(ctx, &params, optFns...)
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

// ListImportedModelsAPIClient is a client that implements the ListImportedModels
// operation.
type ListImportedModelsAPIClient interface {
	ListImportedModels(context.Context, *ListImportedModelsInput, ...func(*Options)) (*ListImportedModelsOutput, error)
}

var _ ListImportedModelsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListImportedModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImportedModels",
	}
}
