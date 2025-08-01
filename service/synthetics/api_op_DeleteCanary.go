// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently deletes the specified canary.
//
// If the canary's ProvisionedResourceCleanup field is set to AUTOMATIC or you
// specify DeleteLambda in this operation as true , CloudWatch Synthetics also
// deletes the Lambda functions and layers that are used by the canary.
//
// Other resources used and created by the canary are not automatically deleted.
// After you delete a canary, you should also delete the following:
//
//   - The CloudWatch alarms created for this canary. These alarms have a name of
//     Synthetics-Alarm-first-198-characters-of-canary-name-canaryId-alarm number
//
//   - Amazon S3 objects and buckets, such as the canary's artifact location.
//
//   - IAM roles created for the canary. If they were created in the console,
//     these roles have the name
//     role/service-role/CloudWatchSyntheticsRole-First-21-Characters-of-CanaryName
//
//   - CloudWatch Logs log groups created for the canary. These logs groups have
//     the name /aws/lambda/cwsyn-First-21-Characters-of-CanaryName
//
// Before you delete a canary, you might want to use GetCanary to display the
// information about this canary. Make note of the information returned by this
// operation so that you can delete these resources after you delete the canary.
func (c *Client) DeleteCanary(ctx context.Context, params *DeleteCanaryInput, optFns ...func(*Options)) (*DeleteCanaryOutput, error) {
	if params == nil {
		params = &DeleteCanaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCanary", params, optFns, c.addOperationDeleteCanaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCanaryInput struct {

	// The name of the canary that you want to delete. To find the names of your
	// canaries, use [DescribeCanaries].
	//
	// [DescribeCanaries]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html
	//
	// This member is required.
	Name *string

	// Specifies whether to also delete the Lambda functions and layers used by this
	// canary. The default is false .
	//
	// Your setting for this parameter is used only if the canary doesn't have
	// AUTOMATIC for its ProvisionedResourceCleanup field. If that field is set to
	// AUTOMATIC , then the Lambda functions and layers will be deleted when this
	// canary is deleted.
	//
	// Type: Boolean
	DeleteLambda bool

	noSmithyDocumentSerde
}

type DeleteCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCanaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCanary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCanary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteCanary"); err != nil {
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
	if err = addOpDeleteCanaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCanary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCanary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteCanary",
	}
}
