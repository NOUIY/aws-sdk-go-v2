// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an environment.
//
// If the environment is associated with an environment account connection, don't
// update or include the protonServiceRoleArn and provisioningRepository parameter
// to update or connect to an environment account connection.
//
// You can only update to a new environment account connection if that connection
// was created in the same environment account that the current environment account
// connection was created in. The account connection must also be associated with
// the current environment.
//
// If the environment isn't associated with an environment account connection,
// don't update or include the environmentAccountConnectionId parameter. You can't
// update or connect the environment to an environment account connection if it
// isn't already associated with an environment connection.
//
// You can update either the environmentAccountConnectionId or protonServiceRoleArn
// parameter and value. You can’t update both.
//
// If the environment was configured for Amazon Web Services-managed provisioning,
// omit the provisioningRepository parameter.
//
// If the environment was configured for self-managed provisioning, specify the
// provisioningRepository parameter and omit the protonServiceRoleArn and
// environmentAccountConnectionId parameters.
//
// For more information, see [Environments] and [Provisioning methods] in the Proton User Guide.
//
// There are four modes for updating an environment. The deploymentType field
// defines the mode.
//
//	NONE
//
// In this mode, a deployment doesn't occur. Only the requested metadata
// parameters are updated.
//
//	CURRENT_VERSION
//
// In this mode, the environment is deployed and updated with the new spec that
// you provide. Only requested parameters are updated. Don’t include minor or major
// version parameters when you use this deployment-type .
//
//	MINOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) minor version of the current major version in use, by
// default. You can also specify a different minor version of the current major
// version in use.
//
//	MAJOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) major and minor version of the current template, by
// default. You can also specify a different major version that's higher than the
// major version in use and a minor version.
//
// [Environments]: https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html
// [Provisioning methods]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironment", params, optFns, c.addOperationUpdateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentInput struct {

	// There are four modes for updating an environment. The deploymentType field
	// defines the mode.
	//
	//     NONE
	//
	// In this mode, a deployment doesn't occur. Only the requested metadata
	// parameters are updated.
	//
	//     CURRENT_VERSION
	//
	// In this mode, the environment is deployed and updated with the new spec that
	// you provide. Only requested parameters are updated. Don’t include major or minor
	// version parameters when you use this deployment-type .
	//
	//     MINOR_VERSION
	//
	// In this mode, the environment is deployed and updated with the published,
	// recommended (latest) minor version of the current major version in use, by
	// default. You can also specify a different minor version of the current major
	// version in use.
	//
	//     MAJOR_VERSION
	//
	// In this mode, the environment is deployed and updated with the published,
	// recommended (latest) major and minor version of the current template, by
	// default. You can also specify a different major version that is higher than the
	// major version in use and a minor version (optional).
	//
	// This member is required.
	DeploymentType types.DeploymentUpdateType

	// The name of the environment to update.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM service role that allows Proton to
	// provision infrastructure using CodeBuild-based provisioning on your behalf.
	CodebuildRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in this environment. It determines the
	// scope of infrastructure that a component can provision.
	//
	// The environment must have a componentRoleArn to allow directly defined
	// components to be associated with the environment.
	//
	// For more information about components, see [Proton components] in the Proton User Guide.
	//
	// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
	ComponentRoleArn *string

	// A description of the environment update.
	Description *string

	// The ID of the environment account connection.
	//
	// You can only update to a new environment account connection if it was created
	// in the same environment account that the current environment account connection
	// was created in and is associated with the current environment.
	EnvironmentAccountConnectionId *string

	// The Amazon Resource Name (ARN) of the Proton service role that allows Proton to
	// make API calls to other services your behalf.
	ProtonServiceRoleArn *string

	// The linked repository that you use to host your rendered infrastructure
	// templates for self-managed provisioning. A linked repository is a repository
	// that has been registered with Proton. For more information, see CreateRepository.
	ProvisioningRepository *types.RepositoryBranchInput

	// The formatted specification that defines the update.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	// The major version of the environment to update.
	TemplateMajorVersion *string

	// The minor version of the environment to update.
	TemplateMinorVersion *string

	noSmithyDocumentSerde
}

type UpdateEnvironmentOutput struct {

	// The environment detail data that's returned by Proton.
	//
	// This member is required.
	Environment *types.Environment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnvironment"); err != nil {
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
	if err = addOpUpdateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnvironment",
	}
}
