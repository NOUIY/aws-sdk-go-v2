// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The ListReviewPolicyResultsForHIT operation retrieves the computed results and
//
// the actions taken in the course of executing your Review Policies for a given
// HIT. For information about how to specify Review Policies when you call
// CreateHIT, see Review Policies. The ListReviewPolicyResultsForHIT operation can
// return results for both Assignment-level and HIT-level review results.
func (c *Client) ListReviewPolicyResultsForHIT(ctx context.Context, params *ListReviewPolicyResultsForHITInput, optFns ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error) {
	if params == nil {
		params = &ListReviewPolicyResultsForHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReviewPolicyResultsForHIT", params, optFns, c.addOperationListReviewPolicyResultsForHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReviewPolicyResultsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewPolicyResultsForHITInput struct {

	// The unique identifier of the HIT to retrieve review results for.
	//
	// This member is required.
	HITId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination token
	NextToken *string

	//  The Policy Level(s) to retrieve review results for - HIT or Assignment. If
	// omitted, the default behavior is to retrieve all data for both policy levels.
	// For a list of all the described policies, see Review Policies.
	PolicyLevels []types.ReviewPolicyLevel

	//  Specify if the operation should retrieve a list of the actions taken executing
	// the Review Policies and their outcomes.
	RetrieveActions *bool

	//  Specify if the operation should retrieve a list of the results computed by the
	// Review Policies.
	RetrieveResults *bool

	noSmithyDocumentSerde
}

type ListReviewPolicyResultsForHITOutput struct {

	//  The name of the Assignment-level Review Policy. This contains only the
	// PolicyName element.
	AssignmentReviewPolicy *types.ReviewPolicy

	//  Contains both ReviewResult and ReviewAction elements for an Assignment.
	AssignmentReviewReport *types.ReviewReport

	// The HITId of the HIT for which results have been returned.
	HITId *string

	// The name of the HIT-level Review Policy. This contains only the PolicyName
	// element.
	HITReviewPolicy *types.ReviewPolicy

	// Contains both ReviewResult and ReviewAction elements for a particular HIT.
	HITReviewReport *types.ReviewReport

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReviewPolicyResultsForHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReviewPolicyResultsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReviewPolicyResultsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReviewPolicyResultsForHIT"); err != nil {
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
	if err = addOpListReviewPolicyResultsForHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(options.Region), middleware.Before); err != nil {
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

// ListReviewPolicyResultsForHITPaginatorOptions is the paginator options for
// ListReviewPolicyResultsForHIT
type ListReviewPolicyResultsForHITPaginatorOptions struct {
	// Limit the number of results returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReviewPolicyResultsForHITPaginator is a paginator for
// ListReviewPolicyResultsForHIT
type ListReviewPolicyResultsForHITPaginator struct {
	options   ListReviewPolicyResultsForHITPaginatorOptions
	client    ListReviewPolicyResultsForHITAPIClient
	params    *ListReviewPolicyResultsForHITInput
	nextToken *string
	firstPage bool
}

// NewListReviewPolicyResultsForHITPaginator returns a new
// ListReviewPolicyResultsForHITPaginator
func NewListReviewPolicyResultsForHITPaginator(client ListReviewPolicyResultsForHITAPIClient, params *ListReviewPolicyResultsForHITInput, optFns ...func(*ListReviewPolicyResultsForHITPaginatorOptions)) *ListReviewPolicyResultsForHITPaginator {
	if params == nil {
		params = &ListReviewPolicyResultsForHITInput{}
	}

	options := ListReviewPolicyResultsForHITPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReviewPolicyResultsForHITPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReviewPolicyResultsForHITPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReviewPolicyResultsForHIT page.
func (p *ListReviewPolicyResultsForHITPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error) {
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
	result, err := p.client.ListReviewPolicyResultsForHIT(ctx, &params, optFns...)
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

// ListReviewPolicyResultsForHITAPIClient is a client that implements the
// ListReviewPolicyResultsForHIT operation.
type ListReviewPolicyResultsForHITAPIClient interface {
	ListReviewPolicyResultsForHIT(context.Context, *ListReviewPolicyResultsForHITInput, ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error)
}

var _ ListReviewPolicyResultsForHITAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReviewPolicyResultsForHIT",
	}
}
