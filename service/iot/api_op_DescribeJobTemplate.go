// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a job template.
func (c *Client) DescribeJobTemplate(ctx context.Context, params *DescribeJobTemplateInput, optFns ...func(*Options)) (*DescribeJobTemplateOutput, error) {
	if params == nil {
		params = &DescribeJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobTemplate", params, optFns, c.addOperationDescribeJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobTemplateInput struct {

	// The unique identifier of the job template.
	//
	// This member is required.
	JobTemplateId *string

	noSmithyDocumentSerde
}

type DescribeJobTemplateOutput struct {

	// The criteria that determine when and how a job abort takes place.
	AbortConfig *types.AbortConfig

	// The time, in seconds since the epoch, when the job template was created.
	CreatedAt *time.Time

	// A description of the job template.
	Description *string

	// The package version Amazon Resource Names (ARNs) that are installed on the
	// device when the job successfully completes. The package version must be in
	// either the Published or Deprecated state when the job deploys. For more
	// information, see [Package version lifecycle].
	//
	// Note:The following Length Constraints relates to a single ARN. Up to 25 package
	// version ARNs are allowed.
	//
	// [Package version lifecycle]: https://docs.aws.amazon.com/iot/latest/developerguide/preparing-to-use-software-package-catalog.html#package-version-lifecycle
	DestinationPackageVersions []string

	// The job document.
	Document *string

	// An S3 link to the job document.
	DocumentSource *string

	// The configuration that determines how many retries are allowed for each failure
	// type for a job.
	JobExecutionsRetryConfig *types.JobExecutionsRetryConfig

	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig *types.JobExecutionsRolloutConfig

	// The ARN of the job template.
	JobTemplateArn *string

	// The unique identifier of the job template.
	JobTemplateId *string

	// Allows you to configure an optional maintenance window for the rollout of a job
	// document to all devices in the target group for a job.
	MaintenanceWindows []types.MaintenanceWindow

	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig *types.PresignedUrlConfig

	// Specifies the amount of time each device has to finish its execution of the
	// job. A timer is started when the job execution status is set to IN_PROGRESS . If
	// the job execution status is not set to another terminal state before the timer
	// expires, it will be automatically set to TIMED_OUT .
	TimeoutConfig *types.TimeoutConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeJobTemplate"); err != nil {
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
	if err = addOpDescribeJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeJobTemplate",
	}
}
