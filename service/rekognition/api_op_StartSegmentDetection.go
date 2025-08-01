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

// Starts asynchronous detection of segment detection in a stored video.
//
// Amazon Rekognition Video can detect segments in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartSegmentDetection returns a job identifier ( JobId ) which you use to get
// the results of the operation. When segment detection is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel .
//
// You can use the Filters (StartSegmentDetectionFilters ) input parameter to specify the minimum detection
// confidence returned in the response. Within Filters , use ShotFilter (StartShotDetectionFilter ) to
// filter detected shots. Use TechnicalCueFilter (StartTechnicalCueDetectionFilter ) to filter technical cues.
//
// To get the results of the segment detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . if so, call GetSegmentDetection and
// pass the job identifier ( JobId ) from the initial call to StartSegmentDetection
// .
//
// For more information, see Detecting video segments in stored video in the
// Amazon Rekognition Developer Guide.
func (c *Client) StartSegmentDetection(ctx context.Context, params *StartSegmentDetectionInput, optFns ...func(*Options)) (*StartSegmentDetectionOutput, error) {
	if params == nil {
		params = &StartSegmentDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSegmentDetection", params, optFns, c.addOperationStartSegmentDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSegmentDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSegmentDetectionInput struct {

	// An array of segment types to detect in the video. Valid values are
	// TECHNICAL_CUE and SHOT.
	//
	// This member is required.
	SegmentTypes []types.SegmentType

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetectionuse Video to specify a video for analysis. The supported
	// file formats are .mp4, .mov and .avi.
	//
	// This member is required.
	Video *types.Video

	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartSegmentDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string

	// Filters for technical cue or shot detection.
	Filters *types.StartSegmentDetectionFilters

	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the segment detection operation. Note that the
	// Amazon SNS topic must have a topic name that begins with AmazonRekognition if
	// you are using the AmazonRekognitionServiceRole permissions policy to access the
	// topic.
	NotificationChannel *types.NotificationChannel

	noSmithyDocumentSerde
}

type StartSegmentDetectionOutput struct {

	// Unique identifier for the segment detection job. The JobId is returned from
	// StartSegmentDetection .
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSegmentDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSegmentDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSegmentDetection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSegmentDetection"); err != nil {
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
	if err = addOpStartSegmentDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSegmentDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSegmentDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSegmentDetection",
	}
}
