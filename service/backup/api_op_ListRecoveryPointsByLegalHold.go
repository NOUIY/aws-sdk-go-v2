// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action returns recovery point ARNs (Amazon Resource Names) of the
// specified legal hold.
func (c *Client) ListRecoveryPointsByLegalHold(ctx context.Context, params *ListRecoveryPointsByLegalHoldInput, optFns ...func(*Options)) (*ListRecoveryPointsByLegalHoldOutput, error) {
	if params == nil {
		params = &ListRecoveryPointsByLegalHoldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecoveryPointsByLegalHold", params, optFns, c.addOperationListRecoveryPointsByLegalHoldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecoveryPointsByLegalHoldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecoveryPointsByLegalHoldInput struct {

	// The ID of the legal hold.
	//
	// This member is required.
	LegalHoldId *string

	// The maximum number of resource list items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return MaxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecoveryPointsByLegalHoldOutput struct {

	// The next item following a partial list of returned resources.
	NextToken *string

	// The recovery points.
	RecoveryPoints []types.RecoveryPointMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecoveryPointsByLegalHoldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecoveryPointsByLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecoveryPointsByLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecoveryPointsByLegalHold"); err != nil {
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
	if err = addOpListRecoveryPointsByLegalHoldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecoveryPointsByLegalHold(options.Region), middleware.Before); err != nil {
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

// ListRecoveryPointsByLegalHoldPaginatorOptions is the paginator options for
// ListRecoveryPointsByLegalHold
type ListRecoveryPointsByLegalHoldPaginatorOptions struct {
	// The maximum number of resource list items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecoveryPointsByLegalHoldPaginator is a paginator for
// ListRecoveryPointsByLegalHold
type ListRecoveryPointsByLegalHoldPaginator struct {
	options   ListRecoveryPointsByLegalHoldPaginatorOptions
	client    ListRecoveryPointsByLegalHoldAPIClient
	params    *ListRecoveryPointsByLegalHoldInput
	nextToken *string
	firstPage bool
}

// NewListRecoveryPointsByLegalHoldPaginator returns a new
// ListRecoveryPointsByLegalHoldPaginator
func NewListRecoveryPointsByLegalHoldPaginator(client ListRecoveryPointsByLegalHoldAPIClient, params *ListRecoveryPointsByLegalHoldInput, optFns ...func(*ListRecoveryPointsByLegalHoldPaginatorOptions)) *ListRecoveryPointsByLegalHoldPaginator {
	if params == nil {
		params = &ListRecoveryPointsByLegalHoldInput{}
	}

	options := ListRecoveryPointsByLegalHoldPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecoveryPointsByLegalHoldPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecoveryPointsByLegalHoldPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecoveryPointsByLegalHold page.
func (p *ListRecoveryPointsByLegalHoldPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecoveryPointsByLegalHoldOutput, error) {
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
	result, err := p.client.ListRecoveryPointsByLegalHold(ctx, &params, optFns...)
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

// ListRecoveryPointsByLegalHoldAPIClient is a client that implements the
// ListRecoveryPointsByLegalHold operation.
type ListRecoveryPointsByLegalHoldAPIClient interface {
	ListRecoveryPointsByLegalHold(context.Context, *ListRecoveryPointsByLegalHoldInput, ...func(*Options)) (*ListRecoveryPointsByLegalHoldOutput, error)
}

var _ ListRecoveryPointsByLegalHoldAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRecoveryPointsByLegalHold(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecoveryPointsByLegalHold",
	}
}
