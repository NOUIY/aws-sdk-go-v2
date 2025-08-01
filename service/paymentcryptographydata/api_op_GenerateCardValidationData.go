// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates card-related validation data using algorithms such as Card
// Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2),
// or Card Security Codes (CSC). For more information, see [Generate card data]in the Amazon Web
// Services Payment Cryptography User Guide.
//
// This operation generates a CVV or CSC value that is printed on a payment credit
// or debit card during card production. The CVV or CSC, PAN (Primary Account
// Number) and expiration date of the card are required to check its validity
// during transaction processing. To begin this operation, a CVK (Card Verification
// Key) encryption key is required. You can use [CreateKey]or [ImportKey] to establish a CVK within
// Amazon Web Services Payment Cryptography. The KeyModesOfUse should be set to
// Generate and Verify for a CVK encryption key.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ImportKey]
//
// # VerifyCardValidationData
//
// [Generate card data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/generate-card-data.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func (c *Client) GenerateCardValidationData(ctx context.Context, params *GenerateCardValidationDataInput, optFns ...func(*Options)) (*GenerateCardValidationDataOutput, error) {
	if params == nil {
		params = &GenerateCardValidationDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateCardValidationData", params, optFns, c.addOperationGenerateCardValidationDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateCardValidationDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateCardValidationDataInput struct {

	// The algorithm for generating CVV or CSC values for the card within Amazon Web
	// Services Payment Cryptography.
	//
	// This member is required.
	GenerationAttributes types.CardGenerationAttributes

	// The keyARN of the CVK encryption key that Amazon Web Services Payment
	// Cryptography uses to generate card data.
	//
	// This member is required.
	KeyIdentifier *string

	// The Primary Account Number (PAN), a unique identifier for a payment credit or
	// debit card that associates the card with a specific account holder.
	//
	// This member is required.
	PrimaryAccountNumber *string

	// The length of the CVV or CSC to be generated. The default value is 3.
	ValidationDataLength *int32

	noSmithyDocumentSerde
}

type GenerateCardValidationDataOutput struct {

	// The keyARN of the CVK encryption key that Amazon Web Services Payment
	// Cryptography uses to generate CVV or CSC.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed.
	//
	// Amazon Web Services Payment Cryptography computes the KCV according to the CMAC
	// specification.
	//
	// This member is required.
	KeyCheckValue *string

	// The CVV or CSC value that Amazon Web Services Payment Cryptography generates
	// for the card.
	//
	// This member is required.
	ValidationData *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateCardValidationDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateCardValidationData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateCardValidationData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateCardValidationData"); err != nil {
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
	if err = addOpGenerateCardValidationDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateCardValidationData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateCardValidationData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateCardValidationData",
	}
}
