// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the policies attached to the specified principal. If you use an Cognito
// identity, the ID must be in [AmazonCognito Identity format].
//
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use ListAttachedPoliciesinstead.
//
// Requires permission to access the [ListPrincipalPolicies] action.
//
// Deprecated: This operation has been deprecated.
//
// [ListPrincipalPolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [AmazonCognito Identity format]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetCredentialsForIdentity.html#API_GetCredentialsForIdentity_RequestSyntax
func (c *Client) ListPrincipalPolicies(ctx context.Context, params *ListPrincipalPoliciesInput, optFns ...func(*Options)) (*ListPrincipalPoliciesOutput, error) {
	if params == nil {
		params = &ListPrincipalPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrincipalPolicies", params, optFns, c.addOperationListPrincipalPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrincipalPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListPrincipalPolicies operation.
type ListPrincipalPoliciesInput struct {

	// The principal. Valid principals are CertificateArn
	// (arn:aws:iot:region:accountId:cert/certificateId), thingGroupArn
	// (arn:aws:iot:region:accountId:thinggroup/groupName) and CognitoId (region:id).
	//
	// This member is required.
	Principal *string

	// Specifies the order for results. If true, results are returned in ascending
	// creation order.
	AscendingOrder bool

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32

	noSmithyDocumentSerde
}

// The output from the ListPrincipalPolicies operation.
type ListPrincipalPoliciesOutput struct {

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string

	// The policies.
	Policies []types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrincipalPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipalPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipalPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPrincipalPolicies"); err != nil {
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
	if err = addOpListPrincipalPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalPolicies(options.Region), middleware.Before); err != nil {
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

// ListPrincipalPoliciesPaginatorOptions is the paginator options for
// ListPrincipalPolicies
type ListPrincipalPoliciesPaginatorOptions struct {
	// The result page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrincipalPoliciesPaginator is a paginator for ListPrincipalPolicies
type ListPrincipalPoliciesPaginator struct {
	options   ListPrincipalPoliciesPaginatorOptions
	client    ListPrincipalPoliciesAPIClient
	params    *ListPrincipalPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListPrincipalPoliciesPaginator returns a new ListPrincipalPoliciesPaginator
func NewListPrincipalPoliciesPaginator(client ListPrincipalPoliciesAPIClient, params *ListPrincipalPoliciesInput, optFns ...func(*ListPrincipalPoliciesPaginatorOptions)) *ListPrincipalPoliciesPaginator {
	if params == nil {
		params = &ListPrincipalPoliciesInput{}
	}

	options := ListPrincipalPoliciesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrincipalPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrincipalPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrincipalPolicies page.
func (p *ListPrincipalPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrincipalPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListPrincipalPolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListPrincipalPoliciesAPIClient is a client that implements the
// ListPrincipalPolicies operation.
type ListPrincipalPoliciesAPIClient interface {
	ListPrincipalPolicies(context.Context, *ListPrincipalPoliciesInput, ...func(*Options)) (*ListPrincipalPoliciesOutput, error)
}

var _ ListPrincipalPoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPrincipalPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPrincipalPolicies",
	}
}
