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

// Returns a list of virtual tape recovery points that are available for the
// specified tape gateway.
//
// A recovery point is a point-in-time view of a virtual tape at which all the
// data on the virtual tape is consistent. If your gateway crashes, virtual tapes
// that have recovery points can be recovered to a new gateway. This operation is
// only supported in the tape gateway type.
func (c *Client) DescribeTapeRecoveryPoints(ctx context.Context, params *DescribeTapeRecoveryPointsInput, optFns ...func(*Options)) (*DescribeTapeRecoveryPointsOutput, error) {
	if params == nil {
		params = &DescribeTapeRecoveryPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTapeRecoveryPoints", params, optFns, c.addOperationDescribeTapeRecoveryPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTapeRecoveryPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapeRecoveryPointsInput
type DescribeTapeRecoveryPointsInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies that the number of virtual tape recovery points that are described be
	// limited to the specified number.
	Limit *int32

	// An opaque string that indicates the position at which to begin describing the
	// virtual tape recovery points.
	Marker *string

	noSmithyDocumentSerde
}

// DescribeTapeRecoveryPointsOutput
type DescribeTapeRecoveryPointsOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// An opaque string that indicates the position at which the virtual tape recovery
	// points that were listed for description ended.
	//
	// Use this marker in your next request to list the next set of virtual tape
	// recovery points in the list. If there are no more recovery points to describe,
	// this field does not appear in the response.
	Marker *string

	// An array of TapeRecoveryPointInfos that are available for the specified gateway.
	TapeRecoveryPointInfos []types.TapeRecoveryPointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTapeRecoveryPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapeRecoveryPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapeRecoveryPoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTapeRecoveryPoints"); err != nil {
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
	if err = addOpDescribeTapeRecoveryPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapeRecoveryPoints(options.Region), middleware.Before); err != nil {
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

// DescribeTapeRecoveryPointsPaginatorOptions is the paginator options for
// DescribeTapeRecoveryPoints
type DescribeTapeRecoveryPointsPaginatorOptions struct {
	// Specifies that the number of virtual tape recovery points that are described be
	// limited to the specified number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTapeRecoveryPointsPaginator is a paginator for
// DescribeTapeRecoveryPoints
type DescribeTapeRecoveryPointsPaginator struct {
	options   DescribeTapeRecoveryPointsPaginatorOptions
	client    DescribeTapeRecoveryPointsAPIClient
	params    *DescribeTapeRecoveryPointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTapeRecoveryPointsPaginator returns a new
// DescribeTapeRecoveryPointsPaginator
func NewDescribeTapeRecoveryPointsPaginator(client DescribeTapeRecoveryPointsAPIClient, params *DescribeTapeRecoveryPointsInput, optFns ...func(*DescribeTapeRecoveryPointsPaginatorOptions)) *DescribeTapeRecoveryPointsPaginator {
	if params == nil {
		params = &DescribeTapeRecoveryPointsInput{}
	}

	options := DescribeTapeRecoveryPointsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTapeRecoveryPointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTapeRecoveryPointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTapeRecoveryPoints page.
func (p *DescribeTapeRecoveryPointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTapeRecoveryPointsOutput, error) {
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
	result, err := p.client.DescribeTapeRecoveryPoints(ctx, &params, optFns...)
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

// DescribeTapeRecoveryPointsAPIClient is a client that implements the
// DescribeTapeRecoveryPoints operation.
type DescribeTapeRecoveryPointsAPIClient interface {
	DescribeTapeRecoveryPoints(context.Context, *DescribeTapeRecoveryPointsInput, ...func(*Options)) (*DescribeTapeRecoveryPointsOutput, error)
}

var _ DescribeTapeRecoveryPointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTapeRecoveryPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTapeRecoveryPoints",
	}
}
