// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a FirewallRuleGroup with a VPC, to provide DNS filtering for the VPC.
func (c *Client) AssociateFirewallRuleGroup(ctx context.Context, params *AssociateFirewallRuleGroupInput, optFns ...func(*Options)) (*AssociateFirewallRuleGroupOutput, error) {
	if params == nil {
		params = &AssociateFirewallRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateFirewallRuleGroup", params, optFns, c.addOperationAssociateFirewallRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateFirewallRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateFirewallRuleGroupInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// The unique identifier of the firewall rule group.
	//
	// This member is required.
	FirewallRuleGroupId *string

	// A name that lets you identify the association, to manage and use it.
	//
	// This member is required.
	Name *string

	// The setting that determines the processing order of the rule group among the
	// rule groups that you associate with the specified VPC. DNS Firewall filters VPC
	// traffic starting from the rule group with the lowest numeric priority setting.
	//
	// You must specify a unique priority for each rule group that you associate with
	// a single VPC. To make it easier to insert rule groups later, leave space between
	// the numbers, for example, use 101, 200, and so on. You can change the priority
	// setting for a rule group association after you create it.
	//
	// The allowed values for Priority are between 100 and 9900.
	//
	// This member is required.
	Priority *int32

	// The unique identifier of the VPC that you want to associate with the rule
	// group.
	//
	// This member is required.
	VpcId *string

	// If enabled, this setting disallows modification or removal of the association,
	// to help prevent against accidentally altering DNS firewall protections. When you
	// create the association, the default setting is DISABLED .
	MutationProtection types.MutationProtectionStatus

	// A list of the tag keys and values that you want to associate with the rule
	// group association.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type AssociateFirewallRuleGroupOutput struct {

	// The association that you just created. The association has an ID that you can
	// use to identify it in other requests, like update and delete.
	FirewallRuleGroupAssociation *types.FirewallRuleGroupAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateFirewallRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateFirewallRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateFirewallRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateFirewallRuleGroup"); err != nil {
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
	if err = addIdempotencyToken_opAssociateFirewallRuleGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateFirewallRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateFirewallRuleGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpAssociateFirewallRuleGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateFirewallRuleGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateFirewallRuleGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateFirewallRuleGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateFirewallRuleGroupInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAssociateFirewallRuleGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateFirewallRuleGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateFirewallRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateFirewallRuleGroup",
	}
}
