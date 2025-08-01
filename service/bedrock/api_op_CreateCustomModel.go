// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom model in Amazon Bedrock. After the model is active, you
// can use it for inference.
//
// To use the model for inference, you must purchase Provisioned Throughput for
// it. You can't use On-demand inference with these custom models. For more
// information about Provisioned Throughput, see [Provisioned Throughput].
//
// The model appears in ListCustomModels with a customizationType of imported . To
// track the status of the new model, you use the GetCustomModel API operation.
// The model can be in the following states:
//
//   - Creating - Initial state during validation and registration
//
//   - Active - Model is ready for use in inference
//
//   - Failed - Creation process encountered an error
//
// # Related APIs
//
// [GetCustomModel]
//
// [ListCustomModels]
//
// [DeleteCustomModel]
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [ListCustomModels]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModels.html
// [DeleteCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModel.html
// [GetCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModel.html
func (c *Client) CreateCustomModel(ctx context.Context, params *CreateCustomModelInput, optFns ...func(*Options)) (*CreateCustomModelOutput, error) {
	if params == nil {
		params = &CreateCustomModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomModel", params, optFns, c.addOperationCreateCustomModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomModelInput struct {

	// A unique name for the custom model.
	//
	// This member is required.
	ModelName *string

	// The data source for the model. The Amazon S3 URI in the model source must be
	// for the Amazon-managed Amazon S3 bucket containing your model artifacts.
	//
	// This member is required.
	ModelSourceConfig types.ModelDataSource

	// A unique, case-sensitive identifier to ensure that the API request completes no
	// more than one time. If this token matches a previous request, Amazon Bedrock
	// ignores the request, but does not return an error. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the customer managed KMS key to encrypt the
	// custom model. If you don't provide a KMS key, Amazon Bedrock uses an Amazon Web
	// Services-managed KMS key to encrypt the model.
	//
	// If you provide a customer managed KMS key, your Amazon Bedrock service role
	// must have permissions to use it. For more information see [Encryption of imported models].
	//
	// [Encryption of imported models]: https://docs.aws.amazon.com/bedrock/latest/userguide/encryption-import-model.html
	ModelKmsKeyArn *string

	// A list of key-value pairs to associate with the custom model resource. You can
	// use these tags to organize and identify your resources.
	//
	// For more information, see [Tagging resources] in the [Amazon Bedrock User Guide].
	//
	// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
	// [Tagging resources]: https://docs.aws.amazon.com/bedrock/latest/userguide/tagging.html
	ModelTags []types.Tag

	// The Amazon Resource Name (ARN) of an IAM service role that Amazon Bedrock
	// assumes to perform tasks on your behalf. This role must have permissions to
	// access the Amazon S3 bucket containing your model artifacts and the KMS key (if
	// specified). For more information, see [Setting up an IAM service role for importing models]in the Amazon Bedrock User Guide.
	//
	// [Setting up an IAM service role for importing models]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-import-iam-role.html
	RoleArn *string

	noSmithyDocumentSerde
}

type CreateCustomModelOutput struct {

	// The Amazon Resource Name (ARN) of the new custom model.
	//
	// This member is required.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCustomModel"); err != nil {
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
	if err = addIdempotencyToken_opCreateCustomModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomModel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomModelInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCustomModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCustomModel",
	}
}
