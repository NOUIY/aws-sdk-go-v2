// Code generated by smithy-go-codegen DO NOT EDIT.

package gameliftstreams

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gameliftstreams/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	This action initiates a new stream session and outputs connection information
//
// that clients can use to access the stream. A stream session refers to an
// instance of a stream that Amazon GameLift Streams transmits from the server to
// the end-user. A stream session runs on a compute resource that a stream group
// has allocated.
//
// To start a new stream session, specify a stream group and application ID, along
// with the transport protocol and signal request settings to use with the stream.
// You must have associated at least one application to the stream group before
// starting a stream session, either when creating the stream group, or by using [AssociateApplications].
//
// For stream groups that have multiple locations, provide a set of locations
// ordered by priority using a Locations parameter. Amazon GameLift Streams will
// start a single stream session in the next available location. An application
// must be finished replicating in a remote location before the remote location can
// host a stream.
//
// If the request is successful, Amazon GameLift Streams begins to prepare the
// stream. Amazon GameLift Streams assigns an Amazon Resource Name (ARN) value to
// the stream session resource and sets the status to ACTIVATING . During the
// stream preparation process, Amazon GameLift Streams queues the request and
// searches for available stream capacity to run the stream. This results in one of
// the following:
//
//   - Amazon GameLift Streams identifies an available compute resource to run the
//     application content and start the stream. When the stream is ready, the stream
//     session's status changes to ACTIVE and includes stream connection information.
//     Provide the connection information to the requesting client to join the stream
//     session.
//
//   - Amazon GameLift Streams doesn't identify an available resource within a
//     certain time, set by ClientToken . In this case, Amazon GameLift Streams stops
//     processing the request, and the stream session object status changes to ERROR
//     with status reason placementTimeout .
//
// [AssociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_AssociateApplications.html
func (c *Client) StartStreamSession(ctx context.Context, params *StartStreamSessionInput, optFns ...func(*Options)) (*StartStreamSessionOutput, error) {
	if params == nil {
		params = &StartStreamSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStreamSession", params, optFns, c.addOperationStartStreamSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStreamSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamSessionInput struct {

	// An [Amazon Resource Name (ARN)] or ID that uniquely identifies the application resource. Example ARN:
	// arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6 . Example
	// ID: a-9ZY8X7Wv6 .
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html
	//
	// This member is required.
	ApplicationIdentifier *string

	// The stream group to run this stream session with.
	//
	// This value is an [Amazon Resource Name (ARN)] or ID that uniquely identifies the stream group resource.
	// Example ARN:
	// arn:aws:gameliftstreams:us-west-2:111122223333:streamgroup/sg-1AB2C3De4 .
	// Example ID: sg-1AB2C3De4 .
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html
	//
	// This member is required.
	Identifier *string

	// The data transport protocol to use for the stream session.
	//
	// This member is required.
	Protocol types.Protocol

	// A WebRTC ICE offer string to use when initializing a WebRTC connection.
	// Typically, the offer is a very long JSON string. Provide the string as a text
	// value in quotes.
	//
	// Amazon GameLift Streams also supports setting the field to
	// "NO_CLIENT_CONNECTION". This will create a session without needing any browser
	// request or Web SDK integration. The session starts up as usual and waits for a
	// reconnection from a browser, which is accomplished using [CreateStreamSessionConnection].
	//
	// [CreateStreamSessionConnection]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_CreateStreamSessionConnection.html
	//
	// This member is required.
	SignalRequest *string

	// A set of options that you can use to control the stream session runtime
	// environment, expressed as a set of key-value pairs. You can use this to
	// configure the application or stream session details. You can also provide custom
	// environment variables that Amazon GameLift Streams passes to your game client.
	//
	// If you want to debug your application with environment variables, we recommend
	// that you do so in a local environment outside of Amazon GameLift Streams. For
	// more information, refer to the Compatibility Guidance in the troubleshooting
	// section of the Developer Guide.
	//
	// AdditionalEnvironmentVariables and AdditionalLaunchArgs have similar purposes.
	// AdditionalEnvironmentVariables passes data using environment variables; while
	// AdditionalLaunchArgs passes data using command-line arguments.
	AdditionalEnvironmentVariables map[string]string

	// A list of CLI arguments that are sent to the streaming server when a stream
	// session launches. You can use this to configure the application or stream
	// session details. You can also provide custom arguments that Amazon GameLift
	// Streams passes to your game client.
	//
	// AdditionalEnvironmentVariables and AdditionalLaunchArgs have similar purposes.
	// AdditionalEnvironmentVariables passes data using environment variables; while
	// AdditionalLaunchArgs passes data using command-line arguments.
	AdditionalLaunchArgs []string

	//  A unique identifier that represents a client request. The request is
	// idempotent, which ensures that an API request completes only once. When users
	// send a request, Amazon GameLift Streams automatically populates this field.
	ClientToken *string

	// Length of time (in seconds) that Amazon GameLift Streams should wait for a
	// client to connect to the stream session. This time span starts when the stream
	// session reaches ACTIVE status. If no client connects before the timeout, Amazon
	// GameLift Streams stops the stream session with status of TERMINATED . Default
	// value is 120.
	ConnectionTimeoutSeconds *int32

	// A human-readable label for the stream session. You can update this value later.
	Description *string

	//  A list of locations, in order of priority, where you want Amazon GameLift
	// Streams to start a stream from. Amazon GameLift Streams selects the location
	// with the next available capacity to start a single stream session in. If this
	// value is empty, Amazon GameLift Streams attempts to start a stream session in
	// the primary location.
	//
	// This value is A set of location names. For example, us-east-1 . For a complete
	// list of locations that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations]in the Amazon
	// GameLift Streams Developer Guide.
	//
	// [Regions, quotas, and limitations]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html
	Locations []string

	// The maximum length of time (in seconds) that Amazon GameLift Streams keeps the
	// stream session open. At this point, Amazon GameLift Streams ends the stream
	// session regardless of any existing client connections. Default value is 43200.
	SessionLengthSeconds *int32

	//  An opaque, unique identifier for an end-user, defined by the developer.
	UserId *string

	noSmithyDocumentSerde
}

type StartStreamSessionOutput struct {

	// A set of options that you can use to control the stream session runtime
	// environment, expressed as a set of key-value pairs. You can use this to
	// configure the application or stream session details. You can also provide custom
	// environment variables that Amazon GameLift Streams passes to your game client.
	//
	// If you want to debug your application with environment variables, we recommend
	// that you do so in a local environment outside of Amazon GameLift Streams. For
	// more information, refer to the Compatibility Guidance in the troubleshooting
	// section of the Developer Guide.
	//
	// AdditionalEnvironmentVariables and AdditionalLaunchArgs have similar purposes.
	// AdditionalEnvironmentVariables passes data using environment variables; while
	// AdditionalLaunchArgs passes data using command-line arguments.
	AdditionalEnvironmentVariables map[string]string

	// A list of CLI arguments that are sent to the streaming server when a stream
	// session launches. You can use this to configure the application or stream
	// session details. You can also provide custom arguments that Amazon GameLift
	// Streams passes to your game client.
	//
	// AdditionalEnvironmentVariables and AdditionalLaunchArgs have similar purposes.
	// AdditionalEnvironmentVariables passes data using environment variables; while
	// AdditionalLaunchArgs passes data using command-line arguments.
	AdditionalLaunchArgs []string

	// An [Amazon Resource Name (ARN)] that uniquely identifies the application resource. Example ARN:
	// arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6 .
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html
	ApplicationArn *string

	// The [Amazon Resource Name (ARN)] that's assigned to a stream session resource. When combined with the
	// stream group resource ID, this value uniquely identifies the stream session
	// across all Amazon Web Services Regions. Format is arn:aws:gameliftstreams:[AWS
	// Region]:[AWS account]:streamsession/[stream group resource ID]/[stream session
	// resource ID] .
	//
	// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html
	Arn *string

	// The maximum length of time (in seconds) that Amazon GameLift Streams keeps the
	// stream session open. At this point, Amazon GameLift Streams ends the stream
	// session regardless of any existing client connections.
	ConnectionTimeoutSeconds *int32

	// A timestamp that indicates when this resource was created. Timestamps are
	// expressed using in ISO8601 format, such as: 2022-12-27T22:29:40+00:00 (UTC).
	CreatedAt *time.Time

	// A human-readable label for the stream session. You can update this value at any
	// time.
	Description *string

	// Provides details about the stream session's exported files.
	ExportFilesMetadata *types.ExportFilesMetadata

	// A timestamp that indicates when this resource was last updated. Timestamps are
	// expressed using in ISO8601 format, such as: 2022-12-27T22:29:40+00:00 (UTC).
	LastUpdatedAt *time.Time

	//  The location where Amazon GameLift Streams is streaming your application from.
	//
	// A location's name. For example, us-east-1 . For a complete list of locations
	// that Amazon GameLift Streams supports, refer to [Regions, quotas, and limitations]in the Amazon GameLift Streams
	// Developer Guide.
	//
	// [Regions, quotas, and limitations]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/regions-quotas.html
	Location *string

	// Access location for log files that your content generates during a stream
	// session. These log files are uploaded to cloud storage location at the end of a
	// stream session. The Amazon GameLift Streams application resource defines which
	// log files to upload.
	LogFileLocationUri *string

	// The data transfer protocol in use with the stream session.
	Protocol types.Protocol

	// The length of time that Amazon GameLift Streams keeps the game session open.
	SessionLengthSeconds *int32

	// The WebRTC ICE offer string that a client generates to initiate a connection to
	// the stream session.
	SignalRequest *string

	// The WebRTC answer string that the stream server generates in response to the
	// SignalRequest .
	SignalResponse *string

	// The current status of the stream session. A stream session can host clients
	// when in ACTIVE status.
	Status types.StreamSessionStatus

	// A short description of the reason the stream session is in ERROR status.
	StatusReason types.StreamSessionStatusReason

	// The unique identifier for the Amazon GameLift Streams stream group that is
	// hosting the stream session. Format example: sg-1AB2C3De4 .
	StreamGroupId *string

	//  An opaque, unique identifier for an end-user, defined by the developer.
	UserId *string

	// The URL of an S3 bucket that stores Amazon GameLift Streams WebSDK files. The
	// URL is used to establish connection with the client.
	WebSdkProtocolUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartStreamSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartStreamSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartStreamSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartStreamSession"); err != nil {
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
	if err = addIdempotencyToken_opStartStreamSessionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartStreamSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamSession(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartStreamSession struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartStreamSession) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartStreamSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartStreamSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartStreamSessionInput ")
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
func addIdempotencyToken_opStartStreamSessionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartStreamSession{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartStreamSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartStreamSession",
	}
}
