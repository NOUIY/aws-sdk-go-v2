// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a tenant database in a DB instance that uses the multi-tenant
// configuration. Only RDS for Oracle container database (CDB) instances are
// supported.
func (c *Client) CreateTenantDatabase(ctx context.Context, params *CreateTenantDatabaseInput, optFns ...func(*Options)) (*CreateTenantDatabaseOutput, error) {
	if params == nil {
		params = &CreateTenantDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTenantDatabase", params, optFns, c.addOperationCreateTenantDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTenantDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTenantDatabaseInput struct {

	// The user-supplied DB instance identifier. RDS creates your tenant database in
	// this DB instance. This parameter isn't case-sensitive.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name for the master user account in your tenant database. RDS creates this
	// user account in the tenant database and grants privileges to the master user.
	// This parameter is case-sensitive.
	//
	// Constraints:
	//
	//   - Must be 1 to 16 letters, numbers, or underscores.
	//
	//   - First character must be a letter.
	//
	//   - Can't be a reserved word for the chosen database engine.
	//
	// This member is required.
	MasterUsername *string

	// The user-supplied name of the tenant database that you want to create in your
	// DB instance. This parameter has the same constraints as DBName in
	// CreateDBInstance .
	//
	// This member is required.
	TenantDBName *string

	// The character set for your tenant database. If you don't specify a value, the
	// character set name defaults to AL32UTF8 .
	CharacterSetName *string

	// Specifies whether to manage the master user password with Amazon Web Services
	// Secrets Manager.
	//
	// For more information, see [Password management with Amazon Web Services Secrets Manager] in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//   - Can't manage the master user password with Amazon Web Services Secrets
	//   Manager if MasterUserPassword is specified.
	//
	// [Password management with Amazon Web Services Secrets Manager]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html
	ManageMasterUserPassword *bool

	// The password for the master user in your tenant database.
	//
	// Constraints:
	//
	//   - Must be 8 to 30 characters.
	//
	//   - Can include any printable ASCII character except forward slash ( / ), double
	//   quote ( " ), at symbol ( @ ), ampersand ( & ), or single quote ( ' ).
	//
	//   - Can't be specified when ManageMasterUserPassword is enabled.
	MasterUserPassword *string

	// The Amazon Web Services KMS key identifier to encrypt a secret that is
	// automatically generated and managed in Amazon Web Services Secrets Manager.
	//
	// This setting is valid only if the master user password is managed by RDS in
	// Amazon Web Services Secrets Manager for the DB instance.
	//
	// The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN,
	// or alias name for the KMS key. To use a KMS key in a different Amazon Web
	// Services account, specify the key ARN or alias ARN.
	//
	// If you don't specify MasterUserSecretKmsKeyId , then the aws/secretsmanager KMS
	// key is used to encrypt the secret. If the secret is in a different Amazon Web
	// Services account, then you can't use the aws/secretsmanager KMS key to encrypt
	// the secret, and you must use a customer managed KMS key.
	//
	// There is a default KMS key for your Amazon Web Services account. Your Amazon
	// Web Services account has a different default KMS key for each Amazon Web
	// Services Region.
	MasterUserSecretKmsKeyId *string

	// The NCHAR value for the tenant database.
	NcharCharacterSetName *string

	// A list of tags.
	//
	// For more information, see [Tagging Amazon RDS resources] in the Amazon RDS User Guide or [Tagging Amazon Aurora and Amazon RDS resources] in the Amazon
	// Aurora User Guide.
	//
	// [Tagging Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	// [Tagging Amazon Aurora and Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTenantDatabaseOutput struct {

	// A tenant database in the DB instance. This data type is an element in the
	// response to the DescribeTenantDatabases action.
	TenantDatabase *types.TenantDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTenantDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTenantDatabase"); err != nil {
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
	if err = addOpCreateTenantDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTenantDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTenantDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTenantDatabase",
	}
}
