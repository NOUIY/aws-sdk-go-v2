// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch configuration.
//
// If you exceed your maximum limit of launch configurations, the call fails. To
// query this limit, call the [DescribeAccountLimits]API. For information about updating this limit, see [Quotas for Amazon EC2 Auto Scaling]
// in the Amazon EC2 Auto Scaling User Guide.
//
// For more information, see [Launch configurations] in the Amazon EC2 Auto Scaling User Guide.
//
// Amazon EC2 Auto Scaling configures instances launched as part of an Auto
// Scaling group using either a launch template or a launch configuration. We
// strongly recommend that you do not use launch configurations. They do not
// provide full functionality for Amazon EC2 Auto Scaling or Amazon EC2. For
// information about using launch templates, see [Launch templates]in the Amazon EC2 Auto Scaling
// User Guide.
//
// [DescribeAccountLimits]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
// [Launch configurations]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html
// [Launch templates]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-templates.html
func (c *Client) CreateLaunchConfiguration(ctx context.Context, params *CreateLaunchConfigurationInput, optFns ...func(*Options)) (*CreateLaunchConfigurationOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfiguration", params, optFns, c.addOperationCreateLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationInput struct {

	// The name of the launch configuration. This name must be unique per Region per
	// account.
	//
	// This member is required.
	LaunchConfigurationName *string

	// Specifies whether to assign a public IPv4 address to the group's instances. If
	// the instance is launched into a default subnet, the default is to assign a
	// public IPv4 address, unless you disabled the option to assign a public IPv4
	// address on the subnet. If the instance is launched into a nondefault subnet, the
	// default is not to assign a public IPv4 address, unless you enabled the option to
	// assign a public IPv4 address on the subnet.
	//
	// If you specify true , each instance in the Auto Scaling group receives a unique
	// public IPv4 address. For more information, see [Provide network connectivity for your Auto Scaling instances using Amazon VPC]in the Amazon EC2 Auto Scaling
	// User Guide.
	//
	// If you specify this property, you must specify at least one subnet for
	// VPCZoneIdentifier when you create your group.
	//
	// [Provide network connectivity for your Auto Scaling instances using Amazon VPC]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html
	AssociatePublicIpAddress *bool

	// The block device mapping entries that define the block devices to attach to the
	// instances at launch. By default, the block devices specified in the block device
	// mapping for the AMI are used. For more information, see [Block device mappings]in the Amazon EC2 User
	// Guide.
	//
	// [Block device mappings]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	BlockDeviceMappings []types.BlockDeviceMapping

	// Available for backward compatibility.
	ClassicLinkVPCId *string

	// Available for backward compatibility.
	ClassicLinkVPCSecurityGroups []string

	// Specifies whether the launch configuration is optimized for EBS I/O ( true ) or
	// not ( false ). The optimization provides dedicated throughput to Amazon EBS and
	// an optimized configuration stack to provide optimal I/O performance. This
	// optimization is not available with all instance types. Additional fees are
	// incurred when you enable EBS optimization for an instance type that is not
	// EBS-optimized by default. For more information, see [Amazon EBS-optimized instances]in the Amazon EC2 User
	// Guide.
	//
	// The default value is false .
	//
	// [Amazon EBS-optimized instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html
	EbsOptimized *bool

	// The name or the Amazon Resource Name (ARN) of the instance profile associated
	// with the IAM role for the instance. The instance profile contains the IAM role.
	// For more information, see [IAM role for applications that run on Amazon EC2 instances]in the Amazon EC2 Auto Scaling User Guide.
	//
	// [IAM role for applications that run on Amazon EC2 instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html
	IamInstanceProfile *string

	// The ID of the Amazon Machine Image (AMI) that was assigned during registration.
	// For more information, see [Find a Linux AMI]in the Amazon EC2 User Guide.
	//
	// If you specify InstanceId , an ImageId is not required.
	//
	// [Find a Linux AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html
	ImageId *string

	// The ID of the instance to use to create the launch configuration. The new
	// launch configuration derives attributes from the instance, except for the block
	// device mapping.
	//
	// To create a launch configuration with a block device mapping or override any
	// other instance attributes, specify them as part of the same request.
	//
	// For more information, see [Create a launch configuration] in the Amazon EC2 Auto Scaling User Guide.
	//
	// [Create a launch configuration]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html
	InstanceId *string

	// Controls whether instances in this group are launched with detailed ( true ) or
	// basic ( false ) monitoring.
	//
	// The default value is true (enabled).
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every
	// minute and your account is charged a fee. When you disable detailed monitoring,
	// CloudWatch generates metrics every 5 minutes. For more information, see [Configure monitoring for Auto Scaling instances]in the
	// Amazon EC2 Auto Scaling User Guide.
	//
	// [Configure monitoring for Auto Scaling instances]: https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html
	InstanceMonitoring *types.InstanceMonitoring

	// Specifies the instance type of the EC2 instance. For information about
	// available instance types, see [Available instance types]in the Amazon EC2 User Guide.
	//
	// If you specify InstanceId , an InstanceType is not required.
	//
	// [Available instance types]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#AvailableInstanceTypes
	InstanceType *string

	// The ID of the kernel associated with the AMI.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see [User provided kernels]in the Amazon EC2 User Guide.
	//
	// [User provided kernels]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html
	KernelId *string

	// The name of the key pair. For more information, see [Amazon EC2 key pairs and Amazon EC2 instances] in the Amazon EC2 User
	// Guide.
	//
	// [Amazon EC2 key pairs and Amazon EC2 instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	KeyName *string

	// The metadata options for the instances. For more information, see [Configure the instance metadata options] in the
	// Amazon EC2 Auto Scaling User Guide.
	//
	// [Configure the instance metadata options]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds
	MetadataOptions *types.InstanceMetadataOptions

	// The tenancy of the instance, either default or dedicated . An instance with
	// dedicated tenancy runs on isolated, single-tenant hardware and can only be
	// launched into a VPC. To launch dedicated instances into a shared tenancy VPC (a
	// VPC with the instance placement tenancy attribute set to default ), you must set
	// the value of this property to dedicated .
	//
	// If you specify PlacementTenancy , you must specify at least one subnet for
	// VPCZoneIdentifier when you create your group.
	//
	// Valid values: default | dedicated
	PlacementTenancy *string

	// The ID of the RAM disk to select.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see [User provided kernels]in the Amazon EC2 User Guide.
	//
	// [User provided kernels]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html
	RamdiskId *string

	// A list that contains the security group IDs to assign to the instances in the
	// Auto Scaling group. For more information, see [Control traffic to your Amazon Web Services resources using security groups]in the Amazon Virtual Private
	// Cloud User Guide.
	//
	// [Control traffic to your Amazon Web Services resources using security groups]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-security-groups.html
	SecurityGroups []string

	// The maximum hourly price to be paid for any Spot Instance launched to fulfill
	// the request. Spot Instances are launched when the price you specify exceeds the
	// current Spot price. For more information, see [Request Spot Instances for fault-tolerant and flexible applications]in the Amazon EC2 Auto Scaling
	// User Guide.
	//
	// Valid Range: Minimum value of 0.001
	//
	// When you change your maximum price by creating a new launch configuration,
	// running instances will continue to run as long as the maximum price for those
	// running instances is higher than the current Spot price.
	//
	// [Request Spot Instances for fault-tolerant and flexible applications]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-template-spot-instances.html
	SpotPrice *string

	// The user data to make available to the launched EC2 instances. For more
	// information, see [Instance metadata and user data](Linux) and [Instance metadata and user data] (Windows). If you are using a command line tool,
	// base64-encoding is performed for you, and you can load the text from a file.
	// Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.
	//
	// [Instance metadata and user data]: https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-instance-metadata.html
	UserData *string

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLaunchConfiguration"); err != nil {
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
	if err = addOpCreateLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLaunchConfiguration",
	}
}
