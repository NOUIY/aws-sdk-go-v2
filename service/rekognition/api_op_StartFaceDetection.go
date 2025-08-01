// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts asynchronous detection of faces in a stored video.
//
// Amazon Rekognition Video can detect faces in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartFaceDetection returns a job identifier ( JobId ) that you use to get the
// results of the operation. When face detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel . To get the results of the face
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetFaceDetection and pass the job identifier ( JobId ) from
// the initial call to StartFaceDetection .
//
// For more information, see Detecting faces in a stored video in the Amazon
// Rekognition Developer Guide.
func (c *Client) StartFaceDetection(ctx context.Context, params *StartFaceDetectionInput, optFns ...func(*Options)) (*StartFaceDetectionOutput, error) {
	if params == nil {
		params = &StartFaceDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFaceDetection", params, optFns, c.addOperationStartFaceDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFaceDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFaceDetectionInput struct {

	// The video in which you want to detect faces. The video must be stored in an
	// Amazon S3 bucket.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartFaceDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// The face attributes you want returned.
	//
	// DEFAULT - The following subset of facial attributes are returned: BoundingBox,
	// Confidence, Pose, Quality and Landmarks.
	//
	// ALL - All facial attributes are returned.
	FaceAttributes types.FaceAttributes

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the face detection operation. The Amazon SNS
	// topic must have a topic name that begins with AmazonRekognition if you are using
	// the AmazonRekognitionServiceRole permissions policy.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartFaceDetectionOutput struct {

	// The identifier for the face detection job. Use JobId to identify the job in a
	// subsequent call to GetFaceDetection .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFaceDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartFaceDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFaceDetection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartFaceDetection"); err != nil {
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
	if err = addOpStartFaceDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFaceDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartFaceDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartFaceDetection",
	}
}
