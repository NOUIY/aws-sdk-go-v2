// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a new custom connector with your Amazon Web Services account. Before
// you can register the connector, you must deploy the associated AWS lambda
// function in your account.
func (c *Client) RegisterConnector(ctx context.Context, params *RegisterConnectorInput, optFns ...func(*Options)) (*RegisterConnectorOutput, error) {
	if params == nil {
		params = &RegisterConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterConnector", params, optFns, c.addOperationRegisterConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterConnectorInput struct {

	// The clientToken parameter is an idempotency token. It ensures that your
	// RegisterConnector request completes only once. You choose the value to pass. For
	// example, if you don't receive a response from your request, you can safely retry
	// the request with the same clientToken parameter value.
	//
	// If you omit a clientToken value, the Amazon Web Services SDK that you are using
	// inserts a value for you. This way, the SDK can safely retry requests multiple
	// times after a network error. You must provide your own value for other use
	// cases.
	//
	// If you specify input parameters that differ from your first request, an error
	// occurs. If you use a different value for clientToken , Amazon AppFlow considers
	// it a new call to RegisterConnector . The token is active for 8 hours.
	ClientToken *string

	//  The name of the connector. The name is unique for each ConnectorRegistration
	// in your Amazon Web Services account.
	ConnectorLabel *string

	// The provisioning type of the connector. Currently the only supported value is
	// LAMBDA.
	ConnectorProvisioningConfig *types.ConnectorProvisioningConfig

	// The provisioning type of the connector. Currently the only supported value is
	// LAMBDA.
	ConnectorProvisioningType types.ConnectorProvisioningType

	// A description about the connector that's being registered.
	Description *string

	noSmithyDocumentSerde
}

type RegisterConnectorOutput struct {

	// The ARN of the connector being registered.
	ConnectorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterConnector"); err != nil {
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
	if err = addIdempotencyToken_opRegisterConnectorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterConnector(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRegisterConnector struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterConnector) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterConnector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterConnectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterConnectorInput ")
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
func addIdempotencyToken_opRegisterConnectorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterConnector{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterConnector",
	}
}
