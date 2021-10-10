package operators

import (
	"container/list"
	"fmt"
	"strings"
)

type LabelFilterOperator string

const (
	Without LabelFilterOperator = "without"
	By      LabelFilterOperator = "by"
)

type FilterAction interface {
	AddFields(fields ...string) *LabelFilter
}

type LabelFilter struct {
	fields     []string // 当做hashSet容器使用
	filterType LabelFilterOperator
}

func ByLabelFilter(fields []string) *LabelFilter {
	return NewLabelFilter(By, fields)
}

func WithoutLabelFilter(fields []string) *LabelFilter {
	return NewLabelFilter(Without, fields)
}

func NewLabelFilter(filterType LabelFilterOperator, fields []string) *LabelFilter {
	return &LabelFilter{
		filterType: filterType,
		fields:     fields,
	}
}

func (m *LabelFilter) AddFields(fields ...string) *LabelFilter {
	m.fields = append(m.fields, fields...)
	return m
}

func (m *LabelFilter) Apply(items *list.List) {
	m.BeforeApply(items)
	if len(m.fields) > 0 {
		items.PushBack(fmt.Sprintf("%s(", m.filterType))
		items.PushBack(strings.Join(m.fields, ", "))
		items.PushBack(")")
	}
}

func (m *LabelFilter) BeforeApply(items *list.List) {
	// no-op
}
