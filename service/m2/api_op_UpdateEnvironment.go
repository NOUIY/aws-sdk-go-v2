// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration details for a specific runtime environment.
func (c *Client) UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironment", params, optFns, c.addOperationUpdateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentInput struct {

	// The unique identifier of the runtime environment that you want to update.
	//
	// This member is required.
	EnvironmentId *string

	// Indicates whether to update the runtime environment during the maintenance
	// window. The default is false. Currently, Amazon Web Services Mainframe
	// Modernization accepts the engineVersion parameter only if
	// applyDuringMaintenanceWindow is true. If any parameter other than engineVersion
	// is provided in UpdateEnvironmentRequest , it will fail if
	// applyDuringMaintenanceWindow is set to true.
	ApplyDuringMaintenanceWindow bool

	// The desired capacity for the runtime environment to update. The minimum
	// possible value is 0 and the maximum is 100.
	DesiredCapacity *int32

	// The version of the runtime engine for the runtime environment.
	EngineVersion *string

	// Forces the updates on the environment. This option is needed if the
	// applications in the environment are not stopped or if there are ongoing
	// application-related activities in the environment.
	//
	// If you use this option, be aware that it could lead to data corruption in the
	// applications, and that you might need to perform repair and recovery procedures
	// for the applications.
	//
	// This option is not needed if the attribute being updated is
	// preferredMaintenanceWindow .
	ForceUpdate bool

	// The instance type for the runtime environment to update.
	InstanceType *string

	// Configures the maintenance window that you want for the runtime environment.
	// The maintenance window must have the format ddd:hh24:mi-ddd:hh24:mi and must be
	// less than 24 hours. The following two examples are valid maintenance windows:
	// sun:23:45-mon:00:15 or sat:01:00-sat:03:00 .
	//
	// If you do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow *string

	noSmithyDocumentSerde
}

type UpdateEnvironmentOutput struct {

	// The unique identifier of the runtime environment that was updated.
	//
	// This member is required.
	EnvironmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnvironment"); err != nil {
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
	if err = addOpUpdateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnvironment",
	}
}
