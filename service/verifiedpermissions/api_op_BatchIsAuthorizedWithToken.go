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

// Makes a series of decisions about multiple authorization requests for one
// token. The principal in this request comes from an external identity source in
// the form of an identity or access token, formatted as a [JSON web token (JWT)]. The information in
// the parameters can also define additional context that Verified Permissions can
// include in the evaluations.
//
// The request is evaluated against all policies in the specified policy store
// that match the entities that you provide in the entities declaration and in the
// token. The result of the decisions is a series of Allow or Deny responses,
// along with the IDs of the policies that produced each decision.
//
// The entities of a BatchIsAuthorizedWithToken API request can contain up to 100
// resources and up to 99 user groups. The requests of a BatchIsAuthorizedWithToken
// API request can contain up to 30 requests.
//
// The BatchIsAuthorizedWithToken operation doesn't have its own IAM permission.
// To authorize this operation for Amazon Web Services principals, include the
// permission verifiedpermissions:IsAuthorizedWithToken in their IAM policies.
//
// [JSON web token (JWT)]: https://wikipedia.org/wiki/JSON_Web_Token
func (c *Client) BatchIsAuthorizedWithToken(ctx context.Context, params *BatchIsAuthorizedWithTokenInput, optFns ...func(*Options)) (*BatchIsAuthorizedWithTokenOutput, error) {
	if params == nil {
		params = &BatchIsAuthorizedWithTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchIsAuthorizedWithToken", params, optFns, c.addOperationBatchIsAuthorizedWithTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchIsAuthorizedWithTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchIsAuthorizedWithTokenInput struct {

	// Specifies the ID of the policy store. Policies in this policy store will be
	// used to make an authorization decision for the input.
	//
	// This member is required.
	PolicyStoreId *string

	// An array of up to 30 requests that you want Verified Permissions to evaluate.
	//
	// This member is required.
	Requests []types.BatchIsAuthorizedWithTokenInputItem

	// Specifies an access token for the principal that you want to authorize in each
	// request. This token is provided to you by the identity provider (IdP) associated
	// with the specified identity source. You must specify either an accessToken , an
	// identityToken , or both.
	//
	// Must be an access token. Verified Permissions returns an error if the token_use
	// claim in the submitted token isn't access .
	AccessToken *string

	// (Optional) Specifies the list of resources and their associated attributes that
	// Verified Permissions can examine when evaluating the policies. These additional
	// entities and their attributes can be referenced and checked by conditional
	// elements in the policies in the specified policy store.
	//
	// You can't include principals in this parameter, only resource and action
	// entities. This parameter can't include any entities of a type that matches the
	// user or group entity types that you defined in your identity source.
	//
	//   - The BatchIsAuthorizedWithToken operation takes principal attributes from
	//   only the identityToken or accessToken passed to the operation.
	//
	//   - For action entities, you can include only their Identifier and EntityType .
	Entities types.EntitiesDefinition

	// Specifies an identity (ID) token for the principal that you want to authorize
	// in each request. This token is provided to you by the identity provider (IdP)
	// associated with the specified identity source. You must specify either an
	// accessToken , an identityToken , or both.
	//
	// Must be an ID token. Verified Permissions returns an error if the token_use
	// claim in the submitted token isn't id .
	IdentityToken *string

	noSmithyDocumentSerde
}

type BatchIsAuthorizedWithTokenOutput struct {

	// A series of Allow or Deny decisions for each request, and the policies that
	// produced them. These results are returned in the order they were requested.
	//
	// This member is required.
	Results []types.BatchIsAuthorizedWithTokenOutputItem

	// The identifier of the principal in the ID or access token.
	Principal *types.EntityIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchIsAuthorizedWithTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchIsAuthorizedWithToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchIsAuthorizedWithToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchIsAuthorizedWithToken"); err != nil {
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
	if err = addOpBatchIsAuthorizedWithTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchIsAuthorizedWithToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchIsAuthorizedWithToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchIsAuthorizedWithToken",
	}
}
