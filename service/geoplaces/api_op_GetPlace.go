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

// GetPlace finds a place by its unique ID. A PlaceId is returned by other place
// operations.
func (c *Client) GetPlace(ctx context.Context, params *GetPlaceInput, optFns ...func(*Options)) (*GetPlaceOutput, error) {
	if params == nil {
		params = &GetPlaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlace", params, optFns, c.addOperationGetPlaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlaceInput struct {

	// The PlaceId of the place you wish to receive the information for.
	//
	// This member is required.
	PlaceId *string

	// A list of optional additional parameters such as time zone that can be
	// requested for each result.
	AdditionalFeatures []types.GetPlaceAdditionalFeature

	// Indicates if the results will be stored. Defaults to SingleUse , if left empty.
	//
	// Storing the response of an GetPlace query is required to comply with service
	// terms, but charged at a higher cost per request. Please review the [user agreement]and [service pricing structure] to
	// determine the correct setting for your use case.
	//
	// [service pricing structure]: https://aws.amazon.com/location/pricing/
	// [user agreement]: https://aws.amazon.com/location/sla/
	IntendedUse types.GetPlaceIntendedUse

	// Optional: The API key to be used for authorization. Either an API key or valid
	// SigV4 signature must be provided when making a request.
	Key *string

	// A list of [BCP 47] compliant language codes for the results to be rendered in. If there
	// is no data for the result in the requested language, data will be returned in
	// the default language for the entry.
	//
	// [BCP 47]: https://en.wikipedia.org/wiki/IETF_language_tag
	Language *string

	// The alpha-2 or alpha-3 character code for the political view of a country. The
	// political view applies to the results of the request to represent unresolved
	// territorial claims through the point of view of the specified country.
	PoliticalView *string

	noSmithyDocumentSerde
}

type GetPlaceOutput struct {

	// The PlaceId of the place you wish to receive the information for.
	//
	// This member is required.
	PlaceId *string

	// A PlaceType is a category that the result place must belong to.
	//
	// This member is required.
	PlaceType types.PlaceType

	// The pricing bucket for which the query is charged at.
	//
	// For more information on pricing, please visit [Amazon Location Service Pricing].
	//
	// [Amazon Location Service Pricing]: https://aws.amazon.com/location/pricing/
	//
	// This member is required.
	PricingBucket *string

	// The localized display name of this result item based on request parameter
	// language .
	//
	// This member is required.
	Title *string

	// Position of the access point in (lng,lat) .
	AccessPoints []types.AccessPoint

	// Indicates known access restrictions on a vehicle access point. The index
	// correlates to an access point and indicates if access through this point has
	// some form of restriction.
	AccessRestrictions []types.AccessRestriction

	// The place's address.
	Address *types.Address

	// Boolean indicating if the address provided has been corrected.
	AddressNumberCorrected *bool

	// The Business Chains associated with the place.
	BusinessChains []types.BusinessChain

	// Categories of results that results must belong to.
	Categories []types.Category

	// List of potential contact methods for the result/place.
	Contacts *types.Contacts

	// List of food types offered by this result.
	FoodTypes []types.FoodType

	// The main address corresponding to a place of type Secondary Address.
	MainAddress *types.RelatedPlace

	// The bounding box enclosing the geometric shape (area or line) that an
	// individual result covers.
	//
	// The bounding box formed is defined as a set of four coordinates: [{westward
	// lng}, {southern lat}, {eastward lng}, {northern lat}]
	MapView []float64

	// List of opening hours objects.
	OpeningHours []types.OpeningHours

	// How the various components of the result's address are pronounced in various
	// languages.
	Phonemes *types.PhonemeDetails

	// The alpha-2 or alpha-3 character code for the political view of a country. The
	// political view applies to the results of the request to represent unresolved
	// territorial claims through the point of view of the specified country.
	PoliticalView *string

	// The position, in longitude and latitude.
	Position []float64

	// Contains details about the postal code of the place/result.
	PostalCodeDetails []types.PostalCodeDetails

	// All secondary addresses that are associated with a main address. A secondary
	// address is one that includes secondary designators, such as a Suite or Unit
	// Number, Building, or Floor information.
	SecondaryAddresses []types.RelatedPlace

	// The time zone in which the place is located.
	TimeZone *types.TimeZone

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPlaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPlace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPlace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPlace"); err != nil {
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
	if err = addOpGetPlaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPlace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPlace",
	}
}
