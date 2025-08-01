// Code generated by smithy-go-codegen DO NOT EDIT.

package s3outposts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all endpoints associated with an Outpost that has been shared by Amazon
// Web Services Resource Access Manager (RAM).
//
// Related actions include:
//
// [CreateEndpoint]
//
// [DeleteEndpoint]
//
// [CreateEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html
// [DeleteEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html
func (c *Client) ListSharedEndpoints(ctx context.Context, params *ListSharedEndpointsInput, optFns ...func(*Options)) (*ListSharedEndpointsOutput, error) {
	if params == nil {
		params = &ListSharedEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSharedEndpoints", params, optFns, c.addOperationListSharedEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSharedEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSharedEndpointsInput struct {

	// The ID of the Amazon Web Services Outpost.
	//
	// This member is required.
	OutpostId *string

	// The maximum number of endpoints that will be returned in the response.
	MaxResults int32

	// If a previous response from this operation included a NextToken value, you can
	// provide that value here to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSharedEndpointsOutput struct {

	// The list of endpoints associated with the specified Outpost that have been
	// shared by Amazon Web Services Resource Access Manager (RAM).
	Endpoints []types.Endpoint

	// If the number of endpoints associated with the specified Outpost exceeds
	// MaxResults , you can include this value in subsequent calls to this operation to
	// retrieve more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSharedEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSharedEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSharedEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSharedEndpoints"); err != nil {
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
	if err = addOpListSharedEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSharedEndpoints(options.Region), middleware.Before); err != nil {
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

// ListSharedEndpointsPaginatorOptions is the paginator options for
// ListSharedEndpoints
type ListSharedEndpointsPaginatorOptions struct {
	// The maximum number of endpoints that will be returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSharedEndpointsPaginator is a paginator for ListSharedEndpoints
type ListSharedEndpointsPaginator struct {
	options   ListSharedEndpointsPaginatorOptions
	client    ListSharedEndpointsAPIClient
	params    *ListSharedEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListSharedEndpointsPaginator returns a new ListSharedEndpointsPaginator
func NewListSharedEndpointsPaginator(client ListSharedEndpointsAPIClient, params *ListSharedEndpointsInput, optFns ...func(*ListSharedEndpointsPaginatorOptions)) *ListSharedEndpointsPaginator {
	if params == nil {
		params = &ListSharedEndpointsInput{}
	}

	options := ListSharedEndpointsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSharedEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSharedEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSharedEndpoints page.
func (p *ListSharedEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSharedEndpointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListSharedEndpoints(ctx, &params, optFns...)
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

// ListSharedEndpointsAPIClient is a client that implements the
// ListSharedEndpoints operation.
type ListSharedEndpointsAPIClient interface {
	ListSharedEndpoints(context.Context, *ListSharedEndpointsInput, ...func(*Options)) (*ListSharedEndpointsOutput, error)
}

var _ ListSharedEndpointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSharedEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSharedEndpoints",
	}
}
