// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides detailed information about a KMS key. You can run DescribeKey on a [customer managed key] or
// an [Amazon Web Services managed key].
//
// This detailed information includes the key ARN, creation date (and deletion
// date, if applicable), the key state, and the origin and expiration date (if any)
// of the key material. It includes fields, like KeySpec , that help you
// distinguish different types of KMS keys. It also displays the key usage
// (encryption, signing, or generating and verifying MACs) and the algorithms that
// the KMS key supports.
//
// For [multi-Region keys], DescribeKey displays the primary key and all related replica keys. For
// KMS keys in [CloudHSM key stores], it includes information about the key store, such as the key
// store ID and the CloudHSM cluster ID. For KMS keys in [external key stores], it includes the custom
// key store ID and the ID of the external key.
//
// DescribeKey does not return the following information:
//
//   - Aliases associated with the KMS key. To get this information, use ListAliases.
//
//   - Whether automatic key rotation is enabled on the KMS key. To get this
//     information, use GetKeyRotationStatus. Also, some key states prevent a KMS key from being
//     automatically rotated. For details, see [How key rotation works]in the Key Management Service
//     Developer Guide.
//
//   - Tags on the KMS key. To get this information, use ListResourceTags.
//
//   - Key policies and grants on the KMS key. To get this information, use GetKeyPolicyand ListGrants.
//
// In general, DescribeKey is a non-mutating operation. It returns data about KMS
// keys, but doesn't change them. However, Amazon Web Services services use
// DescribeKey to create [Amazon Web Services managed keys] from a predefined Amazon Web Services alias with no key
// ID.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DescribeKey] (key policy)
//
// Related operations:
//
// # GetKeyPolicy
//
// # GetKeyRotationStatus
//
// # ListAliases
//
// # ListGrants
//
// # ListKeys
//
// # ListResourceTags
//
// # ListRetirableGrants
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [How key rotation works]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-how-it-works
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [kms:DescribeKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
func (c *Client) DescribeKey(ctx context.Context, params *DescribeKeyInput, optFns ...func(*Options)) (*DescribeKeyOutput, error) {
	if params == nil {
		params = &DescribeKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeKey", params, optFns, c.addOperationDescribeKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeKeyInput struct {

	// Describes the specified KMS key.
	//
	// If you specify a predefined Amazon Web Services alias (an Amazon Web Services
	// alias with no key ID), KMS associates the alias with an [Amazon Web Services managed key]and returns its KeyId
	// and Arn in the response.
	//
	// To specify a KMS key, use its key ID, key ARN, alias name, or alias ARN. When
	// using an alias name, prefix it with "alias/" . To specify a KMS key in a
	// different Amazon Web Services account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Alias name: alias/ExampleAlias
	//
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name
	// and alias ARN, use ListAliases.
	//
	// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
	//
	// This member is required.
	KeyId *string

	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from a new
	// grant that has not yet achieved eventual consistency. For more information, see [Grant token]
	// and [Using a grant token]in the Key Management Service Developer Guide.
	//
	// [Grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
	// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/using-grant-token.html
	GrantTokens []string

	noSmithyDocumentSerde
}

type DescribeKeyOutput struct {

	// Metadata associated with the key.
	KeyMetadata *types.KeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeKey"); err != nil {
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
	if err = addOpDescribeKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeKey",
	}
}
