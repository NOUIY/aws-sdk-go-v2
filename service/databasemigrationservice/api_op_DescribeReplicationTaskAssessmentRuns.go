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

// Returns a paginated list of premigration assessment runs based on filter
// settings.
//
// These filter settings can specify a combination of premigration assessment
// runs, migration tasks, replication instances, and assessment run status values.
//
// This operation doesn't return information about individual assessments. For
// this information, see the DescribeReplicationTaskIndividualAssessments
// operation.
func (c *Client) DescribeReplicationTaskAssessmentRuns(ctx context.Context, params *DescribeReplicationTaskAssessmentRunsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentRunsOutput, error) {
	if params == nil {
		params = &DescribeReplicationTaskAssessmentRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationTaskAssessmentRuns", params, optFns, c.addOperationDescribeReplicationTaskAssessmentRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationTaskAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationTaskAssessmentRunsInput struct {

	// Filters applied to the premigration assessment runs described in the form of
	// key-value pairs.
	//
	// Valid filter names: replication-task-assessment-run-arn , replication-task-arn ,
	// replication-instance-arn , status
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeReplicationTaskAssessmentRunsOutput struct {

	// A pagination token returned for you to pass to a subsequent request. If you
	// pass this token as the Marker value in a subsequent request, the response
	// includes only records beyond the marker, up to the value specified in the
	// request by MaxRecords .
	Marker *string

	// One or more premigration assessment runs as specified by Filters .
	ReplicationTaskAssessmentRuns []types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationTaskAssessmentRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTaskAssessmentRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTaskAssessmentRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicationTaskAssessmentRuns"); err != nil {
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
	if err = addOpDescribeReplicationTaskAssessmentRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentRuns(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationTaskAssessmentRunsPaginatorOptions is the paginator options
// for DescribeReplicationTaskAssessmentRuns
type DescribeReplicationTaskAssessmentRunsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationTaskAssessmentRunsPaginator is a paginator for
// DescribeReplicationTaskAssessmentRuns
type DescribeReplicationTaskAssessmentRunsPaginator struct {
	options   DescribeReplicationTaskAssessmentRunsPaginatorOptions
	client    DescribeReplicationTaskAssessmentRunsAPIClient
	params    *DescribeReplicationTaskAssessmentRunsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationTaskAssessmentRunsPaginator returns a new
// DescribeReplicationTaskAssessmentRunsPaginator
func NewDescribeReplicationTaskAssessmentRunsPaginator(client DescribeReplicationTaskAssessmentRunsAPIClient, params *DescribeReplicationTaskAssessmentRunsInput, optFns ...func(*DescribeReplicationTaskAssessmentRunsPaginatorOptions)) *DescribeReplicationTaskAssessmentRunsPaginator {
	if params == nil {
		params = &DescribeReplicationTaskAssessmentRunsInput{}
	}

	options := DescribeReplicationTaskAssessmentRunsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationTaskAssessmentRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationTaskAssessmentRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationTaskAssessmentRuns page.
func (p *DescribeReplicationTaskAssessmentRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentRunsOutput, error) {
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
	result, err := p.client.DescribeReplicationTaskAssessmentRuns(ctx, &params, optFns...)
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

// DescribeReplicationTaskAssessmentRunsAPIClient is a client that implements the
// DescribeReplicationTaskAssessmentRuns operation.
type DescribeReplicationTaskAssessmentRunsAPIClient interface {
	DescribeReplicationTaskAssessmentRuns(context.Context, *DescribeReplicationTaskAssessmentRunsInput, ...func(*Options)) (*DescribeReplicationTaskAssessmentRunsOutput, error)
}

var _ DescribeReplicationTaskAssessmentRunsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicationTaskAssessmentRuns",
	}
}
