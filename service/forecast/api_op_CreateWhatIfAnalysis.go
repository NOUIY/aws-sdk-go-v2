// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// What-if analysis is a scenario modeling technique where you make a hypothetical
// change to a time series and compare the forecasts generated by these changes
// against the baseline, unchanged time series. It is important to remember that
// the purpose of a what-if analysis is to understand how a forecast can change
// given different modifications to the baseline time series.
//
// For example, imagine you are a clothing retailer who is considering an end of
// season sale to clear space for new styles. After creating a baseline forecast,
// you can use a what-if analysis to investigate how different sales tactics might
// affect your goals.
//
// You could create a scenario where everything is given a 25% markdown, and
// another where everything is given a fixed dollar markdown. You could create a
// scenario where the sale lasts for one week and another where the sale lasts for
// one month. With a what-if analysis, you can compare many different scenarios
// against each other.
//
// Note that a what-if analysis is meant to display what the forecasting model has
// learned and how it will behave in the scenarios that you are evaluating. Do not
// blindly use the results of the what-if analysis to make business decisions. For
// instance, forecasts might not be accurate for novel scenarios where there is no
// reference available to determine whether a forecast is good.
//
// The TimeSeriesSelector object defines the items that you want in the what-if analysis.
func (c *Client) CreateWhatIfAnalysis(ctx context.Context, params *CreateWhatIfAnalysisInput, optFns ...func(*Options)) (*CreateWhatIfAnalysisOutput, error) {
	if params == nil {
		params = &CreateWhatIfAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWhatIfAnalysis", params, optFns, c.addOperationCreateWhatIfAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWhatIfAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWhatIfAnalysisInput struct {

	// The Amazon Resource Name (ARN) of the baseline forecast.
	//
	// This member is required.
	ForecastArn *string

	// The name of the what-if analysis. Each name must be unique.
	//
	// This member is required.
	WhatIfAnalysisName *string

	// A list of [tags] to apply to the what if forecast.
	//
	// [tags]: https://docs.aws.amazon.com/forecast/latest/dg/tagging-forecast-resources.html
	Tags []types.Tag

	// Defines the set of time series that are used in the what-if analysis with a
	// TimeSeriesIdentifiers object. What-if analyses are performed only for the time
	// series in this object.
	//
	// The TimeSeriesIdentifiers object needs the following information:
	//
	//   - DataSource
	//
	//   - Format
	//
	//   - Schema
	TimeSeriesSelector *types.TimeSeriesSelector

	noSmithyDocumentSerde
}

type CreateWhatIfAnalysisOutput struct {

	// The Amazon Resource Name (ARN) of the what-if analysis.
	WhatIfAnalysisArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWhatIfAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWhatIfAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWhatIfAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWhatIfAnalysis"); err != nil {
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
	if err = addOpCreateWhatIfAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWhatIfAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWhatIfAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWhatIfAnalysis",
	}
}
