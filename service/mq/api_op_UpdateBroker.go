// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a pending configuration change to a broker.
func (c *Client) UpdateBroker(ctx context.Context, params *UpdateBrokerInput, optFns ...func(*Options)) (*UpdateBrokerOutput, error) {
	if params == nil {
		params = &UpdateBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBroker", params, optFns, c.addOperationUpdateBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates the broker using the specified properties.
type UpdateBrokerInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string

	// Optional. The authentication strategy used to secure the broker. The default is
	// SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Enables automatic upgrades to new patch versions for brokers as new versions
	// are released and supported by Amazon MQ. Automatic upgrades occur during the
	// scheduled maintenance window or after a manual broker reboot.
	//
	// Must be set to true for ActiveMQ brokers version 5.18 and above and for
	// RabbitMQ brokers version 3.13 and above.
	AutoMinorVersionUpgrade *bool

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// Defines whether this broker is a part of a data replication pair.
	DataReplicationMode types.DataReplicationMode

	// The broker engine version. For more information, see the [ActiveMQ version management] and the [RabbitMQ version management] sections in
	// the Amazon MQ Developer Guide.
	//
	// When upgrading to ActiveMQ version 5.18 and above or RabbitMQ version 3.13 and
	// above, you must have autoMinorVersionUpgrade set to true for the broker.
	//
	// [RabbitMQ version management]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/rabbitmq-version-management.html
	// [ActiveMQ version management]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/activemq-version-management.html
	EngineVersion *string

	// The broker's host instance type to upgrade to. For a list of supported instance
	// types, see [Broker instance types].
	//
	// [Broker instance types]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker.html#broker-instance-types
	HostInstanceType *string

	// Optional. The metadata of the LDAP server used to authenticate and authorize
	// connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	noSmithyDocumentSerde
}

type UpdateBrokerOutput struct {

	// Optional. The authentication strategy used to secure the broker. The default is
	// SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Enables automatic upgrades to new patch versions for brokers as new versions
	// are released and supported by Amazon MQ. Automatic upgrades occur during the
	// scheduled maintenance window or after a manual broker reboot.
	AutoMinorVersionUpgrade *bool

	// Required. The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// The ID of the updated configuration.
	Configuration *types.ConfigurationId

	// The replication details of the data replication-enabled broker. Only returned
	// if dataReplicationMode is set to CRDR.
	DataReplicationMetadata *types.DataReplicationMetadataOutput

	// Describes whether this broker is a part of a data replication pair.
	DataReplicationMode types.DataReplicationMode

	// The broker engine version to upgrade to. For more information, see the [ActiveMQ version management] and the [RabbitMQ version management]
	// sections in the Amazon MQ Developer Guide.
	//
	// [RabbitMQ version management]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/rabbitmq-version-management.html
	// [ActiveMQ version management]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/activemq-version-management.html
	EngineVersion *string

	// The broker's host instance type to upgrade to. For a list of supported instance
	// types, see [Broker instance types].
	//
	// [Broker instance types]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker.html#broker-instance-types
	HostInstanceType *string

	// Optional. The metadata of the LDAP server used to authenticate and authorize
	// connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata *types.LdapServerMetadataOutput

	// The list of information about logs to be enabled for the specified broker.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The pending replication details of the data replication-enabled broker. Only
	// returned if pendingDataReplicationMode is set to CRDR.
	PendingDataReplicationMetadata *types.DataReplicationMetadataOutput

	// Describes whether this broker will be a part of a data replication pair after
	// reboot.
	PendingDataReplicationMode types.DataReplicationMode

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateBroker"); err != nil {
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
	if err = addOpUpdateBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBroker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateBroker",
	}
}
