// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a URL that you can use to connect to the Jupyter server from a notebook
// instance. In the SageMaker AI console, when you choose Open next to a notebook
// instance, SageMaker AI opens a new tab showing the Jupyter server home page from
// the notebook instance. The console uses this API to get the URL and show the
// page.
//
// The IAM role or user used to call this API defines the permissions to access
// the notebook instance. Once the presigned URL is created, no additional
// permission is required to access this URL. IAM authorization policies for this
// API are also enforced for every HTTP request and WebSocket frame that attempts
// to connect to the notebook instance.
//
// You can restrict access to this API and to the URL that it returns to a list of
// IP addresses that you specify. Use the NotIpAddress condition operator and the
// aws:SourceIP condition context key to specify the list of IP addresses that you
// want to have access to the notebook instance. For more information, see [Limit Access to a Notebook Instance by IP Address].
//
// The URL that you get from a call to [CreatePresignedNotebookInstanceUrl] is valid only for 5 minutes. If you try to
// use the URL after the 5-minute limit expires, you are directed to the Amazon Web
// Services console sign-in page.
//
// [CreatePresignedNotebookInstanceUrl]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreatePresignedNotebookInstanceUrl.html
// [Limit Access to a Notebook Instance by IP Address]: https://docs.aws.amazon.com/sagemaker/latest/dg/security_iam_id-based-policy-examples.html#nbi-ip-filter
func (c *Client) CreatePresignedNotebookInstanceUrl(ctx context.Context, params *CreatePresignedNotebookInstanceUrlInput, optFns ...func(*Options)) (*CreatePresignedNotebookInstanceUrlOutput, error) {
	if params == nil {
		params = &CreatePresignedNotebookInstanceUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePresignedNotebookInstanceUrl", params, optFns, c.addOperationCreatePresignedNotebookInstanceUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePresignedNotebookInstanceUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresignedNotebookInstanceUrlInput struct {

	// The name of the notebook instance.
	//
	// This member is required.
	NotebookInstanceName *string

	// The duration of the session, in seconds. The default is 12 hours.
	SessionExpirationDurationInSeconds *int32

	noSmithyDocumentSerde
}

type CreatePresignedNotebookInstanceUrlOutput struct {

	// A JSON object that contains the URL string.
	AuthorizedUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePresignedNotebookInstanceUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePresignedNotebookInstanceUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePresignedNotebookInstanceUrl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePresignedNotebookInstanceUrl"); err != nil {
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
	if err = addOpCreatePresignedNotebookInstanceUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePresignedNotebookInstanceUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePresignedNotebookInstanceUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePresignedNotebookInstanceUrl",
	}
}
