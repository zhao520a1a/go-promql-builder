package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhao520a1a/go-promql-builder.git/metrics"
	"testing"
)

func TestGetCpuUsingRateLpLastWeek(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenThreshold   float64
		wantExpr         string
	}{
		{
			"api", "reportapi", 0,
			`sum(runtime_resource_cpu_usage_current{group="api",job="sla-yewu",service="reportapi"})by(instance)/sum(runtime_resource_cpu_usage_current{group="api",job="sla-yewu",service="reportapi"} offset 1w)by(instance)-1.00`,
		},
		{
			"api", "interactclassapi", 50,
			`(sum(runtime_resource_cpu_usage_current{group="api",job="sla-yewu",service="interactclassapi"})by(instance)>50.00)/sum(runtime_resource_cpu_usage_current{group="api",job="sla-yewu",service="interactclassapi"} offset 1w)by(instance)-1.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetCpuUsingRateLpLastWeek(tt.givenGroupName, tt.givenServiceName, tt.givenThreshold))
	}
}

func TestGetCpuUsingRateLpYesterday(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		wantExpr         string
	}{
		{
			"base", "account",
			`sum(runtime_resource_cpu_usage_current{group="base",job="sla-yewu",service="account"})by(instance)/sum(runtime_resource_cpu_usage_current{group="base",job="sla-yewu",service="account"} offset 1d)by(instance)-1.00`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetCpuUsingRateLpYesterday(tt.givenGroupName, tt.givenServiceName))
	}
}

func TestGetMemoryUsingRateLpYesterday(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		wantExpr         string
	}{
		{
			"base", "account",
			`sum(runtime_resource_memory_current{group="base",job="sla-yewu",service="account"})by(instance)/sum(runtime_resource_memory_current{group="base",job="sla-yewu",service="account"} offset 1d)by(instance)-1.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetMemoryUsingRateLpYesterday(tt.givenGroupName, tt.givenServiceName))
	}
}

func TestGetRoutineUsingRateLpYesterday(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		wantExpr         string
	}{
		{
			"base", "account",
			`sum(runtime_resource_goroutine_current{group="base",job="sla-yewu",service="account"})by(instance)/sum(runtime_resource_goroutine_current{group="base",job="sla-yewu",service="account"} offset 1d)by(instance)-1.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetRoutineUsingRateLpYesterday(tt.givenGroupName, tt.givenServiceName))
	}
}

func TestGetApiReqCountRateLpYesterday(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenApiName     string
		wantExpr         string
	}{
		{
			"base", "innerim", "/InnerIm.InnerImService/BroadcastMsg",
			`sum(irate(api_request_count{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[1m]))by(instance)/sum(irate(api_request_count{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[1m] offset 1d))by(instance)-1.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetApiReqCountRateLpYesterday(tt.givenGroupName, tt.givenServiceName, tt.givenApiName))
	}
}

func TestGetDefaultApiReqCountRate(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenApiName     string
		wantExpr         string
	}{
		{
			"base", "innerim", "/InnerIm.InnerImService/BroadcastMsg",
			`sum(irate(api_request_count{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[30s]))by(instance)`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetDefaultApiReqCountRate(tt.givenGroupName, tt.givenServiceName, tt.givenApiName))
	}
}

func TestGetDefaultApiReqCount(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenApiName     string
		wantExpr         string
	}{
		{
			"base", "innerim", "/InnerIm.InnerImService/BroadcastMsg",
			`sum(increase(api_request_count{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[1d]))>5000.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetDefaultApiReqCount(tt.givenGroupName, tt.givenServiceName, tt.givenApiName, "", 5000))
	}
}

func TestGetApiReqDurationHistogramRateLpYesterday(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenApiName     string
		wantExpr         string
	}{
		{
			"base", "innerim", "/InnerIm.InnerImService/BroadcastMsg",
			`histogram_quantile(0.99,sum(rate(api_request_duration_bucket{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[1m]))by(api, le, instance))/histogram_quantile(0.99,sum(rate(api_request_duration_bucket{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[1m] offset 1d))by(api, le, instance))-1.00`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetApiReqDurationHistogramRateLpYesterday(tt.givenGroupName, tt.givenServiceName, tt.givenApiName))
	}
}

func TestGetApiReqDurationHistogramRate(t *testing.T) {
	tests := []struct {
		givenGroupName   string
		givenServiceName string
		givenApiName     string
		wantExpr         string
	}{
		{
			"base", "innerim", "/InnerIm.InnerImService/BroadcastMsg",
			`histogram_quantile(0.99,sum(rate(api_request_duration_bucket{api="/InnerIm.InnerImService/BroadcastMsg",group_name="base",job="sla-yewu",servname="innerim"}[30s]))by(api, le, instance))`,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantExpr, metrics.GetApiReqDurationHistogramRate(tt.givenGroupName, tt.givenServiceName, tt.givenApiName))
	}
}
