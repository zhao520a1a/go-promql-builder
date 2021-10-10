package metrics

import "github.com/zhao520a1a/go-promql-builder.git/util"

func GetRoutineUsingRateLpYesterday(groupName string, serviceName string) string {
	return GetUsingRateLp(util.RoutineMetricName, groupName, serviceName, 0, 1, 'd')
}
