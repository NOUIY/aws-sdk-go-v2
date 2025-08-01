// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the organizational units (OUs) or accounts that are contained in
// the specified parent OU or root. This operation, along with ListParentsenables you to
// traverse the tree structure that makes up this root.
//
// Always check the NextToken response parameter for a null value when calling a
// List* operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator.
func (c *Client) ListChildren(ctx context.Context, params *ListChildrenInput, optFns ...func(*Options)) (*ListChildrenOutput, error) {
	if params == nil {
		params = &ListChildrenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChildren", params, optFns, c.addOperationListChildrenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChildrenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChildrenInput struct {

	// Filters the output to include only the specified child type.
	//
	// This member is required.
	ChildType types.ChildType

	// The unique identifier (ID) for the parent root or OU whose children you want to
	// list.
	//
	// The [regex pattern] for a parent ID string requires one of the following:
	//
	//   - Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//   letters or digits.
	//
	//   - Organizational unit (OU) - A string that begins with "ou-" followed by from
	//   4 to 32 lowercase letters or digits (the ID of the root that the OU is in). This
	//   string is followed by a second "-" dash and from 8 to 32 additional lowercase
	//   letters or digits.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	ParentId *string

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListChildrenOutput struct {

	// The list of children of the specified parent container.
	Children []types.Child

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChildrenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListChildren{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListChildren{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListChildren"); err != nil {
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
	if err = addOpListChildrenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChildren(options.Region), middleware.Before); err != nil {
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

// ListChildrenPaginatorOptions is the paginator options for ListChildren
type ListChildrenPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChildrenPaginator is a paginator for ListChildren
type ListChildrenPaginator struct {
	options   ListChildrenPaginatorOptions
	client    ListChildrenAPIClient
	params    *ListChildrenInput
	nextToken *string
	firstPage bool
}

// NewListChildrenPaginator returns a new ListChildrenPaginator
func NewListChildrenPaginator(client ListChildrenAPIClient, params *ListChildrenInput, optFns ...func(*ListChildrenPaginatorOptions)) *ListChildrenPaginator {
	if params == nil {
		params = &ListChildrenInput{}
	}

	options := ListChildrenPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChildrenPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChildrenPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChildren page.
func (p *ListChildrenPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChildrenOutput, error) {
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
	result, err := p.client.ListChildren(ctx, &params, optFns...)
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

// ListChildrenAPIClient is a client that implements the ListChildren operation.
type ListChildrenAPIClient interface {
	ListChildren(context.Context, *ListChildrenInput, ...func(*Options)) (*ListChildrenOutput, error)
}

var _ ListChildrenAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListChildren(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListChildren",
	}
}
