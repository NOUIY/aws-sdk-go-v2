// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of virtual tapes that correspond to the specified Amazon
// Resource Names (ARNs). If TapeARN is not specified, returns a description of
// the virtual tapes associated with the specified gateway. This operation is only
// supported for the tape gateway type.
//
// The operation supports pagination. By default, the operation returns a maximum
// of up to 100 tapes. You can optionally specify the Limit field in the body to
// limit the number of tapes in the response. If the number of tapes returned in
// the response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of tapes.
func (c *Client) DescribeTapes(ctx context.Context, params *DescribeTapesInput, optFns ...func(*Options)) (*DescribeTapesOutput, error) {
	if params == nil {
		params = &DescribeTapesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTapes", params, optFns, c.addOperationDescribeTapesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapesInput
type DescribeTapesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies that the number of virtual tapes described be limited to the
	// specified number.
	//
	// Amazon Web Services may impose its own limit, if this field is not set.
	Limit *int32

	// A marker value, obtained in a previous call to DescribeTapes . This marker
	// indicates which page of results to retrieve.
	//
	// If not specified, the first page of results is retrieved.
	Marker *string

	// Specifies one or more unique Amazon Resource Names (ARNs) that represent the
	// virtual tapes you want to describe. If this parameter is not specified, Tape
	// gateway returns a description of all virtual tapes associated with the specified
	// gateway.
	TapeARNs []string

	noSmithyDocumentSerde
}

// DescribeTapesOutput
type DescribeTapesOutput struct {

	// An opaque string that can be used as part of a subsequent DescribeTapes call to
	// retrieve the next page of results.
	//
	// If a response does not contain a marker, then there are no more results to be
	// retrieved.
	Marker *string

	// An array of virtual tape descriptions.
	Tapes []types.Tape

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTapesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTapes"); err != nil {
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
	if err = addOpDescribeTapesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapes(options.Region), middleware.Before); err != nil {
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

// DescribeTapesPaginatorOptions is the paginator options for DescribeTapes
type DescribeTapesPaginatorOptions struct {
	// Specifies that the number of virtual tapes described be limited to the
	// specified number.
	//
	// Amazon Web Services may impose its own limit, if this field is not set.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTapesPaginator is a paginator for DescribeTapes
type DescribeTapesPaginator struct {
	options   DescribeTapesPaginatorOptions
	client    DescribeTapesAPIClient
	params    *DescribeTapesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTapesPaginator returns a new DescribeTapesPaginator
func NewDescribeTapesPaginator(client DescribeTapesAPIClient, params *DescribeTapesInput, optFns ...func(*DescribeTapesPaginatorOptions)) *DescribeTapesPaginator {
	if params == nil {
		params = &DescribeTapesInput{}
	}

	options := DescribeTapesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTapesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTapesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTapes page.
func (p *DescribeTapesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTapesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeTapes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeTapesAPIClient is a client that implements the DescribeTapes operation.
type DescribeTapesAPIClient interface {
	DescribeTapes(context.Context, *DescribeTapesInput, ...func(*Options)) (*DescribeTapesOutput, error)
}

var _ DescribeTapesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTapes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTapes",
	}
}
