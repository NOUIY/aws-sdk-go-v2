// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a VPC with a service network. When you associate a VPC with the
// service network, it enables all the resources within that VPC to be clients and
// communicate with other services in the service network. For more information,
// see [Manage VPC associations]in the Amazon VPC Lattice User Guide.
//
// You can't use this operation if there is a disassociation in progress. If the
// association fails, retry by deleting the association and recreating it.
//
// As a result of this operation, the association gets created in the service
// network account and the VPC owner account.
//
// If you add a security group to the service network and VPC association, the
// association must continue to always have at least one security group. You can
// add or edit security groups at any time. However, to remove all security groups,
// you must first delete the association and recreate it without security groups.
//
// [Manage VPC associations]: https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-vpc-associations
func (c *Client) CreateServiceNetworkVpcAssociation(ctx context.Context, params *CreateServiceNetworkVpcAssociationInput, optFns ...func(*Options)) (*CreateServiceNetworkVpcAssociationOutput, error) {
	if params == nil {
		params = &CreateServiceNetworkVpcAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceNetworkVpcAssociation", params, optFns, c.addOperationCreateServiceNetworkVpcAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceNetworkVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceNetworkVpcAssociationInput struct {

	// The ID or ARN of the service network. You must use an ARN if the resources are
	// in different accounts.
	//
	// This member is required.
	ServiceNetworkIdentifier *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcIdentifier *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you retry a request that completed successfully using the
	// same client token and parameters, the retry succeeds without performing any
	// actions. If the parameters aren't identical, the retry fails.
	ClientToken *string

	// The IDs of the security groups. Security groups aren't added by default. You
	// can add a security group to apply network level controls to control which
	// resources in a VPC are allowed to access the service network and its services.
	// For more information, see [Control traffic to resources using security groups]in the Amazon VPC User Guide.
	//
	// [Control traffic to resources using security groups]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	SecurityGroupIds []string

	// The tags for the association.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateServiceNetworkVpcAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The account that created the association.
	CreatedBy *string

	// The ID of the association.
	Id *string

	// The IDs of the security groups.
	SecurityGroupIds []string

	// The association status.
	Status types.ServiceNetworkVpcAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceNetworkVpcAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateServiceNetworkVpcAssociation"); err != nil {
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
	if err = addIdempotencyToken_opCreateServiceNetworkVpcAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceNetworkVpcAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceNetworkVpcAssociation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceNetworkVpcAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceNetworkVpcAssociationInput ")
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
func addIdempotencyToken_opCreateServiceNetworkVpcAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceNetworkVpcAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateServiceNetworkVpcAssociation",
	}
}
