// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the metadata for the VPC endpoint associations that you have defined.
// If you specify a fireawll, this returns only the endpoint associations for that
// firewall.
//
// Depending on your setting for max results and the number of associations, a
// single call might not return the full list.
func (c *Client) ListVpcEndpointAssociations(ctx context.Context, params *ListVpcEndpointAssociationsInput, optFns ...func(*Options)) (*ListVpcEndpointAssociationsOutput, error) {
	if params == nil {
		params = &ListVpcEndpointAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcEndpointAssociations", params, optFns, c.addOperationListVpcEndpointAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcEndpointAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVpcEndpointAssociationsInput struct {

	// The Amazon Resource Name (ARN) of the firewall.
	//
	// If you don't specify this, Network Firewall retrieves all VPC endpoint
	// associations that you have defined.
	FirewallArn *string

	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVpcEndpointAssociationsOutput struct {

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// The VPC endpoint assocation metadata objects for the firewall that you
	// specified. If you didn't specify a firewall, this is all VPC endpoint
	// associations that you have defined.
	//
	// Depending on your setting for max results and the number of firewalls you have,
	// a single call might not be the full list.
	VpcEndpointAssociations []types.VpcEndpointAssociationMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcEndpointAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVpcEndpointAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVpcEndpointAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVpcEndpointAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcEndpointAssociations(options.Region), middleware.Before); err != nil {
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

// ListVpcEndpointAssociationsPaginatorOptions is the paginator options for
// ListVpcEndpointAssociations
type ListVpcEndpointAssociationsPaginatorOptions struct {
	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVpcEndpointAssociationsPaginator is a paginator for
// ListVpcEndpointAssociations
type ListVpcEndpointAssociationsPaginator struct {
	options   ListVpcEndpointAssociationsPaginatorOptions
	client    ListVpcEndpointAssociationsAPIClient
	params    *ListVpcEndpointAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListVpcEndpointAssociationsPaginator returns a new
// ListVpcEndpointAssociationsPaginator
func NewListVpcEndpointAssociationsPaginator(client ListVpcEndpointAssociationsAPIClient, params *ListVpcEndpointAssociationsInput, optFns ...func(*ListVpcEndpointAssociationsPaginatorOptions)) *ListVpcEndpointAssociationsPaginator {
	if params == nil {
		params = &ListVpcEndpointAssociationsInput{}
	}

	options := ListVpcEndpointAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVpcEndpointAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVpcEndpointAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVpcEndpointAssociations page.
func (p *ListVpcEndpointAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVpcEndpointAssociationsOutput, error) {
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
	result, err := p.client.ListVpcEndpointAssociations(ctx, &params, optFns...)
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

// ListVpcEndpointAssociationsAPIClient is a client that implements the
// ListVpcEndpointAssociations operation.
type ListVpcEndpointAssociationsAPIClient interface {
	ListVpcEndpointAssociations(context.Context, *ListVpcEndpointAssociationsInput, ...func(*Options)) (*ListVpcEndpointAssociationsOutput, error)
}

var _ ListVpcEndpointAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListVpcEndpointAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVpcEndpointAssociations",
	}
}
