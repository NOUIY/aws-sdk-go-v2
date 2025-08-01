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

// Returns a paginated list of metadata model conversions for a migration project.
func (c *Client) DescribeMetadataModelConversions(ctx context.Context, params *DescribeMetadataModelConversionsInput, optFns ...func(*Options)) (*DescribeMetadataModelConversionsOutput, error) {
	if params == nil {
		params = &DescribeMetadataModelConversionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetadataModelConversions", params, optFns, c.addOperationDescribeMetadataModelConversionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetadataModelConversionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetadataModelConversionsInput struct {

	// The migration project name or Amazon Resource Name (ARN).
	//
	// This member is required.
	MigrationProjectIdentifier *string

	// Filters applied to the metadata model conversions described in the form of
	// key-value pairs.
	Filters []types.Filter

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeMetadataModelConversionsOutput struct {

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// A paginated list of metadata model conversions.
	Requests []types.SchemaConversionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMetadataModelConversionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMetadataModelConversions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMetadataModelConversions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMetadataModelConversions"); err != nil {
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
	if err = addOpDescribeMetadataModelConversionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetadataModelConversions(options.Region), middleware.Before); err != nil {
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

// DescribeMetadataModelConversionsPaginatorOptions is the paginator options for
// DescribeMetadataModelConversions
type DescribeMetadataModelConversionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMetadataModelConversionsPaginator is a paginator for
// DescribeMetadataModelConversions
type DescribeMetadataModelConversionsPaginator struct {
	options   DescribeMetadataModelConversionsPaginatorOptions
	client    DescribeMetadataModelConversionsAPIClient
	params    *DescribeMetadataModelConversionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMetadataModelConversionsPaginator returns a new
// DescribeMetadataModelConversionsPaginator
func NewDescribeMetadataModelConversionsPaginator(client DescribeMetadataModelConversionsAPIClient, params *DescribeMetadataModelConversionsInput, optFns ...func(*DescribeMetadataModelConversionsPaginatorOptions)) *DescribeMetadataModelConversionsPaginator {
	if params == nil {
		params = &DescribeMetadataModelConversionsInput{}
	}

	options := DescribeMetadataModelConversionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMetadataModelConversionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMetadataModelConversionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMetadataModelConversions page.
func (p *DescribeMetadataModelConversionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMetadataModelConversionsOutput, error) {
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
	result, err := p.client.DescribeMetadataModelConversions(ctx, &params, optFns...)
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

// DescribeMetadataModelConversionsAPIClient is a client that implements the
// DescribeMetadataModelConversions operation.
type DescribeMetadataModelConversionsAPIClient interface {
	DescribeMetadataModelConversions(context.Context, *DescribeMetadataModelConversionsInput, ...func(*Options)) (*DescribeMetadataModelConversionsOutput, error)
}

var _ DescribeMetadataModelConversionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeMetadataModelConversions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMetadataModelConversions",
	}
}
