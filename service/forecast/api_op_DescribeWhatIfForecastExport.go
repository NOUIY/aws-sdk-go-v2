// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the what-if forecast export created using the CreateWhatIfForecastExport operation.
//
// In addition to listing the properties provided in the CreateWhatIfForecastExport
// request, this operation lists the following properties:
//
//   - CreationTime
//
//   - LastModificationTime
//
//   - Message - If an error occurred, information about the error.
//
//   - Status
func (c *Client) DescribeWhatIfForecastExport(ctx context.Context, params *DescribeWhatIfForecastExportInput, optFns ...func(*Options)) (*DescribeWhatIfForecastExportOutput, error) {
	if params == nil {
		params = &DescribeWhatIfForecastExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWhatIfForecastExport", params, optFns, c.addOperationDescribeWhatIfForecastExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWhatIfForecastExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWhatIfForecastExportInput struct {

	// The Amazon Resource Name (ARN) of the what-if forecast export that you are
	// interested in.
	//
	// This member is required.
	WhatIfForecastExportArn *string

	noSmithyDocumentSerde
}

type DescribeWhatIfForecastExportOutput struct {

	// When the what-if forecast export was created.
	CreationTime *time.Time

	// The destination for an export job. Provide an S3 path, an Identity and Access
	// Management (IAM) role that allows Amazon Forecast to access the location, and an
	// Key Management Service (KMS) key (optional).
	Destination *types.DataDestination

	// The approximate time remaining to complete the what-if forecast export, in
	// minutes.
	EstimatedTimeRemainingInMinutes *int64

	// The format of the exported data, CSV or PARQUET.
	Format *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	//   - CREATE_PENDING - The CreationTime .
	//
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//
	//   - CREATE_STOPPING - The current timestamp.
	//
	//   - CREATE_STOPPED - When the job stopped.
	//
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, an informational message about the error.
	Message *string

	// The status of the what-if forecast. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - CREATE_STOPPING , CREATE_STOPPED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//
	// The Status of the what-if forecast export must be ACTIVE before you can access
	// the forecast export.
	Status *string

	// An array of Amazon Resource Names (ARNs) that represent all of the what-if
	// forecasts exported in this resource.
	WhatIfForecastArns []string

	// The Amazon Resource Name (ARN) of the what-if forecast export.
	WhatIfForecastExportArn *string

	// The name of the what-if forecast export.
	WhatIfForecastExportName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWhatIfForecastExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWhatIfForecastExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWhatIfForecastExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeWhatIfForecastExport"); err != nil {
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
	if err = addOpDescribeWhatIfForecastExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWhatIfForecastExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWhatIfForecastExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeWhatIfForecastExport",
	}
}
