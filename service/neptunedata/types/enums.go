// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Action string

// Enum values for Action
const (
	ActionInitializeReset Action = "initiateDatabaseReset"
	ActionPerformReset    Action = "performDatabaseReset"
)

// Values returns all known values for Action. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Action) Values() []Action {
	return []Action{
		"initiateDatabaseReset",
		"performDatabaseReset",
	}
}

type Encoding string

// Enum values for Encoding
const (
	EncodingGzip Encoding = "gzip"
)

// Values returns all known values for Encoding. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Encoding) Values() []Encoding {
	return []Encoding{
		"gzip",
	}
}

type Format string

// Enum values for Format
const (
	FormatCsv        Format = "csv"
	FormatOpencypher Format = "opencypher"
	FormatNtriples   Format = "ntriples"
	FormatNquads     Format = "nquads"
	FormatRdfxml     Format = "rdfxml"
	FormatTurtle     Format = "turtle"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"csv",
		"opencypher",
		"ntriples",
		"nquads",
		"rdfxml",
		"turtle",
	}
}

type GraphSummaryType string

// Enum values for GraphSummaryType
const (
	GraphSummaryTypeBasic    GraphSummaryType = "basic"
	GraphSummaryTypeDetailed GraphSummaryType = "detailed"
)

// Values returns all known values for GraphSummaryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GraphSummaryType) Values() []GraphSummaryType {
	return []GraphSummaryType{
		"basic",
		"detailed",
	}
}

type IteratorType string

// Enum values for IteratorType
const (
	IteratorTypeAtSequenceNumber    IteratorType = "AT_SEQUENCE_NUMBER"
	IteratorTypeAfterSequenceNumber IteratorType = "AFTER_SEQUENCE_NUMBER"
	IteratorTypeTrimHorizon         IteratorType = "TRIM_HORIZON"
	IteratorTypeLatest              IteratorType = "LATEST"
)

// Values returns all known values for IteratorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IteratorType) Values() []IteratorType {
	return []IteratorType{
		"AT_SEQUENCE_NUMBER",
		"AFTER_SEQUENCE_NUMBER",
		"TRIM_HORIZON",
		"LATEST",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeResume Mode = "RESUME"
	ModeNew    Mode = "NEW"
	ModeAuto   Mode = "AUTO"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"RESUME",
		"NEW",
		"AUTO",
	}
}

type OpenCypherExplainMode string

// Enum values for OpenCypherExplainMode
const (
	OpenCypherExplainModeStatic  OpenCypherExplainMode = "static"
	OpenCypherExplainModeDynamic OpenCypherExplainMode = "dynamic"
	OpenCypherExplainModeDetails OpenCypherExplainMode = "details"
)

// Values returns all known values for OpenCypherExplainMode. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OpenCypherExplainMode) Values() []OpenCypherExplainMode {
	return []OpenCypherExplainMode{
		"static",
		"dynamic",
		"details",
	}
}

type Parallelism string

// Enum values for Parallelism
const (
	ParallelismLow           Parallelism = "LOW"
	ParallelismMedium        Parallelism = "MEDIUM"
	ParallelismHigh          Parallelism = "HIGH"
	ParallelismOversubscribe Parallelism = "OVERSUBSCRIBE"
)

// Values returns all known values for Parallelism. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Parallelism) Values() []Parallelism {
	return []Parallelism{
		"LOW",
		"MEDIUM",
		"HIGH",
		"OVERSUBSCRIBE",
	}
}

type S3BucketRegion string

// Enum values for S3BucketRegion
const (
	S3BucketRegionUsEast1      S3BucketRegion = "us-east-1"
	S3BucketRegionUsEast2      S3BucketRegion = "us-east-2"
	S3BucketRegionUsWest1      S3BucketRegion = "us-west-1"
	S3BucketRegionUsWest2      S3BucketRegion = "us-west-2"
	S3BucketRegionCaCentral1   S3BucketRegion = "ca-central-1"
	S3BucketRegionSaEast1      S3BucketRegion = "sa-east-1"
	S3BucketRegionEuNorth1     S3BucketRegion = "eu-north-1"
	S3BucketRegionEuWest1      S3BucketRegion = "eu-west-1"
	S3BucketRegionEuWest2      S3BucketRegion = "eu-west-2"
	S3BucketRegionEuWest3      S3BucketRegion = "eu-west-3"
	S3BucketRegionEuCentral1   S3BucketRegion = "eu-central-1"
	S3BucketRegionMeSouth1     S3BucketRegion = "me-south-1"
	S3BucketRegionAfSouth1     S3BucketRegion = "af-south-1"
	S3BucketRegionApEast1      S3BucketRegion = "ap-east-1"
	S3BucketRegionApNortheast1 S3BucketRegion = "ap-northeast-1"
	S3BucketRegionApNortheast2 S3BucketRegion = "ap-northeast-2"
	S3BucketRegionApSoutheast1 S3BucketRegion = "ap-southeast-1"
	S3BucketRegionApSoutheast2 S3BucketRegion = "ap-southeast-2"
	S3BucketRegionApSouth1     S3BucketRegion = "ap-south-1"
	S3BucketRegionCnNorth1     S3BucketRegion = "cn-north-1"
	S3BucketRegionCnNorthwest1 S3BucketRegion = "cn-northwest-1"
	S3BucketRegionUsGovWest1   S3BucketRegion = "us-gov-west-1"
	S3BucketRegionUsGovEast1   S3BucketRegion = "us-gov-east-1"
	S3BucketRegionCaWest1      S3BucketRegion = "ca-west-1"
	S3BucketRegionEuSouth2     S3BucketRegion = "eu-south-2"
	S3BucketRegionIlCentral1   S3BucketRegion = "il-central-1"
	S3BucketRegionMeCentral1   S3BucketRegion = "me-central-1"
	S3BucketRegionApNortheast3 S3BucketRegion = "ap-northeast-3"
	S3BucketRegionApSoutheast3 S3BucketRegion = "ap-southeast-3"
	S3BucketRegionApSoutheast4 S3BucketRegion = "ap-southeast-4"
	S3BucketRegionApSoutheast5 S3BucketRegion = "ap-southeast-5"
	S3BucketRegionApSoutheast7 S3BucketRegion = "ap-southeast-7"
	S3BucketRegionMxCentral1   S3BucketRegion = "mx-central-1"
	S3BucketRegionApEast2      S3BucketRegion = "ap-east-2"
	S3BucketRegionApSouth2     S3BucketRegion = "ap-south-2"
	S3BucketRegionEuCentral2   S3BucketRegion = "eu-central-2"
)

// Values returns all known values for S3BucketRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (S3BucketRegion) Values() []S3BucketRegion {
	return []S3BucketRegion{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ca-central-1",
		"sa-east-1",
		"eu-north-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-central-1",
		"me-south-1",
		"af-south-1",
		"ap-east-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-south-1",
		"cn-north-1",
		"cn-northwest-1",
		"us-gov-west-1",
		"us-gov-east-1",
		"ca-west-1",
		"eu-south-2",
		"il-central-1",
		"me-central-1",
		"ap-northeast-3",
		"ap-southeast-3",
		"ap-southeast-4",
		"ap-southeast-5",
		"ap-southeast-7",
		"mx-central-1",
		"ap-east-2",
		"ap-south-2",
		"eu-central-2",
	}
}

type StatisticsAutoGenerationMode string

// Enum values for StatisticsAutoGenerationMode
const (
	StatisticsAutoGenerationModeDisableAutocompute StatisticsAutoGenerationMode = "disableAutoCompute"
	StatisticsAutoGenerationModeEnableAutocompute  StatisticsAutoGenerationMode = "enableAutoCompute"
	StatisticsAutoGenerationModeRefresh            StatisticsAutoGenerationMode = "refresh"
)

// Values returns all known values for StatisticsAutoGenerationMode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StatisticsAutoGenerationMode) Values() []StatisticsAutoGenerationMode {
	return []StatisticsAutoGenerationMode{
		"disableAutoCompute",
		"enableAutoCompute",
		"refresh",
	}
}
