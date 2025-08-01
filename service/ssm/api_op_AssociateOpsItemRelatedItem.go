// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a related item to a Systems Manager OpsCenter OpsItem. For example,
// you can associate an Incident Manager incident or analysis with an OpsItem.
// Incident Manager and OpsCenter are tools in Amazon Web Services Systems Manager.
func (c *Client) AssociateOpsItemRelatedItem(ctx context.Context, params *AssociateOpsItemRelatedItemInput, optFns ...func(*Options)) (*AssociateOpsItemRelatedItemOutput, error) {
	if params == nil {
		params = &AssociateOpsItemRelatedItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateOpsItemRelatedItem", params, optFns, c.addOperationAssociateOpsItemRelatedItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateOpsItemRelatedItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateOpsItemRelatedItemInput struct {

	// The type of association that you want to create between an OpsItem and a
	// resource. OpsCenter supports IsParentOf and RelatesTo association types.
	//
	// This member is required.
	AssociationType *string

	// The ID of the OpsItem to which you want to associate a resource as a related
	// item.
	//
	// This member is required.
	OpsItemId *string

	// The type of resource that you want to associate with an OpsItem. OpsCenter
	// supports the following types:
	//
	// AWS::SSMIncidents::IncidentRecord : an Incident Manager incident.
	//
	// AWS::SSM::Document : a Systems Manager (SSM) document.
	//
	// This member is required.
	ResourceType *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services resource that you
	// want to associate with the OpsItem.
	//
	// This member is required.
	ResourceUri *string

	noSmithyDocumentSerde
}

type AssociateOpsItemRelatedItemOutput struct {

	// The association ID.
	AssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateOpsItemRelatedItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateOpsItemRelatedItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateOpsItemRelatedItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateOpsItemRelatedItem"); err != nil {
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
	if err = addOpAssociateOpsItemRelatedItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateOpsItemRelatedItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateOpsItemRelatedItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateOpsItemRelatedItem",
	}
}
