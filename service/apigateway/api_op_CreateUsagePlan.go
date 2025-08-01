// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a usage plan with the throttle and quota limits, as well as the
// associated API stages, specified in the payload.
func (c *Client) CreateUsagePlan(ctx context.Context, params *CreateUsagePlanInput, optFns ...func(*Options)) (*CreateUsagePlanOutput, error) {
	if params == nil {
		params = &CreateUsagePlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsagePlan", params, optFns, c.addOperationCreateUsagePlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsagePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST request to create a usage plan with the name, description, throttle
// limits and quota limits, as well as the associated API stages, specified in the
// payload.
type CreateUsagePlanInput struct {

	// The name of the usage plan.
	//
	// This member is required.
	Name *string

	// The associated API stages of the usage plan.
	ApiStages []types.ApiStage

	// The description of the usage plan.
	Description *string

	// The quota of the usage plan.
	Quota *types.QuotaSettings

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws: . The tag value
	// can be up to 256 characters.
	Tags map[string]string

	// The throttling limits of the usage plan.
	Throttle *types.ThrottleSettings

	noSmithyDocumentSerde
}

// Represents a usage plan used to specify who can assess associated API stages.
// Optionally, target request rate and quota limits can be set. In some cases
// clients can exceed the targets that you set. Don’t rely on usage plans to
// control costs. Consider using [Amazon Web Services Budgets]to monitor costs and [WAF] to manage API requests.
//
// [Amazon Web Services Budgets]: https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html
// [WAF]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
type CreateUsagePlanOutput struct {

	// The associated API stages of a usage plan.
	ApiStages []types.ApiStage

	// The description of a usage plan.
	Description *string

	// The identifier of a UsagePlan resource.
	Id *string

	// The name of a usage plan.
	Name *string

	// The Amazon Web Services Marketplace product identifier to associate with the
	// usage plan as a SaaS product on the Amazon Web Services Marketplace.
	ProductCode *string

	// The target maximum number of permitted requests per a given unit time interval.
	Quota *types.QuotaSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// A map containing method level throttling information for API stage in a usage
	// plan.
	Throttle *types.ThrottleSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUsagePlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUsagePlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUsagePlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUsagePlan"); err != nil {
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
	if err = addOpCreateUsagePlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsagePlan(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateUsagePlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUsagePlan",
	}
}
