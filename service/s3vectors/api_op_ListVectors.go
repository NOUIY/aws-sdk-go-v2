// Code generated by smithy-go-codegen DO NOT EDIT.

package s3vectors

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/s3vectors/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.
//
// List vectors in the specified vector index. To specify the vector index, you
// can either use both the vector bucket name and the vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// ListVectors operations proceed sequentially; however, for faster performance on
// a large number of vectors in a vector index, applications can request a parallel
// ListVectors operation by providing the segmentCount and segmentIndex parameters.
//
// Permissions You must have the s3vectors:ListVectors permission to use this
// operation. Additional permissions are required based on the request parameters
// you specify:
//
//   - With only s3vectors:ListVectors permission, you can list vector keys when
//     returnData and returnMetadata are both set to false or not specified..
//
//   - If you set returnData or returnMetadata to true, you must have both
//     s3vectors:ListVectors and s3vectors:GetVectors permissions. The request fails
//     with a 403 Forbidden error if you request vector data or metadata without the
//     s3vectors:GetVectors permission.
func (c *Client) ListVectors(ctx context.Context, params *ListVectorsInput, optFns ...func(*Options)) (*ListVectorsOutput, error) {
	if params == nil {
		params = &ListVectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVectors", params, optFns, c.addOperationListVectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVectorsInput struct {

	// The Amazon resource Name (ARN) of the vector index.
	IndexArn *string

	// The name of the vector index.
	IndexName *string

	// The maximum number of vectors to return on a page.
	//
	// If you don't specify maxResults , the ListVectors operation uses a default
	// value of 500.
	//
	// If the processed dataset size exceeds 1 MB before reaching the maxResults
	// value, the operation stops and returns the vectors that are retrieved up to that
	// point, along with a nextToken that you can use in a subsequent request to
	// retrieve the next set of results.
	MaxResults *int32

	// Pagination token from a previous request. The value of this field is empty for
	// an initial request.
	NextToken *string

	// If true, the vector data of each vector will be included in the response. The
	// default value is false .
	ReturnData bool

	// If true, the metadata associated with each vector will be included in the
	// response. The default value is false .
	ReturnMetadata bool

	// For a parallel ListVectors request, segmentCount represents the total number of
	// vector segments into which the ListVectors operation will be divided. The value
	// of segmentCount corresponds to the number of application workers that will
	// perform the parallel ListVectors operation. For example, if you want to use
	// four application threads to list vectors in a vector index, specify a
	// segmentCount value of 4.
	//
	// If you specify a segmentCount value of 1, the ListVectors operation will be
	// sequential rather than parallel.
	//
	// If you specify segmentCount , you must also specify segmentIndex .
	SegmentCount *int32

	// For a parallel ListVectors request, segmentIndex is the index of the segment
	// from which to list vectors in the current request. It identifies an individual
	// segment to be listed by an application worker.
	//
	// Segment IDs are zero-based, so the first segment is always 0. For example, if
	// you want to use four application threads to list vectors in a vector index, then
	// the first thread specifies a segmentIndex value of 0, the second thread
	// specifies 1, and so on.
	//
	// The value of segmentIndex must be less than the value provided for segmentCount
	// .
	//
	// If you provide segmentIndex , you must also provide segmentCount .
	SegmentIndex int32

	// The name of the vector bucket.
	VectorBucketName *string

	noSmithyDocumentSerde
}

type ListVectorsOutput struct {

	// Vectors in the current segment.
	//
	// This member is required.
	Vectors []types.ListOutputVector

	// Pagination token to be used in the subsequent request. The field is empty if no
	// further pagination is required.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVectors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVectors(options.Region), middleware.Before); err != nil {
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

// ListVectorsPaginatorOptions is the paginator options for ListVectors
type ListVectorsPaginatorOptions struct {
	// The maximum number of vectors to return on a page.
	//
	// If you don't specify maxResults , the ListVectors operation uses a default
	// value of 500.
	//
	// If the processed dataset size exceeds 1 MB before reaching the maxResults
	// value, the operation stops and returns the vectors that are retrieved up to that
	// point, along with a nextToken that you can use in a subsequent request to
	// retrieve the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVectorsPaginator is a paginator for ListVectors
type ListVectorsPaginator struct {
	options   ListVectorsPaginatorOptions
	client    ListVectorsAPIClient
	params    *ListVectorsInput
	nextToken *string
	firstPage bool
}

// NewListVectorsPaginator returns a new ListVectorsPaginator
func NewListVectorsPaginator(client ListVectorsAPIClient, params *ListVectorsInput, optFns ...func(*ListVectorsPaginatorOptions)) *ListVectorsPaginator {
	if params == nil {
		params = &ListVectorsInput{}
	}

	options := ListVectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVectors page.
func (p *ListVectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVectorsOutput, error) {
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
	result, err := p.client.ListVectors(ctx, &params, optFns...)
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

// ListVectorsAPIClient is a client that implements the ListVectors operation.
type ListVectorsAPIClient interface {
	ListVectors(context.Context, *ListVectorsInput, ...func(*Options)) (*ListVectorsOutput, error)
}

var _ ListVectorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListVectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVectors",
	}
}
