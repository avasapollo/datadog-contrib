package dynamodb

import (
	"context"
	"testing"

	"github.com/avasapollo/datadog-contrib/aws/v1/dynamodb/mocks"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func TestDynamoDB_PutItemWithContext(t *testing.T) {
	tracer.Start(tracer.WithDebugMode(true))
	defer tracer.Stop()

	ctrl := gomock.NewController(t)
	mockDBApi := mocks.NewMockDynamoDBAPI(ctrl)
	mockDBApi.EXPECT().PutItemWithContext(gomock.Any(), gomock.Any()).Return(nil, nil)
	d := New(mockDBApi)
	input := &dynamodb.PutItemInput{}
	_, err := d.PutItemWithContext(context.Background(), input)
	require.NoError(t, err)
}
