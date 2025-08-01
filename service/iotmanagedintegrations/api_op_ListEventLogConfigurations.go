// Code generated by smithy-go-codegen DO NOT EDIT.

package iotmanagedintegrations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all event log configurations for an account.
func (c *Client) ListEventLogConfigurations(ctx context.Context, params *ListEventLogConfigurationsInput, optFns ...func(*Options)) (*ListEventLogConfigurationsOutput, error) {
	if params == nil {
		params = &ListEventLogConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventLogConfigurations", params, optFns, c.addOperationListEventLogConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventLogConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventLogConfigurationsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A token that can be used to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEventLogConfigurationsOutput struct {

	// A list of each event log configuration and pertinent information.
	EventLogConfigurationList []types.EventLogConfigurationSummary

	// A token that can be used to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventLogConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventLogConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventLogConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventLogConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventLogConfigurations(options.Region), middleware.Before); err != nil {
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

// ListEventLogConfigurationsPaginatorOptions is the paginator options for
// ListEventLogConfigurations
type ListEventLogConfigurationsPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventLogConfigurationsPaginator is a paginator for
// ListEventLogConfigurations
type ListEventLogConfigurationsPaginator struct {
	options   ListEventLogConfigurationsPaginatorOptions
	client    ListEventLogConfigurationsAPIClient
	params    *ListEventLogConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListEventLogConfigurationsPaginator returns a new
// ListEventLogConfigurationsPaginator
func NewListEventLogConfigurationsPaginator(client ListEventLogConfigurationsAPIClient, params *ListEventLogConfigurationsInput, optFns ...func(*ListEventLogConfigurationsPaginatorOptions)) *ListEventLogConfigurationsPaginator {
	if params == nil {
		params = &ListEventLogConfigurationsInput{}
	}

	options := ListEventLogConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventLogConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventLogConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventLogConfigurations page.
func (p *ListEventLogConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventLogConfigurationsOutput, error) {
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
	result, err := p.client.ListEventLogConfigurations(ctx, &params, optFns...)
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

// ListEventLogConfigurationsAPIClient is a client that implements the
// ListEventLogConfigurations operation.
type ListEventLogConfigurationsAPIClient interface {
	ListEventLogConfigurations(context.Context, *ListEventLogConfigurationsInput, ...func(*Options)) (*ListEventLogConfigurationsOutput, error)
}

var _ ListEventLogConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEventLogConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventLogConfigurations",
	}
}
