package metrics

import (
	"github.com/zhao520a1a/go-promql-builder.git/operators"
	"github.com/zhao520a1a/go-promql-builder.git/util"
	"github.com/zhao520a1a/go-promql-builder.git/vectors"
)

// 获取环比昨天Api响应时间比率
func GetApiReqDurationHistogramRateLpYesterday(groupName string, serviceName string, apiName string) string {
	return GetApiReqDurationHistogramRateLp(util.RequestDurationBucketMetricName, groupName, serviceName, apiName, 1, 'd', 0.99)
}

// 获取Api响应时间比率
func GetApiReqDurationHistogramRate(groupName string, serviceName string, apiName string) string {
	labelMap := map[string]string{
		"job":        "sla-yewu",
		"group_name": groupName,
		"servname":   serviceName,
		"api":        apiName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	filterLabelName := []string{"api", "le", "instance"}
	expr := vectors.CreateHistogramSumRate(util.RequestDurationBucketMetricName, labelMatchMap, 30, 's', filterLabelName, 0.99)
	return expr.Build()
}

/*
	histogram_quantile(
		0.99 ,
		sum(
			rate(
				api_request_duration_bucket { job = "sla-yewu" ,
				api = "xxx" ,
				group_name = "xxx" ,
				servname = "xxx" }[1m]
			)
		) by(api , le , instance)
	) / histogram_quantile(
		0.99 ,
		sum(
			rate(
				api_request_duration_bucket { job = "sla-yewu" ,
				api = "xxx" ,
				group_name = "xxx" ,
				servname = "xxx" }[1m] offset 1 d
			)
		) by(api , le , instance)
	) - 1 > xxx
*/
func GetApiReqDurationHistogramRateLp(metricName string, groupName string, serviceName string, apiName string, offestNum int, offsetUnit byte, quantile float64) string {
	labelMap := map[string]string{
		"job":        "sla-yewu",
		"group_name": groupName,
		"servname":   serviceName,
		"api":        apiName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	filterLabelName := []string{"api", "le", "instance"}
	expr := vectors.CreateHistogramSumRate(metricName, labelMatchMap, 1, 'm', filterLabelName, quantile)

	offsetSumExpr := vectors.CreateHistogramSumRateWithOffset(metricName, labelMatchMap, 1, 'm', offestNum, offsetUnit, filterLabelName, quantile)
	expr.OperateWithString(operators.Division, offsetSumExpr.Build())
	expr.OperateWithFloat(operators.Sub, 1)
	return expr.Build()
}
