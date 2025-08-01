// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a list of [HyperParameterTuningJobSummary] objects that describe the hyperparameter tuning jobs launched
// in your account.
//
// [HyperParameterTuningJobSummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobSummary.html
func (c *Client) ListHyperParameterTuningJobs(ctx context.Context, params *ListHyperParameterTuningJobsInput, optFns ...func(*Options)) (*ListHyperParameterTuningJobsOutput, error) {
	if params == nil {
		params = &ListHyperParameterTuningJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHyperParameterTuningJobs", params, optFns, c.addOperationListHyperParameterTuningJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHyperParameterTuningJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHyperParameterTuningJobsInput struct {

	// A filter that returns only tuning jobs that were created after the specified
	// time.
	CreationTimeAfter *time.Time

	// A filter that returns only tuning jobs that were created before the specified
	// time.
	CreationTimeBefore *time.Time

	// A filter that returns only tuning jobs that were modified after the specified
	// time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only tuning jobs that were modified before the specified
	// time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of tuning jobs to return. The default value is 10.
	MaxResults *int32

	// A string in the tuning job name. This filter returns only tuning jobs whose
	// name contains the specified string.
	NameContains *string

	// If the result of the previous ListHyperParameterTuningJobs request was
	// truncated, the response includes a NextToken . To retrieve the next set of
	// tuning jobs, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is Name .
	SortBy types.HyperParameterTuningJobSortByOptions

	// The sort order for results. The default is Ascending .
	SortOrder types.SortOrder

	// A filter that returns only tuning jobs with the specified status.
	StatusEquals types.HyperParameterTuningJobStatus

	noSmithyDocumentSerde
}

type ListHyperParameterTuningJobsOutput struct {

	// A list of [HyperParameterTuningJobSummary] objects that describe the tuning jobs that the
	// ListHyperParameterTuningJobs request returned.
	//
	// [HyperParameterTuningJobSummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobSummary.html
	//
	// This member is required.
	HyperParameterTuningJobSummaries []types.HyperParameterTuningJobSummary

	// If the result of this ListHyperParameterTuningJobs request was truncated, the
	// response includes a NextToken . To retrieve the next set of tuning jobs, use the
	// token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHyperParameterTuningJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHyperParameterTuningJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHyperParameterTuningJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHyperParameterTuningJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHyperParameterTuningJobs(options.Region), middleware.Before); err != nil {
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

// ListHyperParameterTuningJobsPaginatorOptions is the paginator options for
// ListHyperParameterTuningJobs
type ListHyperParameterTuningJobsPaginatorOptions struct {
	// The maximum number of tuning jobs to return. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHyperParameterTuningJobsPaginator is a paginator for
// ListHyperParameterTuningJobs
type ListHyperParameterTuningJobsPaginator struct {
	options   ListHyperParameterTuningJobsPaginatorOptions
	client    ListHyperParameterTuningJobsAPIClient
	params    *ListHyperParameterTuningJobsInput
	nextToken *string
	firstPage bool
}

// NewListHyperParameterTuningJobsPaginator returns a new
// ListHyperParameterTuningJobsPaginator
func NewListHyperParameterTuningJobsPaginator(client ListHyperParameterTuningJobsAPIClient, params *ListHyperParameterTuningJobsInput, optFns ...func(*ListHyperParameterTuningJobsPaginatorOptions)) *ListHyperParameterTuningJobsPaginator {
	if params == nil {
		params = &ListHyperParameterTuningJobsInput{}
	}

	options := ListHyperParameterTuningJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHyperParameterTuningJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHyperParameterTuningJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHyperParameterTuningJobs page.
func (p *ListHyperParameterTuningJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHyperParameterTuningJobsOutput, error) {
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
	result, err := p.client.ListHyperParameterTuningJobs(ctx, &params, optFns...)
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

// ListHyperParameterTuningJobsAPIClient is a client that implements the
// ListHyperParameterTuningJobs operation.
type ListHyperParameterTuningJobsAPIClient interface {
	ListHyperParameterTuningJobs(context.Context, *ListHyperParameterTuningJobsInput, ...func(*Options)) (*ListHyperParameterTuningJobsOutput, error)
}

var _ ListHyperParameterTuningJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListHyperParameterTuningJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHyperParameterTuningJobs",
	}
}
