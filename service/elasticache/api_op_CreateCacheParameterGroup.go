// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the
// CacheParameterGroup.
//
// A newly created CacheParameterGroup is an exact duplicate of the default
// parameter group for the CacheParameterGroupFamily. To customize the newly
// created CacheParameterGroup you can change the values of specific parameters.
// For more information, see:
//
// [ModifyCacheParameterGroup]
//   - in the ElastiCache API Reference.
//
// [Parameters and Parameter Groups]
//   - in the ElastiCache User Guide.
//
// [ModifyCacheParameterGroup]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html
// [Parameters and Parameter Groups]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.html
func (c *Client) CreateCacheParameterGroup(ctx context.Context, params *CreateCacheParameterGroupInput, optFns ...func(*Options)) (*CreateCacheParameterGroupOutput, error) {
	if params == nil {
		params = &CreateCacheParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCacheParameterGroup", params, optFns, c.addOperationCreateCacheParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCacheParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheParameterGroup operation.
type CreateCacheParameterGroupInput struct {

	// The name of the cache parameter group family that the cache parameter group can
	// be used with.
	//
	// Valid values are: valkey8 | valkey7 | memcached1.4 | memcached1.5 | memcached1.6
	// | redis2.6 | redis2.8 | redis3.2 | redis4.0 | redis5.0 | redis6.x | redis7
	//
	// This member is required.
	CacheParameterGroupFamily *string

	// A user-specified name for the cache parameter group.
	//
	// This member is required.
	CacheParameterGroupName *string

	// A user-specified description for the cache parameter group.
	//
	// This member is required.
	Description *string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCacheParameterGroupOutput struct {

	// Represents the output of a CreateCacheParameterGroup operation.
	CacheParameterGroup *types.CacheParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCacheParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCacheParameterGroup"); err != nil {
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
	if err = addOpCreateCacheParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCacheParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCacheParameterGroup",
	}
}
