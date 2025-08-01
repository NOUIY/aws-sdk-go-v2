// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) pre-aggregated data for a standard metric that applies to
// an application.
func (c *Client) GetApplicationDateRangeKpi(ctx context.Context, params *GetApplicationDateRangeKpiInput, optFns ...func(*Options)) (*GetApplicationDateRangeKpiOutput, error) {
	if params == nil {
		params = &GetApplicationDateRangeKpiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplicationDateRangeKpi", params, optFns, c.addOperationGetApplicationDateRangeKpiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationDateRangeKpiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationDateRangeKpiInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The name of the metric, also referred to as a key performance indicator (KPI),
	// to retrieve data for. This value describes the associated metric and consists of
	// two or more terms, which are comprised of lowercase alphanumeric characters,
	// separated by a hyphen. Examples are email-open-rate and
	// successful-delivery-rate. For a list of valid values, see the [Amazon Pinpoint Developer Guide].
	//
	// [Amazon Pinpoint Developer Guide]: https://docs.aws.amazon.com/pinpoint/latest/developerguide/analytics-standard-metrics.html
	//
	// This member is required.
	KpiName *string

	// The last date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-26T20:00:00Z for
	// 8:00 PM UTC July 26, 2019.
	EndTime *time.Time

	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	// The first date and time to retrieve data for, as part of an inclusive date
	// range that filters the query results. This value should be in extended ISO 8601
	// format and use Coordinated Universal Time (UTC), for example:
	// 2019-07-19T20:00:00Z for 8:00 PM UTC July 19, 2019. This value should also be
	// fewer than 90 days from the current day.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetApplicationDateRangeKpiOutput struct {

	// Provides the results of a query that retrieved the data for a standard metric
	// that applies to an application, and provides information about that query.
	//
	// This member is required.
	ApplicationDateRangeKpiResponse *types.ApplicationDateRangeKpiResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationDateRangeKpiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplicationDateRangeKpi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplicationDateRangeKpi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetApplicationDateRangeKpi"); err != nil {
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
	if err = addOpGetApplicationDateRangeKpiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationDateRangeKpi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplicationDateRangeKpi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetApplicationDateRangeKpi",
	}
}
