// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The CreateTable operation adds a new table to your account. In an Amazon Web
// Services account, table names must be unique within each Region. That is, you
// can have two tables with same name if you create the tables in different
// Regions.
//
// CreateTable is an asynchronous operation. Upon receiving a CreateTable request,
// DynamoDB immediately returns a response with a TableStatus of CREATING . After
// the table is created, DynamoDB sets the TableStatus to ACTIVE . You can perform
// read and write operations only on an ACTIVE table.
//
// You can optionally define secondary indexes on the new table, as part of the
// CreateTable operation. If you want to create multiple tables with secondary
// indexes on them, you must create the tables sequentially. Only one table with
// secondary indexes can be in the CREATING state at any given time.
//
// You can use the DescribeTable action to check the table status.
func (c *Client) CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) {
	if params == nil {
		params = &CreateTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTable", params, optFns, c.addOperationCreateTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateTable operation.
type CreateTableInput struct {

	// An array of attributes that describe the key schema for the table and indexes.
	//
	// This member is required.
	AttributeDefinitions []types.AttributeDefinition

	// Specifies the attributes that make up the primary key for a table or an index.
	// The attributes in KeySchema must also be defined in the AttributeDefinitions
	// array. For more information, see [Data Model]in the Amazon DynamoDB Developer Guide.
	//
	// Each KeySchemaElement in the array is composed of:
	//
	//   - AttributeName - The name of this key attribute.
	//
	//   - KeyType - The role that the key attribute will assume:
	//
	//   - HASH - partition key
	//
	//   - RANGE - sort key
	//
	// The partition key of an item is also known as its hash attribute. The term
	// "hash attribute" derives from the DynamoDB usage of an internal hash function to
	// evenly distribute data items across partitions, based on their partition key
	// values.
	//
	// The sort key of an item is also known as its range attribute. The term "range
	// attribute" derives from the way DynamoDB stores items with the same partition
	// key physically close together, in sorted order by the sort key value.
	//
	// For a simple primary key (partition key), you must provide exactly one element
	// with a KeyType of HASH .
	//
	// For a composite primary key (partition key and sort key), you must provide
	// exactly two elements, in this order: The first element must have a KeyType of
	// HASH , and the second element must have a KeyType of RANGE .
	//
	// For more information, see [Working with Tables] in the Amazon DynamoDB Developer Guide.
	//
	// [Data Model]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html
	// [Working with Tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key
	//
	// This member is required.
	KeySchema []types.KeySchemaElement

	// The name of the table to create. You can also provide the Amazon Resource Name
	// (ARN) of the table in this parameter.
	//
	// This member is required.
	TableName *string

	// Controls how you are charged for read and write throughput and how you manage
	// capacity. This setting can be changed later.
	//
	//   - PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for most DynamoDB
	//   workloads. PAY_PER_REQUEST sets the billing mode to [On-demand capacity mode].
	//
	//   - PROVISIONED - We recommend using PROVISIONED for steady workloads with
	//   predictable growth where capacity requirements can be reliably forecasted.
	//   PROVISIONED sets the billing mode to [Provisioned capacity mode].
	//
	// [Provisioned capacity mode]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html
	// [On-demand capacity mode]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html
	BillingMode types.BillingMode

	// Indicates whether deletion protection is to be enabled (true) or disabled
	// (false) on the table.
	DeletionProtectionEnabled *bool

	// One or more global secondary indexes (the maximum is 20) to be created on the
	// table. Each global secondary index in the array includes the following:
	//
	//   - IndexName - The name of the global secondary index. Must be unique only for
	//   this table.
	//
	//   - KeySchema - Specifies the key schema for the global secondary index.
	//
	//   - Projection - Specifies attributes that are copied (projected) from the table
	//   into the index. These are in addition to the primary key attributes and index
	//   key attributes, which are automatically projected. Each attribute specification
	//   is composed of:
	//
	//   - ProjectionType - One of the following:
	//
	//   - KEYS_ONLY - Only the index and primary keys are projected into the index.
	//
	//   - INCLUDE - Only the specified table attributes are projected into the index.
	//   The list of projected attributes is in NonKeyAttributes .
	//
	//   - ALL - All of the table attributes are projected into the index.
	//
	//   - NonKeyAttributes - A list of one or more non-key attribute names that are
	//   projected into the secondary index. The total count of attributes provided in
	//   NonKeyAttributes , summed across all of the secondary indexes, must not exceed
	//   100. If you project the same attribute into two different indexes, this counts
	//   as two distinct attributes when determining the total. This limit only applies
	//   when you specify the ProjectionType of INCLUDE . You still can specify the
	//   ProjectionType of ALL to project all attributes from the source table, even if
	//   the table has more than 100 attributes.
	//
	//   - ProvisionedThroughput - The provisioned throughput settings for the global
	//   secondary index, consisting of read and write capacity units.
	GlobalSecondaryIndexes []types.GlobalSecondaryIndex

	// One or more local secondary indexes (the maximum is 5) to be created on the
	// table. Each index is scoped to a given partition key value. There is a 10 GB
	// size limit per partition key value; otherwise, the size of a local secondary
	// index is unconstrained.
	//
	// Each local secondary index in the array includes the following:
	//
	//   - IndexName - The name of the local secondary index. Must be unique only for
	//   this table.
	//
	//   - KeySchema - Specifies the key schema for the local secondary index. The key
	//   schema must begin with the same partition key as the table.
	//
	//   - Projection - Specifies attributes that are copied (projected) from the table
	//   into the index. These are in addition to the primary key attributes and index
	//   key attributes, which are automatically projected. Each attribute specification
	//   is composed of:
	//
	//   - ProjectionType - One of the following:
	//
	//   - KEYS_ONLY - Only the index and primary keys are projected into the index.
	//
	//   - INCLUDE - Only the specified table attributes are projected into the index.
	//   The list of projected attributes is in NonKeyAttributes .
	//
	//   - ALL - All of the table attributes are projected into the index.
	//
	//   - NonKeyAttributes - A list of one or more non-key attribute names that are
	//   projected into the secondary index. The total count of attributes provided in
	//   NonKeyAttributes , summed across all of the secondary indexes, must not exceed
	//   100. If you project the same attribute into two different indexes, this counts
	//   as two distinct attributes when determining the total. This limit only applies
	//   when you specify the ProjectionType of INCLUDE . You still can specify the
	//   ProjectionType of ALL to project all attributes from the source table, even if
	//   the table has more than 100 attributes.
	LocalSecondaryIndexes []types.LocalSecondaryIndex

	// Sets the maximum number of read and write units for the specified table in
	// on-demand capacity mode. If you use this parameter, you must specify
	// MaxReadRequestUnits , MaxWriteRequestUnits , or both.
	OnDemandThroughput *types.OnDemandThroughput

	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// If you set BillingMode as PROVISIONED , you must specify this property. If you
	// set BillingMode as PAY_PER_REQUEST , you cannot specify this property.
	//
	// For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas] in the
	// Amazon DynamoDB Developer Guide.
	//
	// [Service, Account, and Table Quotas]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html
	ProvisionedThroughput *types.ProvisionedThroughput

	// An Amazon Web Services resource-based policy document in JSON format that will
	// be attached to the table.
	//
	// When you attach a resource-based policy while creating a table, the policy
	// application is strongly consistent.
	//
	// The maximum size supported for a resource-based policy document is 20 KB.
	// DynamoDB counts whitespaces when calculating the size of a policy against this
	// limit. For a full list of all considerations that apply for resource-based
	// policies, see [Resource-based policy considerations].
	//
	// You need to specify the CreateTable and PutResourcePolicy IAM actions for
	// authorizing a user to create a table with a resource-based policy.
	//
	// [Resource-based policy considerations]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html
	ResourcePolicy *string

	// Represents the settings used to enable server-side encryption.
	SSESpecification *types.SSESpecification

	// The settings for DynamoDB Streams on the table. These settings consist of:
	//
	//   - StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true)
	//   or disabled (false).
	//
	//   - StreamViewType - When an item in the table is modified, StreamViewType
	//   determines what information is written to the table's stream. Valid values for
	//   StreamViewType are:
	//
	//   - KEYS_ONLY - Only the key attributes of the modified item are written to the
	//   stream.
	//
	//   - NEW_IMAGE - The entire item, as it appears after it was modified, is written
	//   to the stream.
	//
	//   - OLD_IMAGE - The entire item, as it appeared before it was modified, is
	//   written to the stream.
	//
	//   - NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are
	//   written to the stream.
	StreamSpecification *types.StreamSpecification

	// The table class of the new table. Valid values are STANDARD and
	// STANDARD_INFREQUENT_ACCESS .
	TableClass types.TableClass

	// A list of key-value pairs to label the table. For more information, see [Tagging for DynamoDB].
	//
	// [Tagging for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html
	Tags []types.Tag

	// Represents the warm throughput (in read units per second and write units per
	// second) for creating a table.
	WarmThroughput *types.WarmThroughput

	noSmithyDocumentSerde
}

func (in *CreateTableInput) bindEndpointParams(p *EndpointParameters) {

	p.ResourceArn = in.TableName

}

// Represents the output of a CreateTable operation.
type CreateTableOutput struct {

	// Represents the properties of the table.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTable"); err != nil {
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
	if err = addOpCreateTableDiscoverEndpointMiddleware(stack, options, c); err != nil {
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
	if err = addUserAgentAccountIDEndpointMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTable(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

func addOpCreateTableDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpCreateTableDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpCreateTableDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*CreateTableInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opCreateTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTable",
	}
}
