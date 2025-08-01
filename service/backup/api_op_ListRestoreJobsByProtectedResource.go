// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This returns restore jobs that contain the specified protected resource.
//
// You must include ResourceArn . You can optionally include NextToken , ByStatus ,
// MaxResults , ByRecoveryPointCreationDateAfter , and
// ByRecoveryPointCreationDateBefore .
func (c *Client) ListRestoreJobsByProtectedResource(ctx context.Context, params *ListRestoreJobsByProtectedResourceInput, optFns ...func(*Options)) (*ListRestoreJobsByProtectedResourceOutput, error) {
	if params == nil {
		params = &ListRestoreJobsByProtectedResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRestoreJobsByProtectedResource", params, optFns, c.addOperationListRestoreJobsByProtectedResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRestoreJobsByProtectedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRestoreJobsByProtectedResourceInput struct {

	// Returns only restore jobs that match the specified resource Amazon Resource
	// Name (ARN).
	//
	// This member is required.
	ResourceArn *string

	// Returns only restore jobs of recovery points that were created after the
	// specified date.
	ByRecoveryPointCreationDateAfter *time.Time

	// Returns only restore jobs of recovery points that were created before the
	// specified date.
	ByRecoveryPointCreationDateBefore *time.Time

	// Returns only restore jobs associated with the specified job status.
	ByStatus types.RestoreJobStatus

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request ismade to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRestoreJobsByProtectedResourceOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows youto
	// return more items in your list starting at the location pointed to by the next
	// token
	NextToken *string

	// An array of objects that contain detailed information about jobs to restore
	// saved resources.>
	RestoreJobs []types.RestoreJobsListMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRestoreJobsByProtectedResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRestoreJobsByProtectedResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRestoreJobsByProtectedResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRestoreJobsByProtectedResource"); err != nil {
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
	if err = addOpListRestoreJobsByProtectedResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRestoreJobsByProtectedResource(options.Region), middleware.Before); err != nil {
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

// ListRestoreJobsByProtectedResourcePaginatorOptions is the paginator options for
// ListRestoreJobsByProtectedResource
type ListRestoreJobsByProtectedResourcePaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRestoreJobsByProtectedResourcePaginator is a paginator for
// ListRestoreJobsByProtectedResource
type ListRestoreJobsByProtectedResourcePaginator struct {
	options   ListRestoreJobsByProtectedResourcePaginatorOptions
	client    ListRestoreJobsByProtectedResourceAPIClient
	params    *ListRestoreJobsByProtectedResourceInput
	nextToken *string
	firstPage bool
}

// NewListRestoreJobsByProtectedResourcePaginator returns a new
// ListRestoreJobsByProtectedResourcePaginator
func NewListRestoreJobsByProtectedResourcePaginator(client ListRestoreJobsByProtectedResourceAPIClient, params *ListRestoreJobsByProtectedResourceInput, optFns ...func(*ListRestoreJobsByProtectedResourcePaginatorOptions)) *ListRestoreJobsByProtectedResourcePaginator {
	if params == nil {
		params = &ListRestoreJobsByProtectedResourceInput{}
	}

	options := ListRestoreJobsByProtectedResourcePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRestoreJobsByProtectedResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRestoreJobsByProtectedResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRestoreJobsByProtectedResource page.
func (p *ListRestoreJobsByProtectedResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRestoreJobsByProtectedResourceOutput, error) {
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
	result, err := p.client.ListRestoreJobsByProtectedResource(ctx, &params, optFns...)
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

// ListRestoreJobsByProtectedResourceAPIClient is a client that implements the
// ListRestoreJobsByProtectedResource operation.
type ListRestoreJobsByProtectedResourceAPIClient interface {
	ListRestoreJobsByProtectedResource(context.Context, *ListRestoreJobsByProtectedResourceInput, ...func(*Options)) (*ListRestoreJobsByProtectedResourceOutput, error)
}

var _ ListRestoreJobsByProtectedResourceAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRestoreJobsByProtectedResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRestoreJobsByProtectedResource",
	}
}
