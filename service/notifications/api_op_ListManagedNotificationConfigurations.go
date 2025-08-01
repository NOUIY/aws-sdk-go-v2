// Code generated by smithy-go-codegen DO NOT EDIT.

package notifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/notifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Managed Notification Configurations according to specified
// filters, ordered by creation time in reverse chronological order (newest first).
func (c *Client) ListManagedNotificationConfigurations(ctx context.Context, params *ListManagedNotificationConfigurationsInput, optFns ...func(*Options)) (*ListManagedNotificationConfigurationsOutput, error) {
	if params == nil {
		params = &ListManagedNotificationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedNotificationConfigurations", params, optFns, c.addOperationListManagedNotificationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedNotificationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedNotificationConfigurationsInput struct {

	// The identifier or ARN of the notification channel to filter configurations by.
	ChannelIdentifier *string

	// The maximum number of results to be returned in this call. Defaults to 20.
	MaxResults *int32

	// The start token for paginated calls. Retrieved from the response of a previous
	// ListManagedNotificationChannelAssociations call. Next token uses Base64
	// encoding.
	NextToken *string

	noSmithyDocumentSerde
}

type ListManagedNotificationConfigurationsOutput struct {

	// A list of Managed Notification Configurations matching the request criteria.
	//
	// This member is required.
	ManagedNotificationConfigurations []types.ManagedNotificationConfigurationStructure

	// A pagination token. If a non-null pagination token is returned in a result,
	// pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedNotificationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListManagedNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListManagedNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListManagedNotificationConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedNotificationConfigurations(options.Region), middleware.Before); err != nil {
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

// ListManagedNotificationConfigurationsPaginatorOptions is the paginator options
// for ListManagedNotificationConfigurations
type ListManagedNotificationConfigurationsPaginatorOptions struct {
	// The maximum number of results to be returned in this call. Defaults to 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedNotificationConfigurationsPaginator is a paginator for
// ListManagedNotificationConfigurations
type ListManagedNotificationConfigurationsPaginator struct {
	options   ListManagedNotificationConfigurationsPaginatorOptions
	client    ListManagedNotificationConfigurationsAPIClient
	params    *ListManagedNotificationConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListManagedNotificationConfigurationsPaginator returns a new
// ListManagedNotificationConfigurationsPaginator
func NewListManagedNotificationConfigurationsPaginator(client ListManagedNotificationConfigurationsAPIClient, params *ListManagedNotificationConfigurationsInput, optFns ...func(*ListManagedNotificationConfigurationsPaginatorOptions)) *ListManagedNotificationConfigurationsPaginator {
	if params == nil {
		params = &ListManagedNotificationConfigurationsInput{}
	}

	options := ListManagedNotificationConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedNotificationConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedNotificationConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedNotificationConfigurations page.
func (p *ListManagedNotificationConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedNotificationConfigurationsOutput, error) {
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
	result, err := p.client.ListManagedNotificationConfigurations(ctx, &params, optFns...)
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

// ListManagedNotificationConfigurationsAPIClient is a client that implements the
// ListManagedNotificationConfigurations operation.
type ListManagedNotificationConfigurationsAPIClient interface {
	ListManagedNotificationConfigurations(context.Context, *ListManagedNotificationConfigurationsInput, ...func(*Options)) (*ListManagedNotificationConfigurationsOutput, error)
}

var _ ListManagedNotificationConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListManagedNotificationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListManagedNotificationConfigurations",
	}
}
