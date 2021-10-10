package operators

import (
	"container/list"
	"strconv"
	"strings"
)

type AggregationOperator struct {
	operatorName string
	labelName    string  //标签值
	numK         int     //前或后 k位
	quantile     float64 //分位
}

func (m *AggregationOperator) Apply(source *list.List) {
	m.BeforeApply(source)

	s := strings.Builder{}
	s.WriteString(m.operatorName)
	s.WriteString("(")
	if m.operatorName == "topk" || m.operatorName == "bottomk" {
		s.WriteString(strconv.Itoa(m.numK))
		s.WriteString(",")
	} else if m.operatorName == "quantile" {
		s.WriteString(strconv.FormatFloat(m.quantile, 'g', -1, 64))
		s.WriteString(",")
	} else if m.operatorName == "count_values" {
		s.WriteString("\"" + m.labelName + "\"")
		s.WriteString(",")
	}

	source.PushFront(s.String())
	source.PushBack(")")
}

func (m *AggregationOperator) BeforeApply(source *list.List) {
	// no-op
}

func NewAggregationOperator(name string) *AggregationOperator {
	return &AggregationOperator{
		operatorName: name,
	}
}

func NumKOperator(name string, numK int) *AggregationOperator {
	return &AggregationOperator{
		operatorName: name,
		numK:         numK,
	}
}

func CountValuesOperator(name string, labelName string) *AggregationOperator {
	return &AggregationOperator{
		operatorName: name,
		labelName:    labelName,
	}
}

func QuantileOperator(name string, quantile float64) *AggregationOperator {
	return &AggregationOperator{
		operatorName: name,
		quantile:     quantile,
	}
}
