// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing rule groups namespace within a workspace. A rule groups
// namespace is associated with exactly one rules file. A workspace can have
// multiple rule groups namespaces.
//
// Use this operation only to update existing rule groups namespaces. To create a
// new rule groups namespace, use CreateRuleGroupsNamespace .
//
// You can't use this operation to add tags to an existing rule groups namespace.
// Instead, use TagResource .
func (c *Client) PutRuleGroupsNamespace(ctx context.Context, params *PutRuleGroupsNamespaceInput, optFns ...func(*Options)) (*PutRuleGroupsNamespaceOutput, error) {
	if params == nil {
		params = &PutRuleGroupsNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRuleGroupsNamespace", params, optFns, c.addOperationPutRuleGroupsNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRuleGroupsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutRuleGroupsNamespace operation.
type PutRuleGroupsNamespaceInput struct {

	// The new rules file to use in the namespace. A base64-encoded version of the
	// YAML rule groups file.
	//
	// For details about the rule groups namespace structure, see [RuleGroupsNamespaceData].
	//
	// [RuleGroupsNamespaceData]: https://docs.aws.amazon.com/prometheus/latest/APIReference/yaml-RuleGroupsNamespaceData.html
	//
	// This member is required.
	Data []byte

	// The name of the rule groups namespace that you are updating.
	//
	// This member is required.
	Name *string

	// The ID of the workspace where you are updating the rule groups namespace.
	//
	// This member is required.
	WorkspaceId *string

	// A unique identifier that you can provide to ensure the idempotency of the
	// request. Case-sensitive.
	ClientToken *string

	noSmithyDocumentSerde
}

// Represents the output of a PutRuleGroupsNamespace operation.
type PutRuleGroupsNamespaceOutput struct {

	// The ARN of the rule groups namespace.
	//
	// This member is required.
	Arn *string

	// The name of the rule groups namespace that was updated.
	//
	// This member is required.
	Name *string

	// A structure that includes the current status of the rule groups namespace.
	//
	// This member is required.
	Status *types.RuleGroupsNamespaceStatus

	// The list of tag keys and values that are associated with the namespace.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRuleGroupsNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRuleGroupsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRuleGroupsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRuleGroupsNamespace"); err != nil {
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
	if err = addIdempotencyToken_opPutRuleGroupsNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutRuleGroupsNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRuleGroupsNamespace(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPutRuleGroupsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutRuleGroupsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutRuleGroupsNamespaceInput ")
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
func addIdempotencyToken_opPutRuleGroupsNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutRuleGroupsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutRuleGroupsNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRuleGroupsNamespace",
	}
}
