// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures event selectors (also referred to as basic event selectors) or
// advanced event selectors for your trail. You can use either
// AdvancedEventSelectors or EventSelectors , but not both. If you apply
// AdvancedEventSelectors to a trail, any existing EventSelectors are overwritten.
//
// You can use AdvancedEventSelectors to log management events, data events for
// all resource types, and network activity events.
//
// You can use EventSelectors to log management events and data events for the
// following resource types:
//
//   - AWS::DynamoDB::Table
//
//   - AWS::Lambda::Function
//
//   - AWS::S3::Object
//
// You can't use EventSelectors to log network activity events.
//
// If you want your trail to log Insights events, be sure the event selector or
// advanced event selector enables logging of the Insights event types you want
// configured for your trail. For more information about logging Insights events,
// see [Working with CloudTrail Insights]in the CloudTrail User Guide. By default, trails created without specific
// event selectors are configured to log all read and write management events, and
// no data events or network activity events.
//
// When an event occurs in your account, CloudTrail evaluates the event selectors
// or advanced event selectors in all trails. For each trail, if the event matches
// any event selector, the trail processes and logs the event. If the event doesn't
// match any event selector, the trail doesn't log the event.
//
// Example
//
//   - You create an event selector for a trail and specify that you want to log
//     write-only events.
//
//   - The EC2 GetConsoleOutput and RunInstances API operations occur in your
//     account.
//
//   - CloudTrail evaluates whether the events match your event selectors.
//
//   - The RunInstances is a write-only event and it matches your event selector.
//     The trail logs the event.
//
//   - The GetConsoleOutput is a read-only event that doesn't match your event
//     selector. The trail doesn't log the event.
//
// The PutEventSelectors operation must be called from the Region in which the
// trail was created; otherwise, an InvalidHomeRegionException exception is thrown.
//
// You can configure up to five event selectors for each trail.
//
// You can add advanced event selectors, and conditions for your advanced event
// selectors, up to a maximum of 500 values for all conditions and selectors on a
// trail. For more information, see [Logging management events], [Logging data events], [Logging network activity events], and [Quotas in CloudTrail] in the CloudTrail User Guide.
//
// [Logging network activity events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html
// [Logging management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
// [Quotas in CloudTrail]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html
// [Logging data events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
func (c *Client) PutEventSelectors(ctx context.Context, params *PutEventSelectorsInput, optFns ...func(*Options)) (*PutEventSelectorsOutput, error) {
	if params == nil {
		params = &PutEventSelectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEventSelectors", params, optFns, c.addOperationPutEventSelectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventSelectorsInput struct {

	// Specifies the name of the trail or trail ARN. If you specify a trail name, the
	// string must meet the following requirements:
	//
	//   - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.),
	//   underscores (_), or dashes (-)
	//
	//   - Start with a letter or number, and end with a letter or number
	//
	//   - Be between 3 and 128 characters
	//
	//   - Have no adjacent periods, underscores or dashes. Names like my-_namespace
	//   and my--namespace are not valid.
	//
	//   - Not be in IP address format (for example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must be in the following format.
	//
	//     arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	TrailName *string

	//  Specifies the settings for advanced event selectors. You can use advanced
	// event selectors to log management events, data events for all resource types,
	// and network activity events.
	//
	// You can add advanced event selectors, and conditions for your advanced event
	// selectors, up to a maximum of 500 values for all conditions and selectors on a
	// trail. You can use either AdvancedEventSelectors or EventSelectors , but not
	// both. If you apply AdvancedEventSelectors to a trail, any existing
	// EventSelectors are overwritten. For more information about advanced event
	// selectors, see [Logging data events]and [Logging network activity events] in the CloudTrail User Guide.
	//
	// [Logging network activity events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html
	// [Logging data events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
	AdvancedEventSelectors []types.AdvancedEventSelector

	// Specifies the settings for your event selectors. You can use event selectors to
	// log management events and data events for the following resource types:
	//
	//   - AWS::DynamoDB::Table
	//
	//   - AWS::Lambda::Function
	//
	//   - AWS::S3::Object
	//
	// You can't use event selectors to log network activity events.
	//
	// You can configure up to five event selectors for a trail. You can use either
	// EventSelectors or AdvancedEventSelectors in a PutEventSelectors request, but
	// not both. If you apply EventSelectors to a trail, any existing
	// AdvancedEventSelectors are overwritten.
	EventSelectors []types.EventSelector

	noSmithyDocumentSerde
}

type PutEventSelectorsOutput struct {

	// Specifies the advanced event selectors configured for your trail.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// Specifies the event selectors configured for your trail.
	EventSelectors []types.EventSelector

	// Specifies the ARN of the trail that was updated with event selectors. The
	// following is the format of a trail ARN.
	//
	//     arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEventSelectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEventSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEventSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutEventSelectors"); err != nil {
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
	if err = addOpPutEventSelectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventSelectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEventSelectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutEventSelectors",
	}
}
