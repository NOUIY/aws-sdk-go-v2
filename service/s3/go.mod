module github.com/aws/aws-sdk-go-v2/service/s3

go 1.22

require (
	github.com/aws/aws-sdk-go-v2 v1.37.1
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.1
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.1
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.1
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.0
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.8.1
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.1
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.1
	github.com/aws/smithy-go v1.22.5
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/internal/v4a => ../../internal/v4a/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/checksum => ../../service/internal/checksum/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
