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

// Restores an elastic cluster from a snapshot.
func (c *Client) RestoreClusterFromSnapshot(ctx context.Context, params *RestoreClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreClusterFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreClusterFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreClusterFromSnapshot", params, optFns, c.addOperationRestoreClusterFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreClusterFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreClusterFromSnapshotInput struct {

	// The name of the elastic cluster.
	//
	// This member is required.
	ClusterName *string

	// The ARN identifier of the elastic cluster snapshot.
	//
	// This member is required.
	SnapshotArn *string

	// The KMS key identifier to use to encrypt the new Amazon DocumentDB elastic
	// clusters cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a cluster using the same Amazon account that owns this
	// KMS encryption key, you can use the KMS key alias instead of the ARN as the KMS
	// encryption key.
	//
	// If an encryption key is not specified here, Amazon DocumentDB uses the default
	// encryption key that KMS creates for your account. Your account has a different
	// default encryption key for each Amazon Region.
	KmsKeyId *string

	// The capacity of each shard in the new restored elastic cluster.
	ShardCapacity *int32

	// The number of replica instances applying to all shards in the elastic cluster.
	// A shardInstanceCount value of 1 means there is one writer instance, and any
	// additional instances are replicas that can be used for reads and to improve
	// availability.
	ShardInstanceCount *int32

	// The Amazon EC2 subnet IDs for the elastic cluster.
	SubnetIds []string

	// A list of the tag names to be assigned to the restored elastic cluster, in the
	// form of an array of key-value pairs in which the key is the tag name and the
	// value is the key value.
	Tags map[string]string

	// A list of EC2 VPC security groups to associate with the elastic cluster.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreClusterFromSnapshotOutput struct {

	// Returns information about a the restored elastic cluster.
	//
	// This member is required.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreClusterFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRestoreClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRestoreClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreClusterFromSnapshot"); err != nil {
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
	if err = addOpRestoreClusterFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreClusterFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreClusterFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreClusterFromSnapshot",
	}
}
