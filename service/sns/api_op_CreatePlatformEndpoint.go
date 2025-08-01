// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS.
// CreatePlatformEndpoint requires the PlatformApplicationArn that is returned
// from CreatePlatformApplication . You can use the returned EndpointArn to send a
// message to a mobile app or by the Subscribe action for subscription to a topic.
// The CreatePlatformEndpoint action is idempotent, so if the requester already
// owns an endpoint with the same device token and attributes, that endpoint's ARN
// is returned without creating a new endpoint. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// When using CreatePlatformEndpoint with Baidu, two attributes must be provided:
// ChannelId and UserId. The token field must also contain the ChannelId. For more
// information, see [Creating an Amazon SNS Endpoint for Baidu].
//
// [Creating an Amazon SNS Endpoint for Baidu]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func (c *Client) CreatePlatformEndpoint(ctx context.Context, params *CreatePlatformEndpointInput, optFns ...func(*Options)) (*CreatePlatformEndpointOutput, error) {
	if params == nil {
		params = &CreatePlatformEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlatformEndpoint", params, optFns, c.addOperationCreatePlatformEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlatformEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for CreatePlatformEndpoint action.
type CreatePlatformEndpointInput struct {

	// PlatformApplicationArn returned from CreatePlatformApplication is used to
	// create a an endpoint.
	//
	// This member is required.
	PlatformApplicationArn *string

	// Unique identifier created by the notification service for an app on a device.
	// The specific name for Token will vary, depending on which notification service
	// is being used. For example, when using APNS as the notification service, you
	// need the device token. Alternatively, when using GCM (Firebase Cloud Messaging)
	// or ADM, the device token equivalent is called the registration ID.
	//
	// This member is required.
	Token *string

	// For a list of attributes, see [SetEndpointAttributes]SetEndpointAttributes .
	//
	// [SetEndpointAttributes]: https://docs.aws.amazon.com/sns/latest/api/API_SetEndpointAttributes.html
	Attributes map[string]string

	// Arbitrary user data to associate with the endpoint. Amazon SNS does not use
	// this data. The data must be in UTF-8 format and less than 2KB.
	CustomUserData *string

	noSmithyDocumentSerde
}

// Response from CreateEndpoint action.
type CreatePlatformEndpointOutput struct {

	// EndpointArn returned from CreateEndpoint action.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlatformEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePlatformEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePlatformEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlatformEndpoint"); err != nil {
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
	if err = addOpCreatePlatformEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlatformEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlatformEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlatformEndpoint",
	}
}
