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

// Creates a custom slot type or replaces an existing custom slot type.
//
// To create a custom slot type, specify a name for the slot type and a set of
// enumeration values, which are the values that a slot of this type can assume.
// For more information, see how-it-works.
//
// If you specify the name of an existing slot type, the fields in the request
// replace the existing values in the $LATEST version of the slot type. Amazon Lex
// removes the fields that you don't provide in the request. If you don't specify
// required fields, Amazon Lex throws an exception. When you update the $LATEST
// version of a slot type, if a bot uses the $LATEST version of an intent that
// contains the slot type, the bot's status field is set to NOT_BUILT .
//
// This operation requires permissions for the lex:PutSlotType action.
func (c *Client) PutSlotType(ctx context.Context, params *PutSlotTypeInput, optFns ...func(*Options)) (*PutSlotTypeOutput, error) {
	if params == nil {
		params = &PutSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSlotType", params, optFns, c.addOperationPutSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSlotTypeInput struct {

	// The name of the slot type. The name is not case sensitive.
	//
	// The name can't match a built-in slot type name, or a built-in slot type name
	// with "AMAZON." removed. For example, because there is a built-in slot type
	// called AMAZON.DATE , you can't create a custom slot type called DATE .
	//
	// For a list of built-in slot types, see [Slot Type Reference] in the Alexa Skills Kit.
	//
	// [Slot Type Reference]: https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference
	//
	// This member is required.
	Name *string

	// Identifies a specific revision of the $LATEST version.
	//
	// When you create a new slot type, leave the checksum field blank. If you specify
	// a checksum you get a BadRequestException exception.
	//
	// When you want to update a slot type, set the checksum field to the checksum of
	// the most recent revision of the $LATEST version. If you don't specify the
	// checksum field, or if the checksum does not match the $LATEST version, you get
	// a PreconditionFailedException exception.
	Checksum *string

	// When set to true a new numbered version of the slot type is created. This is
	// the same as calling the CreateSlotTypeVersion operation. If you do not specify
	// createVersion , the default is false .
	CreateVersion *bool

	// A description of the slot type.
	Description *string

	// A list of EnumerationValue objects that defines the values that the slot type
	// can take. Each value can have a list of synonyms , which are additional values
	// that help train the machine learning model about the values that it resolves for
	// a slot.
	//
	// A regular expression slot type doesn't require enumeration values. All other
	// slot types require a list of enumeration values.
	//
	// When Amazon Lex resolves a slot value, it generates a resolution list that
	// contains up to five possible values for the slot. If you are using a Lambda
	// function, this resolution list is passed to the function. If you are not using a
	// Lambda function you can choose to return the value that the user entered or the
	// first value in the resolution list as the slot value. The valueSelectionStrategy
	// field indicates the option to use.
	EnumerationValues []types.EnumerationValue

	// The built-in slot type used as the parent of the slot type. When you define a
	// parent slot type, the new slot type has all of the same configuration as the
	// parent.
	//
	// Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string

	// Configuration information that extends the parent built-in slot type. The
	// configuration is added to the settings for the parent slot type.
	SlotTypeConfigurations []types.SlotTypeConfiguration

	// Determines the slot resolution strategy that Amazon Lex uses to return slot
	// type values. The field can be set to one of the following values:
	//
	//   - ORIGINAL_VALUE - Returns the value entered by the user, if the user value is
	//   similar to the slot value.
	//
	//   - TOP_RESOLUTION - If there is a resolution list for the slot, return the
	//   first value in the resolution list as the slot type value. If there is no
	//   resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy , the default is ORIGINAL_VALUE .
	ValueSelectionStrategy types.SlotValueSelectionStrategy

	noSmithyDocumentSerde
}

type PutSlotTypeOutput struct {

	// Checksum of the $LATEST version of the slot type.
	Checksum *string

	// True if a new version of the slot type was created. If the createVersion field
	// was not specified in the request, the createVersion field is set to false in
	// the response.
	CreateVersion *bool

	// The date that the slot type was created.
	CreatedDate *time.Time

	// A description of the slot type.
	Description *string

	// A list of EnumerationValue objects that defines the values that the slot type
	// can take.
	EnumerationValues []types.EnumerationValue

	// The date that the slot type was updated. When you create a slot type, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time

	// The name of the slot type.
	Name *string

	// The built-in slot type used as the parent of the slot type.
	ParentSlotTypeSignature *string

	// Configuration information that extends the parent built-in slot type.
	SlotTypeConfigurations []types.SlotTypeConfiguration

	// The slot resolution strategy that Amazon Lex uses to determine the value of the
	// slot. For more information, see PutSlotType.
	ValueSelectionStrategy types.SlotValueSelectionStrategy

	// The version of the slot type. For a new slot type, the version is always
	// $LATEST .
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSlotType"); err != nil {
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
	if err = addOpPutSlotTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSlotType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSlotType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSlotType",
	}
}
