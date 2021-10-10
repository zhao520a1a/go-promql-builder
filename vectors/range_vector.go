package vectors

import (
	"container/list"
	"fmt"
	"github.com/zhao520a1a/go-promql-builder.git/functions"
	"github.com/zhao520a1a/go-promql-builder.git/operators"
	"github.com/zhao520a1a/go-promql-builder.git/util"
	"sort"
	"strconv"
)

type RangeVector struct {
	name          string
	labelMatchMap operators.LabelMatchMap
	duration      int
	unit          byte
	stepDuration  int
	stepUnit      byte
	actions       []VectorAction
	offsetNum     int
	offsetUnit    byte
}

func NewRangeVector(name string, duration int, unit byte) *RangeVector {
	return &RangeVector{
		name:     name,
		duration: duration,
		unit:     unit,
		actions:  []VectorAction{},
	}
}

func NewRangeVectorWithLabels(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte) *RangeVector {
	return &RangeVector{
		name:          name,
		labelMatchMap: labelMatchMap,
		duration:      duration,
		unit:          unit,
		actions:       []VectorAction{},
	}
}

func NewOffsetRangeVectorWithLabels(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, offsetNum int, offsetUnit byte) *RangeVector {
	return &RangeVector{
		name:          name,
		labelMatchMap: labelMatchMap,
		duration:      duration,
		unit:          unit,
		actions:       []VectorAction{},
		offsetNum:     offsetNum,
		offsetUnit:    offsetUnit,
	}
}

func NewRangeVectorWithStep(name string, duration int, unit byte, stepDuration int, stepUnit byte) *RangeVector {
	return &RangeVector{
		name:         name,
		duration:     duration,
		unit:         unit,
		stepDuration: stepDuration,
		stepUnit:     stepUnit,
		actions:      []VectorAction{},
	}
}

func NewRangeVectorWithLabelsAndStep(name string, labelMatchMap operators.LabelMatchMap, duration int, unit byte, stepDuration int, stepUnit byte) *RangeVector {
	return &RangeVector{
		name:          name,
		labelMatchMap: labelMatchMap,
		duration:      duration,
		unit:          unit,
		stepDuration:  stepDuration,
		stepUnit:      stepUnit,
		actions:       []VectorAction{},
	}
}

func (m *RangeVector) AddFunction(action VectorAction) *InstantVector {
	m.actions = append(m.actions, action)
	return NewInstantVector(m.Build())
}

func (m *RangeVector) AddAction(action VectorAction) *InstantVector {
	m.actions = append(m.actions, action)
	return NewInstantVector(m.Build())
}

func (m *RangeVector) AbsentOverTime() *InstantVector {
	m.actions = append(m.actions, functions.AbsentOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) AvgOverTime() *InstantVector {
	m.actions = append(m.actions, functions.AvgOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Changes() *InstantVector {
	m.actions = append(m.actions, functions.ChangesFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) CountOverTime() *InstantVector {
	m.actions = append(m.actions, functions.CountOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Delta() *InstantVector {
	m.actions = append(m.actions, functions.DeltaFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Derivative() *InstantVector {
	m.actions = append(m.actions, functions.DerivFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) HoltWinters() *InstantVector {
	m.actions = append(m.actions, functions.HoltWintersFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Increase() *InstantVector {
	m.actions = append(m.actions, functions.IncreaseFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) InstantDelta() *InstantVector {
	m.actions = append(m.actions, functions.IdeltaFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) InstantRate() *InstantVector {
	m.actions = append(m.actions, functions.IrateFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) MaxOverTime() *InstantVector {
	m.actions = append(m.actions, functions.MaxOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) MinOverTime() *InstantVector {
	m.actions = append(m.actions, functions.MinOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) PredictLinear() *InstantVector {
	m.actions = append(m.actions, functions.PredictLinearFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Rate() *InstantVector {
	m.actions = append(m.actions, functions.RateFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Resets() *InstantVector {
	m.actions = append(m.actions, functions.ResetsFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) StandardDeviationOverTime() *InstantVector {
	m.actions = append(m.actions, functions.StandardDeviationOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) StandardVarianceOverTime() *InstantVector {
	m.actions = append(m.actions, functions.StandardVarianceOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) SumOverTime() *InstantVector {
	m.actions = append(m.actions, functions.SumOverTimeFunction())
	return NewInstantVector(m.Build())
}

func (m *RangeVector) Build() string {
	items := m.BuildList()
	res := util.ListToString(items)
	return res
}

func (m *RangeVector) BuildList() *list.List {
	items := &list.List{}
	items.PushFront(m.name)
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
	expr := "[" + strconv.Itoa(m.duration) + string(m.unit)
	if m.stepDuration != 0 && len(string(m.stepUnit)) > 0 {
		expr = expr + ":" + strconv.Itoa(m.duration) + string(m.unit)
	}
	expr += "]"
	items.PushBack(expr)

	if m.offsetNum > 0 {
		expr := " offset " + strconv.Itoa(m.offsetNum) + string(m.offsetUnit)
		items.PushBack(expr)
	}
	for _, actor := range m.actions {
		actor.Apply(items)
	}
	return items
}
