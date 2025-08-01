// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing IdMappingWorkflow . This method is identical to
// CreateIdMappingWorkflow, except it uses an HTTP PUT request instead of a POST
// request, and the IdMappingWorkflow must already exist for the method to succeed.
func (c *Client) UpdateIdMappingWorkflow(ctx context.Context, params *UpdateIdMappingWorkflowInput, optFns ...func(*Options)) (*UpdateIdMappingWorkflowOutput, error) {
	if params == nil {
		params = &UpdateIdMappingWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIdMappingWorkflow", params, optFns, c.addOperationUpdateIdMappingWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIdMappingWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIdMappingWorkflowInput struct {

	// An object which defines the ID mapping technique and any additional
	// configurations.
	//
	// This member is required.
	IdMappingTechniques *types.IdMappingTechniques

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.IdMappingWorkflowInputSource

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// A list of OutputSource objects, each of which contains fields outputS3Path and
	// KMSArn .
	OutputSourceConfig []types.IdMappingWorkflowOutputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access Amazon Web Services resources on your behalf.
	RoleArn *string

	noSmithyDocumentSerde
}

type UpdateIdMappingWorkflowOutput struct {

	// An object which defines the ID mapping technique and any additional
	// configurations.
	//
	// This member is required.
	IdMappingTechniques *types.IdMappingTechniques

	// A list of InputSource objects, which have the fields InputSourceARN and
	// SchemaName .
	//
	// This member is required.
	InputSourceConfig []types.IdMappingWorkflowInputSource

	// The Amazon Resource Name (ARN) of the workflow role. Entity Resolution assumes
	// this role to access Amazon Web Services resources on your behalf.
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	// A description of the workflow.
	Description *string

	// A list of OutputSource objects, each of which contains fields outputS3Path and
	// KMSArn .
	OutputSourceConfig []types.IdMappingWorkflowOutputSource

	// The Amazon Resource Name (ARN) of the IAM role. Entity Resolution assumes this
	// role to access Amazon Web Services resources on your behalf.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIdMappingWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIdMappingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIdMappingWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateIdMappingWorkflow"); err != nil {
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
	if err = addOpUpdateIdMappingWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIdMappingWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIdMappingWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateIdMappingWorkflow",
	}
}
