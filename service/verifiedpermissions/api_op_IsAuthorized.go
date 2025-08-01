// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Makes an authorization decision about a service request described in the
// parameters. The information in the parameters can also define additional context
// that Verified Permissions can include in the evaluation. The request is
// evaluated against all matching policies in the specified policy store. The
// result of the decision is either Allow or Deny , along with a list of the
// policies that resulted in the decision.
func (c *Client) IsAuthorized(ctx context.Context, params *IsAuthorizedInput, optFns ...func(*Options)) (*IsAuthorizedOutput, error) {
	if params == nil {
		params = &IsAuthorizedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IsAuthorized", params, optFns, c.addOperationIsAuthorizedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IsAuthorizedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IsAuthorizedInput struct {

	// Specifies the ID of the policy store. Policies in this policy store will be
	// used to make an authorization decision for the input.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the requested action to be authorized. For example, is the principal
	// authorized to perform this action on the resource?
	Action *types.ActionIdentifier

	// Specifies additional context that can be used to make more granular
	// authorization decisions.
	Context types.ContextDefinition

	// (Optional) Specifies the list of resources and principals and their associated
	// attributes that Verified Permissions can examine when evaluating the policies.
	// These additional entities and their attributes can be referenced and checked by
	// conditional elements in the policies in the specified policy store.
	//
	// You can include only principal and resource entities in this parameter; you
	// can't include actions. You must specify actions in the schema.
	Entities types.EntitiesDefinition

	// Specifies the principal for which the authorization decision is to be made.
	Principal *types.EntityIdentifier

	// Specifies the resource for which the authorization decision is to be made.
	Resource *types.EntityIdentifier

	noSmithyDocumentSerde
}

type IsAuthorizedOutput struct {

	// An authorization decision that indicates if the authorization request should be
	// allowed or denied.
	//
	// This member is required.
	Decision types.Decision

	// The list of determining policies used to make the authorization decision. For
	// example, if there are two matching policies, where one is a forbid and the other
	// is a permit, then the forbid policy will be the determining policy. In the case
	// of multiple matching permit policies then there would be multiple determining
	// policies. In the case that no policies match, and hence the response is DENY,
	// there would be no determining policies.
	//
	// This member is required.
	DeterminingPolicies []types.DeterminingPolicyItem

	// Errors that occurred while making an authorization decision, for example, a
	// policy references an Entity or entity Attribute that does not exist in the
	// slice.
	//
	// This member is required.
	Errors []types.EvaluationErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIsAuthorizedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpIsAuthorized{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpIsAuthorized{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "IsAuthorized"); err != nil {
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
	if err = addOpIsAuthorizedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIsAuthorized(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIsAuthorized(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IsAuthorized",
	}
}
