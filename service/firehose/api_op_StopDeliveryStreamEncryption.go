// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables server-side encryption (SSE) for the Firehose stream.
//
// This operation is asynchronous. It returns immediately. When you invoke it,
// Firehose first sets the encryption status of the stream to DISABLING , and then
// to DISABLED . You can continue to read and write data to your stream while its
// status is DISABLING . It can take up to 5 seconds after the encryption status
// changes to DISABLED before all records written to the Firehose stream are no
// longer subject to encryption. To find out whether a record or a batch of records
// was encrypted, check the response elements PutRecordOutput$Encryptedand PutRecordBatchOutput$Encrypted, respectively.
//
// To check the encryption state of a Firehose stream, use DescribeDeliveryStream.
//
// If SSE is enabled using a customer managed CMK and then you invoke
// StopDeliveryStreamEncryption , Firehose schedules the related KMS grant for
// retirement and then retires it after it ensures that it is finished delivering
// records to the destination.
//
// The StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations
// have a combined limit of 25 calls per Firehose stream per 24 hours. For example,
// you reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same Firehose stream in a 24-hour
// period.
func (c *Client) StopDeliveryStreamEncryption(ctx context.Context, params *StopDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StopDeliveryStreamEncryptionOutput, error) {
	if params == nil {
		params = &StopDeliveryStreamEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDeliveryStreamEncryption", params, optFns, c.addOperationStopDeliveryStreamEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDeliveryStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDeliveryStreamEncryptionInput struct {

	// The name of the Firehose stream for which you want to disable server-side
	// encryption (SSE).
	//
	// This member is required.
	DeliveryStreamName *string

	noSmithyDocumentSerde
}

type StopDeliveryStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopDeliveryStreamEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDeliveryStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDeliveryStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopDeliveryStreamEncryption"); err != nil {
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
	if err = addOpStopDeliveryStreamEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDeliveryStreamEncryption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopDeliveryStreamEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopDeliveryStreamEncryption",
	}
}
