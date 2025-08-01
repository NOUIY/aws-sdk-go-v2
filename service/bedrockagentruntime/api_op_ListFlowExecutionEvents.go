// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists events that occurred during a flow execution. Events provide detailed
// information about the execution progress, including node inputs and outputs,
// flow inputs and outputs, condition results, and failure events.
//
// Flow executions is in preview release for Amazon Bedrock and is subject to
// change.
func (c *Client) ListFlowExecutionEvents(ctx context.Context, params *ListFlowExecutionEventsInput, optFns ...func(*Options)) (*ListFlowExecutionEventsOutput, error) {
	if params == nil {
		params = &ListFlowExecutionEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFlowExecutionEvents", params, optFns, c.addOperationListFlowExecutionEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFlowExecutionEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowExecutionEventsInput struct {

	// The type of events to retrieve. Specify Node for node-level events or Flow for
	// flow-level events.
	//
	// This member is required.
	EventType types.FlowExecutionEventType

	// The unique identifier of the flow execution.
	//
	// This member is required.
	ExecutionIdentifier *string

	// The unique identifier of the flow alias used for the execution.
	//
	// This member is required.
	FlowAliasIdentifier *string

	// The unique identifier of the flow.
	//
	// This member is required.
	FlowIdentifier *string

	// The maximum number of events to return in a single response. If more events
	// exist than the specified maxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int32

	// A token to retrieve the next set of results. This value is returned in the
	// response if more results are available.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFlowExecutionEventsOutput struct {

	// A list of events that occurred during the flow execution. Events can include
	// node inputs and outputs, flow inputs and outputs, condition results, and failure
	// events.
	//
	// This member is required.
	FlowExecutionEvents []types.FlowExecutionEvent

	// A token to retrieve the next set of results. This value is returned if more
	// results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFlowExecutionEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFlowExecutionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFlowExecutionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFlowExecutionEvents"); err != nil {
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
	if err = addOpListFlowExecutionEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFlowExecutionEvents(options.Region), middleware.Before); err != nil {
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

// ListFlowExecutionEventsPaginatorOptions is the paginator options for
// ListFlowExecutionEvents
type ListFlowExecutionEventsPaginatorOptions struct {
	// The maximum number of events to return in a single response. If more events
	// exist than the specified maxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFlowExecutionEventsPaginator is a paginator for ListFlowExecutionEvents
type ListFlowExecutionEventsPaginator struct {
	options   ListFlowExecutionEventsPaginatorOptions
	client    ListFlowExecutionEventsAPIClient
	params    *ListFlowExecutionEventsInput
	nextToken *string
	firstPage bool
}

// NewListFlowExecutionEventsPaginator returns a new
// ListFlowExecutionEventsPaginator
func NewListFlowExecutionEventsPaginator(client ListFlowExecutionEventsAPIClient, params *ListFlowExecutionEventsInput, optFns ...func(*ListFlowExecutionEventsPaginatorOptions)) *ListFlowExecutionEventsPaginator {
	if params == nil {
		params = &ListFlowExecutionEventsInput{}
	}

	options := ListFlowExecutionEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFlowExecutionEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFlowExecutionEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFlowExecutionEvents page.
func (p *ListFlowExecutionEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFlowExecutionEventsOutput, error) {
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
	result, err := p.client.ListFlowExecutionEvents(ctx, &params, optFns...)
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

// ListFlowExecutionEventsAPIClient is a client that implements the
// ListFlowExecutionEvents operation.
type ListFlowExecutionEventsAPIClient interface {
	ListFlowExecutionEvents(context.Context, *ListFlowExecutionEventsInput, ...func(*Options)) (*ListFlowExecutionEventsOutput, error)
}

var _ ListFlowExecutionEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListFlowExecutionEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFlowExecutionEvents",
	}
}
