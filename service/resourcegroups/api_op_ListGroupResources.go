// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Amazon resource names (ARNs) of the resources that are
// members of a specified resource group.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
//   - resource-groups:ListGroupResources
//
//   - cloudformation:DescribeStacks
//
//   - cloudformation:ListStackResources
//
//   - tag:GetResources
func (c *Client) ListGroupResources(ctx context.Context, params *ListGroupResourcesInput, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) {
	if params == nil {
		params = &ListGroupResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupResources", params, optFns, c.addOperationListGroupResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupResourcesInput struct {

	// Filters, formatted as ResourceFilter objects, that you want to apply to a ListGroupResources
	// operation. Filters the results to include only those of the specified resource
	// types.
	//
	//   - resource-type - Filter resources by their type. Specify up to five resource
	//   types in the format AWS::ServiceCode::ResourceType . For example,
	//   AWS::EC2::Instance , or AWS::S3::Bucket .
	//
	// When you specify a resource-type filter for ListGroupResources , Resource Groups
	// validates your filter resource types against the types that are defined in the
	// query associated with the group. For example, if a group contains only S3
	// buckets because its query specifies only that resource type, but your
	// resource-type filter includes EC2 instances, AWS Resource Groups does not filter
	// for EC2 instances. In this case, a ListGroupResources request returns a
	// BadRequestException error with a message similar to the following:
	//
	//     The resource types specified as filters in the request are not valid.
	//
	// The error includes a list of resource types that failed the validation because
	// they are not part of the query associated with the group. This validation
	// doesn't occur when the group query specifies AWS::AllSupported , because a group
	// based on such a query can contain any of the allowed resource types for the
	// query type (tag-based or Amazon CloudFront stack-based queries).
	Filters []types.ResourceFilter

	// The name or the Amazon resource name (ARN) of the resource group.
	Group *string

	//  Deprecated - don't use this parameter. Use the Group request field instead.
	//
	// Deprecated: This field is deprecated, use Group instead.
	GroupName *string

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value provided by a previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupResourcesOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// A list of QueryError objects. Each error contains an ErrorCode and Message .
	// Possible values for ErrorCode are CLOUDFORMATION_STACK_INACTIVE ,
	// CLOUDFORMATION_STACK_NOT_EXISTING , CLOUDFORMATION_STACK_UNASSUMABLE_ROLE and
	// RESOURCE_TYPE_NOT_SUPPORTED .
	QueryErrors []types.QueryError

	//  Deprecated - don't use this parameter. Use the Resources response field
	// instead.
	//
	// Deprecated: This field is deprecated, use Resources instead.
	ResourceIdentifiers []types.ResourceIdentifier

	// An array of resources from which you can determine each resource's identity,
	// type, and group membership status.
	Resources []types.ListGroupResourcesItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGroupResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroupResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroupResources"); err != nil {
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
	if err = addOpListGroupResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupResources(options.Region), middleware.Before); err != nil {
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

// ListGroupResourcesPaginatorOptions is the paginator options for
// ListGroupResources
type ListGroupResourcesPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupResourcesPaginator is a paginator for ListGroupResources
type ListGroupResourcesPaginator struct {
	options   ListGroupResourcesPaginatorOptions
	client    ListGroupResourcesAPIClient
	params    *ListGroupResourcesInput
	nextToken *string
	firstPage bool
}

// NewListGroupResourcesPaginator returns a new ListGroupResourcesPaginator
func NewListGroupResourcesPaginator(client ListGroupResourcesAPIClient, params *ListGroupResourcesInput, optFns ...func(*ListGroupResourcesPaginatorOptions)) *ListGroupResourcesPaginator {
	if params == nil {
		params = &ListGroupResourcesInput{}
	}

	options := ListGroupResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupResources page.
func (p *ListGroupResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) {
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
	result, err := p.client.ListGroupResources(ctx, &params, optFns...)
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

// ListGroupResourcesAPIClient is a client that implements the ListGroupResources
// operation.
type ListGroupResourcesAPIClient interface {
	ListGroupResources(context.Context, *ListGroupResourcesInput, ...func(*Options)) (*ListGroupResourcesOutput, error)
}

var _ ListGroupResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListGroupResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroupResources",
	}
}
