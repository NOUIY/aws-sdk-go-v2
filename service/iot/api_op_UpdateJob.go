// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates supported fields of the specified job.
//
// Requires permission to access the [UpdateJob] action.
//
// [UpdateJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) {
	if params == nil {
		params = &UpdateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJob", params, optFns, c.addOperationUpdateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobInput struct {

	// The ID of the job to be updated.
	//
	// This member is required.
	JobId *string

	// Allows you to create criteria to abort a job.
	AbortConfig *types.AbortConfig

	// A short text description of the job.
	Description *string

	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig *types.JobExecutionsRetryConfig

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *types.JobExecutionsRolloutConfig

	// The namespace used to indicate that a job is a customer-managed job.
	//
	// When you specify a value for this parameter, Amazon Web Services IoT Core sends
	// jobs notifications to MQTT topics that contain the value in the following
	// format.
	//
	//     $aws/things/THING_NAME/jobs/JOB_ID/notify-namespace-NAMESPACE_ID/
	//
	// The namespaceId feature is only supported by IoT Greengrass at this time. For
	// more information, see [Setting up IoT Greengrass core devices.]
	//
	// [Setting up IoT Greengrass core devices.]: https://docs.aws.amazon.com/greengrass/v2/developerguide/setting-up.html
	NamespaceId *string

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *types.PresignedUrlConfig

	// Specifies the amount of time each device has to finish its execution of the
	// job. The timer is started when the job execution status is set to IN_PROGRESS .
	// If the job execution status is not set to another terminal state before the time
	// expires, it will be automatically set to TIMED_OUT .
	TimeoutConfig *types.TimeoutConfig

	noSmithyDocumentSerde
}

type UpdateJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateJob"); err != nil {
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
	if err = addOpUpdateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateJob",
	}
}
