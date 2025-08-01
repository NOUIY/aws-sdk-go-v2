// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a trained model in a collaboration.
func (c *Client) GetCollaborationTrainedModel(ctx context.Context, params *GetCollaborationTrainedModelInput, optFns ...func(*Options)) (*GetCollaborationTrainedModelOutput, error) {
	if params == nil {
		params = &GetCollaborationTrainedModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCollaborationTrainedModel", params, optFns, c.addOperationGetCollaborationTrainedModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCollaborationTrainedModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCollaborationTrainedModelInput struct {

	// The collaboration ID that contains the trained model that you want to return
	// information about.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the trained model that you want to return
	// information about.
	//
	// This member is required.
	TrainedModelArn *string

	// The version identifier of the trained model to retrieve. If not specified, the
	// operation returns information about the latest version of the trained model.
	VersionIdentifier *string

	noSmithyDocumentSerde
}

type GetCollaborationTrainedModelOutput struct {

	// The collaboration ID of the collaboration that contains the trained model.
	//
	// This member is required.
	CollaborationIdentifier *string

	// The Amazon Resource Name (ARN) of the configured model algorithm association
	// that was used to create this trained model.
	//
	// This member is required.
	ConfiguredModelAlgorithmAssociationArn *string

	// The time at which the trained model was created.
	//
	// This member is required.
	CreateTime *time.Time

	// The account ID of the member that created the trained model.
	//
	// This member is required.
	CreatorAccountId *string

	// The membership ID of the member that created the trained model.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the trained model.
	//
	// This member is required.
	Name *string

	// The status of the trained model.
	//
	// This member is required.
	Status types.TrainedModelStatus

	// The Amazon Resource Name (ARN) of the trained model.
	//
	// This member is required.
	TrainedModelArn *string

	// The most recent time at which the trained model was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The description of the trained model.
	Description *string

	// Information about the incremental training data channels used to create this
	// version of the trained model. This includes details about the base model that
	// was used for incremental training and the channel configuration.
	IncrementalTrainingDataChannels []types.IncrementalTrainingDataChannelOutput

	// Status information for the logs.
	LogsStatus types.LogsStatus

	// Details about the status information for the logs.
	LogsStatusDetails *string

	// The status of the model metrics.
	MetricsStatus types.MetricsStatus

	// Details about the status information for the model metrics.
	MetricsStatusDetails *string

	// The EC2 resource configuration that was used to train this model.
	ResourceConfig *types.ResourceConfig

	// Details about the status of a resource.
	StatusDetails *types.StatusDetails

	// The stopping condition that determined when model training ended.
	StoppingCondition *types.StoppingCondition

	// Information about the training container image.
	TrainingContainerImageDigest *string

	// The input mode that was used for accessing the training data when this trained
	// model was created. This indicates how the training data was made available to
	// the training algorithm.
	TrainingInputMode types.TrainingInputMode

	// The version identifier of the trained model. This unique identifier
	// distinguishes this version from other versions of the same trained model.
	VersionIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCollaborationTrainedModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCollaborationTrainedModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCollaborationTrainedModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCollaborationTrainedModel"); err != nil {
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
	if err = addOpGetCollaborationTrainedModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCollaborationTrainedModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCollaborationTrainedModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCollaborationTrainedModel",
	}
}
