// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// To deregister a consumer, provide its ARN. Alternatively, you can provide the
// ARN of the data stream and the name you gave the consumer when you registered
// it. You may also provide all three parameters, as long as they don't conflict
// with each other. If you don't know the name or ARN of the consumer that you want
// to deregister, you can use the ListStreamConsumersoperation to get a list of the descriptions of
// all the consumers that are currently registered with a given data stream. The
// description of a consumer contains its name and ARN.
//
// This operation has a limit of five transactions per second per stream.
func (c *Client) DeregisterStreamConsumer(ctx context.Context, params *DeregisterStreamConsumerInput, optFns ...func(*Options)) (*DeregisterStreamConsumerOutput, error) {
	if params == nil {
		params = &DeregisterStreamConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterStreamConsumer", params, optFns, c.addOperationDeregisterStreamConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterStreamConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterStreamConsumerInput struct {

	// The ARN returned by Kinesis Data Streams when you registered the consumer. If
	// you don't know the ARN of the consumer that you want to deregister, you can use
	// the ListStreamConsumers operation to get a list of the descriptions of all the
	// consumers that are currently registered with a given data stream. The
	// description of a consumer contains its ARN.
	ConsumerARN *string

	// The name that you gave to the consumer.
	ConsumerName *string

	// The ARN of the Kinesis data stream that the consumer is registered with. For
	// more information, see [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces].
	//
	// [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams
	StreamARN *string

	noSmithyDocumentSerde
}

func (in *DeregisterStreamConsumerInput) bindEndpointParams(p *EndpointParameters) {

	p.StreamARN = in.StreamARN
	p.ConsumerARN = in.ConsumerARN
	p.OperationType = ptr.String("control")
}

type DeregisterStreamConsumerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterStreamConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterStreamConsumer"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterStreamConsumer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterStreamConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterStreamConsumer",
	}
}
