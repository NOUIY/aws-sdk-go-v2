// Code generated by smithy-go-codegen DO NOT EDIT.

package codeconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of sync configurations for a specified repository.
func (c *Client) ListSyncConfigurations(ctx context.Context, params *ListSyncConfigurationsInput, optFns ...func(*Options)) (*ListSyncConfigurationsOutput, error) {
	if params == nil {
		params = &ListSyncConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSyncConfigurations", params, optFns, c.addOperationListSyncConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSyncConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSyncConfigurationsInput struct {

	// The ID of the repository link for the requested list of sync configurations.
	//
	// This member is required.
	RepositoryLinkId *string

	// The sync type for the requested list of sync configurations.
	//
	// This member is required.
	SyncType types.SyncConfigurationType

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults int32

	// An enumeration token that allows the operation to batch the results of the
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSyncConfigurationsOutput struct {

	// The list of repository sync definitions returned by the request.
	//
	// This member is required.
	SyncConfigurations []types.SyncConfiguration

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSyncConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListSyncConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListSyncConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSyncConfigurations"); err != nil {
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
	if err = addOpListSyncConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSyncConfigurations(options.Region), middleware.Before); err != nil {
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

// ListSyncConfigurationsPaginatorOptions is the paginator options for
// ListSyncConfigurations
type ListSyncConfigurationsPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSyncConfigurationsPaginator is a paginator for ListSyncConfigurations
type ListSyncConfigurationsPaginator struct {
	options   ListSyncConfigurationsPaginatorOptions
	client    ListSyncConfigurationsAPIClient
	params    *ListSyncConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListSyncConfigurationsPaginator returns a new ListSyncConfigurationsPaginator
func NewListSyncConfigurationsPaginator(client ListSyncConfigurationsAPIClient, params *ListSyncConfigurationsInput, optFns ...func(*ListSyncConfigurationsPaginatorOptions)) *ListSyncConfigurationsPaginator {
	if params == nil {
		params = &ListSyncConfigurationsInput{}
	}

	options := ListSyncConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSyncConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSyncConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSyncConfigurations page.
func (p *ListSyncConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSyncConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListSyncConfigurations(ctx, &params, optFns...)
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

// ListSyncConfigurationsAPIClient is a client that implements the
// ListSyncConfigurations operation.
type ListSyncConfigurationsAPIClient interface {
	ListSyncConfigurations(context.Context, *ListSyncConfigurationsInput, ...func(*Options)) (*ListSyncConfigurationsOutput, error)
}

var _ ListSyncConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSyncConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSyncConfigurations",
	}
}
