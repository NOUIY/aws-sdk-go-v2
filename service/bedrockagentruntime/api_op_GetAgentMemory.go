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

// Gets the sessions stored in the memory of the agent.
func (c *Client) GetAgentMemory(ctx context.Context, params *GetAgentMemoryInput, optFns ...func(*Options)) (*GetAgentMemoryOutput, error) {
	if params == nil {
		params = &GetAgentMemoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAgentMemory", params, optFns, c.addOperationGetAgentMemoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAgentMemoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAgentMemoryInput struct {

	// The unique identifier of an alias of an agent.
	//
	// This member is required.
	AgentAliasId *string

	// The unique identifier of the agent to which the alias belongs.
	//
	// This member is required.
	AgentId *string

	// The unique identifier of the memory.
	//
	// This member is required.
	MemoryId *string

	// The type of memory.
	//
	// This member is required.
	MemoryType types.MemoryType

	// The maximum number of items to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxItems *int32

	// If the total number of results is greater than the maxItems value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetAgentMemoryOutput struct {

	// Contains details of the sessions stored in the memory
	MemoryContents []types.Memory

	// If the total number of results is greater than the maxItems value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAgentMemoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAgentMemory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAgentMemory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAgentMemory"); err != nil {
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
	if err = addOpGetAgentMemoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAgentMemory(options.Region), middleware.Before); err != nil {
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

// GetAgentMemoryPaginatorOptions is the paginator options for GetAgentMemory
type GetAgentMemoryPaginatorOptions struct {
	// The maximum number of items to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAgentMemoryPaginator is a paginator for GetAgentMemory
type GetAgentMemoryPaginator struct {
	options   GetAgentMemoryPaginatorOptions
	client    GetAgentMemoryAPIClient
	params    *GetAgentMemoryInput
	nextToken *string
	firstPage bool
}

// NewGetAgentMemoryPaginator returns a new GetAgentMemoryPaginator
func NewGetAgentMemoryPaginator(client GetAgentMemoryAPIClient, params *GetAgentMemoryInput, optFns ...func(*GetAgentMemoryPaginatorOptions)) *GetAgentMemoryPaginator {
	if params == nil {
		params = &GetAgentMemoryInput{}
	}

	options := GetAgentMemoryPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAgentMemoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAgentMemoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAgentMemory page.
func (p *GetAgentMemoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAgentMemoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetAgentMemory(ctx, &params, optFns...)
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

// GetAgentMemoryAPIClient is a client that implements the GetAgentMemory
// operation.
type GetAgentMemoryAPIClient interface {
	GetAgentMemory(context.Context, *GetAgentMemoryInput, ...func(*Options)) (*GetAgentMemoryOutput, error)
}

var _ GetAgentMemoryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetAgentMemory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAgentMemory",
	}
}
