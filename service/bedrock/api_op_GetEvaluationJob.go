// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an evaluation job, such as the status of the job.
func (c *Client) GetEvaluationJob(ctx context.Context, params *GetEvaluationJobInput, optFns ...func(*Options)) (*GetEvaluationJobOutput, error) {
	if params == nil {
		params = &GetEvaluationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvaluationJob", params, optFns, c.addOperationGetEvaluationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvaluationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvaluationJobInput struct {

	// The Amazon Resource Name (ARN) of the evaluation job you want get information
	// on.
	//
	// This member is required.
	JobIdentifier *string

	noSmithyDocumentSerde
}

type GetEvaluationJobOutput struct {

	// The time the evaluation job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Contains the configuration details of either an automated or human-based
	// evaluation job.
	//
	// This member is required.
	EvaluationConfig types.EvaluationConfig

	// Contains the configuration details of the inference model used for the
	// evaluation job.
	//
	// This member is required.
	InferenceConfig types.EvaluationInferenceConfig

	// The Amazon Resource Name (ARN) of the evaluation job.
	//
	// This member is required.
	JobArn *string

	// The name for the evaluation job.
	//
	// This member is required.
	JobName *string

	// Specifies whether the evaluation job is automated or human-based.
	//
	// This member is required.
	JobType types.EvaluationJobType

	// Contains the configuration details of the Amazon S3 bucket for storing the
	// results of the evaluation job.
	//
	// This member is required.
	OutputDataConfig *types.EvaluationOutputDataConfig

	// The Amazon Resource Name (ARN) of the IAM service role used in the evaluation
	// job.
	//
	// This member is required.
	RoleArn *string

	// The current status of the evaluation job.
	//
	// This member is required.
	Status types.EvaluationJobStatus

	// Specifies whether the evaluation job is for evaluating a model or evaluating a
	// knowledge base (retrieval and response generation).
	ApplicationType types.ApplicationType

	// The Amazon Resource Name (ARN) of the customer managed encryption key specified
	// when the evaluation job was created.
	CustomerEncryptionKeyId *string

	// A list of strings that specify why the evaluation job failed to create.
	FailureMessages []string

	// The description of the evaluation job.
	JobDescription *string

	// The time the evaluation job was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEvaluationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEvaluationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEvaluationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEvaluationJob"); err != nil {
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
	if err = addOpGetEvaluationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvaluationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEvaluationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEvaluationJob",
	}
}
