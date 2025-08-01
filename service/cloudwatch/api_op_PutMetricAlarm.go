// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an alarm and associates it with the specified metric, metric
// math expression, anomaly detection model, or Metrics Insights query. For more
// information about using a Metrics Insights query for an alarm, see [Create alarms on Metrics Insights queries].
//
// Alarms based on anomaly detection models cannot have Auto Scaling actions.
//
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// If you are an IAM user, you must have Amazon EC2 permissions for some alarm
// operations:
//
//   - The iam:CreateServiceLinkedRole permission for all alarms with EC2 actions
//
//   - The iam:CreateServiceLinkedRole permissions to create an alarm with Systems
//     Manager OpsItem or response plan actions.
//
// The first time you create an alarm in the Amazon Web Services Management
// Console, the CLI, or by using the PutMetricAlarm API, CloudWatch creates the
// necessary service-linked role for you. The service-linked roles are called
// AWSServiceRoleForCloudWatchEvents and
// AWSServiceRoleForCloudWatchAlarms_ActionSSM . For more information, see [Amazon Web Services service-linked role].
//
// Each PutMetricAlarm action has a maximum uncompressed payload of 120 KB.
//
// # Cross-account alarms
//
// You can set an alarm on metrics in the current account, or in another account.
// To create a cross-account alarm that watches a metric in a different account,
// you must have completed the following pre-requisites:
//
//   - The account where the metrics are located (the sharing account) must
//     already have a sharing role named CloudWatch-CrossAccountSharingRole. If it does
//     not already have this role, you must create it using the instructions in Set up
//     a sharing account in [Cross-account cross-Region CloudWatch console]. The policy for that role must grant access to the ID
//     of the account where you are creating the alarm.
//
//   - The account where you are creating the alarm (the monitoring account) must
//     already have a service-linked role named AWSServiceRoleForCloudWatchCrossAccount
//     to allow CloudWatch to assume the sharing role in the sharing account. If it
//     does not, you must create it following the directions in Set up a monitoring
//     account in [Cross-account cross-Region CloudWatch console].
//
// [Amazon Web Services service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role
// [Create alarms on Metrics Insights queries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Metrics_Insights_Alarm.html
// [Cross-account cross-Region CloudWatch console]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Cross-Account-Cross-Region.html#enable-cross-account-cross-Region
func (c *Client) PutMetricAlarm(ctx context.Context, params *PutMetricAlarmInput, optFns ...func(*Options)) (*PutMetricAlarmOutput, error) {
	if params == nil {
		params = &PutMetricAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricAlarm", params, optFns, c.addOperationPutMetricAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricAlarmInput struct {

	// The name for the alarm. This name must be unique within the Region.
	//
	// The name must contain only UTF-8 characters, and can't contain ASCII control
	// characters
	//
	// This member is required.
	AlarmName *string

	//  The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand.
	//
	// The values LessThanLowerOrGreaterThanUpperThreshold , LessThanLowerThreshold ,
	// and GreaterThanUpperThreshold are used only for alarms based on anomaly
	// detection models.
	//
	// This member is required.
	ComparisonOperator types.ComparisonOperator

	// The number of periods over which data is compared to the specified threshold.
	// If you are setting an alarm that requires that a number of consecutive data
	// points be breaching to trigger the alarm, this value specifies that number. If
	// you are setting an "M out of N" alarm, this value is the N.
	//
	// This member is required.
	EvaluationPeriods *int32

	// Indicates whether actions should be executed during any changes to the alarm
	// state. The default is TRUE .
	ActionsEnabled *bool

	// The actions to execute when this alarm transitions to the ALARM state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN). Valid
	// values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   -
	//   arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   -
	//   arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Lambda actions:
	//
	//   - Invoke the latest version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name
	//
	//   - Invoke a specific version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:version-number
	//
	//   - Invoke a function by using an alias Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:alias-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	//
	// Start a Amazon Q Developer operational investigation
	//
	//     arn:aws:aiops:region:account-id:investigation-group:investigation-group-id
	AlarmActions []string

	// The description for the alarm.
	AlarmDescription *string

	// The number of data points that must be breaching to trigger the alarm. This is
	// used only if you are setting an "M out of N" alarm. In that case, this value is
	// the M. For more information, see [Evaluating an Alarm]in the Amazon CloudWatch User Guide.
	//
	// [Evaluating an Alarm]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation
	DatapointsToAlarm *int32

	// The dimensions for the metric specified in MetricName .
	Dimensions []types.Dimension

	//  Used only for alarms based on percentiles. If you specify ignore , the alarm
	// state does not change during periods with too few data points to be
	// statistically significant. If you specify evaluate or omit this parameter, the
	// alarm is always evaluated and possibly changes state no matter how many data
	// points are available. For more information, see [Percentile-Based CloudWatch Alarms and Low Data Samples].
	//
	// Valid Values: evaluate | ignore
	//
	// [Percentile-Based CloudWatch Alarms and Low Data Samples]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples
	EvaluateLowSampleCountPercentile *string

	// The extended statistic for the metric specified in MetricName . When you call
	// PutMetricAlarm and specify a MetricName , you must specify either Statistic or
	// ExtendedStatistic but not both.
	//
	// If you specify ExtendedStatistic , the following are valid values:
	//
	//   - p90
	//
	//   - tm90
	//
	//   - tc90
	//
	//   - ts90
	//
	//   - wm90
	//
	//   - IQM
	//
	//   - PR(n:m) where n and m are values of the metric
	//
	//   - TC(X%:X%) where X is between 10 and 90 inclusive.
	//
	//   - TM(X%:X%) where X is between 10 and 90 inclusive.
	//
	//   - TS(X%:X%) where X is between 10 and 90 inclusive.
	//
	//   - WM(X%:X%) where X is between 10 and 90 inclusive.
	//
	// For more information about these extended statistics, see [CloudWatch statistics definitions].
	//
	// [CloudWatch statistics definitions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html
	ExtendedStatistic *string

	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource Name
	// (ARN). Valid values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   -
	//   arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   -
	//   arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Lambda actions:
	//
	//   - Invoke the latest version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name
	//
	//   - Invoke a specific version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:version-number
	//
	//   - Invoke a function by using an alias Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:alias-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	InsufficientDataActions []string

	// The name for the metric associated with the alarm. For each PutMetricAlarm
	// operation, you must specify either MetricName or a Metrics array.
	//
	// If you are creating an alarm based on a math expression, you cannot specify
	// this parameter, or any of the Namespace , Dimensions , Period , Unit , Statistic
	// , or ExtendedStatistic parameters. Instead, you specify all this information in
	// the Metrics array.
	MetricName *string

	// An array of MetricDataQuery structures that enable you to create an alarm based
	// on the result of a metric math expression. For each PutMetricAlarm operation,
	// you must specify either MetricName or a Metrics array.
	//
	// Each item in the Metrics array either retrieves a metric or performs a math
	// expression.
	//
	// One item in the Metrics array is the expression that the alarm watches. You
	// designate this expression by setting ReturnData to true for this object in the
	// array. For more information, see [MetricDataQuery].
	//
	// If you use the Metrics parameter, you cannot include the Namespace , MetricName
	// , Dimensions , Period , Unit , Statistic , or ExtendedStatistic parameters of
	// PutMetricAlarm in the same operation. Instead, you retrieve the metrics you are
	// using in your math expression as part of the Metrics array.
	//
	// [MetricDataQuery]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDataQuery.html
	Metrics []types.MetricDataQuery

	// The namespace for the metric associated specified in MetricName .
	Namespace *string

	// The actions to execute when this alarm transitions to an OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN). Valid
	// values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   -
	//   arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   -
	//   arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// Lambda actions:
	//
	//   - Invoke the latest version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name
	//
	//   - Invoke a specific version of a Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:version-number
	//
	//   - Invoke a function by using an alias Lambda function:
	//   arn:aws:lambda:region:account-id:function:function-name:alias-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	OKActions []string

	// The length, in seconds, used each time the metric specified in MetricName is
	// evaluated. Valid values are 10, 20, 30, and any multiple of 60.
	//
	// Period is required for alarms based on static thresholds. If you are creating
	// an alarm based on a metric math expression, you specify the period for each
	// metric within the objects in the Metrics array.
	//
	// Be sure to specify 10, 20, or 30 only for metrics that are stored by a
	// PutMetricData call with a StorageResolution of 1. If you specify a period of
	// 10, 20, or 30 for a metric that does not have sub-minute resolution, the alarm
	// still attempts to gather data at the period rate that you specify. In this case,
	// it does not receive data for the attempts that do not correspond to a one-minute
	// data resolution, and the alarm might often lapse into INSUFFICENT_DATA status.
	// Specifying 10, 20, or 30 also sets this alarm as a high-resolution alarm, which
	// has a higher charge than other alarms. For more information about pricing, see [Amazon CloudWatch Pricing].
	//
	// An alarm's total current evaluation period can be no longer than seven days, so
	// Period multiplied by EvaluationPeriods can't be more than 604,800 seconds. For
	// alarms with a period of less than one hour (3,600 seconds), the total evaluation
	// period can't be longer than one day (86,400 seconds).
	//
	// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
	Period *int32

	// The statistic for the metric specified in MetricName , other than percentile.
	// For percentile statistics, use ExtendedStatistic . When you call PutMetricAlarm
	// and specify a MetricName , you must specify either Statistic or
	// ExtendedStatistic, but not both.
	Statistic types.Statistic

	// A list of key-value pairs to associate with the alarm. You can associate as
	// many as 50 tags with an alarm. To be able to associate tags with the alarm when
	// you create the alarm, you must have the cloudwatch:TagResource permission.
	//
	// Tags can help you organize and categorize your resources. You can also use them
	// to scope user permissions by granting a user permission to access or change only
	// resources with certain tag values.
	//
	// If you are using this operation to update an existing alarm, any tags you
	// specify in this parameter are ignored. To change the tags of an existing alarm,
	// use [TagResource]or [UntagResource].
	//
	// To use this field to set tags for an alarm when you create it, you must be
	// signed on with both the cloudwatch:PutMetricAlarm and cloudwatch:TagResource
	// permissions.
	//
	// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html
	// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_UntagResource.html
	Tags []types.Tag

	// The value against which the specified statistic is compared.
	//
	// This parameter is required for alarms based on static thresholds, but should
	// not be used for alarms based on anomaly detection models.
	Threshold *float64

	// If this is an alarm based on an anomaly detection model, make this value match
	// the ID of the ANOMALY_DETECTION_BAND function.
	//
	// For an example of how to use this parameter, see the Anomaly Detection Model
	// Alarm example on this page.
	//
	// If your alarm uses this parameter, it cannot have Auto Scaling actions.
	ThresholdMetricId *string

	//  Sets how this alarm is to handle missing data points. If TreatMissingData is
	// omitted, the default behavior of missing is used. For more information, see [Configuring How CloudWatch Alarms Treats Missing Data].
	//
	// Valid Values: breaching | notBreaching | ignore | missing
	//
	// Alarms that evaluate metrics in the AWS/DynamoDB namespace always ignore
	// missing data even if you choose a different option for TreatMissingData . When
	// an AWS/DynamoDB metric has missing data, alarms that evaluate that metric
	// remain in their current state.
	//
	// [Configuring How CloudWatch Alarms Treats Missing Data]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data
	TreatMissingData *string

	// The unit of measure for the statistic. For example, the units for the Amazon
	// EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes that
	// an instance receives on all network interfaces. You can also specify a unit when
	// you create a custom metric. Units help provide conceptual meaning to your data.
	// Metric data points that specify a unit of measure, such as Percent, are
	// aggregated separately. If you are creating an alarm based on a metric math
	// expression, you can specify the unit for each metric (if needed) within the
	// objects in the Metrics array.
	//
	// If you don't specify Unit , CloudWatch retrieves all unit types that have been
	// published for the metric and attempts to evaluate the alarm. Usually, metrics
	// are published with only one unit, so the alarm works as intended.
	//
	// However, if the metric is published with multiple types of units and you don't
	// specify a unit, the alarm's behavior is not defined and it behaves
	// unpredictably.
	//
	// We recommend omitting Unit so that you don't inadvertently specify an incorrect
	// unit that is not published for this metric. Doing so causes the alarm to be
	// stuck in the INSUFFICIENT DATA state.
	Unit types.StandardUnit

	noSmithyDocumentSerde
}

type PutMetricAlarmOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMetricAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutMetricAlarm"); err != nil {
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
	if err = addOpPutMetricAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricAlarm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetricAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutMetricAlarm",
	}
}
