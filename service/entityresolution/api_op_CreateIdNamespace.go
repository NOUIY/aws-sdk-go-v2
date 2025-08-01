// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an ID namespace object which will help customers provide metadata
// explaining their dataset and how to use it. Each ID namespace must have a unique
// name. To modify an existing ID namespace, use the UpdateIdNamespace API.
func (c *Client) CreateIdNamespace(ctx context.Context, params *CreateIdNamespaceInput, optFns ...func(*Options)) (*CreateIdNamespaceOutput, error) {
	if params == nil {
		params = &CreateIdNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdNamespace", params, optFns, c.addOperationCreateIdNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIdNamespaceInput struct {

	// The name of the ID namespace.
	//
	// This member is required.
	IdNamespaceName *string

	// The type of ID namespace. There are two types: SOURCE and TARGET .
	//
	// The SOURCE contains configurations for sourceId data that will be processed in
	// an ID mapping workflow.
	//
	// The TARGET contains a configuration of targetId to which all sourceIds will
	// resolve to.
	//
	// This member is required.
	Type types.IdNamespaceType

	// The description of the ID namespace.
	Description *string

	// Determines the properties of IdMappingWorflow where this IdNamespace can be
	// used as a Source or a Target .
	IdMappingWorkflowProperties []types.IdNamespaceIdMappingWorkflowProperties

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	InputSourceConfig []types.IdNamespaceInputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access the resources defined in this IdNamespace on your behalf as part
	// of the workflow run.
	RoleArn *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateIdNamespaceOutput struct {

	// The timestamp of when the ID namespace was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the ID namespace.
	//
	// This member is required.
	IdNamespaceArn *string

	// The name of the ID namespace.
	//
	// This member is required.
	IdNamespaceName *string

	// The type of ID namespace. There are two types: SOURCE and TARGET .
	//
	// The SOURCE contains configurations for sourceId data that will be processed in
	// an ID mapping workflow.
	//
	// The TARGET contains a configuration of targetId to which all sourceIds will
	// resolve to.
	//
	// This member is required.
	Type types.IdNamespaceType

	// The timestamp of when the ID namespace was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The description of the ID namespace.
	Description *string

	// Determines the properties of IdMappingWorkflow where this IdNamespace can be
	// used as a Source or a Target .
	IdMappingWorkflowProperties []types.IdNamespaceIdMappingWorkflowProperties

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	InputSourceConfig []types.IdNamespaceInputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access the resources defined in inputSourceConfig on your behalf as
	// part of the workflow run.
	RoleArn *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIdNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIdNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIdNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIdNamespace"); err != nil {
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
	if err = addOpCreateIdNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIdNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIdNamespace",
	}
}
