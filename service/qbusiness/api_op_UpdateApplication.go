// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Q Business application.
//
// Amazon Q Business applications may securely transmit data for processing across
// Amazon Web Services Regions within your geography. For more information, see [Cross region inference in Amazon Q Business].
//
// An Amazon Q Apps service-linked role will be created if it's absent in the
// Amazon Web Services account when QAppsConfiguration is enabled in the request.
// For more information, see [Using service-linked roles for Q Apps].
//
// [Using service-linked roles for Q Apps]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html
// [Cross region inference in Amazon Q Business]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html
func (c *Client) UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) {
	if params == nil {
		params = &UpdateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplication", params, optFns, c.addOperationUpdateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationInput struct {

	// The identifier of the Amazon Q Business application.
	//
	// This member is required.
	ApplicationId *string

	// An option to allow end users to upload files directly during chat.
	AttachmentsConfiguration *types.AttachmentsConfiguration

	// An option to enable updating the default subscription type assigned to an
	// Amazon Q Business application using IAM identity federation for user management.
	AutoSubscriptionConfiguration *types.AutoSubscriptionConfiguration

	// A description for the Amazon Q Business application.
	Description *string

	// A name for the Amazon Q Business application.
	DisplayName *string

	//  The Amazon Resource Name (ARN) of the IAM Identity Center instance you are
	// either creating for—or connecting to—your Amazon Q Business application.
	IdentityCenterInstanceArn *string

	// Configuration information about chat response personalization. For more
	// information, see [Personalizing chat responses].
	//
	// [Personalizing chat responses]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html
	PersonalizationConfiguration *types.PersonalizationConfiguration

	// An option to allow end users to create and use Amazon Q Apps in the web
	// experience.
	QAppsConfiguration *types.QAppsConfiguration

	// An Amazon Web Services Identity and Access Management (IAM) role that gives
	// Amazon Q Business permission to access Amazon CloudWatch logs and metrics.
	RoleArn *string

	noSmithyDocumentSerde
}

type UpdateApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateApplication"); err != nil {
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
	if err = addOpUpdateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateApplication",
	}
}
