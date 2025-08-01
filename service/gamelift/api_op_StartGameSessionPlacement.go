// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Makes a request to start a new game session using a game session queue. When
// processing a placement request, Amazon GameLift Servers looks for the best
// possible available resource to host the game session, based on how the queue is
// configured to prioritize factors such as resource cost, latency, and location.
// After selecting an available resource, Amazon GameLift Servers prompts the
// resource to start a game session. A placement request can include a list of
// players to create a set of player sessions. The request can also include
// information to pass to the new game session, such as to specify a game map or
// other options.
//
// # Request options
//
// Use this operation to make the following types of requests.
//
//   - Request a placement using the queue's default prioritization process (see
//     the default prioritization described in [PriorityConfiguration]). Include these required parameters:
//
//   - GameSessionQueueName
//
//   - MaximumPlayerSessionCount
//
//   - PlacementID
//
//   - Request a placement and prioritize based on latency. Include these
//     parameters:
//
//   - Required parameters GameSessionQueueName , MaximumPlayerSessionCount ,
//     PlacementID .
//
//   - PlayerLatencies . Include a set of latency values for destinations in the
//     queue. When a request includes latency data, Amazon GameLift Servers
//     automatically reorder the queue's locations priority list based on lowest
//     available latency values. If a request includes latency data for multiple
//     players, Amazon GameLift Servers calculates each location's average latency for
//     all players and reorders to find the lowest latency across all players.
//
//   - Don't include PriorityConfigurationOverride .
//
//   - Prioritize based on a custom list of locations. If you're using a queue
//     that's configured to prioritize location first (see [PriorityConfiguration]for game session queues),
//     you can optionally use the PriorityConfigurationOverride parameter to substitute
//     a different location priority list for this placement request. Amazon GameLift
//     Servers searches each location on the priority override list to find an
//     available hosting resource for the new game session. Specify a fallback strategy
//     to use in the event that Amazon GameLift Servers fails to place the game session
//     in any of the locations on the override list.
//
//   - Request a placement and prioritized based on a custom list of locations.
//
//   - You can request new player sessions for a group of players. Include the
//     DesiredPlayerSessions parameter and include at minimum a unique player ID for
//     each. You can also include player-specific data to pass to the new game session.
//
// # Result
//
// If successful, this operation generates a new game session placement request
// and adds it to the game session queue for processing. You can track the status
// of individual placement requests by calling [DescribeGameSessionPlacement]or by monitoring queue
// notifications. When the request status is FULFILLED , a new game session has
// started and the placement request is updated with connection information for the
// game session (IP address and port). If the request included player session data,
// Amazon GameLift Servers creates a player session for each player ID in the
// request.
//
// The request results in a InvalidRequestException in the following situations:
//
//   - If the request includes both PlayerLatencies and
//     PriorityConfigurationOverride parameters.
//
//   - If the request includes the PriorityConfigurationOverride parameter and
//     specifies a queue that doesn't prioritize locations.
//
// Amazon GameLift Servers continues to retry each placement request until it
// reaches the queue's timeout setting. If a request times out, you can resubmit
// the request to the same queue or try a different queue.
//
// [DescribeGameSessionPlacement]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeGameSessionPlacement.html
// [PriorityConfiguration]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_PriorityConfiguration.html
func (c *Client) StartGameSessionPlacement(ctx context.Context, params *StartGameSessionPlacementInput, optFns ...func(*Options)) (*StartGameSessionPlacementOutput, error) {
	if params == nil {
		params = &StartGameSessionPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartGameSessionPlacement", params, optFns, c.addOperationStartGameSessionPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartGameSessionPlacementInput struct {

	// Name of the queue to use to place the new game session. You can use either the
	// queue name or ARN value.
	//
	// This member is required.
	GameSessionQueueName *string

	// The maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// This member is required.
	MaximumPlayerSessionCount *int32

	// A unique identifier to assign to the new game session placement. This value is
	// developer-defined. The value must be unique across all Regions and cannot be
	// reused.
	//
	// This member is required.
	PlacementId *string

	// Set of information on each player to create a player session for.
	DesiredPlayerSessions []types.DesiredPlayerSession

	// A set of key-value pairs that can store custom data in a game session. For
	// example: {"Key": "difficulty", "Value": "novice"} .
	GameProperties []types.GameProperty

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process with a request to start a new game
	// session. For more information, see [Start a game session].
	//
	// [Start a game session]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession
	GameSessionData *string

	// A descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	GameSessionName *string

	// A set of values, expressed in milliseconds, that indicates the amount of
	// latency that a player experiences when connected to Amazon Web Services Regions.
	// This information is used to try to place the new game session where it can offer
	// the best possible gameplay experience for the players.
	PlayerLatencies []types.PlayerLatency

	// A prioritized list of locations to use for the game session placement and
	// instructions on how to use it. This list overrides a queue's prioritized
	// location list for this game session placement request only. You can include
	// Amazon Web Services Regions, local zones, and custom locations (for Anywhere
	// fleets). You can choose to limit placements to locations on the override list
	// only, or you can prioritize locations on the override list first and then fall
	// back to the queue's other locations if needed. Choose a fallback strategy to use
	// in the event that Amazon GameLift Servers fails to place a game session in any
	// of the locations on the priority override list.
	PriorityConfigurationOverride *types.PriorityConfigurationOverride

	noSmithyDocumentSerde
}

type StartGameSessionPlacementOutput struct {

	// Object that describes the newly created game session placement. This object
	// includes all the information provided in the request, as well as start/end time
	// stamps and placement status.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartGameSessionPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartGameSessionPlacement"); err != nil {
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
	if err = addOpStartGameSessionPlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartGameSessionPlacement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartGameSessionPlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartGameSessionPlacement",
	}
}
