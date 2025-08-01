// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the transition lifecycle of a recovery point.
//
// The lifecycle defines when a protected resource is transitioned to cold storage
// and when it expires. Backup transitions and expires backups automatically
// according to the lifecycle that you define.
//
// Resource types that can transition to cold storage are listed in the [Feature availability by resource] table.
// Backup ignores this expression for other resource types.
//
// Backups transitioned to cold storage must be stored in cold storage for a
// minimum of 90 days. Therefore, the “retention” setting must be 90 days greater
// than the “transition to cold after days” setting. The “transition to cold after
// days” setting cannot be changed after a backup has been transitioned to cold.
//
// If your lifecycle currently uses the parameters DeleteAfterDays and
// MoveToColdStorageAfterDays , include these parameters and their values when you
// call this operation. Not including them may result in your plan updating with
// null values.
//
// This operation does not support continuous backups.
//
// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
func (c *Client) UpdateRecoveryPointLifecycle(ctx context.Context, params *UpdateRecoveryPointLifecycleInput, optFns ...func(*Options)) (*UpdateRecoveryPointLifecycleOutput, error) {
	if params == nil {
		params = &UpdateRecoveryPointLifecycleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecoveryPointLifecycle", params, optFns, c.addOperationUpdateRecoveryPointLifecycleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecoveryPointLifecycleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecoveryPointLifecycleInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created.
	//
	// This member is required.
	BackupVaultName *string

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	//
	// This member is required.
	RecoveryPointArn *string

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. Backup transitions and expires backups automatically
	// according to the lifecycle that you define.
	//
	// Backups transitioned to cold storage must be stored in cold storage for a
	// minimum of 90 days. Therefore, the “retention” setting must be 90 days greater
	// than the “transition to cold after days” setting. The “transition to cold after
	// days” setting cannot be changed after a backup has been transitioned to cold.
	Lifecycle *types.Lifecycle

	noSmithyDocumentSerde
}

type UpdateRecoveryPointLifecycleOutput struct {

	// An ARN that uniquely identifies a backup vault; for example,
	// arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault .
	BackupVaultArn *string

	// A CalculatedLifecycle object containing DeleteAt and MoveToColdStorageAt
	// timestamps.
	CalculatedLifecycle *types.CalculatedLifecycle

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. Backup transitions and expires backups automatically
	// according to the lifecycle that you define.
	//
	// Backups transitioned to cold storage must be stored in cold storage for a
	// minimum of 90 days. Therefore, the “retention” setting must be 90 days greater
	// than the “transition to cold after days” setting. The “transition to cold after
	// days” setting cannot be changed after a backup has been transitioned to cold.
	//
	// Resource types that can transition to cold storage are listed in the [Feature availability by resource] table.
	// Backup ignores this expression for other resource types.
	//
	// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
	Lifecycle *types.Lifecycle

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRecoveryPointLifecycleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecoveryPointLifecycle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecoveryPointLifecycle{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRecoveryPointLifecycle"); err != nil {
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
	if err = addOpUpdateRecoveryPointLifecycleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecoveryPointLifecycle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecoveryPointLifecycle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRecoveryPointLifecycle",
	}
}
