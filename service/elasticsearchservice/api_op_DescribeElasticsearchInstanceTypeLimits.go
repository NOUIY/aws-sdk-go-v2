// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Describe Elasticsearch Limits for a given InstanceType and
//
// ElasticsearchVersion. When modifying existing Domain, specify the DomainNameto know what
// Limits are supported for modifying.
func (c *Client) DescribeElasticsearchInstanceTypeLimits(ctx context.Context, params *DescribeElasticsearchInstanceTypeLimitsInput, optFns ...func(*Options)) (*DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	if params == nil {
		params = &DescribeElasticsearchInstanceTypeLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticsearchInstanceTypeLimits", params, optFns, c.addOperationDescribeElasticsearchInstanceTypeLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticsearchInstanceTypeLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to DescribeElasticsearchInstanceTypeLimits operation.
type DescribeElasticsearchInstanceTypeLimitsInput struct {

	//  Version of Elasticsearch for which Limits are needed.
	//
	// This member is required.
	ElasticsearchVersion *string

	//  The instance type for an Elasticsearch cluster for which Elasticsearch Limits are
	// needed.
	//
	// This member is required.
	InstanceType types.ESPartitionInstanceType

	//  DomainName represents the name of the Domain that we are trying to modify.
	// This should be present only if we are querying for Elasticsearch Limitsfor existing
	// domain.
	DomainName *string

	noSmithyDocumentSerde
}

// Container for the parameters received from DescribeElasticsearchInstanceTypeLimits operation.
type DescribeElasticsearchInstanceTypeLimitsOutput struct {

	//  Map of Role of the Instance and Limits that are applicable. Role performed by
	// given Instance in Elasticsearch can be one of the following:
	//
	//   - data: If the given InstanceType is used as data node
	//   - master: If the given InstanceType is used as master node
	//   - ultra_warm: If the given InstanceType is used as warm node
	LimitsByRole map[string]types.Limits

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeElasticsearchInstanceTypeLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeElasticsearchInstanceTypeLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeElasticsearchInstanceTypeLimits{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeElasticsearchInstanceTypeLimits"); err != nil {
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
	if err = addOpDescribeElasticsearchInstanceTypeLimitsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticsearchInstanceTypeLimits(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticsearchInstanceTypeLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeElasticsearchInstanceTypeLimits",
	}
}
