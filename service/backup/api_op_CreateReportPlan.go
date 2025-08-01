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

// Creates a report plan. A report plan is a document that contains information
// about the contents of the report and where Backup will deliver it.
//
// If you call CreateReportPlan with a plan that already exists, you receive an
// AlreadyExistsException exception.
func (c *Client) CreateReportPlan(ctx context.Context, params *CreateReportPlanInput, optFns ...func(*Options)) (*CreateReportPlanOutput, error) {
	if params == nil {
		params = &CreateReportPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReportPlan", params, optFns, c.addOperationCreateReportPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReportPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReportPlanInput struct {

	// A structure that contains information about where and how to deliver your
	// reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats
	// of your reports.
	//
	// This member is required.
	ReportDeliveryChannel *types.ReportDeliveryChannel

	// The unique name of the report plan. The name must be between 1 and 256
	// characters, starting with a letter, and consisting of letters (a-z, A-Z),
	// numbers (0-9), and underscores (_).
	//
	// This member is required.
	ReportPlanName *string

	// Identifies the report template for the report. Reports are built using a report
	// template. The report templates are:
	//
	//     RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT |
	//     COPY_JOB_REPORT | RESTORE_JOB_REPORT
	//
	// If the report template is RESOURCE_COMPLIANCE_REPORT or
	// CONTROL_COMPLIANCE_REPORT , this API resource also describes the report coverage
	// by Amazon Web Services Regions and frameworks.
	//
	// This member is required.
	ReportSetting *types.ReportSetting

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to CreateReportPlanInput . Retrying a successful request with
	// the same idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	// An optional description of the report plan with a maximum of 1,024 characters.
	ReportPlanDescription *string

	// The tags to assign to the report plan.
	ReportPlanTags map[string]string

	noSmithyDocumentSerde
}

type CreateReportPlanOutput struct {

	// The date and time a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationTime is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationTime *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	ReportPlanArn *string

	// The unique name of the report plan.
	ReportPlanName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReportPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReportPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReportPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReportPlan"); err != nil {
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
	if err = addIdempotencyToken_opCreateReportPlanMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateReportPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReportPlan(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateReportPlan struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateReportPlan) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateReportPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateReportPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateReportPlanInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateReportPlanMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateReportPlan{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateReportPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReportPlan",
	}
}
