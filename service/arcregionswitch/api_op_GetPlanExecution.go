// Code generated by smithy-go-codegen DO NOT EDIT.

package arcregionswitch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arcregionswitch/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves detailed information about a specific plan execution. You must
// specify the plan ARN and execution ID.
func (c *Client) GetPlanExecution(ctx context.Context, params *GetPlanExecutionInput, optFns ...func(*Options)) (*GetPlanExecutionOutput, error) {
	if params == nil {
		params = &GetPlanExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlanExecution", params, optFns, c.addOperationGetPlanExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlanExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlanExecutionInput struct {

	// The execution identifier of a plan execution.
	//
	// This member is required.
	ExecutionId *string

	// The Amazon Resource Name (ARN) of the plan with the execution to retrieve.
	//
	// This member is required.
	PlanArn *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a nextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's nextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetPlanExecutionOutput struct {

	// The plan execution action. Valid values are Activate , to activate an Amazon Web
	// Services Region, or Deactivate , to deactivate a Region.
	//
	// This member is required.
	ExecutionAction types.ExecutionAction

	// The execution identifier of a plan execution.
	//
	// This member is required.
	ExecutionId *string

	// The Amazon Web Services Region for a plan execution.
	//
	// This member is required.
	ExecutionRegion *string

	// The plan execution state. Provides the state of a plan execution, for example,
	// In Progress or Paused by Operator.
	//
	// This member is required.
	ExecutionState types.ExecutionState

	// The plan execution mode. Valid values are Practice , for testing without making
	// actual changes, or Recovery , for actual traffic shifting and application
	// recovery.
	//
	// This member is required.
	Mode types.ExecutionMode

	// The Amazon Resource Name (ARN) of the plan.
	//
	// This member is required.
	PlanArn *string

	// The time (UTC) when the plan execution started.
	//
	// This member is required.
	StartTime *time.Time

	// The actual recovery time that Region switch calculates for a plan execution.
	// Actual recovery time includes the time for the plan to run added to the time
	// elapsed until the application health alarms that you've specified are healthy
	// again.
	ActualRecoveryTime *string

	// A comment included on the plan execution.
	Comment *string

	// The time (UTC) when the plan execution ended.
	EndTime *time.Time

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a nextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's nextToken response to request the next page of results.
	NextToken *string

	// The details of the Region switch plan.
	Plan *types.Plan

	// The states of the steps in the plan execution.
	StepStates []types.StepState

	// The timestamp when the plan execution was last updated.
	UpdatedAt *time.Time

	// The version for the plan.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPlanExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&smithyRpcv2cbor_serializeOpGetPlanExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&smithyRpcv2cbor_deserializeOpGetPlanExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPlanExecution"); err != nil {
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
	if err = addUserAgentFeatureProtocolRPCV2CBOR(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpGetPlanExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlanExecution(options.Region), middleware.Before); err != nil {
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

// PlanExecutionCompletedWaiterOptions are waiter options for
// PlanExecutionCompletedWaiter
type PlanExecutionCompletedWaiterOptions struct {

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
	// PlanExecutionCompletedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, PlanExecutionCompletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *GetPlanExecutionInput, *GetPlanExecutionOutput, error) (bool, error)
}

// PlanExecutionCompletedWaiter defines the waiters for PlanExecutionCompleted
type PlanExecutionCompletedWaiter struct {
	client GetPlanExecutionAPIClient

	options PlanExecutionCompletedWaiterOptions
}

// NewPlanExecutionCompletedWaiter constructs a PlanExecutionCompletedWaiter.
func NewPlanExecutionCompletedWaiter(client GetPlanExecutionAPIClient, optFns ...func(*PlanExecutionCompletedWaiterOptions)) *PlanExecutionCompletedWaiter {
	options := PlanExecutionCompletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = planExecutionCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PlanExecutionCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PlanExecutionCompleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *PlanExecutionCompletedWaiter) Wait(ctx context.Context, params *GetPlanExecutionInput, maxWaitDur time.Duration, optFns ...func(*PlanExecutionCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PlanExecutionCompleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *PlanExecutionCompletedWaiter) WaitForOutput(ctx context.Context, params *GetPlanExecutionInput, maxWaitDur time.Duration, optFns ...func(*PlanExecutionCompletedWaiterOptions)) (*GetPlanExecutionOutput, error) {
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

		out, err := w.client.GetPlanExecution(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PlanExecutionCompleted waiter")
}

func planExecutionCompletedStateRetryable(ctx context.Context, input *GetPlanExecutionInput, output *GetPlanExecutionOutput, err error) (bool, error) {

	if err == nil {
		v1 := output.ExecutionState
		expectedValue := "completed"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.ExecutionState
		expectedValue := "completedWithExceptions"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		v1 := output.ExecutionState
		expectedValue := "failed"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		v1 := output.ExecutionState
		expectedValue := "canceled"
		var pathValue string
		pathValue = string(v1)
		if pathValue == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		v1 := output.ExecutionState
		expectedValue := "planExecutionTimedOut"
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

// GetPlanExecutionPaginatorOptions is the paginator options for GetPlanExecution
type GetPlanExecutionPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPlanExecutionPaginator is a paginator for GetPlanExecution
type GetPlanExecutionPaginator struct {
	options   GetPlanExecutionPaginatorOptions
	client    GetPlanExecutionAPIClient
	params    *GetPlanExecutionInput
	nextToken *string
	firstPage bool
}

// NewGetPlanExecutionPaginator returns a new GetPlanExecutionPaginator
func NewGetPlanExecutionPaginator(client GetPlanExecutionAPIClient, params *GetPlanExecutionInput, optFns ...func(*GetPlanExecutionPaginatorOptions)) *GetPlanExecutionPaginator {
	if params == nil {
		params = &GetPlanExecutionInput{}
	}

	options := GetPlanExecutionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPlanExecutionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPlanExecutionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPlanExecution page.
func (p *GetPlanExecutionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPlanExecutionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetPlanExecution(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetPlanExecutionAPIClient is a client that implements the GetPlanExecution
// operation.
type GetPlanExecutionAPIClient interface {
	GetPlanExecution(context.Context, *GetPlanExecutionInput, ...func(*Options)) (*GetPlanExecutionOutput, error)
}

var _ GetPlanExecutionAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetPlanExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPlanExecution",
	}
}
