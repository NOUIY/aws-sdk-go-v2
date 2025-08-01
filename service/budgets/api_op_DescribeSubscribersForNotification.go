// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the subscribers that are associated with a notification.
func (c *Client) DescribeSubscribersForNotification(ctx context.Context, params *DescribeSubscribersForNotificationInput, optFns ...func(*Options)) (*DescribeSubscribersForNotificationOutput, error) {
	if params == nil {
		params = &DescribeSubscribersForNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSubscribersForNotification", params, optFns, c.addOperationDescribeSubscribersForNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSubscribersForNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationInput struct {

	// The accountId that is associated with the budget whose subscribers you want
	// descriptions of.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose subscribers you want descriptions of.
	//
	// This member is required.
	BudgetName *string

	// The notification whose subscribers you want to list.
	//
	// This member is required.
	Notification *types.Notification

	// An optional integer that represents how many entries a paginated response
	// contains.
	MaxResults *int32

	// The pagination token that you include in your request to indicate the next set
	// of results that you want to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

// Response of DescribeSubscribersForNotification
type DescribeSubscribersForNotificationOutput struct {

	// The pagination token in the service response that indicates the next set of
	// results that you can retrieve.
	NextToken *string

	// A list of subscribers that are associated with a notification.
	Subscribers []types.Subscriber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSubscribersForNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSubscribersForNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSubscribersForNotification{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSubscribersForNotification"); err != nil {
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
	if err = addOpDescribeSubscribersForNotificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubscribersForNotification(options.Region), middleware.Before); err != nil {
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

// DescribeSubscribersForNotificationPaginatorOptions is the paginator options for
// DescribeSubscribersForNotification
type DescribeSubscribersForNotificationPaginatorOptions struct {
	// An optional integer that represents how many entries a paginated response
	// contains.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSubscribersForNotificationPaginator is a paginator for
// DescribeSubscribersForNotification
type DescribeSubscribersForNotificationPaginator struct {
	options   DescribeSubscribersForNotificationPaginatorOptions
	client    DescribeSubscribersForNotificationAPIClient
	params    *DescribeSubscribersForNotificationInput
	nextToken *string
	firstPage bool
}

// NewDescribeSubscribersForNotificationPaginator returns a new
// DescribeSubscribersForNotificationPaginator
func NewDescribeSubscribersForNotificationPaginator(client DescribeSubscribersForNotificationAPIClient, params *DescribeSubscribersForNotificationInput, optFns ...func(*DescribeSubscribersForNotificationPaginatorOptions)) *DescribeSubscribersForNotificationPaginator {
	if params == nil {
		params = &DescribeSubscribersForNotificationInput{}
	}

	options := DescribeSubscribersForNotificationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSubscribersForNotificationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSubscribersForNotificationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSubscribersForNotification page.
func (p *DescribeSubscribersForNotificationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSubscribersForNotificationOutput, error) {
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
	result, err := p.client.DescribeSubscribersForNotification(ctx, &params, optFns...)
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

// DescribeSubscribersForNotificationAPIClient is a client that implements the
// DescribeSubscribersForNotification operation.
type DescribeSubscribersForNotificationAPIClient interface {
	DescribeSubscribersForNotification(context.Context, *DescribeSubscribersForNotificationInput, ...func(*Options)) (*DescribeSubscribersForNotificationOutput, error)
}

var _ DescribeSubscribersForNotificationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeSubscribersForNotification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSubscribersForNotification",
	}
}
