// Code generated by smithy-go-codegen DO NOT EDIT.

package qldbsession

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qldbsession/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a command to an Amazon QLDB ledger.
//
// Instead of interacting directly with this API, we recommend using the QLDB
// driver or the QLDB shell to execute data transactions on a ledger.
//
//   - If you are working with an AWS SDK, use the QLDB driver. The driver
//     provides a high-level abstraction layer above this QLDB Session data plane and
//     manages SendCommand API calls for you. For information and a list of supported
//     programming languages, see [Getting started with the driver]in the Amazon QLDB Developer Guide.
//
//   - If you are working with the AWS Command Line Interface (AWS CLI), use the
//     QLDB shell. The shell is a command line interface that uses the QLDB driver to
//     interact with a ledger. For information, see [Accessing Amazon QLDB using the QLDB shell].
//
// [Getting started with the driver]: https://docs.aws.amazon.com/qldb/latest/developerguide/getting-started-driver.html
// [Accessing Amazon QLDB using the QLDB shell]: https://docs.aws.amazon.com/qldb/latest/developerguide/data-shell.html
func (c *Client) SendCommand(ctx context.Context, params *SendCommandInput, optFns ...func(*Options)) (*SendCommandOutput, error) {
	if params == nil {
		params = &SendCommandInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendCommand", params, optFns, c.addOperationSendCommandMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendCommandOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendCommandInput struct {

	// Command to abort the current transaction.
	AbortTransaction *types.AbortTransactionRequest

	// Command to commit the specified transaction.
	CommitTransaction *types.CommitTransactionRequest

	// Command to end the current session.
	EndSession *types.EndSessionRequest

	// Command to execute a statement in the specified transaction.
	ExecuteStatement *types.ExecuteStatementRequest

	// Command to fetch a page.
	FetchPage *types.FetchPageRequest

	// Specifies the session token for the current command. A session token is
	// constant throughout the life of the session.
	//
	// To obtain a session token, run the StartSession command. This SessionToken is
	// required for every subsequent command that is issued during the current session.
	SessionToken *string

	// Command to start a new session. A session token is obtained as part of the
	// response.
	StartSession *types.StartSessionRequest

	// Command to start a new transaction.
	StartTransaction *types.StartTransactionRequest

	noSmithyDocumentSerde
}

type SendCommandOutput struct {

	// Contains the details of the aborted transaction.
	AbortTransaction *types.AbortTransactionResult

	// Contains the details of the committed transaction.
	CommitTransaction *types.CommitTransactionResult

	// Contains the details of the ended session.
	EndSession *types.EndSessionResult

	// Contains the details of the executed statement.
	ExecuteStatement *types.ExecuteStatementResult

	// Contains the details of the fetched page.
	FetchPage *types.FetchPageResult

	// Contains the details of the started session that includes a session token. This
	// SessionToken is required for every subsequent command that is issued during the
	// current session.
	StartSession *types.StartSessionResult

	// Contains the details of the started transaction.
	StartTransaction *types.StartTransactionResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendCommandMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendCommand{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendCommand{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendCommand"); err != nil {
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
	if err = addOpSendCommandValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendCommand(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendCommand(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendCommand",
	}
}
