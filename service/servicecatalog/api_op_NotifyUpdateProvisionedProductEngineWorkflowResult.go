// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies the result of the update engine execution.
func (c *Client) NotifyUpdateProvisionedProductEngineWorkflowResult(ctx context.Context, params *NotifyUpdateProvisionedProductEngineWorkflowResultInput, optFns ...func(*Options)) (*NotifyUpdateProvisionedProductEngineWorkflowResultOutput, error) {
	if params == nil {
		params = &NotifyUpdateProvisionedProductEngineWorkflowResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyUpdateProvisionedProductEngineWorkflowResult", params, optFns, c.addOperationNotifyUpdateProvisionedProductEngineWorkflowResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyUpdateProvisionedProductEngineWorkflowResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyUpdateProvisionedProductEngineWorkflowResultInput struct {

	//  The idempotency token that identifies the update engine execution.
	//
	// This member is required.
	IdempotencyToken *string

	//  The identifier of the record.
	//
	// This member is required.
	RecordId *string

	//  The status of the update engine execution.
	//
	// This member is required.
	Status types.EngineWorkflowStatus

	//  The encrypted contents of the update engine execution payload that Service
	// Catalog sends after the Terraform product update workflow starts.
	//
	// This member is required.
	WorkflowToken *string

	//  The reason why the update engine execution failed.
	FailureReason *string

	//  The output of the update engine execution.
	Outputs []types.RecordOutput

	noSmithyDocumentSerde
}

type NotifyUpdateProvisionedProductEngineWorkflowResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyUpdateProvisionedProductEngineWorkflowResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpNotifyUpdateProvisionedProductEngineWorkflowResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpNotifyUpdateProvisionedProductEngineWorkflowResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "NotifyUpdateProvisionedProductEngineWorkflowResult"); err != nil {
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
	if err = addIdempotencyToken_opNotifyUpdateProvisionedProductEngineWorkflowResultMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpNotifyUpdateProvisionedProductEngineWorkflowResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyUpdateProvisionedProductEngineWorkflowResult(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpNotifyUpdateProvisionedProductEngineWorkflowResult struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpNotifyUpdateProvisionedProductEngineWorkflowResult) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpNotifyUpdateProvisionedProductEngineWorkflowResult) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*NotifyUpdateProvisionedProductEngineWorkflowResultInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *NotifyUpdateProvisionedProductEngineWorkflowResultInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opNotifyUpdateProvisionedProductEngineWorkflowResultMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpNotifyUpdateProvisionedProductEngineWorkflowResult{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opNotifyUpdateProvisionedProductEngineWorkflowResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NotifyUpdateProvisionedProductEngineWorkflowResult",
	}
}
