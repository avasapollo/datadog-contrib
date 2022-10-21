package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// assertion the interface Payer to MongoPayer
var _ dynamodbiface.DynamoDBAPI = (*DynamoDB)(nil)

type config struct {
	ServiceName string `envconfig:"DD_SERVICE"`
}

type Option func(opts *config)

type DynamoDB struct {
	config *config
	client dynamodbiface.DynamoDBAPI
}

func New(client dynamodbiface.DynamoDBAPI, opts ...Option) *DynamoDB {
	c := new(config)
	_ = envconfig.Process("", c)
	for _, opt := range opts {
		opt(c)
	}
	return &DynamoDB{
		config: c,
		client: client,
	}
}

func (d *DynamoDB) BatchExecuteStatement(input *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {

	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchExecuteStatementWithContext(context aws.Context, input *dynamodb.BatchExecuteStatementInput, option ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchExecuteStatementRequest(input *dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchGetItemWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, option ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchGetItemPages(input *dynamodb.BatchGetItemInput, f func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchGetItemPagesWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, f func(*dynamodb.BatchGetItemOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchWriteItemWithContext(context aws.Context, input *dynamodb.BatchWriteItemInput, option ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateBackup(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateBackupWithContext(context aws.Context, input *dynamodb.CreateBackupInput, option ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateBackupRequest(input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateGlobalTable(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateGlobalTableWithContext(context aws.Context, input *dynamodb.CreateGlobalTableInput, option ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateGlobalTableRequest(input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateTableWithContext(context aws.Context, input *dynamodb.CreateTableInput, option ...request.Option) (*dynamodb.CreateTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) CreateTableRequest(input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteBackup(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteBackupWithContext(context aws.Context, input *dynamodb.DeleteBackupInput, option ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteBackupRequest(input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteItemWithContext(context aws.Context, input *dynamodb.DeleteItemInput, option ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteTable(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteTableWithContext(context aws.Context, input *dynamodb.DeleteTableInput, option ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DeleteTableRequest(input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeBackup(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeBackupWithContext(context aws.Context, input *dynamodb.DescribeBackupInput, option ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeBackupRequest(input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContinuousBackups(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContinuousBackupsWithContext(context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, option ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContinuousBackupsRequest(input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContributorInsights(input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContributorInsightsWithContext(context aws.Context, input *dynamodb.DescribeContributorInsightsInput, option ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeContributorInsightsRequest(input *dynamodb.DescribeContributorInsightsInput) (*request.Request, *dynamodb.DescribeContributorInsightsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeEndpoints(input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeEndpointsWithContext(context aws.Context, input *dynamodb.DescribeEndpointsInput, option ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeEndpointsRequest(input *dynamodb.DescribeEndpointsInput) (*request.Request, *dynamodb.DescribeEndpointsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeExport(input *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeExportWithContext(context aws.Context, input *dynamodb.DescribeExportInput, option ...request.Option) (*dynamodb.DescribeExportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeExportRequest(input *dynamodb.DescribeExportInput) (*request.Request, *dynamodb.DescribeExportOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTable(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTableWithContext(context aws.Context, input *dynamodb.DescribeGlobalTableInput, option ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTableRequest(input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTableSettings(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTableSettingsWithContext(context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, option ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeGlobalTableSettingsRequest(input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeImport(input *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeImportWithContext(context aws.Context, input *dynamodb.DescribeImportInput, option ...request.Option) (*dynamodb.DescribeImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeImportRequest(input *dynamodb.DescribeImportInput) (*request.Request, *dynamodb.DescribeImportOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeKinesisStreamingDestination(input *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeKinesisStreamingDestinationWithContext(context aws.Context, input *dynamodb.DescribeKinesisStreamingDestinationInput, option ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeKinesisStreamingDestinationRequest(input *dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DescribeKinesisStreamingDestinationOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeLimitsWithContext(context aws.Context, input *dynamodb.DescribeLimitsInput, option ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeLimitsRequest(input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTableWithContext(context aws.Context, input *dynamodb.DescribeTableInput, option ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTableRequest(input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTableReplicaAutoScaling(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTableReplicaAutoScalingWithContext(context aws.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput, option ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTableReplicaAutoScalingRequest(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request, *dynamodb.DescribeTableReplicaAutoScalingOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTimeToLiveWithContext(context aws.Context, input *dynamodb.DescribeTimeToLiveInput, option ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DescribeTimeToLiveRequest(input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DisableKinesisStreamingDestination(input *dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DisableKinesisStreamingDestinationWithContext(context aws.Context, input *dynamodb.DisableKinesisStreamingDestinationInput, option ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) DisableKinesisStreamingDestinationRequest(input *dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) EnableKinesisStreamingDestination(input *dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) EnableKinesisStreamingDestinationWithContext(context aws.Context, input *dynamodb.EnableKinesisStreamingDestinationInput, option ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) EnableKinesisStreamingDestinationRequest(input *dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteStatement(input *dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteStatementWithContext(context aws.Context, input *dynamodb.ExecuteStatementInput, option ...request.Option) (*dynamodb.ExecuteStatementOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteStatementRequest(input *dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteTransaction(input *dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteTransactionWithContext(context aws.Context, input *dynamodb.ExecuteTransactionInput, option ...request.Option) (*dynamodb.ExecuteTransactionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExecuteTransactionRequest(input *dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExportTableToPointInTime(input *dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExportTableToPointInTimeWithContext(context aws.Context, input *dynamodb.ExportTableToPointInTimeInput, option ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ExportTableToPointInTimeRequest(input *dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) GetItemWithContext(context aws.Context, input *dynamodb.GetItemInput, option ...request.Option) (*dynamodb.GetItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ImportTable(input *dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ImportTableWithContext(context aws.Context, input *dynamodb.ImportTableInput, option ...request.Option) (*dynamodb.ImportTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ImportTableRequest(input *dynamodb.ImportTableInput) (*request.Request, *dynamodb.ImportTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListBackups(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListBackupsWithContext(context aws.Context, input *dynamodb.ListBackupsInput, option ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListBackupsRequest(input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListContributorInsights(input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListContributorInsightsWithContext(context aws.Context, input *dynamodb.ListContributorInsightsInput, option ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListContributorInsightsRequest(input *dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListContributorInsightsPages(input *dynamodb.ListContributorInsightsInput, f func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListContributorInsightsPagesWithContext(context aws.Context, input *dynamodb.ListContributorInsightsInput, f func(*dynamodb.ListContributorInsightsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListExports(input *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListExportsWithContext(context aws.Context, input *dynamodb.ListExportsInput, option ...request.Option) (*dynamodb.ListExportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListExportsRequest(input *dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListExportsPages(input *dynamodb.ListExportsInput, f func(*dynamodb.ListExportsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListExportsPagesWithContext(context aws.Context, input *dynamodb.ListExportsInput, f func(*dynamodb.ListExportsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListGlobalTables(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListGlobalTablesWithContext(context aws.Context, input *dynamodb.ListGlobalTablesInput, option ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListGlobalTablesRequest(input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListImports(input *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListImportsWithContext(context aws.Context, input *dynamodb.ListImportsInput, option ...request.Option) (*dynamodb.ListImportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListImportsRequest(input *dynamodb.ListImportsInput) (*request.Request, *dynamodb.ListImportsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListImportsPages(input *dynamodb.ListImportsInput, f func(*dynamodb.ListImportsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListImportsPagesWithContext(context aws.Context, input *dynamodb.ListImportsInput, f func(*dynamodb.ListImportsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTablesWithContext(context aws.Context, input *dynamodb.ListTablesInput, option ...request.Option) (*dynamodb.ListTablesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTablesRequest(input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTablesPages(input *dynamodb.ListTablesInput, f func(*dynamodb.ListTablesOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTablesPagesWithContext(context aws.Context, input *dynamodb.ListTablesInput, f func(*dynamodb.ListTablesOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTagsOfResourceWithContext(context aws.Context, input *dynamodb.ListTagsOfResourceInput, option ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ListTagsOfResourceRequest(input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.PutItemWithContext(context.Background(), input)
}

type tagsParams struct {
	command string
	table   *string
}

func (p *tagsParams) Table() string {
	if p.table != nil {
		return *p.table
	}
	return ""
}

func (d *DynamoDB) setTags(span tracer.Span, params *tagsParams) tracer.Span {
	span.SetTag(ext.ServiceName, d.config.ServiceName)
	span.SetTag(ext.DBType, "nosql")
	span.SetTag(ext.DBName, "dynamodb")
	span.SetTag(ext.ResourceName, params.command)
	if params.table != nil {
		span.SetTag(ext.ResourceName, fmt.Sprintf("%s %s", params.command, params.Table()))
	}
	span.SetOperationName(fmt.Sprintf("dynamodb.%s", params.command))
	span.SetTag("db.operation", params.command)
	span.SetTag("db.dynamodb.table", params.Table())
	return span
}

func (d *DynamoDB) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, option ...request.Option) (*dynamodb.PutItemOutput, error) {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		return d.client.PutItemWithContext(ctx, input, option...)
	}
	defer span.Finish()

	params := &tagsParams{
		command: "PutItem",
		table:   input.TableName,
	}
	span = d.setTags(span, params)

	res, err := d.client.PutItemWithContext(ctx, input)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}

func (d *DynamoDB) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, option ...request.Option) (*dynamodb.QueryOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) QueryPages(input *dynamodb.QueryInput, f func(*dynamodb.QueryOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) QueryPagesWithContext(context aws.Context, input *dynamodb.QueryInput, f func(*dynamodb.QueryOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableFromBackup(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableFromBackupWithContext(context aws.Context, input *dynamodb.RestoreTableFromBackupInput, option ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableFromBackupRequest(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableToPointInTime(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableToPointInTimeWithContext(context aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, option ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) RestoreTableToPointInTimeRequest(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ScanWithContext(context aws.Context, input *dynamodb.ScanInput, option ...request.Option) (*dynamodb.ScanOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ScanPages(input *dynamodb.ScanInput, f func(*dynamodb.ScanOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) ScanPagesWithContext(context aws.Context, input *dynamodb.ScanInput, f func(*dynamodb.ScanOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TagResourceWithContext(context aws.Context, input *dynamodb.TagResourceInput, option ...request.Option) (*dynamodb.TagResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TagResourceRequest(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactGetItemsWithContext(context aws.Context, input *dynamodb.TransactGetItemsInput, option ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactWriteItemsWithContext(context aws.Context, input *dynamodb.TransactWriteItemsInput, option ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UntagResourceWithContext(context aws.Context, input *dynamodb.UntagResourceInput, option ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UntagResourceRequest(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContinuousBackups(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContinuousBackupsWithContext(context aws.Context, input *dynamodb.UpdateContinuousBackupsInput, option ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContinuousBackupsRequest(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContributorInsights(input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContributorInsightsWithContext(context aws.Context, input *dynamodb.UpdateContributorInsightsInput, option ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateContributorInsightsRequest(input *dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTable(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTableWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableInput, option ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTableRequest(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTableSettings(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTableSettingsWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, option ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateGlobalTableSettingsRequest(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateItemWithContext(context aws.Context, input *dynamodb.UpdateItemInput, option ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTableWithContext(context aws.Context, input *dynamodb.UpdateTableInput, option ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTableRequest(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTableReplicaAutoScaling(input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTableReplicaAutoScalingWithContext(context aws.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput, option ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTableReplicaAutoScalingRequest(input *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTimeToLiveWithContext(context aws.Context, input *dynamodb.UpdateTimeToLiveInput, option ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) UpdateTimeToLiveRequest(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) WaitUntilTableExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) WaitUntilTableNotExists(input *dynamodb.DescribeTableInput) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamoDB) WaitUntilTableNotExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}
