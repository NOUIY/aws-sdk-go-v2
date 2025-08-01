// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified time series form for the specified asset.
func (c *Client) DeleteTimeSeriesDataPoints(ctx context.Context, params *DeleteTimeSeriesDataPointsInput, optFns ...func(*Options)) (*DeleteTimeSeriesDataPointsOutput, error) {
	if params == nil {
		params = &DeleteTimeSeriesDataPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTimeSeriesDataPoints", params, optFns, c.addOperationDeleteTimeSeriesDataPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTimeSeriesDataPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTimeSeriesDataPointsInput struct {

	// The ID of the Amazon DataZone domain that houses the asset for which you want
	// to delete a time series form.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the asset for which you want to delete a time series form.
	//
	// This member is required.
	EntityIdentifier *string

	// The type of the asset for which you want to delete a time series form.
	//
	// This member is required.
	EntityType types.TimeSeriesEntityType

	// The name of the time series form that you want to delete.
	//
	// This member is required.
	FormName *string

	// A unique, case-sensitive identifier to ensure idempotency of the request. This
	// field is automatically populated if not provided.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteTimeSeriesDataPointsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTimeSeriesDataPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTimeSeriesDataPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTimeSeriesDataPoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTimeSeriesDataPoints"); err != nil {
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
	if err = addIdempotencyToken_opDeleteTimeSeriesDataPointsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteTimeSeriesDataPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTimeSeriesDataPoints(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteTimeSeriesDataPoints struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteTimeSeriesDataPoints) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteTimeSeriesDataPoints) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteTimeSeriesDataPointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteTimeSeriesDataPointsInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteTimeSeriesDataPointsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteTimeSeriesDataPoints{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteTimeSeriesDataPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTimeSeriesDataPoints",
	}
}
