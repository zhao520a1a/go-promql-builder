package test

import (
	"fmt"
	"github.com/zhao520a1a/go-promql-builder.git/vectors"
	"testing"
)

//  sum (http_requests_total )
func TestInstantVector(t *testing.T) {
	sumExpr := vectors.NewInstantVector("http_requests_total").AddSumOperator().Build()
	fmt.Printf(" %v ", sumExpr)
}
