// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about an ongoing or complete migration from an Amazon Lex V1
// bot to an Amazon Lex V2 bot. Use this operation to view the migration alerts and
// warnings related to the migration.
func (c *Client) GetMigration(ctx context.Context, params *GetMigrationInput, optFns ...func(*Options)) (*GetMigrationOutput, error) {
	if params == nil {
		params = &GetMigrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMigration", params, optFns, c.addOperationGetMigrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMigrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMigrationInput struct {

	// The unique identifier of the migration to view. The migrationID is returned by
	// the operation.
	//
	// This member is required.
	MigrationId *string

	noSmithyDocumentSerde
}

type GetMigrationOutput struct {

	// A list of alerts and warnings that indicate issues with the migration for the
	// Amazon Lex V1 bot to Amazon Lex V2. You receive a warning when an Amazon Lex V1
	// feature has a different implementation if Amazon Lex V2.
	//
	// For more information, see [Migrating a bot] in the Amazon Lex V2 developer guide.
	//
	// [Migrating a bot]: https://docs.aws.amazon.com/lexv2/latest/dg/migrate.html
	Alerts []types.MigrationAlert

	// The unique identifier of the migration. This is the same as the identifier used
	// when calling the GetMigration operation.
	MigrationId *string

	// Indicates the status of the migration. When the status is COMPLETE the
	// migration is finished and the bot is available in Amazon Lex V2. There may be
	// alerts and warnings that need to be resolved to complete the migration.
	MigrationStatus types.MigrationStatus

	// The strategy used to conduct the migration.
	//
	//   - CREATE_NEW - Creates a new Amazon Lex V2 bot and migrates the Amazon Lex V1
	//   bot to the new bot.
	//
	//   - UPDATE_EXISTING - Overwrites the existing Amazon Lex V2 bot metadata and the
	//   locale being migrated. It doesn't change any other locales in the Amazon Lex V2
	//   bot. If the locale doesn't exist, a new locale is created in the Amazon Lex V2
	//   bot.
	MigrationStrategy types.MigrationStrategy

	// The date and time that the migration started.
	MigrationTimestamp *time.Time

	// The locale of the Amazon Lex V1 bot migrated to Amazon Lex V2.
	V1BotLocale types.Locale

	// The name of the Amazon Lex V1 bot migrated to Amazon Lex V2.
	V1BotName *string

	// The version of the Amazon Lex V1 bot migrated to Amazon Lex V2.
	V1BotVersion *string

	// The unique identifier of the Amazon Lex V2 bot that the Amazon Lex V1 is being
	// migrated to.
	V2BotId *string

	// The IAM role that Amazon Lex uses to run the Amazon Lex V2 bot.
	V2BotRole *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMigrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMigration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMigration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMigration"); err != nil {
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
	if err = addOpGetMigrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMigration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMigration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMigration",
	}
}
