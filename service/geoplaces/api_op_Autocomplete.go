// Code generated by smithy-go-codegen DO NOT EDIT.

package geoplaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/geoplaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Autocomplete completes potential places and addresses as the user types, based
// on the partial input. The API enhances the efficiency and accuracy of address by
// completing query based on a few entered keystrokes. It helps you by completing
// partial queries with valid address completion. Also, the API supports the
// filtering of results based on geographic location, country, or specific place
// types, and can be tailored using optional parameters like language and political
// views.
func (c *Client) Autocomplete(ctx context.Context, params *AutocompleteInput, optFns ...func(*Options)) (*AutocompleteOutput, error) {
	if params == nil {
		params = &AutocompleteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Autocomplete", params, optFns, c.addOperationAutocompleteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AutocompleteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AutocompleteInput struct {

	// The free-form text query to match addresses against. This is usually a
	// partially typed address from an end user in an address box or form.
	//
	// The fields QueryText , and QueryID are mutually exclusive.
	//
	// This member is required.
	QueryText *string

	// A list of optional additional parameters that can be requested for each result.
	AdditionalFeatures []types.AutocompleteAdditionalFeature

	// The position in longitude and latitude that the results should be close to.
	// Typically, place results returned are ranked higher the closer they are to this
	// position. Stored in [lng, lat] and in the WSG84 format.
	//
	// The fields BiasPosition , FilterBoundingBox , and FilterCircle are mutually
	// exclusive.
	BiasPosition []float64

	// A structure which contains a set of inclusion/exclusion properties that results
	// must possess in order to be returned as a result.
	Filter *types.AutocompleteFilter

	// Indicates if the results will be stored. Defaults to SingleUse , if left empty.
	IntendedUse types.AutocompleteIntendedUse

	// Optional: The API key to be used for authorization. Either an API key or valid
	// SigV4 signature must be provided when making a request.
	Key *string

	// A list of [BCP 47] compliant language codes for the results to be rendered in. If there
	// is no data for the result in the requested language, data will be returned in
	// the default language for the entry.
	//
	// [BCP 47]: https://en.wikipedia.org/wiki/IETF_language_tag
	Language *string

	// An optional limit for the number of results returned in a single call.
	MaxResults *int32

	// The alpha-2 or alpha-3 character code for the political view of a country. The
	// political view applies to the results of the request to represent unresolved
	// territorial claims through the point of view of the specified country.
	//
	// The following political views are currently supported:
	//
	//   - ARG : Argentina's view on the Southern Patagonian Ice Field and Tierra Del
	//   Fuego, including the Falkland Islands, South Georgia, and South Sandwich Islands
	//
	//   - EGY : Egypt's view on Bir Tawil
	//
	//   - IND : India's view on Gilgit-Baltistan
	//
	//   - KEN : Kenya's view on the Ilemi Triangle
	//
	//   - MAR : Morocco's view on Western Sahara
	//
	//   - RUS : Russia's view on Crimea
	//
	//   - SDN : Sudan's view on the Halaib Triangle
	//
	//   - SRB : Serbia's view on Kosovo, Vukovar, and Sarengrad Islands
	//
	//   - SUR : Suriname's view on the Courantyne Headwaters and Lawa Headwaters
	//
	//   - SYR : Syria's view on the Golan Heights
	//
	//   - TUR : Turkey's view on Cyprus and Northern Cyprus
	//
	//   - TZA : Tanzania's view on Lake Malawi
	//
	//   - URY : Uruguay's view on Rincon de Artigas
	//
	//   - VNM : Vietnam's view on the Paracel Islands and Spratly Islands
	PoliticalView *string

	// The PostalCodeMode affects how postal code results are returned. If a postal
	// code spans multiple localities and this value is empty, partial district or
	// locality information may be returned under a single postal code result entry. If
	// it's populated with the value EnumerateSpannedLocalities , all cities in that
	// postal code are returned.
	PostalCodeMode types.PostalCodeMode

	noSmithyDocumentSerde
}

type AutocompleteOutput struct {

	// The pricing bucket for which the query is charged at.
	//
	// For more information on pricing, please visit [Amazon Location Service Pricing].
	//
	// [Amazon Location Service Pricing]: https://aws.amazon.com/location/pricing/
	//
	// This member is required.
	PricingBucket *string

	// List of places or results returned for a query.
	ResultItems []types.AutocompleteResultItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAutocompleteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAutocomplete{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAutocomplete{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Autocomplete"); err != nil {
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
	if err = addOpAutocompleteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAutocomplete(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAutocomplete(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Autocomplete",
	}
}
