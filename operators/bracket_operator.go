package operators

import (
	"container/list"
)

type BracketOperator struct {
}

func NewBracketOperator() *BracketOperator {
	return &BracketOperator{}
}

func (m *BracketOperator) Apply(items *list.List) {
	m.BeforeApply(items)
	items.PushFront("(")
	items.PushBack(")")
}

func (m *BracketOperator) BeforeApply(items *list.List) {
}
