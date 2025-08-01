// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List the hours of operation overrides.
func (c *Client) ListHoursOfOperationOverrides(ctx context.Context, params *ListHoursOfOperationOverridesInput, optFns ...func(*Options)) (*ListHoursOfOperationOverridesOutput, error) {
	if params == nil {
		params = &ListHoursOfOperationOverridesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHoursOfOperationOverrides", params, optFns, c.addOperationListHoursOfOperationOverridesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHoursOfOperationOverridesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHoursOfOperationOverridesInput struct {

	// The identifier for the hours of operation
	//
	// This member is required.
	HoursOfOperationId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100. Valid Range: Minimum value of 1. Maximum value of 1000.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHoursOfOperationOverridesOutput struct {

	// Information about the hours of operation override.
	HoursOfOperationOverrideList []types.HoursOfOperationOverride

	// The AWS Region where this resource was last modified.
	LastModifiedRegion *string

	// The timestamp when this resource was last modified.
	LastModifiedTime *time.Time

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHoursOfOperationOverridesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHoursOfOperationOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHoursOfOperationOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHoursOfOperationOverrides"); err != nil {
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
	if err = addOpListHoursOfOperationOverridesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHoursOfOperationOverrides(options.Region), middleware.Before); err != nil {
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

// ListHoursOfOperationOverridesPaginatorOptions is the paginator options for
// ListHoursOfOperationOverrides
type ListHoursOfOperationOverridesPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100. Valid Range: Minimum value of 1. Maximum value of 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHoursOfOperationOverridesPaginator is a paginator for
// ListHoursOfOperationOverrides
type ListHoursOfOperationOverridesPaginator struct {
	options   ListHoursOfOperationOverridesPaginatorOptions
	client    ListHoursOfOperationOverridesAPIClient
	params    *ListHoursOfOperationOverridesInput
	nextToken *string
	firstPage bool
}

// NewListHoursOfOperationOverridesPaginator returns a new
// ListHoursOfOperationOverridesPaginator
func NewListHoursOfOperationOverridesPaginator(client ListHoursOfOperationOverridesAPIClient, params *ListHoursOfOperationOverridesInput, optFns ...func(*ListHoursOfOperationOverridesPaginatorOptions)) *ListHoursOfOperationOverridesPaginator {
	if params == nil {
		params = &ListHoursOfOperationOverridesInput{}
	}

	options := ListHoursOfOperationOverridesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHoursOfOperationOverridesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHoursOfOperationOverridesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHoursOfOperationOverrides page.
func (p *ListHoursOfOperationOverridesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHoursOfOperationOverridesOutput, error) {
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
	result, err := p.client.ListHoursOfOperationOverrides(ctx, &params, optFns...)
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

// ListHoursOfOperationOverridesAPIClient is a client that implements the
// ListHoursOfOperationOverrides operation.
type ListHoursOfOperationOverridesAPIClient interface {
	ListHoursOfOperationOverrides(context.Context, *ListHoursOfOperationOverridesInput, ...func(*Options)) (*ListHoursOfOperationOverridesOutput, error)
}

var _ ListHoursOfOperationOverridesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListHoursOfOperationOverrides(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHoursOfOperationOverrides",
	}
}
