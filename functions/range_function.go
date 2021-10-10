package functions

/*
入参是一个区间向量的函数集合
*/

//每个样本数据值变化的次数
func ChangesFunction() *BaseFunction {
	return NewBaseFunction("changes")
}

//计算一个区间向量 v 的第一个元素和最后一个元素之间的差值
func DeltaFunction() *BaseFunction {
	return NewBaseFunction("delta")
}

//它计算最新的 2 个样本值之间的差值
func IdeltaFunction() *BaseFunction {
	return NewBaseFunction("idelta")
}

//使用简单的线性回归计算区间向量 v 中各个时间序列的导数
func DerivFunction() *BaseFunction {
	return NewBaseFunction("deriv")
}

//生成时间序列数据平滑值。平滑因子 sf 越低, 对旧数据的重视程度越高。趋势因子 tf 越高，对数据的趋势的考虑就越多。其中，0< sf, tf <=1。仅适用于 Gauge 类型的时间序列
//holt_winters(v range-vectors, sf scalar, tf scalar)
func HoltWintersFunction() *BaseFunction {
	return NewBaseFunction("holt_winters")
}

//直接计算区间向量 v 在时间窗口内平均增长速率，结果不带有度量指标，只有标签列表。
//注意： rate() 函数与聚合运算符（例如 sum()）或随时间聚合的函数（任何以 _over_time 结尾的函数）一起使用时，必须先执行 irate 函数，然后再进行聚合操作。
func RateFunction() *BaseFunction {
	return NewBaseFunction("rate")
}

//通过区间向量中最后两个样本数据来计算区间向量的瞬时增长率，能够更好的反应样本数据的瞬时变化状态
//注意： irate() 函数与聚合运算符（例如 sum()）或随时间聚合的函数（任何以 _over_time 结尾的函数）一起使用时，必须先执行 irate 函数，然后再进行聚合操作。
func IrateFunction() *BaseFunction {
	return NewBaseFunction("irate")
}

//预测时间序列在t秒后的值，结果不带有度量指标，只有标签列表
func PredictLinearFunction() *BaseFunction {
	return NewBaseFunction("predict_linear")
}

//返回一个计数器重置的次数。两个连续样本之间的值的减少被认为是一次计数器重置
func ResetsFunction() *BaseFunction {
	return NewBaseFunction("resets")
}

//获取区间向量中的第一个和最后一个样本并返回其增长量
func IncreaseFunction() *BaseFunction {
	return NewBaseFunction("increase")
}
