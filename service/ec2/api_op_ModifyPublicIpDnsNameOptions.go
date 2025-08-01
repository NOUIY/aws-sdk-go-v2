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

// Modify public hostname options for a network interface. For more information,
// see [EC2 instance hostnames, DNS names, and domains]in the Amazon EC2 User Guide.
//
// [EC2 instance hostnames, DNS names, and domains]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html
func (c *Client) ModifyPublicIpDnsNameOptions(ctx context.Context, params *ModifyPublicIpDnsNameOptionsInput, optFns ...func(*Options)) (*ModifyPublicIpDnsNameOptionsOutput, error) {
	if params == nil {
		params = &ModifyPublicIpDnsNameOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyPublicIpDnsNameOptions", params, optFns, c.addOperationModifyPublicIpDnsNameOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyPublicIpDnsNameOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyPublicIpDnsNameOptionsInput struct {

	// The public hostname type. For more information, see [EC2 instance hostnames, DNS names, and domains] in the Amazon EC2 User
	// Guide.
	//
	//   - public-dual-stack-dns-name : A dual-stack public hostname for a network
	//   interface. Requests from within the VPC resolve to both the private IPv4 address
	//   and the IPv6 Global Unicast Address of the network interface. Requests from the
	//   internet resolve to both the public IPv4 and the IPv6 GUA address of the network
	//   interface.
	//
	//   - public-ipv4-dns-name : An IPv4-enabled public hostname for a network
	//   interface. Requests from within the VPC resolve to the private primary IPv4
	//   address of the network interface. Requests from the internet resolve to the
	//   public IPv4 address of the network interface.
	//
	//   - public-ipv6-dns-name : An IPv6-enabled public hostname for a network
	//   interface. Requests from within the VPC or from the internet resolve to the IPv6
	//   GUA of the network interface.
	//
	// [EC2 instance hostnames, DNS names, and domains]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html
	//
	// This member is required.
	HostnameType types.PublicIpDnsOption

	// A network interface ID.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Checks whether you have the required permissions for the operation, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyPublicIpDnsNameOptionsOutput struct {

	// Whether or not the request was successful.
	Successful *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyPublicIpDnsNameOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyPublicIpDnsNameOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyPublicIpDnsNameOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyPublicIpDnsNameOptions"); err != nil {
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
	if err = addOpModifyPublicIpDnsNameOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyPublicIpDnsNameOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyPublicIpDnsNameOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyPublicIpDnsNameOptions",
	}
}
