// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about backtracks for a DB cluster.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// This action only applies to Aurora MySQL DB clusters.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func (c *Client) DescribeDBClusterBacktracks(ctx context.Context, params *DescribeDBClusterBacktracksInput, optFns ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error) {
	if params == nil {
		params = &DescribeDBClusterBacktracksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterBacktracks", params, optFns, c.addOperationDescribeDBClusterBacktracksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterBacktracksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterBacktracksInput struct {

	// The DB cluster identifier of the DB cluster to be described. This parameter is
	// stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//   - First character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// If specified, this value is the backtrack identifier of the backtrack to be
	// described.
	//
	// Constraints:
	//
	//   - Must contain a valid universally unique identifier (UUID). For more
	//   information about UUIDs, see [Universally unique identifier].
	//
	// Example: 123e4567-e89b-12d3-a456-426655440000
	//
	// [Universally unique identifier]: https://en.wikipedia.org/wiki/Universally_unique_identifier
	BacktrackIdentifier *string

	// A filter that specifies one or more DB clusters to describe. Supported filters
	// include the following:
	//
	//   - db-cluster-backtrack-id - Accepts backtrack identifiers. The results list
	//   includes information about only the backtracks identified by these identifiers.
	//
	//   - db-cluster-backtrack-status - Accepts any of the following backtrack status
	//   values:
	//
	//   - applying
	//
	//   - completed
	//
	//   - failed
	//
	//   - pending
	//
	// The results list includes information about only the backtracks identified by
	//   these values.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterBacktracks
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeDBClusterBacktracks action.
type DescribeDBClusterBacktracksOutput struct {

	// Contains a list of backtracks for the user.
	DBClusterBacktracks []types.DBClusterBacktrack

	// A pagination token that can be used in a later DescribeDBClusterBacktracks
	// request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterBacktracksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterBacktracks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterBacktracks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterBacktracks"); err != nil {
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
	if err = addOpDescribeDBClusterBacktracksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterBacktracksPaginatorOptions is the paginator options for
// DescribeDBClusterBacktracks
type DescribeDBClusterBacktracksPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterBacktracksPaginator is a paginator for
// DescribeDBClusterBacktracks
type DescribeDBClusterBacktracksPaginator struct {
	options   DescribeDBClusterBacktracksPaginatorOptions
	client    DescribeDBClusterBacktracksAPIClient
	params    *DescribeDBClusterBacktracksInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterBacktracksPaginator returns a new
// DescribeDBClusterBacktracksPaginator
func NewDescribeDBClusterBacktracksPaginator(client DescribeDBClusterBacktracksAPIClient, params *DescribeDBClusterBacktracksInput, optFns ...func(*DescribeDBClusterBacktracksPaginatorOptions)) *DescribeDBClusterBacktracksPaginator {
	if params == nil {
		params = &DescribeDBClusterBacktracksInput{}
	}

	options := DescribeDBClusterBacktracksPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterBacktracksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterBacktracksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterBacktracks page.
func (p *DescribeDBClusterBacktracksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error) {
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
	result, err := p.client.DescribeDBClusterBacktracks(ctx, &params, optFns...)
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

// DescribeDBClusterBacktracksAPIClient is a client that implements the
// DescribeDBClusterBacktracks operation.
type DescribeDBClusterBacktracksAPIClient interface {
	DescribeDBClusterBacktracks(context.Context, *DescribeDBClusterBacktracksInput, ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error)
}

var _ DescribeDBClusterBacktracksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterBacktracks",
	}
}
