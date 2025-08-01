// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorscep

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a SCEP connector. A SCEP connector links Amazon Web Services Private
// Certificate Authority to your SCEP-compatible devices and mobile device
// management (MDM) systems. Before you create a connector, you must complete a set
// of prerequisites, including creation of a private certificate authority (CA) to
// use with this connector. For more information, see [Connector for SCEP prerequisites].
//
// [Connector for SCEP prerequisites]: https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-prerequisites.html
func (c *Client) CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) {
	if params == nil {
		params = &CreateConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnector", params, optFns, c.addOperationCreateConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorInput struct {

	// The Amazon Resource Name (ARN) of the Amazon Web Services Private Certificate
	// Authority certificate authority to use with this connector. Due to security
	// vulnerabilities present in the SCEP protocol, we recommend using a private CA
	// that's dedicated for use with the connector.
	//
	// To retrieve the private CAs associated with your account, you can call [ListCertificateAuthorities] using
	// the Amazon Web Services Private CA API.
	//
	// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
	//
	// This member is required.
	CertificateAuthorityArn *string

	// Custom string that can be used to distinguish between calls to the [CreateChallenge] action.
	// Client tokens for CreateChallenge time out after five minutes. Therefore, if
	// you call CreateChallenge multiple times with the same client token within five
	// minutes, Connector for SCEP recognizes that you are requesting only one
	// challenge and will only respond with one. If you change the client token for
	// each call, Connector for SCEP recognizes that you are requesting multiple
	// challenge passwords.
	//
	// [CreateChallenge]: https://docs.aws.amazon.com/C4SCEP_API/pca-connector-scep/latest/APIReference/API_CreateChallenge.html
	ClientToken *string

	// If you don't supply a value, by default Connector for SCEP creates a connector
	// for general-purpose use. A general-purpose connector is designed to work with
	// clients or endpoints that support the SCEP protocol, except Connector for SCEP
	// for Microsoft Intune. With connectors for general-purpose use, you manage SCEP
	// challenge passwords using Connector for SCEP. For information about
	// considerations and limitations with using Connector for SCEP, see [Considerations and Limitations].
	//
	// If you provide an IntuneConfiguration , Connector for SCEP creates a connector
	// for use with Microsoft Intune, and you manage the challenge passwords using
	// Microsoft Intune. For more information, see [Using Connector for SCEP for Microsoft Intune].
	//
	// [Considerations and Limitations]: https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlc4scep-considerations-limitations.html
	// [Using Connector for SCEP for Microsoft Intune]: https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-intune.html
	MobileDeviceManagement types.MobileDeviceManagement

	// The key-value pairs to associate with the resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateConnectorOutput struct {

	// Returns the Amazon Resource Name (ARN) of the connector.
	ConnectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConnector"); err != nil {
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
	if err = addIdempotencyToken_opCreateConnectorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnector(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateConnector struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConnector) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConnectorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConnectorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateConnector{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConnector",
	}
}
