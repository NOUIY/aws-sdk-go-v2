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

// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose.
//
// A hyperparameter tuning job automatically creates Amazon SageMaker experiments,
// trials, and trial components for each training job that it runs. You can view
// these entities in Amazon SageMaker Studio. For more information, see [View Experiments, Trials, and Trial Components].
//
// Do not include any security-sensitive information including account access IDs,
// secrets, or tokens in any hyperparameter fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by any
// security-sensitive information included in the request hyperparameter variable
// or plain text fields..
//
// [View Experiments, Trials, and Trial Components]: https://docs.aws.amazon.com/sagemaker/latest/dg/experiments-view-compare.html#experiments-view
func (c *Client) CreateHyperParameterTuningJob(ctx context.Context, params *CreateHyperParameterTuningJobInput, optFns ...func(*Options)) (*CreateHyperParameterTuningJobOutput, error) {
	if params == nil {
		params = &CreateHyperParameterTuningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHyperParameterTuningJob", params, optFns, c.addOperationCreateHyperParameterTuningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHyperParameterTuningJobInput struct {

	// The [HyperParameterTuningJobConfig] object that describes the tuning job, including the search strategy, the
	// objective metric used to evaluate training jobs, ranges of parameters to search,
	// and resource limits for the tuning job. For more information, see [How Hyperparameter Tuning Works].
	//
	// [How Hyperparameter Tuning Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html
	// [HyperParameterTuningJobConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html
	//
	// This member is required.
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig

	// The name of the tuning job. This name is the prefix for the names of all
	// training jobs that this tuning job launches. The name must be unique within the
	// same Amazon Web Services account and Amazon Web Services Region. The name must
	// have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % -
	// (hyphen). The name is not case sensitive.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// Configures SageMaker Automatic model tuning (AMT) to automatically find optimal
	// parameters for the following fields:
	//
	// [ParameterRanges]
	//   - : The names and ranges of parameters that a hyperparameter tuning job can
	//   optimize.
	//
	// [ResourceLimits]
	//   - : The maximum resources that can be used for a training job. These
	//   resources include the maximum number of training jobs, the maximum runtime of a
	//   tuning job, and the maximum number of training jobs to run at the same time.
	//
	// [TrainingJobEarlyStoppingType]
	//   - : A flag that specifies whether or not to use early stopping for training
	//   jobs launched by a hyperparameter tuning job.
	//
	// [RetryStrategy]
	//   - : The number of times to retry a training job.
	//
	// [Strategy]
	//   - : Specifies how hyperparameter tuning chooses the combinations of
	//   hyperparameter values to use for the training jobs that it launches.
	//
	// [ConvergenceDetected]
	//   - : A flag to indicate that Automatic model tuning (AMT) has detected model
	//   convergence.
	//
	// [TrainingJobEarlyStoppingType]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-TrainingJobEarlyStoppingType
	// [ConvergenceDetected]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ConvergenceDetected.html
	// [ResourceLimits]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceLimits.html
	// [Strategy]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html
	// [ParameterRanges]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-ParameterRanges
	// [RetryStrategy]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html#sagemaker-Type-HyperParameterTrainingJobDefinition-RetryStrategy
	Autotune *types.Autotune

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see [Tagging Amazon Web Services Resources].
	//
	// Tags that you specify for the tuning job are also added to all training jobs
	// that the tuning job launches.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	// The [HyperParameterTrainingJobDefinition] object that describes the training jobs that this tuning job launches,
	// including static hyperparameters, input data configuration, output data
	// configuration, resource configuration, and stopping condition.
	//
	// [HyperParameterTrainingJobDefinition]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition

	// A list of the [HyperParameterTrainingJobDefinition] objects launched for this tuning job.
	//
	// [HyperParameterTrainingJobDefinition]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html
	TrainingJobDefinitions []types.HyperParameterTrainingJobDefinition

	// Specifies the configuration for starting the hyperparameter tuning job using
	// one or more previous tuning jobs as a starting point. The results of previous
	// tuning jobs are used to inform which combinations of hyperparameters to search
	// over in the new tuning job.
	//
	// All training jobs launched by the new hyperparameter tuning job are evaluated
	// by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM as
	// the WarmStartType value for the warm start configuration, the training job that
	// performs the best in the new tuning job is compared to the best training jobs
	// from the parent tuning jobs. From these, the training job that performs the best
	// as measured by the objective metric is returned as the overall best training
	// job.
	//
	// All training jobs launched by parent hyperparameter tuning jobs and the new
	// hyperparameter tuning jobs count against the limit of training jobs for the
	// tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig

	noSmithyDocumentSerde
}

type CreateHyperParameterTuningJobOutput struct {

	// The Amazon Resource Name (ARN) of the tuning job. SageMaker assigns an ARN to a
	// hyperparameter tuning job when you create it.
	//
	// This member is required.
	HyperParameterTuningJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHyperParameterTuningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHyperParameterTuningJob"); err != nil {
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
	if err = addOpCreateHyperParameterTuningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHyperParameterTuningJob",
	}
}
