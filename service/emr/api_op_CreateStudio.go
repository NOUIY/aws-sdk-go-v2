// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon EMR Studio.
func (c *Client) CreateStudio(ctx context.Context, params *CreateStudioInput, optFns ...func(*Options)) (*CreateStudioOutput, error) {
	if params == nil {
		params = &CreateStudioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudio", params, optFns, c.addOperationCreateStudioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStudioInput struct {

	// Specifies whether the Studio authenticates users using IAM or IAM Identity
	// Center.
	//
	// This member is required.
	AuthMode types.AuthMode

	// The Amazon S3 location to back up Amazon EMR Studio Workspaces and notebook
	// files.
	//
	// This member is required.
	DefaultS3Location *string

	// The ID of the Amazon EMR Studio Engine security group. The Engine security
	// group allows inbound network traffic from the Workspace security group, and it
	// must be in the same VPC specified by VpcId .
	//
	// This member is required.
	EngineSecurityGroupId *string

	// A descriptive name for the Amazon EMR Studio.
	//
	// This member is required.
	Name *string

	// The IAM role that the Amazon EMR Studio assumes. The service role provides a
	// way for Amazon EMR Studio to interoperate with other Amazon Web Services
	// services.
	//
	// This member is required.
	ServiceRole *string

	// A list of subnet IDs to associate with the Amazon EMR Studio. A Studio can have
	// a maximum of 5 subnets. The subnets must belong to the VPC specified by VpcId .
	// Studio users can create a Workspace in any of the specified subnets.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the
	// Studio.
	//
	// This member is required.
	VpcId *string

	// The ID of the Amazon EMR Studio Workspace security group. The Workspace
	// security group allows outbound network traffic to resources in the Engine
	// security group, and it must be in the same VPC specified by VpcId .
	//
	// This member is required.
	WorkspaceSecurityGroupId *string

	// A detailed description of the Amazon EMR Studio.
	Description *string

	// The KMS key identifier (ARN) used to encrypt Amazon EMR Studio workspace and
	// notebook files when backed up to Amazon S3.
	EncryptionKeyArn *string

	//  The ARN of the IAM Identity Center instance to create the Studio application.
	IdcInstanceArn *string

	//  Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL .
	// If the value is set to REQUIRED , users must be explicitly assigned to the
	// Studio application to access the Studio.
	IdcUserAssignment types.IdcUserAssignment

	// The authentication endpoint of your identity provider (IdP). Specify this value
	// when you use IAM authentication and want to let federated users log in to a
	// Studio with the Studio URL and credentials from your IdP. Amazon EMR Studio
	// redirects users to this endpoint to enter credentials.
	IdpAuthUrl *string

	// The name that your identity provider (IdP) uses for its RelayState parameter.
	// For example, RelayState or TargetSource . Specify this value when you use IAM
	// authentication and want to let federated users log in to a Studio using the
	// Studio URL. The RelayState parameter differs by IdP.
	IdpRelayStateParameterName *string

	// A list of tags to associate with the Amazon EMR Studio. Tags are user-defined
	// key-value pairs that consist of a required key string with a maximum of 128
	// characters, and an optional value string with a maximum of 256 characters.
	Tags []types.Tag

	//  A Boolean indicating whether to enable Trusted identity propagation for the
	// Studio. The default value is false .
	TrustedIdentityPropagationEnabled *bool

	// The IAM user role that users and groups assume when logged in to an Amazon EMR
	// Studio. Only specify a UserRole when you use IAM Identity Center
	// authentication. The permissions attached to the UserRole can be scoped down for
	// each user or group using session policies.
	UserRole *string

	noSmithyDocumentSerde
}

type CreateStudioOutput struct {

	// The ID of the Amazon EMR Studio.
	StudioId *string

	// The unique Studio access URL.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStudioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateStudio"); err != nil {
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
	if err = addOpCreateStudioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudio(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStudio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateStudio",
	}
}
