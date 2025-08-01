// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Get details about a Node in the specified Cluster.
func (c *Client) DescribeNode(ctx context.Context, params *DescribeNodeInput, optFns ...func(*Options)) (*DescribeNodeOutput, error) {
	if params == nil {
		params = &DescribeNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNode", params, optFns, c.addOperationDescribeNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeNodeRequest
type DescribeNodeInput struct {

	// The ID of the cluster
	//
	// This member is required.
	ClusterId *string

	// The ID of the node.
	//
	// This member is required.
	NodeId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DescribeNodeResponse
type DescribeNodeOutput struct {

	// The ARN of the Node. It is automatically assigned when the Node is created.
	Arn *string

	// An array of IDs. Each ID is one ChannelPlacementGroup that is associated with
	// this Node. Empty if the Node is not yet associated with any groups.
	ChannelPlacementGroups []string

	// The ID of the Cluster that the Node belongs to.
	ClusterId *string

	// The current connection state of the Node.
	ConnectionState types.NodeConnectionState

	// The unique ID of the Node. Unique in the Cluster. The ID is the resource-id
	// portion of the ARN.
	Id *string

	// The ARN of the EC2 instance hosting the Node.
	InstanceArn *string

	// The name that you specified for the Node.
	Name *string

	// Documentation update needed
	NodeInterfaceMappings []types.NodeInterfaceMapping

	// The initial role current role of the Node in the Cluster. ACTIVE means the Node
	// is available for encoding. BACKUP means the Node is a redundant Node and might
	// get used if an ACTIVE Node fails.
	Role types.NodeRole

	// An array of SDI source mappings. Each mapping connects one logical SdiSource to
	// the physical SDI card and port that the physical SDI source uses.
	SdiSourceMappings []types.SdiSourceMapping

	// The current state of the Node.
	State types.NodeState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeNode"); err != nil {
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
	if err = addOpDescribeNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNode(options.Region), middleware.Before); err != nil {
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

// NodeDeregisteredWaiterOptions are waiter options for NodeDeregisteredWaiter
type NodeDeregisteredWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NodeDeregisteredWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NodeDeregisteredWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNodeInput, *DescribeNodeOutput, error) (bool, error)
}

// NodeDeregisteredWaiter defines the waiters for NodeDeregistered
type NodeDeregisteredWaiter struct {
	client DescribeNodeAPIClient

	options NodeDeregisteredWaiterOptions
}

// NewNodeDeregisteredWaiter constructs a NodeDeregisteredWaiter.
func NewNodeDeregisteredWaiter(client DescribeNodeAPIClient, optFns ...func(*NodeDeregisteredWaiterOptions)) *NodeDeregisteredWaiter {
	options := NodeDeregisteredWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = nodeDeregisteredStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NodeDeregisteredWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NodeDeregistered waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *NodeDeregisteredWaiter) Wait(ctx context.Context, params *DescribeNodeInput, maxWaitDur time.Duration, optFns ...func(*NodeDeregisteredWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NodeDeregistered waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *NodeDeregisteredWaiter) WaitForOutput(ctx context.Context, params *DescribeNodeInput, maxWaitDur time.Duration, optFns ...func(*NodeDeregisteredWaiterOptions)) (*DescribeNodeOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNode(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NodeDeregistered waiter")
}

func nodeDeregisteredStateRetryable(ctx context.Context, input *DescribeNodeInput, output *DescribeNodeOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "DEREGISTERED"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "DEREGISTERING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "DRAINING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// NodeRegisteredWaiterOptions are waiter options for NodeRegisteredWaiter
type NodeRegisteredWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NodeRegisteredWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, NodeRegisteredWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNodeInput, *DescribeNodeOutput, error) (bool, error)
}

// NodeRegisteredWaiter defines the waiters for NodeRegistered
type NodeRegisteredWaiter struct {
	client DescribeNodeAPIClient

	options NodeRegisteredWaiterOptions
}

// NewNodeRegisteredWaiter constructs a NodeRegisteredWaiter.
func NewNodeRegisteredWaiter(client DescribeNodeAPIClient, optFns ...func(*NodeRegisteredWaiterOptions)) *NodeRegisteredWaiter {
	options := NodeRegisteredWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = nodeRegisteredStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NodeRegisteredWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NodeRegistered waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *NodeRegisteredWaiter) Wait(ctx context.Context, params *DescribeNodeInput, maxWaitDur time.Duration, optFns ...func(*NodeRegisteredWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NodeRegistered waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *NodeRegisteredWaiter) WaitForOutput(ctx context.Context, params *DescribeNodeInput, maxWaitDur time.Duration, optFns ...func(*NodeRegisteredWaiterOptions)) (*DescribeNodeOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNode(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NodeRegistered waiter")
}

func nodeRegisteredStateRetryable(ctx context.Context, input *DescribeNodeInput, output *DescribeNodeOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "ACTIVE"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "REGISTERING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.NotFoundException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "REGISTRATION_FAILED"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// DescribeNodeAPIClient is a client that implements the DescribeNode operation.
type DescribeNodeAPIClient interface {
	DescribeNode(context.Context, *DescribeNodeInput, ...func(*Options)) (*DescribeNodeOutput, error)
}

var _ DescribeNodeAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeNode",
	}
}
