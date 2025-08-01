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

//	End of support notice: On May 20, 2026, Amazon Web Services will end support
//
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Provides descriptions of the schemas discovered by your Fleet Advisor
// collectors.
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
func (c *Client) DescribeFleetAdvisorSchemaObjectSummary(ctx context.Context, params *DescribeFleetAdvisorSchemaObjectSummaryInput, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemaObjectSummaryOutput, error) {
	if params == nil {
		params = &DescribeFleetAdvisorSchemaObjectSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAdvisorSchemaObjectSummary", params, optFns, c.addOperationDescribeFleetAdvisorSchemaObjectSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAdvisorSchemaObjectSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetAdvisorSchemaObjectSummaryInput struct {

	//  If you specify any of the following filters, the output includes information
	// for only those schema objects that meet the filter criteria:
	//
	//   - schema-id – The ID of the schema, for example
	//   d4610ac5-e323-4ad9-bc50-eaf7249dfe9d .
	//
	// Example: describe-fleet-advisor-schema-object-summary --filter
	// Name="schema-id",Values="50"
	Filters []types.Filter

	//  End of support notice: On May 20, 2026, Amazon Web Services will end support
	// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
	// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
	// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
	//
	// Sets the maximum number of records returned in the response.
	//
	// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
	MaxRecords *int32

	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetAdvisorSchemaObjectSummaryOutput struct {

	// A collection of FleetAdvisorSchemaObjectResponse objects.
	FleetAdvisorSchemaObjects []types.FleetAdvisorSchemaObjectResponse

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetAdvisorSchemaObjectSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAdvisorSchemaObjectSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAdvisorSchemaObjectSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetAdvisorSchemaObjectSummary"); err != nil {
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
	if err = addOpDescribeFleetAdvisorSchemaObjectSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAdvisorSchemaObjectSummary(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAdvisorSchemaObjectSummaryPaginatorOptions is the paginator
// options for DescribeFleetAdvisorSchemaObjectSummary
type DescribeFleetAdvisorSchemaObjectSummaryPaginatorOptions struct {
	//  End of support notice: On May 20, 2026, Amazon Web Services will end support
	// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
	// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
	// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
	//
	// Sets the maximum number of records returned in the response.
	//
	// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAdvisorSchemaObjectSummaryPaginator is a paginator for
// DescribeFleetAdvisorSchemaObjectSummary
type DescribeFleetAdvisorSchemaObjectSummaryPaginator struct {
	options   DescribeFleetAdvisorSchemaObjectSummaryPaginatorOptions
	client    DescribeFleetAdvisorSchemaObjectSummaryAPIClient
	params    *DescribeFleetAdvisorSchemaObjectSummaryInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAdvisorSchemaObjectSummaryPaginator returns a new
// DescribeFleetAdvisorSchemaObjectSummaryPaginator
func NewDescribeFleetAdvisorSchemaObjectSummaryPaginator(client DescribeFleetAdvisorSchemaObjectSummaryAPIClient, params *DescribeFleetAdvisorSchemaObjectSummaryInput, optFns ...func(*DescribeFleetAdvisorSchemaObjectSummaryPaginatorOptions)) *DescribeFleetAdvisorSchemaObjectSummaryPaginator {
	if params == nil {
		params = &DescribeFleetAdvisorSchemaObjectSummaryInput{}
	}

	options := DescribeFleetAdvisorSchemaObjectSummaryPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetAdvisorSchemaObjectSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAdvisorSchemaObjectSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetAdvisorSchemaObjectSummary page.
func (p *DescribeFleetAdvisorSchemaObjectSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAdvisorSchemaObjectSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeFleetAdvisorSchemaObjectSummary(ctx, &params, optFns...)
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

// DescribeFleetAdvisorSchemaObjectSummaryAPIClient is a client that implements
// the DescribeFleetAdvisorSchemaObjectSummary operation.
type DescribeFleetAdvisorSchemaObjectSummaryAPIClient interface {
	DescribeFleetAdvisorSchemaObjectSummary(context.Context, *DescribeFleetAdvisorSchemaObjectSummaryInput, ...func(*Options)) (*DescribeFleetAdvisorSchemaObjectSummaryOutput, error)
}

var _ DescribeFleetAdvisorSchemaObjectSummaryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeFleetAdvisorSchemaObjectSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetAdvisorSchemaObjectSummary",
	}
}
