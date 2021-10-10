package operators

import (
	"container/list"
)

type BinaryOperator string

const (
	Add      BinaryOperator = "+"
	Sub      BinaryOperator = "-"
	Multi    BinaryOperator = "*"
	Division BinaryOperator = "/"
	Mod      BinaryOperator = "%"
	Power    BinaryOperator = "^"
)

const (
	Equal          BinaryOperator = "=="
	NotEqual       BinaryOperator = "!="
	GreaterThan    BinaryOperator = ">"
	LessThan       BinaryOperator = "<"
	GreaterOrEqual BinaryOperator = ">="
	LessOrEqual    BinaryOperator = "<="
)

const (
	And    BinaryOperator = "and"
	Or     BinaryOperator = "or"
	Unless BinaryOperator = "unless"
)

type BinaryOperatorFilter struct {
	filterType BinaryOperator
	expression string
}

func NewBinaryOperatorFilter(operator BinaryOperator, expression string) *BinaryOperatorFilter {
	return &BinaryOperatorFilter{
		filterType: operator,
		expression: expression,
	}
}

func (m *BinaryOperatorFilter) Apply(items *list.List) {
	m.BeforeApply(items)
	items.PushBack(m.filterType)
	items.PushBack(m.expression)
}

func (m *BinaryOperatorFilter) BeforeApply(items *list.List) {
	// no-op
}
