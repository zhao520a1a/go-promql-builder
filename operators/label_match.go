package operators

type LabelMatchOperator string
type LabelMatchMap map[string]string

const (
	EqualMatch    LabelMatchOperator = "="
	NotEqualMatch LabelMatchOperator = "!="
	RegexMatch    LabelMatchOperator = "=~"
	NotRegexMatch LabelMatchOperator = "!~"
)

type LabelMatch struct {
	matchOperator LabelMatchOperator
	labelMap      map[string]string
}

func GetEqualMatchLabels(labelMap map[string]string) LabelMatchMap {
	return NewLabelMatch(EqualMatch, labelMap).Build()
}

func GetNotEqualMatchLabels(labelMap map[string]string) LabelMatchMap {
	return NewLabelMatch(NotEqualMatch, labelMap).Build()
}

func GetRegexMatchLabels(labelMap map[string]string) LabelMatchMap {
	return NewLabelMatch(RegexMatch, labelMap).Build()
}

func GetNotRegexMatchLabels(labelMap map[string]string) LabelMatchMap {
	return NewLabelMatch(NotRegexMatch, labelMap).Build()
}

func NewLabelMatch(matchOperator LabelMatchOperator, labelMap map[string]string) *LabelMatch {
	return &LabelMatch{
		matchOperator: matchOperator,
		labelMap:      labelMap,
	}
}

func (m *LabelMatch) Build() LabelMatchMap {
	if len(m.labelMap) > 0 {
		res := make(map[string]string, len(m.labelMap))
		for key, value := range m.labelMap {
			res[key+string(m.matchOperator)] = value
		}
		return res
	}
	return m.labelMap
}
