// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon OpenSearch Service domain. For more information, see [Creating and managing Amazon OpenSearch Service domains].
//
// [Creating and managing Amazon OpenSearch Service domains]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html
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

	// Name of the OpenSearch Service domain to create. Domain names are unique across
	// the domains owned by an account within an Amazon Web Services Region.
	//
	// This member is required.
	DomainName *string

	// Options for all machine learning features for the specified domain.
	AIMLOptions *types.AIMLOptionsInput

	// Identity and Access Management (IAM) policy document specifying the access
	// policies for the new domain.
	AccessPolicies *string

	// Key-value pairs to specify advanced configuration options. The following
	// key-value pairs are supported:
	//
	//   - "rest.action.multi.allow_explicit_index": "true" | "false" - Note the use of
	//   a string rather than a boolean. Specifies whether explicit references to indexes
	//   are allowed inside the body of HTTP requests. If you want to configure access
	//   policies for domain sub-resources, such as specific indexes and domain APIs, you
	//   must disable this property. Default is true.
	//
	//   - "indices.fielddata.cache.size": "80" - Note the use of a string rather than
	//   a boolean. Specifies the percentage of heap space allocated to field data.
	//   Default is unbounded.
	//
	//   - "indices.query.bool.max_clause_count": "1024" - Note the use of a string
	//   rather than a boolean. Specifies the maximum number of clauses allowed in a
	//   Lucene boolean query. Default is 1,024. Queries with more than the permitted
	//   number of clauses result in a TooManyClauses error.
	//
	//   - "override_main_response_version": "true" | "false" - Note the use of a
	//   string rather than a boolean. Specifies whether the domain reports its version
	//   as 7.10 to allow Elasticsearch OSS clients and plugins to continue working with
	//   it. Default is false when creating a domain and true when upgrading a domain.
	//
	// For more information, see [Advanced cluster parameters].
	//
	// [Advanced cluster parameters]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options
	AdvancedOptions map[string]string

	// Options for fine-grained access control.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Options for Auto-Tune.
	AutoTuneOptions *types.AutoTuneOptionsInput

	// Container for the cluster configuration of a domain.
	ClusterConfig *types.ClusterConfig

	// Key-value pairs to configure Amazon Cognito authentication. For more
	// information, see [Configuring Amazon Cognito authentication for OpenSearch Dashboards].
	//
	// [Configuring Amazon Cognito authentication for OpenSearch Dashboards]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html
	CognitoOptions *types.CognitoOptions

	// Additional options for the domain endpoint, such as whether to require HTTPS
	// for all traffic.
	DomainEndpointOptions *types.DomainEndpointOptions

	// Container for the parameters required to enable EBS-based storage for an
	// OpenSearch Service domain.
	EBSOptions *types.EBSOptions

	// Key-value pairs to enable encryption at rest.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engine
	// version for the OpenSearch Service domain. For example, OpenSearch_1.0 or
	// Elasticsearch_7.9 . For more information, see [Creating and managing Amazon OpenSearch Service domains].
	//
	// [Creating and managing Amazon OpenSearch Service domains]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains
	EngineVersion *string

	// Specify either dual stack or IPv4 as your IP address type. Dual stack allows
	// you to share domain resources across IPv4 and IPv6 address types, and is the
	// recommended option. If you set your IP address type to dual stack, you can't
	// change your address type later.
	IPAddressType types.IPAddressType

	// Configuration options for enabling and managing IAM Identity Center integration
	// within a domain.
	IdentityCenterOptions *types.IdentityCenterOptionsInput

	// Key-value pairs to configure log publishing.
	LogPublishingOptions map[string]types.LogPublishingOption

	// Enables node-to-node encryption.
	NodeToNodeEncryptionOptions *types.NodeToNodeEncryptionOptions

	// Specifies a daily 10-hour time block during which OpenSearch Service can
	// perform configuration changes on the domain, including service software updates
	// and Auto-Tune enhancements that require a blue/green deployment. If no options
	// are specified, the default start time of 10:00 P.M. local time (for the Region
	// that the domain is created in) is used.
	OffPeakWindowOptions *types.OffPeakWindowOptions

	// DEPRECATED. Container for the parameters required to configure automated
	// snapshots of domain indexes.
	SnapshotOptions *types.SnapshotOptions

	// Software update options for the domain.
	SoftwareUpdateOptions *types.SoftwareUpdateOptions

	// List of tags to add to the domain upon creation.
	TagList []types.Tag

	// Container for the values required to configure VPC access domains. If you don't
	// specify these values, OpenSearch Service creates the domain with a public
	// endpoint. For more information, see [Launching your Amazon OpenSearch Service domains using a VPC].
	//
	// [Launching your Amazon OpenSearch Service domains using a VPC]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html
	VPCOptions *types.VPCOptions

	noSmithyDocumentSerde
}

// The result of a CreateDomain operation. Contains the status of the newly
// created domain.
type CreateDomainOutput struct {

	// The status of the newly created domain.
	DomainStatus *types.DomainStatus

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
