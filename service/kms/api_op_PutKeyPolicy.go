// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a key policy to the specified KMS key.
//
// For more information about key policies, see [Key Policies] in the Key Management Service
// Developer Guide. For help writing and formatting a JSON policy document, see the
// [IAM JSON Policy Reference]in the Identity and Access Management User Guide . For examples of adding a key
// policy in multiple programming languages, see [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:PutKeyPolicy] (key policy)
//
// Related operations: GetKeyPolicy
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [IAM JSON Policy Reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
// [kms:PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_PutKeyPolicy_section.html
// [Key Policies]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func (c *Client) PutKeyPolicy(ctx context.Context, params *PutKeyPolicyInput, optFns ...func(*Options)) (*PutKeyPolicyOutput, error) {
	if params == nil {
		params = &PutKeyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutKeyPolicy", params, optFns, c.addOperationPutKeyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutKeyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutKeyPolicyInput struct {

	// Sets the key policy on the specified KMS key.
	//
	// Specify the key ID or key ARN of the KMS key.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// The key policy to attach to the KMS key.
	//
	// The key policy must meet the following criteria:
	//
	//   - The key policy must allow the calling principal to make a subsequent
	//   PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key
	//   becomes unmanageable. For more information, see [Default key policy]in the Key Management Service
	//   Developer Guide. (To omit this condition, set BypassPolicyLockoutSafetyCheck
	//   to true.)
	//
	//   - Each statement in the key policy must contain one or more principals. The
	//   principals in the key policy must exist and be visible to KMS. When you create a
	//   new Amazon Web Services principal, you might need to enforce a delay before
	//   including the new principal in a key policy because the new principal might not
	//   be immediately visible to KMS. For more information, see [Changes that I make are not always immediately visible]in the Amazon Web
	//   Services Identity and Access Management User Guide.
	//
	// If either of the required Resource or Action elements are missing from a key
	// policy statement, the policy statement has no effect. When a key policy
	// statement is missing one of these elements, the KMS console correctly reports an
	// error, but the PutKeyPolicy API request succeeds, even though the policy
	// statement is ineffective.
	//
	// For more information on required key policy elements, see [Elements in a key policy] in the Key
	// Management Service Developer Guide.
	//
	// A key policy document can include only the following characters:
	//
	//   - Printable ASCII characters from the space character ( \u0020 ) through the
	//   end of the ASCII character range.
	//
	//   - Printable characters in the Basic Latin and Latin-1 Supplement character
	//   set (through \u00FF ).
	//
	//   - The tab ( \u0009 ), line feed ( \u000A ), and carriage return ( \u000D )
	//   special characters
	//
	// If the key policy exceeds the length constraint, KMS returns a
	// LimitExceededException .
	//
	// For information about key policies, see [Key policies in KMS] in the Key Management Service
	// Developer Guide.For help writing and formatting a JSON policy document, see the [IAM JSON Policy Reference]
	// in the Identity and Access Management User Guide .
	//
	// [Key policies in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
	// [Elements in a key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-overview.html#key-policy-elements
	// [IAM JSON Policy Reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
	// [Default key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key
	// [Changes that I make are not always immediately visible]: https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency
	//
	// This member is required.
	Policy *string

	// Skips ("bypasses") the key policy lockout safety check. The default value is
	// false.
	//
	// Setting this value to true increases the risk that the KMS key becomes
	// unmanageable. Do not set this value to true indiscriminately.
	//
	// For more information, see [Default key policy] in the Key Management Service Developer Guide.
	//
	// Use this parameter only when you intend to prevent the principal that is making
	// the request from making a subsequent [PutKeyPolicy]request on the KMS key.
	//
	// [Default key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key
	// [PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html
	BypassPolicyLockoutSafetyCheck bool

	// The name of the key policy. If no policy name is specified, the default value
	// is default . The only valid value is default .
	PolicyName *string

	noSmithyDocumentSerde
}

type PutKeyPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutKeyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutKeyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutKeyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutKeyPolicy"); err != nil {
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
	if err = addOpPutKeyPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutKeyPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutKeyPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutKeyPolicy",
	}
}
