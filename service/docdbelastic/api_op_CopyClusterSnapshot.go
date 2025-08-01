// Code generated by smithy-go-codegen DO NOT EDIT.

package docdbelastic

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies a snapshot of an elastic cluster.
func (c *Client) CopyClusterSnapshot(ctx context.Context, params *CopyClusterSnapshotInput, optFns ...func(*Options)) (*CopyClusterSnapshotOutput, error) {
	if params == nil {
		params = &CopyClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyClusterSnapshot", params, optFns, c.addOperationCopyClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyClusterSnapshotInput struct {

	// The Amazon Resource Name (ARN) identifier of the elastic cluster snapshot.
	//
	// This member is required.
	SnapshotArn *string

	// The identifier of the new elastic cluster snapshot to create from the source
	// cluster snapshot. This parameter is not case sensitive.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: elastic-cluster-snapshot-5
	//
	// This member is required.
	TargetSnapshotName *string

	// Set to true to copy all tags from the source cluster snapshot to the target
	// elastic cluster snapshot. The default is false .
	CopyTags *bool

	// The Amazon Web Services KMS key ID for an encrypted elastic cluster snapshot.
	// The Amazon Web Services KMS key ID is the Amazon Resource Name (ARN), Amazon Web
	// Services KMS key identifier, or the Amazon Web Services KMS key alias for the
	// Amazon Web Services KMS encryption key.
	//
	// If you copy an encrypted elastic cluster snapshot from your Amazon Web Services
	// account, you can specify a value for KmsKeyId to encrypt the copy with a new
	// Amazon Web ServicesS KMS encryption key. If you don't specify a value for
	// KmsKeyId , then the copy of the elastic cluster snapshot is encrypted with the
	// same AWS KMS key as the source elastic cluster snapshot.
	//
	// To copy an encrypted elastic cluster snapshot to another Amazon Web Services
	// region, set KmsKeyId to the Amazon Web Services KMS key ID that you want to use
	// to encrypt the copy of the elastic cluster snapshot in the destination region.
	// Amazon Web Services KMS encryption keys are specific to the Amazon Web Services
	// region that they are created in, and you can't use encryption keys from one
	// Amazon Web Services region in another Amazon Web Services region.
	//
	// If you copy an unencrypted elastic cluster snapshot and specify a value for the
	// KmsKeyId parameter, an error is returned.
	KmsKeyId *string

	// The tags to be assigned to the elastic cluster snapshot.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CopyClusterSnapshotOutput struct {

	// Returns information about a specific elastic cluster snapshot.
	//
	// This member is required.
	Snapshot *types.ClusterSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCopyClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCopyClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyClusterSnapshot"); err != nil {
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
	if err = addOpCopyClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyClusterSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyClusterSnapshot",
	}
}
