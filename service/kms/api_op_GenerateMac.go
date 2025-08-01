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

// Generates a hash-based message authentication code (HMAC) for a message using
// an HMAC KMS key and a MAC algorithm that the key supports. HMAC KMS keys and the
// HMAC algorithms that KMS uses conform to industry standards defined in [RFC 2104].
//
// You can use value that GenerateMac returns in the VerifyMac operation to demonstrate
// that the original message has not changed. Also, because a secret key is used to
// create the hash, you can verify that the party that generated the hash has the
// required secret key. You can also use the raw result to implement HMAC-based
// algorithms such as key derivation functions. This operation is part of KMS
// support for HMAC KMS keys. For details, see [HMAC keys in KMS]in the Key Management Service
// Developer Guide .
//
// Best practices recommend that you limit the time during which any signing
// mechanism, including an HMAC, is effective. This deters an attack where the
// actor uses a signed message to establish validity repeatedly or long after the
// message is superseded. HMAC tags do not include a timestamp, but you can include
// a timestamp in the token or message to help you detect when its time to refresh
// the HMAC.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateMac] (key policy)
//
// Related operations: VerifyMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:GenerateMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
func (c *Client) GenerateMac(ctx context.Context, params *GenerateMacInput, optFns ...func(*Options)) (*GenerateMacOutput, error) {
	if params == nil {
		params = &GenerateMacInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateMac", params, optFns, c.addOperationGenerateMacMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateMacOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateMacInput struct {

	// The HMAC KMS key to use in the operation. The MAC algorithm computes the HMAC
	// for the message and the key as described in [RFC 2104].
	//
	// To identify an HMAC KMS key, use the DescribeKey operation and see the KeySpec field in
	// the response.
	//
	// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
	//
	// This member is required.
	KeyId *string

	// The MAC algorithm used in the operation.
	//
	// The algorithm must be compatible with the HMAC KMS key that you specify. To
	// find the MAC algorithms that your HMAC KMS key supports, use the DescribeKeyoperation and
	// see the MacAlgorithms field in the DescribeKey response.
	//
	// This member is required.
	MacAlgorithm types.MacAlgorithmSpec

	// The message to be hashed. Specify a message of up to 4,096 bytes.
	//
	// GenerateMac and VerifyMac do not provide special handling for message digests. If you
	// generate an HMAC for a hash digest of a message, you must verify the HMAC of the
	// same hash digest.
	//
	// This member is required.
	Message []byte

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your permissions] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your permissions]: https://docs.aws.amazon.com/kms/latest/developerguide/testing-permissions.html
	DryRun *bool

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

type GenerateMacOutput struct {

	// The HMAC KMS key used in the operation.
	KeyId *string

	// The hash-based message authentication code (HMAC) that was generated for the
	// specified message, HMAC KMS key, and MAC algorithm.
	//
	// This is the standard, raw HMAC defined in [RFC 2104].
	//
	// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
	Mac []byte

	// The MAC algorithm that was used to generate the HMAC.
	MacAlgorithm types.MacAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateMacMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGenerateMac{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGenerateMac{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateMac"); err != nil {
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
	if err = addOpGenerateMacValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateMac(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateMac(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateMac",
	}
}
