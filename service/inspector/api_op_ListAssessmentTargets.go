// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the ARNs of the assessment targets within this AWS account. For more
// information about assessment targets, see [Amazon Inspector Assessment Targets].
//
// [Amazon Inspector Assessment Targets]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html
func (c *Client) ListAssessmentTargets(ctx context.Context, params *ListAssessmentTargetsInput, optFns ...func(*Options)) (*ListAssessmentTargetsOutput, error) {
	if params == nil {
		params = &ListAssessmentTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentTargets", params, optFns, c.addOperationListAssessmentTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentTargetsInput struct {

	// You can use this parameter to specify a subset of data to be included in the
	// action's response.
	//
	// For a record to match a filter, all specified filter attributes must match.
	// When multiple values are specified for a filter attribute, any of the values can
	// match.
	Filter *types.AssessmentTargetFilter

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListAssessmentTargets action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// NextToken from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentTargetsOutput struct {

	// A list of ARNs that specifies the assessment targets that are returned by the
	// action.
	//
	// This member is required.
	AssessmentTargetArns []string

	//  When a response is generated, if there is more data to be listed, this
	// parameter is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssessmentTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssessmentTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssessmentTargets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentTargets(options.Region), middleware.Before); err != nil {
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

// ListAssessmentTargetsPaginatorOptions is the paginator options for
// ListAssessmentTargets
type ListAssessmentTargetsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentTargetsPaginator is a paginator for ListAssessmentTargets
type ListAssessmentTargetsPaginator struct {
	options   ListAssessmentTargetsPaginatorOptions
	client    ListAssessmentTargetsAPIClient
	params    *ListAssessmentTargetsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentTargetsPaginator returns a new ListAssessmentTargetsPaginator
func NewListAssessmentTargetsPaginator(client ListAssessmentTargetsAPIClient, params *ListAssessmentTargetsInput, optFns ...func(*ListAssessmentTargetsPaginatorOptions)) *ListAssessmentTargetsPaginator {
	if params == nil {
		params = &ListAssessmentTargetsInput{}
	}

	options := ListAssessmentTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentTargets page.
func (p *ListAssessmentTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentTargetsOutput, error) {
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
	result, err := p.client.ListAssessmentTargets(ctx, &params, optFns...)
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

// ListAssessmentTargetsAPIClient is a client that implements the
// ListAssessmentTargets operation.
type ListAssessmentTargetsAPIClient interface {
	ListAssessmentTargets(context.Context, *ListAssessmentTargetsInput, ...func(*Options)) (*ListAssessmentTargetsOutput, error)
}

var _ ListAssessmentTargetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAssessmentTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssessmentTargets",
	}
}
