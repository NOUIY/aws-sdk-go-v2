// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets recommendations for reservation purchases. These recommendations might
// help you to reduce your costs. Reservations provide a discounted hourly rate (up
// to 75%) compared to On-Demand pricing.
//
// Amazon Web Services generates your recommendations by identifying your
// On-Demand usage during a specific time period and collecting your usage into
// categories that are eligible for a reservation. After Amazon Web Services has
// these categories, it simulates every combination of reservations in each
// category of usage to identify the best number of each type of Reserved Instance
// (RI) to purchase to maximize your estimated savings.
//
// For example, Amazon Web Services automatically aggregates your Amazon EC2
// Linux, shared tenancy, and c4 family usage in the US West (Oregon) Region and
// recommends that you buy size-flexible regional reservations to apply to the c4
// family usage. Amazon Web Services recommends the smallest size instance in an
// instance family. This makes it easier to purchase a size-flexible Reserved
// Instance (RI). Amazon Web Services also shows the equal number of normalized
// units. This way, you can purchase any instance size that you want. For this
// example, your RI recommendation is for c4.large because that is the smallest
// size instance in the c4 instance family.
func (c *Client) GetReservationPurchaseRecommendation(ctx context.Context, params *GetReservationPurchaseRecommendationInput, optFns ...func(*Options)) (*GetReservationPurchaseRecommendationOutput, error) {
	if params == nil {
		params = &GetReservationPurchaseRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservationPurchaseRecommendation", params, optFns, c.addOperationGetReservationPurchaseRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservationPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservationPurchaseRecommendationInput struct {

	// The specific service that you want recommendations for.
	//
	// This member is required.
	Service *string

	// The account ID that's associated with the recommendation.
	AccountId *string

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the management account and member accounts
	// if the value is set to PAYER . If the value is LINKED , recommendations are
	// calculated for individual member accounts only.
	AccountScope types.AccountScope

	// Use Expression to filter in various Cost Explorer APIs.
	//
	// Not all Expression types are supported in each API. Refer to the documentation
	// for each specific API to see what is supported.
	//
	// There are two patterns:
	//
	//   - Simple dimension values.
	//
	//   - There are three types of simple dimension values: CostCategories , Tags ,
	//   and Dimensions .
	//
	//   - Specify the CostCategories field to define a filter that acts on Cost
	//   Categories.
	//
	//   - Specify the Tags field to define a filter that acts on Cost Allocation Tags.
	//
	//   - Specify the Dimensions field to define a filter that acts on the [DimensionValues]
	//   DimensionValues .
	//
	//   - For each filter type, you can set the dimension name and values for the
	//   filters that you plan to use.
	//
	//   - For example, you can filter for REGION==us-east-1 OR REGION==us-west-1 . For
	//   GetRightsizingRecommendation , the Region is a full name (for example,
	//   REGION==US East (N. Virginia) .
	//
	//   - The corresponding Expression for this example is as follows: {
	//   "Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ] } }
	//
	//   - As shown in the previous example, lists of dimension values are combined
	//   with OR when applying the filter.
	//
	//   - You can also set different match options to further control how the filter
	//   behaves. Not all APIs support match options. Refer to the documentation for each
	//   specific API to see what is supported.
	//
	//   - For example, you can filter for linked account names that start with "a".
	//
	//   - The corresponding Expression for this example is as follows: {
	//   "Dimensions": { "Key": "LINKED_ACCOUNT_NAME", "MatchOptions": [ "STARTS_WITH" ],
	//   "Values": [ "a" ] } }
	//
	//   - Compound Expression types with logical operations.
	//
	//   - You can use multiple Expression types and the logical operators AND/OR/NOT
	//   to create a list of one or more Expression objects. By doing this, you can
	//   filter by more advanced options.
	//
	//   - For example, you can filter by ((REGION == us-east-1 OR REGION ==
	//   us-west-1) OR (TAG.Type == Type1)) AND (USAGE_TYPE != DataTransfer) .
	//
	//   - The corresponding Expression for this example is as follows: { "And": [
	//   {"Or": [ {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1"
	//   ] }}, {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not":
	//   {"Dimensions": { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }
	//
	// Because each Expression can have only one operator, the service returns an error
	//   if more than one is specified. The following example shows an Expression
	//   object that creates an error: { "And": [ ... ], "Dimensions": { "Key":
	//   "USAGE_TYPE", "Values": [ "DataTransfer" ] } }
	//
	// The following is an example of the corresponding error message: "Expression has
	//   more than one roots. Only one root operator is allowed for each expression: And,
	//   Or, Not, Dimensions, Tags, CostCategories"
	//
	// For the GetRightsizingRecommendation action, a combination of OR and NOT isn't
	// supported. OR isn't supported between different dimensions, or dimensions and
	// tags. NOT operators aren't supported. Dimensions are also limited to
	// LINKED_ACCOUNT , REGION , or RIGHTSIZING_TYPE .
	//
	// For the GetReservationPurchaseRecommendation action, only NOT is supported. AND
	// and OR aren't supported. Dimensions are limited to LINKED_ACCOUNT .
	//
	// [DimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_DimensionValues.html
	Filter *types.Expression

	// The number of previous days that you want Amazon Web Services to consider when
	// it calculates your recommendations.
	LookbackPeriodInDays types.LookbackPeriodInDays

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	// The reservation purchase option that you want recommendations for.
	PaymentOption types.PaymentOption

	// The hardware specifications for the service instances that you want
	// recommendations for, such as standard or convertible Amazon EC2 instances.
	ServiceSpecification *types.ServiceSpecification

	// The reservation term that you want recommendations for.
	TermInYears types.TermInYears

	noSmithyDocumentSerde
}

type GetReservationPurchaseRecommendationOutput struct {

	// Information about this specific recommendation call, such as the time stamp for
	// when Cost Explorer generated this recommendation.
	Metadata *types.ReservationPurchaseRecommendationMetadata

	// The pagination token for the next set of retrievable results.
	NextPageToken *string

	// Recommendations for reservations to purchase.
	Recommendations []types.ReservationPurchaseRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservationPurchaseRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReservationPurchaseRecommendation"); err != nil {
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
	if err = addOpGetReservationPurchaseRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReservationPurchaseRecommendation",
	}
}
