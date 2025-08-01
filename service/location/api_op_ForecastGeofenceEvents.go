// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action forecasts future geofence events that are likely to occur within a
// specified time horizon if a device continues moving at its current speed. Each
// forecasted event is associated with a geofence from a provided geofence
// collection. A forecast event can have one of the following states:
//
// ENTER : The device position is outside the referenced geofence, but the device
// may cross into the geofence during the forecasting time horizon if it maintains
// its current speed.
//
// EXIT : The device position is inside the referenced geofence, but the device may
// leave the geofence during the forecasted time horizon if the device maintains
// it's current speed.
//
// IDLE :The device is inside the geofence, and it will remain inside the geofence
// through the end of the time horizon if the device maintains it's current speed.
//
// Heading direction is not considered in the current version. The API takes a
// conservative approach and includes events that can occur for any heading.
func (c *Client) ForecastGeofenceEvents(ctx context.Context, params *ForecastGeofenceEventsInput, optFns ...func(*Options)) (*ForecastGeofenceEventsOutput, error) {
	if params == nil {
		params = &ForecastGeofenceEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ForecastGeofenceEvents", params, optFns, c.addOperationForecastGeofenceEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ForecastGeofenceEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ForecastGeofenceEventsInput struct {

	// The name of the geofence collection.
	//
	// This member is required.
	CollectionName *string

	// Represents the device's state, including its current position and speed. When
	// speed is omitted, this API performs a containment check. The containment check
	// operation returns IDLE events for geofences where the device is currently
	// inside of, but no other events.
	//
	// This member is required.
	DeviceState *types.ForecastGeofenceEventsDeviceState

	// The distance unit used for the NearestDistance property returned in a
	// forecasted event. The measurement system must match for DistanceUnit and
	// SpeedUnit ; if Kilometers is specified for DistanceUnit , then SpeedUnit must
	// be KilometersPerHour .
	//
	// Default Value: Kilometers
	DistanceUnit types.DistanceUnit

	// An optional limit for the number of resources returned in a single call.
	//
	// Default value: 20
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page.
	//
	// Default value: null
	NextToken *string

	// The speed unit for the device captured by the device state. The measurement
	// system must match for DistanceUnit and SpeedUnit ; if Kilometers is specified
	// for DistanceUnit , then SpeedUnit must be KilometersPerHour .
	//
	// Default Value: KilometersPerHour .
	SpeedUnit types.SpeedUnit

	// The forward-looking time window for forecasting, specified in minutes. The API
	// only returns events that are predicted to occur within this time horizon. When
	// no value is specified, this API performs a containment check. The containment
	// check operation returns IDLE events for geofences where the device is currently
	// inside of, but no other events.
	TimeHorizonMinutes *float64

	noSmithyDocumentSerde
}

type ForecastGeofenceEventsOutput struct {

	// The distance unit for the forecasted events.
	//
	// This member is required.
	DistanceUnit types.DistanceUnit

	// The list of forecasted events.
	//
	// This member is required.
	ForecastedEvents []types.ForecastedEvent

	// The speed unit for the forecasted events.
	//
	// This member is required.
	SpeedUnit types.SpeedUnit

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationForecastGeofenceEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpForecastGeofenceEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpForecastGeofenceEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ForecastGeofenceEvents"); err != nil {
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
	if err = addEndpointPrefix_opForecastGeofenceEventsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpForecastGeofenceEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opForecastGeofenceEvents(options.Region), middleware.Before); err != nil {
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

// ForecastGeofenceEventsPaginatorOptions is the paginator options for
// ForecastGeofenceEvents
type ForecastGeofenceEventsPaginatorOptions struct {
	// An optional limit for the number of resources returned in a single call.
	//
	// Default value: 20
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ForecastGeofenceEventsPaginator is a paginator for ForecastGeofenceEvents
type ForecastGeofenceEventsPaginator struct {
	options   ForecastGeofenceEventsPaginatorOptions
	client    ForecastGeofenceEventsAPIClient
	params    *ForecastGeofenceEventsInput
	nextToken *string
	firstPage bool
}

// NewForecastGeofenceEventsPaginator returns a new ForecastGeofenceEventsPaginator
func NewForecastGeofenceEventsPaginator(client ForecastGeofenceEventsAPIClient, params *ForecastGeofenceEventsInput, optFns ...func(*ForecastGeofenceEventsPaginatorOptions)) *ForecastGeofenceEventsPaginator {
	if params == nil {
		params = &ForecastGeofenceEventsInput{}
	}

	options := ForecastGeofenceEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ForecastGeofenceEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ForecastGeofenceEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ForecastGeofenceEvents page.
func (p *ForecastGeofenceEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ForecastGeofenceEventsOutput, error) {
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
	result, err := p.client.ForecastGeofenceEvents(ctx, &params, optFns...)
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

type endpointPrefix_opForecastGeofenceEventsMiddleware struct {
}

func (*endpointPrefix_opForecastGeofenceEventsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opForecastGeofenceEventsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opForecastGeofenceEventsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opForecastGeofenceEventsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ForecastGeofenceEventsAPIClient is a client that implements the
// ForecastGeofenceEvents operation.
type ForecastGeofenceEventsAPIClient interface {
	ForecastGeofenceEvents(context.Context, *ForecastGeofenceEventsInput, ...func(*Options)) (*ForecastGeofenceEventsOutput, error)
}

var _ ForecastGeofenceEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opForecastGeofenceEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ForecastGeofenceEvents",
	}
}
