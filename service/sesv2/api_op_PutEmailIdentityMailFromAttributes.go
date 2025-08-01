// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to enable or disable the custom Mail-From domain configuration for an
// email identity.
func (c *Client) PutEmailIdentityMailFromAttributes(ctx context.Context, params *PutEmailIdentityMailFromAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityMailFromAttributesOutput, error) {
	if params == nil {
		params = &PutEmailIdentityMailFromAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailIdentityMailFromAttributes", params, optFns, c.addOperationPutEmailIdentityMailFromAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailIdentityMailFromAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to configure the custom MAIL FROM domain for a verified identity.
type PutEmailIdentityMailFromAttributesInput struct {

	// The verified email identity.
	//
	// This member is required.
	EmailIdentity *string

	// The action to take if the required MX record isn't found when you send an
	// email. When you set this value to UseDefaultValue , the mail is sent using
	// amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage
	// , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't
	// attempt to deliver the email.
	//
	// These behaviors are taken when the custom MAIL FROM domain configuration is in
	// the Pending , Failed , and TemporaryFailure states.
	BehaviorOnMxFailure types.BehaviorOnMxFailure

	//  The custom MAIL FROM domain that you want the verified identity to use. The
	// MAIL FROM domain must meet the following criteria:
	//
	//   - It has to be a subdomain of the verified identity.
	//
	//   - It can't be used to receive email.
	//
	//   - It can't be used in a "From" address if the MAIL FROM domain is a
	//   destination for feedback forwarding emails.
	MailFromDomain *string

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutEmailIdentityMailFromAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailIdentityMailFromAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityMailFromAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityMailFromAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutEmailIdentityMailFromAttributes"); err != nil {
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
	if err = addOpPutEmailIdentityMailFromAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityMailFromAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailIdentityMailFromAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutEmailIdentityMailFromAttributes",
	}
}
