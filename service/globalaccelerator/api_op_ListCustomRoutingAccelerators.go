// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the custom routing accelerators for an Amazon Web Services account.
func (c *Client) ListCustomRoutingAccelerators(ctx context.Context, params *ListCustomRoutingAcceleratorsInput, optFns ...func(*Options)) (*ListCustomRoutingAcceleratorsOutput, error) {
	if params == nil {
		params = &ListCustomRoutingAcceleratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomRoutingAccelerators", params, optFns, c.addOperationListCustomRoutingAcceleratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomRoutingAcceleratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomRoutingAcceleratorsInput struct {

	// The number of custom routing Global Accelerator objects that you want to return
	// with this call. The default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomRoutingAcceleratorsOutput struct {

	// The list of custom routing accelerators for a customer account.
	Accelerators []types.CustomRoutingAccelerator

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomRoutingAcceleratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomRoutingAccelerators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomRoutingAccelerators{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomRoutingAccelerators"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomRoutingAccelerators(options.Region), middleware.Before); err != nil {
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

// ListCustomRoutingAcceleratorsPaginatorOptions is the paginator options for
// ListCustomRoutingAccelerators
type ListCustomRoutingAcceleratorsPaginatorOptions struct {
	// The number of custom routing Global Accelerator objects that you want to return
	// with this call. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomRoutingAcceleratorsPaginator is a paginator for
// ListCustomRoutingAccelerators
type ListCustomRoutingAcceleratorsPaginator struct {
	options   ListCustomRoutingAcceleratorsPaginatorOptions
	client    ListCustomRoutingAcceleratorsAPIClient
	params    *ListCustomRoutingAcceleratorsInput
	nextToken *string
	firstPage bool
}

// NewListCustomRoutingAcceleratorsPaginator returns a new
// ListCustomRoutingAcceleratorsPaginator
func NewListCustomRoutingAcceleratorsPaginator(client ListCustomRoutingAcceleratorsAPIClient, params *ListCustomRoutingAcceleratorsInput, optFns ...func(*ListCustomRoutingAcceleratorsPaginatorOptions)) *ListCustomRoutingAcceleratorsPaginator {
	if params == nil {
		params = &ListCustomRoutingAcceleratorsInput{}
	}

	options := ListCustomRoutingAcceleratorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomRoutingAcceleratorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomRoutingAcceleratorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomRoutingAccelerators page.
func (p *ListCustomRoutingAcceleratorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomRoutingAcceleratorsOutput, error) {
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
	result, err := p.client.ListCustomRoutingAccelerators(ctx, &params, optFns...)
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

// ListCustomRoutingAcceleratorsAPIClient is a client that implements the
// ListCustomRoutingAccelerators operation.
type ListCustomRoutingAcceleratorsAPIClient interface {
	ListCustomRoutingAccelerators(context.Context, *ListCustomRoutingAcceleratorsInput, ...func(*Options)) (*ListCustomRoutingAcceleratorsOutput, error)
}

var _ ListCustomRoutingAcceleratorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCustomRoutingAccelerators(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomRoutingAccelerators",
	}
}
