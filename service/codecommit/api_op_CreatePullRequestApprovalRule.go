// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an approval rule for a pull request.
func (c *Client) CreatePullRequestApprovalRule(ctx context.Context, params *CreatePullRequestApprovalRuleInput, optFns ...func(*Options)) (*CreatePullRequestApprovalRuleOutput, error) {
	if params == nil {
		params = &CreatePullRequestApprovalRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePullRequestApprovalRule", params, optFns, c.addOperationCreatePullRequestApprovalRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePullRequestApprovalRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePullRequestApprovalRuleInput struct {

	// The content of the approval rule, including the number of approvals needed and
	// the structure of an approval pool defined for approvals, if any. For more
	// information about approval pools, see the CodeCommit User Guide.
	//
	// When you create the content of the approval rule, you can specify approvers in
	// an approval pool in one of two ways:
	//
	//   - CodeCommitApprovers: This option only requires an Amazon Web Services
	//   account and a resource. It can be used for both IAM users and federated access
	//   users whose name matches the provided resource name. This is a very powerful
	//   option that offers a great deal of flexibility. For example, if you specify the
	//   Amazon Web Services account 123456789012 and Mary_Major, all of the following
	//   would be counted as approvals coming from that user:
	//
	//   - An IAM user in the account (arn:aws:iam::123456789012:user/Mary_Major)
	//
	//   - A federated user identified in IAM as Mary_Major
	//   (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//
	// This option does not recognize an active session of someone assuming the role
	//   of CodeCommitReview with a role session name of Mary_Major
	//   (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major) unless you
	//   include a wildcard (*Mary_Major).
	//
	//   - Fully qualified ARN: This option allows you to specify the fully qualified
	//   Amazon Resource Name (ARN) of the IAM user or role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see [IAM Identifiers] in the IAM
	// User Guide.
	//
	// [IAM Identifiers]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html
	//
	// This member is required.
	ApprovalRuleContent *string

	// The name for the approval rule.
	//
	// This member is required.
	ApprovalRuleName *string

	// The system-generated ID of the pull request for which you want to create the
	// approval rule.
	//
	// This member is required.
	PullRequestId *string

	noSmithyDocumentSerde
}

type CreatePullRequestApprovalRuleOutput struct {

	// Information about the created approval rule.
	//
	// This member is required.
	ApprovalRule *types.ApprovalRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePullRequestApprovalRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePullRequestApprovalRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePullRequestApprovalRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePullRequestApprovalRule"); err != nil {
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
	if err = addOpCreatePullRequestApprovalRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePullRequestApprovalRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePullRequestApprovalRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePullRequestApprovalRule",
	}
}
