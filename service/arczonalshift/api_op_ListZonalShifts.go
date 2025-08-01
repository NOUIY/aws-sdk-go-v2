// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all active and completed zonal shifts in Amazon Application Recovery
// Controller in your Amazon Web Services account in this Amazon Web Services
// Region. ListZonalShifts returns customer-initiated zonal shifts, as well as
// practice run zonal shifts that ARC started on your behalf for zonal autoshift.
//
// For more information about listing autoshifts, see [">ListAutoshifts].
//
// [">ListAutoshifts]: https://docs.aws.amazon.com/arc-zonal-shift/latest/api/API_ListAutoshifts.html
func (c *Client) ListZonalShifts(ctx context.Context, params *ListZonalShiftsInput, optFns ...func(*Options)) (*ListZonalShiftsOutput, error) {
	if params == nil {
		params = &ListZonalShiftsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListZonalShifts", params, optFns, c.addOperationListZonalShiftsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListZonalShiftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListZonalShiftsInput struct {

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a nextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's nextToken response to request the next page of results.
	NextToken *string

	// The identifier for the resource that you want to list zonal shifts for. The
	// identifier is the Amazon Resource Name (ARN) for the resource.
	ResourceIdentifier *string

	// A status for a zonal shift.
	//
	// The Status for a zonal shift can have one of the following values:
	//
	//   - ACTIVE: The zonal shift has been started and is active.
	//
	//   - EXPIRED: The zonal shift has expired (the expiry time was exceeded).
	//
	//   - CANCELED: The zonal shift was canceled.
	Status types.ZonalShiftStatus

	noSmithyDocumentSerde
}

type ListZonalShiftsOutput struct {

	// The items in the response list.
	Items []types.ZonalShiftSummary

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a nextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's nextToken response to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListZonalShiftsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListZonalShifts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListZonalShifts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListZonalShifts"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListZonalShifts(options.Region), middleware.Before); err != nil {
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

// ListZonalShiftsPaginatorOptions is the paginator options for ListZonalShifts
type ListZonalShiftsPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListZonalShiftsPaginator is a paginator for ListZonalShifts
type ListZonalShiftsPaginator struct {
	options   ListZonalShiftsPaginatorOptions
	client    ListZonalShiftsAPIClient
	params    *ListZonalShiftsInput
	nextToken *string
	firstPage bool
}

// NewListZonalShiftsPaginator returns a new ListZonalShiftsPaginator
func NewListZonalShiftsPaginator(client ListZonalShiftsAPIClient, params *ListZonalShiftsInput, optFns ...func(*ListZonalShiftsPaginatorOptions)) *ListZonalShiftsPaginator {
	if params == nil {
		params = &ListZonalShiftsInput{}
	}

	options := ListZonalShiftsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListZonalShiftsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListZonalShiftsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListZonalShifts page.
func (p *ListZonalShiftsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListZonalShiftsOutput, error) {
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
	result, err := p.client.ListZonalShifts(ctx, &params, optFns...)
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

// ListZonalShiftsAPIClient is a client that implements the ListZonalShifts
// operation.
type ListZonalShiftsAPIClient interface {
	ListZonalShifts(context.Context, *ListZonalShiftsInput, ...func(*Options)) (*ListZonalShiftsOutput, error)
}

var _ ListZonalShiftsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListZonalShifts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListZonalShifts",
	}
}
