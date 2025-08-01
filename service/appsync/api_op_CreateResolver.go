// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Resolver object.
//
// A resolver converts incoming requests into a format that a data source can
// understand, and converts the data source's responses into GraphQL.
func (c *Client) CreateResolver(ctx context.Context, params *CreateResolverInput, optFns ...func(*Options)) (*CreateResolverOutput, error) {
	if params == nil {
		params = &CreateResolverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolver", params, optFns, c.addOperationCreateResolverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverInput struct {

	// The ID for the GraphQL API for which the resolver is being created.
	//
	// This member is required.
	ApiId *string

	// The name of the field to attach the resolver to.
	//
	// This member is required.
	FieldName *string

	// The name of the Type .
	//
	// This member is required.
	TypeName *string

	// The caching configuration for the resolver.
	CachingConfig *types.CachingConfig

	// The resolver code that contains the request and response functions. When code
	// is used, the runtime is required. The runtime value must be APPSYNC_JS .
	Code *string

	// The name of the data source for which the resolver is being created.
	DataSourceName *string

	// The resolver type.
	//
	//   - UNIT: A UNIT resolver type. A UNIT resolver is the default resolver type.
	//   You can use a UNIT resolver to run a GraphQL query against a single data source.
	//
	//   - PIPELINE: A PIPELINE resolver type. You can use a PIPELINE resolver to
	//   invoke a series of Function objects in a serial manner. You can use a pipeline
	//   resolver to run a GraphQL query against multiple data sources.
	Kind types.ResolverKind

	// The maximum batching size for a resolver.
	MaxBatchSize int32

	// Enables or disables enhanced resolver metrics for specified resolvers. Note
	// that metricsConfig won't be used unless the resolverLevelMetricsBehavior value
	// is set to PER_RESOLVER_METRICS . If the resolverLevelMetricsBehavior is set to
	// FULL_REQUEST_RESOLVER_METRICS instead, metricsConfig will be ignored. However,
	// you can still set its value.
	//
	// metricsConfig can be ENABLED or DISABLED .
	MetricsConfig types.ResolverLevelMetricsConfig

	// The PipelineConfig .
	PipelineConfig *types.PipelineConfig

	// The mapping template to use for requests.
	//
	// A resolver uses a request mapping template to convert a GraphQL expression into
	// a format that a data source can understand. Mapping templates are written in
	// Apache Velocity Template Language (VTL).
	//
	// VTL request mapping templates are optional when using an Lambda data source.
	// For all other data sources, VTL request and response mapping templates are
	// required.
	RequestMappingTemplate *string

	// The mapping template to use for responses from the data source.
	ResponseMappingTemplate *string

	// Describes a runtime used by an Amazon Web Services AppSync pipeline resolver or
	// Amazon Web Services AppSync function. Specifies the name and version of the
	// runtime to use. Note that if a runtime is specified, code must also be
	// specified.
	Runtime *types.AppSyncRuntime

	// The SyncConfig for a resolver attached to a versioned data source.
	SyncConfig *types.SyncConfig

	noSmithyDocumentSerde
}

type CreateResolverOutput struct {

	// The Resolver object.
	Resolver *types.Resolver

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResolverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResolver{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResolver{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateResolver"); err != nil {
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
	if err = addOpCreateResolverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolver(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResolver(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateResolver",
	}
}
