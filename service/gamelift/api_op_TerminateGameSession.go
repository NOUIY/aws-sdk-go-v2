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

// Ends a game session that's currently in progress. Use this action to terminate
// any game session that isn't in ERROR status. Terminating a game session is the
// most efficient way to free up a server process when it's hosting a game session
// that's in a bad state or not ending properly. You can use this action to
// terminate a game session that's being hosted on any type of Amazon GameLift
// Servers fleet compute, including computes for managed EC2, managed container,
// and Anywhere fleets. The game server must be integrated with Amazon GameLift
// Servers server SDK 5.x or greater.
//
// # Request options
//
// Request termination for a single game session. Provide the game session ID and
// the termination mode. There are two potential methods for terminating a game
// session:
//
//   - Initiate a graceful termination using the normal game session shutdown
//     sequence. With this mode, the Amazon GameLift Servers service prompts the server
//     process that's hosting the game session by calling the server SDK callback
//     method OnProcessTerminate() . The callback implementation is part of the
//     custom game server code. It might involve a variety of actions to gracefully end
//     a game session, such as notifying players, before stopping the server process.
//
//   - Force an immediate game session termination. With this mode, the Amazon
//     GameLift Servers service takes action to stop the server process, which ends the
//     game session without the normal game session shutdown sequence.
//
// # Results
//
// If successful, game session termination is initiated. During this activity, the
// game session status is changed to TERMINATING . When completed, the server
// process that was hosting the game session has been stopped and replaced with a
// new server process that's ready to host a new game session. The old game
// session's status is changed to TERMINATED with a status reason that indicates
// the termination method used.
//
// # Learn more
//
// [Add Amazon GameLift Servers to your game server]
//
// Amazon GameLift Servers server SDK 5 reference guide for OnProcessTerminate() ([C++]
// ) ([C#] ) ([Unreal] ) ([Go] )
//
// [C#]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-csharp-initsdk.html
// [C++]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-cpp-initsdk.html
// [Add Amazon GameLift Servers to your game server]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html
// [Unreal]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-unreal-initsdk.html
// [Go]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk-go-initsdk.html
func (c *Client) TerminateGameSession(ctx context.Context, params *TerminateGameSessionInput, optFns ...func(*Options)) (*TerminateGameSessionOutput, error) {
	if params == nil {
		params = &TerminateGameSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateGameSession", params, optFns, c.addOperationTerminateGameSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateGameSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateGameSessionInput struct {

	// A unique identifier for the game session to be terminated. A game session ARN
	// has the following format: arn:aws:gamelift:::gamesession// .
	//
	// This member is required.
	GameSessionId *string

	// The method to use to terminate the game session. Available methods include:
	//
	//   - TRIGGER_ON_PROCESS_TERMINATE – Prompts the Amazon GameLift Servers service
	//   to send an OnProcessTerminate() callback to the server process and initiate
	//   the normal game session shutdown sequence. The OnProcessTerminate method,
	//   which is implemented in the game server code, must include a call to the server
	//   SDK action ProcessEnding() , which is how the server process signals to Amazon
	//   GameLift Servers that a game session is ending. If the server process doesn't
	//   call ProcessEnding() , the game session termination won't conclude
	//   successfully.
	//
	//   - FORCE_TERMINATE – Prompts the Amazon GameLift Servers service to stop the
	//   server process immediately. Amazon GameLift Servers takes action (depending on
	//   the type of fleet) to shut down the server process without the normal game
	//   session shutdown sequence.
	//
	// This method is not available for game sessions that are running on Anywhere
	//   fleets unless the fleet is deployed with the Amazon GameLift Servers Agent. In
	//   this scenario, a force terminate request results in an invalid or bad request
	//   exception.
	//
	// This member is required.
	TerminationMode types.TerminationMode

	noSmithyDocumentSerde
}

type TerminateGameSessionOutput struct {

	// Properties describing a game session.
	//
	// A game session in ACTIVE status can host players. When a game session ends, its
	// status is set to TERMINATED .
	//
	// Amazon GameLift Servers retains a game session resource for 30 days after the
	// game session ends. You can reuse idempotency token values after this time. Game
	// session logs are retained for 14 days.
	//
	// [All APIs by task]
	//
	// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
	GameSession *types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTerminateGameSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateGameSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateGameSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TerminateGameSession"); err != nil {
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
	if err = addOpTerminateGameSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateGameSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTerminateGameSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TerminateGameSession",
	}
}
