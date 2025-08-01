// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a [version] from the current code and configuration of a function. Use versions
// to create a snapshot of your function code and configuration that doesn't
// change.
//
// Lambda doesn't publish a version if the function's configuration and code
// haven't changed since the last version. Use UpdateFunctionCodeor UpdateFunctionConfiguration to update the function before
// publishing a version.
//
// Clients can invoke versions directly or with an alias. To create an alias, use CreateAlias.
//
// [version]: https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html
func (c *Client) PublishVersion(ctx context.Context, params *PublishVersionInput, optFns ...func(*Options)) (*PublishVersionOutput, error) {
	if params == nil {
		params = &PublishVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishVersion", params, optFns, c.addOperationPublishVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishVersionInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name - MyFunction .
	//
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction .
	//
	//   - Partial ARN - 123456789012:function:MyFunction .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Only publish a version if the hash value matches the value that's specified.
	// Use this option to avoid publishing a version if the function code has changed
	// since you last updated it. You can get the hash for the version that you
	// uploaded from the output of UpdateFunctionCode.
	CodeSha256 *string

	// A description for the version to override the description in the function
	// configuration.
	Description *string

	// Only update the function if the revision ID matches the ID that's specified.
	// Use this option to avoid publishing a version if the function configuration has
	// changed since you last updated it.
	RevisionId *string

	noSmithyDocumentSerde
}

// Details about a function's configuration.
type PublishVersionOutput struct {

	// The instruction set architecture that the function supports. Architecture is a
	// string array with one of the valid values. The default architecture value is
	// x86_64 .
	Architectures []types.Architecture

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's [environment variables]. Omitted from CloudTrail logs.
	//
	// [environment variables]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
	Environment *types.EnvironmentResponse

	// The size of the function's /tmp directory in MB. The default value is 512, but
	// can be any whole number between 512 and 10,240 MB. For more information, see [Configuring ephemeral storage (console)].
	//
	// [Configuring ephemeral storage (console)]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-ephemeral-storage
	EphemeralStorage *types.EphemeralStorage

	// Connection settings for an [Amazon EFS file system].
	//
	// [Amazon EFS file system]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html
	FileSystemConfigs []types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin running your function.
	Handler *string

	// The function's image configuration values.
	ImageConfigResponse *types.ImageConfigResponse

	// The ARN of the Key Management Service (KMS) customer managed key that's used to
	// encrypt the following resources:
	//
	//   - The function's [environment variables].
	//
	//   - The function's [Lambda SnapStart]snapshots.
	//
	//   - When used with SourceKMSKeyArn , the unzipped version of the .zip deployment
	//   package that's used for function invocations. For more information, see [Specifying a customer managed key for Lambda].
	//
	//   - The optimized version of the container image that's used for function
	//   invocations. Note that this is not the same key that's used to protect your
	//   container image in the Amazon Elastic Container Registry (Amazon ECR). For more
	//   information, see [Function lifecycle].
	//
	// If you don't provide a customer managed key, Lambda uses an [Amazon Web Services owned key] or an [Amazon Web Services managed key].
	//
	// [Amazon Web Services owned key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk
	// [Specifying a customer managed key for Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/encrypt-zip-package.html#enable-zip-custom-encryption
	// [Lambda SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html
	// [environment variables]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption
	// [Function lifecycle]: https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-lifecycle
	// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
	KMSKeyArn *string

	// The date and time that the function was last updated, in [ISO-8601 format]
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// [ISO-8601 format]: https://www.w3.org/TR/NOTE-datetime
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's [layers].
	//
	// [layers]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
	Layers []types.Layer

	// The function's Amazon CloudWatch Logs configuration settings.
	LoggingConfig *types.LoggingConfig

	// For Lambda@Edge functions, the ARN of the main function.
	MasterArn *string

	// The amount of memory available to the function at runtime.
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set Zip
	// for .zip file archive.
	PackageType types.PackageType

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The identifier of the function's [runtime]. Runtime is required if the deployment
	// package is a .zip file archive. Specifying a runtime results in an error if
	// you're deploying a function using a container image.
	//
	// The following list includes deprecated runtimes. Lambda blocks creating new
	// functions and updating existing functions shortly after each runtime is
	// deprecated. For more information, see [Runtime use after deprecation].
	//
	// For a list of all currently supported runtimes, see [Supported runtimes].
	//
	// [Runtime use after deprecation]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-deprecation-levels
	// [runtime]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
	// [Supported runtimes]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtimes-supported
	Runtime types.Runtime

	// The ARN of the runtime and any errors that occured.
	RuntimeVersionConfig *types.RuntimeVersionConfig

	// The ARN of the signing job.
	SigningJobArn *string

	// The ARN of the signing profile version.
	SigningProfileVersionArn *string

	// Set ApplyOn to PublishedVersions to create a snapshot of the initialized
	// execution environment when you publish a function version. For more information,
	// see [Improving startup performance with Lambda SnapStart].
	//
	// [Improving startup performance with Lambda SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html
	SnapStart *types.SnapStartResponse

	// The current state of the function. When the state is Inactive , you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating ,
	// you can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublishVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PublishVersion"); err != nil {
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
	if err = addOpPublishVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PublishVersion",
	}
}
