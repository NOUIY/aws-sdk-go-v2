// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a budget and, if included, notifications and subscribers.
//
// Only one of BudgetLimit or PlannedBudgetLimits can be present in the syntax at
// one time. Use the syntax that matches your use case. The Request Syntax section
// shows the BudgetLimit syntax. For PlannedBudgetLimits , see the [Examples] section.
//
// Similarly, only one set of filter and metric selections can be present in the
// syntax at one time. Either FilterExpression and Metrics or CostFilters and
// CostTypes , not both or a different combination. We recommend using
// FilterExpression and Metrics as they provide more flexible and powerful
// filtering capabilities. The Request Syntax section shows the FilterExpression /
// Metrics syntax.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CreateBudget.html#API_CreateBudget_Examples
func (c *Client) CreateBudget(ctx context.Context, params *CreateBudgetInput, optFns ...func(*Options)) (*CreateBudgetOutput, error) {
	if params == nil {
		params = &CreateBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBudget", params, optFns, c.addOperationCreateBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of CreateBudget
type CreateBudgetInput struct {

	// The accountId that is associated with the budget.
	//
	// This member is required.
	AccountId *string

	// The budget object that you want to create.
	//
	// This member is required.
	Budget *types.Budget

	// A notification that you want to associate with a budget. A budget can have up
	// to five notifications, and each notification can have one SNS subscriber and up
	// to 10 email subscribers. If you include notifications and subscribers in your
	// CreateBudget call, Amazon Web Services creates the notifications and subscribers
	// for you.
	NotificationsWithSubscribers []types.NotificationWithSubscribers

	// An optional list of tags to associate with the specified budget. Each tag
	// consists of a key and a value, and each key must be unique for the resource.
	ResourceTags []types.ResourceTag

	noSmithyDocumentSerde
}

// Response of CreateBudget
type CreateBudgetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBudget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBudget"); err != nil {
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
	if err = addOpCreateBudgetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBudget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBudget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBudget",
	}
}
