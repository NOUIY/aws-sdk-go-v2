// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new text message and sends it to a recipient's phone number.
// SendTextMessage only sends an SMS message to one recipient each time it is
// invoked.
//
// SMS throughput limits are measured in Message Parts per Second (MPS). Your MPS
// limit depends on the destination country of your messages, as well as the type
// of phone number (origination number) that you use to send the message. For more
// information about MPS, see [Message Parts per Second (MPS) limits]in the AWS End User Messaging SMS User Guide.
//
// [Message Parts per Second (MPS) limits]: https://docs.aws.amazon.com/sms-voice/latest/userguide/sms-limitations-mps.html
func (c *Client) SendTextMessage(ctx context.Context, params *SendTextMessageInput, optFns ...func(*Options)) (*SendTextMessageOutput, error) {
	if params == nil {
		params = &SendTextMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendTextMessage", params, optFns, c.addOperationSendTextMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendTextMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTextMessageInput struct {

	// The destination phone number in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The name of the configuration set to use. This can be either the
	// ConfigurationSetName or ConfigurationSetArn.
	ConfigurationSetName *string

	// You can specify custom data in this field. If you do, that data is logged to
	// the event destination.
	Context map[string]string

	// This field is used for any country-specific registration requirements.
	// Currently, this setting is only used when you send messages to recipients in
	// India using a sender ID. For more information see [Special requirements for sending SMS messages to recipients in India].
	//
	//   - IN_ENTITY_ID The entity ID or Principal Entity (PE) ID that you received
	//   after completing the sender ID registration process.
	//
	//   - IN_TEMPLATE_ID The template ID that you received after completing the sender
	//   ID registration process.
	//
	// Make sure that the Template ID that you specify matches your message template
	//   exactly. If your message doesn't match the template that you provided during the
	//   registration process, the mobile carriers might reject your message.
	//
	// [Special requirements for sending SMS messages to recipients in India]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-senderid-india.html
	DestinationCountryParameters map[string]string

	// When set to true, the message is checked and validated, but isn't sent to the
	// end recipient. You are not charged for using DryRun .
	//
	// The Message Parts per Second (MPS) limit when using DryRun is five. If your
	// origination identity has a lower MPS limit then the lower MPS limit is used. For
	// more information about MPS limits, see [Message Parts per Second (MPS) limits]in the AWS End User Messaging SMS User
	// Guide..
	//
	// [Message Parts per Second (MPS) limits]: https://docs.aws.amazon.com/sms-voice/latest/userguide/sms-limitations-mps.html
	DryRun bool

	// When you register a short code in the US, you must specify a program name. If
	// you don’t have a US short code, omit this attribute.
	Keyword *string

	// The maximum amount that you want to spend, in US dollars, per each text
	// message. If the calculated amount to send the text message is greater than
	// MaxPrice , the message is not sent and an error is returned.
	MaxPrice *string

	// The body of the text message.
	MessageBody *string

	// Set to true to enable message feedback for the message. When a user receives
	// the message you need to update the message status using PutMessageFeedback.
	MessageFeedbackEnabled *bool

	// The type of message. Valid values are for messages that are critical or
	// time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	MessageType types.MessageType

	// The origination identity of the message. This can be either the PhoneNumber,
	// PhoneNumberId, PhoneNumberArn, SenderId, SenderIdArn, PoolId, or PoolArn.
	//
	// If you are using a shared AWS End User Messaging SMS and Voice resource then
	// you must use the full Amazon Resource Name(ARN).
	OriginationIdentity *string

	// The unique identifier for the protect configuration.
	ProtectConfigurationId *string

	// How long the text message is valid for, in seconds. By default this is 72
	// hours. If the messages isn't handed off before the TTL expires we stop
	// attempting to hand off the message and return TTL_EXPIRED event.
	TimeToLive *int32

	noSmithyDocumentSerde
}

type SendTextMessageOutput struct {

	// The unique identifier for the message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendTextMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendTextMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendTextMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendTextMessage"); err != nil {
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
	if err = addOpSendTextMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendTextMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendTextMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendTextMessage",
	}
}
