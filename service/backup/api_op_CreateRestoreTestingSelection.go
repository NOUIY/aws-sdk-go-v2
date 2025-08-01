// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This request can be sent after CreateRestoreTestingPlan request returns
// successfully. This is the second part of creating a resource testing plan, and
// it must be completed sequentially.
//
// This consists of RestoreTestingSelectionName , ProtectedResourceType , and one
// of the following:
//
//   - ProtectedResourceArns
//
//   - ProtectedResourceConditions
//
// Each protected resource type can have one single value.
//
// A restore testing selection can include a wildcard value ("*") for
// ProtectedResourceArns along with ProtectedResourceConditions . Alternatively,
// you can include up to 30 specific protected resource ARNs in
// ProtectedResourceArns .
//
// Cannot select by both protected resource types AND specific ARNs. Request will
// fail if both are included.
func (c *Client) CreateRestoreTestingSelection(ctx context.Context, params *CreateRestoreTestingSelectionInput, optFns ...func(*Options)) (*CreateRestoreTestingSelectionOutput, error) {
	if params == nil {
		params = &CreateRestoreTestingSelectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRestoreTestingSelection", params, optFns, c.addOperationCreateRestoreTestingSelectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRestoreTestingSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRestoreTestingSelectionInput struct {

	// Input the restore testing plan name that was returned from the related
	// CreateRestoreTestingPlan request.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// This consists of RestoreTestingSelectionName , ProtectedResourceType , and one
	// of the following:
	//
	//   - ProtectedResourceArns
	//
	//   - ProtectedResourceConditions
	//
	// Each protected resource type can have one single value.
	//
	// A restore testing selection can include a wildcard value ("*") for
	// ProtectedResourceArns along with ProtectedResourceConditions . Alternatively,
	// you can include up to 30 specific protected resource ARNs in
	// ProtectedResourceArns .
	//
	// This member is required.
	RestoreTestingSelection *types.RestoreTestingSelectionForCreate

	// This is an optional unique string that identifies the request and allows failed
	// requests to be retried without the risk of running the operation twice. If used,
	// this parameter must contain 1 to 50 alphanumeric or '-_.' characters.
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type CreateRestoreTestingSelectionOutput struct {

	// The time that the resource testing selection was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The ARN of the restore testing plan with which the restore testing selection is
	// associated.
	//
	// This member is required.
	RestoreTestingPlanArn *string

	// The name of the restore testing plan.
	//
	// The name cannot be changed after creation. The name consists of only
	// alphanumeric characters and underscores. Maximum length is 50.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// The name of the restore testing selection for the related restore testing plan.
	//
	// This member is required.
	RestoreTestingSelectionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRestoreTestingSelectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRestoreTestingSelection"); err != nil {
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
	if err = addOpCreateRestoreTestingSelectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRestoreTestingSelection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRestoreTestingSelection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRestoreTestingSelection",
	}
}
