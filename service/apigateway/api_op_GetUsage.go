// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the usage data of a usage plan in a specified time interval.
func (c *Client) GetUsage(ctx context.Context, params *GetUsageInput, optFns ...func(*Options)) (*GetUsageOutput, error) {
	if params == nil {
		params = &GetUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsage", params, optFns, c.addOperationGetUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get the usage data of a usage plan in a specified time
// interval.
type GetUsageInput struct {

	// The ending date (e.g., 2016-12-31) of the usage data.
	//
	// This member is required.
	EndDate *string

	// The starting date (e.g., 2016-01-01) of the usage data.
	//
	// This member is required.
	StartDate *string

	// The Id of the usage plan associated with the usage data.
	//
	// This member is required.
	UsagePlanId *string

	// The Id of the API key associated with the resultant usage data.
	KeyId *string

	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents the usage data of a usage plan.
type GetUsageOutput struct {

	// The ending date of the usage data.
	EndDate *string

	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]} , where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota] .
	Items map[string][][]int64

	// The current pagination position in the paged result set.
	Position *string

	// The starting date of the usage data.
	StartDate *string

	// The plan Id associated with this usage data.
	UsagePlanId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUsage"); err != nil {
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
	if err = addOpGetUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsage(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

// GetUsagePaginatorOptions is the paginator options for GetUsage
type GetUsagePaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUsagePaginator is a paginator for GetUsage
type GetUsagePaginator struct {
	options   GetUsagePaginatorOptions
	client    GetUsageAPIClient
	params    *GetUsageInput
	nextToken *string
	firstPage bool
}

// NewGetUsagePaginator returns a new GetUsagePaginator
func NewGetUsagePaginator(client GetUsageAPIClient, params *GetUsageInput, optFns ...func(*GetUsagePaginatorOptions)) *GetUsagePaginator {
	if params == nil {
		params = &GetUsageInput{}
	}

	options := GetUsagePaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUsagePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUsagePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUsage page.
func (p *GetUsagePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUsageOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetUsage(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetUsageAPIClient is a client that implements the GetUsage operation.
type GetUsageAPIClient interface {
	GetUsage(context.Context, *GetUsageInput, ...func(*Options)) (*GetUsageOutput, error)
}

var _ GetUsageAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUsage",
	}
}
