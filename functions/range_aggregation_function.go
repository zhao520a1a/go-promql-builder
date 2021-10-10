package functions

/*
入参是一个区间向量，出参一个瞬时向量 的聚合函数集合

下列函数允许随着时间的推移聚合给定范围向量的每个序列，并返回具有每个序列聚合结果的即时向量：
https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
*/
func BaseOverTimeFunction(aggregation string) *BaseFunction {
	return NewBaseFunction(aggregation + "_over_time")
}

func AbsentOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("absent")
}

func AvgOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("avg")
}

func MinOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("min")
}

func MaxOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("max")
}

func SumOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("sum")
}

func CountOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("count")
}

func QuantileOverTimeFunction(quantile float64) *BaseFunction {
	return BaseOverTimeFunction("quantile").AddFloatPreArgument(quantile)
}

func StandardDeviationOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("stddev")
}

func StandardVarianceOverTimeFunction() *BaseFunction {
	return BaseOverTimeFunction("stdvar")
}
