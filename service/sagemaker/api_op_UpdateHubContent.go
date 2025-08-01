// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates SageMaker hub content (either a Model or Notebook resource).
//
// You can update the metadata that describes the resource. In addition to the
// required request fields, specify at least one of the following fields to update:
//
//   - HubContentDescription
//
//   - HubContentDisplayName
//
//   - HubContentMarkdown
//
//   - HubContentSearchKeywords
//
//   - SupportStatus
//
// For more information about hubs, see [Private curated hubs for foundation model access control in JumpStart].
//
// If you want to update a ModelReference resource in your hub, use the
// UpdateHubContentResource API instead.
//
// [Private curated hubs for foundation model access control in JumpStart]: https://docs.aws.amazon.com/sagemaker/latest/dg/jumpstart-curated-hubs.html
func (c *Client) UpdateHubContent(ctx context.Context, params *UpdateHubContentInput, optFns ...func(*Options)) (*UpdateHubContentOutput, error) {
	if params == nil {
		params = &UpdateHubContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateHubContent", params, optFns, c.addOperationUpdateHubContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateHubContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateHubContentInput struct {

	// The name of the hub content resource that you want to update.
	//
	// This member is required.
	HubContentName *string

	// The content type of the resource that you want to update. Only specify a Model
	// or Notebook resource for this API. To update a ModelReference , use the
	// UpdateHubContentReference API instead.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The hub content version that you want to update. For example, if you have two
	// versions of a resource in your hub, you can update the second version.
	//
	// This member is required.
	HubContentVersion *string

	// The name of the SageMaker hub that contains the hub content you want to update.
	// You can optionally use the hub ARN instead.
	//
	// This member is required.
	HubName *string

	// The description of the hub content.
	HubContentDescription *string

	// The display name of the hub content.
	HubContentDisplayName *string

	// A string that provides a description of the hub content. This string can
	// include links, tables, and standard markdown formatting.
	HubContentMarkdown *string

	// The searchable keywords of the hub content.
	HubContentSearchKeywords []string

	// Indicates the current status of the hub content resource.
	SupportStatus types.HubContentSupportStatus

	noSmithyDocumentSerde
}

type UpdateHubContentOutput struct {

	// The ARN of the private model hub that contains the updated hub content.
	//
	// This member is required.
	HubArn *string

	// The ARN of the hub content resource that was updated.
	//
	// This member is required.
	HubContentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateHubContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateHubContent"); err != nil {
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
	if err = addOpUpdateHubContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHubContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateHubContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateHubContent",
	}
}
