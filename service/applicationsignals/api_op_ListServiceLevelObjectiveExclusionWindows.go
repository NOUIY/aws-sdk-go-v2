// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationsignals

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all exclusion windows configured for a specific SLO.
func (c *Client) ListServiceLevelObjectiveExclusionWindows(ctx context.Context, params *ListServiceLevelObjectiveExclusionWindowsInput, optFns ...func(*Options)) (*ListServiceLevelObjectiveExclusionWindowsOutput, error) {
	if params == nil {
		params = &ListServiceLevelObjectiveExclusionWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceLevelObjectiveExclusionWindows", params, optFns, c.addOperationListServiceLevelObjectiveExclusionWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceLevelObjectiveExclusionWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceLevelObjectiveExclusionWindowsInput struct {

	// The ID of the SLO to list exclusion windows for.
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	MaxResults *int32

	// Include this value, if it was returned by the previous operation, to get the
	// next set of service level objectives.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceLevelObjectiveExclusionWindowsOutput struct {

	// A list of exclusion windows configured for the SLO.
	//
	// This member is required.
	ExclusionWindows []types.ExclusionWindow

	// Include this value, if it was returned by the previous operation, to get the
	// next set of service level objectives.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceLevelObjectiveExclusionWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServiceLevelObjectiveExclusionWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServiceLevelObjectiveExclusionWindows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceLevelObjectiveExclusionWindows"); err != nil {
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
	if err = addOpListServiceLevelObjectiveExclusionWindowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceLevelObjectiveExclusionWindows(options.Region), middleware.Before); err != nil {
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

// ListServiceLevelObjectiveExclusionWindowsPaginatorOptions is the paginator
// options for ListServiceLevelObjectiveExclusionWindows
type ListServiceLevelObjectiveExclusionWindowsPaginatorOptions struct {
	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceLevelObjectiveExclusionWindowsPaginator is a paginator for
// ListServiceLevelObjectiveExclusionWindows
type ListServiceLevelObjectiveExclusionWindowsPaginator struct {
	options   ListServiceLevelObjectiveExclusionWindowsPaginatorOptions
	client    ListServiceLevelObjectiveExclusionWindowsAPIClient
	params    *ListServiceLevelObjectiveExclusionWindowsInput
	nextToken *string
	firstPage bool
}

// NewListServiceLevelObjectiveExclusionWindowsPaginator returns a new
// ListServiceLevelObjectiveExclusionWindowsPaginator
func NewListServiceLevelObjectiveExclusionWindowsPaginator(client ListServiceLevelObjectiveExclusionWindowsAPIClient, params *ListServiceLevelObjectiveExclusionWindowsInput, optFns ...func(*ListServiceLevelObjectiveExclusionWindowsPaginatorOptions)) *ListServiceLevelObjectiveExclusionWindowsPaginator {
	if params == nil {
		params = &ListServiceLevelObjectiveExclusionWindowsInput{}
	}

	options := ListServiceLevelObjectiveExclusionWindowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceLevelObjectiveExclusionWindowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceLevelObjectiveExclusionWindowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceLevelObjectiveExclusionWindows page.
func (p *ListServiceLevelObjectiveExclusionWindowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceLevelObjectiveExclusionWindowsOutput, error) {
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
	result, err := p.client.ListServiceLevelObjectiveExclusionWindows(ctx, &params, optFns...)
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

// ListServiceLevelObjectiveExclusionWindowsAPIClient is a client that implements
// the ListServiceLevelObjectiveExclusionWindows operation.
type ListServiceLevelObjectiveExclusionWindowsAPIClient interface {
	ListServiceLevelObjectiveExclusionWindows(context.Context, *ListServiceLevelObjectiveExclusionWindowsInput, ...func(*Options)) (*ListServiceLevelObjectiveExclusionWindowsOutput, error)
}

var _ ListServiceLevelObjectiveExclusionWindowsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListServiceLevelObjectiveExclusionWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceLevelObjectiveExclusionWindows",
	}
}
