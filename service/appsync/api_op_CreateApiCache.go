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

// Creates a cache for the GraphQL API.
func (c *Client) CreateApiCache(ctx context.Context, params *CreateApiCacheInput, optFns ...func(*Options)) (*CreateApiCacheOutput, error) {
	if params == nil {
		params = &CreateApiCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiCache", params, optFns, c.addOperationCreateApiCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateApiCache operation.
type CreateApiCacheInput struct {

	// Caching behavior.
	//
	//   - FULL_REQUEST_CACHING: All requests from the same user are cached.
	//   Individual resolvers are automatically cached. All API calls will try to return
	//   responses from the cache.
	//
	//   - PER_RESOLVER_CACHING: Individual resolvers that you specify are cached.
	//
	//   - OPERATION_LEVEL_CACHING: Full requests are cached together and returned
	//   without executing resolvers.
	//
	// This member is required.
	ApiCachingBehavior types.ApiCachingBehavior

	// The GraphQL API ID.
	//
	// This member is required.
	ApiId *string

	// TTL in seconds for cache entries.
	//
	// Valid values are 1–3,600 seconds.
	//
	// This member is required.
	Ttl int64

	// The cache instance type. Valid values are
	//
	//   - SMALL
	//
	//   - MEDIUM
	//
	//   - LARGE
	//
	//   - XLARGE
	//
	//   - LARGE_2X
	//
	//   - LARGE_4X
	//
	//   - LARGE_8X (not available in all regions)
	//
	//   - LARGE_12X
	//
	// Historically, instance types were identified by an EC2-style value. As of July
	// 2020, this is deprecated, and the generic identifiers above should be used.
	//
	// The following legacy instance types are available, but their use is discouraged:
	//
	//   - T2_SMALL: A t2.small instance type.
	//
	//   - T2_MEDIUM: A t2.medium instance type.
	//
	//   - R4_LARGE: A r4.large instance type.
	//
	//   - R4_XLARGE: A r4.xlarge instance type.
	//
	//   - R4_2XLARGE: A r4.2xlarge instance type.
	//
	//   - R4_4XLARGE: A r4.4xlarge instance type.
	//
	//   - R4_8XLARGE: A r4.8xlarge instance type.
	//
	// This member is required.
	Type types.ApiCacheType

	// At-rest encryption flag for cache. You cannot update this setting after
	// creation.
	//
	// Deprecated: atRestEncryptionEnabled attribute is deprecated. Encryption at rest
	// is always enabled.
	AtRestEncryptionEnabled bool

	// Controls how cache health metrics will be emitted to CloudWatch. Cache health
	// metrics include:
	//
	//   - NetworkBandwidthOutAllowanceExceeded: The network packets dropped because
	//   the throughput exceeded the aggregated bandwidth limit. This is useful for
	//   diagnosing bottlenecks in a cache configuration.
	//
	//   - EngineCPUUtilization: The CPU utilization (percentage) allocated to the
	//   Redis process. This is useful for diagnosing bottlenecks in a cache
	//   configuration.
	//
	// Metrics will be recorded by API ID. You can set the value to ENABLED or DISABLED
	// .
	HealthMetricsConfig types.CacheHealthMetricsConfig

	// Transit encryption flag when connecting to cache. You cannot update this
	// setting after creation.
	//
	// Deprecated: transitEncryptionEnabled attribute is deprecated. Encryption in
	// transit is always enabled.
	TransitEncryptionEnabled bool

	noSmithyDocumentSerde
}

// Represents the output of a CreateApiCache operation.
type CreateApiCacheOutput struct {

	// The ApiCache object.
	ApiCache *types.ApiCache

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApiCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiCache{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApiCache"); err != nil {
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
	if err = addOpCreateApiCacheValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiCache(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApiCache(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApiCache",
	}
}
