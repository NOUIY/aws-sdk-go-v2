// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon EKS cluster configuration. Your cluster continues to function
// during the update. The response output includes an update ID that you can use to
// track the status of your cluster update with DescribeUpdate .
//
// You can use this operation to do the following actions:
//
//   - You can use this API operation to enable or disable exporting the
//     Kubernetes control plane logs for your cluster to CloudWatch Logs. By default,
//     cluster control plane logs aren't exported to CloudWatch Logs. For more
//     information, see [Amazon EKS Cluster control plane logs]in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
//
//	exported control plane logs. For more information, see [CloudWatch Pricing].
//
//	- You can also use this API operation to enable or disable public and private
//	access to your cluster's Kubernetes API server endpoint. By default, public
//	access is enabled, and private access is disabled. For more information, see [Cluster API server endpoint]
//	in the Amazon EKS User Guide .
//
//	- You can also use this API operation to choose different subnets and
//	security groups for the cluster. You must specify at least two subnets that are
//	in different Availability Zones. You can't change which VPC the subnets are
//	from, the subnets must be in the same VPC as the subnets that the cluster was
//	created with. For more information about the VPC requirements, see [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]in the
//	Amazon EKS User Guide .
//
//	- You can also use this API operation to enable or disable ARC zonal shift.
//	If zonal shift is enabled, Amazon Web Services configures zonal autoshift for
//	the cluster.
//
//	- You can also use this API operation to add, change, or remove the
//	configuration in the cluster for EKS Hybrid Nodes. To remove the configuration,
//	use the remoteNetworkConfig key with an object containing both subkeys with
//	empty arrays for each. Here is an inline example: "remoteNetworkConfig": {
//	"remoteNodeNetworks": [], "remotePodNetworks": [] } .
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful ), the cluster status moves to Active .
//
// [Amazon EKS Cluster control plane logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
//
// [Cluster API server endpoint]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html
func (c *Client) UpdateClusterConfig(ctx context.Context, params *UpdateClusterConfigInput, optFns ...func(*Options)) (*UpdateClusterConfigOutput, error) {
	if params == nil {
		params = &UpdateClusterConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClusterConfig", params, optFns, c.addOperationUpdateClusterConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterConfigInput struct {

	// The name of the Amazon EKS cluster to update.
	//
	// This member is required.
	Name *string

	// The access configuration for the cluster.
	AccessConfig *types.UpdateAccessConfigRequest

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// Update the configuration of the compute capability of your EKS Auto Mode
	// cluster. For example, enable the capability.
	ComputeConfig *types.ComputeConfigRequest

	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig *types.KubernetesNetworkConfigRequest

	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs . By default, cluster control plane logs aren't exported to
	// CloudWatch Logs . For more information, see [Amazon EKS cluster control plane logs]in the Amazon EKS User Guide .
	//
	// CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
	// exported control plane logs. For more information, see [CloudWatch Pricing].
	//
	// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
	// [Amazon EKS cluster control plane logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
	Logging *types.Logging

	// The configuration in the cluster for EKS Hybrid Nodes. You can add, change, or
	// remove this configuration after the cluster is created.
	RemoteNetworkConfig *types.RemoteNetworkConfigRequest

	// An object representing the VPC configuration to use for an Amazon EKS cluster.
	ResourcesVpcConfig *types.VpcConfigRequest

	// Update the configuration of the block storage capability of your EKS Auto Mode
	// cluster. For example, enable the capability.
	StorageConfig *types.StorageConfigRequest

	// You can enable or disable extended support for clusters currently on standard
	// support. You cannot disable extended support once it starts. You must enable
	// extended support before your cluster exits standard support.
	UpgradePolicy *types.UpgradePolicyRequest

	// Enable or disable ARC zonal shift for the cluster. If zonal shift is enabled,
	// Amazon Web Services configures zonal autoshift for the cluster.
	//
	// Zonal shift is a feature of Amazon Application Recovery Controller (ARC). ARC
	// zonal shift is designed to be a temporary measure that allows you to move
	// traffic for a resource away from an impaired AZ until the zonal shift expires or
	// you cancel it. You can extend the zonal shift if necessary.
	//
	// You can start a zonal shift for an EKS cluster, or you can allow Amazon Web
	// Services to do it for you by enabling zonal autoshift. This shift updates the
	// flow of east-to-west network traffic in your cluster to only consider network
	// endpoints for Pods running on worker nodes in healthy AZs. Additionally, any ALB
	// or NLB handling ingress traffic for applications in your EKS cluster will
	// automatically route traffic to targets in the healthy AZs. For more information
	// about zonal shift in EKS, see [Learn about Amazon Application Recovery Controller (ARC) Zonal Shift in Amazon EKS]in the Amazon EKS User Guide .
	//
	// [Learn about Amazon Application Recovery Controller (ARC) Zonal Shift in Amazon EKS]: https://docs.aws.amazon.com/eks/latest/userguide/zone-shift.html
	ZonalShiftConfig *types.ZonalShiftConfigRequest

	noSmithyDocumentSerde
}

type UpdateClusterConfigOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClusterConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClusterConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateClusterConfig"); err != nil {
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
	if err = addIdempotencyToken_opUpdateClusterConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateClusterConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterConfig(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateClusterConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateClusterConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateClusterConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateClusterConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateClusterConfigInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateClusterConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateClusterConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateClusterConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateClusterConfig",
	}
}
