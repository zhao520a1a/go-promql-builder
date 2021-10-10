package metrics

import (
	"github.com/zhao520a1a/go-promql-builder.git/util"
)

// 获取环比上周Cpu使用率
func GetCpuUsingRateLpLastWeek(groupName string, serviceName string, threshold float64) string {
	return GetUsingRateLp(util.CpuMetricName, groupName, serviceName, threshold, 1, 'w')
}

// 获取环比昨天Cpu使用率
func GetCpuUsingRateLpYesterday(groupName string, serviceName string) string {
	return GetUsingRateLp(util.CpuMetricName, groupName, serviceName, 0, 1, 'd')
}
