// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies one existing metric definition for CloudWatch RUM extended metrics.
// For more information about extended metrics, see [BatchCreateRumMetricsDefinitions].
//
// [BatchCreateRumMetricsDefinitions]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricsDefinitions.html
func (c *Client) UpdateRumMetricDefinition(ctx context.Context, params *UpdateRumMetricDefinitionInput, optFns ...func(*Options)) (*UpdateRumMetricDefinitionOutput, error) {
	if params == nil {
		params = &UpdateRumMetricDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRumMetricDefinition", params, optFns, c.addOperationUpdateRumMetricDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRumMetricDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRumMetricDefinitionInput struct {

	// The name of the CloudWatch RUM app monitor that sends these metrics.
	//
	// This member is required.
	AppMonitorName *string

	// The destination to send the metrics to. Valid values are CloudWatch and
	// Evidently . If you specify Evidently , you must also specify the ARN of the
	// CloudWatchEvidently experiment that will receive the metrics and an IAM role
	// that has permission to write to the experiment.
	//
	// This member is required.
	Destination types.MetricDestination

	// A structure that contains the new definition that you want to use for this
	// metric.
	//
	// This member is required.
	MetricDefinition *types.MetricDefinitionRequest

	// The ID of the metric definition to update.
	//
	// This member is required.
	MetricDefinitionId *string

	// This parameter is required if Destination is Evidently . If Destination is
	// CloudWatch , do not use this parameter.
	//
	// This parameter specifies the ARN of the Evidently experiment that is to receive
	// the metrics. You must have already defined this experiment as a valid
	// destination. For more information, see [PutRumMetricsDestination].
	//
	// [PutRumMetricsDestination]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html
	DestinationArn *string

	noSmithyDocumentSerde
}

type UpdateRumMetricDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRumMetricDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRumMetricDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRumMetricDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRumMetricDefinition"); err != nil {
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
	if err = addOpUpdateRumMetricDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRumMetricDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRumMetricDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRumMetricDefinition",
	}
}
