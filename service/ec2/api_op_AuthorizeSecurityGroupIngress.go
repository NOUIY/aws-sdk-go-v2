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

// Adds the specified inbound (ingress) rules to a security group.
//
// An inbound rule permits instances to receive traffic from the specified IPv4 or
// IPv6 address range, the IP address ranges that are specified by a prefix list,
// or the instances that are associated with a destination security group. For more
// information, see [Security group rules].
//
// You must specify exactly one of the following sources: an IPv4 or IPv6 address
// range, a prefix list, or a security group. You must specify a protocol for each
// rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a
// port or port range. If the protocol is ICMP or ICMPv6, you must also specify the
// ICMP/ICMPv6 type and code.
//
// Rule changes are propagated to instances associated with the security group as
// quickly as possible. However, a small delay might occur.
//
// For examples of rules that you can add to security groups for specific access
// scenarios, see [Security group rules for different use cases]in the Amazon EC2 User Guide.
//
// For more information about security group quotas, see [Amazon VPC quotas] in the Amazon VPC User
// Guide.
//
// [Amazon VPC quotas]: https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html
// [Security group rules]: https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html
// [Security group rules for different use cases]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html
func (c *Client) AuthorizeSecurityGroupIngress(ctx context.Context, params *AuthorizeSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &AuthorizeSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeSecurityGroupIngress", params, optFns, c.addOperationAuthorizeSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeSecurityGroupIngressInput struct {

	// The IPv4 address range, in CIDR format.
	//
	// Amazon Web Services [canonicalizes] IPv4 and IPv6 CIDRs. For example, if you specify
	// 100.68.0.18/18 for the CIDR block, Amazon Web Services canonicalizes the CIDR
	// block to 100.68.0.0/18. Any subsequent DescribeSecurityGroups and
	// DescribeSecurityGroupRules calls will return the canonicalized form of the CIDR
	// block. Additionally, if you attempt to add another rule with the non-canonical
	// form of the CIDR (such as 100.68.0.18/18) and there is already a rule for the
	// canonicalized form of the CIDR block (such as 100.68.0.0/18), the API throws an
	// duplicate rule error.
	//
	// To specify an IPv6 address range, use IP permissions instead.
	//
	// To specify multiple rules and descriptions for the rules, use IP permissions
	// instead.
	//
	// [canonicalizes]: https://en.wikipedia.org/wiki/Canonicalization
	CidrIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// If the protocol is TCP or UDP, this is the start of the port range. If the
	// protocol is ICMP, this is the ICMP type or -1 (all ICMP types).
	//
	// To specify multiple rules and descriptions for the rules, use IP permissions
	// instead.
	FromPort *int32

	// The ID of the security group.
	GroupId *string

	// [Default VPC] The name of the security group. For security groups for a default
	// VPC you can specify either the ID or the name of the security group. For
	// security groups for a nondefault VPC, you must specify the ID of the security
	// group.
	GroupName *string

	// The permissions for the security group rules.
	IpPermissions []types.IpPermission

	// The IP protocol name ( tcp , udp , icmp ) or number (see [Protocol Numbers]). To specify all
	// protocols, use -1 .
	//
	// To specify icmpv6 , use IP permissions instead.
	//
	// If you specify a protocol other than one of the supported values, traffic is
	// allowed on all ports, regardless of any ports that you specify.
	//
	// To specify multiple rules and descriptions for the rules, use IP permissions
	// instead.
	//
	// [Protocol Numbers]: http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	IpProtocol *string

	// [Default VPC] The name of the source security group.
	//
	// The rule grants full ICMP, UDP, and TCP access. To create a rule with a
	// specific protocol and port range, specify a set of IP permissions instead.
	SourceSecurityGroupName *string

	// The Amazon Web Services account ID for the source security group, if the source
	// security group is in a different account.
	//
	// The rule grants full ICMP, UDP, and TCP access. To create a rule with a
	// specific protocol and port range, use IP permissions instead.
	SourceSecurityGroupOwnerId *string

	// The tags applied to the security group rule.
	TagSpecifications []types.TagSpecification

	// If the protocol is TCP or UDP, this is the end of the port range. If the
	// protocol is ICMP, this is the ICMP code or -1 (all ICMP codes). If the start
	// port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	//
	// To specify multiple rules and descriptions for the rules, use IP permissions
	// instead.
	ToPort *int32

	noSmithyDocumentSerde
}

type AuthorizeSecurityGroupIngressOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool

	// Information about the inbound (ingress) security group rules that were added.
	SecurityGroupRules []types.SecurityGroupRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAuthorizeSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAuthorizeSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AuthorizeSecurityGroupIngress"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AuthorizeSecurityGroupIngress",
	}
}
