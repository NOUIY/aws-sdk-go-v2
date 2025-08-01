// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Fetches the temporarily cached result of an SQL statement in JSON format. The
// ExecuteStatement or BatchExecuteStatement operation that ran the SQL statement
// must have specified ResultFormat as JSON , or let the format default to JSON. A
// token is returned to page through the statement results.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
func (c *Client) GetStatementResult(ctx context.Context, params *GetStatementResultInput, optFns ...func(*Options)) (*GetStatementResultOutput, error) {
	if params == nil {
		params = &GetStatementResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStatementResult", params, optFns, c.addOperationGetStatementResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStatementResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStatementResultInput struct {

	// The identifier of the SQL statement whose results are to be fetched. This value
	// is a universally unique identifier (UUID) generated by Amazon Redshift Data API.
	// A suffix indicates then number of the SQL statement. For example,
	// d9b6c0c9-0747-4bf4-b142-e8883122f766:2 has a suffix of :2 that indicates the
	// second SQL statement of a batch query. This identifier is returned by
	// BatchExecuteStatment , ExecuteStatment , and ListStatements .
	//
	// This member is required.
	Id *string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetStatementResultOutput struct {

	// The results of the SQL statement in JSON format.
	//
	// This member is required.
	Records [][]types.Field

	// The properties (metadata) of a column.
	ColumnMetadata []types.ColumnMetadata

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The total number of rows in the result set returned from a query. You can use
	// this number to estimate the number of calls to the GetStatementResult operation
	// needed to page through the results.
	TotalNumRows int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStatementResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetStatementResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetStatementResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetStatementResult"); err != nil {
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
	if err = addOpGetStatementResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStatementResult(options.Region), middleware.Before); err != nil {
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

// GetStatementResultPaginatorOptions is the paginator options for
// GetStatementResult
type GetStatementResultPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetStatementResultPaginator is a paginator for GetStatementResult
type GetStatementResultPaginator struct {
	options   GetStatementResultPaginatorOptions
	client    GetStatementResultAPIClient
	params    *GetStatementResultInput
	nextToken *string
	firstPage bool
}

// NewGetStatementResultPaginator returns a new GetStatementResultPaginator
func NewGetStatementResultPaginator(client GetStatementResultAPIClient, params *GetStatementResultInput, optFns ...func(*GetStatementResultPaginatorOptions)) *GetStatementResultPaginator {
	if params == nil {
		params = &GetStatementResultInput{}
	}

	options := GetStatementResultPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetStatementResultPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetStatementResultPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetStatementResult page.
func (p *GetStatementResultPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetStatementResultOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetStatementResult(ctx, &params, optFns...)
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

// GetStatementResultAPIClient is a client that implements the GetStatementResult
// operation.
type GetStatementResultAPIClient interface {
	GetStatementResult(context.Context, *GetStatementResultInput, ...func(*Options)) (*GetStatementResultOutput, error)
}

var _ GetStatementResultAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetStatementResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetStatementResult",
	}
}
