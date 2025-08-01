// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates (registers) a data catalog with the specified name and properties.
// Catalogs created are visible to all users of the same Amazon Web Services
// account.
//
// For a FEDERATED catalog, this API operation creates the following resources.
//
//   - CFN Stack Name with a maximum length of 128 characters and prefix
//     athenafederatedcatalog-CATALOG_NAME_SANITIZED with length 23 characters.
//
//   - Lambda Function Name with a maximum length of 64 characters and prefix
//     athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
//
//   - Glue Connection Name with a maximum length of 255 characters and a prefix
//     athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
func (c *Client) CreateDataCatalog(ctx context.Context, params *CreateDataCatalogInput, optFns ...func(*Options)) (*CreateDataCatalogOutput, error) {
	if params == nil {
		params = &CreateDataCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataCatalog", params, optFns, c.addOperationCreateDataCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataCatalogInput struct {

	// The name of the data catalog to create. The catalog name must be unique for the
	// Amazon Web Services account and can use a maximum of 127 alphanumeric,
	// underscore, at sign, or hyphen characters. The remainder of the length
	// constraint of 256 is reserved for use by Athena.
	//
	// For FEDERATED type the catalog name has following considerations and limits:
	//
	//   - The catalog name allows special characters such as _ , @ , \ , - . These
	//   characters are replaced with a hyphen (-) when creating the CFN Stack Name and
	//   with an underscore (_) when creating the Lambda Function and Glue Connection
	//   Name.
	//
	//   - The catalog name has a theoretical limit of 128 characters. However, since
	//   we use it to create other resources that allow less characters and we prepend a
	//   prefix to it, the actual catalog name limit for FEDERATED catalog is 64 - 23 =
	//   41 characters.
	//
	// This member is required.
	Name *string

	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for an
	// Glue Data Catalog, and HIVE for an external Apache Hive metastore. FEDERATED is
	// a federated catalog for which Athena creates the connection and the Lambda
	// function for you based on the parameters that you pass.
	//
	// For FEDERATED type, we do not support IAM identity center.
	//
	// This member is required.
	Type types.DataCatalogType

	// A description of the data catalog to be created.
	Description *string

	// Specifies the Lambda function or functions to use for creating the data
	// catalog. This is a mapping whose values depend on the catalog type.
	//
	//   - For the HIVE data catalog type, use the following syntax. The
	//   metadata-function parameter is required. The sdk-version parameter is optional
	//   and defaults to the currently supported version.
	//
	// metadata-function=lambda_arn, sdk-version=version_number
	//
	//   - For the LAMBDA data catalog type, use one of the following sets of required
	//   parameters, but not both.
	//
	//   - If you have one Lambda function that processes metadata and another for
	//   reading the actual data, use the following syntax. Both parameters are required.
	//
	// metadata-function=lambda_arn, record-function=lambda_arn
	//
	//   - If you have a composite Lambda function that processes both metadata and
	//   data, use the following syntax to specify your Lambda function.
	//
	// function=lambda_arn
	//
	//   - The GLUE type takes a catalog ID parameter and is required. The catalog_id
	//   is the account ID of the Amazon Web Services account to which the Glue Data
	//   Catalog belongs.
	//
	// catalog-id=catalog_id
	//
	//   - The GLUE data catalog type also applies to the default AwsDataCatalog that
	//   already exists in your account, of which you can have only one and cannot
	//   modify.
	//
	//   - The FEDERATED data catalog type uses one of the following parameters, but
	//   not both. Use connection-arn for an existing Glue connection. Use
	//   connection-type and connection-properties to specify the configuration setting
	//   for a new connection.
	//
	//   - connection-arn:
	//
	//   - lambda-role-arn (optional): The execution role to use for the Lambda
	//   function. If not provided, one is created.
	//
	//   - connection-type:MYSQL|REDSHIFT|...., connection-properties:""
	//
	// For , use escaped JSON text, as in the following example.
	//
	//   "{\"spill_bucket\":\"my_spill\",\"spill_prefix\":\"athena-spill\",\"host\":\"abc12345.snowflakecomputing.com\",\"port\":\"1234\",\"warehouse\":\"DEV_WH\",\"database\":\"TEST\",\"schema\":\"PUBLIC\",\"SecretArn\":\"arn:aws:secretsmanager:ap-south-1:111122223333:secret:snowflake-XHb67j\"}"
	Parameters map[string]string

	// A list of comma separated tags to add to the data catalog that is created. All
	// the resources that are created by the CreateDataCatalog API operation with
	// FEDERATED type will have the tag federated_athena_datacatalog="true" . This
	// includes the CFN Stack, Glue Connection, Athena DataCatalog, and all the
	// resources created as part of the CFN Stack (Lambda Function, IAM
	// policies/roles).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataCatalogOutput struct {

	// Contains information about a data catalog in an Amazon Web Services account.
	//
	// In the Athena console, data catalogs are listed as "data sources" on the Data
	// sources page under the Data source name column.
	DataCatalog *types.DataCatalog

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataCatalog"); err != nil {
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
	if err = addOpCreateDataCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataCatalog",
	}
}
