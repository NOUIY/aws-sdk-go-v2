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

// Starts the classification and analysis of an input document.
// StartLendingAnalysis initiates the classification and analysis of a packet of
// lending documents. StartLendingAnalysis operates on a document file located in
// an Amazon S3 bucket.
//
// StartLendingAnalysis can analyze text in documents that are in one of the
// following formats: JPEG, PNG, TIFF, PDF. Use DocumentLocation to specify the
// bucket name and the file name of the document.
//
// StartLendingAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When the text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If the status is SUCCEEDED you can call either
// GetLendingAnalysis or GetLendingAnalysisSummary and provide the JobId to obtain
// the results of the analysis.
//
// If using OutputConfig to specify an Amazon S3 bucket, the output will be
// contained within the specified prefix in a directory labeled with the job-id. In
// the directory there are 3 sub-directories:
//
//   - detailedResponse (contains the GetLendingAnalysis response)
//
//   - summaryResponse (for the GetLendingAnalysisSummary response)
//
//   - splitDocuments (documents split across logical boundaries)
func (c *Client) StartLendingAnalysis(ctx context.Context, params *StartLendingAnalysisInput, optFns ...func(*Options)) (*StartLendingAnalysisOutput, error) {
	if params == nil {
		params = &StartLendingAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLendingAnalysis", params, optFns, c.addOperationStartLendingAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLendingAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLendingAnalysisInput struct {

	// The Amazon S3 bucket that contains the document to be processed. It's used by
	// asynchronous operations.
	//
	// The input document can be an image file in JPEG or PNG format. It can also be a
	// file in PDF format.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// The idempotent token that you use to identify the start request. If you use the
	// same token with multiple StartLendingAnalysis requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see [Calling Amazon Textract Asynchronous Operations].
	//
	// [Calling Amazon Textract Asynchronous Operations]: https://docs.aws.amazon.com/textract/latest/dg/api-sync.html
	ClientRequestToken *string

	// An identifier that you specify to be included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The KMS key used to encrypt the inference results. This can be in either Key ID
	// or Key Alias format. When a KMS key is provided, the KMS key will be used for
	// server-side encryption of the objects in the customer bucket. When this
	// parameter is not enabled, the result will be encrypted server side, using
	// SSE-S3.
	KMSKeyId *string

	// The Amazon Simple Notification Service (Amazon SNS) topic to which Amazon
	// Textract publishes the completion status of an asynchronous document operation.
	NotificationChannel *types.NotificationChannel

	// Sets whether or not your output will go to a user created bucket. Used to set
	// the name of the bucket, and the prefix on the output file.
	//
	// OutputConfig is an optional parameter which lets you adjust where your output
	// will be placed. By default, Amazon Textract will store the results internally
	// and can only be accessed by the Get API operations. With OutputConfig enabled,
	// you can set the name of the bucket the output will be sent to the file prefix of
	// the results where you can download your results. Additionally, you can set the
	// KMSKeyID parameter to a customer master key (CMK) to encrypt your output.
	// Without this parameter set Amazon Textract will encrypt server-side using the
	// AWS managed CMK for Amazon S3.
	//
	// Decryption of Customer Content is necessary for processing of the documents by
	// Amazon Textract. If your account is opted out under an AI services opt out
	// policy then all unencrypted Customer Content is immediately and permanently
	// deleted after the Customer Content has been processed by the service. No copy of
	// of the output is retained by Amazon Textract. For information about how to opt
	// out, see [Managing AI services opt-out policy.]
	//
	// For more information on data privacy, see the [Data Privacy FAQ].
	//
	// [Managing AI services opt-out policy.]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html
	// [Data Privacy FAQ]: https://aws.amazon.com/compliance/data-privacy-faq/
	OutputConfig *types.OutputConfig

	noSmithyDocumentSerde
}

type StartLendingAnalysisOutput struct {

	// A unique identifier for the lending or text-detection job. The JobId is
	// returned from StartLendingAnalysis . A JobId value is only valid for 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartLendingAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartLendingAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLendingAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartLendingAnalysis"); err != nil {
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
	if err = addOpStartLendingAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLendingAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartLendingAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartLendingAnalysis",
	}
}
