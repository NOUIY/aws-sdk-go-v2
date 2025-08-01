// Code generated by smithy-go-codegen DO NOT EDIT.

package georoutes

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/georoutes/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Use CalculateRouteMatrix to compute results for all pairs of Origins to
//
// Destinations. Each row corresponds to one entry in Origins. Each entry in the
// row corresponds to the route from that entry in Origins to an entry in
// Destinations positions.
func (c *Client) CalculateRouteMatrix(ctx context.Context, params *CalculateRouteMatrixInput, optFns ...func(*Options)) (*CalculateRouteMatrixOutput, error) {
	if params == nil {
		params = &CalculateRouteMatrixInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CalculateRouteMatrix", params, optFns, c.addOperationCalculateRouteMatrixMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CalculateRouteMatrixOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CalculateRouteMatrixInput struct {

	// List of destinations for the route.
	//
	// Route calculations are billed for each origin and destination pair. If you use
	// a large matrix of origins and destinations, your costs will increase
	// accordingly. See [Amazon Location's pricing page]for more information.
	//
	// [Amazon Location's pricing page]: https://docs.aws.amazon.com/location/latest/developerguide/routes-pricing.html`
	//
	// This member is required.
	Destinations []types.RouteMatrixDestination

	// The position in longitude and latitude for the origin.
	//
	// Route calculations are billed for each origin and destination pair. Using a
	// large amount of Origins in a request can lead you to incur unexpected charges.
	// See [Amazon Location's pricing page]for more information.
	//
	// [Amazon Location's pricing page]: https://docs.aws.amazon.com/location/latest/developerguide/routes-pricing.html`
	//
	// This member is required.
	Origins []types.RouteMatrixOrigin

	// Boundary within which the matrix is to be calculated. All data, origins and
	// destinations outside the boundary are considered invalid.
	//
	// When request routing boundary was set as AutoCircle, the response routing
	// boundary will return Circle derived from the AutoCircle settings.
	//
	// This member is required.
	RoutingBoundary *types.RouteMatrixBoundary

	// Features that are allowed while calculating a route.
	Allow *types.RouteMatrixAllowOptions

	// Features that are avoided while calculating a route. Avoidance is on a
	// best-case basis. If an avoidance can't be satisfied for a particular case, it
	// violates the avoidance and the returned response produces a notice for the
	// violation.
	Avoid *types.RouteMatrixAvoidanceOptions

	// Uses the current time as the time of departure.
	DepartNow *bool

	// Time of departure from thr origin.
	//
	// Time format: YYYY-MM-DDThh:mm:ss.sssZ | YYYY-MM-DDThh:mm:ss.sss+hh:mm
	//
	// Examples:
	//
	//     2020-04-22T17:57:24Z
	//
	//     2020-04-22T17:57:24+02:00
	DepartureTime *string

	// Features to be strictly excluded while calculating the route.
	Exclude *types.RouteMatrixExclusionOptions

	// Optional: The API key to be used for authorization. Either an API key or valid
	// SigV4 signature must be provided when making a request.
	Key *string

	// Specifies the optimization criteria for calculating a route.
	//
	// Default Value: FastestRoute
	OptimizeRoutingFor types.RoutingObjective

	// Traffic related options.
	Traffic *types.RouteMatrixTrafficOptions

	// Specifies the mode of transport when calculating a route. Used in estimating
	// the speed of travel and road compatibility.
	//
	// Default Value: Car
	TravelMode types.RouteMatrixTravelMode

	// Travel mode related options for the provided travel mode.
	TravelModeOptions *types.RouteMatrixTravelModeOptions

	noSmithyDocumentSerde
}

type CalculateRouteMatrixOutput struct {

	// The count of error results in the route matrix. If this number is 0, all routes
	// were calculated successfully.
	//
	// This member is required.
	ErrorCount *int32

	// The pricing bucket for which the query is charged at.
	//
	// This member is required.
	PricingBucket *string

	// The calculated route matrix containing the results for all pairs of Origins to
	// Destination positions. Each row corresponds to one entry in Origins. Each entry
	// in the row corresponds to the route from that entry in Origins to an entry in
	// Destination positions.
	//
	// This member is required.
	RouteMatrix [][]types.RouteMatrixEntry

	// Boundary within which the matrix is to be calculated. All data, origins and
	// destinations outside the boundary are considered invalid.
	//
	// When request routing boundary was set as AutoCircle, the response routing
	// boundary will return Circle derived from the AutoCircle settings.
	//
	// This member is required.
	RoutingBoundary *types.RouteMatrixBoundary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCalculateRouteMatrixMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCalculateRouteMatrix{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCalculateRouteMatrix{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CalculateRouteMatrix"); err != nil {
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
	if err = addOpCalculateRouteMatrixValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCalculateRouteMatrix(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCalculateRouteMatrix(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CalculateRouteMatrix",
	}
}
