// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the options for a VPN tunnel in an Amazon Web Services Site-to-Site
// VPN connection. You can modify multiple options for a tunnel in a single
// request, but you can only modify one tunnel at a time. For more information, see
// [Site-to-Site VPN tunnel options for your Site-to-Site VPN connection]in the Amazon Web Services Site-to-Site VPN User Guide.
//
// [Site-to-Site VPN tunnel options for your Site-to-Site VPN connection]: https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNTunnels.html
func (c *Client) ModifyVpnTunnelOptions(ctx context.Context, params *ModifyVpnTunnelOptionsInput, optFns ...func(*Options)) (*ModifyVpnTunnelOptionsOutput, error) {
	if params == nil {
		params = &ModifyVpnTunnelOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpnTunnelOptions", params, optFns, c.addOperationModifyVpnTunnelOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpnTunnelOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnTunnelOptionsInput struct {

	// The tunnel options to modify.
	//
	// This member is required.
	TunnelOptions *types.ModifyVpnTunnelOptionsSpecification

	// The ID of the Amazon Web Services Site-to-Site VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	// The external IP address of the VPN tunnel.
	//
	// This member is required.
	VpnTunnelOutsideIpAddress *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Specifies the storage mode for the pre-shared key (PSK). Valid values are
	// Standard (stored in Site-to-Site VPN service) or SecretsManager (stored in
	// Amazon Web Services Secrets Manager).
	PreSharedKeyStorage *string

	// Choose whether or not to trigger immediate tunnel replacement. This is only
	// applicable when turning on or off EnableTunnelLifecycleControl .
	//
	// Valid values: True | False
	SkipTunnelReplacement *bool

	noSmithyDocumentSerde
}

type ModifyVpnTunnelOptionsOutput struct {

	// Information about the VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpnTunnelOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnTunnelOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnTunnelOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyVpnTunnelOptions"); err != nil {
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
	if err = addOpModifyVpnTunnelOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnTunnelOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpnTunnelOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyVpnTunnelOptions",
	}
}
