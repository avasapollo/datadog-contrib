package main

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

type tagsParams struct{
    command string
    table   *string
}

func(p *tagsParams) Table() string{
    if p.table == nil {
        return ""
    }
    return *p.table
}

func (d *DynamoDB) tags(params *tagsParams) []tracer.StartSpanOption {
	resourceName := params.command
	if params.table != nil {
		resourceName = fmt.Sprintf("%s %s", params.command, params.Table())
	}

	return []tracer.StartSpanOption{
		tracer.ServiceName(d.config.ServiceName),
		tracer.ResourceName(resourceName),
		tracer.Tag(ext.DBType, "nosql"),
		tracer.Tag(ext.DBName, "dynamodb"),
		tracer.Tag("db.operation", params.command),
		tracer.Tag("db.dynamodb.table", params.Table()),
	}
}



func(d *DynamoDB)  BatchExecuteStatement(in *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput,error){
	params := &tagsParams{
		command: "BatchExecuteStatement",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchExecuteStatement", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchExecuteStatement(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchExecuteStatementWithContext(ctx aws.Context,in *dynamodb.BatchExecuteStatementInput,opts ...request.Option) (*dynamodb.BatchExecuteStatementOutput,error){
	params := &tagsParams{
		command: "BatchExecuteStatementWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.BatchExecuteStatementWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchExecuteStatementWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchExecuteStatementRequest(in *dynamodb.BatchExecuteStatementInput) (*request.Request,*dynamodb.BatchExecuteStatementOutput){
	params := &tagsParams{
		command: "BatchExecuteStatementRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchExecuteStatementRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchExecuteStatementRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchGetItem(in *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput,error){
	params := &tagsParams{
		command: "BatchGetItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchGetItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchGetItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchGetItemWithContext(ctx aws.Context,in *dynamodb.BatchGetItemInput,opts ...request.Option) (*dynamodb.BatchGetItemOutput,error){
	params := &tagsParams{
		command: "BatchGetItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.BatchGetItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchGetItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchGetItemRequest(in *dynamodb.BatchGetItemInput) (*request.Request,*dynamodb.BatchGetItemOutput){
	params := &tagsParams{
		command: "BatchGetItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchGetItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchGetItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchGetItemPages(in *dynamodb.BatchGetItemInput, ) (error){
	params := &tagsParams{
		command: "BatchGetItemPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchGetItemPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchGetItemPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchGetItemPagesWithContext(ctx aws.Context,in *dynamodb.BatchGetItemInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "BatchGetItemPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.BatchGetItemPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchGetItemPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchWriteItem(in *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput,error){
	params := &tagsParams{
		command: "BatchWriteItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchWriteItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchWriteItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchWriteItemWithContext(ctx aws.Context,in *dynamodb.BatchWriteItemInput,opts ...request.Option) (*dynamodb.BatchWriteItemOutput,error){
	params := &tagsParams{
		command: "BatchWriteItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.BatchWriteItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchWriteItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  BatchWriteItemRequest(in *dynamodb.BatchWriteItemInput) (*request.Request,*dynamodb.BatchWriteItemOutput){
	params := &tagsParams{
		command: "BatchWriteItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.BatchWriteItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.BatchWriteItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateBackup(in *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput,error){
	params := &tagsParams{
		command: "CreateBackup",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateBackup", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateBackup(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateBackupWithContext(ctx aws.Context,in *dynamodb.CreateBackupInput,opts ...request.Option) (*dynamodb.CreateBackupOutput,error){
	params := &tagsParams{
		command: "CreateBackupWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.CreateBackupWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateBackupWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateBackupRequest(in *dynamodb.CreateBackupInput) (*request.Request,*dynamodb.CreateBackupOutput){
	params := &tagsParams{
		command: "CreateBackupRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateBackupRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateBackupRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateGlobalTable(in *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput,error){
	params := &tagsParams{
		command: "CreateGlobalTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateGlobalTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateGlobalTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateGlobalTableWithContext(ctx aws.Context,in *dynamodb.CreateGlobalTableInput,opts ...request.Option) (*dynamodb.CreateGlobalTableOutput,error){
	params := &tagsParams{
		command: "CreateGlobalTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.CreateGlobalTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateGlobalTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateGlobalTableRequest(in *dynamodb.CreateGlobalTableInput) (*request.Request,*dynamodb.CreateGlobalTableOutput){
	params := &tagsParams{
		command: "CreateGlobalTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateGlobalTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateGlobalTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateTable(in *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput,error){
	params := &tagsParams{
		command: "CreateTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateTableWithContext(ctx aws.Context,in *dynamodb.CreateTableInput,opts ...request.Option) (*dynamodb.CreateTableOutput,error){
	params := &tagsParams{
		command: "CreateTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.CreateTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  CreateTableRequest(in *dynamodb.CreateTableInput) (*request.Request,*dynamodb.CreateTableOutput){
	params := &tagsParams{
		command: "CreateTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.CreateTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.CreateTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteBackup(in *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput,error){
	params := &tagsParams{
		command: "DeleteBackup",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteBackup", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteBackup(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteBackupWithContext(ctx aws.Context,in *dynamodb.DeleteBackupInput,opts ...request.Option) (*dynamodb.DeleteBackupOutput,error){
	params := &tagsParams{
		command: "DeleteBackupWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DeleteBackupWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteBackupWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteBackupRequest(in *dynamodb.DeleteBackupInput) (*request.Request,*dynamodb.DeleteBackupOutput){
	params := &tagsParams{
		command: "DeleteBackupRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteBackupRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteBackupRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteItem(in *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput,error){
	params := &tagsParams{
		command: "DeleteItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteItemWithContext(ctx aws.Context,in *dynamodb.DeleteItemInput,opts ...request.Option) (*dynamodb.DeleteItemOutput,error){
	params := &tagsParams{
		command: "DeleteItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DeleteItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteItemRequest(in *dynamodb.DeleteItemInput) (*request.Request,*dynamodb.DeleteItemOutput){
	params := &tagsParams{
		command: "DeleteItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteTable(in *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput,error){
	params := &tagsParams{
		command: "DeleteTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteTableWithContext(ctx aws.Context,in *dynamodb.DeleteTableInput,opts ...request.Option) (*dynamodb.DeleteTableOutput,error){
	params := &tagsParams{
		command: "DeleteTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DeleteTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DeleteTableRequest(in *dynamodb.DeleteTableInput) (*request.Request,*dynamodb.DeleteTableOutput){
	params := &tagsParams{
		command: "DeleteTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DeleteTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DeleteTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeBackup(in *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput,error){
	params := &tagsParams{
		command: "DescribeBackup",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeBackup", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeBackup(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeBackupWithContext(ctx aws.Context,in *dynamodb.DescribeBackupInput,opts ...request.Option) (*dynamodb.DescribeBackupOutput,error){
	params := &tagsParams{
		command: "DescribeBackupWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeBackupWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeBackupWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeBackupRequest(in *dynamodb.DescribeBackupInput) (*request.Request,*dynamodb.DescribeBackupOutput){
	params := &tagsParams{
		command: "DescribeBackupRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeBackupRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeBackupRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContinuousBackups(in *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput,error){
	params := &tagsParams{
		command: "DescribeContinuousBackups",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeContinuousBackups", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContinuousBackups(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContinuousBackupsWithContext(ctx aws.Context,in *dynamodb.DescribeContinuousBackupsInput,opts ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput,error){
	params := &tagsParams{
		command: "DescribeContinuousBackupsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeContinuousBackupsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContinuousBackupsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContinuousBackupsRequest(in *dynamodb.DescribeContinuousBackupsInput) (*request.Request,*dynamodb.DescribeContinuousBackupsOutput){
	params := &tagsParams{
		command: "DescribeContinuousBackupsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeContinuousBackupsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContinuousBackupsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContributorInsights(in *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput,error){
	params := &tagsParams{
		command: "DescribeContributorInsights",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeContributorInsights", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContributorInsights(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContributorInsightsWithContext(ctx aws.Context,in *dynamodb.DescribeContributorInsightsInput,opts ...request.Option) (*dynamodb.DescribeContributorInsightsOutput,error){
	params := &tagsParams{
		command: "DescribeContributorInsightsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeContributorInsightsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContributorInsightsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeContributorInsightsRequest(in *dynamodb.DescribeContributorInsightsInput) (*request.Request,*dynamodb.DescribeContributorInsightsOutput){
	params := &tagsParams{
		command: "DescribeContributorInsightsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeContributorInsightsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeContributorInsightsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeEndpoints(in *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput,error){
	params := &tagsParams{
		command: "DescribeEndpoints",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeEndpoints", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeEndpoints(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeEndpointsWithContext(ctx aws.Context,in *dynamodb.DescribeEndpointsInput,opts ...request.Option) (*dynamodb.DescribeEndpointsOutput,error){
	params := &tagsParams{
		command: "DescribeEndpointsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeEndpointsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeEndpointsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeEndpointsRequest(in *dynamodb.DescribeEndpointsInput) (*request.Request,*dynamodb.DescribeEndpointsOutput){
	params := &tagsParams{
		command: "DescribeEndpointsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeEndpointsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeEndpointsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeExport(in *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput,error){
	params := &tagsParams{
		command: "DescribeExport",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeExport", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeExport(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeExportWithContext(ctx aws.Context,in *dynamodb.DescribeExportInput,opts ...request.Option) (*dynamodb.DescribeExportOutput,error){
	params := &tagsParams{
		command: "DescribeExportWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeExportWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeExportWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeExportRequest(in *dynamodb.DescribeExportInput) (*request.Request,*dynamodb.DescribeExportOutput){
	params := &tagsParams{
		command: "DescribeExportRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeExportRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeExportRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTable(in *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput,error){
	params := &tagsParams{
		command: "DescribeGlobalTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeGlobalTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTableWithContext(ctx aws.Context,in *dynamodb.DescribeGlobalTableInput,opts ...request.Option) (*dynamodb.DescribeGlobalTableOutput,error){
	params := &tagsParams{
		command: "DescribeGlobalTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeGlobalTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTableRequest(in *dynamodb.DescribeGlobalTableInput) (*request.Request,*dynamodb.DescribeGlobalTableOutput){
	params := &tagsParams{
		command: "DescribeGlobalTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeGlobalTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTableSettings(in *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput,error){
	params := &tagsParams{
		command: "DescribeGlobalTableSettings",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeGlobalTableSettings", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTableSettings(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTableSettingsWithContext(ctx aws.Context,in *dynamodb.DescribeGlobalTableSettingsInput,opts ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput,error){
	params := &tagsParams{
		command: "DescribeGlobalTableSettingsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeGlobalTableSettingsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTableSettingsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeGlobalTableSettingsRequest(in *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request,*dynamodb.DescribeGlobalTableSettingsOutput){
	params := &tagsParams{
		command: "DescribeGlobalTableSettingsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeGlobalTableSettingsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeGlobalTableSettingsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeImport(in *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput,error){
	params := &tagsParams{
		command: "DescribeImport",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeImport", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeImport(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeImportWithContext(ctx aws.Context,in *dynamodb.DescribeImportInput,opts ...request.Option) (*dynamodb.DescribeImportOutput,error){
	params := &tagsParams{
		command: "DescribeImportWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeImportWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeImportWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeImportRequest(in *dynamodb.DescribeImportInput) (*request.Request,*dynamodb.DescribeImportOutput){
	params := &tagsParams{
		command: "DescribeImportRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeImportRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeImportRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeKinesisStreamingDestination(in *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "DescribeKinesisStreamingDestination",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeKinesisStreamingDestination", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeKinesisStreamingDestination(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeKinesisStreamingDestinationWithContext(ctx aws.Context,in *dynamodb.DescribeKinesisStreamingDestinationInput,opts ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "DescribeKinesisStreamingDestinationWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeKinesisStreamingDestinationWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeKinesisStreamingDestinationWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeKinesisStreamingDestinationRequest(in *dynamodb.DescribeKinesisStreamingDestinationInput) (*request.Request,*dynamodb.DescribeKinesisStreamingDestinationOutput){
	params := &tagsParams{
		command: "DescribeKinesisStreamingDestinationRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeKinesisStreamingDestinationRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeKinesisStreamingDestinationRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeLimits(in *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput,error){
	params := &tagsParams{
		command: "DescribeLimits",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeLimits", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeLimits(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeLimitsWithContext(ctx aws.Context,in *dynamodb.DescribeLimitsInput,opts ...request.Option) (*dynamodb.DescribeLimitsOutput,error){
	params := &tagsParams{
		command: "DescribeLimitsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeLimitsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeLimitsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeLimitsRequest(in *dynamodb.DescribeLimitsInput) (*request.Request,*dynamodb.DescribeLimitsOutput){
	params := &tagsParams{
		command: "DescribeLimitsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeLimitsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeLimitsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTable(in *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput,error){
	params := &tagsParams{
		command: "DescribeTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTableWithContext(ctx aws.Context,in *dynamodb.DescribeTableInput,opts ...request.Option) (*dynamodb.DescribeTableOutput,error){
	params := &tagsParams{
		command: "DescribeTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTableRequest(in *dynamodb.DescribeTableInput) (*request.Request,*dynamodb.DescribeTableOutput){
	params := &tagsParams{
		command: "DescribeTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTableReplicaAutoScaling(in *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput,error){
	params := &tagsParams{
		command: "DescribeTableReplicaAutoScaling",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTableReplicaAutoScaling", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTableReplicaAutoScaling(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTableReplicaAutoScalingWithContext(ctx aws.Context,in *dynamodb.DescribeTableReplicaAutoScalingInput,opts ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput,error){
	params := &tagsParams{
		command: "DescribeTableReplicaAutoScalingWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeTableReplicaAutoScalingWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTableReplicaAutoScalingWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTableReplicaAutoScalingRequest(in *dynamodb.DescribeTableReplicaAutoScalingInput) (*request.Request,*dynamodb.DescribeTableReplicaAutoScalingOutput){
	params := &tagsParams{
		command: "DescribeTableReplicaAutoScalingRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTableReplicaAutoScalingRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTableReplicaAutoScalingRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTimeToLive(in *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput,error){
	params := &tagsParams{
		command: "DescribeTimeToLive",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTimeToLive", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTimeToLive(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTimeToLiveWithContext(ctx aws.Context,in *dynamodb.DescribeTimeToLiveInput,opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput,error){
	params := &tagsParams{
		command: "DescribeTimeToLiveWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DescribeTimeToLiveWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTimeToLiveWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DescribeTimeToLiveRequest(in *dynamodb.DescribeTimeToLiveInput) (*request.Request,*dynamodb.DescribeTimeToLiveOutput){
	params := &tagsParams{
		command: "DescribeTimeToLiveRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DescribeTimeToLiveRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DescribeTimeToLiveRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DisableKinesisStreamingDestination(in *dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "DisableKinesisStreamingDestination",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DisableKinesisStreamingDestination", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DisableKinesisStreamingDestination(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DisableKinesisStreamingDestinationWithContext(ctx aws.Context,in *dynamodb.DisableKinesisStreamingDestinationInput,opts ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "DisableKinesisStreamingDestinationWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.DisableKinesisStreamingDestinationWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DisableKinesisStreamingDestinationWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  DisableKinesisStreamingDestinationRequest(in *dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request,*dynamodb.DisableKinesisStreamingDestinationOutput){
	params := &tagsParams{
		command: "DisableKinesisStreamingDestinationRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.DisableKinesisStreamingDestinationRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.DisableKinesisStreamingDestinationRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  EnableKinesisStreamingDestination(in *dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "EnableKinesisStreamingDestination",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.EnableKinesisStreamingDestination", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.EnableKinesisStreamingDestination(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  EnableKinesisStreamingDestinationWithContext(ctx aws.Context,in *dynamodb.EnableKinesisStreamingDestinationInput,opts ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput,error){
	params := &tagsParams{
		command: "EnableKinesisStreamingDestinationWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.EnableKinesisStreamingDestinationWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.EnableKinesisStreamingDestinationWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  EnableKinesisStreamingDestinationRequest(in *dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request,*dynamodb.EnableKinesisStreamingDestinationOutput){
	params := &tagsParams{
		command: "EnableKinesisStreamingDestinationRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.EnableKinesisStreamingDestinationRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.EnableKinesisStreamingDestinationRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteStatement(in *dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput,error){
	params := &tagsParams{
		command: "ExecuteStatement",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExecuteStatement", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteStatement(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteStatementWithContext(ctx aws.Context,in *dynamodb.ExecuteStatementInput,opts ...request.Option) (*dynamodb.ExecuteStatementOutput,error){
	params := &tagsParams{
		command: "ExecuteStatementWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ExecuteStatementWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteStatementWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteStatementRequest(in *dynamodb.ExecuteStatementInput) (*request.Request,*dynamodb.ExecuteStatementOutput){
	params := &tagsParams{
		command: "ExecuteStatementRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExecuteStatementRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteStatementRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteTransaction(in *dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput,error){
	params := &tagsParams{
		command: "ExecuteTransaction",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExecuteTransaction", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteTransaction(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteTransactionWithContext(ctx aws.Context,in *dynamodb.ExecuteTransactionInput,opts ...request.Option) (*dynamodb.ExecuteTransactionOutput,error){
	params := &tagsParams{
		command: "ExecuteTransactionWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ExecuteTransactionWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteTransactionWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExecuteTransactionRequest(in *dynamodb.ExecuteTransactionInput) (*request.Request,*dynamodb.ExecuteTransactionOutput){
	params := &tagsParams{
		command: "ExecuteTransactionRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExecuteTransactionRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExecuteTransactionRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExportTableToPointInTime(in *dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput,error){
	params := &tagsParams{
		command: "ExportTableToPointInTime",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExportTableToPointInTime", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExportTableToPointInTime(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExportTableToPointInTimeWithContext(ctx aws.Context,in *dynamodb.ExportTableToPointInTimeInput,opts ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput,error){
	params := &tagsParams{
		command: "ExportTableToPointInTimeWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ExportTableToPointInTimeWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExportTableToPointInTimeWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ExportTableToPointInTimeRequest(in *dynamodb.ExportTableToPointInTimeInput) (*request.Request,*dynamodb.ExportTableToPointInTimeOutput){
	params := &tagsParams{
		command: "ExportTableToPointInTimeRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ExportTableToPointInTimeRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ExportTableToPointInTimeRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  GetItem(in *dynamodb.GetItemInput) (*dynamodb.GetItemOutput,error){
	params := &tagsParams{
		command: "GetItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.GetItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.GetItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  GetItemWithContext(ctx aws.Context,in *dynamodb.GetItemInput,opts ...request.Option) (*dynamodb.GetItemOutput,error){
	params := &tagsParams{
		command: "GetItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.GetItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.GetItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  GetItemRequest(in *dynamodb.GetItemInput) (*request.Request,*dynamodb.GetItemOutput){
	params := &tagsParams{
		command: "GetItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.GetItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.GetItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ImportTable(in *dynamodb.ImportTableInput) (*dynamodb.ImportTableOutput,error){
	params := &tagsParams{
		command: "ImportTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ImportTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ImportTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ImportTableWithContext(ctx aws.Context,in *dynamodb.ImportTableInput,opts ...request.Option) (*dynamodb.ImportTableOutput,error){
	params := &tagsParams{
		command: "ImportTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ImportTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ImportTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ImportTableRequest(in *dynamodb.ImportTableInput) (*request.Request,*dynamodb.ImportTableOutput){
	params := &tagsParams{
		command: "ImportTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ImportTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ImportTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListBackups(in *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput,error){
	params := &tagsParams{
		command: "ListBackups",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListBackups", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListBackups(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListBackupsWithContext(ctx aws.Context,in *dynamodb.ListBackupsInput,opts ...request.Option) (*dynamodb.ListBackupsOutput,error){
	params := &tagsParams{
		command: "ListBackupsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListBackupsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListBackupsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListBackupsRequest(in *dynamodb.ListBackupsInput) (*request.Request,*dynamodb.ListBackupsOutput){
	params := &tagsParams{
		command: "ListBackupsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListBackupsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListBackupsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListContributorInsights(in *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput,error){
	params := &tagsParams{
		command: "ListContributorInsights",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListContributorInsights", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListContributorInsights(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListContributorInsightsWithContext(ctx aws.Context,in *dynamodb.ListContributorInsightsInput,opts ...request.Option) (*dynamodb.ListContributorInsightsOutput,error){
	params := &tagsParams{
		command: "ListContributorInsightsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListContributorInsightsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListContributorInsightsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListContributorInsightsRequest(in *dynamodb.ListContributorInsightsInput) (*request.Request,*dynamodb.ListContributorInsightsOutput){
	params := &tagsParams{
		command: "ListContributorInsightsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListContributorInsightsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListContributorInsightsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListContributorInsightsPages(in *dynamodb.ListContributorInsightsInput, ) (error){
	params := &tagsParams{
		command: "ListContributorInsightsPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListContributorInsightsPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListContributorInsightsPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListContributorInsightsPagesWithContext(ctx aws.Context,in *dynamodb.ListContributorInsightsInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "ListContributorInsightsPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListContributorInsightsPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListContributorInsightsPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListExports(in *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput,error){
	params := &tagsParams{
		command: "ListExports",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListExports", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListExports(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListExportsWithContext(ctx aws.Context,in *dynamodb.ListExportsInput,opts ...request.Option) (*dynamodb.ListExportsOutput,error){
	params := &tagsParams{
		command: "ListExportsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListExportsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListExportsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListExportsRequest(in *dynamodb.ListExportsInput) (*request.Request,*dynamodb.ListExportsOutput){
	params := &tagsParams{
		command: "ListExportsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListExportsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListExportsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListExportsPages(in *dynamodb.ListExportsInput, ) (error){
	params := &tagsParams{
		command: "ListExportsPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListExportsPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListExportsPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListExportsPagesWithContext(ctx aws.Context,in *dynamodb.ListExportsInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "ListExportsPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListExportsPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListExportsPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListGlobalTables(in *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput,error){
	params := &tagsParams{
		command: "ListGlobalTables",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListGlobalTables", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListGlobalTables(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListGlobalTablesWithContext(ctx aws.Context,in *dynamodb.ListGlobalTablesInput,opts ...request.Option) (*dynamodb.ListGlobalTablesOutput,error){
	params := &tagsParams{
		command: "ListGlobalTablesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListGlobalTablesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListGlobalTablesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListGlobalTablesRequest(in *dynamodb.ListGlobalTablesInput) (*request.Request,*dynamodb.ListGlobalTablesOutput){
	params := &tagsParams{
		command: "ListGlobalTablesRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListGlobalTablesRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListGlobalTablesRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListImports(in *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput,error){
	params := &tagsParams{
		command: "ListImports",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListImports", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListImports(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListImportsWithContext(ctx aws.Context,in *dynamodb.ListImportsInput,opts ...request.Option) (*dynamodb.ListImportsOutput,error){
	params := &tagsParams{
		command: "ListImportsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListImportsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListImportsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListImportsRequest(in *dynamodb.ListImportsInput) (*request.Request,*dynamodb.ListImportsOutput){
	params := &tagsParams{
		command: "ListImportsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListImportsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListImportsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListImportsPages(in *dynamodb.ListImportsInput, ) (error){
	params := &tagsParams{
		command: "ListImportsPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListImportsPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListImportsPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListImportsPagesWithContext(ctx aws.Context,in *dynamodb.ListImportsInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "ListImportsPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListImportsPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListImportsPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTables(in *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput,error){
	params := &tagsParams{
		command: "ListTables",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListTables", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTables(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTablesWithContext(ctx aws.Context,in *dynamodb.ListTablesInput,opts ...request.Option) (*dynamodb.ListTablesOutput,error){
	params := &tagsParams{
		command: "ListTablesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListTablesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTablesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTablesRequest(in *dynamodb.ListTablesInput) (*request.Request,*dynamodb.ListTablesOutput){
	params := &tagsParams{
		command: "ListTablesRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListTablesRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTablesRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTablesPages(in *dynamodb.ListTablesInput, ) (error){
	params := &tagsParams{
		command: "ListTablesPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListTablesPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTablesPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTablesPagesWithContext(ctx aws.Context,in *dynamodb.ListTablesInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "ListTablesPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListTablesPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTablesPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTagsOfResource(in *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput,error){
	params := &tagsParams{
		command: "ListTagsOfResource",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListTagsOfResource", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTagsOfResource(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTagsOfResourceWithContext(ctx aws.Context,in *dynamodb.ListTagsOfResourceInput,opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput,error){
	params := &tagsParams{
		command: "ListTagsOfResourceWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ListTagsOfResourceWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTagsOfResourceWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ListTagsOfResourceRequest(in *dynamodb.ListTagsOfResourceInput) (*request.Request,*dynamodb.ListTagsOfResourceOutput){
	params := &tagsParams{
		command: "ListTagsOfResourceRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ListTagsOfResourceRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ListTagsOfResourceRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput,error){
	params := &tagsParams{
		command: "PutItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.PutItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.PutItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  PutItemWithContext(ctx aws.Context,in *dynamodb.PutItemInput,opts ...request.Option) (*dynamodb.PutItemOutput,error){
	params := &tagsParams{
		command: "PutItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.PutItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.PutItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  PutItemRequest(in *dynamodb.PutItemInput) (*request.Request,*dynamodb.PutItemOutput){
	params := &tagsParams{
		command: "PutItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.PutItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.PutItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  Query(in *dynamodb.QueryInput) (*dynamodb.QueryOutput,error){
	params := &tagsParams{
		command: "Query",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.Query", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.Query(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  QueryWithContext(ctx aws.Context,in *dynamodb.QueryInput,opts ...request.Option) (*dynamodb.QueryOutput,error){
	params := &tagsParams{
		command: "QueryWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.QueryWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.QueryWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  QueryRequest(in *dynamodb.QueryInput) (*request.Request,*dynamodb.QueryOutput){
	params := &tagsParams{
		command: "QueryRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.QueryRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.QueryRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  QueryPages(in *dynamodb.QueryInput, ) (error){
	params := &tagsParams{
		command: "QueryPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.QueryPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.QueryPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  QueryPagesWithContext(ctx aws.Context,in *dynamodb.QueryInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "QueryPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.QueryPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.QueryPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableFromBackup(in *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput,error){
	params := &tagsParams{
		command: "RestoreTableFromBackup",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.RestoreTableFromBackup", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableFromBackup(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableFromBackupWithContext(ctx aws.Context,in *dynamodb.RestoreTableFromBackupInput,opts ...request.Option) (*dynamodb.RestoreTableFromBackupOutput,error){
	params := &tagsParams{
		command: "RestoreTableFromBackupWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.RestoreTableFromBackupWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableFromBackupWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableFromBackupRequest(in *dynamodb.RestoreTableFromBackupInput) (*request.Request,*dynamodb.RestoreTableFromBackupOutput){
	params := &tagsParams{
		command: "RestoreTableFromBackupRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.RestoreTableFromBackupRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableFromBackupRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableToPointInTime(in *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput,error){
	params := &tagsParams{
		command: "RestoreTableToPointInTime",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.RestoreTableToPointInTime", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableToPointInTime(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableToPointInTimeWithContext(ctx aws.Context,in *dynamodb.RestoreTableToPointInTimeInput,opts ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput,error){
	params := &tagsParams{
		command: "RestoreTableToPointInTimeWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.RestoreTableToPointInTimeWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableToPointInTimeWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  RestoreTableToPointInTimeRequest(in *dynamodb.RestoreTableToPointInTimeInput) (*request.Request,*dynamodb.RestoreTableToPointInTimeOutput){
	params := &tagsParams{
		command: "RestoreTableToPointInTimeRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.RestoreTableToPointInTimeRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.RestoreTableToPointInTimeRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  Scan(in *dynamodb.ScanInput) (*dynamodb.ScanOutput,error){
	params := &tagsParams{
		command: "Scan",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.Scan", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.Scan(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ScanWithContext(ctx aws.Context,in *dynamodb.ScanInput,opts ...request.Option) (*dynamodb.ScanOutput,error){
	params := &tagsParams{
		command: "ScanWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ScanWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ScanWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ScanRequest(in *dynamodb.ScanInput) (*request.Request,*dynamodb.ScanOutput){
	params := &tagsParams{
		command: "ScanRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ScanRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ScanRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ScanPages(in *dynamodb.ScanInput, ) (error){
	params := &tagsParams{
		command: "ScanPages",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.ScanPages", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ScanPages(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  ScanPagesWithContext(ctx aws.Context,in *dynamodb.ScanInput, ,opts ...request.Option) (error){
	params := &tagsParams{
		command: "ScanPagesWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.ScanPagesWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.ScanPagesWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TagResource(in *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput,error){
	params := &tagsParams{
		command: "TagResource",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TagResource", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TagResource(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TagResourceWithContext(ctx aws.Context,in *dynamodb.TagResourceInput,opts ...request.Option) (*dynamodb.TagResourceOutput,error){
	params := &tagsParams{
		command: "TagResourceWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.TagResourceWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TagResourceWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TagResourceRequest(in *dynamodb.TagResourceInput) (*request.Request,*dynamodb.TagResourceOutput){
	params := &tagsParams{
		command: "TagResourceRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TagResourceRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TagResourceRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactGetItems(in *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput,error){
	params := &tagsParams{
		command: "TransactGetItems",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TransactGetItems", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactGetItems(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactGetItemsWithContext(ctx aws.Context,in *dynamodb.TransactGetItemsInput,opts ...request.Option) (*dynamodb.TransactGetItemsOutput,error){
	params := &tagsParams{
		command: "TransactGetItemsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.TransactGetItemsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactGetItemsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactGetItemsRequest(in *dynamodb.TransactGetItemsInput) (*request.Request,*dynamodb.TransactGetItemsOutput){
	params := &tagsParams{
		command: "TransactGetItemsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TransactGetItemsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactGetItemsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactWriteItems(in *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput,error){
	params := &tagsParams{
		command: "TransactWriteItems",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TransactWriteItems", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactWriteItems(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactWriteItemsWithContext(ctx aws.Context,in *dynamodb.TransactWriteItemsInput,opts ...request.Option) (*dynamodb.TransactWriteItemsOutput,error){
	params := &tagsParams{
		command: "TransactWriteItemsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.TransactWriteItemsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactWriteItemsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  TransactWriteItemsRequest(in *dynamodb.TransactWriteItemsInput) (*request.Request,*dynamodb.TransactWriteItemsOutput){
	params := &tagsParams{
		command: "TransactWriteItemsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.TransactWriteItemsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.TransactWriteItemsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UntagResource(in *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput,error){
	params := &tagsParams{
		command: "UntagResource",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UntagResource", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UntagResource(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UntagResourceWithContext(ctx aws.Context,in *dynamodb.UntagResourceInput,opts ...request.Option) (*dynamodb.UntagResourceOutput,error){
	params := &tagsParams{
		command: "UntagResourceWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UntagResourceWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UntagResourceWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UntagResourceRequest(in *dynamodb.UntagResourceInput) (*request.Request,*dynamodb.UntagResourceOutput){
	params := &tagsParams{
		command: "UntagResourceRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UntagResourceRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UntagResourceRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContinuousBackups(in *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput,error){
	params := &tagsParams{
		command: "UpdateContinuousBackups",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateContinuousBackups", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContinuousBackups(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContinuousBackupsWithContext(ctx aws.Context,in *dynamodb.UpdateContinuousBackupsInput,opts ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput,error){
	params := &tagsParams{
		command: "UpdateContinuousBackupsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateContinuousBackupsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContinuousBackupsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContinuousBackupsRequest(in *dynamodb.UpdateContinuousBackupsInput) (*request.Request,*dynamodb.UpdateContinuousBackupsOutput){
	params := &tagsParams{
		command: "UpdateContinuousBackupsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateContinuousBackupsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContinuousBackupsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContributorInsights(in *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput,error){
	params := &tagsParams{
		command: "UpdateContributorInsights",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateContributorInsights", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContributorInsights(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContributorInsightsWithContext(ctx aws.Context,in *dynamodb.UpdateContributorInsightsInput,opts ...request.Option) (*dynamodb.UpdateContributorInsightsOutput,error){
	params := &tagsParams{
		command: "UpdateContributorInsightsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateContributorInsightsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContributorInsightsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateContributorInsightsRequest(in *dynamodb.UpdateContributorInsightsInput) (*request.Request,*dynamodb.UpdateContributorInsightsOutput){
	params := &tagsParams{
		command: "UpdateContributorInsightsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateContributorInsightsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateContributorInsightsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTable(in *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput,error){
	params := &tagsParams{
		command: "UpdateGlobalTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateGlobalTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTableWithContext(ctx aws.Context,in *dynamodb.UpdateGlobalTableInput,opts ...request.Option) (*dynamodb.UpdateGlobalTableOutput,error){
	params := &tagsParams{
		command: "UpdateGlobalTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateGlobalTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTableRequest(in *dynamodb.UpdateGlobalTableInput) (*request.Request,*dynamodb.UpdateGlobalTableOutput){
	params := &tagsParams{
		command: "UpdateGlobalTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateGlobalTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTableSettings(in *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput,error){
	params := &tagsParams{
		command: "UpdateGlobalTableSettings",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateGlobalTableSettings", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTableSettings(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTableSettingsWithContext(ctx aws.Context,in *dynamodb.UpdateGlobalTableSettingsInput,opts ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput,error){
	params := &tagsParams{
		command: "UpdateGlobalTableSettingsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateGlobalTableSettingsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTableSettingsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateGlobalTableSettingsRequest(in *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request,*dynamodb.UpdateGlobalTableSettingsOutput){
	params := &tagsParams{
		command: "UpdateGlobalTableSettingsRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateGlobalTableSettingsRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateGlobalTableSettingsRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateItem(in *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput,error){
	params := &tagsParams{
		command: "UpdateItem",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateItem", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateItem(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateItemWithContext(ctx aws.Context,in *dynamodb.UpdateItemInput,opts ...request.Option) (*dynamodb.UpdateItemOutput,error){
	params := &tagsParams{
		command: "UpdateItemWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateItemWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateItemWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateItemRequest(in *dynamodb.UpdateItemInput) (*request.Request,*dynamodb.UpdateItemOutput){
	params := &tagsParams{
		command: "UpdateItemRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateItemRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateItemRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTable(in *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput,error){
	params := &tagsParams{
		command: "UpdateTable",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTable", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTable(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTableWithContext(ctx aws.Context,in *dynamodb.UpdateTableInput,opts ...request.Option) (*dynamodb.UpdateTableOutput,error){
	params := &tagsParams{
		command: "UpdateTableWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateTableWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTableWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTableRequest(in *dynamodb.UpdateTableInput) (*request.Request,*dynamodb.UpdateTableOutput){
	params := &tagsParams{
		command: "UpdateTableRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTableRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTableRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTableReplicaAutoScaling(in *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput,error){
	params := &tagsParams{
		command: "UpdateTableReplicaAutoScaling",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTableReplicaAutoScaling", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTableReplicaAutoScaling(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTableReplicaAutoScalingWithContext(ctx aws.Context,in *dynamodb.UpdateTableReplicaAutoScalingInput,opts ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput,error){
	params := &tagsParams{
		command: "UpdateTableReplicaAutoScalingWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateTableReplicaAutoScalingWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTableReplicaAutoScalingWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTableReplicaAutoScalingRequest(in *dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request,*dynamodb.UpdateTableReplicaAutoScalingOutput){
	params := &tagsParams{
		command: "UpdateTableReplicaAutoScalingRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTableReplicaAutoScalingRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTableReplicaAutoScalingRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTimeToLive(in *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput,error){
	params := &tagsParams{
		command: "UpdateTimeToLive",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTimeToLive", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTimeToLive(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTimeToLiveWithContext(ctx aws.Context,in *dynamodb.UpdateTimeToLiveInput,opts ...request.Option) (*dynamodb.UpdateTimeToLiveOutput,error){
	params := &tagsParams{
		command: "UpdateTimeToLiveWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.UpdateTimeToLiveWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTimeToLiveWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  UpdateTimeToLiveRequest(in *dynamodb.UpdateTimeToLiveInput) (*request.Request,*dynamodb.UpdateTimeToLiveOutput){
	params := &tagsParams{
		command: "UpdateTimeToLiveRequest",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.UpdateTimeToLiveRequest", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.UpdateTimeToLiveRequest(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  WaitUntilTableExists(in *dynamodb.DescribeTableInput) (error){
	params := &tagsParams{
		command: "WaitUntilTableExists",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.WaitUntilTableExists", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.WaitUntilTableExists(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  WaitUntilTableExistsWithContext(ctx aws.Context,in *dynamodb.DescribeTableInput,opts ...request.WaiterOption) (error){
	params := &tagsParams{
		command: "WaitUntilTableExistsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.WaitUntilTableExistsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.WaitUntilTableExistsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  WaitUntilTableNotExists(in *dynamodb.DescribeTableInput) (error){
	params := &tagsParams{
		command: "WaitUntilTableNotExists",
		table:   in.TableName,
	}

	span, _ := tracer.StartSpanFromContext(context.Background(), "dynamodb.WaitUntilTableNotExists", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.WaitUntilTableNotExists(in)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}




func(d *DynamoDB)  WaitUntilTableNotExistsWithContext(ctx aws.Context,in *dynamodb.DescribeTableInput,opts ...request.WaiterOption) (error){
	params := &tagsParams{
		command: "WaitUntilTableNotExistsWithContext",
		table:   in.TableName,
	}

	span, ctx := tracer.StartSpanFromContext(ctx, "dynamodb.WaitUntilTableNotExistsWithContext", d.tags(params)...)
	defer span.Finish()

	res, err := d.client.WaitUntilTableNotExistsWithContext(ctx, in, opts...)
	if err != nil {
		span.SetTag(ext.Error, err)
		return nil, err
	}
	return res, nil
}



