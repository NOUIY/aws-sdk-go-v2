// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	"sync"
)

// Invokes an alias of a flow to run the inputs that you specify and return the
// output of each node as a stream. If there's an error, the error is returned. For
// more information, see [Test a flow in Amazon Bedrock]in the [Amazon Bedrock User Guide].
//
// The CLI doesn't support streaming operations in Amazon Bedrock, including
// InvokeFlow .
//
// [Test a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-test.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func (c *Client) InvokeFlow(ctx context.Context, params *InvokeFlowInput, optFns ...func(*Options)) (*InvokeFlowOutput, error) {
	if params == nil {
		params = &InvokeFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeFlow", params, optFns, c.addOperationInvokeFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeFlowInput struct {

	// The unique identifier of the flow alias.
	//
	// This member is required.
	FlowAliasIdentifier *string

	// The unique identifier of the flow.
	//
	// This member is required.
	FlowIdentifier *string

	// A list of objects, each containing information about an input into the flow.
	//
	// This member is required.
	Inputs []types.FlowInput

	// Specifies whether to return the trace for the flow or not. Traces track inputs
	// and outputs for nodes in the flow. For more information, see [Track each step in your prompt flow by viewing its trace in Amazon Bedrock].
	//
	// [Track each step in your prompt flow by viewing its trace in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-trace.html
	EnableTrace *bool

	// The unique identifier for the current flow execution. If you don't provide a
	// value, Amazon Bedrock creates the identifier for you.
	ExecutionId *string

	// Model performance settings for the request.
	ModelPerformanceConfiguration *types.ModelPerformanceConfiguration

	noSmithyDocumentSerde
}

type InvokeFlowOutput struct {

	// The unique identifier for the current flow execution.
	ExecutionId *string

	eventStream *InvokeFlowEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *InvokeFlowOutput) GetStream() *InvokeFlowEventStream {
	return o.eventStream
}

func (c *Client) addOperationInvokeFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeFlow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeFlow"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamInvokeFlowMiddleware(stack, options); err != nil {
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
	if err = addOpInvokeFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeFlow",
	}
}

// InvokeFlowEventStream provides the event stream handling for the InvokeFlow operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewInvokeFlowEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type InvokeFlowEventStream struct {
	// FlowResponseStreamReader is the EventStream reader for the FlowResponseStream
	// events. This value is automatically set by the SDK when the API call is made Use
	// this member when unit testing your code with the SDK to mock out the EventStream
	// Reader.
	//
	// Must not be nil.
	Reader FlowResponseStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewInvokeFlowEventStream initializes an InvokeFlowEventStream.
// This function should only be used for testing and mocking the InvokeFlowEventStream
// stream within your application.
//
// The Reader member must be set before reading events from the stream.
func NewInvokeFlowEventStream(optFns ...func(*InvokeFlowEventStream)) *InvokeFlowEventStream {
	es := &InvokeFlowEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Events returns a channel to read events from.
func (es *InvokeFlowEventStream) Events() <-chan types.FlowResponseStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *InvokeFlowEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *InvokeFlowEventStream) safeClose() {
	close(es.done)

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *InvokeFlowEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *InvokeFlowEventStream) waitStreamClose() {
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
