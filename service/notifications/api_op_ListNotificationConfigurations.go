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

// Returns a list of abbreviated NotificationConfigurations according to specified
// filters, in reverse chronological order (newest first).
func (c *Client) ListNotificationConfigurations(ctx context.Context, params *ListNotificationConfigurationsInput, optFns ...func(*Options)) (*ListNotificationConfigurationsOutput, error) {
	if params == nil {
		params = &ListNotificationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotificationConfigurations", params, optFns, c.addOperationListNotificationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotificationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotificationConfigurationsInput struct {

	// The Amazon Resource Name (ARN) of the Channel to match.
	ChannelArn *string

	// The matched event source.
	//
	// Must match one of the valid EventBridge sources. Only Amazon Web Services
	// service sourced events are supported. For example, aws.ec2 and aws.cloudwatch .
	// For more information, see [Event delivery from Amazon Web Services services]in the Amazon EventBridge User Guide.
	//
	// [Event delivery from Amazon Web Services services]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-service-event.html#eb-service-event-delivery-level
	EventRuleSource *string

	// The maximum number of results to be returned in this call. Defaults to 20.
	MaxResults *int32

	// The start token for paginated calls. Retrieved from the response of a previous
	// ListEventRules call. Next token uses Base64 encoding.
	NextToken *string

	// The NotificationConfiguration status to match.
	//
	//   - Values:
	//
	//   - ACTIVE
	//
	//   - All EventRules are ACTIVE and any call can be run.
	//
	//   - PARTIALLY_ACTIVE
	//
	//   - Some EventRules are ACTIVE and some are INACTIVE . Any call can be run.
	//
	//   - Any call can be run.
	//
	//   - INACTIVE
	//
	//   - All EventRules are INACTIVE and any call can be run.
	//
	//   - DELETING
	//
	//   - This NotificationConfiguration is being deleted.
	//
	//   - Only GET and LIST calls can be run.
	Status types.NotificationConfigurationStatus

	noSmithyDocumentSerde
}

type ListNotificationConfigurationsOutput struct {

	// The NotificationConfigurations in the account.
	//
	// This member is required.
	NotificationConfigurations []types.NotificationConfigurationStructure

	// A pagination token. If a non-null pagination token is returned in a result,
	// pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotificationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNotificationConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotificationConfigurations(options.Region), middleware.Before); err != nil {
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

// ListNotificationConfigurationsPaginatorOptions is the paginator options for
// ListNotificationConfigurations
type ListNotificationConfigurationsPaginatorOptions struct {
	// The maximum number of results to be returned in this call. Defaults to 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotificationConfigurationsPaginator is a paginator for
// ListNotificationConfigurations
type ListNotificationConfigurationsPaginator struct {
	options   ListNotificationConfigurationsPaginatorOptions
	client    ListNotificationConfigurationsAPIClient
	params    *ListNotificationConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListNotificationConfigurationsPaginator returns a new
// ListNotificationConfigurationsPaginator
func NewListNotificationConfigurationsPaginator(client ListNotificationConfigurationsAPIClient, params *ListNotificationConfigurationsInput, optFns ...func(*ListNotificationConfigurationsPaginatorOptions)) *ListNotificationConfigurationsPaginator {
	if params == nil {
		params = &ListNotificationConfigurationsInput{}
	}

	options := ListNotificationConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotificationConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotificationConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotificationConfigurations page.
func (p *ListNotificationConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotificationConfigurationsOutput, error) {
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
	result, err := p.client.ListNotificationConfigurations(ctx, &params, optFns...)
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

// ListNotificationConfigurationsAPIClient is a client that implements the
// ListNotificationConfigurations operation.
type ListNotificationConfigurationsAPIClient interface {
	ListNotificationConfigurations(context.Context, *ListNotificationConfigurationsInput, ...func(*Options)) (*ListNotificationConfigurationsOutput, error)
}

var _ ListNotificationConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListNotificationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNotificationConfigurations",
	}
}
