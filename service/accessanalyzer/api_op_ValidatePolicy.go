// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests the validation of a policy and returns a list of findings. The
// findings help you identify issues and provide actionable recommendations to
// resolve the issue and enable you to author functional policies that meet
// security best practices.
func (c *Client) ValidatePolicy(ctx context.Context, params *ValidatePolicyInput, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
	if params == nil {
		params = &ValidatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidatePolicy", params, optFns, c.addOperationValidatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidatePolicyInput struct {

	// The JSON policy document to use as the content for the policy.
	//
	// This member is required.
	PolicyDocument *string

	// The type of policy to validate. Identity policies grant permissions to IAM
	// principals. Identity policies include managed and inline policies for IAM roles,
	// users, and groups.
	//
	// Resource policies grant permissions on Amazon Web Services resources. Resource
	// policies include trust policies for IAM roles and bucket policies for Amazon S3
	// buckets. You can provide a generic input such as identity policy or resource
	// policy or a specific input such as managed policy or Amazon S3 bucket policy.
	//
	// Service control policies (SCPs) are a type of organization policy attached to
	// an Amazon Web Services organization, organizational unit (OU), or an account.
	//
	// This member is required.
	PolicyType types.PolicyType

	// The locale to use for localizing the findings.
	Locale types.Locale

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	// The type of resource to attach to your resource policy. Specify a value for the
	// policy validation resource type only if the policy type is RESOURCE_POLICY . For
	// example, to validate a resource policy to attach to an Amazon S3 bucket, you can
	// choose AWS::S3::Bucket for the policy validation resource type.
	//
	// For resource types not supported as valid values, IAM Access Analyzer runs
	// policy checks that apply to all resource policies. For example, to validate a
	// resource policy to attach to a KMS key, do not specify a value for the policy
	// validation resource type and IAM Access Analyzer will run policy checks that
	// apply to all resource policies.
	ValidatePolicyResourceType types.ValidatePolicyResourceType

	noSmithyDocumentSerde
}

type ValidatePolicyOutput struct {

	// The list of findings in a policy returned by IAM Access Analyzer based on its
	// suite of policy checks.
	//
	// This member is required.
	Findings []types.ValidatePolicyFinding

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ValidatePolicy"); err != nil {
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
	if err = addOpValidatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePolicy(options.Region), middleware.Before); err != nil {
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

// ValidatePolicyPaginatorOptions is the paginator options for ValidatePolicy
type ValidatePolicyPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ValidatePolicyPaginator is a paginator for ValidatePolicy
type ValidatePolicyPaginator struct {
	options   ValidatePolicyPaginatorOptions
	client    ValidatePolicyAPIClient
	params    *ValidatePolicyInput
	nextToken *string
	firstPage bool
}

// NewValidatePolicyPaginator returns a new ValidatePolicyPaginator
func NewValidatePolicyPaginator(client ValidatePolicyAPIClient, params *ValidatePolicyInput, optFns ...func(*ValidatePolicyPaginatorOptions)) *ValidatePolicyPaginator {
	if params == nil {
		params = &ValidatePolicyInput{}
	}

	options := ValidatePolicyPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ValidatePolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ValidatePolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ValidatePolicy page.
func (p *ValidatePolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ValidatePolicy(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ValidatePolicyAPIClient is a client that implements the ValidatePolicy
// operation.
type ValidatePolicyAPIClient interface {
	ValidatePolicy(context.Context, *ValidatePolicyInput, ...func(*Options)) (*ValidatePolicyOutput, error)
}

var _ ValidatePolicyAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opValidatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ValidatePolicy",
	}
}
