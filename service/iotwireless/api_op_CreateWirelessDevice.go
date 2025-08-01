// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provisions a wireless device.
func (c *Client) CreateWirelessDevice(ctx context.Context, params *CreateWirelessDeviceInput, optFns ...func(*Options)) (*CreateWirelessDeviceOutput, error) {
	if params == nil {
		params = &CreateWirelessDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWirelessDevice", params, optFns, c.addOperationCreateWirelessDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWirelessDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWirelessDeviceInput struct {

	// The name of the destination to assign to the new wireless device.
	//
	// This member is required.
	DestinationName *string

	// The wireless device type.
	//
	// This member is required.
	Type types.WirelessDeviceType

	// Each resource must have a unique client request token. The client token is used
	// to implement idempotency. It ensures that the request completes no more than one
	// time. If you retry a request with the same token and the same parameters, the
	// request will complete successfully. However, if you try to create a new resource
	// using the same token but different parameters, an HTTP 409 conflict occurs. If
	// you omit this value, AWS SDKs will automatically generate a unique client
	// request. For more information about idempotency, see [Ensuring idempotency in Amazon EC2 API requests].
	//
	// [Ensuring idempotency in Amazon EC2 API requests]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
	ClientRequestToken *string

	// The description of the new resource.
	Description *string

	// The device configuration information to use to create the wireless device.
	LoRaWAN *types.LoRaWANDevice

	// The name of the new resource.
	//
	// The following special characters aren't accepted: <>^#~$
	Name *string

	// FPort values for the GNSS, stream, and ClockSync functions of the positioning
	// information.
	Positioning types.PositioningConfigStatus

	// The device configuration information to use to create the Sidewalk device.
	Sidewalk *types.SidewalkCreateWirelessDevice

	// The tags to attach to the new wireless device. Tags are metadata that you can
	// use to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWirelessDeviceOutput struct {

	// The Amazon Resource Name of the new resource.
	Arn *string

	// The ID of the new wireless device.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWirelessDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWirelessDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWirelessDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWirelessDevice"); err != nil {
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
	if err = addIdempotencyToken_opCreateWirelessDeviceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWirelessDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWirelessDevice(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWirelessDevice struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWirelessDevice) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWirelessDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWirelessDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWirelessDeviceInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWirelessDeviceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWirelessDevice{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWirelessDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWirelessDevice",
	}
}
