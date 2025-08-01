// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the Remote Desktop Services (RDS) License Server endpoints
func (c *Client) ListLicenseServerEndpoints(ctx context.Context, params *ListLicenseServerEndpointsInput, optFns ...func(*Options)) (*ListLicenseServerEndpointsOutput, error) {
	if params == nil {
		params = &ListLicenseServerEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLicenseServerEndpoints", params, optFns, c.addOperationListLicenseServerEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLicenseServerEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLicenseServerEndpointsInput struct {

	// You can use the following filters to streamline results:
	//
	//   - IdentityProviderArn
	Filters []types.Filter

	// The maximum number of results to return from a single request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLicenseServerEndpointsOutput struct {

	// An array of LicenseServerEndpoint resources that contain detailed information
	// about the RDS License Servers that meet the request criteria.
	LicenseServerEndpoints []types.LicenseServerEndpoint

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLicenseServerEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLicenseServerEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLicenseServerEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLicenseServerEndpoints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLicenseServerEndpoints(options.Region), middleware.Before); err != nil {
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

// ListLicenseServerEndpointsPaginatorOptions is the paginator options for
// ListLicenseServerEndpoints
type ListLicenseServerEndpointsPaginatorOptions struct {
	// The maximum number of results to return from a single request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLicenseServerEndpointsPaginator is a paginator for
// ListLicenseServerEndpoints
type ListLicenseServerEndpointsPaginator struct {
	options   ListLicenseServerEndpointsPaginatorOptions
	client    ListLicenseServerEndpointsAPIClient
	params    *ListLicenseServerEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListLicenseServerEndpointsPaginator returns a new
// ListLicenseServerEndpointsPaginator
func NewListLicenseServerEndpointsPaginator(client ListLicenseServerEndpointsAPIClient, params *ListLicenseServerEndpointsInput, optFns ...func(*ListLicenseServerEndpointsPaginatorOptions)) *ListLicenseServerEndpointsPaginator {
	if params == nil {
		params = &ListLicenseServerEndpointsInput{}
	}

	options := ListLicenseServerEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLicenseServerEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLicenseServerEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLicenseServerEndpoints page.
func (p *ListLicenseServerEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLicenseServerEndpointsOutput, error) {
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
	result, err := p.client.ListLicenseServerEndpoints(ctx, &params, optFns...)
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

// ListLicenseServerEndpointsAPIClient is a client that implements the
// ListLicenseServerEndpoints operation.
type ListLicenseServerEndpointsAPIClient interface {
	ListLicenseServerEndpoints(context.Context, *ListLicenseServerEndpointsInput, ...func(*Options)) (*ListLicenseServerEndpointsOutput, error)
}

var _ ListLicenseServerEndpointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLicenseServerEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLicenseServerEndpoints",
	}
}
