// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Schedules a service software update for an Amazon OpenSearch Service domain.
// For more information, see [Service software updates in Amazon OpenSearch Service].
//
// [Service software updates in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html
func (c *Client) StartServiceSoftwareUpdate(ctx context.Context, params *StartServiceSoftwareUpdateInput, optFns ...func(*Options)) (*StartServiceSoftwareUpdateOutput, error) {
	if params == nil {
		params = &StartServiceSoftwareUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartServiceSoftwareUpdate", params, optFns, c.addOperationStartServiceSoftwareUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartServiceSoftwareUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the StartServiceSoftwareUpdate
// operation.
type StartServiceSoftwareUpdateInput struct {

	// The name of the domain that you want to update to the latest service software.
	//
	// This member is required.
	DomainName *string

	// The Epoch timestamp when you want the service software update to start. You
	// only need to specify this parameter if you set ScheduleAt to TIMESTAMP .
	DesiredStartTime *int64

	// When to start the service software update.
	//
	//   - NOW - Immediately schedules the update to happen in the current hour if
	//   there's capacity available.
	//
	//   - TIMESTAMP - Lets you specify a custom date and time to apply the update. If
	//   you specify this value, you must also provide a value for DesiredStartTime .
	//
	//   - OFF_PEAK_WINDOW - Marks the update to be picked up during an upcoming
	//   off-peak window. There's no guarantee that the update will happen during the
	//   next immediate window. Depending on capacity, it might happen in subsequent
	//   days.
	//
	// Default: NOW if you don't specify a value for DesiredStartTime , and TIMESTAMP
	// if you do.
	ScheduleAt types.ScheduleAt

	noSmithyDocumentSerde
}

// Represents the output of a StartServiceSoftwareUpdate operation. Contains the
// status of the update.
type StartServiceSoftwareUpdateOutput struct {

	// The current status of the OpenSearch Service software update.
	ServiceSoftwareOptions *types.ServiceSoftwareOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartServiceSoftwareUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartServiceSoftwareUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartServiceSoftwareUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartServiceSoftwareUpdate"); err != nil {
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
	if err = addOpStartServiceSoftwareUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartServiceSoftwareUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartServiceSoftwareUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartServiceSoftwareUpdate",
	}
}
