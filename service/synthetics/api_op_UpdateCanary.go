// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of a canary that has already been created.
//
// You can't use this operation to update the tags of an existing canary. To
// change the tags of an existing canary, use [TagResource].
//
// When you use the dryRunId field when updating a canary, the only other field
// you can provide is the Schedule . Adding any other field will thrown an
// exception.
//
// [TagResource]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_TagResource.html
func (c *Client) UpdateCanary(ctx context.Context, params *UpdateCanaryInput, optFns ...func(*Options)) (*UpdateCanaryOutput, error) {
	if params == nil {
		params = &UpdateCanaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCanary", params, optFns, c.addOperationUpdateCanaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCanaryInput struct {

	// The name of the canary that you want to update. To find the names of your
	// canaries, use [DescribeCanaries].
	//
	// You cannot change the name of a canary that has already been created.
	//
	// [DescribeCanaries]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html
	//
	// This member is required.
	Name *string

	// A structure that contains the configuration for canary artifacts, including the
	// encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig *types.ArtifactConfigInput

	// The location in Amazon S3 where Synthetics stores artifacts from the test runs
	// of this canary. Artifacts include the log file, screenshots, and HAR files. The
	// name of the Amazon S3 bucket can't include a period (.).
	ArtifactS3Location *string

	// A structure that includes the entry point from which the canary should start
	// running your script. If the script is stored in an Amazon S3 bucket, the bucket
	// name, key, and version are also included.
	Code *types.CanaryCodeInput

	// Update the existing canary using the updated configurations from the DryRun
	// associated with the DryRunId.
	//
	// When you use the dryRunId field when updating a canary, the only other field
	// you can provide is the Schedule . Adding any other field will thrown an
	// exception.
	DryRunId *string

	// The ARN of the IAM role to be used to run the canary. This role must already
	// exist, and must include lambda.amazonaws.com as a principal in the trust
	// policy. The role must also have the following permissions:
	//
	//   - s3:PutObject
	//
	//   - s3:GetBucketLocation
	//
	//   - s3:ListAllMyBuckets
	//
	//   - cloudwatch:PutMetricData
	//
	//   - logs:CreateLogGroup
	//
	//   - logs:CreateLogStream
	//
	//   - logs:CreateLogStream
	ExecutionRoleArn *string

	// The number of days to retain data about failed runs of this canary.
	//
	// This setting affects the range of information returned by [GetCanaryRuns], as well as the
	// range of information displayed in the Synthetics console.
	//
	// [GetCanaryRuns]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanaryRuns.html
	FailureRetentionPeriodInDays *int32

	// Specifies whether to also delete the Lambda functions and layers used by this
	// canary when the canary is deleted.
	//
	// If the value of this parameter is OFF , then the value of the DeleteLambda
	// parameter of the [DeleteCanary]operation determines whether the Lambda functions and layers
	// will be deleted.
	//
	// [DeleteCanary]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DeleteCanary.html
	ProvisionedResourceCleanup types.ProvisionedResourceCleanupSetting

	// A structure that contains the timeout value that is used for each individual
	// run of the canary.
	//
	// Environment variable keys and values are encrypted at rest using Amazon Web
	// Services owned KMS keys. However, the environment variables are not encrypted on
	// the client side. Do not store sensitive information in them.
	RunConfig *types.CanaryRunConfigInput

	// Specifies the runtime version to use for the canary. For a list of valid
	// runtime versions and for more information about runtime versions, see [Canary Runtime Versions].
	//
	// [Canary Runtime Versions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html
	RuntimeVersion *string

	// A structure that contains information about how often the canary is to run, and
	// when these runs are to stop.
	Schedule *types.CanaryScheduleInput

	// The number of days to retain data about successful runs of this canary.
	//
	// This setting affects the range of information returned by [GetCanaryRuns], as well as the
	// range of information displayed in the Synthetics console.
	//
	// [GetCanaryRuns]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanaryRuns.html
	SuccessRetentionPeriodInDays *int32

	// Defines the screenshots to use as the baseline for comparisons during visual
	// monitoring comparisons during future runs of this canary. If you omit this
	// parameter, no changes are made to any baseline screenshots that the canary might
	// be using already.
	//
	// Visual monitoring is supported only on canaries running the
	// syn-puppeteer-node-3.2 runtime or later. For more information, see [Visual monitoring]and [Visual monitoring blueprint]
	//
	// [Visual monitoring]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html
	// [Visual monitoring blueprint]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html
	VisualReference *types.VisualReferenceInput

	// If this canary is to test an endpoint in a VPC, this structure contains
	// information about the subnet and security groups of the VPC endpoint. For more
	// information, see [Running a Canary in a VPC].
	//
	// [Running a Canary in a VPC]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html
	VpcConfig *types.VpcConfigInput

	noSmithyDocumentSerde
}

type UpdateCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCanaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCanary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCanary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCanary"); err != nil {
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
	if err = addOpUpdateCanaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCanary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCanary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCanary",
	}
}
