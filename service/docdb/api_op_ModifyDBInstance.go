// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies settings for an instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request.
func (c *Client) ModifyDBInstance(ctx context.Context, params *ModifyDBInstanceInput, optFns ...func(*Options)) (*ModifyDBInstanceOutput, error) {
	if params == nil {
		params = &ModifyDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBInstance", params, optFns, c.addOperationModifyDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBInstance.
type ModifyDBInstanceInput struct {

	// The instance identifier. This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DBInstance .
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Specifies whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the instance.
	//
	// If this parameter is set to false , changes to the instance are applied during
	// the next maintenance window. Some parameter changes can cause an outage and are
	// applied on the next reboot.
	//
	// Default: false
	ApplyImmediately *bool

	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not
	// perform minor version upgrades regardless of the value set.
	AutoMinorVersionUpgrade *bool

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string

	// Specifies whether the DB instance is restarted when you rotate your SSL/TLS
	// certificate.
	//
	// By default, the DB instance is restarted when you rotate your SSL/TLS
	// certificate. The certificate is not updated until the DB instance is restarted.
	//
	// Set this parameter only if you are not using SSL/TLS to connect to the DB
	// instance.
	//
	// If you are using SSL/TLS to connect to the DB instance, see [Updating Your Amazon DocumentDB TLS Certificates] and [Encrypting Data in Transit] in the Amazon
	// DocumentDB Developer Guide.
	//
	// [Updating Your Amazon DocumentDB TLS Certificates]: https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html
	// [Encrypting Data in Transit]: https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html
	CertificateRotationRestart *bool

	// A value that indicates whether to copy all tags from the DB instance to
	// snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// The new compute and memory capacity of the instance; for example, db.r5.large .
	// Not all instance classes are available in all Amazon Web Services Regions.
	//
	// If you modify the instance class, an outage occurs during the change. The
	// change is applied during the next maintenance window, unless ApplyImmediately
	// is specified as true for this request.
	//
	// Default: Uses existing setting.
	DBInstanceClass *string

	// A value that indicates whether to enable Performance Insights for the DB
	// Instance. For more information, see [Using Amazon Performance Insights].
	//
	// [Using Amazon Performance Insights]: https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html
	EnablePerformanceInsights *bool

	//  The new instance identifier for the instance when renaming an instance. When
	// you change the instance identifier, an instance reboot occurs immediately if you
	// set Apply Immediately to true . It occurs during the next maintenance window if
	// you set Apply Immediately to false . This value is stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	NewDBInstanceIdentifier *string

	// The KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the
	// KMS key.
	//
	// If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// DocumentDB uses your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services region.
	PerformanceInsightsKMSKeyId *string

	// The weekly time range (in UTC) during which system maintenance can occur, which
	// might result in an outage. Changing this parameter doesn't result in an outage
	// except in the following situation, and the change is asynchronously applied as
	// soon as possible. If there are pending actions that cause a reboot, and the
	// maintenance window is changed to include the current time, changing this
	// parameter causes a reboot of the instance. If you are moving this window to the
	// current time, there must be at least 30 minutes between the current time and end
	// of the window to ensure that pending changes are applied.
	//
	// Default: Uses existing setting.
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Must be at least 30 minutes.
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance.
	//
	// Default: 1
	//
	// Valid values: 0-15
	PromotionTier *int32

	noSmithyDocumentSerde
}

type ModifyDBInstanceOutput struct {

	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBInstance"); err != nil {
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
	if err = addOpModifyDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBInstance",
	}
}
