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

// Creates an AWS Chatbot configuration for Microsoft Teams.
func (c *Client) CreateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *CreateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*CreateMicrosoftTeamsChannelConfigurationOutput, error) {
	if params == nil {
		params = &CreateMicrosoftTeamsChannelConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMicrosoftTeamsChannelConfiguration", params, optFns, c.addOperationCreateMicrosoftTeamsChannelConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMicrosoftTeamsChannelConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMicrosoftTeamsChannelConfigurationInput struct {

	// The ID of the Microsoft Teams channel.
	//
	// This member is required.
	ChannelId *string

	// The name of the configuration.
	//
	// This member is required.
	ConfigurationName *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	//  The ID of the Microsoft Teams authorized with AWS Chatbot.
	//
	// To get the team ID, you must perform the initial authorization flow with
	// Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team
	// ID from the console. For more information, see [Step 1: Configure a Microsoft Teams client]in the AWS Chatbot Administrator
	// Guide.
	//
	// [Step 1: Configure a Microsoft Teams client]: https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup
	//
	// This member is required.
	TeamId *string

	// The ID of the Microsoft Teams tenant.
	//
	// This member is required.
	TenantId *string

	// The name of the Microsoft Teams channel.
	ChannelName *string

	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS
	// managed AdministratorAccess policy is applied by default if this is not set.
	GuardrailPolicyArns []string

	// Logging levels include ERROR , INFO , or NONE .
	LoggingLevel *string

	// The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications
	// to AWS Chatbot.
	SnsTopicArns []string

	// A map of tags assigned to a resource. A tag is a string-to-string map of
	// key-value pairs.
	Tags []types.Tag

	// The name of the Microsoft Teams Team.
	TeamName *string

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

type CreateMicrosoftTeamsChannelConfigurationOutput struct {

	// The configuration for a Microsoft Teams channel configured with AWS Chatbot.
	ChannelConfiguration *types.TeamsChannelConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMicrosoftTeamsChannelConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMicrosoftTeamsChannelConfiguration"); err != nil {
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
	if err = addOpCreateMicrosoftTeamsChannelConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMicrosoftTeamsChannelConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMicrosoftTeamsChannelConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMicrosoftTeamsChannelConfiguration",
	}
}
