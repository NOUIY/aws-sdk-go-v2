// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all members of the specified project.
func (c *Client) ListProjectMemberships(ctx context.Context, params *ListProjectMembershipsInput, optFns ...func(*Options)) (*ListProjectMembershipsOutput, error) {
	if params == nil {
		params = &ListProjectMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProjectMemberships", params, optFns, c.addOperationListProjectMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProjectMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProjectMembershipsInput struct {

	// The identifier of the Amazon DataZone domain in which you want to list project
	// memberships.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the project whose memberships you want to list.
	//
	// This member is required.
	ProjectIdentifier *string

	// The maximum number of memberships to return in a single call to
	// ListProjectMemberships . When the number of memberships to be listed is greater
	// than the value of MaxResults , the response contains a NextToken value that you
	// can use in a subsequent call to ListProjectMemberships to list the next set of
	// memberships.
	MaxResults *int32

	// When the number of memberships is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of memberships, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListProjectMemberships to list the next set of memberships.
	NextToken *string

	// The method by which you want to sort the project memberships.
	SortBy types.SortFieldProject

	// The sort order of the project memberships.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListProjectMembershipsOutput struct {

	// The members of the project.
	//
	// This member is required.
	Members []types.ProjectMember

	// When the number of memberships is greater than the default value for the
	// MaxResults parameter, or if you explicitly specify a value for MaxResults that
	// is less than the number of memberships, the response includes a pagination token
	// named NextToken . You can specify this NextToken value in a subsequent call to
	// ListProjectMemberships to list the next set of memberships.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProjectMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProjectMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProjectMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProjectMemberships"); err != nil {
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
	if err = addOpListProjectMembershipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProjectMemberships(options.Region), middleware.Before); err != nil {
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

// ListProjectMembershipsPaginatorOptions is the paginator options for
// ListProjectMemberships
type ListProjectMembershipsPaginatorOptions struct {
	// The maximum number of memberships to return in a single call to
	// ListProjectMemberships . When the number of memberships to be listed is greater
	// than the value of MaxResults , the response contains a NextToken value that you
	// can use in a subsequent call to ListProjectMemberships to list the next set of
	// memberships.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProjectMembershipsPaginator is a paginator for ListProjectMemberships
type ListProjectMembershipsPaginator struct {
	options   ListProjectMembershipsPaginatorOptions
	client    ListProjectMembershipsAPIClient
	params    *ListProjectMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListProjectMembershipsPaginator returns a new ListProjectMembershipsPaginator
func NewListProjectMembershipsPaginator(client ListProjectMembershipsAPIClient, params *ListProjectMembershipsInput, optFns ...func(*ListProjectMembershipsPaginatorOptions)) *ListProjectMembershipsPaginator {
	if params == nil {
		params = &ListProjectMembershipsInput{}
	}

	options := ListProjectMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProjectMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProjectMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProjectMemberships page.
func (p *ListProjectMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProjectMembershipsOutput, error) {
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
	result, err := p.client.ListProjectMemberships(ctx, &params, optFns...)
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

// ListProjectMembershipsAPIClient is a client that implements the
// ListProjectMemberships operation.
type ListProjectMembershipsAPIClient interface {
	ListProjectMemberships(context.Context, *ListProjectMembershipsInput, ...func(*Options)) (*ListProjectMembershipsOutput, error)
}

var _ ListProjectMembershipsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListProjectMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProjectMemberships",
	}
}
