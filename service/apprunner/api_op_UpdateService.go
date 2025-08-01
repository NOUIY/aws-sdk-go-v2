// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an App Runner service. You can update the source configuration and
// instance configuration of the service. You can also update the ARN of the auto
// scaling configuration resource that's associated with the service. However, you
// can't change the name or the encryption configuration of the service. These can
// be set only when you create the service.
//
// To update the tags applied to your service, use the separate actions TagResource and UntagResource.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
func (c *Client) UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) {
	if params == nil {
		params = &UpdateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateService", params, optFns, c.addOperationUpdateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want to
	// update.
	//
	// This member is required.
	ServiceArn *string

	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration
	// resource that you want to associate with the App Runner service.
	AutoScalingConfigurationArn *string

	// The settings for the health check that App Runner performs to monitor the
	// health of the App Runner service.
	HealthCheckConfiguration *types.HealthCheckConfiguration

	// The runtime configuration to apply to instances (scaling units) of your service.
	InstanceConfiguration *types.InstanceConfiguration

	// Configuration settings related to network traffic of the web application that
	// the App Runner service runs.
	NetworkConfiguration *types.NetworkConfiguration

	// The observability configuration of your service.
	ObservabilityConfiguration *types.ServiceObservabilityConfiguration

	// The source configuration to apply to the App Runner service.
	//
	// You can change the configuration of the code or image repository that the
	// service uses. However, you can't switch from code to image or the other way
	// around. This means that you must provide the same structure member of
	// SourceConfiguration that you originally included when you created the service.
	// Specifically, you can include either CodeRepository or ImageRepository . To
	// update the source configuration, set the values to members of the structure that
	// you include.
	SourceConfiguration *types.SourceConfiguration

	noSmithyDocumentSerde
}

type UpdateServiceOutput struct {

	// The unique ID of the asynchronous operation that this request started. You can
	// use it combined with the ListOperationscall to track the operation's progress.
	//
	// This member is required.
	OperationId *string

	// A description of the App Runner service updated by this request. All
	// configuration values in the returned Service structure reflect configuration
	// changes that are being applied by this request.
	//
	// This member is required.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateService"); err != nil {
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
	if err = addOpUpdateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateService",
	}
}
