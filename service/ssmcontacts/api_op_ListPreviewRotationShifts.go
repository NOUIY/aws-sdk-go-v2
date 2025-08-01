// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of shifts based on rotation configuration parameters.
//
// The Incident Manager primarily uses this operation to populate the Preview
// calendar. It is not typically run by end users.
func (c *Client) ListPreviewRotationShifts(ctx context.Context, params *ListPreviewRotationShiftsInput, optFns ...func(*Options)) (*ListPreviewRotationShiftsOutput, error) {
	if params == nil {
		params = &ListPreviewRotationShiftsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPreviewRotationShifts", params, optFns, c.addOperationListPreviewRotationShiftsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPreviewRotationShiftsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPreviewRotationShiftsInput struct {

	// The date and time a rotation shift would end.
	//
	// This member is required.
	EndTime *time.Time

	// The contacts that would be assigned to a rotation.
	//
	// This member is required.
	Members []string

	// Information about how long a rotation would last before restarting at the
	// beginning of the shift order.
	//
	// This member is required.
	Recurrence *types.RecurrenceSettings

	// The time zone the rotation’s activity would be based on, in Internet Assigned
	// Numbers Authority (IANA) format. For example: "America/Los_Angeles", "UTC", or
	// "Asia/Seoul".
	//
	// This member is required.
	TimeZoneId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that can be specified in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. This token is used to get the next set of results.
	NextToken *string

	// Information about changes that would be made in a rotation override.
	Overrides []types.PreviewOverride

	// The date and time a rotation would begin. The first shift is calculated from
	// this date and time.
	RotationStartTime *time.Time

	// Used to filter the range of calculated shifts before sending the response back
	// to the user.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListPreviewRotationShiftsOutput struct {

	// The token for the next set of items to return. This token is used to get the
	// next set of results.
	NextToken *string

	// Details about a rotation shift, including times, types, and contacts.
	RotationShifts []types.RotationShift

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPreviewRotationShiftsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPreviewRotationShifts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPreviewRotationShifts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPreviewRotationShifts"); err != nil {
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
	if err = addOpListPreviewRotationShiftsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPreviewRotationShifts(options.Region), middleware.Before); err != nil {
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

// ListPreviewRotationShiftsPaginatorOptions is the paginator options for
// ListPreviewRotationShifts
type ListPreviewRotationShiftsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that can be specified in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPreviewRotationShiftsPaginator is a paginator for ListPreviewRotationShifts
type ListPreviewRotationShiftsPaginator struct {
	options   ListPreviewRotationShiftsPaginatorOptions
	client    ListPreviewRotationShiftsAPIClient
	params    *ListPreviewRotationShiftsInput
	nextToken *string
	firstPage bool
}

// NewListPreviewRotationShiftsPaginator returns a new
// ListPreviewRotationShiftsPaginator
func NewListPreviewRotationShiftsPaginator(client ListPreviewRotationShiftsAPIClient, params *ListPreviewRotationShiftsInput, optFns ...func(*ListPreviewRotationShiftsPaginatorOptions)) *ListPreviewRotationShiftsPaginator {
	if params == nil {
		params = &ListPreviewRotationShiftsInput{}
	}

	options := ListPreviewRotationShiftsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPreviewRotationShiftsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPreviewRotationShiftsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPreviewRotationShifts page.
func (p *ListPreviewRotationShiftsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPreviewRotationShiftsOutput, error) {
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
	result, err := p.client.ListPreviewRotationShifts(ctx, &params, optFns...)
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

// ListPreviewRotationShiftsAPIClient is a client that implements the
// ListPreviewRotationShifts operation.
type ListPreviewRotationShiftsAPIClient interface {
	ListPreviewRotationShifts(context.Context, *ListPreviewRotationShiftsInput, ...func(*Options)) (*ListPreviewRotationShiftsOutput, error)
}

var _ ListPreviewRotationShiftsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPreviewRotationShifts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPreviewRotationShifts",
	}
}
