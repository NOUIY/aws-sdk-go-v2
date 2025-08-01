// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a job to import a resource to Amazon Lex.
func (c *Client) StartImport(ctx context.Context, params *StartImportInput, optFns ...func(*Options)) (*StartImportOutput, error) {
	if params == nil {
		params = &StartImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImport", params, optFns, c.addOperationStartImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportInput struct {

	// Specifies the action that the StartImport operation should take when there is
	// an existing resource with the same name.
	//
	//   - FAIL_ON_CONFLICT - The import operation is stopped on the first conflict
	//   between a resource in the import file and an existing resource. The name of the
	//   resource causing the conflict is in the failureReason field of the response to
	//   the GetImport operation.
	//
	// OVERWRITE_LATEST - The import operation proceeds even if there is a conflict
	//   with an existing resource. The $LASTEST version of the existing resource is
	//   overwritten with the data from the import file.
	//
	// This member is required.
	MergeStrategy types.MergeStrategy

	// A zip archive in binary format. The archive should contain one file, a JSON
	// file containing the resource to import. The resource should match the type
	// specified in the resourceType field.
	//
	// This member is required.
	Payload []byte

	// Specifies the type of resource to export. Each resource also exports any
	// resources that it depends on.
	//
	//   - A bot exports dependent intents.
	//
	//   - An intent exports dependent slot types.
	//
	// This member is required.
	ResourceType types.ResourceType

	// A list of tags to add to the imported bot. You can only add tags when you
	// import a bot, you can't add tags to an intent or slot type.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartImportOutput struct {

	// A timestamp for the date and time that the import job was requested.
	CreatedDate *time.Time

	// The identifier for the specific import job.
	ImportId *string

	// The status of the import job. If the status is FAILED , you can get the reason
	// for the failure using the GetImport operation.
	ImportStatus types.ImportStatus

	// The action to take when there is a merge conflict.
	MergeStrategy types.MergeStrategy

	// The name given to the import job.
	Name *string

	// The type of resource to import.
	ResourceType types.ResourceType

	// A list of tags added to the imported bot.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartImport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartImport"); err != nil {
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
	if err = addOpStartImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartImport",
	}
}
