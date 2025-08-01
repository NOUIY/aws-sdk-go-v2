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

// Describes a forecast export job created using the CreateForecastExportJob operation.
//
// In addition to listing the properties provided by the user in the
// CreateForecastExportJob request, this operation lists the following properties:
//
//   - CreationTime
//
//   - LastModificationTime
//
//   - Status
//
//   - Message - If an error occurred, information about the error.
func (c *Client) DescribeForecastExportJob(ctx context.Context, params *DescribeForecastExportJobInput, optFns ...func(*Options)) (*DescribeForecastExportJobOutput, error) {
	if params == nil {
		params = &DescribeForecastExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeForecastExportJob", params, optFns, c.addOperationDescribeForecastExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeForecastExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeForecastExportJobInput struct {

	// The Amazon Resource Name (ARN) of the forecast export job.
	//
	// This member is required.
	ForecastExportJobArn *string

	noSmithyDocumentSerde
}

type DescribeForecastExportJobOutput struct {

	// When the forecast export job was created.
	CreationTime *time.Time

	// The path to the Amazon Simple Storage Service (Amazon S3) bucket where the
	// forecast is exported.
	Destination *types.DataDestination

	// The Amazon Resource Name (ARN) of the exported forecast.
	ForecastArn *string

	// The ARN of the forecast export job.
	ForecastExportJobArn *string

	// The name of the forecast export job.
	ForecastExportJobName *string

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

	// The status of the forecast export job. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - CREATE_STOPPING , CREATE_STOPPED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//
	// The Status of the forecast export job must be ACTIVE before you can access the
	// forecast in your S3 bucket.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeForecastExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeForecastExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeForecastExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeForecastExportJob"); err != nil {
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
	if err = addOpDescribeForecastExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeForecastExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeForecastExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeForecastExportJob",
	}
}
