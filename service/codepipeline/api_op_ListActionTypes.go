// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a summary of all CodePipeline action types associated with your account.
func (c *Client) ListActionTypes(ctx context.Context, params *ListActionTypesInput, optFns ...func(*Options)) (*ListActionTypesOutput, error) {
	if params == nil {
		params = &ListActionTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActionTypes", params, optFns, c.addOperationListActionTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActionTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListActionTypes action.
type ListActionTypesInput struct {

	// Filters the list of action types to those created by a specified entity.
	ActionOwnerFilter types.ActionOwner

	// An identifier that was returned from the previous list action types call, which
	// can be used to return the next set of action types in the list.
	NextToken *string

	// The Region to filter on for the list of action types.
	RegionFilter *string

	noSmithyDocumentSerde
}

// Represents the output of a ListActionTypes action.
type ListActionTypesOutput struct {

	// Provides details of the action types.
	//
	// This member is required.
	ActionTypes []types.ActionType

	// If the amount of returned information is significantly large, an identifier is
	// also returned. It can be used in a subsequent list action types call to return
	// the next set of action types in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListActionTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListActionTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActionTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListActionTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActionTypes(options.Region), middleware.Before); err != nil {
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

// ListActionTypesPaginatorOptions is the paginator options for ListActionTypes
type ListActionTypesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListActionTypesPaginator is a paginator for ListActionTypes
type ListActionTypesPaginator struct {
	options   ListActionTypesPaginatorOptions
	client    ListActionTypesAPIClient
	params    *ListActionTypesInput
	nextToken *string
	firstPage bool
}

// NewListActionTypesPaginator returns a new ListActionTypesPaginator
func NewListActionTypesPaginator(client ListActionTypesAPIClient, params *ListActionTypesInput, optFns ...func(*ListActionTypesPaginatorOptions)) *ListActionTypesPaginator {
	if params == nil {
		params = &ListActionTypesInput{}
	}

	options := ListActionTypesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListActionTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListActionTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListActionTypes page.
func (p *ListActionTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListActionTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListActionTypes(ctx, &params, optFns...)
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

// ListActionTypesAPIClient is a client that implements the ListActionTypes
// operation.
type ListActionTypesAPIClient interface {
	ListActionTypes(context.Context, *ListActionTypesInput, ...func(*Options)) (*ListActionTypesOutput, error)
}

var _ ListActionTypesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListActionTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListActionTypes",
	}
}
