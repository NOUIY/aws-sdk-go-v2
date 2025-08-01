// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a listener for the specified Application Load Balancer, Network Load
// Balancer, or Gateway Load Balancer.
//
// For more information, see the following:
//
// [Listeners for your Application Load Balancers]
//
// [Listeners for your Network Load Balancers]
//
// [Listeners for your Gateway Load Balancers]
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple listeners with the same settings, each call
// succeeds.
//
// [Listeners for your Gateway Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-listeners.html
// [Listeners for your Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
// [Listeners for your Network Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html
func (c *Client) CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) {
	if params == nil {
		params = &CreateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateListener", params, optFns, c.addOperationCreateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateListenerInput struct {

	// The actions for the default rule.
	//
	// This member is required.
	DefaultActions []types.Action

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	// [TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN)
	// policy. You can specify one policy name. The following are the possible values:
	//
	//   - HTTP1Only
	//
	//   - HTTP2Only
	//
	//   - HTTP2Optional
	//
	//   - HTTP2Preferred
	//
	//   - None
	//
	// For more information, see [ALPN policies] in the Network Load Balancers Guide.
	//
	// [ALPN policies]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html#alpn-policies
	AlpnPolicy []string

	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN but
	// do not set IsDefault .
	Certificates []types.Certificate

	// The mutual authentication configuration information.
	MutualAuthentication *types.MutualAuthenticationAttributes

	// The port on which the load balancer is listening. You can't specify a port for
	// a Gateway Load Balancer.
	Port *int32

	// The protocol for connections from clients to the load balancer. For Application
	// Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t
	// specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You can't
	// specify a protocol for a Gateway Load Balancer.
	Protocol types.ProtocolEnum

	// [HTTPS and TLS listeners] The security policy that defines which protocols and
	// ciphers are supported.
	//
	// For more information, see [Security policies] in the Application Load Balancers Guide and [Security policies] in the
	// Network Load Balancers Guide.
	//
	// [Security policies]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html
	SslPolicy *string

	// The tags to assign to the listener.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateListenerOutput struct {

	// Information about the listener.
	Listeners []types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateListener{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateListener"); err != nil {
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
	if err = addOpCreateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateListener",
	}
}
