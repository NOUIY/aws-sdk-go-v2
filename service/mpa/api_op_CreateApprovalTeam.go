// Code generated by smithy-go-codegen DO NOT EDIT.

package mpa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mpa/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new approval team. For more information, see [Approval team] in the Multi-party
// approval User Guide.
//
// [Approval team]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func (c *Client) CreateApprovalTeam(ctx context.Context, params *CreateApprovalTeamInput, optFns ...func(*Options)) (*CreateApprovalTeamOutput, error) {
	if params == nil {
		params = &CreateApprovalTeamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApprovalTeam", params, optFns, c.addOperationCreateApprovalTeamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApprovalTeamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApprovalTeamInput struct {

	// An ApprovalStrategy object. Contains details for how the team grants approval.
	//
	// This member is required.
	ApprovalStrategy types.ApprovalStrategy

	// An array of ApprovalTeamRequesterApprovers objects. Contains details for the
	// approvers in the team.
	//
	// This member is required.
	Approvers []types.ApprovalTeamRequestApprover

	// Description for the team.
	//
	// This member is required.
	Description *string

	// Name of the team.
	//
	// This member is required.
	Name *string

	// An array of PolicyReference objects. Contains a list of policies that define
	// the permissions for team resources.
	//
	// The protected operation for a service integration might require specific
	// permissions. For more information, see [How other services work with Multi-party approval]in the Multi-party approval User Guide.
	//
	// [How other services work with Multi-party approval]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-integrations.html
	//
	// This member is required.
	Policies []types.PolicyReference

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If not provided, the Amazon Web Services populates this field.
	//
	// What is idempotency?
	//
	// When you make a mutating API request, the request typically returns a result
	// before the operation's asynchronous workflows have completed. Operations might
	// also time out or encounter other server issues before they complete, even though
	// the request has already returned a result. This could make it difficult to
	// determine whether the request succeeded or not, and could lead to multiple
	// retries to ensure that the operation completes successfully. However, if the
	// original request and the subsequent retries are successful, the operation is
	// completed multiple times. This means that you might create more resources than
	// you intended.
	//
	// Idempotency ensures that an API request completes no more than one time. With
	// an idempotent request, if the original request completes successfully, any
	// subsequent retries complete successfully without performing any further actions.
	ClientToken *string

	// Tags you want to attach to the team.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateApprovalTeamOutput struct {

	// Amazon Resource Name (ARN) for the team that was created.
	Arn *string

	// Timestamp when the team was created.
	CreationTime *time.Time

	// Name of the team that was created.
	Name *string

	// Version ID for the team that was created. When a team is updated, the version
	// ID changes.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApprovalTeamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApprovalTeam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApprovalTeam{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApprovalTeam"); err != nil {
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
	if err = addIdempotencyToken_opCreateApprovalTeamMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateApprovalTeamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApprovalTeam(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateApprovalTeam struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateApprovalTeam) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateApprovalTeam) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateApprovalTeamInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateApprovalTeamInput ")
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
func addIdempotencyToken_opCreateApprovalTeamMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateApprovalTeam{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateApprovalTeam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApprovalTeam",
	}
}
