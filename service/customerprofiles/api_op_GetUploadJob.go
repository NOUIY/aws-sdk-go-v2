// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This API retrieves the details of a specific upload job.
func (c *Client) GetUploadJob(ctx context.Context, params *GetUploadJobInput, optFns ...func(*Options)) (*GetUploadJobOutput, error) {
	if params == nil {
		params = &GetUploadJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUploadJob", params, optFns, c.addOperationGetUploadJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUploadJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUploadJobInput struct {

	// The unique name of the domain containing the upload job.
	//
	// This member is required.
	DomainName *string

	// The unique identifier of the upload job to retrieve.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetUploadJobOutput struct {

	// The timestamp when the upload job was completed.
	CompletedAt *time.Time

	// The timestamp when the upload job was created.
	CreatedAt *time.Time

	// The expiry duration for the profiles ingested with the upload job.
	DataExpiry *int32

	// The unique name of the upload job. Could be a file name to identify the upload
	// job.
	DisplayName *string

	// The mapping between CSV Columns and Profile Object attributes for the upload
	// job.
	Fields map[string]types.ObjectTypeField

	// The unique identifier of the upload job.
	JobId *string

	// The summary of results for the upload job, including the number of updated,
	// created, and failed records.
	ResultsSummary *types.ResultsSummary

	// The status describing the status for the upload job. The following are Valid
	// Values:
	//
	//   - CREATED: The upload job has been created, but has not started processing
	//   yet.
	//
	//   - IN_PROGRESS: The upload job is currently in progress, ingesting and
	//   processing the profile data.
	//
	//   - PARTIALLY_SUCCEEDED: The upload job has successfully completed the
	//   ingestion and processing of all profile data.
	//
	//   - SUCCEEDED: The upload job has successfully completed the ingestion and
	//   processing of all profile data.
	//
	//   - FAILED: The upload job has failed to complete.
	//
	//   - STOPPED: The upload job has been manually stopped or terminated before
	//   completion.
	Status types.UploadJobStatus

	// The reason for the current status of the upload job. Possible reasons:
	//
	//   - VALIDATION_FAILURE: The upload job has encountered an error or issue and
	//   was unable to complete the profile data ingestion.
	//
	//   - INTERNAL_FAILURE: Failure caused from service side
	StatusReason types.StatusReason

	// The unique key columns used for de-duping the keys in the upload job.
	UniqueKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUploadJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUploadJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUploadJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUploadJob"); err != nil {
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
	if err = addOpGetUploadJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUploadJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUploadJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUploadJob",
	}
}
