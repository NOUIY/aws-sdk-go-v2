// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Autopilot job also referred to as Autopilot experiment or AutoML job
// V2.
//
// An AutoML job in SageMaker AI is a fully automated process that allows you to
// build machine learning models with minimal effort and machine learning
// expertise. When initiating an AutoML job, you provide your data and optionally
// specify parameters tailored to your use case. SageMaker AI then automates the
// entire model development lifecycle, including data preprocessing, model
// training, tuning, and evaluation. AutoML jobs are designed to simplify and
// accelerate the model building process by automating various tasks and exploring
// different combinations of machine learning algorithms, data preprocessing
// techniques, and hyperparameter values. The output of an AutoML job comprises one
// or more trained models ready for deployment and inference. Additionally,
// SageMaker AI AutoML jobs generate a candidate model leaderboard, allowing you to
// select the best-performing model for deployment.
//
// For more information about AutoML jobs, see [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html] in the SageMaker AI developer
// guide.
//
// AutoML jobs V2 support various problem types such as regression, binary, and
// multiclass classification with tabular data, text and image classification,
// time-series forecasting, and fine-tuning of large language models (LLMs) for
// text generation.
//
// [CreateAutoMLJobV2]and [DescribeAutoMLJobV2] are new versions of [CreateAutoMLJob] and [DescribeAutoMLJob] which offer backward compatibility.
//
// CreateAutoMLJobV2 can manage tabular problem types identical to those of its
// previous version CreateAutoMLJob , as well as time-series forecasting,
// non-tabular problem types such as image or text classification, and text
// generation (LLMs fine-tuning).
//
// Find guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in [Migrate a CreateAutoMLJob to CreateAutoMLJobV2].
//
// For the list of available problem types supported by CreateAutoMLJobV2 , see [AutoMLProblemTypeConfig].
//
// You can find the best-performing model after you run an AutoML job V2 by
// calling [DescribeAutoMLJobV2].
//
// [CreateAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html
// [DescribeAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html
// [DescribeAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html
// [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [Migrate a CreateAutoMLJob to CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment.html#autopilot-create-experiment-api-migrate-v1-v2
// [AutoMLProblemTypeConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLProblemTypeConfig.html
func (c *Client) CreateAutoMLJobV2(ctx context.Context, params *CreateAutoMLJobV2Input, optFns ...func(*Options)) (*CreateAutoMLJobV2Output, error) {
	if params == nil {
		params = &CreateAutoMLJobV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoMLJobV2", params, optFns, c.addOperationCreateAutoMLJobV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoMLJobV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoMLJobV2Input struct {

	// An array of channel objects describing the input data and their location. Each
	// channel is a named input source. Similar to the [InputDataConfig]attribute in the CreateAutoMLJob
	// input parameters. The supported formats depend on the problem type:
	//
	//   - For tabular problem types: S3Prefix , ManifestFile .
	//
	//   - For image classification: S3Prefix , ManifestFile , AugmentedManifestFile .
	//
	//   - For text classification: S3Prefix .
	//
	//   - For time-series forecasting: S3Prefix .
	//
	//   - For text generation (LLMs fine-tuning): S3Prefix .
	//
	// [InputDataConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html#sagemaker-CreateAutoMLJob-request-InputDataConfig
	//
	// This member is required.
	AutoMLJobInputDataConfig []types.AutoMLJobChannel

	// Identifies an Autopilot job. The name must be unique to your account and is
	// case insensitive.
	//
	// This member is required.
	AutoMLJobName *string

	// Defines the configuration settings of one of the supported problem types.
	//
	// This member is required.
	AutoMLProblemTypeConfig types.AutoMLProblemTypeConfig

	// Provides information about encryption and the Amazon S3 output path needed to
	// store artifacts from an AutoML job.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the role that is used to access the data.
	//
	// This member is required.
	RoleArn *string

	// Specifies the compute configuration for the AutoML job V2.
	AutoMLComputeConfig *types.AutoMLComputeConfig

	// Specifies a metric to minimize or maximize as the objective of a job. If not
	// specified, the default objective metric depends on the problem type. For the
	// list of default values per problem type, see [AutoMLJobObjective].
	//
	//   - For tabular problem types: You must either provide both the
	//   AutoMLJobObjective and indicate the type of supervised learning problem in
	//   AutoMLProblemTypeConfig ( TabularJobConfig.ProblemType ), or none at all.
	//
	//   - For text generation problem types (LLMs fine-tuning): Fine-tuning language
	//   models in Autopilot does not require setting the AutoMLJobObjective field.
	//   Autopilot fine-tunes LLMs without requiring multiple candidates to be trained
	//   and evaluated. Instead, using your dataset, Autopilot directly fine-tunes your
	//   target model to enhance a default objective metric, the cross-entropy loss.
	//   After fine-tuning a language model, you can evaluate the quality of its
	//   generated text using different metrics. For a list of the available metrics, see
	//   [Metrics for fine-tuning LLMs in Autopilot].
	//
	// [Metrics for fine-tuning LLMs in Autopilot]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-llms-finetuning-metrics.html
	// [AutoMLJobObjective]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLJobObjective.html
	AutoMLJobObjective *types.AutoMLJobObjective

	// This structure specifies how to split the data into train and validation
	// datasets.
	//
	// The validation and training datasets must contain the same headers. For jobs
	// created by calling CreateAutoMLJob , the validation dataset must be less than 2
	// GB in size.
	//
	// This attribute must not be set for the time-series forecasting problem type, as
	// Autopilot automatically splits the input dataset into training and validation
	// sets.
	DataSplitConfig *types.AutoMLDataSplitConfig

	// Specifies how to generate the endpoint name for an automatic one-click
	// Autopilot model deployment.
	ModelDeployConfig *types.ModelDeployConfig

	// The security configuration for traffic encryption or Amazon VPC settings.
	SecurityConfig *types.AutoMLSecurityConfig

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, such as by purpose, owner, or environment.
	// For more information, see [Tagging Amazon Web ServicesResources]. Tag keys must be unique per resource.
	//
	// [Tagging Amazon Web ServicesResources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAutoMLJobV2Output struct {

	// The unique ARN assigned to the AutoMLJob when it is created.
	//
	// This member is required.
	AutoMLJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoMLJobV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoMLJobV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoMLJobV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAutoMLJobV2"); err != nil {
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
	if err = addOpCreateAutoMLJobV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoMLJobV2(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoMLJobV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAutoMLJobV2",
	}
}
