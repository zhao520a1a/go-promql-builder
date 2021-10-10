package metrics

import (
	"github.com/zhao520a1a/go-promql-builder.git/operators"
	"github.com/zhao520a1a/go-promql-builder.git/vectors"
)

/*
	sum(
		runtime_resource_goroutine_current {
			job = "sla-yewu" ,
			group = "xxx" ,
			service = "xxx" }
	) by(instance) / sum(
		runtime_resource_goroutine_current {
			job = "sla-yewu" ,
			group = "xxx" ,
			service = "xxx" } offset xxx
	) by(instance) - 1
*/

func GetUsingRateLp(metricName string, groupName string, serviceName string, threshold float64, offsetNum int, offsetUnit byte) string {
	labelMap := map[string]string{
		"job":     "sla-yewu",
		"group":   groupName,
		"service": serviceName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	expr := vectors.CreateSumByInstance(metricName, labelMatchMap)
	if threshold != 0 {
		expr.OperateWithFloat(operators.GreaterThan, threshold).WithBracket()
	}
	offsetSumExpr := vectors.CreateOffsetSumByInstance(metricName, labelMatchMap, offsetNum, offsetUnit)
	expr.OperateWithString(operators.Division, offsetSumExpr.Build())
	expr.OperateWithFloat(operators.Sub, 1)
	return expr.Build()
}

/*
	sum(
		irate(
			api_request_count {
				job = "sla-yewu" ,
				api = "xxx" ,
				group_name = "xxx" ,
				servname = "xxx" }[xxx]
		)
	) by(instance) / sum(
		irate(
			api_request_count {
				job = "sla-yewu" ,
				api = "xxx" ,
				group_name = "xxx" ,
				servname = "xxx" }[xxx] offset xxx
		)
	) by(instance) - 1
*/
func GetApiReqCountRateLp(metricName string, groupName string, serviceName string, apiName string, duration int, unit byte, offestNum int, offsetUnit byte) string {
	labelMap := map[string]string{
		"job":        "sla-yewu",
		"group_name": groupName,
		"servname":   serviceName,
		"api":        apiName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	expr := vectors.CreateSumIrateByInstance(metricName, labelMatchMap, duration, unit)
	offsetSumExpr := vectors.CreateSumIrateWithOffsetByInstance(metricName, labelMatchMap, duration, unit, offestNum, offsetUnit)
	expr.OperateWithString(operators.Division, offsetSumExpr.Build())
	expr.OperateWithFloat(operators.Sub, 1)
	return expr.Build()
}

/*
	sum(
		irate(
			api_request_count {
				job = "sla-yewu" ,
				api = "xxx" ,
				group_name = "xxx" ,
				servname = "xxx" }[xxx]
		)
	) by(instance)
*/
func GetApiReqCountRate(metricName string, groupName string, serviceName string, apiName string, duration int, unit byte) string {
	labelMap := map[string]string{
		"job":        "sla-yewu",
		"group_name": groupName,
		"servname":   serviceName,
		"api":        apiName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	expr := vectors.CreateSumIrateByInstance(metricName, labelMatchMap, duration, unit)
	return expr.Build()
}

func GetApiReqCount(metricName string, groupName string, serviceName string, apiName string, filterLabelName string, filterThreshold float64) string {
	labelMap := map[string]string{
		"job":        "sla-yewu",
		"group_name": groupName,
		"servname":   serviceName,
		"api":        apiName,
	}
	labelMatchMap := operators.GetEqualMatchLabels(labelMap)
	return CreateSumIncrForDay(metricName, labelMatchMap, filterLabelName, filterThreshold)
}

/*
	sum (
		increase(
			xxx { xxx = "xxx", xxx = "xxx", }[xxx]
		)
	) by(xxx) > [1d]
*/

func CreateSumIncrForDay(metricName string, labelMap map[string]string, filterLabelName string, filterThreshold float64) string {
	expr := vectors.CreateSumIncr(metricName, labelMap, 1, 'd', filterLabelName, filterThreshold)
	return expr.Build()
}
