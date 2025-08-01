// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves entries from a fleet's event log. Fleet events are initiated by
// changes in status, such as during fleet creation and termination, changes in
// capacity, etc. If a fleet has multiple locations, events are also initiated by
// changes to status and capacity in remote locations.
//
// You can specify a time range to limit the result set. Use the pagination
// parameters to retrieve results as a set of sequential pages.
//
// If successful, a collection of event log entries matching the request are
// returned.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
func (c *Client) DescribeFleetEvents(ctx context.Context, params *DescribeFleetEventsInput, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) {
	if params == nil {
		params = &DescribeFleetEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetEvents", params, optFns, c.addOperationDescribeFleetEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetEventsInput struct {

	// A unique identifier for the fleet to get event logs for. You can use either the
	// fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The most recent date to retrieve event logs for. If no end time is specified,
	// this call returns entries from the specified start time up to the present.
	// Format is a number expressed in Unix time as milliseconds (ex:
	// "1469498468.057").
	EndTime *time.Time

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// The earliest date to retrieve event logs for. If no start time is specified,
	// this call returns entries starting from when the fleet was created to the
	// specified end time. Format is a number expressed in Unix time as milliseconds
	// (ex: "1469498468.057").
	StartTime *time.Time

	noSmithyDocumentSerde
}

type DescribeFleetEventsOutput struct {

	// A collection of objects containing event log entries for the specified fleet.
	Events []types.Event

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetEvents"); err != nil {
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
	if err = addOpDescribeFleetEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetEvents(options.Region), middleware.Before); err != nil {
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

// DescribeFleetEventsPaginatorOptions is the paginator options for
// DescribeFleetEvents
type DescribeFleetEventsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetEventsPaginator is a paginator for DescribeFleetEvents
type DescribeFleetEventsPaginator struct {
	options   DescribeFleetEventsPaginatorOptions
	client    DescribeFleetEventsAPIClient
	params    *DescribeFleetEventsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetEventsPaginator returns a new DescribeFleetEventsPaginator
func NewDescribeFleetEventsPaginator(client DescribeFleetEventsAPIClient, params *DescribeFleetEventsInput, optFns ...func(*DescribeFleetEventsPaginatorOptions)) *DescribeFleetEventsPaginator {
	if params == nil {
		params = &DescribeFleetEventsInput{}
	}

	options := DescribeFleetEventsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetEvents page.
func (p *DescribeFleetEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeFleetEvents(ctx, &params, optFns...)
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

// DescribeFleetEventsAPIClient is a client that implements the
// DescribeFleetEvents operation.
type DescribeFleetEventsAPIClient interface {
	DescribeFleetEvents(context.Context, *DescribeFleetEventsInput, ...func(*Options)) (*DescribeFleetEventsOutput, error)
}

var _ DescribeFleetEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeFleetEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetEvents",
	}
}
