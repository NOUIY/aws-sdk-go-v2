// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ivsrealtime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about all storage configurations in your account, in
// the AWS region where the API request is processed.
func (c *Client) ListStorageConfigurations(ctx context.Context, params *ListStorageConfigurationsInput, optFns ...func(*Options)) (*ListStorageConfigurationsOutput, error) {
	if params == nil {
		params = &ListStorageConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStorageConfigurations", params, optFns, c.addOperationListStorageConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStorageConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStorageConfigurationsInput struct {

	// Maximum number of storage configurations to return. Default: your service quota
	// or 100, whichever is smaller.
	MaxResults *int32

	// The first storage configuration to retrieve. This is used for pagination; see
	// the nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStorageConfigurationsOutput struct {

	// List of the matching storage configurations.
	//
	// This member is required.
	StorageConfigurations []types.StorageConfigurationSummary

	// If there are more storage configurations than maxResults , use nextToken in the
	// request to get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStorageConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStorageConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStorageConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStorageConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStorageConfigurations(options.Region), middleware.Before); err != nil {
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

// ListStorageConfigurationsPaginatorOptions is the paginator options for
// ListStorageConfigurations
type ListStorageConfigurationsPaginatorOptions struct {
	// Maximum number of storage configurations to return. Default: your service quota
	// or 100, whichever is smaller.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStorageConfigurationsPaginator is a paginator for ListStorageConfigurations
type ListStorageConfigurationsPaginator struct {
	options   ListStorageConfigurationsPaginatorOptions
	client    ListStorageConfigurationsAPIClient
	params    *ListStorageConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListStorageConfigurationsPaginator returns a new
// ListStorageConfigurationsPaginator
func NewListStorageConfigurationsPaginator(client ListStorageConfigurationsAPIClient, params *ListStorageConfigurationsInput, optFns ...func(*ListStorageConfigurationsPaginatorOptions)) *ListStorageConfigurationsPaginator {
	if params == nil {
		params = &ListStorageConfigurationsInput{}
	}

	options := ListStorageConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStorageConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStorageConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStorageConfigurations page.
func (p *ListStorageConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStorageConfigurationsOutput, error) {
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
	result, err := p.client.ListStorageConfigurations(ctx, &params, optFns...)
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

// ListStorageConfigurationsAPIClient is a client that implements the
// ListStorageConfigurations operation.
type ListStorageConfigurationsAPIClient interface {
	ListStorageConfigurations(context.Context, *ListStorageConfigurationsInput, ...func(*Options)) (*ListStorageConfigurationsOutput, error)
}

var _ ListStorageConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStorageConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStorageConfigurations",
	}
}
