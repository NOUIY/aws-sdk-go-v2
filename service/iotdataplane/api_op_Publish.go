// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Publishes an MQTT message.
//
// Requires permission to access the [Publish] action.
//
// For more information about MQTT messages, see [MQTT Protocol] in the IoT Developer Guide.
//
// For more information about messaging costs, see [Amazon Web Services IoT Core pricing - Messaging].
//
// [MQTT Protocol]: http://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html
// [Amazon Web Services IoT Core pricing - Messaging]: http://aws.amazon.com/iot-core/pricing/#Messaging
// [Publish]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) Publish(ctx context.Context, params *PublishInput, optFns ...func(*Options)) (*PublishOutput, error) {
	if params == nil {
		params = &PublishInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Publish", params, optFns, c.addOperationPublishMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the Publish operation.
type PublishInput struct {

	// The name of the MQTT topic.
	//
	// This member is required.
	Topic *string

	// A UTF-8 encoded string that describes the content of the publishing message.
	ContentType *string

	// The base64-encoded binary data used by the sender of the request message to
	// identify which request the response message is for when it's received.
	// correlationData is an HTTP header value in the API.
	CorrelationData *string

	// A user-defined integer value that represents the message expiry interval in
	// seconds. If absent, the message doesn't expire. For more information about the
	// limits of messageExpiry , see [Amazon Web Services IoT Core message broker and protocol limits and quotas] from the Amazon Web Services Reference Guide.
	//
	// [Amazon Web Services IoT Core message broker and protocol limits and quotas]: https://docs.aws.amazon.com/general/latest/gr/iot-core.html#message-broker-limits
	MessageExpiry int64

	// The message body. MQTT accepts text, binary, and empty (null) message payloads.
	//
	// Publishing an empty (null) payload with retain = true deletes the retained
	// message identified by topic from Amazon Web Services IoT Core.
	Payload []byte

	// An Enum string value that indicates whether the payload is formatted as UTF-8.
	// payloadFormatIndicator is an HTTP header value in the API.
	PayloadFormatIndicator types.PayloadFormatIndicator

	// The Quality of Service (QoS) level. The default QoS level is 0.
	Qos int32

	// A UTF-8 encoded string that's used as the topic name for a response message.
	// The response topic is used to describe the topic which the receiver should
	// publish to as part of the request-response flow. The topic must not contain
	// wildcard characters.
	ResponseTopic *string

	// A Boolean value that determines whether to set the RETAIN flag when the message
	// is published.
	//
	// Setting the RETAIN flag causes the message to be retained and sent to new
	// subscribers to the topic.
	//
	// Valid values: true | false
	//
	// Default value: false
	Retain bool

	// A JSON string that contains an array of JSON objects. If you don’t use Amazon
	// Web Services SDK or CLI, you must encode the JSON string to base64 format before
	// adding it to the HTTP header. userProperties is an HTTP header value in the API.
	//
	// The following example userProperties parameter is a JSON string which
	// represents two User Properties. Note that it needs to be base64-encoded:
	//
	//     [{"deviceName": "alpha"}, {"deviceCnt": "45"}]
	//
	// This value conforms to the media type: application/json
	UserProperties *string

	noSmithyDocumentSerde
}

type PublishOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublish{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublish{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Publish"); err != nil {
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
	if err = addOpPublishValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublish(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublish(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Publish",
	}
}
