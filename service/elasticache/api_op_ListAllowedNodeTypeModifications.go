// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available node types that you can scale with your cluster's
// replication group's current node type.
//
// When you use the ModifyCacheCluster or ModifyReplicationGroup operations to
// scale your cluster or replication group, the value of the CacheNodeType
// parameter must be one of the node types returned by this operation.
func (c *Client) ListAllowedNodeTypeModifications(ctx context.Context, params *ListAllowedNodeTypeModificationsInput, optFns ...func(*Options)) (*ListAllowedNodeTypeModificationsOutput, error) {
	if params == nil {
		params = &ListAllowedNodeTypeModificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAllowedNodeTypeModifications", params, optFns, c.addOperationListAllowedNodeTypeModificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAllowedNodeTypeModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input parameters for the ListAllowedNodeTypeModifications operation.
type ListAllowedNodeTypeModificationsInput struct {

	// The name of the cluster you want to scale up to a larger node instanced type.
	// ElastiCache uses the cluster id to identify the current node type of this
	// cluster and from that to create a list of node types you can scale up to.
	//
	// You must provide a value for either the CacheClusterId or the ReplicationGroupId
	// .
	CacheClusterId *string

	// The name of the replication group want to scale up to a larger node type.
	// ElastiCache uses the replication group id to identify the current node type
	// being used by this replication group, and from that to create a list of node
	// types you can scale up to.
	//
	// You must provide a value for either the CacheClusterId or the ReplicationGroupId
	// .
	ReplicationGroupId *string

	noSmithyDocumentSerde
}

// Represents the allowed node types you can use to modify your cluster or
// replication group.
type ListAllowedNodeTypeModificationsOutput struct {

	// A string list, each element of which specifies a cache node type which you can
	// use to scale your cluster or replication group. When scaling down a Valkey or
	// Redis OSS cluster or replication group using ModifyCacheCluster or
	// ModifyReplicationGroup, use a value from this list for the CacheNodeType
	// parameter.
	ScaleDownModifications []string

	// A string list, each element of which specifies a cache node type which you can
	// use to scale your cluster or replication group.
	//
	// When scaling up a Valkey or Redis OSS cluster or replication group using
	// ModifyCacheCluster or ModifyReplicationGroup , use a value from this list for
	// the CacheNodeType parameter.
	ScaleUpModifications []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAllowedNodeTypeModificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListAllowedNodeTypeModifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListAllowedNodeTypeModifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAllowedNodeTypeModifications"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAllowedNodeTypeModifications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAllowedNodeTypeModifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAllowedNodeTypeModifications",
	}
}
