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

// Reserves open slots in a game session for a group of players. New player
// sessions can be created in any game session with an open slot that is in ACTIVE
// status and has a player creation policy of ACCEPT_ALL . To add a single player
// to a game session, use [CreatePlayerSession]
//
// To create player sessions, specify a game session ID and a list of player IDs.
// Optionally, provide a set of player data for each player ID.
//
// If successful, a slot is reserved in the game session for each player, and new
// PlayerSession objects are returned with player session IDs. Each player
// references their player session ID when sending a connection request to the game
// session, and the game server can use it to validate the player reservation with
// the Amazon GameLift Servers service. Player sessions cannot be updated.
//
// The maximum number of players per game session is 200. It is not adjustable.
//
// # Related actions
//
// [All APIs by task]
//
// [CreatePlayerSession]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreatePlayerSession.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
func (c *Client) CreatePlayerSessions(ctx context.Context, params *CreatePlayerSessionsInput, optFns ...func(*Options)) (*CreatePlayerSessionsOutput, error) {
	if params == nil {
		params = &CreatePlayerSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlayerSessions", params, optFns, c.addOperationCreatePlayerSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlayerSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlayerSessionsInput struct {

	// A unique identifier for the game session to add players to.
	//
	// This member is required.
	GameSessionId *string

	// List of unique identifiers for the players to be added.
	//
	// This member is required.
	PlayerIds []string

	// Map of string pairs, each specifying a player ID and a set of developer-defined
	// information related to the player. Amazon GameLift Servers does not use this
	// data, so it can be formatted as needed for use in the game. Any player data
	// strings for player IDs that are not included in the PlayerIds parameter are
	// ignored.
	PlayerDataMap map[string]string

	noSmithyDocumentSerde
}

type CreatePlayerSessionsOutput struct {

	// A collection of player session objects created for the added players.
	PlayerSessions []types.PlayerSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlayerSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePlayerSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePlayerSessions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlayerSessions"); err != nil {
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
	if err = addOpCreatePlayerSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlayerSessions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlayerSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlayerSessions",
	}
}
