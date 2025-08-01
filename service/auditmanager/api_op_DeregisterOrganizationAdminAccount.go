// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified Amazon Web Services account as a delegated administrator
// for Audit Manager.
//
// When you remove a delegated administrator from your Audit Manager settings, you
// continue to have access to the evidence that you previously collected under that
// account. This is also the case when you deregister a delegated administrator
// from Organizations. However, Audit Manager stops collecting and attaching
// evidence to that delegated administrator account moving forward.
//
// Keep in mind the following cleanup task if you use evidence finder:
//
// Before you use your management account to remove a delegated administrator,
// make sure that the current delegated administrator account signs in to Audit
// Manager and disables evidence finder first. Disabling evidence finder
// automatically deletes the event data store that was created in their account
// when they enabled evidence finder. If this task isn’t completed, the event data
// store remains in their account. In this case, we recommend that the original
// delegated administrator goes to CloudTrail Lake and manually [deletes the event data store].
//
// This cleanup task is necessary to ensure that you don't end up with multiple
// event data stores. Audit Manager ignores an unused event data store after you
// remove or change a delegated administrator account. However, the unused event
// data store continues to incur storage costs from CloudTrail Lake if you don't
// delete it.
//
// When you deregister a delegated administrator account for Audit Manager, the
// data for that account isn’t deleted. If you want to delete resource data for a
// delegated administrator account, you must perform that task separately before
// you deregister the account. Either, you can do this in the Audit Manager
// console. Or, you can use one of the delete API operations that are provided by
// Audit Manager.
//
// To delete your Audit Manager resource data, see the following instructions:
//
// [DeleteAssessment]
//   - (see also: [Deleting an assessment]in the Audit Manager User Guide)
//
// [DeleteAssessmentFramework]
//   - (see also: [Deleting a custom framework]in the Audit Manager User Guide)
//
// [DeleteAssessmentFrameworkShare]
//   - (see also: [Deleting a share request]in the Audit Manager User Guide)
//
// [DeleteAssessmentReport]
//   - (see also: [Deleting an assessment report]in the Audit Manager User Guide)
//
// [DeleteControl]
//   - (see also: [Deleting a custom control]in the Audit Manager User Guide)
//
// At this time, Audit Manager doesn't provide an option to delete evidence for a
// specific delegated administrator. Instead, when your management account
// deregisters Audit Manager, we perform a cleanup for the current delegated
// administrator account at the time of deregistration.
//
// [DeleteControl]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteControl.html
// [deletes the event data store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-eds-disable-termination.html
// [DeleteAssessmentFrameworkShare]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFrameworkShare.html
// [Deleting a custom control]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-controls.html
// [Deleting an assessment report]: https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#delete-assessment-report-steps
// [Deleting a custom framework]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-custom-framework.html
// [Deleting a share request]: https://docs.aws.amazon.com/audit-manager/latest/userguide/deleting-shared-framework-requests.html
// [Deleting an assessment]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-assessment.html
// [DeleteAssessment]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessment.html
// [DeleteAssessmentReport]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentReport.html
// [DeleteAssessmentFramework]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFramework.html
func (c *Client) DeregisterOrganizationAdminAccount(ctx context.Context, params *DeregisterOrganizationAdminAccountInput, optFns ...func(*Options)) (*DeregisterOrganizationAdminAccountOutput, error) {
	if params == nil {
		params = &DeregisterOrganizationAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterOrganizationAdminAccount", params, optFns, c.addOperationDeregisterOrganizationAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterOrganizationAdminAccountInput struct {

	//  The identifier for the administrator account.
	AdminAccountId *string

	noSmithyDocumentSerde
}

type DeregisterOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterOrganizationAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeregisterOrganizationAdminAccount"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeregisterOrganizationAdminAccount",
	}
}
