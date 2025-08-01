// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a Stage.
func (c *Client) GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) {
	if params == nil {
		params = &GetStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStage", params, optFns, c.addOperationGetStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStageInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The stage name. Stage names can only contain alphanumeric characters, hyphens,
	// and underscores. Maximum length is 128 characters.
	//
	// This member is required.
	StageName *string

	noSmithyDocumentSerde
}

type GetStageOutput struct {

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether a stage is managed by API Gateway. If you created an API
	// using quick create, the $default stage is managed by API Gateway. You can't
	// modify the $default stage.
	ApiGatewayManaged *bool

	// Specifies whether updates to an API automatically trigger a new deployment. The
	// default value is false.
	AutoDeploy *bool

	// The identifier of a client certificate for a Stage. Supported only for
	// WebSocket APIs.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// Default route settings for the stage.
	DefaultRouteSettings *types.RouteSettings

	// The identifier of the Deployment that the Stage is associated with. Can't be
	// updated if autoDeploy is enabled.
	DeploymentId *string

	// The description of the stage.
	Description *string

	// Describes the status of the last deployment of a stage. Supported only for
	// stages with autoDeploy enabled.
	LastDeploymentStatusMessage *string

	// The timestamp when the stage was last updated.
	LastUpdatedDate *time.Time

	// Route settings for the stage, by routeKey.
	RouteSettings map[string]types.RouteSettings

	// The name of the stage.
	StageName *string

	// A map that defines the stage variables for a stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	StageVariables map[string]string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetStage"); err != nil {
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
	if err = addOpGetStageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetStage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetStage",
	}
}
