// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of read set activation jobs.
func (c *Client) ListReadSetActivationJobs(ctx context.Context, params *ListReadSetActivationJobsInput, optFns ...func(*Options)) (*ListReadSetActivationJobsOutput, error) {
	if params == nil {
		params = &ListReadSetActivationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadSetActivationJobs", params, optFns, c.addOperationListReadSetActivationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadSetActivationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadSetActivationJobsInput struct {

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// A filter to apply to the list.
	Filter *types.ActivateReadSetFilter

	// The maximum number of read set activation jobs to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadSetActivationJobsOutput struct {

	// A list of jobs.
	ActivationJobs []types.ActivateReadSetJobItem

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadSetActivationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadSetActivationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadSetActivationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReadSetActivationJobs"); err != nil {
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
	if err = addEndpointPrefix_opListReadSetActivationJobsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListReadSetActivationJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadSetActivationJobs(options.Region), middleware.Before); err != nil {
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

// ListReadSetActivationJobsPaginatorOptions is the paginator options for
// ListReadSetActivationJobs
type ListReadSetActivationJobsPaginatorOptions struct {
	// The maximum number of read set activation jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadSetActivationJobsPaginator is a paginator for ListReadSetActivationJobs
type ListReadSetActivationJobsPaginator struct {
	options   ListReadSetActivationJobsPaginatorOptions
	client    ListReadSetActivationJobsAPIClient
	params    *ListReadSetActivationJobsInput
	nextToken *string
	firstPage bool
}

// NewListReadSetActivationJobsPaginator returns a new
// ListReadSetActivationJobsPaginator
func NewListReadSetActivationJobsPaginator(client ListReadSetActivationJobsAPIClient, params *ListReadSetActivationJobsInput, optFns ...func(*ListReadSetActivationJobsPaginatorOptions)) *ListReadSetActivationJobsPaginator {
	if params == nil {
		params = &ListReadSetActivationJobsInput{}
	}

	options := ListReadSetActivationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadSetActivationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadSetActivationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadSetActivationJobs page.
func (p *ListReadSetActivationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadSetActivationJobsOutput, error) {
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
	result, err := p.client.ListReadSetActivationJobs(ctx, &params, optFns...)
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

type endpointPrefix_opListReadSetActivationJobsMiddleware struct {
}

func (*endpointPrefix_opListReadSetActivationJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListReadSetActivationJobsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListReadSetActivationJobsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListReadSetActivationJobsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListReadSetActivationJobsAPIClient is a client that implements the
// ListReadSetActivationJobs operation.
type ListReadSetActivationJobsAPIClient interface {
	ListReadSetActivationJobs(context.Context, *ListReadSetActivationJobsInput, ...func(*Options)) (*ListReadSetActivationJobsOutput, error)
}

var _ ListReadSetActivationJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListReadSetActivationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReadSetActivationJobs",
	}
}
