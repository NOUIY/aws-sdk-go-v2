// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts the replication task.
//
// For more information about DMS tasks, see [Working with Migration Tasks] in the Database Migration Service
// User Guide.
//
// [Working with Migration Tasks]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html
func (c *Client) StartReplicationTask(ctx context.Context, params *StartReplicationTaskInput, optFns ...func(*Options)) (*StartReplicationTaskOutput, error) {
	if params == nil {
		params = &StartReplicationTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReplicationTask", params, optFns, c.addOperationStartReplicationTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReplicationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReplicationTaskInput struct {

	// The Amazon Resource Name (ARN) of the replication task to be started.
	//
	// This member is required.
	ReplicationTaskArn *string

	// The type of replication task to start.
	//
	// start-replication is the only valid action that can be used for the first time
	// a task with the migration type of full-load full-load, full-load-and-cdc or cdc
	// is run. Any other action used for the first time on a given task, such as
	// resume-processing and reload-target will result in data errors.
	//
	// You can also use ReloadTables to reload specific tables that failed during migration
	// instead of restarting the task.
	//
	// For a full-load task, the resume-processing option will reload any tables that
	// were partially loaded or not yet loaded during the full load phase.
	//
	// For a full-load-and-cdc task, DMS migrates table data, and then applies data
	// changes that occur on the source. To load all the tables again, and start
	// capturing source changes, use reload-target . Otherwise use resume-processing ,
	// to replicate the changes from the last stop position.
	//
	// For a cdc only task, to start from a specific position, you must use
	// start-replication and also specify the start position. Check the source endpoint
	// DMS documentation for any limitations. For example, not all sources support
	// starting from a time.
	//
	// resume-processing is only available for previously executed tasks.
	//
	// This member is required.
	StartReplicationTaskType types.StartReplicationTaskTypeValue

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, or LSN/SCN format.
	//
	// Date Example: --cdc-start-position “2018-03-08T12:12:12”
	//
	// Checkpoint Example: --cdc-start-position
	// "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	//
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373”
	//
	// When you use this task setting with a source PostgreSQL database, a logical
	// replication slot should already be created and associated with the source
	// endpoint. You can verify this by setting the slotName extra connection
	// attribute to the name of this logical replication slot. For more information,
	// see [Extra Connection Attributes When Using PostgreSQL as a Source for DMS].
	//
	// [Extra Connection Attributes When Using PostgreSQL as a Source for DMS]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib
	CdcStartPosition *string

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation to
	// start. Specifying both values results in an error.
	//
	// Timestamp Example: --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time

	// Indicates when you want a change data capture (CDC) operation to stop. The
	// value can be either server time or commit time.
	//
	// Server time example: --cdc-stop-position “server_time:2018-02-09T12:12:12”
	//
	// Commit time example: --cdc-stop-position “commit_time:2018-02-09T12:12:12“
	CdcStopPosition *string

	noSmithyDocumentSerde
}

type StartReplicationTaskOutput struct {

	// The replication task started.
	ReplicationTask *types.ReplicationTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReplicationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartReplicationTask"); err != nil {
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
	if err = addOpStartReplicationTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartReplicationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartReplicationTask",
	}
}
