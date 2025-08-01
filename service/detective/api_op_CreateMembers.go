// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// CreateMembers is used to send invitations to accounts. For the organization
// behavior graph, the Detective administrator account uses CreateMembers to
// enable organization accounts as member accounts.
//
// For invited accounts, CreateMembers sends a request to invite the specified
// Amazon Web Services accounts to be member accounts in the behavior graph. This
// operation can only be called by the administrator account for a behavior graph.
//
// CreateMembers verifies the accounts and then invites the verified accounts. The
// administrator can optionally specify to not send invitation emails to the member
// accounts. This would be used when the administrator manages their member
// accounts centrally.
//
// For organization accounts in the organization behavior graph, CreateMembers
// attempts to enable the accounts. The organization accounts do not receive
// invitations.
//
// The request provides the behavior graph ARN and the list of accounts to invite
// or to enable.
//
// The response separates the requested accounts into two lists:
//
//   - The accounts that CreateMembers was able to process. For invited accounts,
//     includes member accounts that are being verified, that have passed verification
//     and are to be invited, and that have failed verification. For organization
//     accounts in the organization behavior graph, includes accounts that can be
//     enabled and that cannot be enabled.
//
//   - The accounts that CreateMembers was unable to process. This list includes
//     accounts that were already invited to be member accounts in the behavior graph.
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	if params == nil {
		params = &CreateMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembers", params, optFns, c.addOperationCreateMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// The list of Amazon Web Services accounts to invite or to enable. You can invite
	// or enable up to 50 accounts at a time. For each invited account, the account
	// list contains the account identifier and the Amazon Web Services account root
	// user email address. For organization accounts in the organization behavior
	// graph, the email address is not required.
	//
	// This member is required.
	Accounts []types.Account

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	// if set to true , then the invited accounts do not receive email notifications.
	// By default, this is set to false , and the invited accounts receive email
	// notifications.
	//
	// Organization accounts in the organization behavior graph do not receive email
	// notifications.
	DisableEmailNotification bool

	// Customized message text to include in the invitation email message to the
	// invited member accounts.
	Message *string

	noSmithyDocumentSerde
}

type CreateMembersOutput struct {

	// The set of member account invitation or enablement requests that Detective was
	// able to process. This includes accounts that are being verified, that failed
	// verification, and that passed verification and are being sent an invitation or
	// are being enabled.
	Members []types.MemberDetail

	// The list of accounts for which Detective was unable to process the invitation
	// or enablement request. For each account, the list provides the reason why the
	// request could not be processed. The list includes accounts that are already
	// member accounts in the behavior graph.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMembers"); err != nil {
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
	if err = addOpCreateMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMembers",
	}
}
