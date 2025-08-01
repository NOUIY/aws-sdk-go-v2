// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a consolidated report of your workloads.
//
// You can optionally choose to include workloads that have been shared with you.
func (c *Client) GetConsolidatedReport(ctx context.Context, params *GetConsolidatedReportInput, optFns ...func(*Options)) (*GetConsolidatedReportOutput, error) {
	if params == nil {
		params = &GetConsolidatedReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConsolidatedReport", params, optFns, c.addOperationGetConsolidatedReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConsolidatedReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConsolidatedReportInput struct {

	// The format of the consolidated report.
	//
	// For PDF , Base64String is returned. For JSON , Metrics is returned.
	//
	// This member is required.
	Format types.ReportFormat

	// Set to true to have shared resources included in the report.
	IncludeSharedResources *bool

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConsolidatedReportOutput struct {

	// The Base64-encoded string representation of a lens review report.
	//
	// This data can be used to create a PDF file.
	//
	// Only returned by GetConsolidatedReport when PDF format is requested.
	Base64String *string

	// The metrics that make up the consolidated report.
	//
	// Only returned when JSON format is requested.
	Metrics []types.ConsolidatedReportMetric

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConsolidatedReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConsolidatedReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConsolidatedReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConsolidatedReport"); err != nil {
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
	if err = addOpGetConsolidatedReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConsolidatedReport(options.Region), middleware.Before); err != nil {
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

// GetConsolidatedReportPaginatorOptions is the paginator options for
// GetConsolidatedReport
type GetConsolidatedReportPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConsolidatedReportPaginator is a paginator for GetConsolidatedReport
type GetConsolidatedReportPaginator struct {
	options   GetConsolidatedReportPaginatorOptions
	client    GetConsolidatedReportAPIClient
	params    *GetConsolidatedReportInput
	nextToken *string
	firstPage bool
}

// NewGetConsolidatedReportPaginator returns a new GetConsolidatedReportPaginator
func NewGetConsolidatedReportPaginator(client GetConsolidatedReportAPIClient, params *GetConsolidatedReportInput, optFns ...func(*GetConsolidatedReportPaginatorOptions)) *GetConsolidatedReportPaginator {
	if params == nil {
		params = &GetConsolidatedReportInput{}
	}

	options := GetConsolidatedReportPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConsolidatedReportPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConsolidatedReportPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConsolidatedReport page.
func (p *GetConsolidatedReportPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConsolidatedReportOutput, error) {
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
	result, err := p.client.GetConsolidatedReport(ctx, &params, optFns...)
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

// GetConsolidatedReportAPIClient is a client that implements the
// GetConsolidatedReport operation.
type GetConsolidatedReportAPIClient interface {
	GetConsolidatedReport(context.Context, *GetConsolidatedReportInput, ...func(*Options)) (*GetConsolidatedReportOutput, error)
}

var _ GetConsolidatedReportAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetConsolidatedReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConsolidatedReport",
	}
}
