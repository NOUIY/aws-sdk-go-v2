// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe and the transcription results are streamed to your
// application.
//
// The following parameters are required:
//
//   - language-code or identify-language or identify-multiple-language
//
//   - media-encoding
//
//   - sample-rate
//
// For more information on streaming with Amazon Transcribe, see [Transcribing streaming audio].
//
// [Transcribing streaming audio]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html
func (c *Client) StartStreamTranscription(ctx context.Context, params *StartStreamTranscriptionInput, optFns ...func(*Options)) (*StartStreamTranscriptionOutput, error) {
	if params == nil {
		params = &StartStreamTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartStreamTranscription", params, optFns, c.addOperationStartStreamTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartStreamTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamTranscriptionInput struct {

	// Specify the encoding of your input audio. Supported formats are:
	//
	//   - FLAC
	//
	//   - OPUS-encoded audio in an Ogg container
	//
	//   - PCM (only signed 16-bit little-endian audio formats, which does not include
	//   WAV)
	//
	// For more information, see [Media formats].
	//
	// [Media formats]: https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio
	//
	// This member is required.
	MediaEncoding types.MediaEncoding

	// The sample rate of the input audio (in hertz). Low-quality audio, such as
	// telephone audio, is typically around 8,000 Hz. High-quality audio typically
	// ranges from 16,000 Hz to 48,000 Hz. Note that the sample rate you specify must
	// match that of your audio.
	//
	// This member is required.
	MediaSampleRateHertz *int32

	// Labels all personally identifiable information (PII) identified in your
	// transcript.
	//
	// Content identification is performed at the segment level; PII specified in
	// PiiEntityTypes is flagged upon complete transcription of an audio segment. If
	// you don't include PiiEntityTypes in your request, all PII is identified.
	//
	// You can’t set ContentIdentificationType and ContentRedactionType in the same
	// request. If you set both, your request returns a BadRequestException .
	//
	// For more information, see [Redacting or identifying personally identifiable information].
	//
	// [Redacting or identifying personally identifiable information]: https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html
	ContentIdentificationType types.ContentIdentificationType

	// Redacts all personally identifiable information (PII) identified in your
	// transcript.
	//
	// Content redaction is performed at the segment level; PII specified in
	// PiiEntityTypes is redacted upon complete transcription of an audio segment. If
	// you don't include PiiEntityTypes in your request, all PII is redacted.
	//
	// You can’t set ContentRedactionType and ContentIdentificationType in the same
	// request. If you set both, your request returns a BadRequestException .
	//
	// For more information, see [Redacting or identifying personally identifiable information].
	//
	// [Redacting or identifying personally identifiable information]: https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html
	ContentRedactionType types.ContentRedactionType

	// Enables channel identification in multi-channel audio.
	//
	// Channel identification transcribes the audio on each channel independently,
	// then appends the output for each channel into one transcript.
	//
	// If you have multi-channel audio and do not enable channel identification, your
	// audio is transcribed in a continuous manner and your transcript is not separated
	// by channel.
	//
	// If you include EnableChannelIdentification in your request, you must also
	// include NumberOfChannels .
	//
	// For more information, see [Transcribing multi-channel audio].
	//
	// [Transcribing multi-channel audio]: https://docs.aws.amazon.com/transcribe/latest/dg/channel-id.html
	EnableChannelIdentification bool

	// Enables partial result stabilization for your transcription. Partial result
	// stabilization can reduce latency in your output, but may impact accuracy. For
	// more information, see [Partial-result stabilization].
	//
	// [Partial-result stabilization]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization
	EnablePartialResultsStabilization bool

	// Enables automatic language identification for your transcription.
	//
	// If you include IdentifyLanguage , you must include a list of language codes,
	// using LanguageOptions , that you think may be present in your audio stream.
	//
	// You can also include a preferred language using PreferredLanguage . Adding a
	// preferred language can help Amazon Transcribe identify the language faster than
	// if you omit this parameter.
	//
	// If you have multi-channel audio that contains different languages on each
	// channel, and you've enabled channel identification, automatic language
	// identification identifies the dominant language on each audio channel.
	//
	// Note that you must include either LanguageCode or IdentifyLanguage or
	// IdentifyMultipleLanguages in your request. If you include more than one of these
	// parameters, your transcription job fails.
	//
	// Streaming language identification can't be combined with custom language models
	// or redaction.
	IdentifyLanguage bool

	// Enables automatic multi-language identification in your transcription job
	// request. Use this parameter if your stream contains more than one language. If
	// your stream contains only one language, use IdentifyLanguage instead.
	//
	// If you include IdentifyMultipleLanguages , you must include a list of language
	// codes, using LanguageOptions , that you think may be present in your stream.
	//
	// If you want to apply a custom vocabulary or a custom vocabulary filter to your
	// automatic multiple language identification request, include VocabularyNames or
	// VocabularyFilterNames .
	//
	// Note that you must include one of LanguageCode , IdentifyLanguage , or
	// IdentifyMultipleLanguages in your request. If you include more than one of these
	// parameters, your transcription job fails.
	IdentifyMultipleLanguages bool

	// Specify the language code that represents the language spoken in your audio.
	//
	// If you're unsure of the language spoken in your audio, consider using
	// IdentifyLanguage to enable automatic language identification.
	//
	// For a list of languages supported with Amazon Transcribe streaming, refer to
	// the [Supported languages]table.
	//
	// [Supported languages]: https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html
	LanguageCode types.LanguageCode

	// Specify the name of the custom language model that you want to use when
	// processing your transcription. Note that language model names are case
	// sensitive.
	//
	// The language of the specified language model must match the language code you
	// specify in your transcription request. If the languages don't match, the custom
	// language model isn't applied. There are no errors or warnings associated with a
	// language mismatch.
	//
	// For more information, see [Custom language models].
	//
	// [Custom language models]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-language-models.html
	LanguageModelName *string

	// Specify two or more language codes that represent the languages you think may
	// be present in your media; including more than five is not recommended.
	//
	// Including language options can improve the accuracy of language identification.
	//
	// If you include LanguageOptions in your request, you must also include
	// IdentifyLanguage or IdentifyMultipleLanguages .
	//
	// For a list of languages supported with Amazon Transcribe streaming, refer to
	// the [Supported languages]table.
	//
	// You can only include one language dialect per language per stream. For example,
	// you cannot include en-US and en-AU in the same request.
	//
	// [Supported languages]: https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html
	LanguageOptions *string

	// Specify the number of channels in your audio stream. This value must be 2 , as
	// only two channels are supported. If your audio doesn't contain multiple
	// channels, do not include this parameter in your request.
	//
	// If you include NumberOfChannels in your request, you must also include
	// EnableChannelIdentification .
	NumberOfChannels *int32

	// Specify the level of stability to use when you enable partial results
	// stabilization ( EnablePartialResultsStabilization ).
	//
	// Low stability provides the highest accuracy. High stability transcribes faster,
	// but with slightly lower accuracy.
	//
	// For more information, see [Partial-result stabilization].
	//
	// [Partial-result stabilization]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization
	PartialResultsStability types.PartialResultsStability

	// Specify which types of personally identifiable information (PII) you want to
	// redact in your transcript. You can include as many types as you'd like, or you
	// can select ALL .
	//
	// Values must be comma-separated and can include: ADDRESS , BANK_ACCOUNT_NUMBER ,
	// BANK_ROUTING , CREDIT_DEBIT_CVV , CREDIT_DEBIT_EXPIRY , CREDIT_DEBIT_NUMBER ,
	// EMAIL , NAME , PHONE , PIN , SSN , or ALL .
	//
	// Note that if you include PiiEntityTypes in your request, you must also include
	// ContentIdentificationType or ContentRedactionType .
	//
	// If you include ContentRedactionType or ContentIdentificationType in your
	// request, but do not include PiiEntityTypes , all PII is redacted or identified.
	PiiEntityTypes *string

	// Specify a preferred language from the subset of languages codes you specified
	// in LanguageOptions .
	//
	// You can only use this parameter if you've included IdentifyLanguage and
	// LanguageOptions in your request.
	PreferredLanguage types.LanguageCode

	// Specify a name for your transcription session. If you don't include this
	// parameter in your request, Amazon Transcribe generates an ID and returns it in
	// the response.
	SessionId *string

	// Enables speaker partitioning (diarization) in your transcription output.
	// Speaker partitioning labels the speech from individual speakers in your media
	// file.
	//
	// For more information, see [Partitioning speakers (diarization)].
	//
	// [Partitioning speakers (diarization)]: https://docs.aws.amazon.com/transcribe/latest/dg/diarization.html
	ShowSpeakerLabel bool

	// Specify how you want your vocabulary filter applied to your transcript.
	//
	// To replace words with *** , choose mask .
	//
	// To delete words, choose remove .
	//
	// To flag words without changing them, choose tag .
	VocabularyFilterMethod types.VocabularyFilterMethod

	// Specify the name of the custom vocabulary filter that you want to use when
	// processing your transcription. Note that vocabulary filter names are case
	// sensitive.
	//
	// If the language of the specified custom vocabulary filter doesn't match the
	// language identified in your media, the vocabulary filter is not applied to your
	// transcription.
	//
	// This parameter is not intended for use with the IdentifyLanguage parameter. If
	// you're including IdentifyLanguage in your request and want to use one or more
	// vocabulary filters with your transcription, use the VocabularyFilterNames
	// parameter instead.
	//
	// For more information, see [Using vocabulary filtering with unwanted words].
	//
	// [Using vocabulary filtering with unwanted words]: https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html
	VocabularyFilterName *string

	// Specify the names of the custom vocabulary filters that you want to use when
	// processing your transcription. Note that vocabulary filter names are case
	// sensitive.
	//
	// If none of the languages of the specified custom vocabulary filters match the
	// language identified in your media, your job fails.
	//
	// This parameter is only intended for use with the IdentifyLanguage parameter. If
	// you're not including IdentifyLanguage in your request and want to use a custom
	// vocabulary filter with your transcription, use the VocabularyFilterName
	// parameter instead.
	//
	// For more information, see [Using vocabulary filtering with unwanted words].
	//
	// [Using vocabulary filtering with unwanted words]: https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html
	VocabularyFilterNames *string

	// Specify the name of the custom vocabulary that you want to use when processing
	// your transcription. Note that vocabulary names are case sensitive.
	//
	// If the language of the specified custom vocabulary doesn't match the language
	// identified in your media, the custom vocabulary is not applied to your
	// transcription.
	//
	// This parameter is not intended for use with the IdentifyLanguage parameter. If
	// you're including IdentifyLanguage in your request and want to use one or more
	// custom vocabularies with your transcription, use the VocabularyNames parameter
	// instead.
	//
	// For more information, see [Custom vocabularies].
	//
	// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
	VocabularyName *string

	// Specify the names of the custom vocabularies that you want to use when
	// processing your transcription. Note that vocabulary names are case sensitive.
	//
	// If none of the languages of the specified custom vocabularies match the
	// language identified in your media, your job fails.
	//
	// This parameter is only intended for use with the IdentifyLanguage parameter. If
	// you're not including IdentifyLanguage in your request and want to use a custom
	// vocabulary with your transcription, use the VocabularyName parameter instead.
	//
	// For more information, see [Custom vocabularies].
	//
	// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
	VocabularyNames *string

	noSmithyDocumentSerde
}

type StartStreamTranscriptionOutput struct {

	// Shows whether content identification was enabled for your transcription.
	ContentIdentificationType types.ContentIdentificationType

	// Shows whether content redaction was enabled for your transcription.
	ContentRedactionType types.ContentRedactionType

	// Shows whether channel identification was enabled for your transcription.
	EnableChannelIdentification bool

	// Shows whether partial results stabilization was enabled for your transcription.
	EnablePartialResultsStabilization bool

	// Shows whether automatic language identification was enabled for your
	// transcription.
	IdentifyLanguage bool

	// Shows whether automatic multi-language identification was enabled for your
	// transcription.
	IdentifyMultipleLanguages bool

	// Provides the language code that you specified in your request.
	LanguageCode types.LanguageCode

	// Provides the name of the custom language model that you specified in your
	// request.
	LanguageModelName *string

	// Provides the language codes that you specified in your request.
	LanguageOptions *string

	// Provides the media encoding you specified in your request.
	MediaEncoding types.MediaEncoding

	// Provides the sample rate that you specified in your request.
	MediaSampleRateHertz *int32

	// Provides the number of channels that you specified in your request.
	NumberOfChannels *int32

	// Provides the stabilization level used for your transcription.
	PartialResultsStability types.PartialResultsStability

	// Lists the PII entity types you specified in your request.
	PiiEntityTypes *string

	// Provides the preferred language that you specified in your request.
	PreferredLanguage types.LanguageCode

	// Provides the identifier for your streaming request.
	RequestId *string

	// Provides the identifier for your transcription session.
	SessionId *string

	// Shows whether speaker partitioning was enabled for your transcription.
	ShowSpeakerLabel bool

	// Provides the vocabulary filtering method used in your transcription.
	VocabularyFilterMethod types.VocabularyFilterMethod

	// Provides the name of the custom vocabulary filter that you specified in your
	// request.
	VocabularyFilterName *string

	// Provides the names of the custom vocabulary filters that you specified in your
	// request.
	VocabularyFilterNames *string

	// Provides the name of the custom vocabulary that you specified in your request.
	VocabularyName *string

	// Provides the names of the custom vocabularies that you specified in your
	// request.
	VocabularyNames *string

	eventStream *StartStreamTranscriptionEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartStreamTranscriptionOutput) GetStream() *StartStreamTranscriptionEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartStreamTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartStreamTranscription"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addEventStreamStartStreamTranscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = addContentSHA256Header(stack); err != nil {
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
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
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
	if err = addOpStartStreamTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartStreamTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartStreamTranscription",
	}
}

// StartStreamTranscriptionEventStream provides the event stream handling for the StartStreamTranscription operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartStreamTranscriptionEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartStreamTranscriptionEventStream struct {
	// AudioStreamWriter is the EventStream writer for the AudioStream events. This
	// value is automatically set by the SDK when the API call is made Use this member
	// when unit testing your code with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer AudioStreamWriter

	// TranscriptResultStreamReader is the EventStream reader for the
	// TranscriptResultStream events. This value is automatically set by the SDK when
	// the API call is made Use this member when unit testing your code with the SDK to
	// mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader TranscriptResultStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartStreamTranscriptionEventStream initializes an StartStreamTranscriptionEventStream.
// This function should only be used for testing and mocking the StartStreamTranscriptionEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartStreamTranscriptionEventStream(optFns ...func(*StartStreamTranscriptionEventStream)) *StartStreamTranscriptionEventStream {
	es := &StartStreamTranscriptionEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartStreamTranscriptionEventStream) Send(ctx context.Context, event types.AudioStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartStreamTranscriptionEventStream) Events() <-chan types.TranscriptResultStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartStreamTranscriptionEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartStreamTranscriptionEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartStreamTranscriptionEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartStreamTranscriptionEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
