// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an eventbridge rule template to monitor events and send notifications
// to your targeted resources.
func (c *Client) CreateEventBridgeRuleTemplate(ctx context.Context, params *CreateEventBridgeRuleTemplateInput, optFns ...func(*Options)) (*CreateEventBridgeRuleTemplateOutput, error) {
	if params == nil {
		params = &CreateEventBridgeRuleTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventBridgeRuleTemplate", params, optFns, c.addOperationCreateEventBridgeRuleTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventBridgeRuleTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for CreateEventBridgeRuleTemplateRequest
type CreateEventBridgeRuleTemplateInput struct {

	// The type of event to match with the rule.
	//
	// This member is required.
	EventType types.EventBridgeRuleTemplateEventType

	// An eventbridge rule template group's identifier. Can be either be its id or
	// current name.
	//
	// This member is required.
	GroupIdentifier *string

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	//
	// This member is required.
	Name *string

	// A resource's optional description.
	Description *string

	// Placeholder documentation for __listOfEventBridgeRuleTemplateTarget
	EventTargets []types.EventBridgeRuleTemplateTarget

	// An ID that you assign to a create request. This ID ensures idempotency when
	// creating resources.
	RequestId *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Placeholder documentation for CreateEventBridgeRuleTemplateResponse
type CreateEventBridgeRuleTemplateOutput struct {

	// An eventbridge rule template's ARN (Amazon Resource Name)
	Arn *string

	// Placeholder documentation for __timestampIso8601
	CreatedAt *time.Time

	// A resource's optional description.
	Description *string

	// Placeholder documentation for __listOfEventBridgeRuleTemplateTarget
	EventTargets []types.EventBridgeRuleTemplateTarget

	// The type of event to match with the rule.
	EventType types.EventBridgeRuleTemplateEventType

	// An eventbridge rule template group's id. AWS provided template groups have ids
	// that start with aws-
	GroupId *string

	// An eventbridge rule template's id. AWS provided templates have ids that start
	// with aws-
	Id *string

	// Placeholder documentation for __timestampIso8601
	ModifiedAt *time.Time

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	Name *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventBridgeRuleTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEventBridgeRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEventBridgeRuleTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEventBridgeRuleTemplate"); err != nil {
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
	if err = addIdempotencyToken_opCreateEventBridgeRuleTemplateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEventBridgeRuleTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventBridgeRuleTemplate(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEventBridgeRuleTemplate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEventBridgeRuleTemplate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEventBridgeRuleTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEventBridgeRuleTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEventBridgeRuleTemplateInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEventBridgeRuleTemplateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEventBridgeRuleTemplate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEventBridgeRuleTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEventBridgeRuleTemplate",
	}
}
