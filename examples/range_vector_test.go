package test

import (
	"github.com/zhao520a1a/go-promql-builder.git/operators"
	"github.com/zhao520a1a/go-promql-builder.git/vectors"
	"log"
	"testing"
)

//output: absent_over_time(http_requests_total[15m])
func TestRangeVector(t *testing.T) {
	s := vectors.NewRangeVector("http_requests_total", 15, 'm').AbsentOverTime()
	log.Printf(" %v", s.Build())
}

//output: sum(sum_over_time(http_requests_total[15m:15m])) without (instance)
func TestFunctionWithBy(t *testing.T) {
	s1 := vectors.NewRangeVectorWithStep("http_requests_total", 15, 'm', 1, 'm').SumOverTime().AddSumOperator().AddLabelFilter(operators.WithoutLabelFilter([]string{"instance"}))
	log.Printf(" %v", s1.Build())
}
