// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves information about an HSM. You can identify the HSM by its ARN or its
// serial number.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
func (c *Client) DescribeHsm(ctx context.Context, params *DescribeHsmInput, optFns ...func(*Options)) (*DescribeHsmOutput, error) {
	if params == nil {
		params = &DescribeHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHsm", params, optFns, c.addOperationDescribeHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeHsm operation.
type DescribeHsmInput struct {

	// The ARN of the HSM. Either the HsmArn or the SerialNumber parameter must be
	// specified.
	HsmArn *string

	// The serial number of the HSM. Either the HsmArn or the HsmSerialNumber
	// parameter must be specified.
	HsmSerialNumber *string

	noSmithyDocumentSerde
}

// Contains the output of the DescribeHsm operation.
type DescribeHsmOutput struct {

	// The Availability Zone that the HSM is in.
	AvailabilityZone *string

	// The identifier of the elastic network interface (ENI) attached to the HSM.
	EniId *string

	// The IP address assigned to the HSM's ENI.
	EniIp *string

	// The ARN of the HSM.
	HsmArn *string

	// The HSM model type.
	HsmType *string

	// The ARN of the IAM role assigned to the HSM.
	IamRoleArn *string

	// The list of partitions on the HSM.
	Partitions []string

	// The serial number of the HSM.
	SerialNumber *string

	// The date and time that the server certificate was last updated.
	ServerCertLastUpdated *string

	// The URI of the certificate server.
	ServerCertUri *string

	// The HSM software version.
	SoftwareVersion *string

	// The date and time that the SSH key was last updated.
	SshKeyLastUpdated *string

	// The public SSH key.
	SshPublicKey *string

	// The status of the HSM.
	Status types.HsmStatus

	// Contains additional information about the status of the HSM.
	StatusDetails *string

	// The identifier of the subnet that the HSM is in.
	SubnetId *string

	// The subscription end date.
	SubscriptionEndDate *string

	// The subscription start date.
	SubscriptionStartDate *string

	// Specifies the type of subscription for the HSM.
	//
	//   - PRODUCTION - The HSM is being used in a production environment.
	//
	//   - TRIAL - The HSM is being used in a product trial.
	SubscriptionType types.SubscriptionType

	// The name of the HSM vendor.
	VendorName *string

	// The identifier of the VPC that the HSM is in.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHsm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHsm"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHsm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHsm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHsm",
	}
}
