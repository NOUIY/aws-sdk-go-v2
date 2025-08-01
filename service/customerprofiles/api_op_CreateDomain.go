// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a domain, which is a container for all customer data, such as customer
// profile attributes, object types, profile keys, and encryption keys. You can
// create multiple domains, and each domain can have multiple third-party
// integrations.
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// Use this API or [UpdateDomain] to enable [identity resolution]: set Matching to true.
//
// To prevent cross-service impersonation when you call this API, see [Cross-service confused deputy prevention] for sample
// policies that you should apply.
//
// It is not possible to associate a Customer Profiles domain with an Amazon
// Connect Instance directly from the API. If you would like to create a domain and
// associate a Customer Profiles domain, use the Amazon Connect admin website. For
// more information, see [Enable Customer Profiles].
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// [UpdateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html
// [Enable Customer Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-customer-profiles.html#enable-customer-profiles-step1
// [Cross-service confused deputy prevention]: https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html
// [identity resolution]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, c.addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The default number of days until the data within the domain expires.
	//
	// This member is required.
	DefaultExpirationDays *int32

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications. You must set up a
	// policy on the DeadLetterQueue for the SendMessage operation to enable Amazon
	// Connect Customer Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The process of matching duplicate profiles. If Matching = true , Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains.
	//
	// After the Identity Resolution Job completes, use the [GetMatches] API to return and review
	// the results. Or, if you have configured ExportingConfig in the MatchingRequest ,
	// you can download the results from S3.
	//
	// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
	Matching *types.MatchingRequest

	// The process of matching duplicate profiles using the Rule-Based matching. If
	// RuleBasedMatching = true, Amazon Connect Customer Profiles will start to match
	// and merge your profiles according to your configuration in the
	// RuleBasedMatchingRequest . You can use the ListRuleBasedMatches and
	// GetSimilarProfiles API to return and review the results. Also, if you have
	// configured ExportingConfig in the RuleBasedMatchingRequest , you can download
	// the results from S3.
	RuleBasedMatching *types.RuleBasedMatchingRequest

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The default number of days until the data within the domain expires.
	//
	// This member is required.
	DefaultExpirationDays *int32

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The process of matching duplicate profiles. If Matching = true , Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains.
	//
	// After the Identity Resolution Job completes, use the [GetMatches] API to return and review
	// the results. Or, if you have configured ExportingConfig in the MatchingRequest ,
	// you can download the results from S3.
	//
	// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
	Matching *types.MatchingResponse

	// The process of matching duplicate profiles using the Rule-Based matching. If
	// RuleBasedMatching = true, Amazon Connect Customer Profiles will start to match
	// and merge your profiles according to your configuration in the
	// RuleBasedMatchingRequest . You can use the ListRuleBasedMatches and
	// GetSimilarProfiles API to return and review the results. Also, if you have
	// configured ExportingConfig in the RuleBasedMatchingRequest , you can download
	// the results from S3.
	RuleBasedMatching *types.RuleBasedMatchingResponse

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDomain"); err != nil {
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
	if err = addOpCreateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDomain",
	}
}
