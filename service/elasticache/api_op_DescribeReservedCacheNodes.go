// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about reserved cache nodes for this account, or about a
// specified reserved cache node.
func (c *Client) DescribeReservedCacheNodes(ctx context.Context, params *DescribeReservedCacheNodesInput, optFns ...func(*Options)) (*DescribeReservedCacheNodesOutput, error) {
	if params == nil {
		params = &DescribeReservedCacheNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedCacheNodes", params, optFns, c.addOperationDescribeReservedCacheNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedCacheNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeReservedCacheNodes operation.
type DescribeReservedCacheNodesInput struct {

	// The cache node type filter value. Use this parameter to show only those
	// reservations matching the specified cache node type.
	//
	// The following node types are supported by ElastiCache. Generally speaking, the
	// current generation types provide more memory and computational power at lower
	// cost when compared to their equivalent previous generation counterparts.
	//
	//   - General purpose:
	//
	//   - Current generation:
	//
	// M7g node types: cache.m7g.large , cache.m7g.xlarge , cache.m7g.2xlarge ,
	//   cache.m7g.4xlarge , cache.m7g.8xlarge , cache.m7g.12xlarge ,
	//   cache.m7g.16xlarge
	//
	// For region availability, see [Supported Node Types]
	//
	// M6g node types (available only for Redis OSS engine version 5.0.6 onward and
	//   for Memcached engine version 1.5.16 onward):
	//
	// cache.m6g.large , cache.m6g.xlarge , cache.m6g.2xlarge , cache.m6g.4xlarge ,
	//   cache.m6g.8xlarge , cache.m6g.12xlarge , cache.m6g.16xlarge
	//
	// M5 node types: cache.m5.large , cache.m5.xlarge , cache.m5.2xlarge ,
	//   cache.m5.4xlarge , cache.m5.12xlarge , cache.m5.24xlarge
	//
	// M4 node types: cache.m4.large , cache.m4.xlarge , cache.m4.2xlarge ,
	//   cache.m4.4xlarge , cache.m4.10xlarge
	//
	// T4g node types (available only for Redis OSS engine version 5.0.6 onward and
	//   Memcached engine version 1.5.16 onward): cache.t4g.micro , cache.t4g.small ,
	//   cache.t4g.medium
	//
	// T3 node types: cache.t3.micro , cache.t3.small , cache.t3.medium
	//
	// T2 node types: cache.t2.micro , cache.t2.small , cache.t2.medium
	//
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.)
	//
	// T1 node types: cache.t1.micro
	//
	// M1 node types: cache.m1.small , cache.m1.medium , cache.m1.large ,
	//   cache.m1.xlarge
	//
	// M3 node types: cache.m3.medium , cache.m3.large , cache.m3.xlarge ,
	//   cache.m3.2xlarge
	//
	//   - Compute optimized:
	//
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.)
	//
	// C1 node types: cache.c1.xlarge
	//
	//   - Memory optimized:
	//
	//   - Current generation:
	//
	// R7g node types: cache.r7g.large , cache.r7g.xlarge , cache.r7g.2xlarge ,
	//   cache.r7g.4xlarge , cache.r7g.8xlarge , cache.r7g.12xlarge ,
	//   cache.r7g.16xlarge
	//
	// For region availability, see [Supported Node Types]
	//
	// R6g node types (available only for Redis OSS engine version 5.0.6 onward and
	//   for Memcached engine version 1.5.16 onward): cache.r6g.large ,
	//   cache.r6g.xlarge , cache.r6g.2xlarge , cache.r6g.4xlarge , cache.r6g.8xlarge ,
	//   cache.r6g.12xlarge , cache.r6g.16xlarge
	//
	// R5 node types: cache.r5.large , cache.r5.xlarge , cache.r5.2xlarge ,
	//   cache.r5.4xlarge , cache.r5.12xlarge , cache.r5.24xlarge
	//
	// R4 node types: cache.r4.large , cache.r4.xlarge , cache.r4.2xlarge ,
	//   cache.r4.4xlarge , cache.r4.8xlarge , cache.r4.16xlarge
	//
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.)
	//
	// M2 node types: cache.m2.xlarge , cache.m2.2xlarge , cache.m2.4xlarge
	//
	// R3 node types: cache.r3.large , cache.r3.xlarge , cache.r3.2xlarge ,
	//
	// cache.r3.4xlarge , cache.r3.8xlarge
	//
	// Additional node type info
	//
	//   - All current generation instance types are created in Amazon VPC by default.
	//
	//   - Valkey or Redis OSS append-only files (AOF) are not supported for T1 or T2
	//   instances.
	//
	//   - Valkey or Redis OSS Multi-AZ with automatic failover is not supported on T1
	//   instances.
	//
	//   - The configuration variables appendonly and appendfsync are not supported on
	//   Valkey, or on Redis OSS version 2.8.22 and later.
	//
	// [Supported Node Types]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion
	CacheNodeType *string

	// The duration filter value, specified in years or seconds. Use this parameter to
	// show only reservations for this duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int32

	// The offering type filter value. Use this parameter to show only the available
	// offerings matching the specified offering type.
	//
	// Valid values: "Light Utilization"|"Medium Utilization"|"Heavy Utilization"|"All
	// Upfront"|"Partial Upfront"| "No Upfront"
	OfferingType *string

	// The product description filter value. Use this parameter to show only those
	// reservations matching the specified product description.
	ProductDescription *string

	// The reserved cache node identifier filter value. Use this parameter to show
	// only the reservation that matches the specified reservation ID.
	ReservedCacheNodeId *string

	// The offering identifier filter value. Use this parameter to show only purchased
	// reservations matching the specified offering identifier.
	ReservedCacheNodesOfferingId *string

	noSmithyDocumentSerde
}

// Represents the output of a DescribeReservedCacheNodes operation.
type DescribeReservedCacheNodesOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// A list of reserved cache nodes. Each element in the list contains detailed
	// information about one node.
	ReservedCacheNodes []types.ReservedCacheNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedCacheNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedCacheNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedCacheNodes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedCacheNodes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedCacheNodes(options.Region), middleware.Before); err != nil {
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

// DescribeReservedCacheNodesPaginatorOptions is the paginator options for
// DescribeReservedCacheNodes
type DescribeReservedCacheNodesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedCacheNodesPaginator is a paginator for
// DescribeReservedCacheNodes
type DescribeReservedCacheNodesPaginator struct {
	options   DescribeReservedCacheNodesPaginatorOptions
	client    DescribeReservedCacheNodesAPIClient
	params    *DescribeReservedCacheNodesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedCacheNodesPaginator returns a new
// DescribeReservedCacheNodesPaginator
func NewDescribeReservedCacheNodesPaginator(client DescribeReservedCacheNodesAPIClient, params *DescribeReservedCacheNodesInput, optFns ...func(*DescribeReservedCacheNodesPaginatorOptions)) *DescribeReservedCacheNodesPaginator {
	if params == nil {
		params = &DescribeReservedCacheNodesInput{}
	}

	options := DescribeReservedCacheNodesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedCacheNodesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedCacheNodesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedCacheNodes page.
func (p *DescribeReservedCacheNodesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedCacheNodesOutput, error) {
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
	result, err := p.client.DescribeReservedCacheNodes(ctx, &params, optFns...)
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

// DescribeReservedCacheNodesAPIClient is a client that implements the
// DescribeReservedCacheNodes operation.
type DescribeReservedCacheNodesAPIClient interface {
	DescribeReservedCacheNodes(context.Context, *DescribeReservedCacheNodesInput, ...func(*Options)) (*DescribeReservedCacheNodesOutput, error)
}

var _ DescribeReservedCacheNodesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReservedCacheNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedCacheNodes",
	}
}
