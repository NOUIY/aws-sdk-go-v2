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

// Retrieves the reservation coverage for your account, which you can use to see
// how much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon
// Relational Database Service, or Amazon Redshift usage is covered by a
// reservation. An organization's management account can see the coverage of the
// associated member accounts. This supports dimensions, Cost Categories, and
// nested expressions. For any time period, you can filter data about reservation
// usage by the following dimensions:
//
//   - AZ
//
//   - CACHE_ENGINE
//
//   - DATABASE_ENGINE
//
//   - DEPLOYMENT_OPTION
//
//   - INSTANCE_TYPE
//
//   - LINKED_ACCOUNT
//
//   - OPERATING_SYSTEM
//
//   - PLATFORM
//
//   - REGION
//
//   - SERVICE
//
//   - TAG
//
//   - TENANCY
//
// To determine valid values for a dimension, use the GetDimensionValues
// operation.
func (c *Client) GetReservationCoverage(ctx context.Context, params *GetReservationCoverageInput, optFns ...func(*Options)) (*GetReservationCoverageOutput, error) {
	if params == nil {
		params = &GetReservationCoverageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservationCoverage", params, optFns, c.addOperationGetReservationCoverageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservationCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// You can use the following request parameters to query for how much of your
// instance usage a reservation covered.
type GetReservationCoverageInput struct {

	// The start and end dates of the period that you want to retrieve data about
	// reservation coverage for. You can retrieve data for a maximum of 13 months: the
	// last 12 months and the current month. The start date is inclusive, but the end
	// date is exclusive. For example, if start is 2017-01-01 and end is 2017-05-01 ,
	// then the cost and usage data is retrieved from 2017-01-01 up to and including
	// 2017-04-30 but not including 2017-05-01 .
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters utilization data by dimensions. You can filter by the following
	// dimensions:
	//
	//   - AZ
	//
	//   - CACHE_ENGINE
	//
	//   - DATABASE_ENGINE
	//
	//   - DEPLOYMENT_OPTION
	//
	//   - INSTANCE_TYPE
	//
	//   - LINKED_ACCOUNT
	//
	//   - OPERATING_SYSTEM
	//
	//   - PLATFORM
	//
	//   - REGION
	//
	//   - SERVICE
	//
	//   - TAG
	//
	//   - TENANCY
	//
	// GetReservationCoverage uses the same [Expression] object as the other operations, but only
	// AND is supported among each dimension. You can nest only one level deep. If
	// there are multiple values for a dimension, they are OR'd together.
	//
	// If you don't provide a SERVICE filter, Cost Explorer defaults to EC2.
	//
	// Cost category is also supported.
	//
	// [Expression]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html
	Filter *types.Expression

	// The granularity of the Amazon Web Services cost data for the reservation. Valid
	// values are MONTHLY and DAILY .
	//
	// If GroupBy is set, Granularity can't be set. If Granularity isn't set, the
	// response object doesn't include Granularity , either MONTHLY or DAILY .
	//
	// The GetReservationCoverage operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity

	// You can group the data by the following attributes:
	//
	//   - AZ
	//
	//   - CACHE_ENGINE
	//
	//   - DATABASE_ENGINE
	//
	//   - DEPLOYMENT_OPTION
	//
	//   - INSTANCE_TYPE
	//
	//   - INVOICING_ENTITY
	//
	//   - LINKED_ACCOUNT
	//
	//   - OPERATING_SYSTEM
	//
	//   - PLATFORM
	//
	//   - REGION
	//
	//   - TENANCY
	GroupBy []types.GroupDefinition

	// The maximum number of objects that you returned for this request. If more
	// objects are available, in the response, Amazon Web Services provides a
	// NextPageToken value that you can use in a subsequent call to get the next batch
	// of objects.
	MaxResults *int32

	// The measurement that you want your reservation coverage reported in.
	//
	// Valid values are Hour , Unit , and Cost . You can use multiple values in a
	// request.
	Metrics []string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The value by which you want to sort the data.
	//
	// The following values are supported for Key :
	//
	//   - OnDemandCost
	//
	//   - CoverageHoursPercentage
	//
	//   - OnDemandHours
	//
	//   - ReservedHours
	//
	//   - TotalRunningHours
	//
	//   - CoverageNormalizedUnitsPercentage
	//
	//   - OnDemandNormalizedUnits
	//
	//   - ReservedNormalizedUnits
	//
	//   - TotalRunningNormalizedUnits
	//
	//   - Time
	//
	// Supported values for SortOrder are ASCENDING or DESCENDING .
	SortBy *types.SortDefinition

	noSmithyDocumentSerde
}

type GetReservationCoverageOutput struct {

	// The amount of time that your reservations covered.
	//
	// This member is required.
	CoveragesByTime []types.CoverageByTime

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// The total amount of instance usage that a reservation covered.
	Total *types.Coverage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservationCoverageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReservationCoverage"); err != nil {
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
	if err = addOpGetReservationCoverageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationCoverage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReservationCoverage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReservationCoverage",
	}
}
