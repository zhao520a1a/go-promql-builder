package metrics

import (
	"github.com/zhao520a1a/go-promql-builder.git/util"
)

// 获取环比昨天内存使用率
func GetMemoryUsingRateLpYesterday(groupName string, serviceName string) string {
	return GetUsingRateLp(util.MemoryMetricName, groupName, serviceName, 0, 1, 'd')
}
