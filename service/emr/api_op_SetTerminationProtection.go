// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// SetTerminationProtection locks a cluster (job flow) so the Amazon EC2 instances
// in the cluster cannot be terminated by user intervention, an API call, or in the
// event of a job-flow error. The cluster still terminates upon successful
// completion of the job flow. Calling SetTerminationProtection on a cluster is
// similar to calling the Amazon EC2 DisableAPITermination API on all Amazon EC2
// instances in a cluster.
//
// SetTerminationProtection is used to prevent accidental termination of a cluster
// and to ensure that in the event of an error, the instances persist so that you
// can recover any data stored in their ephemeral instance storage.
//
// To terminate a cluster that has been locked by setting SetTerminationProtection
// to true , you must first unlock the job flow by a subsequent call to
// SetTerminationProtection in which you set the value to false .
//
// For more information, see [Managing Cluster Termination] in the Amazon EMR Management Guide.
//
// [Managing Cluster Termination]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_TerminationProtection.html
func (c *Client) SetTerminationProtection(ctx context.Context, params *SetTerminationProtectionInput, optFns ...func(*Options)) (*SetTerminationProtectionOutput, error) {
	if params == nil {
		params = &SetTerminationProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTerminationProtection", params, optFns, c.addOperationSetTerminationProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTerminationProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input argument to the TerminationProtection operation.
type SetTerminationProtectionInput struct {

	//  A list of strings that uniquely identify the clusters to protect. This
	// identifier is returned by RunJobFlowand can also be obtained from DescribeJobFlows .
	//
	// This member is required.
	JobFlowIds []string

	// A Boolean that indicates whether to protect the cluster and prevent the Amazon
	// EC2 instances in the cluster from shutting down due to API calls, user
	// intervention, or job-flow error.
	//
	// This member is required.
	TerminationProtected *bool

	noSmithyDocumentSerde
}

type SetTerminationProtectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTerminationProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetTerminationProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTerminationProtection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetTerminationProtection"); err != nil {
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
	if err = addOpSetTerminationProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTerminationProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTerminationProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetTerminationProtection",
	}
}
