// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates environment network to connect to your internal network by using a
// transit gateway. This API supports request to create a transit gateway
// attachment from FinSpace VPC to your transit gateway ID and create a custom
// Route-53 outbound resolvers.
//
// Once you send a request to update a network, you cannot change it again.
// Network update might require termination of any clusters that are running in the
// existing network.
func (c *Client) UpdateKxEnvironmentNetwork(ctx context.Context, params *UpdateKxEnvironmentNetworkInput, optFns ...func(*Options)) (*UpdateKxEnvironmentNetworkOutput, error) {
	if params == nil {
		params = &UpdateKxEnvironmentNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKxEnvironmentNetwork", params, optFns, c.addOperationUpdateKxEnvironmentNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKxEnvironmentNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKxEnvironmentNetworkInput struct {

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	// A list of DNS server name and server IP. This is used to set up Route-53
	// outbound resolvers.
	CustomDNSConfiguration []types.CustomDNSServer

	// Specifies the transit gateway and network configuration to connect the kdb
	// environment to an internal network.
	TransitGatewayConfiguration *types.TransitGatewayConfiguration

	noSmithyDocumentSerde
}

type UpdateKxEnvironmentNetworkOutput struct {

	// The identifier of the availability zones where subnets for the environment are
	// created.
	AvailabilityZoneIds []string

	// The unique identifier of the AWS account that is used to create the kdb
	// environment.
	AwsAccountId *string

	// The timestamp at which the kdb environment was created in FinSpace.
	CreationTimestamp *time.Time

	// A list of DNS server name and server IP. This is used to set up Route-53
	// outbound resolvers.
	CustomDNSConfiguration []types.CustomDNSServer

	// A unique identifier for the AWS environment infrastructure account.
	DedicatedServiceAccountId *string

	// The description of the environment.
	Description *string

	// The status of DNS configuration.
	DnsStatus types.DnsStatus

	// The ARN identifier of the environment.
	EnvironmentArn *string

	// A unique identifier for the kdb environment.
	EnvironmentId *string

	// Specifies the error message that appears if a flow fails.
	ErrorMessage *string

	// The KMS key ID to encrypt your data in the FinSpace environment.
	KmsKeyId *string

	// The name of the kdb environment.
	Name *string

	// The status of the kdb environment.
	Status types.EnvironmentStatus

	// The status of the network configuration.
	TgwStatus types.TgwStatus

	// The structure of the transit gateway and network configuration that is used to
	// connect the kdb environment to an internal network.
	TransitGatewayConfiguration *types.TransitGatewayConfiguration

	// The timestamp at which the kdb environment was updated.
	UpdateTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKxEnvironmentNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateKxEnvironmentNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateKxEnvironmentNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKxEnvironmentNetwork"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opUpdateKxEnvironmentNetworkMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateKxEnvironmentNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKxEnvironmentNetwork(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateKxEnvironmentNetwork struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateKxEnvironmentNetwork) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateKxEnvironmentNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateKxEnvironmentNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateKxEnvironmentNetworkInput ")
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
func addIdempotencyToken_opUpdateKxEnvironmentNetworkMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateKxEnvironmentNetwork{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateKxEnvironmentNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKxEnvironmentNetwork",
	}
}
