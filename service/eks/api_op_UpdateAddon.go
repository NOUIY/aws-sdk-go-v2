// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon EKS add-on.
func (c *Client) UpdateAddon(ctx context.Context, params *UpdateAddonInput, optFns ...func(*Options)) (*UpdateAddonOutput, error) {
	if params == nil {
		params = &UpdateAddonInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAddon", params, optFns, c.addOperationUpdateAddonMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAddonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAddonInput struct {

	// The name of the add-on. The name must match one of the names returned by [ListAddons]
	// ListAddons .
	//
	// [ListAddons]: https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html
	//
	// This member is required.
	AddonName *string

	// The name of your cluster.
	//
	// This member is required.
	ClusterName *string

	// The version of the add-on. The version must match one of the versions returned
	// by [DescribeAddonVersions]DescribeAddonVersions .
	//
	// [DescribeAddonVersions]: https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html
	AddonVersion *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// The set of configuration values for the add-on that's created. The values that
	// you provide are validated against the schema returned by
	// DescribeAddonConfiguration .
	ConfigurationValues *string

	// An array of EKS Pod Identity associations to be updated. Each association maps
	// a Kubernetes service account to an IAM role. If this value is left blank, no
	// change. If an empty array is provided, existing associations owned by the add-on
	// are deleted.
	//
	// For more information, see [Attach an IAM Role to an Amazon EKS add-on using EKS Pod Identity] in the Amazon EKS User Guide.
	//
	// [Attach an IAM Role to an Amazon EKS add-on using EKS Pod Identity]: https://docs.aws.amazon.com/eks/latest/userguide/add-ons-iam.html
	PodIdentityAssociations []types.AddonPodIdentityAssociations

	// How to resolve field value conflicts for an Amazon EKS add-on if you've changed
	// a value from the Amazon EKS default value. Conflicts are handled based on the
	// option you choose:
	//
	//   - None – Amazon EKS doesn't change the value. The update might fail.
	//
	//   - Overwrite – Amazon EKS overwrites the changed value back to the Amazon EKS
	//   default value.
	//
	//   - Preserve – Amazon EKS preserves the value. If you choose this option, we
	//   recommend that you test any field and value changes on a non-production cluster
	//   before updating the add-on on your production cluster.
	ResolveConflicts types.ResolveConflicts

	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's
	// service account. The role must be assigned the IAM permissions required by the
	// add-on. If you don't specify an existing IAM role, then the add-on uses the
	// permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role]in the
	// Amazon EKS User Guide.
	//
	// To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC)
	// provider created for your cluster. For more information, see [Enabling IAM roles for service accounts on your cluster]in the Amazon EKS
	// User Guide.
	//
	// [Enabling IAM roles for service accounts on your cluster]: https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html
	// [Amazon EKS node IAM role]: https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html
	ServiceAccountRoleArn *string

	noSmithyDocumentSerde
}

type UpdateAddonOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAddonMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAddon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAddon{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAddon"); err != nil {
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
	if err = addIdempotencyToken_opUpdateAddonMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateAddonValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAddon(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateAddon struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAddon) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAddonInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateAddonMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAddon{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAddon(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAddon",
	}
}
