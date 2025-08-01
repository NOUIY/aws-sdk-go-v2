// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the associations for the transit gateway multicast
// domain.
func (c *Client) GetTransitGatewayMulticastDomainAssociations(ctx context.Context, params *GetTransitGatewayMulticastDomainAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayMulticastDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayMulticastDomainAssociations", params, optFns, c.addOperationGetTransitGatewayMulticastDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayMulticastDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayMulticastDomainAssociationsInput struct {

	// The ID of the transit gateway multicast domain.
	//
	// This member is required.
	TransitGatewayMulticastDomainId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//   - resource-id - The ID of the resource.
	//
	//   - resource-type - The type of resource. The valid value is: vpc .
	//
	//   - state - The state of the subnet association. Valid values are associated |
	//   associating | disassociated | disassociating .
	//
	//   - subnet-id - The ID of the subnet.
	//
	//   - transit-gateway-attachment-id - The id of the transit gateway attachment.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTransitGatewayMulticastDomainAssociationsOutput struct {

	// Information about the multicast domain associations.
	MulticastDomainAssociations []types.TransitGatewayMulticastDomainAssociation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransitGatewayMulticastDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetTransitGatewayMulticastDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetTransitGatewayMulticastDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTransitGatewayMulticastDomainAssociations"); err != nil {
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
	if err = addOpGetTransitGatewayMulticastDomainAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayMulticastDomainAssociationsPaginatorOptions is the paginator
// options for GetTransitGatewayMulticastDomainAssociations
type GetTransitGatewayMulticastDomainAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayMulticastDomainAssociationsPaginator is a paginator for
// GetTransitGatewayMulticastDomainAssociations
type GetTransitGatewayMulticastDomainAssociationsPaginator struct {
	options   GetTransitGatewayMulticastDomainAssociationsPaginatorOptions
	client    GetTransitGatewayMulticastDomainAssociationsAPIClient
	params    *GetTransitGatewayMulticastDomainAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayMulticastDomainAssociationsPaginator returns a new
// GetTransitGatewayMulticastDomainAssociationsPaginator
func NewGetTransitGatewayMulticastDomainAssociationsPaginator(client GetTransitGatewayMulticastDomainAssociationsAPIClient, params *GetTransitGatewayMulticastDomainAssociationsInput, optFns ...func(*GetTransitGatewayMulticastDomainAssociationsPaginatorOptions)) *GetTransitGatewayMulticastDomainAssociationsPaginator {
	if params == nil {
		params = &GetTransitGatewayMulticastDomainAssociationsInput{}
	}

	options := GetTransitGatewayMulticastDomainAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayMulticastDomainAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayMulticastDomainAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTransitGatewayMulticastDomainAssociations page.
func (p *GetTransitGatewayMulticastDomainAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayMulticastDomainAssociationsOutput, error) {
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
	result, err := p.client.GetTransitGatewayMulticastDomainAssociations(ctx, &params, optFns...)
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

// GetTransitGatewayMulticastDomainAssociationsAPIClient is a client that
// implements the GetTransitGatewayMulticastDomainAssociations operation.
type GetTransitGatewayMulticastDomainAssociationsAPIClient interface {
	GetTransitGatewayMulticastDomainAssociations(context.Context, *GetTransitGatewayMulticastDomainAssociationsInput, ...func(*Options)) (*GetTransitGatewayMulticastDomainAssociationsOutput, error)
}

var _ GetTransitGatewayMulticastDomainAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetTransitGatewayMulticastDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTransitGatewayMulticastDomainAssociations",
	}
}
