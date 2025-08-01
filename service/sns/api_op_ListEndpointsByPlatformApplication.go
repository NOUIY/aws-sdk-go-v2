// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the endpoints and endpoint attributes for devices in a supported push
// notification service, such as GCM (Firebase Cloud Messaging) and APNS. The
// results for ListEndpointsByPlatformApplication are paginated and return a
// limited list of endpoints, up to 100. If additional records are available after
// the first page results, then a NextToken string will be returned. To receive the
// next page, you call ListEndpointsByPlatformApplication again using the
// NextToken string received from the previous call. When there are no more records
// to return, NextToken will be null. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// This action is throttled at 30 transactions per second (TPS).
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func (c *Client) ListEndpointsByPlatformApplication(ctx context.Context, params *ListEndpointsByPlatformApplicationInput, optFns ...func(*Options)) (*ListEndpointsByPlatformApplicationOutput, error) {
	if params == nil {
		params = &ListEndpointsByPlatformApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointsByPlatformApplication", params, optFns, c.addOperationListEndpointsByPlatformApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointsByPlatformApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListEndpointsByPlatformApplication action.
type ListEndpointsByPlatformApplicationInput struct {

	// PlatformApplicationArn for ListEndpointsByPlatformApplicationInput action.
	//
	// This member is required.
	PlatformApplicationArn *string

	// NextToken string is used when calling ListEndpointsByPlatformApplication action
	// to retrieve additional records that are available after the first page results.
	NextToken *string

	noSmithyDocumentSerde
}

// Response for ListEndpointsByPlatformApplication action.
type ListEndpointsByPlatformApplicationOutput struct {

	// Endpoints returned for ListEndpointsByPlatformApplication action.
	Endpoints []types.Endpoint

	// NextToken string is returned when calling ListEndpointsByPlatformApplication
	// action if additional records are available after the first page results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEndpointsByPlatformApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListEndpointsByPlatformApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListEndpointsByPlatformApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEndpointsByPlatformApplication"); err != nil {
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
	if err = addOpListEndpointsByPlatformApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointsByPlatformApplication(options.Region), middleware.Before); err != nil {
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

// ListEndpointsByPlatformApplicationPaginatorOptions is the paginator options for
// ListEndpointsByPlatformApplication
type ListEndpointsByPlatformApplicationPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEndpointsByPlatformApplicationPaginator is a paginator for
// ListEndpointsByPlatformApplication
type ListEndpointsByPlatformApplicationPaginator struct {
	options   ListEndpointsByPlatformApplicationPaginatorOptions
	client    ListEndpointsByPlatformApplicationAPIClient
	params    *ListEndpointsByPlatformApplicationInput
	nextToken *string
	firstPage bool
}

// NewListEndpointsByPlatformApplicationPaginator returns a new
// ListEndpointsByPlatformApplicationPaginator
func NewListEndpointsByPlatformApplicationPaginator(client ListEndpointsByPlatformApplicationAPIClient, params *ListEndpointsByPlatformApplicationInput, optFns ...func(*ListEndpointsByPlatformApplicationPaginatorOptions)) *ListEndpointsByPlatformApplicationPaginator {
	if params == nil {
		params = &ListEndpointsByPlatformApplicationInput{}
	}

	options := ListEndpointsByPlatformApplicationPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEndpointsByPlatformApplicationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEndpointsByPlatformApplicationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEndpointsByPlatformApplication page.
func (p *ListEndpointsByPlatformApplicationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEndpointsByPlatformApplicationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListEndpointsByPlatformApplication(ctx, &params, optFns...)
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

// ListEndpointsByPlatformApplicationAPIClient is a client that implements the
// ListEndpointsByPlatformApplication operation.
type ListEndpointsByPlatformApplicationAPIClient interface {
	ListEndpointsByPlatformApplication(context.Context, *ListEndpointsByPlatformApplicationInput, ...func(*Options)) (*ListEndpointsByPlatformApplicationOutput, error)
}

var _ ListEndpointsByPlatformApplicationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEndpointsByPlatformApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEndpointsByPlatformApplication",
	}
}
