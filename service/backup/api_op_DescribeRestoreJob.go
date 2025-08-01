// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata associated with a restore job that is specified by a job ID.
func (c *Client) DescribeRestoreJob(ctx context.Context, params *DescribeRestoreJobInput, optFns ...func(*Options)) (*DescribeRestoreJobOutput, error) {
	if params == nil {
		params = &DescribeRestoreJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRestoreJob", params, optFns, c.addOperationDescribeRestoreJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRestoreJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRestoreJobInput struct {

	// Uniquely identifies the job that restores a recovery point.
	//
	// This member is required.
	RestoreJobId *string

	noSmithyDocumentSerde
}

type DescribeRestoreJobOutput struct {

	// Returns the account ID that owns the restore job.
	AccountId *string

	// The size, in bytes, of the restored resource.
	BackupSizeInBytes *int64

	// The date and time that a job to restore a recovery point is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CompletionDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time

	// Contains identifying information about the creation of a restore job.
	CreatedBy *types.RestoreJobCreator

	// The Amazon Resource Name (ARN) of the resource that was created by the restore
	// job.
	//
	// The format of the ARN depends on the resource type of the backed-up resource.
	CreatedResourceArn *string

	// The date and time that a restore job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// The status of the data generated by the restore test.
	DeletionStatus types.RestoreDeletionStatus

	// This describes the restore job deletion status.
	DeletionStatusMessage *string

	// The amount of time in minutes that a job restoring a recovery point is expected
	// to take.
	ExpectedCompletionTimeMinutes *int64

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access .
	IamRoleArn *string

	// Contains an estimated percentage that is complete of a job at the time the job
	// status was queried.
	PercentDone *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string

	// The creation date of the recovery point made by the specifed restore job.
	RecoveryPointCreationDate *time.Time

	// Returns metadata associated with a restore job listed by resource type.
	ResourceType *string

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string

	// Status code specifying the state of the job that is initiated by Backup to
	// restore a recovery point.
	Status types.RestoreJobStatus

	// A message showing the status of a job to restore a recovery point.
	StatusMessage *string

	// The status of validation run on the indicated restore job.
	ValidationStatus types.RestoreValidationStatus

	// The status message.
	ValidationStatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRestoreJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRestoreJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRestoreJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRestoreJob"); err != nil {
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
	if err = addOpDescribeRestoreJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRestoreJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRestoreJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRestoreJob",
	}
}
