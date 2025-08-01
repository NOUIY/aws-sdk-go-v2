// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploys the EndpointConfig specified in the request to a new fleet of
// instances. SageMaker shifts endpoint traffic to the new instances with the
// updated endpoint configuration and then deletes the old instances using the
// previous EndpointConfig (there is no availability loss). For more information
// about how to control the update and traffic shifting process, see [Update models in production].
//
// When SageMaker receives the request, it sets the endpoint status to Updating .
// After updating the endpoint, it sets the status to InService . To check the
// status of an endpoint, use the [DescribeEndpoint]API.
//
// You must not delete an EndpointConfig in use by an endpoint that is live or
// while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. To update an endpoint, you must create a new EndpointConfig .
//
// If you delete the EndpointConfig of an endpoint that is active or being created
// or updated you may lose visibility into the instance type the endpoint is using.
// The endpoint must be deleted in order to stop incurring charges.
//
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
// [Update models in production]: https://docs.aws.amazon.com/sagemaker/latest/dg/deployment-guardrails.html
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	if params == nil {
		params = &UpdateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpoint", params, optFns, c.addOperationUpdateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointInput struct {

	// The name of the new endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	// The name of the endpoint whose configuration you want to update.
	//
	// This member is required.
	EndpointName *string

	// The deployment configuration for an endpoint, which contains the desired
	// deployment strategy and rollback configurations.
	DeploymentConfig *types.DeploymentConfig

	// When you are updating endpoint resources with RetainAllVariantProperties , whose
	// value is set to true , ExcludeRetainedVariantProperties specifies the list of
	// type [VariantProperty]to override with the values provided by EndpointConfig . If you don't
	// specify a value for ExcludeRetainedVariantProperties , no variant properties are
	// overridden.
	//
	// [VariantProperty]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html
	ExcludeRetainedVariantProperties []types.VariantProperty

	// When updating endpoint resources, enables or disables the retention of [variant properties], such
	// as the instance count or the variant weight. To retain the variant properties of
	// an endpoint when updating it, set RetainAllVariantProperties to true . To use
	// the variant properties specified in a new EndpointConfig call when updating an
	// endpoint, set RetainAllVariantProperties to false . The default is false .
	//
	// [variant properties]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html
	RetainAllVariantProperties *bool

	// Specifies whether to reuse the last deployment configuration. The default value
	// is false (the configuration is not reused).
	RetainDeploymentConfig *bool

	noSmithyDocumentSerde
}

type UpdateEndpointOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEndpoint"); err != nil {
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
	if err = addOpUpdateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEndpoint",
	}
}
