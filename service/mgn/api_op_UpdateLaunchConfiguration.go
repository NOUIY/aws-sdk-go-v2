// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates multiple LaunchConfigurations by Source Server ID.
//
// bootMode valid values are LEGACY_BIOS | UEFI
func (c *Client) UpdateLaunchConfiguration(ctx context.Context, params *UpdateLaunchConfigurationInput, optFns ...func(*Options)) (*UpdateLaunchConfigurationOutput, error) {
	if params == nil {
		params = &UpdateLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLaunchConfiguration", params, optFns, c.addOperationUpdateLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLaunchConfigurationInput struct {

	// Update Launch configuration by Source Server ID request.
	//
	// This member is required.
	SourceServerID *string

	// Update Launch configuration Account ID.
	AccountID *string

	// Update Launch configuration boot mode request.
	BootMode types.BootMode

	// Update Launch configuration copy Private IP request.
	CopyPrivateIp *bool

	// Update Launch configuration copy Tags request.
	CopyTags *bool

	// Enable map auto tagging.
	EnableMapAutoTagging *bool

	// Update Launch configuration launch disposition request.
	LaunchDisposition types.LaunchDisposition

	// Update Launch configuration licensing request.
	Licensing *types.Licensing

	// Launch configuration map auto tagging MPE ID.
	MapAutoTaggingMpeID *string

	// Update Launch configuration name request.
	Name *string

	// Post Launch Actions to executed on the Test or Cutover instance.
	PostLaunchActions *types.PostLaunchActions

	// Update Launch configuration Target instance right sizing request.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	noSmithyDocumentSerde
}

type UpdateLaunchConfigurationOutput struct {

	// Launch configuration boot mode.
	BootMode types.BootMode

	// Copy Private IP during Launch Configuration.
	CopyPrivateIp *bool

	// Copy Tags during Launch Configuration.
	CopyTags *bool

	// Launch configuration EC2 Launch template ID.
	Ec2LaunchTemplateID *string

	// Enable map auto tagging.
	EnableMapAutoTagging *bool

	// Launch disposition for launch configuration.
	LaunchDisposition types.LaunchDisposition

	// Launch configuration OS licensing.
	Licensing *types.Licensing

	// Map auto tagging MPE ID.
	MapAutoTaggingMpeID *string

	// Launch configuration name.
	Name *string

	// Post Launch Actions to executed on the Test or Cutover instance.
	PostLaunchActions *types.PostLaunchActions

	// Launch configuration Source Server ID.
	SourceServerID *string

	// Launch configuration Target instance type right sizing method.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLaunchConfiguration"); err != nil {
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
	if err = addOpUpdateLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLaunchConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLaunchConfiguration",
	}
}
