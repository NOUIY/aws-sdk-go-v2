// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates values stored in the domain configuration. Domain configurations for
// default endpoints can't be updated.
//
// Requires permission to access the [UpdateDomainConfiguration] action.
//
// [UpdateDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) UpdateDomainConfiguration(ctx context.Context, params *UpdateDomainConfigurationInput, optFns ...func(*Options)) (*UpdateDomainConfigurationOutput, error) {
	if params == nil {
		params = &UpdateDomainConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainConfiguration", params, optFns, c.addOperationUpdateDomainConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainConfigurationInput struct {

	// The name of the domain configuration to be updated.
	//
	// This member is required.
	DomainConfigurationName *string

	// An enumerated string that speciﬁes the application-layer protocol.
	//
	//   - SECURE_MQTT - MQTT over TLS.
	//
	//   - MQTT_WSS - MQTT over WebSocket.
	//
	//   - HTTPS - HTTP over TLS.
	//
	//   - DEFAULT - Use a combination of port and Application Layer Protocol
	//   Negotiation (ALPN) to specify application_layer protocol. For more information,
	//   see [Device communication protocols].
	//
	// [Device communication protocols]: https://docs.aws.amazon.com/iot/latest/developerguide/protocols.html
	ApplicationProtocol types.ApplicationProtocol

	// An enumerated string that speciﬁes the authentication type.
	//
	//   - CUSTOM_AUTH_X509 - Use custom authentication and authorization with
	//   additional details from the X.509 client certificate.
	//
	//   - CUSTOM_AUTH - Use custom authentication and authorization. For more
	//   information, see [Custom authentication and authorization].
	//
	//   - AWS_X509 - Use X.509 client certificates without custom authentication and
	//   authorization. For more information, see [X.509 client certificates].
	//
	//   - AWS_SIGV4 - Use Amazon Web Services Signature Version 4. For more
	//   information, see [IAM users, groups, and roles].
	//
	//   - DEFAULT - Use a combination of port and Application Layer Protocol
	//   Negotiation (ALPN) to specify authentication type. For more information, see [Device communication protocols].
	//
	// [Custom authentication and authorization]: https://docs.aws.amazon.com/iot/latest/developerguide/custom-authentication.html
	// [X.509 client certificates]: https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html
	// [IAM users, groups, and roles]: https://docs.aws.amazon.com/iot/latest/developerguide/custom-authentication.html
	// [Device communication protocols]: https://docs.aws.amazon.com/iot/latest/developerguide/protocols.html
	AuthenticationType types.AuthenticationType

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *types.AuthorizerConfig

	// An object that speciﬁes the client certificate conﬁguration for a domain.
	ClientCertificateConfig *types.ClientCertificateConfig

	// The status to which the domain configuration should be updated.
	DomainConfigurationStatus types.DomainConfigurationStatus

	// Removes the authorization configuration from a domain.
	RemoveAuthorizerConfig bool

	// The server certificate configuration.
	ServerCertificateConfig *types.ServerCertificateConfig

	// An object that specifies the TLS configuration for a domain.
	TlsConfig *types.TlsConfig

	noSmithyDocumentSerde
}

type UpdateDomainConfigurationOutput struct {

	// The ARN of the domain configuration that was updated.
	DomainConfigurationArn *string

	// The name of the domain configuration that was updated.
	DomainConfigurationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDomainConfiguration"); err != nil {
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
	if err = addOpUpdateDomainConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDomainConfiguration",
	}
}
