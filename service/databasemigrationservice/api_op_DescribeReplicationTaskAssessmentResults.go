// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the task assessment results from the Amazon S3 bucket that DMS creates
// in your Amazon Web Services account. This action always returns the latest
// results.
//
// For more information about DMS task assessments, see [Creating a task assessment report] in the Database Migration
// Service User Guide.
//
// [Creating a task assessment report]: https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.AssessmentReport.html
func (c *Client) DescribeReplicationTaskAssessmentResults(ctx context.Context, params *DescribeReplicationTaskAssessmentResultsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentResultsOutput, error) {
	if params == nil {
		params = &DescribeReplicationTaskAssessmentResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationTaskAssessmentResults", params, optFns, c.addOperationDescribeReplicationTaskAssessmentResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationTaskAssessmentResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationTaskAssessmentResultsInput struct {

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) string that uniquely identifies the task. When
	// this input parameter is specified, the API returns only one result and ignore
	// the values of the MaxRecords and Marker parameters.
	ReplicationTaskArn *string

	noSmithyDocumentSerde
}

type DescribeReplicationTaskAssessmentResultsOutput struct {

	// - The Amazon S3 bucket where the task assessment report is located.
	BucketName *string

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	//  The task assessment report.
	ReplicationTaskAssessmentResults []types.ReplicationTaskAssessmentResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationTaskAssessmentResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTaskAssessmentResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTaskAssessmentResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicationTaskAssessmentResults"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentResults(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationTaskAssessmentResultsPaginatorOptions is the paginator
// options for DescribeReplicationTaskAssessmentResults
type DescribeReplicationTaskAssessmentResultsPaginatorOptions struct {
	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationTaskAssessmentResultsPaginator is a paginator for
// DescribeReplicationTaskAssessmentResults
type DescribeReplicationTaskAssessmentResultsPaginator struct {
	options   DescribeReplicationTaskAssessmentResultsPaginatorOptions
	client    DescribeReplicationTaskAssessmentResultsAPIClient
	params    *DescribeReplicationTaskAssessmentResultsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationTaskAssessmentResultsPaginator returns a new
// DescribeReplicationTaskAssessmentResultsPaginator
func NewDescribeReplicationTaskAssessmentResultsPaginator(client DescribeReplicationTaskAssessmentResultsAPIClient, params *DescribeReplicationTaskAssessmentResultsInput, optFns ...func(*DescribeReplicationTaskAssessmentResultsPaginatorOptions)) *DescribeReplicationTaskAssessmentResultsPaginator {
	if params == nil {
		params = &DescribeReplicationTaskAssessmentResultsInput{}
	}

	options := DescribeReplicationTaskAssessmentResultsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationTaskAssessmentResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationTaskAssessmentResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationTaskAssessmentResults page.
func (p *DescribeReplicationTaskAssessmentResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentResultsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeReplicationTaskAssessmentResults(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeReplicationTaskAssessmentResultsAPIClient is a client that implements
// the DescribeReplicationTaskAssessmentResults operation.
type DescribeReplicationTaskAssessmentResultsAPIClient interface {
	DescribeReplicationTaskAssessmentResults(context.Context, *DescribeReplicationTaskAssessmentResultsInput, ...func(*Options)) (*DescribeReplicationTaskAssessmentResultsOutput, error)
}

var _ DescribeReplicationTaskAssessmentResultsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicationTaskAssessmentResults",
	}
}
