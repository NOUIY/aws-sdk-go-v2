// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration profile, which is information that enables AppConfig to
// access the configuration source. Valid configuration sources include the
// following:
//
//   - Configuration data in YAML, JSON, and other formats stored in the AppConfig
//     hosted configuration store
//
//   - Configuration data stored as objects in an Amazon Simple Storage Service
//     (Amazon S3) bucket
//
//   - Pipelines stored in CodePipeline
//
//   - Secrets stored in Secrets Manager
//
//   - Standard and secure string parameters stored in Amazon Web Services Systems
//     Manager Parameter Store
//
//   - Configuration data in SSM documents stored in the Systems Manager document
//     store
//
// A configuration profile includes the following information:
//
//   - The URI location of the configuration data.
//
//   - The Identity and Access Management (IAM) role that provides access to the
//     configuration data.
//
//   - A validator for the configuration data. Available validators include either
//     a JSON Schema or an Amazon Web Services Lambda function.
//
// For more information, see [Create a Configuration and a Configuration Profile] in the AppConfig User Guide.
//
// [Create a Configuration and a Configuration Profile]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html
func (c *Client) CreateConfigurationProfile(ctx context.Context, params *CreateConfigurationProfileInput, optFns ...func(*Options)) (*CreateConfigurationProfileOutput, error) {
	if params == nil {
		params = &CreateConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationProfile", params, optFns, c.addOperationCreateConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfigurationProfileInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// A URI to locate the configuration. You can specify the following:
	//
	//   - For the AppConfig hosted configuration store and for feature flags, specify
	//   hosted .
	//
	//   - For an Amazon Web Services Systems Manager Parameter Store parameter,
	//   specify either the parameter name in the format ssm-parameter:// or the ARN.
	//
	//   - For an Amazon Web Services CodePipeline pipeline, specify the URI in the
	//   following format: codepipeline ://.
	//
	//   - For an Secrets Manager secret, specify the URI in the following format:
	//   secretsmanager ://.
	//
	//   - For an Amazon S3 object, specify the URI in the following format: s3:/// .
	//   Here is an example: s3://amzn-s3-demo-bucket/my-app/us-east-1/my-config.json
	//
	//   - For an SSM document, specify either the document name in the format
	//   ssm-document:// or the Amazon Resource Name (ARN).
	//
	// This member is required.
	LocationUri *string

	// A name for the configuration profile.
	//
	// This member is required.
	Name *string

	// A description of the configuration profile.
	Description *string

	// The identifier for an Key Management Service key to encrypt new configuration
	// data versions in the AppConfig hosted configuration store. This attribute is
	// only used for hosted configuration types. The identifier can be an KMS key ID,
	// alias, or the Amazon Resource Name (ARN) of the key ID or alias. To encrypt data
	// managed in other configuration stores, see the documentation for how to specify
	// an KMS key for that particular service.
	KmsKeyIdentifier *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri .
	//
	// A retrieval role ARN is not required for configurations stored in CodePipeline
	// or the AppConfig hosted configuration store. It is required for all other
	// sources that store your configuration.
	RetrievalRoleArn *string

	// Metadata to assign to the configuration profile. Tags help organize and
	// categorize your AppConfig resources. Each tag consists of a key and an optional
	// value, both of which you define.
	Tags map[string]string

	// The type of configurations contained in the profile. AppConfig supports feature
	// flags and freeform configurations. We recommend you create feature flag
	// configurations to enable or disable new features and freeform configurations to
	// distribute configurations to an application. When calling this API, enter one of
	// the following values for Type :
	//
	//     AWS.AppConfig.FeatureFlags
	//
	//     AWS.Freeform
	Type *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	noSmithyDocumentSerde
}

type CreateConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile description.
	Description *string

	// The configuration profile ID.
	Id *string

	// The Amazon Resource Name of the Key Management Service key to encrypt new
	// configuration data versions in the AppConfig hosted configuration store. This
	// attribute is only used for hosted configuration types. To encrypt data managed
	// in other configuration stores, see the documentation for how to specify an KMS
	// key for that particular service.
	KmsKeyArn *string

	// The Key Management Service key identifier (key ID, key alias, or key ARN)
	// provided when the resource was created or updated.
	KmsKeyIdentifier *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri .
	RetrievalRoleArn *string

	// The type of configurations contained in the profile. AppConfig supports feature
	// flags and freeform configurations. We recommend you create feature flag
	// configurations to enable or disable new features and freeform configurations to
	// distribute configurations to an application. When calling this API, enter one of
	// the following values for Type :
	//
	//     AWS.AppConfig.FeatureFlags
	//
	//     AWS.Freeform
	Type *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConfigurationProfile"); err != nil {
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
	if err = addOpCreateConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConfigurationProfile",
	}
}
