// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a SageMaker HyperPod cluster. SageMaker HyperPod is a capability of
// SageMaker for creating and managing persistent clusters for developing large
// machine learning models, such as large language models (LLMs) and diffusion
// models. To learn more, see [Amazon SageMaker HyperPod]in the Amazon SageMaker Developer Guide.
//
// [Amazon SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod.html
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The name for the new SageMaker HyperPod cluster.
	//
	// This member is required.
	ClusterName *string

	// The instance groups to be created in the SageMaker HyperPod cluster.
	InstanceGroups []types.ClusterInstanceGroupSpecification

	// The node recovery mode for the SageMaker HyperPod cluster. When set to Automatic
	// , SageMaker HyperPod will automatically reboot or replace faulty nodes when
	// issues are detected. When set to None , cluster administrators will need to
	// manually manage any faulty cluster instances.
	NodeRecovery types.ClusterNodeRecovery

	// The type of orchestrator to use for the SageMaker HyperPod cluster. Currently,
	// the only supported value is "eks" , which is to use an Amazon Elastic Kubernetes
	// Service (EKS) cluster as the orchestrator.
	Orchestrator *types.ClusterOrchestrator

	// The specialized instance groups for training models like Amazon Nova to be
	// created in the SageMaker HyperPod cluster.
	RestrictedInstanceGroups []types.ClusterRestrictedInstanceGroupSpecification

	// Custom tags for managing the SageMaker HyperPod cluster as an Amazon Web
	// Services resource. You can add tags to your cluster in the same way you add them
	// in other Amazon Web Services services that support tagging. To learn more about
	// tagging Amazon Web Services resources in general, see [Tagging Amazon Web Services Resources User Guide].
	//
	// [Tagging Amazon Web Services Resources User Guide]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html
	Tags []types.Tag

	// Specifies the Amazon Virtual Private Cloud (VPC) that is associated with the
	// Amazon SageMaker HyperPod cluster. You can control access to and from your
	// resources by configuring your VPC. For more information, see [Give SageMaker access to resources in your Amazon VPC].
	//
	// When your Amazon VPC and subnets support IPv6, network communications differ
	// based on the cluster orchestration platform:
	//
	//   - Slurm-orchestrated clusters automatically configure nodes with dual IPv6
	//   and IPv4 addresses, allowing immediate IPv6 network communications.
	//
	//   - In Amazon EKS-orchestrated clusters, nodes receive dual-stack addressing,
	//   but pods can only use IPv6 when the Amazon EKS cluster is explicitly
	//   IPv6-enabled. For information about deploying an IPv6 Amazon EKS cluster, see [Amazon EKS IPv6 Cluster Deployment]
	//   .
	//
	// Additional resources for IPv6 configuration:
	//
	//   - For information about adding IPv6 support to your VPC, see to [IPv6 Support for VPC].
	//
	//   - For information about creating a new IPv6-compatible VPC, see [Amazon VPC Creation Guide].
	//
	//   - To configure SageMaker HyperPod with a custom Amazon VPC, see [Custom Amazon VPC Setup for SageMaker HyperPod].
	//
	// [Give SageMaker access to resources in your Amazon VPC]: https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html
	// [IPv6 Support for VPC]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-migrate-ipv6.html
	// [Amazon EKS IPv6 Cluster Deployment]: https://docs.aws.amazon.com/eks/latest/userguide/deploy-ipv6-cluster.html#_deploy_an_ipv6_cluster_with_eksctl
	// [Amazon VPC Creation Guide]: https://docs.aws.amazon.com/vpc/latest/userguide/create-vpc.html
	// [Custom Amazon VPC Setup for SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-prerequisites.html#sagemaker-hyperpod-prerequisites-optional-vpc
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	//
	// This member is required.
	ClusterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCluster"); err != nil {
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCluster",
	}
}
