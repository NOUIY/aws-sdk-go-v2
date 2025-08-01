// Code generated by smithy-go-codegen DO NOT EDIT.

package iotmanagedintegrations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Relays third-party device events for a connector such as a new device or a
// device state change event.
func (c *Client) SendConnectorEvent(ctx context.Context, params *SendConnectorEventInput, optFns ...func(*Options)) (*SendConnectorEventOutput, error) {
	if params == nil {
		params = &SendConnectorEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendConnectorEvent", params, optFns, c.addOperationSendConnectorEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendConnectorEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendConnectorEventInput struct {

	// The id of the connector between the third-party cloud provider and IoT managed
	// integrations.
	//
	// This member is required.
	ConnectorId *string

	// The Open Connectivity Foundation (OCF) operation requested to be performed on
	// the managed thing.
	//
	// The field op can have a value of "I" or "U". The field "cn" will contain the
	// capability types.
	//
	// This member is required.
	Operation types.ConnectorEventOperation

	// The third-party device id as defined by the connector. This device id must not
	// contain personal identifiable information (PII).
	//
	// This parameter is used for cloud-to-cloud devices only.
	ConnectorDeviceId *string

	// The id for the device discovery job.
	DeviceDiscoveryId *string

	// The list of devices.
	Devices []types.Device

	// The device endpoint.
	MatterEndpoint *types.MatterEndpoint

	// The device state change event payload.
	//
	// This parameter will include the following three fields:
	//
	//   - uri : schema auc://<PARTNER-DEVICE-ID>/ResourcePath (The Resourcepath
	//   corresponds to an OCF resource.)
	//
	//   - op : For device state changes, this field must populate as n+d .
	//
	//   - cn : The content depends on the OCF resource referenced in ResourcePath .
	Message *string

	// The Open Connectivity Foundation (OCF) security specification version for the
	// operation being requested on the managed thing. For more information, see [OCF Security Specification].
	//
	// [OCF Security Specification]: https://openconnectivity.org/specs/OCF_Security_Specification_v1.0.0.pdf
	OperationVersion *string

	// The status code of the Open Connectivity Foundation (OCF) operation being
	// performed on the managed thing.
	StatusCode *int32

	// The trace request identifier used to correlate a command request and response.
	// This is specified by the device owner, but will be generated by IoT managed
	// integrations if not provided by the device owner.
	TraceId *string

	// The id of the third-party cloud provider.
	UserId *string

	noSmithyDocumentSerde
}

type SendConnectorEventOutput struct {

	// The id of the connector between the third-party cloud provider and IoT managed
	// integrations.
	//
	// This member is required.
	ConnectorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendConnectorEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendConnectorEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendConnectorEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendConnectorEvent"); err != nil {
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
	if err = addOpSendConnectorEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendConnectorEvent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendConnectorEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendConnectorEvent",
	}
}
