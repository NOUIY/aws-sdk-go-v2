module github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue

go 1.22

require (
	github.com/aws/aws-sdk-go-v2 v1.36.6
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.44.1
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.26.1
)

require github.com/aws/smithy-go v1.22.4

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/aws/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
