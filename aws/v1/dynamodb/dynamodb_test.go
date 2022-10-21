package dynamodb

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func TestDynamoDB_PutItemWithContext(t *testing.T) {

	tracer.new
	type fields struct {
		client dynamodbiface.DynamoDBAPI
	}
	type args struct {
		ctx    aws.Context
		input  *dynamodb.PutItemInput
		option []request.Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dynamodb.PutItemOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamoDB{
				client: tt.fields.client,
			}
			got, err := d.PutItemWithContext(tt.args.ctx, tt.args.input, tt.args.option...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutItemWithContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutItemWithContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}
