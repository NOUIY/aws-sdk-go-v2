// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an account-level data protection policy, subscription filter policy,
// field index policy, transformer policy, or metric extraction policy that applies
// to all log groups or a subset of log groups in the account.
//
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are creating.
//
//   - To create a data protection policy, you must have the
//     logs:PutDataProtectionPolicy and logs:PutAccountPolicy permissions.
//
//   - To create a subscription filter policy, you must have the
//     logs:PutSubscriptionFilter and logs:PutAccountPolicy permissions.
//
//   - To create a transformer policy, you must have the logs:PutTransformer and
//     logs:PutAccountPolicy permissions.
//
//   - To create a field index policy, you must have the logs:PutIndexPolicy and
//     logs:PutAccountPolicy permissions.
//
//   - To create a metric extraction policy, you must have the
//     logs:PutMetricExtractionPolicy and logs:PutAccountPolicy permissions.
//
// # Data protection policy
//
// A data protection policy can help safeguard sensitive data that's ingested by
// your log groups by auditing and masking the sensitive log data. Each account can
// have only one account-level data protection policy.
//
// Sensitive data is detected and masked when it is ingested into a log group.
// When you set a data protection policy, log events ingested into the log groups
// before that time are not masked.
//
// If you use PutAccountPolicy to create a data protection policy for your whole
// account, it applies to both existing log groups and all log groups that are
// created later in this account. The account-level policy is applied to existing
// log groups with eventual consistency. It might take up to 5 minutes before
// sensitive data in existing log groups begins to be masked.
//
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a [GetLogEvents]or [FilterLogEvents] operation with the unmask parameter set to true to
// view the unmasked log events. Users with the logs:Unmask can also view unmasked
// data in the CloudWatch Logs console by running a CloudWatch Logs Insights query
// with the unmask query command.
//
// For more information, including a list of types of data that can be audited and
// masked, see [Protect sensitive log data with masking].
//
// To use the PutAccountPolicy operation for a data protection policy, you must be
// signed on with the logs:PutDataProtectionPolicy and logs:PutAccountPolicy
// permissions.
//
// The PutAccountPolicy operation applies to all log groups in the account. You
// can use [PutDataProtectionPolicy]to create a data protection policy that applies to just one log group.
// If a log group has its own data protection policy and the account also has an
// account-level data protection policy, then the two policies are cumulative. Any
// sensitive term specified in either policy is masked.
//
// # Subscription filter policy
//
// A subscription filter policy sets up a real-time feed of log events from
// CloudWatch Logs to other Amazon Web Services services. Account-level
// subscription filter policies apply to both existing log groups and log groups
// that are created later in this account. Supported destinations are Kinesis Data
// Streams, Firehose, and Lambda. When log events are sent to the receiving
// service, they are Base64 encoded and compressed with the GZIP format.
//
// The following destinations are supported for subscription filters:
//
//   - An Kinesis Data Streams data stream in the same account as the subscription
//     policy, for same-account delivery.
//
//   - An Firehose data stream in the same account as the subscription policy, for
//     same-account delivery.
//
//   - A Lambda function in the same account as the subscription policy, for
//     same-account delivery.
//
//   - A logical destination in a different account created with [PutDestination], for
//     cross-account delivery. Kinesis Data Streams and Firehose are supported as
//     logical destinations.
//
// Each account can have one account-level subscription filter policy per Region.
// If you are updating an existing filter, you must specify the correct name in
// PolicyName . To perform a PutAccountPolicy subscription filter operation for
// any destination except a Lambda function, you must also have the iam:PassRole
// permission.
//
// # Transformer policy
//
// Creates or updates a log transformer policy for your account. You use log
// transformers to transform log events into a different format, making them easier
// for you to process and analyze. You can also transform logs from different
// sources into standardized formats that contain relevant, source-specific
// information. After you have created a transformer, CloudWatch Logs performs this
// transformation at the time of log ingestion. You can then refer to the
// transformed versions of the logs during operations such as querying with
// CloudWatch Logs Insights or creating metric filters or subscription filters.
//
// You can also use a transformer to copy metadata from metadata keys into the log
// events themselves. This metadata can include log group name, log stream name,
// account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor
// applies one type of transformation to the log events ingested into this log
// group. For more information about the available processors to use in a
// transformer, see [Processors that you can use].
//
// Having log events in standardized format enables visibility across your
// applications for your log analysis, reporting, and alarming needs. CloudWatch
// Logs provides transformation for common log types with out-of-the-box
// transformation templates for major Amazon Web Services log sources such as VPC
// flow logs, Lambda, and Amazon RDS. You can use pre-built transformation
// templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can have one account-level transformer policy that applies to all log
// groups in the account. Or you can create as many as 20 account-level transformer
// policies that are each scoped to a subset of log groups with the
// selectionCriteria parameter. If you have multiple account-level transformer
// policies with selection criteria, no two of them can use the same or overlapping
// log group name prefixes. For example, if you have one policy filtered to log
// groups that start with my-log , you can't have another field index policy
// filtered to my-logpprod or my-logging .
//
// You can also set up a transformer at the log-group level. For more information,
// see [PutTransformer]. If there is both a log-group level transformer created with PutTransformer
// and an account-level transformer that could apply to the same log group, the log
// group uses only the log-group level transformer. It ignores the account-level
// transformer.
//
// # Field index policy
//
// You can use field index policies to create indexes on fields found in log
// events in the log group. Creating field indexes can help lower the scan volume
// for CloudWatch Logs Insights queries that reference those fields, because these
// queries attempt to skip the processing of log events that are known to not match
// the indexed field. Good fields to index are fields that you often need to query
// for and fields or values that match only a small fraction of the total log
// events. Common examples of indexes include request ID, session ID, user IDs, or
// instance IDs. For more information, see [Create field indexes to improve query performance and reduce costs]
//
// To find the fields that are in your log group events, use the [GetLogGroupFields] operation.
//
// For example, suppose you have created a field index for requestId . Then, any
// CloudWatch Logs Insights query on that log group that includes requestId =
// value or requestId in [value, value, ...] will attempt to process only the log
// events where the indexed field matches the specified value.
//
// Matches of log events to the names of indexed fields are case-sensitive. For
// example, an indexed field of RequestId won't match a log event containing
// requestId .
//
// You can have one account-level field index policy that applies to all log
// groups in the account. Or you can create as many as 20 account-level field index
// policies that are each scoped to a subset of log groups with the
// selectionCriteria parameter. If you have multiple account-level index policies
// with selection criteria, no two of them can use the same or overlapping log
// group name prefixes. For example, if you have one policy filtered to log groups
// that start with my-log , you can't have another field index policy filtered to
// my-logpprod or my-logging .
//
// If you create an account-level field index policy in a monitoring account in
// cross-account observability, the policy is applied only to the monitoring
// account and not to any source accounts.
//
// If you want to create a field index policy for a single log group, you can use [PutIndexPolicy]
// instead of PutAccountPolicy . If you do so, that log group will use only that
// log-group level policy, and will ignore the account-level policy that you create
// with [PutAccountPolicy].
//
// # Metric extraction policy
//
// A metric extraction policy controls whether CloudWatch Metrics can be created
// through the Embedded Metrics Format (EMF) for log groups in your account. By
// default, EMF metric creation is enabled for all log groups. You can use metric
// extraction policies to disable EMF metric creation for your entire account or
// specific log groups.
//
// When a policy disables EMF metric creation for a log group, log events in the
// EMF format are still ingested, but no CloudWatch Metrics are created from them.
//
// Creating a policy disables metrics for AWS features that use EMF to create
// metrics, such as CloudWatch Container Insights and CloudWatch Application
// Signals. To prevent turning off those features by accident, we recommend that
// you exclude the underlying log-groups through a selection-criteria such as
// LogGroupNamePrefix NOT IN ["/aws/containerinsights",
// "/aws/ecs/containerinsights", "/aws/application-signals/data"] .
//
// Each account can have either one account-level metric extraction policy that
// applies to all log groups, or up to 5 policies that are each scoped to a subset
// of log groups with the selectionCriteria parameter. The selection criteria
// supports filtering by LogGroupName and LogGroupNamePrefix using the operators IN
// and NOT IN . You can specify up to 50 values in each IN or NOT IN list.
//
// The selection criteria can be specified in these formats:
//
//	LogGroupName IN ["log-group-1", "log-group-2"]
//
//	LogGroupNamePrefix NOT IN ["/aws/prefix1", "/aws/prefix2"]
//
// If you have multiple account-level metric extraction policies with selection
// criteria, no two of them can have overlapping criteria. For example, if you have
// one policy with selection criteria LogGroupNamePrefix IN ["my-log"] , you can't
// have another metric extraction policy with selection criteria
// LogGroupNamePrefix IN ["/my-log-prod"] or LogGroupNamePrefix IN ["/my-logging"]
// , as the set of log groups matching these prefixes would be a subset of the log
// groups matching the first policy's prefix, creating an overlap.
//
// When using NOT IN , only one policy with this operator is allowed per account.
//
// When combining policies with IN and NOT IN operators, the overlap check ensures
// that policies don't have conflicting effects. Two policies with IN and NOT IN
// operators do not overlap if and only if every value in the IN policy is
// completely contained within some value in the NOT IN policy. For example:
//
//   - If you have a NOT IN policy for prefix "/aws/lambda" , you can create an IN
//     policy for the exact log group name "/aws/lambda/function1" because the set of
//     log groups matching "/aws/lambda/function1" is a subset of the log groups
//     matching "/aws/lambda" .
//
//   - If you have a NOT IN policy for prefix "/aws/lambda" , you cannot create an
//     IN policy for prefix "/aws" because the set of log groups matching "/aws" is
//     not a subset of the log groups matching "/aws/lambda" .
//
// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
// [PutTransformer]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html
// [PutIndexPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutIndexPolicy.html
// [PutDataProtectionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDataProtectionPolicy.html
// [Protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
// [FilterLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html
// [GetLogGroupFields]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogGroupFields.html
// [Processors that you can use]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
// [Create field indexes to improve query performance and reduce costs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
func (c *Client) PutAccountPolicy(ctx context.Context, params *PutAccountPolicyInput, optFns ...func(*Options)) (*PutAccountPolicyOutput, error) {
	if params == nil {
		params = &PutAccountPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountPolicy", params, optFns, c.addOperationPutAccountPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountPolicyInput struct {

	// Specify the policy, in JSON.
	//
	// Data protection policy
	//
	// A data protection policy must include two JSON blocks:
	//
	//   - The first block must include both a DataIdentifer array and an Operation
	//   property with an Audit action. The DataIdentifer array lists the types of
	//   sensitive data that you want to mask. For more information about the available
	//   options, see [Types of data that you can mask].
	//
	// The Operation property with an Audit action is required to find the sensitive
	//   data terms. This Audit action must contain a FindingsDestination object. You
	//   can optionally use that FindingsDestination object to list one or more
	//   destinations to send audit findings to. If you specify destinations such as log
	//   groups, Firehose streams, and S3 buckets, they must already exist.
	//
	//   - The second block must include both a DataIdentifer array and an Operation
	//   property with an Deidentify action. The DataIdentifer array must exactly match
	//   the DataIdentifer array in the first block of the policy.
	//
	// The Operation property with the Deidentify action is what actually masks the
	//   data, and it must contain the "MaskConfig": {} object. The "MaskConfig": {}
	//   object must be empty.
	//
	// For an example data protection policy, see the Examples section on this page.
	//
	// The contents of the two DataIdentifer arrays must match exactly.
	//
	// In addition to the two JSON blocks, the policyDocument can also include Name ,
	// Description , and Version fields. The Name is different than the operation's
	// policyName parameter, and is used as a dimension when CloudWatch Logs reports
	// audit findings metrics to CloudWatch.
	//
	// The JSON specified in policyDocument can be up to 30,720 characters long.
	//
	// Subscription filter policy
	//
	// A subscription filter policy can include the following attributes in a JSON
	// block:
	//
	//   - DestinationArn The ARN of the destination to deliver log events to.
	//   Supported destinations are:
	//
	//   - An Kinesis Data Streams data stream in the same account as the subscription
	//   policy, for same-account delivery.
	//
	//   - An Firehose data stream in the same account as the subscription policy, for
	//   same-account delivery.
	//
	//   - A Lambda function in the same account as the subscription policy, for
	//   same-account delivery.
	//
	//   - A logical destination in a different account created with [PutDestination], for
	//   cross-account delivery. Kinesis Data Streams and Firehose are supported as
	//   logical destinations.
	//
	//   - RoleArn The ARN of an IAM role that grants CloudWatch Logs permissions to
	//   deliver ingested log events to the destination stream. You don't need to provide
	//   the ARN when you are working with a logical destination for cross-account
	//   delivery.
	//
	//   - FilterPattern A filter pattern for subscribing to a filtered stream of log
	//   events.
	//
	//   - Distribution The method used to distribute log data to the destination. By
	//   default, log data is grouped by log stream, but the grouping can be set to
	//   Random for a more even distribution. This property is only applicable when the
	//   destination is an Kinesis Data Streams data stream.
	//
	// Transformer policy
	//
	// A transformer policy must include one JSON block with the array of processors
	// and their configurations. For more information about available processors, see [Processors that you can use]
	// .
	//
	// Field index policy
	//
	// A field index filter policy can include the following attribute in a JSON block:
	//
	//   - Fields The array of field indexes to create.
	//
	// It must contain at least one field index.
	//
	// The following is an example of an index policy document that creates two
	// indexes, RequestId and TransactionId .
	//
	//     "policyDocument": "{ \"Fields\": [ \"RequestId\", \"TransactionId\" ] }"
	//
	// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
	// [Processors that you can use]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors
	// [Types of data that you can mask]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-types.html
	//
	// This member is required.
	PolicyDocument *string

	// A name for the policy. This must be unique within the account.
	//
	// This member is required.
	PolicyName *string

	// The type of policy that you're creating or updating.
	//
	// This member is required.
	PolicyType types.PolicyType

	// Currently the only valid value for this parameter is ALL , which specifies that
	// the data protection policy applies to all log groups in the account. If you omit
	// this parameter, the default of ALL is used.
	Scope types.Scope

	// Use this parameter to apply the new policy to a subset of log groups in the
	// account.
	//
	// Specifying selectionCriteria is valid only when you specify
	// SUBSCRIPTION_FILTER_POLICY , FIELD_INDEX_POLICY or TRANSFORMER_POLICY for
	// policyType .
	//
	// If policyType is SUBSCRIPTION_FILTER_POLICY , the only supported
	// selectionCriteria filter is LogGroupName NOT IN []
	//
	// If policyType is FIELD_INDEX_POLICY or TRANSFORMER_POLICY , the only supported
	// selectionCriteria filter is LogGroupNamePrefix
	//
	// The selectionCriteria string can be up to 25KB in length. The length is
	// determined by using its UTF-8 bytes.
	//
	// Using the selectionCriteria parameter with SUBSCRIPTION_FILTER_POLICY is useful
	// to help prevent infinite loops. For more information, see [Log recursion prevention].
	//
	// [Log recursion prevention]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscriptions-recursion-prevention.html
	SelectionCriteria *string

	noSmithyDocumentSerde
}

type PutAccountPolicyOutput struct {

	// The account policy that you created.
	AccountPolicy *types.AccountPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountPolicy"); err != nil {
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
	if err = addOpPutAccountPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountPolicy",
	}
}
