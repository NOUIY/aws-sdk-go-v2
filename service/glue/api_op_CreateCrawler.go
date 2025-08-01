// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new crawler with specified targets, role, configuration, and optional
// schedule. At least one crawl target must be specified, in the s3Targets field,
// the jdbcTargets field, or the DynamoDBTargets field.
func (c *Client) CreateCrawler(ctx context.Context, params *CreateCrawlerInput, optFns ...func(*Options)) (*CreateCrawlerOutput, error) {
	if params == nil {
		params = &CreateCrawlerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCrawler", params, optFns, c.addOperationCreateCrawlerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCrawlerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCrawlerInput struct {

	// Name of the new crawler.
	//
	// This member is required.
	Name *string

	// The IAM role or Amazon Resource Name (ARN) of an IAM role used by the new
	// crawler to access customer resources.
	//
	// This member is required.
	Role *string

	// A list of collection of targets to crawl.
	//
	// This member is required.
	Targets *types.CrawlerTargets

	// A list of custom classifiers that the user has registered. By default, all
	// built-in classifiers are included in a crawl, but these custom classifiers
	// always override the default classifiers for a given classification.
	Classifiers []string

	// Crawler configuration information. This versioned JSON string allows users to
	// specify aspects of a crawler's behavior. For more information, see [Setting crawler configuration options].
	//
	// [Setting crawler configuration options]: https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html
	Configuration *string

	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string

	// The Glue database where results are written, such as:
	// arn:aws:daylight:us-east-1::database/sometable/* .
	DatabaseName *string

	// A description of the new crawler.
	Description *string

	// Specifies Lake Formation configuration settings for the crawler.
	LakeFormationConfiguration *types.LakeFormationConfiguration

	// Specifies data lineage configuration settings for the crawler.
	LineageConfiguration *types.LineageConfiguration

	// A policy that specifies whether to crawl the entire dataset again, or to crawl
	// only folders that were added since the last crawler run.
	RecrawlPolicy *types.RecrawlPolicy

	// A cron expression used to specify the schedule (see [Time-Based Schedules for Jobs and Crawlers]. For example, to run
	// something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *) .
	//
	// [Time-Based Schedules for Jobs and Crawlers]: https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html
	Schedule *string

	// The policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *types.SchemaChangePolicy

	// The table prefix used for catalog tables that are created.
	TablePrefix *string

	// The tags to use with this crawler request. You may use tags to limit access to
	// the crawler. For more information about tags in Glue, see [Amazon Web Services Tags in Glue]in the developer
	// guide.
	//
	// [Amazon Web Services Tags in Glue]: https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCrawlerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCrawlerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCrawler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCrawler{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCrawler"); err != nil {
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
	if err = addOpCreateCrawlerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCrawler(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCrawler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCrawler",
	}
}
