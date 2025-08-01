// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration for an application to use grants. Conceptually grants
// are authorization to request actions related to tokens. This configuration will
// be used when parties are requesting and receiving tokens during the trusted
// identity propagation process. For more information on the IAM Identity Center
// supported grant workflows, see [SAML 2.0 and OAuth 2.0].
//
// A grant is created between your applications and Identity Center instance which
// enables an application to use specified mechanisms to obtain tokens. These
// tokens are used by your applications to gain access to Amazon Web Services
// resources on behalf of users. The following elements are within these exchanges:
//
//   - Requester - The application requesting access to Amazon Web Services
//     resources.
//
//   - Subject - Typically the user that is requesting access to Amazon Web
//     Services resources.
//
//   - Grant - Conceptually, a grant is authorization to access Amazon Web
//     Services resources. These grants authorize token generation for authenticating
//     access to the requester and for the request to make requests on behalf of the
//     subjects. There are four types of grants:
//
//   - AuthorizationCode - Allows an application to request authorization through
//     a series of user-agent redirects.
//
//   - JWT bearer - Authorizes an application to exchange a JSON Web Token that
//     came from an external identity provider. To learn more, see [RFC 6479].
//
//   - Refresh token - Enables application to request new access tokens to replace
//     expiring or expired access tokens.
//
//   - Exchange token - A grant that requests tokens from the authorization server
//     by providing a ‘subject’ token with access scope authorizing trusted identity
//     propagation to this application. To learn more, see [RFC 8693].
//
//   - Authorization server - IAM Identity Center requests tokens.
//
// User credentials are never shared directly within these exchanges. Instead,
// applications use grants to request access tokens from IAM Identity Center. For
// more information, see [RFC 6479].
//
// Use cases
//
//   - Connecting to custom applications.
//
//   - Configuring an Amazon Web Services service to make calls to another Amazon
//     Web Services services using JWT tokens.
//
// [RFC 6479]: https://datatracker.ietf.org/doc/html/rfc6749
// [SAML 2.0 and OAuth 2.0]: https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps-saml2-oauth2.html
// [RFC 8693]: https://datatracker.ietf.org/doc/html/rfc8693
func (c *Client) PutApplicationGrant(ctx context.Context, params *PutApplicationGrantInput, optFns ...func(*Options)) (*PutApplicationGrantOutput, error) {
	if params == nil {
		params = &PutApplicationGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutApplicationGrant", params, optFns, c.addOperationPutApplicationGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutApplicationGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutApplicationGrantInput struct {

	// Specifies the ARN of the application to update.
	//
	// This member is required.
	ApplicationArn *string

	// Specifies a structure that describes the grant to update.
	//
	// This member is required.
	Grant types.Grant

	// Specifies the type of grant to update.
	//
	// This member is required.
	GrantType types.GrantType

	noSmithyDocumentSerde
}

type PutApplicationGrantOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutApplicationGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutApplicationGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutApplicationGrant{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutApplicationGrant"); err != nil {
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
	if err = addOpPutApplicationGrantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutApplicationGrant(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutApplicationGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutApplicationGrant",
	}
}
