// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the SMB security strategy level for an Amazon S3 file gateway. This
// action is only supported for Amazon S3 file gateways.
//
// For information about configuring this setting using the Amazon Web Services
// console, see [Setting a security level for your gateway]in the Amazon S3 File Gateway User Guide.
//
// A higher security strategy level can affect performance of the gateway.
//
// [Setting a security level for your gateway]: https://docs.aws.amazon.com/filegateway/latest/files3/security-strategy.html
func (c *Client) UpdateSMBSecurityStrategy(ctx context.Context, params *UpdateSMBSecurityStrategyInput, optFns ...func(*Options)) (*UpdateSMBSecurityStrategyOutput, error) {
	if params == nil {
		params = &UpdateSMBSecurityStrategyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSMBSecurityStrategy", params, optFns, c.addOperationUpdateSMBSecurityStrategyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSMBSecurityStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSMBSecurityStrategyInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies the type of security strategy.
	//
	// ClientSpecified : If you choose this option, requests are established based on
	// what is negotiated by the client. This option is recommended when you want to
	// maximize compatibility across different clients in your environment. Supported
	// only for S3 File Gateway.
	//
	// MandatorySigning : If you choose this option, File Gateway only allows
	// connections from SMBv2 or SMBv3 clients that have signing enabled. This option
	// works with SMB clients on Microsoft Windows Vista, Windows Server 2008 or newer.
	//
	// MandatoryEncryption : If you choose this option, File Gateway only allows
	// connections from SMBv3 clients that have encryption enabled. This option is
	// recommended for environments that handle sensitive data. This option works with
	// SMB clients on Microsoft Windows 8, Windows Server 2012 or newer.
	//
	// MandatoryEncryptionNoAes128 : If you choose this option, File Gateway only
	// allows connections from SMBv3 clients that use 256-bit AES encryption
	// algorithms. 128-bit algorithms are not allowed. This option is recommended for
	// environments that handle sensitive data. It works with SMB clients on Microsoft
	// Windows 8, Windows Server 2012, or later.
	//
	// This member is required.
	SMBSecurityStrategy types.SMBSecurityStrategy

	noSmithyDocumentSerde
}

type UpdateSMBSecurityStrategyOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to return a
	// list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSMBSecurityStrategyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSMBSecurityStrategy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSMBSecurityStrategy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSMBSecurityStrategy"); err != nil {
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
	if err = addOpUpdateSMBSecurityStrategyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSMBSecurityStrategy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSMBSecurityStrategy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSMBSecurityStrategy",
	}
}
