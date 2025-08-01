// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates or updates the policy schema in the specified policy store. The schema
// is used to validate any Cedar policies and policy templates submitted to the
// policy store. Any changes to the schema validate only policies and templates
// submitted after the schema change. Existing policies and templates are not
// re-evaluated against the changed schema. If you later update a policy, then it
// is evaluated against the new schema at that time.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
func (c *Client) PutSchema(ctx context.Context, params *PutSchemaInput, optFns ...func(*Options)) (*PutSchemaOutput, error) {
	if params == nil {
		params = &PutSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSchema", params, optFns, c.addOperationPutSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSchemaInput struct {

	// Specifies the definition of the schema to be stored. The schema definition must
	// be written in Cedar schema JSON.
	//
	// This member is required.
	Definition types.SchemaDefinition

	// Specifies the ID of the policy store in which to place the schema.
	//
	// This member is required.
	PolicyStoreId *string

	noSmithyDocumentSerde
}

type PutSchemaOutput struct {

	// The date and time that the schema was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the schema was last updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// Identifies the namespaces of the entities referenced by this schema.
	//
	// This member is required.
	Namespaces []string

	// The unique ID of the policy store that contains the schema.
	//
	// This member is required.
	PolicyStoreId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutSchema{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSchema"); err != nil {
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
	if err = addOpPutSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSchema",
	}
}
