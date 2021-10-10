package metrics

import "github.com/zhao520a1a/go-promql-builder.git/util"

//获取环比昨天请求数量的增长率
func GetApiReqCountRateLpYesterday(groupName string, serviceName string, apiName string) string {
	return GetApiReqCountRateLp(util.ApiRequestCountMetricName, groupName, serviceName, apiName, 1, 'm', 1, 'd')
}

//获取30s内请求数量的增长率
func GetDefaultApiReqCountRate(groupName string, serviceName string, apiName string) string {
	return GetApiReqCountRate(util.ApiRequestCountMetricName, groupName, serviceName, apiName, 30, 's')
}

//获取1 day的请求增量
func GetDefaultApiReqCount(groupName string, serviceName string, apiName string, filterLabelName string, filterThreshold float64) string {
	return GetApiReqCount(util.ApiRequestCountMetricName, groupName, serviceName, apiName, filterLabelName, filterThreshold)
}
