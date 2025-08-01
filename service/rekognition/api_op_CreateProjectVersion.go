// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of Amazon Rekognition project (like a Custom Labels model
// or a custom adapter) and begins training. Models and adapters are managed as
// part of a Rekognition project. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the project version.
//
// The FeatureConfig operation argument allows you to configure specific model or
// adapter settings. You can provide a description to the project version by using
// the VersionDescription argment. Training can take a while to complete. You can
// get the current status by calling DescribeProjectVersions. Training completed successfully if the
// value of the Status field is TRAINING_COMPLETED . Once training has successfully
// completed, call DescribeProjectVersionsto get the training results and evaluate the model.
//
// This operation requires permissions to perform the
// rekognition:CreateProjectVersion action.
//
// The following applies only to projects with Amazon Rekognition Custom Labels as
// the chosen feature:
//
// You can train a model in a project that doesn't have associated datasets by
// specifying manifest files in the TrainingData and TestingData fields.
//
// If you open the console after training a model with manifest files, Amazon
// Rekognition Custom Labels creates the datasets for you using the most recent
// manifest files. You can no longer train a model version for the project by
// specifying manifest files.
//
// Instead of training with a project without associated datasets, we recommend
// that you use the manifest files to create training and test datasets for the
// project.
func (c *Client) CreateProjectVersion(ctx context.Context, params *CreateProjectVersionInput, optFns ...func(*Options)) (*CreateProjectVersionOutput, error) {
	if params == nil {
		params = &CreateProjectVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProjectVersion", params, optFns, c.addOperationCreateProjectVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectVersionInput struct {

	// The Amazon S3 bucket location to store the results of training. The bucket can
	// be any S3 bucket in your AWS account. You need s3:PutObject permission on the
	// bucket.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The ARN of the Amazon Rekognition project that will manage the project version
	// you want to train.
	//
	// This member is required.
	ProjectArn *string

	// A name for the version of the project version. This value must be unique.
	//
	// This member is required.
	VersionName *string

	// Feature-specific configuration of the training job. If the job configuration
	// does not match the feature type associated with the project, an
	// InvalidParameterException is returned.
	FeatureConfig *types.CustomizationFeatureConfig

	// The identifier for your AWS Key Management Service key (AWS KMS key). You can
	// supply the Amazon Resource Name (ARN) of your KMS key, the ID of your KMS key,
	// an alias for your KMS key, or an alias ARN. The key is used to encrypt training
	// images, test images, and manifest files copied into the service for the project
	// version. Your source images are unaffected. The key is also used to encrypt
	// training results and manifest files written to the output Amazon S3 bucket (
	// OutputConfig ).
	//
	// If you choose to use your own KMS key, you need the following permissions on
	// the KMS key.
	//
	//   - kms:CreateGrant
	//
	//   - kms:DescribeKey
	//
	//   - kms:GenerateDataKey
	//
	//   - kms:Decrypt
	//
	// If you don't specify a value for KmsKeyId , images copied into the service are
	// encrypted using a key that AWS owns and manages.
	KmsKeyId *string

	//  A set of tags (key-value pairs) that you want to attach to the project
	// version.
	Tags map[string]string

	// Specifies an external manifest that the service uses to test the project
	// version. If you specify TestingData you must also specify TrainingData . The
	// project must not have any associated datasets.
	TestingData *types.TestingData

	// Specifies an external manifest that the services uses to train the project
	// version. If you specify TrainingData you must also specify TestingData . The
	// project must not have any associated datasets.
	TrainingData *types.TrainingData

	// A description applied to the project version being created.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateProjectVersionOutput struct {

	// The ARN of the model or the project version that was created. Use
	// DescribeProjectVersion to get the current status of the training operation.
	ProjectVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProjectVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProjectVersion"); err != nil {
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
	if err = addOpCreateProjectVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProjectVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProjectVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProjectVersion",
	}
}
