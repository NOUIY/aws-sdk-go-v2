// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a time series (data stream).
//
// To identify a time series, do one of the following:
//
//   - If the time series isn't associated with an asset property, specify the
//     alias of the time series.
//
//   - If the time series is associated with an asset property, specify one of the
//     following:
//
//   - The alias of the time series.
//
//   - The assetId and propertyId that identifies the asset property.
func (c *Client) DescribeTimeSeries(ctx context.Context, params *DescribeTimeSeriesInput, optFns ...func(*Options)) (*DescribeTimeSeriesOutput, error) {
	if params == nil {
		params = &DescribeTimeSeriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTimeSeries", params, optFns, c.addOperationDescribeTimeSeriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTimeSeriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTimeSeriesInput struct {

	// The alias that identifies the time series.
	Alias *string

	// The ID of the asset in which the asset property was created. This can be either
	// the actual ID in UUID format, or else externalId: followed by the external ID,
	// if it has one. For more information, see [Referencing objects with external IDs]in the IoT SiteWise User Guide.
	//
	// [Referencing objects with external IDs]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references
	AssetId *string

	// The ID of the asset property. This can be either the actual ID in UUID format,
	// or else externalId: followed by the external ID, if it has one. For more
	// information, see [Referencing objects with external IDs]in the IoT SiteWise User Guide.
	//
	// [Referencing objects with external IDs]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references
	PropertyId *string

	noSmithyDocumentSerde
}

type DescribeTimeSeriesOutput struct {

	// The data type of the time series.
	//
	// If you specify STRUCT , you must also specify dataTypeSpec to identify the type
	// of the structure for this time series.
	//
	// This member is required.
	DataType types.PropertyDataType

	// The [ARN] of the time series, which has the following format.
	//
	//     arn:${Partition}:iotsitewise:${Region}:${Account}:time-series/${TimeSeriesId}
	//
	// [ARN]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	TimeSeriesArn *string

	// The date that the time series was created, in Unix epoch time.
	//
	// This member is required.
	TimeSeriesCreationDate *time.Time

	// The ID of the time series.
	//
	// This member is required.
	TimeSeriesId *string

	// The date that the time series was last updated, in Unix epoch time.
	//
	// This member is required.
	TimeSeriesLastUpdateDate *time.Time

	// The alias that identifies the time series.
	Alias *string

	// The ID of the asset in which the asset property was created.
	AssetId *string

	// The data type of the structure for this time series. This parameter is required
	// for time series that have the STRUCT data type.
	//
	// The options for this parameter depend on the type of the composite model in
	// which you created the asset property that is associated with your time series.
	// Use AWS/ALARM_STATE for alarm state in alarm composite models.
	DataTypeSpec *string

	// The ID of the asset property, in UUID format.
	PropertyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTimeSeriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTimeSeries"); err != nil {
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
	if err = addEndpointPrefix_opDescribeTimeSeriesMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTimeSeries(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeTimeSeriesMiddleware struct {
}

func (*endpointPrefix_opDescribeTimeSeriesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeTimeSeriesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribeTimeSeriesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeTimeSeriesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTimeSeries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTimeSeries",
	}
}
