// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves glyphs used to display labels on a map.
func (c *Client) GetMapGlyphs(ctx context.Context, params *GetMapGlyphsInput, optFns ...func(*Options)) (*GetMapGlyphsOutput, error) {
	if params == nil {
		params = &GetMapGlyphsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapGlyphs", params, optFns, c.addOperationGetMapGlyphsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapGlyphsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapGlyphsInput struct {

	// A comma-separated list of fonts to load glyphs from in order of preference. For
	// example, Noto Sans Regular, Arial Unicode .
	//
	// Valid font stacks for [Esri] styles:
	//
	//   - VectorEsriDarkGrayCanvas – Ubuntu Medium Italic | Ubuntu Medium | Ubuntu
	//   Italic | Ubuntu Regular | Ubuntu Bold
	//
	//   - VectorEsriLightGrayCanvas – Ubuntu Italic | Ubuntu Regular | Ubuntu Light |
	//   Ubuntu Bold
	//
	//   - VectorEsriTopographic – Noto Sans Italic | Noto Sans Regular | Noto Sans
	//   Bold | Noto Serif Regular | Roboto Condensed Light Italic
	//
	//   - VectorEsriStreets – Arial Regular | Arial Italic | Arial Bold
	//
	//   - VectorEsriNavigation – Arial Regular | Arial Italic | Arial Bold
	//
	// Valid font stacks for [HERE Technologies] styles:
	//
	//   - VectorHereContrast – Fira GO Regular | Fira GO Bold
	//
	//   - VectorHereExplore, VectorHereExploreTruck, HybridHereExploreSatellite –
	//   Fira GO Italic | Fira GO Map | Fira GO Map Bold | Noto Sans CJK JP Bold |
	//   Noto Sans CJK JP Light | Noto Sans CJK JP Regular
	//
	// Valid font stacks for [GrabMaps] styles:
	//
	//   - VectorGrabStandardLight, VectorGrabStandardDark – Noto Sans Regular | Noto
	//   Sans Medium | Noto Sans Bold
	//
	// Valid font stacks for [Open Data] styles:
	//
	//   - VectorOpenDataStandardLight, VectorOpenDataStandardDark,
	//   VectorOpenDataVisualizationLight, VectorOpenDataVisualizationDark – Amazon
	//   Ember Regular,Noto Sans Regular | Amazon Ember Bold,Noto Sans Bold | Amazon
	//   Ember Medium,Noto Sans Medium | Amazon Ember Regular Italic,Noto Sans Italic |
	//   Amazon Ember Condensed RC Regular,Noto Sans Regular | Amazon Ember Condensed
	//   RC Bold,Noto Sans Bold | Amazon Ember Regular,Noto Sans Regular,Noto Sans
	//   Arabic Regular | Amazon Ember Condensed RC Bold,Noto Sans Bold,Noto Sans
	//   Arabic Condensed Bold | Amazon Ember Bold,Noto Sans Bold,Noto Sans Arabic Bold
	//   | Amazon Ember Regular Italic,Noto Sans Italic,Noto Sans Arabic Regular |
	//   Amazon Ember Condensed RC Regular,Noto Sans Regular,Noto Sans Arabic Condensed
	//   Regular | Amazon Ember Medium,Noto Sans Medium,Noto Sans Arabic Medium
	//
	// The fonts used by the Open Data map styles are combined fonts that use Amazon
	// Ember for most glyphs but Noto Sans for glyphs unsupported by Amazon Ember .
	//
	// [Esri]: https://docs.aws.amazon.com/location/previous/developerguide/esri.html
	// [HERE Technologies]: https://docs.aws.amazon.com/location/previous/developerguide/HERE.html
	// [GrabMaps]: https://docs.aws.amazon.com/location/previous/developerguide/grab.html
	// [Open Data]: https://docs.aws.amazon.com/location/previous/developerguide/open-data.html
	//
	// This member is required.
	FontStack *string

	// A Unicode range of characters to download glyphs for. Each response will
	// contain 256 characters. For example, 0–255 includes all characters from range
	// U+0000 to 00FF . Must be aligned to multiples of 256.
	//
	// This member is required.
	FontUnicodeRange *string

	// The map resource associated with the glyph ﬁle.
	//
	// This member is required.
	MapName *string

	// The optional [API key] to authorize the request.
	//
	// [API key]: https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html
	Key *string

	noSmithyDocumentSerde
}

type GetMapGlyphsOutput struct {

	// The glyph, as binary blob.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The map glyph content type. For example, application/octet-stream .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapGlyphsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapGlyphs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapGlyphs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMapGlyphs"); err != nil {
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
	if err = addEndpointPrefix_opGetMapGlyphsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMapGlyphsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapGlyphs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapGlyphsMiddleware struct {
}

func (*endpointPrefix_opGetMapGlyphsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapGlyphsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetMapGlyphsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetMapGlyphsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetMapGlyphs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMapGlyphs",
	}
}
