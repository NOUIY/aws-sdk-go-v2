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

// Create a machine learning algorithm that you can use in SageMaker and list in
// the Amazon Web Services Marketplace.
func (c *Client) CreateAlgorithm(ctx context.Context, params *CreateAlgorithmInput, optFns ...func(*Options)) (*CreateAlgorithmOutput, error) {
	if params == nil {
		params = &CreateAlgorithmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlgorithm", params, optFns, c.addOperationCreateAlgorithmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAlgorithmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAlgorithmInput struct {

	// The name of the algorithm.
	//
	// This member is required.
	AlgorithmName *string

	// Specifies details about training jobs run by this algorithm, including the
	// following:
	//
	//   - The Amazon ECR path of the container and the version digest of the
	//   algorithm.
	//
	//   - The hyperparameters that the algorithm supports.
	//
	//   - The instance types that the algorithm supports for training.
	//
	//   - Whether the algorithm supports distributed training.
	//
	//   - The metrics that the algorithm emits to Amazon CloudWatch.
	//
	//   - Which metrics that the algorithm emits can be used as the objective metric
	//   for hyperparameter tuning jobs.
	//
	//   - The input channels that the algorithm supports for training data. For
	//   example, an algorithm might support train , validation , and test channels.
	//
	// This member is required.
	TrainingSpecification *types.TrainingSpecification

	// A description of the algorithm.
	AlgorithmDescription *string

	// Whether to certify the algorithm so that it can be listed in Amazon Web
	// Services Marketplace.
	CertifyForMarketplace *bool

	// Specifies details about inference jobs that the algorithm runs, including the
	// following:
	//
	//   - The Amazon ECR paths of containers that contain the inference code and
	//   model artifacts.
	//
	//   - The instance types that the algorithm supports for transform jobs and
	//   real-time endpoints used for inference.
	//
	//   - The input and output content formats that the algorithm supports for
	//   inference.
	InferenceSpecification *types.InferenceSpecification

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see [Tagging Amazon Web Services Resources].
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	// Specifies configurations for one or more training jobs and that SageMaker runs
	// to test the algorithm's training code and, optionally, one or more batch
	// transform jobs that SageMaker runs to test the algorithm's inference code.
	ValidationSpecification *types.AlgorithmValidationSpecification

	noSmithyDocumentSerde
}

type CreateAlgorithmOutput struct {

	// The Amazon Resource Name (ARN) of the new algorithm.
	//
	// This member is required.
	AlgorithmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAlgorithmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAlgorithm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAlgorithm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAlgorithm"); err != nil {
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
	if err = addOpCreateAlgorithmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlgorithm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAlgorithm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAlgorithm",
	}
}
