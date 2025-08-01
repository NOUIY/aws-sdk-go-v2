// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the shard count of the specified stream to the specified number of
// shards. This API is only supported for the data streams with the provisioned
// capacity mode.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// Updating the shard count is an asynchronous operation. Upon receiving the
// request, Kinesis Data Streams returns immediately and sets the status of the
// stream to UPDATING . After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE . Depending on the size of the stream, the
// scaling action could take a few minutes to complete. You can continue to read
// and write data to your stream while its status is UPDATING .
//
// To update the shard count, Kinesis Data Streams performs splits or merges on
// individual shards. This can cause short-lived shards to be created, in addition
// to the final shards. These short-lived shards count towards your total shard
// limit for your account in the Region.
//
// When using this operation, we recommend that you specify a target shard count
// that is a multiple of 25% (25%, 50%, 75%, 100%). You can specify any target
// value within your shard limit. However, if you specify a target that isn't a
// multiple of 25%, the scaling action might take longer to complete.
//
// This operation has the following default limits. By default, you cannot do the
// following:
//
//   - Scale more than ten times per rolling 24-hour period per stream
//
//   - Scale up to more than double your current shard count for a stream
//
//   - Scale down below half your current shard count for a stream
//
//   - Scale up to more than 10000 shards in a stream
//
//   - Scale a stream with more than 10000 shards down unless the result is less
//     than 10000 shards
//
//   - Scale up to more than the shard limit for your account
//
//   - Make over 10 TPS. TPS over 10 will trigger the LimitExceededException
//
// For the default limits for an Amazon Web Services account, see [Streams Limits] in the Amazon
// Kinesis Data Streams Developer Guide. To request an increase in the call rate
// limit, the shard limit for this API, or your overall shard limit, use the [limits form].
//
// [limits form]: https://console.aws.amazon.com/support/v1#/case/create?issueType=service-limit-increase&limitType=service-code-kinesis
// [Streams Limits]: https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html
func (c *Client) UpdateShardCount(ctx context.Context, params *UpdateShardCountInput, optFns ...func(*Options)) (*UpdateShardCountOutput, error) {
	if params == nil {
		params = &UpdateShardCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateShardCount", params, optFns, c.addOperationUpdateShardCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateShardCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateShardCountInput struct {

	// The scaling type. Uniform scaling creates shards of equal size.
	//
	// This member is required.
	ScalingType types.ScalingType

	// The new number of shards. This value has the following default limits. By
	// default, you cannot do the following:
	//
	//   - Set this value to more than double your current shard count for a stream.
	//
	//   - Set this value below half your current shard count for a stream.
	//
	//   - Set this value to more than 10000 shards in a stream (the default limit for
	//   shard count per stream is 10000 per account per region), unless you request a
	//   limit increase.
	//
	//   - Scale a stream with more than 10000 shards down unless you set this value
	//   to less than 10000 shards.
	//
	// This member is required.
	TargetShardCount *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *UpdateShardCountInput) bindEndpointParams(p *EndpointParameters) {

	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

type UpdateShardCountOutput struct {

	// The current number of shards.
	CurrentShardCount *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream.
	StreamName *string

	// The updated number of shards.
	TargetShardCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateShardCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateShardCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateShardCount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateShardCount"); err != nil {
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
	if err = addOpUpdateShardCountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateShardCount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateShardCount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateShardCount",
	}
}
