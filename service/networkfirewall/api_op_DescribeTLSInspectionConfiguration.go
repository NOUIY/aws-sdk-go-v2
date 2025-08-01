// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the data objects for the specified TLS inspection configuration.
func (c *Client) DescribeTLSInspectionConfiguration(ctx context.Context, params *DescribeTLSInspectionConfigurationInput, optFns ...func(*Options)) (*DescribeTLSInspectionConfigurationOutput, error) {
	if params == nil {
		params = &DescribeTLSInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTLSInspectionConfiguration", params, optFns, c.addOperationDescribeTLSInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTLSInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTLSInspectionConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the TLS inspection configuration.
	//
	// You must specify the ARN or the name, and you can specify both.
	TLSInspectionConfigurationArn *string

	// The descriptive name of the TLS inspection configuration. You can't change the
	// name of a TLS inspection configuration after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	TLSInspectionConfigurationName *string

	noSmithyDocumentSerde
}

type DescribeTLSInspectionConfigurationOutput struct {

	// The high-level properties of a TLS inspection configuration. This, along with
	// the TLSInspectionConfiguration, define the TLS inspection configuration. You can retrieve all objects for
	// a TLS inspection configuration by calling DescribeTLSInspectionConfiguration.
	//
	// This member is required.
	TLSInspectionConfigurationResponse *types.TLSInspectionConfigurationResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the TLS inspection configuration. The token marks the state
	// of the TLS inspection configuration resource at the time of the request.
	//
	// To make changes to the TLS inspection configuration, you provide the token in
	// your request. Network Firewall uses the token to ensure that the TLS inspection
	// configuration hasn't changed since you last retrieved it. If it has changed, the
	// operation fails with an InvalidTokenException . If this happens, retrieve the
	// TLS inspection configuration again to get a current copy of it with a current
	// token. Reapply your changes as needed, then try the operation again using the
	// new token.
	//
	// This member is required.
	UpdateToken *string

	// The object that defines a TLS inspection configuration. This, along with TLSInspectionConfigurationResponse,
	// define the TLS inspection configuration. You can retrieve all objects for a TLS
	// inspection configuration by calling DescribeTLSInspectionConfiguration.
	//
	// Network Firewall uses a TLS inspection configuration to decrypt traffic.
	// Network Firewall re-encrypts the traffic before sending it to its destination.
	//
	// To use a TLS inspection configuration, you add it to a new Network Firewall
	// firewall policy, then you apply the firewall policy to a firewall. Network
	// Firewall acts as a proxy service to decrypt and inspect the traffic traveling
	// through your firewalls. You can reference a TLS inspection configuration from
	// more than one firewall policy, and you can use a firewall policy in more than
	// one firewall. For more information about using TLS inspection configurations,
	// see [Inspecting SSL/TLS traffic with TLS inspection configurations]in the Network Firewall Developer Guide.
	//
	// [Inspecting SSL/TLS traffic with TLS inspection configurations]: https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html
	TLSInspectionConfiguration *types.TLSInspectionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTLSInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTLSInspectionConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTLSInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTLSInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTLSInspectionConfiguration",
	}
}
