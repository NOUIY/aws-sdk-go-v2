// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one or more Amazon Lightsail instances.
//
// The create instances operation supports tag-based access control via request
// tags. For more information, see the [Lightsail Developer Guide].
//
// [Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
func (c *Client) CreateInstances(ctx context.Context, params *CreateInstancesInput, optFns ...func(*Options)) (*CreateInstancesOutput, error) {
	if params == nil {
		params = &CreateInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstances", params, optFns, c.addOperationCreateInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstancesInput struct {

	// The Availability Zone in which to create your instance. Use the following
	// format: us-east-2a (case sensitive). You can get a list of Availability Zones
	// by using the [get regions]operation. Be sure to add the include Availability Zones parameter
	// to your request.
	//
	// [get regions]: http://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRegions.html
	//
	// This member is required.
	AvailabilityZone *string

	// The ID for a virtual private server image ( app_wordpress_x_x or app_lamp_x_x ).
	// Use the get blueprints operation to return a list of available images (or
	// blueprints).
	//
	// Use active blueprints when creating new instances. Inactive blueprints are
	// listed to support customers with existing instances and are not necessarily
	// available to create new instances. Blueprints are marked inactive when they
	// become outdated due to operating system updates or new application releases.
	//
	// This member is required.
	BlueprintId *string

	// The bundle of specification information for your virtual private server (or
	// instance), including the pricing plan ( medium_x_x ).
	//
	// This member is required.
	BundleId *string

	// The names to use for your new Lightsail instances. Separate multiple values
	// using quotation marks and commas, for example:
	// ["MyFirstInstance","MySecondInstance"]
	//
	// This member is required.
	InstanceNames []string

	// An array of objects representing the add-ons to enable for the new instance.
	AddOns []types.AddOnRequest

	// (Discontinued) The name for your custom image.
	//
	// In releases prior to June 12, 2017, this parameter was ignored by the API. It
	// is now discontinued.
	//
	// Deprecated: This member has been deprecated.
	CustomImageName *string

	// The IP address type for the instance.
	//
	// The possible values are ipv4 for IPv4 only, ipv6 for IPv6 only, and dualstack
	// for IPv4 and IPv6.
	//
	// The default value is dualstack .
	IpAddressType types.IpAddressType

	// The name of your key pair.
	KeyPairName *string

	// The tag keys and optional values to add to the resource during create.
	//
	// Use the TagResource action to tag a resource after it's created.
	Tags []types.Tag

	// A launch script you can create that configures a server with additional user
	// data. For example, you might want to run apt-get -y update .
	//
	// Depending on the machine image you choose, the command to get software on your
	// instance varies. Amazon Linux and CentOS use yum , Debian and Ubuntu use apt-get
	// , and FreeBSD uses pkg . For a complete list, see the [Amazon Lightsail Developer Guide].
	//
	// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/compare-options-choose-lightsail-instance-image
	UserData *string

	noSmithyDocumentSerde
}

type CreateInstancesOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateInstances"); err != nil {
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
	if err = addOpCreateInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateInstances",
	}
}
