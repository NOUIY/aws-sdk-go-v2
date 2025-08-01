// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a DataSource object. A DataSource references data that can be used to
// perform CreateMLModel , CreateEvaluation , or CreateBatchPrediction operations.
//
// CreateDataSourceFromS3 is an asynchronous operation. In response to
// CreateDataSourceFromS3 , Amazon Machine Learning (Amazon ML) immediately returns
// and sets the DataSource status to PENDING . After the DataSource has been
// created and is ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in the COMPLETED or PENDING state can be used to perform only
// CreateMLModel , CreateEvaluation or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// The observation data used in a DataSource should be ready to use; that is, it
// should have a consistent structure, and missing data values should be kept to a
// minimum. The observation data must reside in one or more .csv files in an Amazon
// Simple Storage Service (Amazon S3) location, along with a schema that describes
// the data items by name and type. The same schema must be used for all of the
// data files referenced by the DataSource .
//
// After the DataSource has been created, it's ready to use in evaluations and
// batch predictions. If you plan to use the DataSource to train an MLModel , the
// DataSource also needs a recipe. A recipe describes how each input variable will
// be used in training an MLModel . Will the variable be included or excluded from
// training? Will the variable be manipulated; for example, will it be combined
// with another variable or will it be split apart into word combinations? The
// recipe provides answers to these questions.
func (c *Client) CreateDataSourceFromS3(ctx context.Context, params *CreateDataSourceFromS3Input, optFns ...func(*Options)) (*CreateDataSourceFromS3Output, error) {
	if params == nil {
		params = &CreateDataSourceFromS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSourceFromS3", params, optFns, c.addOperationCreateDataSourceFromS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceFromS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceFromS3Input struct {

	// A user-supplied identifier that uniquely identifies the DataSource .
	//
	// This member is required.
	DataSourceId *string

	// The data specification of a DataSource :
	//
	//   - DataLocationS3 - The Amazon S3 location of the observation data.
	//
	//   - DataSchemaLocationS3 - The Amazon S3 location of the DataSchema .
	//
	//   - DataSchema - A JSON string representing the schema. This is not required if
	//   DataSchemaUri is specified.
	//
	//   - DataRearrangement - A JSON string that represents the splitting and
	//   rearrangement requirements for the Datasource .
	//
	// Sample - "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// This member is required.
	DataSpec *types.S3DataSpec

	// The compute statistics for a DataSource . The statistics are generated from the
	// observation data referenced by a DataSource . Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if the
	// DataSource needs to be used for MLModel training.
	ComputeStatistics bool

	// A user-supplied name or description of the DataSource .
	DataSourceName *string

	noSmithyDocumentSerde
}

//	Represents the output of a CreateDataSourceFromS3 operation, and is an
//
// acknowledgement that Amazon ML received the request.
//
// The CreateDataSourceFromS3 operation is asynchronous. You can poll for updates
// by using the GetBatchPrediction operation and checking the Status parameter.
type CreateDataSourceFromS3Output struct {

	// A user-supplied ID that uniquely identifies the DataSource . This value should
	// be identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceFromS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSourceFromS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSourceFromS3{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataSourceFromS3"); err != nil {
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
	if err = addOpCreateDataSourceFromS3ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSourceFromS3(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataSourceFromS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataSourceFromS3",
	}
}
