// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the network resource relationships for the specified global network.
func (c *Client) GetNetworkResourceRelationships(ctx context.Context, params *GetNetworkResourceRelationshipsInput, optFns ...func(*Options)) (*GetNetworkResourceRelationshipsOutput, error) {
	if params == nil {
		params = &GetNetworkResourceRelationshipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetNetworkResourceRelationships", params, optFns, c.addOperationGetNetworkResourceRelationshipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetNetworkResourceRelationshipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNetworkResourceRelationshipsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Web Services account ID.
	AccountId *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// The ID of a core network.
	CoreNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ARN of the registered gateway.
	RegisteredGatewayArn *string

	// The ARN of the gateway.
	ResourceArn *string

	// The resource type.
	//
	// The following are the supported resource types for Direct Connect:
	//
	//   - dxcon
	//
	//   - dx-gateway
	//
	//   - dx-vif
	//
	// The following are the supported resource types for Network Manager:
	//
	//   - attachment
	//
	//   - connect-peer
	//
	//   - connection
	//
	//   - core-network
	//
	//   - device
	//
	//   - link
	//
	//   - peering
	//
	//   - site
	//
	// The following are the supported resource types for Amazon VPC:
	//
	//   - customer-gateway
	//
	//   - transit-gateway
	//
	//   - transit-gateway-attachment
	//
	//   - transit-gateway-connect-peer
	//
	//   - transit-gateway-route-table
	//
	//   - vpn-connection
	ResourceType *string

	noSmithyDocumentSerde
}

type GetNetworkResourceRelationshipsOutput struct {

	// The token for the next page of results.
	NextToken *string

	// The resource relationships.
	Relationships []types.Relationship

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetNetworkResourceRelationshipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetNetworkResourceRelationships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetNetworkResourceRelationships{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetNetworkResourceRelationships"); err != nil {
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
	if err = addOpGetNetworkResourceRelationshipsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetNetworkResourceRelationships(options.Region), middleware.Before); err != nil {
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

// GetNetworkResourceRelationshipsPaginatorOptions is the paginator options for
// GetNetworkResourceRelationships
type GetNetworkResourceRelationshipsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetNetworkResourceRelationshipsPaginator is a paginator for
// GetNetworkResourceRelationships
type GetNetworkResourceRelationshipsPaginator struct {
	options   GetNetworkResourceRelationshipsPaginatorOptions
	client    GetNetworkResourceRelationshipsAPIClient
	params    *GetNetworkResourceRelationshipsInput
	nextToken *string
	firstPage bool
}

// NewGetNetworkResourceRelationshipsPaginator returns a new
// GetNetworkResourceRelationshipsPaginator
func NewGetNetworkResourceRelationshipsPaginator(client GetNetworkResourceRelationshipsAPIClient, params *GetNetworkResourceRelationshipsInput, optFns ...func(*GetNetworkResourceRelationshipsPaginatorOptions)) *GetNetworkResourceRelationshipsPaginator {
	if params == nil {
		params = &GetNetworkResourceRelationshipsInput{}
	}

	options := GetNetworkResourceRelationshipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetNetworkResourceRelationshipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetNetworkResourceRelationshipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetNetworkResourceRelationships page.
func (p *GetNetworkResourceRelationshipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetNetworkResourceRelationshipsOutput, error) {
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
	result, err := p.client.GetNetworkResourceRelationships(ctx, &params, optFns...)
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

// GetNetworkResourceRelationshipsAPIClient is a client that implements the
// GetNetworkResourceRelationships operation.
type GetNetworkResourceRelationshipsAPIClient interface {
	GetNetworkResourceRelationships(context.Context, *GetNetworkResourceRelationshipsInput, ...func(*Options)) (*GetNetworkResourceRelationshipsOutput, error)
}

var _ GetNetworkResourceRelationshipsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetNetworkResourceRelationships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetNetworkResourceRelationships",
	}
}
