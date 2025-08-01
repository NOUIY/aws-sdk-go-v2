// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
//
// Adds or updates user-defined contact information associated with the specified
// contact. At least one field to be updated must be present in the request.
//
// You can add or update user-defined contact information for both ongoing and
// completed contacts.
func (c *Client) UpdateContact(ctx context.Context, params *UpdateContactInput, optFns ...func(*Options)) (*UpdateContactOutput, error) {
	if params == nil {
		params = &UpdateContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContact", params, optFns, c.addOperationUpdateContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContactInput struct {

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with your contact center.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	//
	// This member is required.
	InstanceId *string

	// The endpoint of the customer for which the contact was initiated. For external
	// audio contacts, this is usually the end customer's phone number. This value can
	// only be updated for external audio contacts. For more information, see [Amazon Connect Contact Lens integration]in the
	// Amazon Connect Administrator Guide.
	//
	// [Amazon Connect Contact Lens integration]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-integration.html
	CustomerEndpoint *types.Endpoint

	// The description of the contact.
	Description *string

	// The name of the contact.
	Name *string

	//  Information about the queue associated with a contact. This parameter can only
	// be updated for external audio contacts. It is used when you integrate
	// third-party systems with Contact Lens for analytics. For more information, see [Amazon Connect Contact Lens integration]
	// in the Amazon Connect Administrator Guide.
	//
	// [Amazon Connect Contact Lens integration]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-integration.html
	QueueInfo *types.QueueInfoInput

	// Well-formed data on contact, shown to agents on Contact Control Panel (CCP).
	References map[string]types.Reference

	// A set of system defined key-value pairs stored on individual contact segments
	// (unique contact ID) using an attribute map. The attributes are standard Amazon
	// Connect attributes. They can be accessed in flows.
	//
	// Attribute keys can include only alphanumeric, -, and _.
	//
	// This field can be used to show channel subtype, such as connect:Guide .
	//
	// Currently Contact Expiry is the only segment attribute which can be updated by
	// using the UpdateContact API.
	SegmentAttributes map[string]types.SegmentAttributeValue

	// External system endpoint for the contact was initiated. For external audio
	// contacts, this is the phone number of the external system such as the contact
	// center. This value can only be updated for external audio contacts. For more
	// information, see [Amazon Connect Contact Lens integration]in the Amazon Connect Administrator Guide.
	//
	// [Amazon Connect Contact Lens integration]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-integration.html
	SystemEndpoint *types.Endpoint

	// Information about the agent associated with a contact. This parameter can only
	// be updated for external audio contacts. It is used when you integrate
	// third-party systems with Contact Lens for analytics. For more information, see [Amazon Connect Contact Lens integration]
	// in the Amazon Connect Administrator Guide.
	//
	// [Amazon Connect Contact Lens integration]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-integration.html
	UserInfo *types.UserInfo

	noSmithyDocumentSerde
}

type UpdateContactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContact"); err != nil {
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
	if err = addOpUpdateContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContact",
	}
}
