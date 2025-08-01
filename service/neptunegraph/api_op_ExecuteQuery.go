// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/document"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"strings"
)

// Execute an openCypher query.
//
// When invoking this operation in a Neptune Analytics cluster, the IAM user or
// role making the request must have a policy attached that allows one of the
// following IAM actions in that cluster, depending on the query:
//
//   - neptune-graph:ReadDataViaQuery
//
//   - neptune-graph:WriteDataViaQuery
//
//   - neptune-graph:DeleteDataViaQuery
func (c *Client) ExecuteQuery(ctx context.Context, params *ExecuteQueryInput, optFns ...func(*Options)) (*ExecuteQueryOutput, error) {
	if params == nil {
		params = &ExecuteQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteQuery", params, optFns, c.addOperationExecuteQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteQueryInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// The query language the query is written in. Currently only openCypher is
	// supported.
	//
	// This member is required.
	Language types.QueryLanguage

	// The query string to be executed.
	//
	// This member is required.
	QueryString *string

	// The explain mode parameter returns a query explain instead of the actual query
	// results. A query explain can be used to gather insights about the query
	// execution such as planning decisions, time spent on each operator, solutions
	// flowing etc.
	ExplainMode types.ExplainMode

	// The data parameters the query can use in JSON format. For example: {"name":
	// "john", "age": 20}. (optional)
	Parameters map[string]document.Interface

	// Query plan cache is a feature that saves the query plan and reuses it on
	// successive executions of the same query. This reduces query latency, and works
	// for both READ and UPDATE queries. The plan cache is an LRU cache with a 5
	// minute TTL and a capacity of 1000.
	PlanCache types.PlanCacheType

	// Specifies the query timeout duration, in milliseconds. (optional)
	QueryTimeoutMilliseconds *int32

	noSmithyDocumentSerde
}

func (in *ExecuteQueryInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("DataPlane")
}

type ExecuteQueryOutput struct {

	// The query results.
	//
	// This member is required.
	Payload io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExecuteQuery"); err != nil {
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
	if err = addEndpointPrefix_opExecuteQueryMiddleware(stack); err != nil {
		return err
	}
	if err = addOpExecuteQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteQuery(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opExecuteQueryMiddleware struct {
}

func (*endpointPrefix_opExecuteQueryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opExecuteQueryMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ExecuteQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.GraphIdentifier == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("GraphIdentifier forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.GraphIdentifier) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("GraphIdentifier forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.GraphIdentifier)}
	} else {
		prefix.WriteString(*input.GraphIdentifier)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opExecuteQueryMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opExecuteQueryMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opExecuteQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExecuteQuery",
	}
}
