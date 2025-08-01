// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Writes multiple data records into a Firehose stream in a single call, which can
// achieve higher throughput per producer than when writing single records. To
// write single data records into a Firehose stream, use PutRecord. Applications using
// these operations are referred to as producers.
//
// Firehose accumulates and publishes a particular metric for a customer account
// in one minute intervals. It is possible that the bursts of incoming
// bytes/records ingested to a Firehose stream last only for a few seconds. Due to
// this, the actual spikes in the traffic might not be fully visible in the
// customer's 1 minute CloudWatch metrics.
//
// For information about service quota, see [Amazon Firehose Quota].
//
// Each PutRecordBatch request supports up to 500 records. Each record in the request can be as
// large as 1,000 KB (before base64 encoding), up to a limit of 4 MB for the entire
// request. These limits cannot be changed.
//
// You must specify the name of the Firehose stream and the data record when using PutRecord
// . The data record consists of a data blob that can be up to 1,000 KB in size,
// and any kind of data. For example, it could be a segment from a log file,
// geographic location data, website clickstream data, and so on.
//
// For multi record de-aggregation, you can not put more than 500 records even if
// the data blob length is less than 1000 KiB. If you include more than 500
// records, the request succeeds but the record de-aggregation doesn't work as
// expected and transformation lambda is invoked with the complete base64 encoded
// data blob instead of de-aggregated base64 decoded records.
//
// Firehose buffers records before delivering them to the destination. To
// disambiguate the data blobs at the destination, a common solution is to use
// delimiters in the data, such as a newline ( \n ) or some other character unique
// within the data. This allows the consumer application to parse individual data
// items when reading the data from the destination.
//
// The PutRecordBatch response includes a count of failed records, FailedPutCount , and an array
// of responses, RequestResponses . Even if the PutRecordBatch call succeeds, the value of
// FailedPutCount may be greater than 0, indicating that there are records for
// which the operation didn't succeed. Each entry in the RequestResponses array
// provides additional information about the processed record. It directly
// correlates with a record in the request array using the same ordering, from the
// top to the bottom. The response array always includes the same number of records
// as the request array. RequestResponses includes both successfully and
// unsuccessfully processed records. Firehose tries to process all records in each PutRecordBatch
// request. A single record failure does not stop the processing of subsequent
// records.
//
// A successfully processed record includes a RecordId value, which is unique for
// the record. An unsuccessfully processed record includes ErrorCode and
// ErrorMessage values. ErrorCode reflects the type of error, and is one of the
// following values: ServiceUnavailableException or InternalFailure . ErrorMessage
// provides more detailed information about the error.
//
// If there is an internal server error or a timeout, the write might have
// completed or it might have failed. If FailedPutCount is greater than 0, retry
// the request, resending only those records that might have failed processing.
// This minimizes the possible duplicate records and also reduces the total bytes
// sent (and corresponding charges). We recommend that you handle any duplicates at
// the destination.
//
// If PutRecordBatch throws ServiceUnavailableException , the API is automatically reinvoked
// (retried) 3 times. If the exception persists, it is possible that the throughput
// limits have been exceeded for the Firehose stream.
//
// Re-invoking the Put API operations (for example, PutRecord and PutRecordBatch)
// can result in data duplicates. For larger data assets, allow for a longer time
// out before retrying Put API operations.
//
// Data records sent to Firehose are stored for 24 hours from the time they are
// added to a Firehose stream as it attempts to send the records to the
// destination. If the destination is unreachable for more than 24 hours, the data
// is no longer available.
//
// Don't concatenate two or more base64 strings to form the data fields of your
// records. Instead, concatenate the raw data, then perform base64 encoding.
//
// [Amazon Firehose Quota]: https://docs.aws.amazon.com/firehose/latest/dev/limits.html
func (c *Client) PutRecordBatch(ctx context.Context, params *PutRecordBatchInput, optFns ...func(*Options)) (*PutRecordBatchOutput, error) {
	if params == nil {
		params = &PutRecordBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecordBatch", params, optFns, c.addOperationPutRecordBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecordBatchInput struct {

	// The name of the Firehose stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// One or more records.
	//
	// This member is required.
	Records []types.Record

	noSmithyDocumentSerde
}

type PutRecordBatchOutput struct {

	// The number of records that might have failed processing. This number might be
	// greater than 0 even if the PutRecordBatchcall succeeds. Check FailedPutCount to determine
	// whether there are records that you need to resend.
	//
	// This member is required.
	FailedPutCount *int32

	// The results array. For each record, the index of the response element is the
	// same as the index used in the request array.
	//
	// This member is required.
	RequestResponses []types.PutRecordBatchResponseEntry

	// Indicates whether server-side encryption (SSE) was enabled during this
	// operation.
	Encrypted *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecordBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecordBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecordBatch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRecordBatch"); err != nil {
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
	if err = addOpPutRecordBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecordBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRecordBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRecordBatch",
	}
}
