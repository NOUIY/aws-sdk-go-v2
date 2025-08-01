// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a configuration profile.
//
// To prevent users from unintentionally deleting actively-used configuration
// profiles, enable [deletion protection].
//
// [deletion protection]: https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
func (c *Client) DeleteConfigurationProfile(ctx context.Context, params *DeleteConfigurationProfileInput, optFns ...func(*Options)) (*DeleteConfigurationProfileOutput, error) {
	if params == nil {
		params = &DeleteConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationProfile", params, optFns, c.addOperationDeleteConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConfigurationProfileInput struct {

	// The application ID that includes the configuration profile you want to delete.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the configuration profile you want to delete.
	//
	// This member is required.
	ConfigurationProfileId *string

	// A parameter to configure deletion protection. Deletion protection prevents a
	// user from deleting a configuration profile if your application has called either
	// [GetLatestConfiguration]or for the configuration profile during the specified interval.
	//
	// This parameter supports the following values:
	//
	//   - BYPASS : Instructs AppConfig to bypass the deletion protection check and
	//   delete a configuration profile even if deletion protection would have otherwise
	//   prevented it.
	//
	//   - APPLY : Instructs the deletion protection check to run, even if deletion
	//   protection is disabled at the account level. APPLY also forces the deletion
	//   protection check to run against resources created in the past hour, which are
	//   normally excluded from deletion protection checks.
	//
	//   - ACCOUNT_DEFAULT : The default setting, which instructs AppConfig to
	//   implement the deletion protection value specified in the UpdateAccountSettings
	//   API.
	//
	// [GetLatestConfiguration]: https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html
	DeletionProtectionCheck types.DeletionProtectionCheck

	noSmithyDocumentSerde
}

type DeleteConfigurationProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteConfigurationProfile"); err != nil {
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
	if err = addOpDeleteConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteConfigurationProfile",
	}
}
