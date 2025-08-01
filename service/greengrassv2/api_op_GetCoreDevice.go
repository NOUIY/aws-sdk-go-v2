// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves metadata for a Greengrass core device.
//
// IoT Greengrass relies on individual devices to send status updates to the
// Amazon Web Services Cloud. If the IoT Greengrass Core software isn't running on
// the device, or if device isn't connected to the Amazon Web Services Cloud, then
// the reported status of that device might not reflect its current status. The
// status timestamp indicates when the device status was last updated.
//
// Core devices send status updates at the following times:
//
//   - When the IoT Greengrass Core software starts
//
//   - When the core device receives a deployment from the Amazon Web Services
//     Cloud
//
//   - When the status of any component on the core device becomes BROKEN
//
//   - At a [regular interval that you can configure], which defaults to 24 hours
//
//   - For IoT Greengrass Core v2.7.0, the core device sends status updates upon
//     local deployment and cloud deployment
//
// [regular interval that you can configure]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html#greengrass-nucleus-component-configuration-fss
func (c *Client) GetCoreDevice(ctx context.Context, params *GetCoreDeviceInput, optFns ...func(*Options)) (*GetCoreDeviceOutput, error) {
	if params == nil {
		params = &GetCoreDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoreDevice", params, optFns, c.addOperationGetCoreDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoreDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoreDeviceInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	noSmithyDocumentSerde
}

type GetCoreDeviceOutput struct {

	// The computer architecture of the core device.
	Architecture *string

	// The name of the core device. This is also the name of the IoT thing.
	CoreDeviceThingName *string

	// The version of the IoT Greengrass Core software that the core device runs. This
	// version is equivalent to the version of the Greengrass nucleus component that
	// runs on the core device. For more information, see the [Greengrass nucleus component]in the IoT Greengrass V2
	// Developer Guide.
	//
	// [Greengrass nucleus component]: https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html
	CoreVersion *string

	// The time at which the core device's status last updated, expressed in ISO 8601
	// format.
	LastStatusUpdateTimestamp *time.Time

	// The operating system platform that the core device runs.
	Platform *string

	// The runtime for the core device. The runtime can be:
	//
	//   - aws_nucleus_classic
	//
	//   - aws_nucleus_lite
	Runtime *string

	// The status of the core device. The core device status can be:
	//
	//   - HEALTHY – The IoT Greengrass Core software and all components run on the
	//   core device without issue.
	//
	//   - UNHEALTHY – The IoT Greengrass Core software or a component is in a failed
	//   state on the core device.
	Status types.CoreDeviceStatus

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see [Tag your resources]in the IoT Greengrass V2 Developer Guide.
	//
	// [Tag your resources]: https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCoreDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCoreDevice"); err != nil {
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
	if err = addOpGetCoreDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoreDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCoreDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCoreDevice",
	}
}
