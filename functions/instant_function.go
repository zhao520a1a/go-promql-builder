package functions

import (
	"github.com/zhao520a1a/go-promql-builder.git/util"
	"log"
)

/*
入参是一个瞬时向量的函数集合
*/

//返回所有样本的绝对值
func AbsFunction() *BaseFunction {
	return NewBaseFunction("abs")
}

//有样本数据，则返回空向量，无样本数据，则返回样本值为1
func AbsentFunction() *BaseFunction {
	return NewBaseFunction("absent")
}

//将样本值向上取整数
func CeilFunction() *BaseFunction {
	return NewBaseFunction("ceil")
}

//将样本值向下取整数
func FloorFunction() *BaseFunction {
	return NewBaseFunction("floor")
}

//输入一个瞬时向量和最大值，样本数据值若大于 max，则改为 max，否则不变
func ClampMaxFunction() *BaseFunction {
	return NewBaseFunction("clamp_max")
}

//输入一个瞬时向量和最小值，样本数据值若小于 min，则改为 min，否则不变
func ClampMinFunction() *BaseFunction {
	return NewBaseFunction("clamp_min")
}

//返回各个样本值的 e 的指数值，即 e 的 N 次方。当 N 的值足够大时会返回 +Inf
func ExponentialFunction() *BaseFunction {
	return NewBaseFunction("exp")
}

//分位数计算：计算 φ (0 ≤ φ ≤ 1) 分位数的样本的最大值
//直方图指标类型自动提供带有 _bucket 后缀和相应标签的时间序列。每个样本中必须包含名为"le"的label, 用来表示每个 bucket 的上边界，没有 le 标签的样本会被忽略。
func HistogramQuantileFunction(quantile float64) *BaseFunction {
	fun := "HistogramQuantileFunction"
	if quantile > 1 || quantile < 0 {
		log.Fatalln(util.NewInvalidParamErr(fun, "quantile invalid").Error())
	}
	return NewBaseFunction("histogram_quantile").AddFloatPreArgument(quantile)
}

//通过将样本中一些标签中的值链接起来，在样本中生成一个新的标签
//label_join(v instant-vectors, dst_label string, separator string, src_label_1 string, src_label_2 string, ...)
func LabelJoinFunction(dstLabel string, separator string, srcLabels []string) *BaseFunction {
	args := []string{dstLabel, separator}
	args = append(args, srcLabels...)
	return NewBaseFunction("label_join").AddStringArguments(args...)
}

//为时间序列添加额外的标签
func LabelReplaceFunction(dstLabel string, replacement string, srcLabel string, regex string) *BaseFunction {
	return NewBaseFunction("label_replace").AddStringArguments(dstLabel, replacement, srcLabel, regex)
}

//计算瞬时向量 v 中所有样本数据的自然对数
func LnFunction() *BaseFunction {
	return NewBaseFunction("ln")
}

//计算瞬时向量 v 中所有样本数据的二进制对数
func Log2Function() *BaseFunction {
	return NewBaseFunction("log2")
}

// 计算瞬时向量 v 中所有样本数据的十进制对数
func Log10Function() *BaseFunction {
	return NewBaseFunction("log10")
}

//返回向量中所有样本值的最接近的整数
func RoundFunction() *BaseFunction {
	return NewBaseFunction("round")
}

//返回其唯一的时间序列的值作为一个标量，如果度量指标的样本数量大于 1 或者等于 0, 则返回 NaN
func ScalarFunction() *BaseFunction {
	return NewBaseFunction("scalar")
}

//对向量按元素的值进行升序排序
func SortAscFunction() *BaseFunction {
	return NewBaseFunction("sort")
}

//对向量按元素的值进行降序排序
func SortDescFunction() *BaseFunction {
	return NewBaseFunction("sort_desc")
}

//计算向量 v 中所有元素的平方根
func SqrtFunction() *BaseFunction {
	return NewBaseFunction("sqrt")
}

//返回向量 v 中的每个样本的时间戳（从 1970-01-01 到现在的秒数）,该函数从 Prometheus 2.0 版本开始引入
func TimestampFunction() *BaseFunction {
	return NewBaseFunction("timestamp")
}
