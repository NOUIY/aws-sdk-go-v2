// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a knowledge base.
//
// When using this API, you cannot reuse [Amazon AppIntegrations] DataIntegrations with external knowledge
// bases such as Salesforce and ServiceNow. If you do, you'll get an
// InvalidRequestException error.
//
// For example, you're programmatically managing your external knowledge base, and
// you want to add or remove one of the fields that is being ingested from
// Salesforce. Do the following:
//
//   - Call [DeleteKnowledgeBase].
//
//   - Call [DeleteDataIntegration].
//
//   - Call [CreateDataIntegration]to recreate the DataIntegration or a create different one.
//
//   - Call CreateKnowledgeBase.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteKnowledgeBase]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_DeleteKnowledgeBase.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func (c *Client) CreateKnowledgeBase(ctx context.Context, params *CreateKnowledgeBaseInput, optFns ...func(*Options)) (*CreateKnowledgeBaseOutput, error) {
	if params == nil {
		params = &CreateKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKnowledgeBase", params, optFns, c.addOperationCreateKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKnowledgeBaseInput struct {

	// The type of knowledge base. Only CUSTOM knowledge bases allow you to upload
	// your own content. EXTERNAL knowledge bases support integrations with third-party
	// systems whose content is synchronized automatically.
	//
	// This member is required.
	KnowledgeBaseType types.KnowledgeBaseType

	// The name of the knowledge base.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: http://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// The description.
	Description *string

	// Information about how to render the content.
	RenderingConfiguration *types.RenderingConfiguration

	// The configuration information for the customer managed key used for encryption.
	//
	// This KMS key must have a policy that allows kms:CreateGrant , kms:DescribeKey ,
	// kms:Decrypt , and kms:GenerateDataKey* permissions to the IAM identity using
	// the key to invoke Amazon Q in Connect.
	//
	// For more information about setting up a customer managed key for Amazon Q in
	// Connect, see [Enable Amazon Q in Connect for your instance].
	//
	// [Enable Amazon Q in Connect for your instance]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-q.html
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The source of the knowledge base content. Only set this argument for EXTERNAL
	// or Managed knowledge bases.
	SourceConfiguration types.SourceConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Contains details about how to ingest the documents in a data source.
	VectorIngestionConfiguration *types.VectorIngestionConfiguration

	noSmithyDocumentSerde
}

type CreateKnowledgeBaseOutput struct {

	// The knowledge base.
	KnowledgeBase *types.KnowledgeBaseData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKnowledgeBase"); err != nil {
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
	if err = addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKnowledgeBase(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKnowledgeBase struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKnowledgeBase) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKnowledgeBase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKnowledgeBaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKnowledgeBaseInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKnowledgeBase{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKnowledgeBase",
	}
}
