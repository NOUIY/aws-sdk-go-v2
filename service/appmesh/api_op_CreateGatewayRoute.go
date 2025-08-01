// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a gateway route.
//
// A gateway route is attached to a virtual gateway and routes traffic to an
// existing virtual service. If a route matches a request, it can distribute
// traffic to a target virtual service.
//
// For more information about gateway routes, see [Gateway routes].
//
// [Gateway routes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html
func (c *Client) CreateGatewayRoute(ctx context.Context, params *CreateGatewayRouteInput, optFns ...func(*Options)) (*CreateGatewayRouteOutput, error) {
	if params == nil {
		params = &CreateGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGatewayRoute", params, optFns, c.addOperationCreateGatewayRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGatewayRouteInput struct {

	// The name to use for the gateway route.
	//
	// This member is required.
	GatewayRouteName *string

	// The name of the service mesh to create the gateway route in.
	//
	// This member is required.
	MeshName *string

	// The gateway route specification to apply.
	//
	// This member is required.
	Spec *types.GatewayRouteSpec

	// The name of the virtual gateway to associate the gateway route with. If the
	// virtual gateway is in a shared mesh, then you must be the owner of the virtual
	// gateway resource.
	//
	// This member is required.
	VirtualGatewayName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then the account that you specify must share the
	// mesh with your account before you can create the resource in the service mesh.
	// For more information about mesh sharing, see [Working with shared meshes].
	//
	// [Working with shared meshes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html
	MeshOwner *string

	// Optional metadata that you can apply to the gateway route to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef

	noSmithyDocumentSerde
}

type CreateGatewayRouteOutput struct {

	// The full description of your gateway route following the create call.
	//
	// This member is required.
	GatewayRoute *types.GatewayRouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGatewayRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGatewayRoute"); err != nil {
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
	if err = addIdempotencyToken_opCreateGatewayRouteMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateGatewayRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGatewayRoute(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateGatewayRoute struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateGatewayRoute) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateGatewayRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateGatewayRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateGatewayRouteInput ")
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
func addIdempotencyToken_opCreateGatewayRouteMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateGatewayRoute{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateGatewayRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGatewayRoute",
	}
}
