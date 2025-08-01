// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	This operation is only valid for legacy predictors created with
//
// CreatePredictor. If you are not using a legacy predictor, use DescribeAutoPredictor.
//
// Describes a predictor created using the CreatePredictor operation.
//
// In addition to listing the properties provided in the CreatePredictor request,
// this operation lists the following properties:
//
//   - DatasetImportJobArns - The dataset import jobs used to import training data.
//
//   - AutoMLAlgorithmArns - If AutoML is performed, the algorithms that were
//     evaluated.
//
//   - CreationTime
//
//   - LastModificationTime
//
//   - Status
//
//   - Message - If an error occurred, information about the error.
func (c *Client) DescribePredictor(ctx context.Context, params *DescribePredictorInput, optFns ...func(*Options)) (*DescribePredictorOutput, error) {
	if params == nil {
		params = &DescribePredictorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePredictor", params, optFns, c.addOperationDescribePredictorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePredictorInput struct {

	// The Amazon Resource Name (ARN) of the predictor that you want information about.
	//
	// This member is required.
	PredictorArn *string

	noSmithyDocumentSerde
}

type DescribePredictorOutput struct {

	// The Amazon Resource Name (ARN) of the algorithm used for model training.
	AlgorithmArn *string

	// When PerformAutoML is specified, the ARN of the chosen algorithm.
	AutoMLAlgorithmArns []string

	//  The LatencyOptimized AutoML override strategy is only available in private
	// beta. Contact Amazon Web Services Support or your account manager to learn more
	// about access privileges.
	//
	// The AutoML strategy used to train the predictor. Unless LatencyOptimized is
	// specified, the AutoML strategy optimizes predictor accuracy.
	//
	// This parameter is only valid for predictors trained using AutoML.
	AutoMLOverrideStrategy types.AutoMLOverrideStrategy

	// When the model training task was created.
	CreationTime *time.Time

	// An array of the ARNs of the dataset import jobs used to import training data
	// for the predictor.
	DatasetImportJobArns []string

	// An Key Management Service (KMS) key and the Identity and Access Management
	// (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *types.EncryptionConfig

	// The estimated time remaining in minutes for the predictor training job to
	// complete.
	EstimatedTimeRemainingInMinutes *int64

	// Used to override the default evaluation parameters of the specified algorithm.
	// Amazon Forecast evaluates a predictor by splitting a dataset into training data
	// and testing data. The evaluation parameters define how to perform the split and
	// the number of iterations.
	EvaluationParameters *types.EvaluationParameters

	// The featurization configuration.
	FeaturizationConfig *types.FeaturizationConfig

	// The number of time-steps of the forecast. The forecast horizon is also called
	// the prediction length.
	ForecastHorizon *int32

	// The forecast types used during predictor training. Default value is
	// ["0.1","0.5","0.9"]
	ForecastTypes []string

	// The hyperparameter override values for the algorithm.
	HPOConfig *types.HyperParameterTuningJobConfig

	// Describes the dataset group that contains the data to use to train the
	// predictor.
	InputDataConfig *types.InputDataConfig

	// Whether the predictor was created with CreateAutoPredictor.
	IsAutoPredictor *bool

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	//   - CREATE_PENDING - The CreationTime .
	//
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//
	//   - CREATE_STOPPING - The current timestamp.
	//
	//   - CREATE_STOPPED - When the job stopped.
	//
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The accuracy metric used to optimize the predictor.
	OptimizationMetric types.OptimizationMetric

	// Whether the predictor is set to perform AutoML.
	PerformAutoML *bool

	// Whether the predictor is set to perform hyperparameter optimization (HPO).
	PerformHPO *bool

	// The ARN of the predictor.
	PredictorArn *string

	// Details on the the status and results of the backtests performed to evaluate
	// the accuracy of the predictor. You specify the number of backtests to perform
	// when you call the operation.
	PredictorExecutionDetails *types.PredictorExecutionDetails

	// The name of the predictor.
	PredictorName *string

	// The status of the predictor. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//
	//   - CREATE_STOPPING , CREATE_STOPPED
	//
	// The Status of the predictor must be ACTIVE before you can use the predictor to
	// create a forecast.
	Status *string

	// The default training parameters or overrides selected during model training.
	// When running AutoML or choosing HPO with CNN-QR or DeepAR+, the optimized values
	// for the chosen hyperparameters are returned. For more information, see aws-forecast-choosing-recipes.
	TrainingParameters map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePredictorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePredictor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePredictor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePredictor"); err != nil {
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
	if err = addOpDescribePredictorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePredictor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePredictor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePredictor",
	}
}
