// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the asynchronous detection of text in a document. Amazon Textract can
// detect lines of text and the words that make up a line of text.
//
// StartDocumentTextDetection can analyze text in documents that are in JPEG, PNG,
// TIFF, and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocationto
// specify the bucket name and file name of the document.
//
// StartDocumentTextDetection returns a job identifier ( JobId ) that you use to
// get the results of the operation. When text detection is finished, Amazon
// Textract publishes a completion status to the Amazon Simple Notification Service
// (Amazon SNS) topic that you specify in NotificationChannel . To get the results
// of the text detection operation, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetDocumentTextDetection, and pass the job identifier (
// JobId ) from the initial call to StartDocumentTextDetection .
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
func (c *Client) StartDocumentTextDetection(ctx context.Context, params *StartDocumentTextDetectionInput, optFns ...func(*Options)) (*StartDocumentTextDetectionOutput, error) {
	if params == nil {
		params = &StartDocumentTextDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDocumentTextDetection", params, optFns, c.addOperationStartDocumentTextDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDocumentTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDocumentTextDetectionInput struct {

	// The location of the document to be processed.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// The idempotent token that's used to identify the start request. If you use the
	// same token with multiple StartDocumentTextDetection requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see [Calling Amazon Textract Asynchronous Operations].
	//
	// [Calling Amazon Textract Asynchronous Operations]: https://docs.aws.amazon.com/textract/latest/dg/api-async.html
	ClientRequestToken *string

	// An identifier that you specify that's included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The KMS key used to encrypt the inference results. This can be in either Key ID
	// or Key Alias format. When a KMS key is provided, the KMS key will be used for
	// server-side encryption of the objects in the customer bucket. When this
	// parameter is not enabled, the result will be encrypted server side,using SSE-S3.
	KMSKeyId *string

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the
	// completion status of the operation to.
	NotificationChannel *types.NotificationChannel

	// Sets if the output will go to a customer defined bucket. By default Amazon
	// Textract will save the results internally to be accessed with the
	// GetDocumentTextDetection operation.
	OutputConfig *types.OutputConfig

	noSmithyDocumentSerde
}

type StartDocumentTextDetectionOutput struct {

	// The identifier of the text detection job for the document. Use JobId to
	// identify the job in a subsequent call to GetDocumentTextDetection . A JobId
	// value is only valid for 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDocumentTextDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDocumentTextDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDocumentTextDetection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartDocumentTextDetection"); err != nil {
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
	if err = addOpStartDocumentTextDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDocumentTextDetection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDocumentTextDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDocumentTextDetection",
	}
}
