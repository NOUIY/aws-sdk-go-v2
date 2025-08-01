// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Amazon Redshift parameter groups, including parameter groups
// you created and the default parameter group. For each parameter group, the
// response includes the parameter group name, description, and parameter group
// family name. You can optionally specify a name to retrieve the description of a
// specific parameter group.
//
// For more information about parameters and parameter groups, go to [Amazon Redshift Parameter Groups] in the
// Amazon Redshift Cluster Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all parameter groups that match any combination of the
// specified keys and values. For example, if you have owner and environment for
// tag keys, and admin and test for tag values, all parameter groups that have any
// combination of those values are returned.
//
// If both tag keys and values are omitted from the request, parameter groups are
// returned regardless of whether they have tag keys or values associated with
// them.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func (c *Client) DescribeClusterParameterGroups(ctx context.Context, params *DescribeClusterParameterGroupsInput, optFns ...func(*Options)) (*DescribeClusterParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeClusterParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterParameterGroups", params, optFns, c.addOperationDescribeClusterParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterParameterGroupsInput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterParameterGroupsrequest exceed the value specified in
	// MaxRecords , Amazon Web Services returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The name of a specific parameter group for which to return details. By default,
	// details about all parameter groups and the default parameter group are returned.
	ParameterGroupName *string

	// A tag key or keys for which you want to return all matching cluster parameter
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have parameter groups that are tagged with keys called owner and
	// environment . If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the parameter groups that have either or both
	// of these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching cluster
	// parameter groups that are associated with the specified tag value or values. For
	// example, suppose that you have parameter groups that are tagged with values
	// called admin and test . If you specify both of these tag values in the request,
	// Amazon Redshift returns a response with the parameter groups that have either or
	// both of these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

// Contains the output from the DescribeClusterParameterGroups action.
type DescribeClusterParameterGroupsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of ClusterParameterGroup instances. Each instance describes one cluster parameter group.
	ParameterGroups []types.ClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClusterParameterGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterParameterGroups(options.Region), middleware.Before); err != nil {
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

// DescribeClusterParameterGroupsPaginatorOptions is the paginator options for
// DescribeClusterParameterGroups
type DescribeClusterParameterGroupsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterParameterGroupsPaginator is a paginator for
// DescribeClusterParameterGroups
type DescribeClusterParameterGroupsPaginator struct {
	options   DescribeClusterParameterGroupsPaginatorOptions
	client    DescribeClusterParameterGroupsAPIClient
	params    *DescribeClusterParameterGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterParameterGroupsPaginator returns a new
// DescribeClusterParameterGroupsPaginator
func NewDescribeClusterParameterGroupsPaginator(client DescribeClusterParameterGroupsAPIClient, params *DescribeClusterParameterGroupsInput, optFns ...func(*DescribeClusterParameterGroupsPaginatorOptions)) *DescribeClusterParameterGroupsPaginator {
	if params == nil {
		params = &DescribeClusterParameterGroupsInput{}
	}

	options := DescribeClusterParameterGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterParameterGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterParameterGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusterParameterGroups page.
func (p *DescribeClusterParameterGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterParameterGroupsOutput, error) {
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
	result, err := p.client.DescribeClusterParameterGroups(ctx, &params, optFns...)
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

// DescribeClusterParameterGroupsAPIClient is a client that implements the
// DescribeClusterParameterGroups operation.
type DescribeClusterParameterGroupsAPIClient interface {
	DescribeClusterParameterGroups(context.Context, *DescribeClusterParameterGroupsInput, ...func(*Options)) (*DescribeClusterParameterGroupsOutput, error)
}

var _ DescribeClusterParameterGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeClusterParameterGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClusterParameterGroups",
	}
}
