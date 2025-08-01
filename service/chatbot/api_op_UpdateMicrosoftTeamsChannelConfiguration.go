// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chatbot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Microsoft Teams channel configuration.
func (c *Client) UpdateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *UpdateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*UpdateMicrosoftTeamsChannelConfigurationOutput, error) {
	if params == nil {
		params = &UpdateMicrosoftTeamsChannelConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMicrosoftTeamsChannelConfiguration", params, optFns, c.addOperationUpdateMicrosoftTeamsChannelConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMicrosoftTeamsChannelConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMicrosoftTeamsChannelConfigurationInput struct {

	// The ID of the Microsoft Teams channel.
	//
	// This member is required.
	ChannelId *string

	// The Amazon Resource Name (ARN) of the TeamsChannelConfiguration to update.
	//
	// This member is required.
	ChatConfigurationArn *string

	// The name of the Microsoft Teams channel.
	ChannelName *string

	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS
	// managed AdministratorAccess policy is applied by default if this is not set.
	GuardrailPolicyArns []string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	IamRoleArn *string

	// Logging levels include ERROR , INFO , or NONE .
	LoggingLevel *string

	// The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications
	// to AWS Chatbot.
	SnsTopicArns []string

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

type UpdateMicrosoftTeamsChannelConfigurationOutput struct {

	// The configuration for a Microsoft Teams channel configured with AWS Chatbot.
	ChannelConfiguration *types.TeamsChannelConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMicrosoftTeamsChannelConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMicrosoftTeamsChannelConfiguration"); err != nil {
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
	if err = addOpUpdateMicrosoftTeamsChannelConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMicrosoftTeamsChannelConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMicrosoftTeamsChannelConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMicrosoftTeamsChannelConfiguration",
	}
}
