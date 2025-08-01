// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a software update for a core or group of cores (specified as an IoT
// thing group.) Use this to update the OTA Agent as well as the Greengrass core
// software. It makes use of the IoT Jobs feature which provides additional
// commands to manage a Greengrass core software update job.
func (c *Client) CreateSoftwareUpdateJob(ctx context.Context, params *CreateSoftwareUpdateJobInput, optFns ...func(*Options)) (*CreateSoftwareUpdateJobOutput, error) {
	if params == nil {
		params = &CreateSoftwareUpdateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSoftwareUpdateJob", params, optFns, c.addOperationCreateSoftwareUpdateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSoftwareUpdateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSoftwareUpdateJobInput struct {

	// The IAM Role that Greengrass will use to create pre-signed URLs pointing
	// towards the update artifact.
	//
	// This member is required.
	S3UrlSignerRole *string

	// The piece of software on the Greengrass core that will be updated.
	//
	// This member is required.
	SoftwareToUpdate types.SoftwareToUpdate

	// The ARNs of the targets (IoT things or IoT thing groups) that this update will
	// be applied to.
	//
	// This member is required.
	UpdateTargets []string

	// The architecture of the cores which are the targets of an update.
	//
	// This member is required.
	UpdateTargetsArchitecture types.UpdateTargetsArchitecture

	// The operating system of the cores which are the targets of an update.
	//
	// This member is required.
	UpdateTargetsOperatingSystem types.UpdateTargetsOperatingSystem

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// The minimum level of log statements that should be logged by the OTA Agent
	// during an update.
	UpdateAgentLogLevel types.UpdateAgentLogLevel

	noSmithyDocumentSerde
}

type CreateSoftwareUpdateJobOutput struct {

	// The IoT Job ARN corresponding to this update.
	IotJobArn *string

	// The IoT Job Id corresponding to this update.
	IotJobId *string

	// The software version installed on the device or devices after the update.
	PlatformSoftwareVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSoftwareUpdateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSoftwareUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSoftwareUpdateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSoftwareUpdateJob"); err != nil {
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
	if err = addOpCreateSoftwareUpdateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSoftwareUpdateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSoftwareUpdateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSoftwareUpdateJob",
	}
}
