// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a deployment.
func (c *Client) StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) {
	if params == nil {
		params = &StartDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeployment", params, optFns, c.addOperationStartDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeploymentInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The configuration version to deploy. If deploying an AppConfig hosted
	// configuration version, you can specify either the version number or version
	// label. For all other configurations, you must specify the version number.
	//
	// This member is required.
	ConfigurationVersion *string

	// The deployment strategy ID.
	//
	// This member is required.
	DeploymentStrategyId *string

	// The environment ID.
	//
	// This member is required.
	EnvironmentId *string

	// A description of the deployment.
	Description *string

	// A map of dynamic extension parameter names to values to pass to associated
	// extensions with PRE_START_DEPLOYMENT actions.
	DynamicExtensionParameters map[string]string

	// The KMS key identifier (key ID, key alias, or key ARN). AppConfig uses this ID
	// to encrypt the configuration data using a customer managed key.
	KmsKeyIdentifier *string

	// Metadata to assign to the deployment. Tags help organize and categorize your
	// AppConfig resources. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartDeploymentOutput struct {

	// The ID of the application that was deployed.
	ApplicationId *string

	// A list of extensions that were processed as part of the deployment. The
	// extensions that were previously associated to the configuration profile,
	// environment, or the application when StartDeployment was called.
	AppliedExtensions []types.AppliedExtension

	// The time the deployment completed.
	CompletedAt *time.Time

	// Information about the source location of the configuration.
	ConfigurationLocationUri *string

	// The name of the configuration.
	ConfigurationName *string

	// The ID of the configuration profile that was deployed.
	ConfigurationProfileId *string

	// The configuration version that was deployed.
	ConfigurationVersion *string

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes int32

	// The sequence number of the deployment.
	DeploymentNumber int32

	// The ID of the deployment strategy that was deployed.
	DeploymentStrategyId *string

	// The description of the deployment.
	Description *string

	// The ID of the environment that was deployed.
	EnvironmentId *string

	// A list containing all events related to a deployment. The most recent events
	// are displayed first.
	EventLog []types.DeploymentEvent

	// The amount of time that AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic rollback.
	FinalBakeTimeInMinutes int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The Amazon Resource Name of the Key Management Service key used to encrypt
	// configuration data. You can encrypt secrets stored in Secrets Manager, Amazon
	// Simple Storage Service (Amazon S3) objects encrypted with SSE-KMS, or secure
	// string parameters stored in Amazon Web Services Systems Manager Parameter Store.
	KmsKeyArn *string

	// The Key Management Service key identifier (key ID, key alias, or key ARN)
	// provided when the resource was created or updated.
	KmsKeyIdentifier *string

	// The percentage of targets for which the deployment is available.
	PercentageComplete *float32

	// The time the deployment started.
	StartedAt *time.Time

	// The state of the deployment.
	State types.DeploymentState

	// A user-defined label for an AppConfig hosted configuration version.
	VersionLabel *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartDeployment"); err != nil {
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
	if err = addOpStartDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDeployment",
	}
}
