// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables you to write your time-series data into Timestream. You can specify a
// single data point or a batch of data points to be inserted into the system.
// Timestream offers you a flexible schema that auto detects the column names and
// data types for your Timestream tables based on the dimension names and data
// types of the data points you specify when invoking writes into the database.
//
// Timestream supports eventual consistency read semantics. This means that when
// you query data immediately after writing a batch of data into Timestream, the
// query results might not reflect the results of a recently completed write
// operation. The results may also include some stale data. If you repeat the query
// request after a short time, the results should return the latest data. [Service quotas apply].
//
// See [code sample] for details.
//
// # Upserts
//
// You can use the Version parameter in a WriteRecords request to update data
// points. Timestream tracks a version number with each record. Version defaults
// to 1 when it's not specified for the record in the request. Timestream updates
// an existing record’s measure value along with its Version when it receives a
// write request with a higher Version number for that record. When it receives an
// update request where the measure value is the same as that of the existing
// record, Timestream still updates Version , if it is greater than the existing
// value of Version . You can update a data point as many times as desired, as long
// as the value of Version continuously increases.
//
// For example, suppose you write a new record without indicating Version in the
// request. Timestream stores this record, and set Version to 1 . Now, suppose you
// try to update this record with a WriteRecords request of the same record with a
// different measure value but, like before, do not provide Version . In this case,
// Timestream will reject this update with a RejectedRecordsException since the
// updated record’s version is not greater than the existing value of Version.
//
// However, if you were to resend the update request with Version set to 2 ,
// Timestream would then succeed in updating the record’s value, and the Version
// would be set to 2 . Next, suppose you sent a WriteRecords request with this
// same record and an identical measure value, but with Version set to 3 . In this
// case, Timestream would only update Version to 3 . Any further updates would need
// to send a version number greater than 3 , or the update requests would receive a
// RejectedRecordsException .
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.write.html
func (c *Client) WriteRecords(ctx context.Context, params *WriteRecordsInput, optFns ...func(*Options)) (*WriteRecordsOutput, error) {
	if params == nil {
		params = &WriteRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "WriteRecords", params, optFns, c.addOperationWriteRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*WriteRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type WriteRecordsInput struct {

	// The name of the Timestream database.
	//
	// This member is required.
	DatabaseName *string

	// An array of records that contain the unique measure, dimension, time, and
	// version attributes for each time-series data point.
	//
	// This member is required.
	Records []types.Record

	// The name of the Timestream table.
	//
	// This member is required.
	TableName *string

	// A record that contains the common measure, dimension, time, and version
	// attributes shared across all the records in the request. The measure and
	// dimension attributes specified will be merged with the measure and dimension
	// attributes in the records object when the data is written into Timestream.
	// Dimensions may not overlap, or a ValidationException will be thrown. In other
	// words, a record must contain dimensions with unique names.
	CommonAttributes *types.Record

	noSmithyDocumentSerde
}

type WriteRecordsOutput struct {

	// Information on the records ingested by this request.
	RecordsIngested *types.RecordsIngested

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationWriteRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpWriteRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpWriteRecords{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "WriteRecords"); err != nil {
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
	if err = addOpWriteRecordsDiscoverEndpointMiddleware(stack, options, c); err != nil {
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
	if err = addOpWriteRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opWriteRecords(options.Region), middleware.Before); err != nil {
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

func addOpWriteRecordsDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpWriteRecordsDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpWriteRecordsDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*WriteRecordsInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("Timestream Write.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	if err != nil {
		return internalEndpointDiscovery.WeightedAddress{}, err
	}

	weighted, ok := endpoint.GetValidAddress()
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")
	}
	return weighted, nil
}

func newServiceMetadataMiddleware_opWriteRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "WriteRecords",
	}
}
