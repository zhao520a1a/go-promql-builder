package vectors

import (
	"container/list"
	"github.com/zhao520a1a/go-promql-builder.git/functions"
	"github.com/zhao520a1a/go-promql-builder.git/operators"
)

type VectorAction interface {
	Apply(*list.List)
	BeforeApply(*list.List)
}

/*
	sum(
		xxx {
        	xxx = "xxx" }
	) by(instance)
*/

func CreateSumByInstance(name string, labelMatchMap operators.LabelMatchMap) *InstantVector {
	return NewInstantVectorWithLabels(name, labelMatchMap).AddSumOperator().AddLabelFilter(operators.ByLabelFilter([]string{"instance"}))
}

/*
	sum(
		xxx {
        	xxx = "xxx" } offset xxx
	) by(instance)
*/
func CreateOffsetSumByInstance(name string, labelMatchMap operators.LabelMatchMap, offsetNum int, offsetUnit byte) *InstantVector {
	return NewOffsetInstantVector(name, labelMatchMap, offsetNum, offsetUnit).AddSumOperator().AddLabelFilter(operators.ByLabelFilter([]string{"instance"}))
}

/*
	sum (
		increase(
			xxx { xxx = "xxx", xxx = "xxx", }[xxx]
		)
	) by(xxx) > xxx
*/
func CreateSumIncr(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, filterLabelName string, filterThreshold float64) *InstantVector {
	v := NewRangeVectorWithLabels(name, labelMatchMap, duration, unit).AddFunction(functions.IncreaseFunction()).AddSumOperator()
	if len(filterLabelName) > 0 {
		v.AddLabelFilter(operators.ByLabelFilter([]string{filterLabelName}))
	}
	v.OperateWithFloat(operators.GreaterThan, filterThreshold)
	return v
}

/*
	sum(
		irate(
			xxx {
				xxx = "xxx" }[xxx]
	) by(instance)
*/
func CreateSumIrateByInstance(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte) *InstantVector {
	return NewRangeVectorWithLabels(name, labelMatchMap, duration, unit).AddFunction(functions.IrateFunction()).AddSumOperator().AddLabelFilter(operators.ByLabelFilter([]string{"instance"}))
}

/*
	sum(
		irate(
			xxx {
				xxx = "xxx" }[xxx] offset xxx
	) by(instance)
*/
func CreateSumIrateWithOffsetByInstance(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, offsetNum int, offsetUnit byte) *InstantVector {
	return NewOffsetRangeVectorWithLabels(name, labelMatchMap, duration, unit, offsetNum, offsetUnit).AddFunction(functions.IrateFunction()).AddSumOperator().AddLabelFilter(operators.ByLabelFilter([]string{"instance"}))
}

/*
	sum(
		rate(
			xxx {
			xxx = "xxx" }[xxx] offset xxx
		)
	) by(xxxx)
*/
func CreateSumRate(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, offsetNum int, offsetUnit byte, filterLabelName []string) *InstantVector {
	v := NewOffsetRangeVectorWithLabels(name, labelMatchMap, duration, unit, offsetNum, offsetUnit).AddFunction(functions.RateFunction()).AddSumOperator()
	if len(filterLabelName) > 0 {
		v.AddLabelFilter(operators.ByLabelFilter(filterLabelName))
	}
	return v
}

/*
	histogram_quantile(
		0.99 ,
		sum(
			rate(
				xxx {
				xxx = "xxx" }[xxx]
			)
		) by(xxxx)
	)
*/
func CreateHistogramSumRate(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, filterLabelName []string, quantile float64) *InstantVector {
	return CreateSumRate(name, labelMatchMap, duration, unit, 0, 0, filterLabelName).AddFunction(functions.HistogramQuantileFunction(quantile))
}

/*
	histogram_quantile(
		xxx ,
		sum(
			rate(
				xxx {
				xxx = "xxx" }[xxx] offset xxx
			)
		) by(xxxx)
	)
*/
func CreateHistogramSumRateWithOffset(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, offestNum int, offsetUnit byte, filterLabelName []string, quantile float64) *InstantVector {
	return CreateSumRate(name, labelMatchMap, duration, unit, offestNum, offsetUnit, filterLabelName).AddFunction(functions.HistogramQuantileFunction(quantile))
}
