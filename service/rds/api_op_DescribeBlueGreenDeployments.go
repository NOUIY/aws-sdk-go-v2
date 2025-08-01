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

// Describes one or more blue/green deployments.
//
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
func (c *Client) DescribeBlueGreenDeployments(ctx context.Context, params *DescribeBlueGreenDeploymentsInput, optFns ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error) {
	if params == nil {
		params = &DescribeBlueGreenDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBlueGreenDeployments", params, optFns, c.addOperationDescribeBlueGreenDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBlueGreenDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBlueGreenDeploymentsInput struct {

	// The blue/green deployment identifier. If you specify this parameter, the
	// response only includes information about the specific blue/green deployment.
	// This parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//   - Must match an existing blue/green deployment identifier.
	BlueGreenDeploymentIdentifier *string

	// A filter that specifies one or more blue/green deployments to describe.
	//
	// Valid Values:
	//
	//   - blue-green-deployment-identifier - Accepts system-generated identifiers for
	//   blue/green deployments. The results list only includes information about the
	//   blue/green deployments with the specified identifiers.
	//
	//   - blue-green-deployment-name - Accepts user-supplied names for blue/green
	//   deployments. The results list only includes information about the blue/green
	//   deployments with the specified names.
	//
	//   - source - Accepts source databases for a blue/green deployment. The results
	//   list only includes information about the blue/green deployments with the
	//   specified source databases.
	//
	//   - target - Accepts target databases for a blue/green deployment. The results
	//   list only includes information about the blue/green deployments with the
	//   specified target databases.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeBlueGreenDeployments
	// request. If you specify this parameter, the response only includes records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints:
	//
	//   - Must be a minimum of 20.
	//
	//   - Can't exceed 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeBlueGreenDeploymentsOutput struct {

	// A list of blue/green deployments in the current account and Amazon Web Services
	// Region.
	BlueGreenDeployments []types.BlueGreenDeployment

	// A pagination token that can be used in a later DescribeBlueGreenDeployments
	// request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBlueGreenDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeBlueGreenDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeBlueGreenDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBlueGreenDeployments"); err != nil {
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
	if err = addOpDescribeBlueGreenDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBlueGreenDeployments(options.Region), middleware.Before); err != nil {
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

// DescribeBlueGreenDeploymentsPaginatorOptions is the paginator options for
// DescribeBlueGreenDeployments
type DescribeBlueGreenDeploymentsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints:
	//
	//   - Must be a minimum of 20.
	//
	//   - Can't exceed 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBlueGreenDeploymentsPaginator is a paginator for
// DescribeBlueGreenDeployments
type DescribeBlueGreenDeploymentsPaginator struct {
	options   DescribeBlueGreenDeploymentsPaginatorOptions
	client    DescribeBlueGreenDeploymentsAPIClient
	params    *DescribeBlueGreenDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBlueGreenDeploymentsPaginator returns a new
// DescribeBlueGreenDeploymentsPaginator
func NewDescribeBlueGreenDeploymentsPaginator(client DescribeBlueGreenDeploymentsAPIClient, params *DescribeBlueGreenDeploymentsInput, optFns ...func(*DescribeBlueGreenDeploymentsPaginatorOptions)) *DescribeBlueGreenDeploymentsPaginator {
	if params == nil {
		params = &DescribeBlueGreenDeploymentsInput{}
	}

	options := DescribeBlueGreenDeploymentsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBlueGreenDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBlueGreenDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBlueGreenDeployments page.
func (p *DescribeBlueGreenDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error) {
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
	result, err := p.client.DescribeBlueGreenDeployments(ctx, &params, optFns...)
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

// DescribeBlueGreenDeploymentsAPIClient is a client that implements the
// DescribeBlueGreenDeployments operation.
type DescribeBlueGreenDeploymentsAPIClient interface {
	DescribeBlueGreenDeployments(context.Context, *DescribeBlueGreenDeploymentsInput, ...func(*Options)) (*DescribeBlueGreenDeploymentsOutput, error)
}

var _ DescribeBlueGreenDeploymentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeBlueGreenDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBlueGreenDeployments",
	}
}
