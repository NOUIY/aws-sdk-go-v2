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

// Updates the specified portfolio share. You can use this API to enable or
// disable TagOptions sharing or Principal sharing for an existing portfolio
// share.
//
// The portfolio share cannot be updated if the CreatePortfolioShare operation is
// IN_PROGRESS , as the share is not available to recipient entities. In this case,
// you must wait for the portfolio share to be completed.
//
// You must provide the accountId or organization node in the input, but not both.
//
// If the portfolio is shared to both an external account and an organization
// node, and both shares need to be updated, you must invoke UpdatePortfolioShare
// separately for each share type.
//
// This API cannot be used for removing the portfolio share. You must use
// DeletePortfolioShare API for that action.
//
// When you associate a principal with portfolio, a potential privilege escalation
// path may occur when that portfolio is then shared with other accounts. For a
// user in a recipient account who is not an Service Catalog Admin, but still has
// the ability to create Principals (Users/Groups/Roles), that user could create a
// role that matches a principal name association for the portfolio. Although this
// user may not know which principal names are associated through Service Catalog,
// they may be able to guess the user. If this potential escalation path is a
// concern, then Service Catalog recommends using PrincipalType as IAM . With this
// configuration, the PrincipalARN must already exist in the recipient account
// before it can be associated.
func (c *Client) UpdatePortfolioShare(ctx context.Context, params *UpdatePortfolioShareInput, optFns ...func(*Options)) (*UpdatePortfolioShareOutput, error) {
	if params == nil {
		params = &UpdatePortfolioShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePortfolioShare", params, optFns, c.addOperationUpdatePortfolioShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePortfolioShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePortfolioShareInput struct {

	// The unique identifier of the portfolio for which the share will be updated.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// The Amazon Web Services account Id of the recipient account. This field is
	// required when updating an external account to account type share.
	AccountId *string

	// Information about the organization node.
	OrganizationNode *types.OrganizationNode

	// A flag to enables or disables Principals sharing in the portfolio. If this
	// field is not provided, the current state of the Principals sharing on the
	// portfolio share will not be modified.
	SharePrincipals *bool

	// Enables or disables TagOptions sharing for the portfolio share. If this field
	// is not provided, the current state of TagOptions sharing on the portfolio share
	// will not be modified.
	ShareTagOptions *bool

	noSmithyDocumentSerde
}

type UpdatePortfolioShareOutput struct {

	// The token that tracks the status of the UpdatePortfolioShare operation for
	// external account to account or organizational type sharing.
	PortfolioShareToken *string

	// The status of UpdatePortfolioShare operation. You can also obtain the operation
	// status using DescribePortfolioShareStatus API.
	Status types.ShareStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePortfolioShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePortfolioShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePortfolioShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePortfolioShare"); err != nil {
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
	if err = addOpUpdatePortfolioShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePortfolioShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePortfolioShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePortfolioShare",
	}
}
