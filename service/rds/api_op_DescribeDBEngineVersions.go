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

// Describes the properties of specific versions of DB engines.
func (c *Client) DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBEngineVersions", params, optFns, c.addOperationDescribeDBEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBEngineVersionsInput struct {

	// The name of a specific DB parameter group family to return details for.
	//
	// Constraints:
	//
	//   - If supplied, must match an existing DB parameter group family.
	DBParameterGroupFamily *string

	// Specifies whether to return only the default version of the specified engine or
	// the engine and major version combination.
	DefaultOnly *bool

	// The database engine to return version details for.
	//
	// Valid Values:
	//
	//   - aurora-mysql
	//
	//   - aurora-postgresql
	//
	//   - custom-oracle-ee
	//
	//   - custom-oracle-ee-cdb
	//
	//   - custom-oracle-se2
	//
	//   - custom-oracle-se2-cdb
	//
	//   - db2-ae
	//
	//   - db2-se
	//
	//   - mariadb
	//
	//   - mysql
	//
	//   - oracle-ee
	//
	//   - oracle-ee-cdb
	//
	//   - oracle-se2
	//
	//   - oracle-se2-cdb
	//
	//   - postgres
	//
	//   - sqlserver-ee
	//
	//   - sqlserver-se
	//
	//   - sqlserver-ex
	//
	//   - sqlserver-web
	Engine *string

	// A specific database engine version to return details for.
	//
	// Example: 5.1.49
	EngineVersion *string

	// A filter that specifies one or more DB engine versions to describe.
	//
	// Supported filters:
	//
	//   - db-parameter-group-family - Accepts parameter groups family names. The
	//   results list only includes information about the DB engine versions for these
	//   parameter group families.
	//
	//   - engine - Accepts engine names. The results list only includes information
	//   about the DB engine versions for these engines.
	//
	//   - engine-mode - Accepts DB engine modes. The results list only includes
	//   information about the DB engine versions for these engine modes. Valid DB engine
	//   modes are the following:
	//
	//   - global
	//
	//   - multimaster
	//
	//   - parallelquery
	//
	//   - provisioned
	//
	//   - serverless
	//
	//   - engine-version - Accepts engine versions. The results list only includes
	//   information about the DB engine versions for these engine versions.
	//
	//   - status - Accepts engine version statuses. The results list only includes
	//   information about the DB engine versions for these statuses. Valid statuses are
	//   the following:
	//
	//   - available
	//
	//   - deprecated
	Filters []types.Filter

	// Specifies whether to also list the engine versions that aren't available. The
	// default is to list only available engine versions.
	IncludeAll *bool

	// Specifies whether to list the supported character sets for each engine version.
	//
	// If this parameter is enabled and the requested engine supports the
	// CharacterSetName parameter for CreateDBInstance , the response includes a list
	// of supported character sets for each engine version.
	//
	// For RDS Custom, the default is not to list supported character sets. If you
	// enable this parameter, RDS Custom returns no results.
	ListSupportedCharacterSets *bool

	// Specifies whether to list the supported time zones for each engine version.
	//
	// If this parameter is enabled and the requested engine supports the TimeZone
	// parameter for CreateDBInstance , the response includes a list of supported time
	// zones for each engine version.
	//
	// For RDS Custom, the default is not to list supported time zones. If you enable
	// this parameter, RDS Custom returns no results.
	ListSupportedTimezones *bool

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

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the DescribeDBEngineVersions
// action.
type DescribeDBEngineVersionsOutput struct {

	// A list of DBEngineVersion elements.
	DBEngineVersions []types.DBEngineVersion

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBEngineVersions"); err != nil {
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
	if err = addOpDescribeDBEngineVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBEngineVersions(options.Region), middleware.Before); err != nil {
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

// DescribeDBEngineVersionsPaginatorOptions is the paginator options for
// DescribeDBEngineVersions
type DescribeDBEngineVersionsPaginatorOptions struct {
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

// DescribeDBEngineVersionsPaginator is a paginator for DescribeDBEngineVersions
type DescribeDBEngineVersionsPaginator struct {
	options   DescribeDBEngineVersionsPaginatorOptions
	client    DescribeDBEngineVersionsAPIClient
	params    *DescribeDBEngineVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBEngineVersionsPaginator returns a new
// DescribeDBEngineVersionsPaginator
func NewDescribeDBEngineVersionsPaginator(client DescribeDBEngineVersionsAPIClient, params *DescribeDBEngineVersionsInput, optFns ...func(*DescribeDBEngineVersionsPaginatorOptions)) *DescribeDBEngineVersionsPaginator {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	options := DescribeDBEngineVersionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBEngineVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBEngineVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBEngineVersions page.
func (p *DescribeDBEngineVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
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
	result, err := p.client.DescribeDBEngineVersions(ctx, &params, optFns...)
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

// DescribeDBEngineVersionsAPIClient is a client that implements the
// DescribeDBEngineVersions operation.
type DescribeDBEngineVersionsAPIClient interface {
	DescribeDBEngineVersions(context.Context, *DescribeDBEngineVersionsInput, ...func(*Options)) (*DescribeDBEngineVersionsOutput, error)
}

var _ DescribeDBEngineVersionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeDBEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBEngineVersions",
	}
}
