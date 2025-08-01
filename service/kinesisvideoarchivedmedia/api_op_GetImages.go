// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a list of images corresponding to each timestamp for a given time
// range, sampling interval, and image format configuration.
func (c *Client) GetImages(ctx context.Context, params *GetImagesInput, optFns ...func(*Options)) (*GetImagesOutput, error) {
	if params == nil {
		params = &GetImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImages", params, optFns, c.addOperationGetImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImagesInput struct {

	// The end timestamp for the range of images to be generated. If the time range
	// between StartTimestamp and EndTimestamp is more than 300 seconds above
	// StartTimestamp , you will receive an IllegalArgumentException .
	//
	// This member is required.
	EndTimestamp *time.Time

	// The format that will be used to encode the image.
	//
	// This member is required.
	Format types.Format

	// The origin of the Server or Producer timestamps to use to generate the images.
	//
	// This member is required.
	ImageSelectorType types.ImageSelectorType

	// The starting point from which the images should be generated. This
	// StartTimestamp must be within an inclusive range of timestamps for an image to
	// be returned.
	//
	// This member is required.
	StartTimestamp *time.Time

	// The list of a key-value pair structure that contains extra parameters that can
	// be applied when the image is generated. The FormatConfig key is the JPEGQuality
	// , which indicates the JPEG quality key to be used to generate the image. The
	// FormatConfig value accepts ints from 1 to 100. If the value is 1, the image will
	// be generated with less quality and the best compression. If the value is 100,
	// the image will be generated with the best quality and less compression. If no
	// value is provided, the default value of the JPEGQuality key will be set to 80.
	FormatConfig map[string]string

	// The height of the output image that is used in conjunction with the WidthPixels
	// parameter. When both HeightPixels and WidthPixels parameters are provided, the
	// image will be stretched to fit the specified aspect ratio. If only the
	// HeightPixels parameter is provided, its original aspect ratio will be used to
	// calculate the WidthPixels ratio. If neither parameter is provided, the original
	// image size will be returned.
	HeightPixels *int32

	// The maximum number of images to be returned by the API.
	//
	// The default limit is 25 images per API response. Providing a MaxResults greater
	// than this value will result in a page size of 25. Any additional results will be
	// paginated.
	MaxResults *int64

	// A token that specifies where to start paginating the next set of Images. This
	// is the GetImages:NextToken from a previously truncated response.
	NextToken *string

	// The time interval in milliseconds (ms) at which the images need to be generated
	// from the stream. The minimum value that can be provided is 200 ms (5 images per
	// second). If the timestamp range is less than the sampling interval, the image
	// from the startTimestamp will be returned if available.
	SamplingInterval *int32

	// The Amazon Resource Name (ARN) of the stream from which to retrieve the images.
	// You must specify either the StreamName or the StreamARN .
	StreamARN *string

	// The name of the stream from which to retrieve the images. You must specify
	// either the StreamName or the StreamARN .
	StreamName *string

	// The width of the output image that is used in conjunction with the HeightPixels
	// parameter. When both WidthPixels and HeightPixels parameters are provided, the
	// image will be stretched to fit the specified aspect ratio. If only the
	// WidthPixels parameter is provided or if only the HeightPixels is provided, a
	// ValidationException will be thrown. If neither parameter is provided, the
	// original image size from the stream will be returned.
	WidthPixels *int32

	noSmithyDocumentSerde
}

type GetImagesOutput struct {

	// The list of images generated from the video stream. If there is no media
	// available for the given timestamp, the NO_MEDIA error will be listed in the
	// output. If an error occurs while the image is being generated, the MEDIA_ERROR
	// will be listed in the output as the cause of the missing image.
	Images []types.Image

	// The encrypted token that was used in the request to get more images.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImages{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetImages"); err != nil {
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
	if err = addOpGetImagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImages(options.Region), middleware.Before); err != nil {
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

// GetImagesPaginatorOptions is the paginator options for GetImages
type GetImagesPaginatorOptions struct {
	// The maximum number of images to be returned by the API.
	//
	// The default limit is 25 images per API response. Providing a MaxResults greater
	// than this value will result in a page size of 25. Any additional results will be
	// paginated.
	Limit int64

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetImagesPaginator is a paginator for GetImages
type GetImagesPaginator struct {
	options   GetImagesPaginatorOptions
	client    GetImagesAPIClient
	params    *GetImagesInput
	nextToken *string
	firstPage bool
}

// NewGetImagesPaginator returns a new GetImagesPaginator
func NewGetImagesPaginator(client GetImagesAPIClient, params *GetImagesInput, optFns ...func(*GetImagesPaginatorOptions)) *GetImagesPaginator {
	if params == nil {
		params = &GetImagesInput{}
	}

	options := GetImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetImages page.
func (p *GetImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetImagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int64
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetImages(ctx, &params, optFns...)
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

// GetImagesAPIClient is a client that implements the GetImages operation.
type GetImagesAPIClient interface {
	GetImages(context.Context, *GetImagesInput, ...func(*Options)) (*GetImagesOutput, error)
}

var _ GetImagesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetImages",
	}
}
