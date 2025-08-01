// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all signing platforms available in AWS Signer that match the request
// parameters. If additional jobs remain to be listed, Signer returns a nextToken
// value. Use this value in subsequent calls to ListSigningJobs to fetch the
// remaining values. You can continue calling ListSigningJobs with your maxResults
// parameter and with new values that Signer returns in the nextToken parameter
// until all of your signing jobs have been returned.
func (c *Client) ListSigningPlatforms(ctx context.Context, params *ListSigningPlatformsInput, optFns ...func(*Options)) (*ListSigningPlatformsOutput, error) {
	if params == nil {
		params = &ListSigningPlatformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSigningPlatforms", params, optFns, c.addOperationListSigningPlatformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSigningPlatformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningPlatformsInput struct {

	// The category type of a signing platform.
	Category *string

	// The maximum number of results to be returned by this operation.
	MaxResults *int32

	// Value for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string

	// Any partner entities connected to a signing platform.
	Partner *string

	// The validation template that is used by the target signing platform.
	Target *string

	noSmithyDocumentSerde
}

type ListSigningPlatformsOutput struct {

	// Value for specifying the next set of paginated results to return.
	NextToken *string

	// A list of all platforms that match the request parameters.
	Platforms []types.SigningPlatform

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSigningPlatformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSigningPlatforms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSigningPlatforms{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSigningPlatforms"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningPlatforms(options.Region), middleware.Before); err != nil {
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

// ListSigningPlatformsPaginatorOptions is the paginator options for
// ListSigningPlatforms
type ListSigningPlatformsPaginatorOptions struct {
	// The maximum number of results to be returned by this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSigningPlatformsPaginator is a paginator for ListSigningPlatforms
type ListSigningPlatformsPaginator struct {
	options   ListSigningPlatformsPaginatorOptions
	client    ListSigningPlatformsAPIClient
	params    *ListSigningPlatformsInput
	nextToken *string
	firstPage bool
}

// NewListSigningPlatformsPaginator returns a new ListSigningPlatformsPaginator
func NewListSigningPlatformsPaginator(client ListSigningPlatformsAPIClient, params *ListSigningPlatformsInput, optFns ...func(*ListSigningPlatformsPaginatorOptions)) *ListSigningPlatformsPaginator {
	if params == nil {
		params = &ListSigningPlatformsInput{}
	}

	options := ListSigningPlatformsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSigningPlatformsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSigningPlatformsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSigningPlatforms page.
func (p *ListSigningPlatformsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSigningPlatformsOutput, error) {
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
	result, err := p.client.ListSigningPlatforms(ctx, &params, optFns...)
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

// ListSigningPlatformsAPIClient is a client that implements the
// ListSigningPlatforms operation.
type ListSigningPlatformsAPIClient interface {
	ListSigningPlatforms(context.Context, *ListSigningPlatformsInput, ...func(*Options)) (*ListSigningPlatformsOutput, error)
}

var _ ListSigningPlatformsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSigningPlatforms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSigningPlatforms",
	}
}
