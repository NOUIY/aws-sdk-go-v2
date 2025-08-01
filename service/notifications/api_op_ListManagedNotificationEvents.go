// Code generated by smithy-go-codegen DO NOT EDIT.

package notifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/notifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of Managed Notification Events according to specified filters,
// ordered by creation time in reverse chronological order (newest first).
func (c *Client) ListManagedNotificationEvents(ctx context.Context, params *ListManagedNotificationEventsInput, optFns ...func(*Options)) (*ListManagedNotificationEventsOutput, error) {
	if params == nil {
		params = &ListManagedNotificationEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedNotificationEvents", params, optFns, c.addOperationListManagedNotificationEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedNotificationEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedNotificationEventsInput struct {

	// Latest time of events to return from this call.
	EndTime *time.Time

	// The locale code of the language used for the retrieved NotificationEvent. The
	// default locale is English (en_US).
	Locale types.LocaleCode

	// The maximum number of results to be returned in this call. Defaults to 20.
	MaxResults *int32

	// The start token for paginated calls. Retrieved from the response of a previous
	// ListManagedNotificationChannelAssociations call. Next token uses Base64 encoding.
	NextToken *string

	// The Organizational Unit Id that an Amazon Web Services account belongs to.
	OrganizationalUnitId *string

	// The Amazon Web Services account ID associated with the Managed Notification
	// Events.
	RelatedAccount *string

	// The Amazon Web Services service the event originates from. For example
	// aws.cloudwatch.
	Source *string

	// The earliest time of events to return from this call.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListManagedNotificationEventsOutput struct {

	// A list of Managed Notification Events matching the request criteria.
	//
	// This member is required.
	ManagedNotificationEvents []types.ManagedNotificationEventOverview

	// A pagination token. If a non-null pagination token is returned in a result,
	// pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedNotificationEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListManagedNotificationEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListManagedNotificationEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListManagedNotificationEvents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedNotificationEvents(options.Region), middleware.Before); err != nil {
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

// ListManagedNotificationEventsPaginatorOptions is the paginator options for
// ListManagedNotificationEvents
type ListManagedNotificationEventsPaginatorOptions struct {
	// The maximum number of results to be returned in this call. Defaults to 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedNotificationEventsPaginator is a paginator for
// ListManagedNotificationEvents
type ListManagedNotificationEventsPaginator struct {
	options   ListManagedNotificationEventsPaginatorOptions
	client    ListManagedNotificationEventsAPIClient
	params    *ListManagedNotificationEventsInput
	nextToken *string
	firstPage bool
}

// NewListManagedNotificationEventsPaginator returns a new
// ListManagedNotificationEventsPaginator
func NewListManagedNotificationEventsPaginator(client ListManagedNotificationEventsAPIClient, params *ListManagedNotificationEventsInput, optFns ...func(*ListManagedNotificationEventsPaginatorOptions)) *ListManagedNotificationEventsPaginator {
	if params == nil {
		params = &ListManagedNotificationEventsInput{}
	}

	options := ListManagedNotificationEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedNotificationEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedNotificationEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedNotificationEvents page.
func (p *ListManagedNotificationEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedNotificationEventsOutput, error) {
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
	result, err := p.client.ListManagedNotificationEvents(ctx, &params, optFns...)
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

// ListManagedNotificationEventsAPIClient is a client that implements the
// ListManagedNotificationEvents operation.
type ListManagedNotificationEventsAPIClient interface {
	ListManagedNotificationEvents(context.Context, *ListManagedNotificationEventsInput, ...func(*Options)) (*ListManagedNotificationEventsOutput, error)
}

var _ ListManagedNotificationEventsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListManagedNotificationEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListManagedNotificationEvents",
	}
}
