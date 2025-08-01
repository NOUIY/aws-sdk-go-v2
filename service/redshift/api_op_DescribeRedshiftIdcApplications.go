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

// Lists the Amazon Redshift IAM Identity Center applications.
func (c *Client) DescribeRedshiftIdcApplications(ctx context.Context, params *DescribeRedshiftIdcApplicationsInput, optFns ...func(*Options)) (*DescribeRedshiftIdcApplicationsOutput, error) {
	if params == nil {
		params = &DescribeRedshiftIdcApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRedshiftIdcApplications", params, optFns, c.addOperationDescribeRedshiftIdcApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRedshiftIdcApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRedshiftIdcApplicationsInput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// The ARN for the Redshift application that integrates with IAM Identity Center.
	RedshiftIdcApplicationArn *string

	noSmithyDocumentSerde
}

type DescribeRedshiftIdcApplicationsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// The list of Amazon Redshift IAM Identity Center applications.
	RedshiftIdcApplications []types.RedshiftIdcApplication

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRedshiftIdcApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeRedshiftIdcApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeRedshiftIdcApplications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRedshiftIdcApplications"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRedshiftIdcApplications(options.Region), middleware.Before); err != nil {
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

// DescribeRedshiftIdcApplicationsPaginatorOptions is the paginator options for
// DescribeRedshiftIdcApplications
type DescribeRedshiftIdcApplicationsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRedshiftIdcApplicationsPaginator is a paginator for
// DescribeRedshiftIdcApplications
type DescribeRedshiftIdcApplicationsPaginator struct {
	options   DescribeRedshiftIdcApplicationsPaginatorOptions
	client    DescribeRedshiftIdcApplicationsAPIClient
	params    *DescribeRedshiftIdcApplicationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRedshiftIdcApplicationsPaginator returns a new
// DescribeRedshiftIdcApplicationsPaginator
func NewDescribeRedshiftIdcApplicationsPaginator(client DescribeRedshiftIdcApplicationsAPIClient, params *DescribeRedshiftIdcApplicationsInput, optFns ...func(*DescribeRedshiftIdcApplicationsPaginatorOptions)) *DescribeRedshiftIdcApplicationsPaginator {
	if params == nil {
		params = &DescribeRedshiftIdcApplicationsInput{}
	}

	options := DescribeRedshiftIdcApplicationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRedshiftIdcApplicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRedshiftIdcApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRedshiftIdcApplications page.
func (p *DescribeRedshiftIdcApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRedshiftIdcApplicationsOutput, error) {
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
	result, err := p.client.DescribeRedshiftIdcApplications(ctx, &params, optFns...)
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

// DescribeRedshiftIdcApplicationsAPIClient is a client that implements the
// DescribeRedshiftIdcApplications operation.
type DescribeRedshiftIdcApplicationsAPIClient interface {
	DescribeRedshiftIdcApplications(context.Context, *DescribeRedshiftIdcApplicationsInput, ...func(*Options)) (*DescribeRedshiftIdcApplicationsOutput, error)
}

var _ DescribeRedshiftIdcApplicationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeRedshiftIdcApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRedshiftIdcApplications",
	}
}
