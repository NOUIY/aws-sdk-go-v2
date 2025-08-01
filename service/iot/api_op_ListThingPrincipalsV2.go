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

// Lists the principals associated with the specified thing. A principal can be an
// X.509 certificate or an Amazon Cognito ID.
//
// Requires permission to access the [ListThingPrincipals] action.
//
// [ListThingPrincipals]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListThingPrincipalsV2(ctx context.Context, params *ListThingPrincipalsV2Input, optFns ...func(*Options)) (*ListThingPrincipalsV2Output, error) {
	if params == nil {
		params = &ListThingPrincipalsV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingPrincipalsV2", params, optFns, c.addOperationListThingPrincipalsV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingPrincipalsV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingPrincipalsV2Input struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// The type of the relation you want to filter in the response. If no value is
	// provided in this field, the response will list all principals, including both
	// the EXCLUSIVE_THING and NON_EXCLUSIVE_THING attachment types.
	//
	//   - EXCLUSIVE_THING - Attaches the specified principal to the specified thing,
	//   exclusively. The thing will be the only thing that’s attached to the principal.
	//
	//   - NON_EXCLUSIVE_THING - Attaches the specified principal to the specified
	//   thing. Multiple things can be attached to the principal.
	ThingPrincipalType types.ThingPrincipalType

	noSmithyDocumentSerde
}

type ListThingPrincipalsV2Output struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// A list of thingPrincipalObject that represents the principal and the type of
	// relation it has with the thing.
	ThingPrincipalObjects []types.ThingPrincipalObject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingPrincipalsV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingPrincipalsV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingPrincipalsV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListThingPrincipalsV2"); err != nil {
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
	if err = addOpListThingPrincipalsV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingPrincipalsV2(options.Region), middleware.Before); err != nil {
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

// ListThingPrincipalsV2PaginatorOptions is the paginator options for
// ListThingPrincipalsV2
type ListThingPrincipalsV2PaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingPrincipalsV2Paginator is a paginator for ListThingPrincipalsV2
type ListThingPrincipalsV2Paginator struct {
	options   ListThingPrincipalsV2PaginatorOptions
	client    ListThingPrincipalsV2APIClient
	params    *ListThingPrincipalsV2Input
	nextToken *string
	firstPage bool
}

// NewListThingPrincipalsV2Paginator returns a new ListThingPrincipalsV2Paginator
func NewListThingPrincipalsV2Paginator(client ListThingPrincipalsV2APIClient, params *ListThingPrincipalsV2Input, optFns ...func(*ListThingPrincipalsV2PaginatorOptions)) *ListThingPrincipalsV2Paginator {
	if params == nil {
		params = &ListThingPrincipalsV2Input{}
	}

	options := ListThingPrincipalsV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingPrincipalsV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingPrincipalsV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingPrincipalsV2 page.
func (p *ListThingPrincipalsV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingPrincipalsV2Output, error) {
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
	result, err := p.client.ListThingPrincipalsV2(ctx, &params, optFns...)
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

// ListThingPrincipalsV2APIClient is a client that implements the
// ListThingPrincipalsV2 operation.
type ListThingPrincipalsV2APIClient interface {
	ListThingPrincipalsV2(context.Context, *ListThingPrincipalsV2Input, ...func(*Options)) (*ListThingPrincipalsV2Output, error)
}

var _ ListThingPrincipalsV2APIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListThingPrincipalsV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListThingPrincipalsV2",
	}
}
