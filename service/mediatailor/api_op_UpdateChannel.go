// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func (c *Client) UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) {
	if params == nil {
		params = &UpdateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChannel", params, optFns, c.addOperationUpdateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChannelInput struct {

	// The name of the channel.
	//
	// This member is required.
	ChannelName *string

	// The channel's output properties.
	//
	// This member is required.
	Outputs []types.RequestOutputItem

	// The list of audiences defined in channel.
	Audiences []string

	// The slate used to fill gaps between programs in the schedule. You must
	// configure filler slate if your channel uses the LINEAR PlaybackMode .
	// MediaTailor doesn't support filler slate for channels using the LOOP PlaybackMode
	// .
	FillerSlate *types.SlateSource

	//  The time-shifted viewing configuration you want to associate to the channel.
	TimeShiftConfiguration *types.TimeShiftConfiguration

	noSmithyDocumentSerde
}

type UpdateChannelOutput struct {

	// The Amazon Resource Name (ARN) associated with the channel.
	Arn *string

	// The list of audiences defined in channel.
	Audiences []string

	// The name of the channel.
	ChannelName *string

	// Returns the state whether the channel is running or not.
	ChannelState types.ChannelState

	// The timestamp of when the channel was created.
	CreationTime *time.Time

	// The slate used to fill gaps between programs in the schedule. You must
	// configure filler slate if your channel uses the LINEAR PlaybackMode .
	// MediaTailor doesn't support filler slate for channels using the LOOP PlaybackMode
	// .
	FillerSlate *types.SlateSource

	// The timestamp that indicates when the channel was last modified.
	LastModifiedTime *time.Time

	// The channel's output properties.
	Outputs []types.ResponseOutputItem

	// The type of playback mode for this channel.
	//
	// LINEAR - Programs play back-to-back only once.
	//
	// LOOP - Programs play back-to-back in an endless loop. When the last program in
	// the schedule plays, playback loops back to the first program in the schedule.
	PlaybackMode *string

	// The tags to assign to the channel. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources].
	//
	// [Tagging AWS Elemental MediaTailor Resources]: https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html
	Tags map[string]string

	// The tier associated with this Channel.
	Tier *string

	//  The time-shifted viewing configuration for the channel.
	TimeShiftConfiguration *types.TimeShiftConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateChannel"); err != nil {
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
	if err = addOpUpdateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateChannel",
	}
}
