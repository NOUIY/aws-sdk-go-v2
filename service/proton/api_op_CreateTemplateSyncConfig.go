// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set up a template to create new template versions automatically by tracking a
// linked repository. A linked repository is a repository that has been registered
// with Proton. For more information, see CreateRepository.
//
// When a commit is pushed to your linked repository, Proton checks for changes to
// your repository template bundles. If it detects a template bundle change, a new
// major or minor version of its template is created, if the version doesn’t
// already exist. For more information, see [Template sync configurations]in the Proton User Guide.
//
// [Template sync configurations]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-sync-configs.html
func (c *Client) CreateTemplateSyncConfig(ctx context.Context, params *CreateTemplateSyncConfigInput, optFns ...func(*Options)) (*CreateTemplateSyncConfigOutput, error) {
	if params == nil {
		params = &CreateTemplateSyncConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTemplateSyncConfig", params, optFns, c.addOperationCreateTemplateSyncConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTemplateSyncConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTemplateSyncConfigInput struct {

	// The repository branch for your template.
	//
	// This member is required.
	Branch *string

	// The repository name (for example, myrepos/myrepo ).
	//
	// This member is required.
	RepositoryName *string

	// The provider type for your repository.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The name of your registered template.
	//
	// This member is required.
	TemplateName *string

	// The type of the registered template.
	//
	// This member is required.
	TemplateType types.TemplateType

	// A repository subdirectory path to your template bundle directory. When
	// included, Proton limits the template bundle search to this repository directory.
	Subdirectory *string

	noSmithyDocumentSerde
}

type CreateTemplateSyncConfigOutput struct {

	// The template sync configuration detail data that's returned by Proton.
	TemplateSyncConfig *types.TemplateSyncConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTemplateSyncConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTemplateSyncConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTemplateSyncConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTemplateSyncConfig"); err != nil {
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
	if err = addOpCreateTemplateSyncConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTemplateSyncConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTemplateSyncConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTemplateSyncConfig",
	}
}
