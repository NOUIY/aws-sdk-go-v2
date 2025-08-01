// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the endpoint attributes for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
// information, see [Using Amazon SNS Mobile Push Notifications].
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func (c *Client) GetEndpointAttributes(ctx context.Context, params *GetEndpointAttributesInput, optFns ...func(*Options)) (*GetEndpointAttributesOutput, error) {
	if params == nil {
		params = &GetEndpointAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEndpointAttributes", params, optFns, c.addOperationGetEndpointAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEndpointAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetEndpointAttributes action.
type GetEndpointAttributesInput struct {

	// EndpointArn for GetEndpointAttributes input.
	//
	// This member is required.
	EndpointArn *string

	noSmithyDocumentSerde
}

// Response from GetEndpointAttributes of the EndpointArn .
type GetEndpointAttributesOutput struct {

	// Attributes include the following:
	//
	//   - CustomUserData – arbitrary user data to associate with the endpoint. Amazon
	//   SNS does not use this data. The data must be in UTF-8 format and less than 2KB.
	//
	//   - Enabled – flag that enables/disables delivery to the endpoint. Amazon SNS
	//   will set this to false when a notification service indicates to Amazon SNS that
	//   the endpoint is invalid. Users can set it back to true, typically after updating
	//   Token.
	//
	//   - Token – device token, also referred to as a registration id, for an app and
	//   mobile device. This is returned from the notification service when an app and
	//   mobile device are registered with the notification service.
	//
	// The device token for the iOS platform is returned in lowercase.
	Attributes map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEndpointAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetEndpointAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetEndpointAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEndpointAttributes"); err != nil {
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
	if err = addOpGetEndpointAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEndpointAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEndpointAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEndpointAttributes",
	}
}
