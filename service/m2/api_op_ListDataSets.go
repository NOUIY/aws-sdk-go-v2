// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the data sets imported for a specific application. In Amazon Web Services
// Mainframe Modernization, data sets are associated with applications deployed on
// runtime environments. This is known as importing data sets. Currently, Amazon
// Web Services Mainframe Modernization can import data sets into catalogs using [CreateDataSetImportTask].
//
// [CreateDataSetImportTask]: https://docs.aws.amazon.com/m2/latest/APIReference/API_CreateDataSetImportTask.html
func (c *Client) ListDataSets(ctx context.Context, params *ListDataSetsInput, optFns ...func(*Options)) (*ListDataSetsOutput, error) {
	if params == nil {
		params = &ListDataSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSets", params, optFns, c.addOperationListDataSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSetsInput struct {

	// The unique identifier of the application for which you want to list the
	// associated data sets.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of objects to return.
	MaxResults *int32

	// Filter dataset name matching the specified pattern. Can use * and % as wild
	// cards.
	NameFilter *string

	// A pagination token returned from a previous call to this operation. This
	// specifies the next item to return. To return to the beginning of the list,
	// exclude this parameter.
	NextToken *string

	// The prefix of the data set name, which you can use to filter the list of data
	// sets.
	Prefix *string

	noSmithyDocumentSerde
}

type ListDataSetsOutput struct {

	// The list of data sets, containing information including the creation time, the
	// data set name, the data set organization, the data set format, and the last time
	// the data set was referenced or updated.
	//
	// This member is required.
	DataSets []types.DataSetSummary

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to this operation to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataSets"); err != nil {
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
	if err = addOpListDataSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSets(options.Region), middleware.Before); err != nil {
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

// ListDataSetsPaginatorOptions is the paginator options for ListDataSets
type ListDataSetsPaginatorOptions struct {
	// The maximum number of objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSetsPaginator is a paginator for ListDataSets
type ListDataSetsPaginator struct {
	options   ListDataSetsPaginatorOptions
	client    ListDataSetsAPIClient
	params    *ListDataSetsInput
	nextToken *string
	firstPage bool
}

// NewListDataSetsPaginator returns a new ListDataSetsPaginator
func NewListDataSetsPaginator(client ListDataSetsAPIClient, params *ListDataSetsInput, optFns ...func(*ListDataSetsPaginatorOptions)) *ListDataSetsPaginator {
	if params == nil {
		params = &ListDataSetsInput{}
	}

	options := ListDataSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSets page.
func (p *ListDataSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSetsOutput, error) {
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
	result, err := p.client.ListDataSets(ctx, &params, optFns...)
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

// ListDataSetsAPIClient is a client that implements the ListDataSets operation.
type ListDataSetsAPIClient interface {
	ListDataSets(context.Context, *ListDataSetsInput, ...func(*Options)) (*ListDataSetsOutput, error)
}

var _ ListDataSetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDataSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataSets",
	}
}
