// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API permits filtering on the ControlId and HomeRegion fields.
func (c *Client) DescribeHomeRegionControls(ctx context.Context, params *DescribeHomeRegionControlsInput, optFns ...func(*Options)) (*DescribeHomeRegionControlsOutput, error) {
	if params == nil {
		params = &DescribeHomeRegionControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHomeRegionControls", params, optFns, c.addOperationDescribeHomeRegionControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHomeRegionControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHomeRegionControlsInput struct {

	// The ControlID is a unique identifier string of your HomeRegionControl object.
	ControlId *string

	// The name of the home region you'd like to view.
	HomeRegion *string

	// The maximum number of filtering results to display per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, more results are available. To
	// retrieve the next page of results, make the call again using the returned token
	// in NextToken .
	NextToken *string

	// The target parameter specifies the identifier to which the home region is
	// applied, which is always of type ACCOUNT . It applies the home region to the
	// current ACCOUNT .
	Target *types.Target

	noSmithyDocumentSerde
}

type DescribeHomeRegionControlsOutput struct {

	// An array that contains your HomeRegionControl objects.
	HomeRegionControls []types.HomeRegionControl

	// If a NextToken was returned by a previous call, more results are available. To
	// retrieve the next page of results, make the call again using the returned token
	// in NextToken .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHomeRegionControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHomeRegionControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHomeRegionControls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHomeRegionControls"); err != nil {
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
	if err = addOpDescribeHomeRegionControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHomeRegionControls(options.Region), middleware.Before); err != nil {
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

// DescribeHomeRegionControlsPaginatorOptions is the paginator options for
// DescribeHomeRegionControls
type DescribeHomeRegionControlsPaginatorOptions struct {
	// The maximum number of filtering results to display per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeHomeRegionControlsPaginator is a paginator for
// DescribeHomeRegionControls
type DescribeHomeRegionControlsPaginator struct {
	options   DescribeHomeRegionControlsPaginatorOptions
	client    DescribeHomeRegionControlsAPIClient
	params    *DescribeHomeRegionControlsInput
	nextToken *string
	firstPage bool
}

// NewDescribeHomeRegionControlsPaginator returns a new
// DescribeHomeRegionControlsPaginator
func NewDescribeHomeRegionControlsPaginator(client DescribeHomeRegionControlsAPIClient, params *DescribeHomeRegionControlsInput, optFns ...func(*DescribeHomeRegionControlsPaginatorOptions)) *DescribeHomeRegionControlsPaginator {
	if params == nil {
		params = &DescribeHomeRegionControlsInput{}
	}

	options := DescribeHomeRegionControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeHomeRegionControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeHomeRegionControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeHomeRegionControls page.
func (p *DescribeHomeRegionControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeHomeRegionControlsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeHomeRegionControls(ctx, &params, optFns...)
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

// DescribeHomeRegionControlsAPIClient is a client that implements the
// DescribeHomeRegionControls operation.
type DescribeHomeRegionControlsAPIClient interface {
	DescribeHomeRegionControls(context.Context, *DescribeHomeRegionControlsInput, ...func(*Options)) (*DescribeHomeRegionControlsOutput, error)
}

var _ DescribeHomeRegionControlsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeHomeRegionControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHomeRegionControls",
	}
}
