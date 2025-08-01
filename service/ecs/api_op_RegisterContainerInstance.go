// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for
// use outside of the agent.
//
// Registers an EC2 instance into the specified cluster. This instance becomes
// available to place containers on.
func (c *Client) RegisterContainerInstance(ctx context.Context, params *RegisterContainerInstanceInput, optFns ...func(*Options)) (*RegisterContainerInstanceOutput, error) {
	if params == nil {
		params = &RegisterContainerInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterContainerInstance", params, optFns, c.addOperationRegisterContainerInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterContainerInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterContainerInstanceInput struct {

	// The container instance attributes that this container instance supports.
	Attributes []types.Attribute

	// The short name or full Amazon Resource Name (ARN) of the cluster to register
	// your container instance with. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// The ARN of the container instance (if it was previously registered).
	ContainerInstanceArn *string

	// The instance identity document for the EC2 instance to register. This document
	// can be found by running the following command from the instance: curl
	// http://169.254.169.254/latest/dynamic/instance-identity/document/
	InstanceIdentityDocument *string

	// The instance identity document signature for the EC2 instance to register. This
	// signature can be found by running the following command from the instance: curl
	// http://169.254.169.254/latest/dynamic/instance-identity/signature/
	InstanceIdentityDocumentSignature *string

	// The devices that are available on the container instance. The only supported
	// device type is a GPU.
	PlatformDevices []types.PlatformDevice

	// The metadata that you apply to the container instance to help you categorize
	// and organize them. Each tag consists of a key and an optional value. You define
	// both.
	//
	// The following basic restrictions apply to tags:
	//
	//   - Maximum number of tags per resource - 50
	//
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//
	//   - Tag keys and values are case-sensitive.
	//
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	// The resources available on the instance.
	TotalResources []types.Resource

	// The version information for the Amazon ECS container agent and Docker daemon
	// that runs on the container instance.
	VersionInfo *types.VersionInfo

	noSmithyDocumentSerde
}

type RegisterContainerInstanceOutput struct {

	// The container instance that was registered.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterContainerInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterContainerInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterContainerInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterContainerInstance"); err != nil {
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
	if err = addOpRegisterContainerInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterContainerInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterContainerInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterContainerInstance",
	}
}
