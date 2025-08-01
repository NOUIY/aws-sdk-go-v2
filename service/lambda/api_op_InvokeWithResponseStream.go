// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	"sync"
)

// Configure your Lambda functions to stream response payloads back to clients.
// For more information, see [Configuring a Lambda function to stream responses].
//
// This operation requires permission for the [lambda:InvokeFunction] action. For details on how to set
// up permissions for cross-account invocations, see [Granting function access to other accounts].
//
// [Configuring a Lambda function to stream responses]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html
// [lambda:InvokeFunction]: https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html
// [Granting function access to other accounts]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xaccountinvoke
func (c *Client) InvokeWithResponseStream(ctx context.Context, params *InvokeWithResponseStreamInput, optFns ...func(*Options)) (*InvokeWithResponseStreamOutput, error) {
	if params == nil {
		params = &InvokeWithResponseStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeWithResponseStream", params, optFns, c.addOperationInvokeWithResponseStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeWithResponseStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeWithResponseStreamInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – my-function .
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN – 123456789012:function:my-function .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Up to 3,583 bytes of base64-encoded data about the invoking client to pass to
	// the function in the context object.
	ClientContext *string

	// Use one of the following options:
	//
	//   - RequestResponse (default) – Invoke the function synchronously. Keep the
	//   connection open until the function returns a response or times out. The API
	//   operation response includes the function response and additional data.
	//
	//   - DryRun – Validate parameter values and verify that the IAM user or role has
	//   permission to invoke the function.
	InvocationType types.ResponseStreamingInvocationType

	// Set to Tail to include the execution log in the response. Applies to
	// synchronously invoked functions only.
	LogType types.LogType

	// The JSON that you want to provide to your Lambda function as input.
	//
	// You can enter the JSON directly. For example, --payload '{ "key": "value" }' .
	// You can also specify a file path. For example, --payload file://payload.json .
	Payload []byte

	// The alias name.
	Qualifier *string

	noSmithyDocumentSerde
}

type InvokeWithResponseStreamOutput struct {

	// The version of the function that executed. When you invoke a function with an
	// alias, this indicates which version the alias resolved to.
	ExecutedVersion *string

	// The type of data the stream is returning.
	ResponseStreamContentType *string

	// For a successful request, the HTTP status code is in the 200 range. For the
	// RequestResponse invocation type, this status code is 200. For the DryRun
	// invocation type, this status code is 204.
	StatusCode int32

	eventStream *InvokeWithResponseStreamEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *InvokeWithResponseStreamOutput) GetStream() *InvokeWithResponseStreamEventStream {
	return o.eventStream
}

func (c *Client) addOperationInvokeWithResponseStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeWithResponseStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeWithResponseStream"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamInvokeWithResponseStreamMiddleware(stack, options); err != nil {
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
	if err = addOpInvokeWithResponseStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeWithResponseStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeWithResponseStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeWithResponseStream",
	}
}

// InvokeWithResponseStreamEventStream provides the event stream handling for the InvokeWithResponseStream operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewInvokeWithResponseStreamEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type InvokeWithResponseStreamEventStream struct {
	// InvokeWithResponseStreamResponseEventReader is the EventStream reader for the
	// InvokeWithResponseStreamResponseEvent events. This value is automatically set by
	// the SDK when the API call is made Use this member when unit testing your code
	// with the SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader InvokeWithResponseStreamResponseEventReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewInvokeWithResponseStreamEventStream initializes an InvokeWithResponseStreamEventStream.
// This function should only be used for testing and mocking the InvokeWithResponseStreamEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewInvokeWithResponseStreamEventStream(optFns ...func(*InvokeWithResponseStreamEventStream)) *InvokeWithResponseStreamEventStream {
	es := &InvokeWithResponseStreamEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *InvokeWithResponseStreamEventStream) Events() <-chan types.InvokeWithResponseStreamResponseEvent {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *InvokeWithResponseStreamEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *InvokeWithResponseStreamEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *InvokeWithResponseStreamEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *InvokeWithResponseStreamEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
