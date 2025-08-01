// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns summary information about operations performed on a stack set.
//
// This API provides eventually consistent reads meaning it may take some time but
// will eventually return the most up-to-date data.
func (c *Client) ListStackSetOperations(ctx context.Context, params *ListStackSetOperationsInput, optFns ...func(*Options)) (*ListStackSetOperationsOutput, error) {
	if params == nil {
		params = &ListStackSetOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackSetOperations", params, optFns, c.addOperationListStackSetOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackSetOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackSetOperationsInput struct {

	// The name or unique ID of the stack set that you want to get operation summaries
	// for.
	//
	// This member is required.
	StackSetName *string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account.
	//
	// By default, SELF is specified. Use SELF for stack sets with self-managed
	// permissions.
	//
	//   - If you are signed in to the management account, specify SELF .
	//
	//   - If you are signed in to a delegated administrator account, specify
	//   DELEGATED_ADMIN .
	//
	// Your Amazon Web Services account must be registered as a delegated
	//   administrator in the management account. For more information, see [Register a delegated administrator]in the
	//   CloudFormation User Guide.
	//
	// [Register a delegated administrator]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html
	CallAs types.CallAs

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int32

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListStackSetOperations again and assign that
	// token to the request object's NextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null .
	NextToken *string

	noSmithyDocumentSerde
}

type ListStackSetOperationsOutput struct {

	// If the request doesn't return all results, NextToken is set to a token. To
	// retrieve the next set of results, call ListOperationResults again and assign
	// that token to the request object's NextToken parameter. If there are no
	// remaining results, NextToken is set to null .
	NextToken *string

	// A list of StackSetOperationSummary structures that contain summary information
	// about operations for the specified stack set.
	Summaries []types.StackSetOperationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStackSetOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackSetOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackSetOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStackSetOperations"); err != nil {
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
	if err = addOpListStackSetOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackSetOperations(options.Region), middleware.Before); err != nil {
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

// ListStackSetOperationsPaginatorOptions is the paginator options for
// ListStackSetOperations
type ListStackSetOperationsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStackSetOperationsPaginator is a paginator for ListStackSetOperations
type ListStackSetOperationsPaginator struct {
	options   ListStackSetOperationsPaginatorOptions
	client    ListStackSetOperationsAPIClient
	params    *ListStackSetOperationsInput
	nextToken *string
	firstPage bool
}

// NewListStackSetOperationsPaginator returns a new ListStackSetOperationsPaginator
func NewListStackSetOperationsPaginator(client ListStackSetOperationsAPIClient, params *ListStackSetOperationsInput, optFns ...func(*ListStackSetOperationsPaginatorOptions)) *ListStackSetOperationsPaginator {
	if params == nil {
		params = &ListStackSetOperationsInput{}
	}

	options := ListStackSetOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStackSetOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStackSetOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStackSetOperations page.
func (p *ListStackSetOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStackSetOperationsOutput, error) {
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
	result, err := p.client.ListStackSetOperations(ctx, &params, optFns...)
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

// ListStackSetOperationsAPIClient is a client that implements the
// ListStackSetOperations operation.
type ListStackSetOperationsAPIClient interface {
	ListStackSetOperations(context.Context, *ListStackSetOperationsInput, ...func(*Options)) (*ListStackSetOperationsOutput, error)
}

var _ ListStackSetOperationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStackSetOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStackSetOperations",
	}
}
