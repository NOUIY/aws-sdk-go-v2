// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels the transfer of a domain from the current Amazon Web Services account
// to another Amazon Web Services account. You initiate a transfer betweenAmazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// You must cancel the transfer before the other Amazon Web Services account
// accepts the transfer using [AcceptDomainTransferFromAnotherAwsAccount].
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func (c *Client) CancelDomainTransferToAnotherAwsAccount(ctx context.Context, params *CancelDomainTransferToAnotherAwsAccountInput, optFns ...func(*Options)) (*CancelDomainTransferToAnotherAwsAccountOutput, error) {
	if params == nil {
		params = &CancelDomainTransferToAnotherAwsAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelDomainTransferToAnotherAwsAccount", params, optFns, c.addOperationCancelDomainTransferToAnotherAwsAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelDomainTransferToAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CancelDomainTransferToAnotherAwsAccount request includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountInput struct {

	// The name of the domain for which you want to cancel the transfer to another
	// Amazon Web Services account.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The CancelDomainTransferToAnotherAwsAccount response includes the following
// element.
type CancelDomainTransferToAnotherAwsAccountOutput struct {

	// The identifier that TransferDomainToAnotherAwsAccount returned to track the
	// progress of the request. Because the transfer request was canceled, the value is
	// no longer valid, and you can't use GetOperationDetail to query the operation
	// status.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelDomainTransferToAnotherAwsAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelDomainTransferToAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelDomainTransferToAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelDomainTransferToAnotherAwsAccount"); err != nil {
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
	if err = addOpCancelDomainTransferToAnotherAwsAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelDomainTransferToAnotherAwsAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelDomainTransferToAnotherAwsAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelDomainTransferToAnotherAwsAccount",
	}
}
