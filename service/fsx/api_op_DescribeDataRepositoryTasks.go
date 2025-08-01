// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository tasks, if one or more TaskIds values are provided in the
// request, or if filters are used in the request. You can use filters to narrow
// the response to include just tasks for specific file systems or caches, or tasks
// in a specific lifecycle state. Otherwise, it returns all data repository tasks
// owned by your Amazon Web Services account in the Amazon Web Services Region of
// the endpoint that you're calling.
//
// When retrieving all tasks, you can paginate the response by using the optional
// MaxResults parameter to limit the number of tasks returned in a response. If
// more tasks remain, a NextToken value is returned in the response. In this case,
// send a later request with the NextToken request parameter set to the value of
// NextToken from the last response.
func (c *Client) DescribeDataRepositoryTasks(ctx context.Context, params *DescribeDataRepositoryTasksInput, optFns ...func(*Options)) (*DescribeDataRepositoryTasksOutput, error) {
	if params == nil {
		params = &DescribeDataRepositoryTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataRepositoryTasks", params, optFns, c.addOperationDescribeDataRepositoryTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataRepositoryTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataRepositoryTasksInput struct {

	// (Optional) You can use filters to narrow the DescribeDataRepositoryTasks
	// response to include just tasks for specific file systems, or tasks in a specific
	// lifecycle state.
	Filters []types.DataRepositoryTaskFilter

	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	MaxResults *int32

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	// (Optional) IDs of the tasks whose descriptions you want to retrieve (String).
	TaskIds []string

	noSmithyDocumentSerde
}

type DescribeDataRepositoryTasksOutput struct {

	// The collection of data repository task descriptions returned.
	DataRepositoryTasks []types.DataRepositoryTask

	// (Optional) Opaque pagination token returned from a previous operation (String).
	// If present, this token indicates from what point you can continue processing the
	// request, where the previous NextToken value left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataRepositoryTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataRepositoryTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataRepositoryTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDataRepositoryTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataRepositoryTasks(options.Region), middleware.Before); err != nil {
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

// DescribeDataRepositoryTasksPaginatorOptions is the paginator options for
// DescribeDataRepositoryTasks
type DescribeDataRepositoryTasksPaginatorOptions struct {
	// The maximum number of resources to return in the response. This value must be
	// an integer greater than zero.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDataRepositoryTasksPaginator is a paginator for
// DescribeDataRepositoryTasks
type DescribeDataRepositoryTasksPaginator struct {
	options   DescribeDataRepositoryTasksPaginatorOptions
	client    DescribeDataRepositoryTasksAPIClient
	params    *DescribeDataRepositoryTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeDataRepositoryTasksPaginator returns a new
// DescribeDataRepositoryTasksPaginator
func NewDescribeDataRepositoryTasksPaginator(client DescribeDataRepositoryTasksAPIClient, params *DescribeDataRepositoryTasksInput, optFns ...func(*DescribeDataRepositoryTasksPaginatorOptions)) *DescribeDataRepositoryTasksPaginator {
	if params == nil {
		params = &DescribeDataRepositoryTasksInput{}
	}

	options := DescribeDataRepositoryTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDataRepositoryTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDataRepositoryTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDataRepositoryTasks page.
func (p *DescribeDataRepositoryTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDataRepositoryTasksOutput, error) {
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
	result, err := p.client.DescribeDataRepositoryTasks(ctx, &params, optFns...)
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

// DescribeDataRepositoryTasksAPIClient is a client that implements the
// DescribeDataRepositoryTasks operation.
type DescribeDataRepositoryTasksAPIClient interface {
	DescribeDataRepositoryTasks(context.Context, *DescribeDataRepositoryTasksInput, ...func(*Options)) (*DescribeDataRepositoryTasksOutput, error)
}

var _ DescribeDataRepositoryTasksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeDataRepositoryTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDataRepositoryTasks",
	}
}
