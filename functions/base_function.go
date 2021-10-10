package functions

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type BaseFunction struct {
	name         string
	arguments    []string
	preArguments []string
}

func NewBaseFunction(name string) *BaseFunction {
	return &BaseFunction{name: name}
}

func (m *BaseFunction) AddFloatArgument(arg float64) *BaseFunction {
	m.arguments = append(m.arguments, strconv.FormatFloat(arg, 'g', -1, 64))
	return m
}

func (m *BaseFunction) AddStringArguments(args ...string) *BaseFunction {
	if len(args) > 0 {
		s := fmt.Sprintf("\" %s \"", strings.Join(args, "\", \""))
		m.arguments = append(m.arguments, s)
	}
	return m
}

func (m *BaseFunction) AddFloatPreArgument(arg float64) *BaseFunction {
	m.preArguments = append(m.arguments, strconv.FormatFloat(arg, 'g', -1, 64))
	return m
}

func (m *BaseFunction) AddStringPreArgument(arg string) *BaseFunction {
	m.preArguments = append(m.arguments, "\""+arg+"\"")
	return m
}

func (m *BaseFunction) Apply(source *list.List) {
	m.BeforeApply(source)
	s := strings.Builder{}
	s.WriteString(m.name)
	s.WriteString("(")
	if len(m.preArguments) > 0 {
		s.WriteString(strings.Join(m.preArguments, ","))
		s.WriteString(",")
	}
	source.PushFront(s.String())
	if len(m.arguments) > 0 {
		source.PushBack(",")
		source.PushBack(strings.Join(m.arguments, ","))
	}
	source.PushBack(")")
}

func (m *BaseFunction) BeforeApply(source *list.List) {

}
