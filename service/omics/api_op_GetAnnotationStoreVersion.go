// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves the metadata for an annotation store version.
func (c *Client) GetAnnotationStoreVersion(ctx context.Context, params *GetAnnotationStoreVersionInput, optFns ...func(*Options)) (*GetAnnotationStoreVersionOutput, error) {
	if params == nil {
		params = &GetAnnotationStoreVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnnotationStoreVersion", params, optFns, c.addOperationGetAnnotationStoreVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnnotationStoreVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnnotationStoreVersionInput struct {

	//  The name given to an annotation store version to distinguish it from others.
	//
	// This member is required.
	Name *string

	//  The name given to an annotation store version to distinguish it from others.
	//
	// This member is required.
	VersionName *string

	noSmithyDocumentSerde
}

type GetAnnotationStoreVersionOutput struct {

	//  The time stamp for when an annotation store version was created.
	//
	// This member is required.
	CreationTime *time.Time

	//  The description for an annotation store version.
	//
	// This member is required.
	Description *string

	//  The annotation store version ID.
	//
	// This member is required.
	Id *string

	//  The name of the annotation store.
	//
	// This member is required.
	Name *string

	//  The status of an annotation store version.
	//
	// This member is required.
	Status types.VersionStatus

	//  The status of an annotation store version.
	//
	// This member is required.
	StatusMessage *string

	//  The store ID for annotation store version.
	//
	// This member is required.
	StoreId *string

	//  Any tags associated with an annotation store version.
	//
	// This member is required.
	Tags map[string]string

	//  The time stamp for when an annotation store version was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	//  The Arn for the annotation store.
	//
	// This member is required.
	VersionArn *string

	//  The name given to an annotation store version to distinguish it from others.
	//
	// This member is required.
	VersionName *string

	//  The size of the annotation store version in Bytes.
	//
	// This member is required.
	VersionSizeBytes *int64

	//  The options for an annotation store version.
	VersionOptions types.VersionOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnnotationStoreVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAnnotationStoreVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAnnotationStoreVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAnnotationStoreVersion"); err != nil {
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
	if err = addEndpointPrefix_opGetAnnotationStoreVersionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAnnotationStoreVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnnotationStoreVersion(options.Region), middleware.Before); err != nil {
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

// AnnotationStoreVersionCreatedWaiterOptions are waiter options for
// AnnotationStoreVersionCreatedWaiter
type AnnotationStoreVersionCreatedWaiterOptions struct {

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
	// AnnotationStoreVersionCreatedWaiter will use default minimum delay of 30
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AnnotationStoreVersionCreatedWaiter will use default max delay of
	// 600 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
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
	Retryable func(context.Context, *GetAnnotationStoreVersionInput, *GetAnnotationStoreVersionOutput, error) (bool, error)
}

// AnnotationStoreVersionCreatedWaiter defines the waiters for
// AnnotationStoreVersionCreated
type AnnotationStoreVersionCreatedWaiter struct {
	client GetAnnotationStoreVersionAPIClient

	options AnnotationStoreVersionCreatedWaiterOptions
}

// NewAnnotationStoreVersionCreatedWaiter constructs a
// AnnotationStoreVersionCreatedWaiter.
func NewAnnotationStoreVersionCreatedWaiter(client GetAnnotationStoreVersionAPIClient, optFns ...func(*AnnotationStoreVersionCreatedWaiterOptions)) *AnnotationStoreVersionCreatedWaiter {
	options := AnnotationStoreVersionCreatedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = annotationStoreVersionCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnnotationStoreVersionCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnnotationStoreVersionCreated waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *AnnotationStoreVersionCreatedWaiter) Wait(ctx context.Context, params *GetAnnotationStoreVersionInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreVersionCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnnotationStoreVersionCreated
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *AnnotationStoreVersionCreatedWaiter) WaitForOutput(ctx context.Context, params *GetAnnotationStoreVersionInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreVersionCreatedWaiterOptions)) (*GetAnnotationStoreVersionOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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

		out, err := w.client.GetAnnotationStoreVersion(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnnotationStoreVersionCreated waiter")
}

func annotationStoreVersionCreatedStateRetryable(ctx context.Context, input *GetAnnotationStoreVersionInput, output *GetAnnotationStoreVersionOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.Status
		expectedValue := "ACTIVE"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.Status
		expectedValue := "CREATING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		v1 := output.Status
		expectedValue := "UPDATING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		v1 := output.Status
		expectedValue := "FAILED"
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

// AnnotationStoreVersionDeletedWaiterOptions are waiter options for
// AnnotationStoreVersionDeletedWaiter
type AnnotationStoreVersionDeletedWaiterOptions struct {

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
	// AnnotationStoreVersionDeletedWaiter will use default minimum delay of 30
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AnnotationStoreVersionDeletedWaiter will use default max delay of
	// 600 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
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
	Retryable func(context.Context, *GetAnnotationStoreVersionInput, *GetAnnotationStoreVersionOutput, error) (bool, error)
}

// AnnotationStoreVersionDeletedWaiter defines the waiters for
// AnnotationStoreVersionDeleted
type AnnotationStoreVersionDeletedWaiter struct {
	client GetAnnotationStoreVersionAPIClient

	options AnnotationStoreVersionDeletedWaiterOptions
}

// NewAnnotationStoreVersionDeletedWaiter constructs a
// AnnotationStoreVersionDeletedWaiter.
func NewAnnotationStoreVersionDeletedWaiter(client GetAnnotationStoreVersionAPIClient, optFns ...func(*AnnotationStoreVersionDeletedWaiterOptions)) *AnnotationStoreVersionDeletedWaiter {
	options := AnnotationStoreVersionDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = annotationStoreVersionDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnnotationStoreVersionDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnnotationStoreVersionDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *AnnotationStoreVersionDeletedWaiter) Wait(ctx context.Context, params *GetAnnotationStoreVersionInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreVersionDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnnotationStoreVersionDeleted
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *AnnotationStoreVersionDeletedWaiter) WaitForOutput(ctx context.Context, params *GetAnnotationStoreVersionInput, maxWaitDur time.Duration, optFns ...func(*AnnotationStoreVersionDeletedWaiterOptions)) (*GetAnnotationStoreVersionOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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

		out, err := w.client.GetAnnotationStoreVersion(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnnotationStoreVersionDeleted waiter")
}

func annotationStoreVersionDeletedStateRetryable(ctx context.Context, input *GetAnnotationStoreVersionInput, output *GetAnnotationStoreVersionOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.Status
		expectedValue := "DELETED"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.Status
		expectedValue := "DELETING"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

type endpointPrefix_opGetAnnotationStoreVersionMiddleware struct {
}

func (*endpointPrefix_opGetAnnotationStoreVersionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAnnotationStoreVersionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAnnotationStoreVersionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAnnotationStoreVersionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetAnnotationStoreVersionAPIClient is a client that implements the
// GetAnnotationStoreVersion operation.
type GetAnnotationStoreVersionAPIClient interface {
	GetAnnotationStoreVersion(context.Context, *GetAnnotationStoreVersionInput, ...func(*Options)) (*GetAnnotationStoreVersionOutput, error)
}

var _ GetAnnotationStoreVersionAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetAnnotationStoreVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAnnotationStoreVersion",
	}
}
