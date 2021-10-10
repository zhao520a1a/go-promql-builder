package vectors

import (
	"container/list"
	"fmt"
	"github.com/zhao520a1a/go-promql-builder.git/functions"
	"github.com/zhao520a1a/go-promql-builder.git/operators"
	"github.com/zhao520a1a/go-promql-builder.git/util"
	"log"
	"sort"
	"strconv"
)

type Vector interface {
	Build() string
}

type InstantVector struct {
	name          string
	labelMatchMap operators.LabelMatchMap
	actions       []VectorAction
	offsetNum     int
	offsetUnit    byte
}

func NewInstantVector(name string) *InstantVector {
	return &InstantVector{
		name:    name,
		actions: []VectorAction{},
	}
}

func NewInstantVectorWithLabels(name string, labelMatchMap operators.LabelMatchMap) *InstantVector {
	return &InstantVector{
		name:          name,
		labelMatchMap: labelMatchMap,
		actions:       []VectorAction{},
	}
}

func EmptyInstantVector() *InstantVector {
	return &InstantVector{actions: []VectorAction{}}
}

func NewOffsetInstantVector(name string, labelMatchMap operators.LabelMatchMap, offsetNum int, unit byte) *InstantVector {
	return &InstantVector{
		name:          name,
		labelMatchMap: labelMatchMap,
		actions:       []VectorAction{},
		offsetNum:     offsetNum,
		offsetUnit:    unit,
	}
}

func NewScalarVector(value float64) *InstantVector {
	return EmptyInstantVector().AddFunction(functions.VectorFunction(value))
}

func NewTimeVector() *InstantVector {
	return EmptyInstantVector().AddFunction(functions.TimeFunction())
}

func (m *InstantVector) Build() string {
	items := m.buildList()
	res := util.ListToString(items)
	return res
}

func (m *InstantVector) buildList() *list.List {
	items := &list.List{}
	items.PushBack(m.name)

	if len(m.labelMatchMap) > 0 {
		items.PushBack("{")

		var keys []string
		for k, _ := range m.labelMatchMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var count int
		for _, key := range keys {
			count += 1
			value := m.labelMatchMap[key]
			items.PushBack(fmt.Sprintf("%s\"%s\"", key, value))
			if count < len(m.labelMatchMap) {
				items.PushBack(",")
			}
		}
		items.PushBack("}")
	}

	if m.offsetNum != 0 && len(string(m.offsetUnit)) > 0 {
		expr := " offset " + strconv.Itoa(m.offsetNum) + string(m.offsetUnit)
		items.PushBack(expr)
	}

	for _, actor := range m.actions {
		actor.Apply(items)
	}
	return items
}

func (m *InstantVector) AddLabelFilter(filter *operators.LabelFilter) *InstantVector {
	m.actions = append(m.actions, filter)
	return m
}

func (m *InstantVector) AddFunction(action VectorAction) *InstantVector {
	m.actions = append(m.actions, action)
	return m
}

func (m *InstantVector) AddSumOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("sum"))
	return m
}

func (m *InstantVector) AddMinOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("min"))
	return m
}

func (m *InstantVector) AddMaxOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("max"))
	return m
}
func (m *InstantVector) AddAvgOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("avg"))
	return m
}

func (m *InstantVector) AddGroupOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("group"))
	return m
}

func (m *InstantVector) AddStddevOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("stddev"))
	return m
}

func (m *InstantVector) AddCountOperator() *InstantVector {
	m.actions = append(m.actions, operators.NewAggregationOperator("count"))
	return m
}

func (m *InstantVector) AddCountValuesOperator(labelName string) *InstantVector {
	m.actions = append(m.actions, operators.CountValuesOperator("count_values", labelName))
	return m
}

func (m *InstantVector) AddBottomkOperator(numK int) *InstantVector {
	m.actions = append(m.actions, operators.NumKOperator("bottomk", numK))
	return m
}

func (m *InstantVector) AddTopkOperator(numK int) *InstantVector {
	m.actions = append(m.actions, operators.NumKOperator("topk", numK))
	return m
}

func (m *InstantVector) AddQuantileOperator(quantile float64) *InstantVector {
	fun := "QuantileOperator -->"
	if quantile > 1 || quantile < 0 {
		log.Fatalln("%s", util.NewInvalidParamErr(fun, "quantile  invalid").Error())
	}
	m.actions = append(m.actions, operators.QuantileOperator("quantile", quantile))
	return m
}

func (m *InstantVector) Absent() *InstantVector {
	m.actions = append(m.actions, functions.AbsentFunction())
	return m
}

func (m *InstantVector) Abs() *InstantVector {
	m.actions = append(m.actions, functions.AbsFunction())
	return m
}

func (m *InstantVector) Ln() *InstantVector {
	m.actions = append(m.actions, functions.LnFunction())
	return m
}

func (m *InstantVector) Log2() *InstantVector {
	m.actions = append(m.actions, functions.Log2Function())
	return m
}

func (m *InstantVector) Log10() *InstantVector {
	m.actions = append(m.actions, functions.Log10Function())
	return m
}

func (m *InstantVector) Ceil() *InstantVector {
	m.actions = append(m.actions, functions.CeilFunction())
	return m
}

func (m *InstantVector) ClampMax() *InstantVector {
	m.actions = append(m.actions, functions.ClampMaxFunction())
	return m
}

func (m *InstantVector) ClampMin() *InstantVector {
	m.actions = append(m.actions, functions.ClampMinFunction())
	return m
}

func (m *InstantVector) DayOfMonth() *InstantVector {
	m.actions = append(m.actions, functions.DayOfMonthFunction())
	return m
}

func (m *InstantVector) DayOfWeek() *InstantVector {
	m.actions = append(m.actions, functions.DayOfWeekFunction())
	return m
}

func (m *InstantVector) DaysInMonth() *InstantVector {
	m.actions = append(m.actions, functions.DaysInMonthFunction())
	return m
}

func (m *InstantVector) Floor() *InstantVector {
	m.actions = append(m.actions, functions.FloorFunction())
	return m
}

func (m *InstantVector) HistogramQuantile(quantile float64) *InstantVector {
	m.actions = append(m.actions, functions.HistogramQuantileFunction(quantile))
	return m
}

func (m *InstantVector) Hour() *InstantVector {
	m.actions = append(m.actions, functions.HourFunction())
	return m
}

func (m *InstantVector) LabelJoin(dstLabel string, separator string, itemsLabels []string) *InstantVector {
	m.actions = append(m.actions, functions.LabelJoinFunction(dstLabel, separator, itemsLabels))
	return m
}

func (m *InstantVector) LabelReplace(dstLabel string, replacement string, srcLabel string, regex string) *InstantVector {
	m.actions = append(m.actions, functions.LabelReplaceFunction(dstLabel, replacement, srcLabel, regex))
	return m
}

func (m *InstantVector) Minute() *InstantVector {
	m.actions = append(m.actions, functions.MinuteFunction())
	return m
}

func (m *InstantVector) Month() *InstantVector {
	m.actions = append(m.actions, functions.MonthFunction())
	return m
}

func (m *InstantVector) Round() *InstantVector {
	m.actions = append(m.actions, functions.RoundFunction())
	return m
}

func (m *InstantVector) Scalar() *InstantVector {
	m.actions = append(m.actions, functions.ScalarFunction())
	return m
}

func (m *InstantVector) SortAsc() *InstantVector {
	m.actions = append(m.actions, functions.SortAscFunction())
	return m
}

func (m *InstantVector) SortDesc() *InstantVector {
	m.actions = append(m.actions, functions.SortDescFunction())
	return m
}

func (m *InstantVector) Sqrt() *InstantVector {
	m.actions = append(m.actions, functions.SqrtFunction())
	return m
}

func (m *InstantVector) Timestamp() *InstantVector {
	m.actions = append(m.actions, functions.TimestampFunction())
	return m
}

func (m *InstantVector) Year() *InstantVector {
	m.actions = append(m.actions, functions.YearFunction())
	return m
}

func (m *InstantVector) OperateWithFloat(operator operators.BinaryOperator, threshold float64) *InstantVector {
	m.actions = append(m.actions, operators.NewBinaryOperatorFilter(operator, strconv.FormatFloat(threshold, 'f', 2, 64)))
	return m
}

func (m *InstantVector) OperateWithString(operator operators.BinaryOperator, expr string) *InstantVector {
	m.actions = append(m.actions, operators.NewBinaryOperatorFilter(operator, expr))
	return m
}

func (m *InstantVector) WithBracket() *InstantVector {
	m.actions = append(m.actions, operators.NewBracketOperator())
	return m
}
