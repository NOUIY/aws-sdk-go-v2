// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the following configuration parameters of the Hadoop Distributed File
// System (HDFS) transfer location that you're using with DataSync.
//
// For more information, see [Configuring DataSync transfers with an HDFS cluster].
//
// [Configuring DataSync transfers with an HDFS cluster]: https://docs.aws.amazon.com/datasync/latest/userguide/create-hdfs-location.html
func (c *Client) UpdateLocationHdfs(ctx context.Context, params *UpdateLocationHdfsInput, optFns ...func(*Options)) (*UpdateLocationHdfsOutput, error) {
	if params == nil {
		params = &UpdateLocationHdfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationHdfs", params, optFns, c.addOperationUpdateLocationHdfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationHdfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationHdfsInput struct {

	// The Amazon Resource Name (ARN) of the source HDFS cluster location.
	//
	// This member is required.
	LocationArn *string

	// The Amazon Resource Names (ARNs) of the DataSync agents that can connect to
	// your HDFS cluster.
	AgentArns []string

	// The type of authentication used to determine the identity of the user.
	AuthenticationType types.HdfsAuthenticationType

	// The size of the data blocks to write into the HDFS cluster.
	BlockSize *int32

	// The Kerberos key table (keytab) that contains mappings between the defined
	// Kerberos principal and the encrypted keys. You can load the keytab from a file
	// by providing the file's address.
	KerberosKeytab []byte

	// The krb5.conf file that contains the Kerberos configuration information. You
	// can load the krb5.conf file by providing the file's address. If you're using
	// the CLI, it performs the base64 encoding for you. Otherwise, provide the
	// base64-encoded text.
	KerberosKrb5Conf []byte

	// The Kerberos principal with access to the files and folders on the HDFS
	// cluster.
	KerberosPrincipal *string

	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string

	// The NameNode that manages the HDFS namespace. The NameNode performs operations
	// such as opening, closing, and renaming files and directories. The NameNode
	// contains the information to map blocks of data to the DataNodes. You can use
	// only one NameNode.
	NameNodes []types.HdfsNameNode

	// The Quality of Protection (QOP) configuration specifies the Remote Procedure
	// Call (RPC) and data transfer privacy settings configured on the Hadoop
	// Distributed File System (HDFS) cluster.
	QopConfiguration *types.QopConfiguration

	// The number of DataNodes to replicate the data to when writing to the HDFS
	// cluster.
	ReplicationFactor *int32

	// The user name used to identify the client on the host operating system.
	SimpleUser *string

	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from
	// or write data to the HDFS cluster.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationHdfsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationHdfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationHdfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationHdfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationHdfs"); err != nil {
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
	if err = addOpUpdateLocationHdfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationHdfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationHdfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationHdfs",
	}
}
