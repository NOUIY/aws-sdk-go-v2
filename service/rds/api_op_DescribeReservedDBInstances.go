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

// Returns information about reserved DB instances for this account, or about a
// specified reserved DB instance.
func (c *Client) DescribeReservedDBInstances(ctx context.Context, params *DescribeReservedDBInstancesInput, optFns ...func(*Options)) (*DescribeReservedDBInstancesOutput, error) {
	if params == nil {
		params = &DescribeReservedDBInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedDBInstances", params, optFns, c.addOperationDescribeReservedDBInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedDBInstancesInput struct {

	// The DB instance class filter value. Specify this parameter to show only those
	// reservations matching the specified DB instances class.
	DBInstanceClass *string

	// The duration filter value, specified in years or seconds. Specify this
	// parameter to show only reservations for this duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// The lease identifier filter value. Specify this parameter to show only the
	// reservation that matches the specified lease ID.
	//
	// Amazon Web Services Support might request the lease ID for an issue related to
	// a reserved DB instance.
	LeaseId *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// Specifies whether to show only those reservations that support Multi-AZ.
	MultiAZ *bool

	// The offering type filter value. Specify this parameter to show only the
	// available offerings matching the specified offering type.
	//
	// Valid Values: "Partial Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string

	// The product description filter value. Specify this parameter to show only those
	// reservations matching the specified product description.
	ProductDescription *string

	// The reserved DB instance identifier filter value. Specify this parameter to
	// show only the reservation that matches the specified reservation ID.
	ReservedDBInstanceId *string

	// The offering identifier filter value. Specify this parameter to show only
	// purchased reservations matching the specified offering identifier.
	ReservedDBInstancesOfferingId *string

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeReservedDBInstances action.
type DescribeReservedDBInstancesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// A list of reserved DB instances.
	ReservedDBInstances []types.ReservedDBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedDBInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedDBInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedDBInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedDBInstances"); err != nil {
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
	if err = addOpDescribeReservedDBInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedDBInstances(options.Region), middleware.Before); err != nil {
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

// DescribeReservedDBInstancesPaginatorOptions is the paginator options for
// DescribeReservedDBInstances
type DescribeReservedDBInstancesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedDBInstancesPaginator is a paginator for
// DescribeReservedDBInstances
type DescribeReservedDBInstancesPaginator struct {
	options   DescribeReservedDBInstancesPaginatorOptions
	client    DescribeReservedDBInstancesAPIClient
	params    *DescribeReservedDBInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedDBInstancesPaginator returns a new
// DescribeReservedDBInstancesPaginator
func NewDescribeReservedDBInstancesPaginator(client DescribeReservedDBInstancesAPIClient, params *DescribeReservedDBInstancesInput, optFns ...func(*DescribeReservedDBInstancesPaginatorOptions)) *DescribeReservedDBInstancesPaginator {
	if params == nil {
		params = &DescribeReservedDBInstancesInput{}
	}

	options := DescribeReservedDBInstancesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedDBInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedDBInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedDBInstances page.
func (p *DescribeReservedDBInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedDBInstancesOutput, error) {
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
	result, err := p.client.DescribeReservedDBInstances(ctx, &params, optFns...)
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

// DescribeReservedDBInstancesAPIClient is a client that implements the
// DescribeReservedDBInstances operation.
type DescribeReservedDBInstancesAPIClient interface {
	DescribeReservedDBInstances(context.Context, *DescribeReservedDBInstancesInput, ...func(*Options)) (*DescribeReservedDBInstancesOutput, error)
}

var _ DescribeReservedDBInstancesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReservedDBInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedDBInstances",
	}
}
