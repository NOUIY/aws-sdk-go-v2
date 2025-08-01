// Code generated by smithy-go-codegen DO NOT EDIT.

package socialmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/socialmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Upload a media file to the WhatsApp service. Only the specified
// originationPhoneNumberId has the permissions to send the media file when using [SendWhatsAppMessage]
// . You must use either sourceS3File or sourceS3PresignedUrl for the source. If
// both or neither are specified then an InvalidParameterException is returned.
//
// [SendWhatsAppMessage]: https://docs.aws.amazon.com/social-messaging/latest/APIReference/API_SendWhatsAppMessage.html
func (c *Client) PostWhatsAppMessageMedia(ctx context.Context, params *PostWhatsAppMessageMediaInput, optFns ...func(*Options)) (*PostWhatsAppMessageMediaOutput, error) {
	if params == nil {
		params = &PostWhatsAppMessageMediaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PostWhatsAppMessageMedia", params, optFns, c.addOperationPostWhatsAppMessageMediaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PostWhatsAppMessageMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostWhatsAppMessageMediaInput struct {

	// The ID of the phone number to associate with the WhatsApp media file. The phone
	// number identifiers are formatted as
	// phone-number-id-01234567890123456789012345678901 . Use [GetLinkedWhatsAppBusinessAccount] to find a phone
	// number's id.
	//
	// [GetLinkedWhatsAppBusinessAccount]: https://docs.aws.amazon.com/social-messaging/latest/APIReference/API_GetLinkedWhatsAppBusinessAccountPhoneNumber.html
	//
	// This member is required.
	OriginationPhoneNumberId *string

	// The source S3 url for the media file.
	SourceS3File *types.S3File

	// The source presign url of the media file.
	SourceS3PresignedUrl *types.S3PresignedUrl

	noSmithyDocumentSerde
}

type PostWhatsAppMessageMediaOutput struct {

	// The unique identifier of the posted WhatsApp message.
	MediaId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPostWhatsAppMessageMediaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPostWhatsAppMessageMedia{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPostWhatsAppMessageMedia{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PostWhatsAppMessageMedia"); err != nil {
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
	if err = addOpPostWhatsAppMessageMediaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPostWhatsAppMessageMedia(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPostWhatsAppMessageMedia(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PostWhatsAppMessageMedia",
	}
}
