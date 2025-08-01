// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about the managed permission and its associations to any
// resource shares that use this managed permission. This lets you see which
// resource shares use which versions of the specified managed permission.
func (c *Client) ListPermissionAssociations(ctx context.Context, params *ListPermissionAssociationsInput, optFns ...func(*Options)) (*ListPermissionAssociationsOutput, error) {
	if params == nil {
		params = &ListPermissionAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionAssociations", params, optFns, c.addOperationListPermissionAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionAssociationsInput struct {

	// Specifies that you want to list only those associations with resource shares
	// that match this status.
	AssociationStatus types.ResourceShareAssociationStatus

	// When true , specifies that you want to list only those associations with
	// resource shares that use the default version of the specified managed
	// permission.
	//
	// When false (the default value), lists associations with resource shares that
	// use any version of the specified managed permission.
	DefaultVersion *bool

	// Specifies that you want to list only those associations with resource shares
	// that have a featureSet with this value.
	FeatureSet types.PermissionFeatureSet

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies the [Amazon Resource Name (ARN)] of the managed permission.
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	PermissionArn *string

	// Specifies that you want to list only those associations with resource shares
	// that use this version of the managed permission. If you don't provide a value
	// for this parameter, then the operation returns information about associations
	// with resource shares that use any version of the managed permission.
	PermissionVersion *int32

	// Specifies that you want to list only those associations with resource shares
	// that include at least one resource of this resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

type ListPermissionAssociationsOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// A structure with information about this customer managed permission.
	Permissions []types.AssociatedPermission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPermissionAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPermissionAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPermissionAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionAssociations(options.Region), middleware.Before); err != nil {
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

// ListPermissionAssociationsPaginatorOptions is the paginator options for
// ListPermissionAssociations
type ListPermissionAssociationsPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionAssociationsPaginator is a paginator for
// ListPermissionAssociations
type ListPermissionAssociationsPaginator struct {
	options   ListPermissionAssociationsPaginatorOptions
	client    ListPermissionAssociationsAPIClient
	params    *ListPermissionAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionAssociationsPaginator returns a new
// ListPermissionAssociationsPaginator
func NewListPermissionAssociationsPaginator(client ListPermissionAssociationsAPIClient, params *ListPermissionAssociationsInput, optFns ...func(*ListPermissionAssociationsPaginatorOptions)) *ListPermissionAssociationsPaginator {
	if params == nil {
		params = &ListPermissionAssociationsInput{}
	}

	options := ListPermissionAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissionAssociations page.
func (p *ListPermissionAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionAssociationsOutput, error) {
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
	result, err := p.client.ListPermissionAssociations(ctx, &params, optFns...)
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

// ListPermissionAssociationsAPIClient is a client that implements the
// ListPermissionAssociations operation.
type ListPermissionAssociationsAPIClient interface {
	ListPermissionAssociations(context.Context, *ListPermissionAssociationsInput, ...func(*Options)) (*ListPermissionAssociationsOutput, error)
}

var _ ListPermissionAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPermissionAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPermissionAssociations",
	}
}
