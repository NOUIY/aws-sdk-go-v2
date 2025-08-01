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

// Gets details about a multiplex.
func (c *Client) DescribeMultiplex(ctx context.Context, params *DescribeMultiplexInput, optFns ...func(*Options)) (*DescribeMultiplexOutput, error) {
	if params == nil {
		params = &DescribeMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMultiplex", params, optFns, c.addOperationDescribeMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeMultiplexRequest
type DescribeMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DescribeMultiplexResponse
type DescribeMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []string

	// A list of the multiplex output destinations.
	Destinations []types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32

	// The number of programs in the multiplex.
	ProgramCount *int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMultiplex"); err != nil {
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
	if err = addOpDescribeMultiplexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMultiplex(options.Region), middleware.Before); err != nil {
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

// MultiplexCreatedWaiterOptions are waiter options for MultiplexCreatedWaiter
type MultiplexCreatedWaiterOptions struct {

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
	// MultiplexCreatedWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, MultiplexCreatedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexCreatedWaiter defines the waiters for MultiplexCreated
type MultiplexCreatedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexCreatedWaiterOptions
}

// NewMultiplexCreatedWaiter constructs a MultiplexCreatedWaiter.
func NewMultiplexCreatedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexCreatedWaiterOptions)) *MultiplexCreatedWaiter {
	options := MultiplexCreatedWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexCreated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexCreatedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for MultiplexCreated waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *MultiplexCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexCreatedWaiterOptions)) (*DescribeMultiplexOutput, error) {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for MultiplexCreated waiter")
}

func multiplexCreatedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "IDLE"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "CREATING"
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

	if err == nil {
		v1 := output.State
		expectedValue := "CREATE_FAILED"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// MultiplexDeletedWaiterOptions are waiter options for MultiplexDeletedWaiter
type MultiplexDeletedWaiterOptions struct {

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
	// MultiplexDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, MultiplexDeletedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexDeletedWaiter defines the waiters for MultiplexDeleted
type MultiplexDeletedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexDeletedWaiterOptions
}

// NewMultiplexDeletedWaiter constructs a MultiplexDeletedWaiter.
func NewMultiplexDeletedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexDeletedWaiterOptions)) *MultiplexDeletedWaiter {
	options := MultiplexDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexDeletedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for MultiplexDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *MultiplexDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexDeletedWaiterOptions)) (*DescribeMultiplexOutput, error) {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for MultiplexDeleted waiter")
}

func multiplexDeletedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "DELETED"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "DELETING"
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

// MultiplexRunningWaiterOptions are waiter options for MultiplexRunningWaiter
type MultiplexRunningWaiterOptions struct {

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
	// MultiplexRunningWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, MultiplexRunningWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexRunningWaiter defines the waiters for MultiplexRunning
type MultiplexRunningWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexRunningWaiterOptions
}

// NewMultiplexRunningWaiter constructs a MultiplexRunningWaiter.
func NewMultiplexRunningWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexRunningWaiterOptions)) *MultiplexRunningWaiter {
	options := MultiplexRunningWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexRunning waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexRunningWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexRunningWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for MultiplexRunning waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *MultiplexRunningWaiter) WaitForOutput(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexRunningWaiterOptions)) (*DescribeMultiplexOutput, error) {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for MultiplexRunning waiter")
}

func multiplexRunningStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "RUNNING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "STARTING"
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

// MultiplexStoppedWaiterOptions are waiter options for MultiplexStoppedWaiter
type MultiplexStoppedWaiterOptions struct {

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
	// MultiplexStoppedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, MultiplexStoppedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeMultiplexInput, *DescribeMultiplexOutput, error) (bool, error)
}

// MultiplexStoppedWaiter defines the waiters for MultiplexStopped
type MultiplexStoppedWaiter struct {
	client DescribeMultiplexAPIClient

	options MultiplexStoppedWaiterOptions
}

// NewMultiplexStoppedWaiter constructs a MultiplexStoppedWaiter.
func NewMultiplexStoppedWaiter(client DescribeMultiplexAPIClient, optFns ...func(*MultiplexStoppedWaiterOptions)) *MultiplexStoppedWaiter {
	options := MultiplexStoppedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = multiplexStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &MultiplexStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for MultiplexStopped waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MultiplexStoppedWaiter) Wait(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for MultiplexStopped waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *MultiplexStoppedWaiter) WaitForOutput(ctx context.Context, params *DescribeMultiplexInput, maxWaitDur time.Duration, optFns ...func(*MultiplexStoppedWaiterOptions)) (*DescribeMultiplexOutput, error) {
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

		out, err := w.client.DescribeMultiplex(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for MultiplexStopped waiter")
}

func multiplexStoppedStateRetryable(ctx context.Context, input *DescribeMultiplexInput, output *DescribeMultiplexOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.State
		expectedValue := "IDLE"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.State
		expectedValue := "STOPPING"
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

// DescribeMultiplexAPIClient is a client that implements the DescribeMultiplex
// operation.
type DescribeMultiplexAPIClient interface {
	DescribeMultiplex(context.Context, *DescribeMultiplexInput, ...func(*Options)) (*DescribeMultiplexOutput, error)
}

var _ DescribeMultiplexAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeMultiplex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMultiplex",
	}
}
