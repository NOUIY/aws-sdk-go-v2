// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/b2bi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified parameters for a transformer. A transformer can take an
// EDI file as input and transform it into a JSON-or XML-formatted document.
// Alternatively, a transformer can take a JSON-or XML-formatted document as input
// and transform it into an EDI file.
func (c *Client) UpdateTransformer(ctx context.Context, params *UpdateTransformerInput, optFns ...func(*Options)) (*UpdateTransformerOutput, error) {
	if params == nil {
		params = &UpdateTransformerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTransformer", params, optFns, c.addOperationUpdateTransformerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTransformerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTransformerInput struct {

	// Specifies the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	// Specifies the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	EdiType types.EdiType

	// Specifies that the currently supported file formats for EDI transformations are
	// JSON and XML .
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	FileFormat types.FileFormat

	// To update, specify the InputConversion object, which contains the format
	// options for the inbound transformation.
	InputConversion *types.InputConversion

	// Specify the structure that contains the mapping template and its language
	// (either XSLT or JSONATA).
	Mapping *types.Mapping

	// Specifies the mapping template for the transformer. This template is used to
	// map the parsed EDI file using JSONata or XSLT.
	//
	// This parameter is available for backwards compatibility. Use the [Mapping] data type
	// instead.
	//
	// [Mapping]: https://docs.aws.amazon.com/b2bi/latest/APIReference/API_Mapping.html
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	MappingTemplate *string

	// Specify a new name for the transformer, if you want to update it.
	Name *string

	// To update, specify the OutputConversion object, which contains the format
	// options for the outbound transformation.
	OutputConversion *types.OutputConversion

	// Specifies a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	SampleDocument *string

	// Specify a structure that contains the Amazon S3 bucket and an array of the
	// corresponding keys used to identify the location for your sample documents.
	SampleDocuments *types.SampleDocuments

	// Specifies the transformer's status. You can update the state of the transformer
	// from inactive to active .
	Status types.TransformerStatus

	noSmithyDocumentSerde
}

type UpdateTransformerOutput struct {

	// Returns a timestamp for creation date and time of the transformer.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns a timestamp for last time the transformer was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// Returns the name of the transformer.
	//
	// This member is required.
	Name *string

	// Returns the state of the newly created transformer. The transformer can be
	// either active or inactive . For the transformer to be used in a capability, its
	// status must active .
	//
	// This member is required.
	Status types.TransformerStatus

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	TransformerArn *string

	// Returns the system-assigned unique identifier for the transformer.
	//
	// This member is required.
	TransformerId *string

	// Returns the details for the EDI standard that is being used for the
	// transformer. Currently, only X12 is supported. X12 is a set of standards and
	// corresponding messages that define specific business documents.
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	EdiType types.EdiType

	// Returns that the currently supported file formats for EDI transformations are
	// JSON and XML .
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	FileFormat types.FileFormat

	// Returns the InputConversion object, which contains the format options for the
	// inbound transformation.
	InputConversion *types.InputConversion

	// Returns the structure that contains the mapping template and its language
	// (either XSLT or JSONATA).
	Mapping *types.Mapping

	// Returns the mapping template for the transformer. This template is used to map
	// the parsed EDI file using JSONata or XSLT.
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	MappingTemplate *string

	// Returns the OutputConversion object, which contains the format options for the
	// outbound transformation.
	OutputConversion *types.OutputConversion

	// Returns a sample EDI document that is used by a transformer as a guide for
	// processing the EDI data.
	//
	// Deprecated: This is a legacy trait. Please use input-conversion or
	// output-conversion.
	SampleDocument *string

	// Returns a structure that contains the Amazon S3 bucket and an array of the
	// corresponding keys used to identify the location for your sample documents.
	SampleDocuments *types.SampleDocuments

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTransformerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTransformer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTransformer"); err != nil {
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
	if err = addOpUpdateTransformerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTransformer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTransformer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTransformer",
	}
}
