// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a rule in Amazon DataZone. A rule is a formal agreement that enforces
// specific requirements across user workflows (e.g., publishing assets to the
// catalog, requesting subscriptions, creating projects) within the Amazon DataZone
// data portal. These rules help maintain consistency, ensure compliance, and
// uphold governance standards in data management processes. For instance, a
// metadata enforcement rule can specify the required information for creating a
// subscription request or publishing a data asset to the catalog, ensuring
// alignment with organizational standards.
func (c *Client) CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) {
	if params == nil {
		params = &CreateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRule", params, optFns, c.addOperationCreateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleInput struct {

	// The action of the rule.
	//
	// This member is required.
	Action types.RuleAction

	// The detail of the rule.
	//
	// This member is required.
	Detail types.RuleDetail

	// The ID of the domain where the rule is created.
	//
	// This member is required.
	DomainIdentifier *string

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The scope of the rule.
	//
	// This member is required.
	Scope *types.RuleScope

	// The target of the rule.
	//
	// This member is required.
	Target types.RuleTarget

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The description of the rule.
	Description *string

	noSmithyDocumentSerde
}

type CreateRuleOutput struct {

	// The action of the rule.
	//
	// This member is required.
	Action types.RuleAction

	// The timestamp at which the rule is created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user who creates the rule.
	//
	// This member is required.
	CreatedBy *string

	// The detail of the rule.
	//
	// This member is required.
	Detail types.RuleDetail

	// The ID of the rule.
	//
	// This member is required.
	Identifier *string

	// The name of the rule.
	//
	// This member is required.
	Name *string

	// The type of the rule.
	//
	// This member is required.
	RuleType types.RuleType

	// The scope of the rule.
	//
	// This member is required.
	Scope *types.RuleScope

	// The target of the rule.
	//
	// This member is required.
	Target types.RuleTarget

	// The description of the rule.
	Description *string

	// The target type of the rule.
	TargetType types.RuleTargetType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRule"); err != nil {
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
	if err = addIdempotencyToken_opCreateRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRuleInput ")
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
func addIdempotencyToken_opCreateRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRule",
	}
}
