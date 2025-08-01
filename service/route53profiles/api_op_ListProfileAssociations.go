// Code generated by smithy-go-codegen DO NOT EDIT.

package route53profiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53profiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the VPCs that the specified Route 53 Profile is associated with.
func (c *Client) ListProfileAssociations(ctx context.Context, params *ListProfileAssociationsInput, optFns ...func(*Options)) (*ListProfileAssociationsOutput, error) {
	if params == nil {
		params = &ListProfileAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileAssociations", params, optFns, c.addOperationListProfileAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfileAssociationsInput struct {

	//  The maximum number of objects that you want to return for this request. If
	// more objects are available, in the response, a NextToken value, which you can
	// use in a subsequent call to get the next batch of objects, is provided.
	//
	// If you don't specify a value for MaxResults , up to 100 objects are returned.
	MaxResults *int32

	//  For the first call to this list request, omit this value.
	//
	// When you request a list of objects, at most the number of objects specified by
	// MaxResults is returned. If more objects are available for retrieval, a NextToken
	// value is returned in the response. To retrieve the next batch of objects, use
	// the token that was returned for the prior request in your next request.
	NextToken *string

	//  ID of the Profile.
	ProfileId *string

	//  ID of the VPC.
	ResourceId *string

	noSmithyDocumentSerde
}

type ListProfileAssociationsOutput struct {

	//  If more than MaxResults profile associations match the specified criteria, you
	// can submit another ListProfileAssociations request to get the next group of
	// results. In the next request, specify the value of NextToken from the previous
	// response.
	NextToken *string

	//  A complex type that containts settings information about the profile's VPC
	// associations.
	ProfileAssociations []types.ProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProfileAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileAssociations(options.Region), middleware.Before); err != nil {
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

// ListProfileAssociationsPaginatorOptions is the paginator options for
// ListProfileAssociations
type ListProfileAssociationsPaginatorOptions struct {
	//  The maximum number of objects that you want to return for this request. If
	// more objects are available, in the response, a NextToken value, which you can
	// use in a subsequent call to get the next batch of objects, is provided.
	//
	// If you don't specify a value for MaxResults , up to 100 objects are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfileAssociationsPaginator is a paginator for ListProfileAssociations
type ListProfileAssociationsPaginator struct {
	options   ListProfileAssociationsPaginatorOptions
	client    ListProfileAssociationsAPIClient
	params    *ListProfileAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListProfileAssociationsPaginator returns a new
// ListProfileAssociationsPaginator
func NewListProfileAssociationsPaginator(client ListProfileAssociationsAPIClient, params *ListProfileAssociationsInput, optFns ...func(*ListProfileAssociationsPaginatorOptions)) *ListProfileAssociationsPaginator {
	if params == nil {
		params = &ListProfileAssociationsInput{}
	}

	options := ListProfileAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfileAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfileAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfileAssociations page.
func (p *ListProfileAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfileAssociationsOutput, error) {
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
	result, err := p.client.ListProfileAssociations(ctx, &params, optFns...)
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

// ListProfileAssociationsAPIClient is a client that implements the
// ListProfileAssociations operation.
type ListProfileAssociationsAPIClient interface {
	ListProfileAssociations(context.Context, *ListProfileAssociationsInput, ...func(*Options)) (*ListProfileAssociationsOutput, error)
}

var _ ListProfileAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListProfileAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProfileAssociations",
	}
}
