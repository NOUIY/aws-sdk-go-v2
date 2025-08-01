// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of medical transcription jobs that match the specified
// criteria. If no criteria are specified, all medical transcription jobs are
// returned.
//
// To get detailed information about a specific medical transcription job, use the
// operation.
func (c *Client) ListMedicalTranscriptionJobs(ctx context.Context, params *ListMedicalTranscriptionJobsInput, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) {
	if params == nil {
		params = &ListMedicalTranscriptionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMedicalTranscriptionJobs", params, optFns, c.addOperationListMedicalTranscriptionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMedicalTranscriptionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMedicalTranscriptionJobsInput struct {

	// Returns only the medical transcription jobs that contain the specified string.
	// The search is not case sensitive.
	JobNameContains *string

	// The maximum number of medical transcription jobs to return in each page of
	// results. If there are fewer results than the value that you specify, only the
	// actual results are returned. If you do not specify a value, a default of 5 is
	// used.
	MaxResults *int32

	// If your ListMedicalTranscriptionJobs request returns more results than can be
	// displayed, NextToken is displayed in the response with an associated string. To
	// get the next page of results, copy this string and repeat your request,
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Returns only medical transcription jobs with the specified status. Jobs are
	// ordered by creation date, with the newest job first. If you do not include
	// Status , all medical transcription jobs are returned.
	Status types.TranscriptionJobStatus

	noSmithyDocumentSerde
}

type ListMedicalTranscriptionJobsOutput struct {

	// Provides a summary of information about each result.
	MedicalTranscriptionJobSummaries []types.MedicalTranscriptionJobSummary

	// If NextToken is present in your response, it indicates that not all results are
	// displayed. To view the next set of results, copy the string associated with the
	// NextToken parameter in your results output, then run your request again
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Lists all medical transcription jobs that have the status specified in your
	// request. Jobs are ordered by creation date, with the newest job first.
	Status types.TranscriptionJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMedicalTranscriptionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMedicalTranscriptionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMedicalTranscriptionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMedicalTranscriptionJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(options.Region), middleware.Before); err != nil {
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

// ListMedicalTranscriptionJobsPaginatorOptions is the paginator options for
// ListMedicalTranscriptionJobs
type ListMedicalTranscriptionJobsPaginatorOptions struct {
	// The maximum number of medical transcription jobs to return in each page of
	// results. If there are fewer results than the value that you specify, only the
	// actual results are returned. If you do not specify a value, a default of 5 is
	// used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMedicalTranscriptionJobsPaginator is a paginator for
// ListMedicalTranscriptionJobs
type ListMedicalTranscriptionJobsPaginator struct {
	options   ListMedicalTranscriptionJobsPaginatorOptions
	client    ListMedicalTranscriptionJobsAPIClient
	params    *ListMedicalTranscriptionJobsInput
	nextToken *string
	firstPage bool
}

// NewListMedicalTranscriptionJobsPaginator returns a new
// ListMedicalTranscriptionJobsPaginator
func NewListMedicalTranscriptionJobsPaginator(client ListMedicalTranscriptionJobsAPIClient, params *ListMedicalTranscriptionJobsInput, optFns ...func(*ListMedicalTranscriptionJobsPaginatorOptions)) *ListMedicalTranscriptionJobsPaginator {
	if params == nil {
		params = &ListMedicalTranscriptionJobsInput{}
	}

	options := ListMedicalTranscriptionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMedicalTranscriptionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMedicalTranscriptionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMedicalTranscriptionJobs page.
func (p *ListMedicalTranscriptionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) {
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
	result, err := p.client.ListMedicalTranscriptionJobs(ctx, &params, optFns...)
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

// ListMedicalTranscriptionJobsAPIClient is a client that implements the
// ListMedicalTranscriptionJobs operation.
type ListMedicalTranscriptionJobsAPIClient interface {
	ListMedicalTranscriptionJobs(context.Context, *ListMedicalTranscriptionJobsInput, ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error)
}

var _ ListMedicalTranscriptionJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMedicalTranscriptionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMedicalTranscriptionJobs",
	}
}
