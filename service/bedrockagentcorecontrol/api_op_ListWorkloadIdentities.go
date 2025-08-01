// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentcorecontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all workload identities in your account.
func (c *Client) ListWorkloadIdentities(ctx context.Context, params *ListWorkloadIdentitiesInput, optFns ...func(*Options)) (*ListWorkloadIdentitiesOutput, error) {
	if params == nil {
		params = &ListWorkloadIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkloadIdentities", params, optFns, c.addOperationListWorkloadIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkloadIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkloadIdentitiesInput struct {

	// Maximum number of results to return.
	MaxResults *int32

	// Pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkloadIdentitiesOutput struct {

	// The list of workload identities.
	//
	// This member is required.
	WorkloadIdentities []types.WorkloadIdentityType

	// Pagination token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkloadIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkloadIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkloadIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkloadIdentities"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkloadIdentities(options.Region), middleware.Before); err != nil {
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

// ListWorkloadIdentitiesPaginatorOptions is the paginator options for
// ListWorkloadIdentities
type ListWorkloadIdentitiesPaginatorOptions struct {
	// Maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkloadIdentitiesPaginator is a paginator for ListWorkloadIdentities
type ListWorkloadIdentitiesPaginator struct {
	options   ListWorkloadIdentitiesPaginatorOptions
	client    ListWorkloadIdentitiesAPIClient
	params    *ListWorkloadIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListWorkloadIdentitiesPaginator returns a new ListWorkloadIdentitiesPaginator
func NewListWorkloadIdentitiesPaginator(client ListWorkloadIdentitiesAPIClient, params *ListWorkloadIdentitiesInput, optFns ...func(*ListWorkloadIdentitiesPaginatorOptions)) *ListWorkloadIdentitiesPaginator {
	if params == nil {
		params = &ListWorkloadIdentitiesInput{}
	}

	options := ListWorkloadIdentitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkloadIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkloadIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkloadIdentities page.
func (p *ListWorkloadIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkloadIdentitiesOutput, error) {
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
	result, err := p.client.ListWorkloadIdentities(ctx, &params, optFns...)
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

// ListWorkloadIdentitiesAPIClient is a client that implements the
// ListWorkloadIdentities operation.
type ListWorkloadIdentitiesAPIClient interface {
	ListWorkloadIdentities(context.Context, *ListWorkloadIdentitiesInput, ...func(*Options)) (*ListWorkloadIdentitiesOutput, error)
}

var _ ListWorkloadIdentitiesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkloadIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkloadIdentities",
	}
}
