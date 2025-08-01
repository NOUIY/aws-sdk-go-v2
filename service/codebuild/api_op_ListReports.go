// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of ARNs for the reports in the current Amazon Web Services
//
// account.
func (c *Client) ListReports(ctx context.Context, params *ListReportsInput, optFns ...func(*Options)) (*ListReportsOutput, error) {
	if params == nil {
		params = &ListReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReports", params, optFns, c.addOperationListReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportsInput struct {

	//  A ReportFilter object used to filter the returned reports.
	Filter *types.ReportFilter

	//  The maximum number of paginated reports returned per response. Use nextToken
	// to iterate pages in the list of returned Report objects. The default value is
	// 100.
	MaxResults *int32

	//  During a previous call, the maximum number of items that can be returned is
	// the value specified in maxResults . If there more items in the list, then a
	// unique string called a nextToken is returned. To get the next batch of items in
	// the list, call this operation again, adding the next token to the call. To get
	// all of the items in the list, keep calling this operation with each subsequent
	// next token that is returned, until no more next tokens are returned.
	NextToken *string

	//  Specifies the sort order for the list of returned reports. Valid values are:
	//
	//   - ASCENDING : return reports in chronological order based on their creation
	//   date.
	//
	//   - DESCENDING : return reports in the reverse chronological order based on
	//   their creation date.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListReportsOutput struct {

	//  During a previous call, the maximum number of items that can be returned is
	// the value specified in maxResults . If there more items in the list, then a
	// unique string called a nextToken is returned. To get the next batch of items in
	// the list, call this operation again, adding the next token to the call. To get
	// all of the items in the list, keep calling this operation with each subsequent
	// next token that is returned, until no more next tokens are returned.
	NextToken *string

	//  The list of returned ARNs for the reports in the current Amazon Web Services
	// account.
	Reports []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReports"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReports(options.Region), middleware.Before); err != nil {
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

// ListReportsPaginatorOptions is the paginator options for ListReports
type ListReportsPaginatorOptions struct {
	//  The maximum number of paginated reports returned per response. Use nextToken
	// to iterate pages in the list of returned Report objects. The default value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReportsPaginator is a paginator for ListReports
type ListReportsPaginator struct {
	options   ListReportsPaginatorOptions
	client    ListReportsAPIClient
	params    *ListReportsInput
	nextToken *string
	firstPage bool
}

// NewListReportsPaginator returns a new ListReportsPaginator
func NewListReportsPaginator(client ListReportsAPIClient, params *ListReportsInput, optFns ...func(*ListReportsPaginatorOptions)) *ListReportsPaginator {
	if params == nil {
		params = &ListReportsInput{}
	}

	options := ListReportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReports page.
func (p *ListReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReportsOutput, error) {
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
	result, err := p.client.ListReports(ctx, &params, optFns...)
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

// ListReportsAPIClient is a client that implements the ListReports operation.
type ListReportsAPIClient interface {
	ListReports(context.Context, *ListReportsInput, ...func(*Options)) (*ListReportsOutput, error)
}

var _ ListReportsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReports",
	}
}
