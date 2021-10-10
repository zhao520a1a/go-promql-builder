 
## 简介
go-promql-builder 是用Go编码的方式来生成复杂的PromQL，同时结合公司业务提供了通用的PromQL语句生成方法，因此减少编写相同语句和硬编码带来的隐患。
> 注：因为目前网上还没有类似的开源工具库可供使用，因此我们打算自己去维护、开发该工具包。


## 使用

项目主要分为三个独立的部分：

vectors - 表达式类型

functions - 内置函数

operators - 操作符

### 瞬时向量表达式(Instant vector)

Example #1:

```go
//output: sum (http_requests_total )
func TestInstantVector(t *testing.T) {
 	sumExpr := vector2.NewInstantVector("http_requests_total").Sum().Build()
 	log.Printf(" %v", sumExpr)
}
```


### 范围向量表达式(Range vector)

Example #1:
```go
//output: absent_over_time(http_requests_total[15m])
func TestRangeVector(t *testing.T) {
	s := vector2.NewRangeVector("http_requests_total", 15, 'm').AbsentOverTime()
	log.Printf(" %v", s.Build())
}
```

Example #2:
 ``` go
//output: sum(sum_over_time(http_requests_total[15m:1m])) without (instance)
func TestFunctionWithBy(t *testing.T) {
	s1 := vectors.NewRangeVectorWithStep("http_requests_total", 15, 'm', 1, 'm').SumOverTime().AddSumOperator().AddLabelFilter(operators.WithoutLabelFilter([]string{"instance"}))
	log.Printf(" %v", s1.Build())
}
 ```

 


### 已封装的表达式 (vectors/vector_action.go)
```go
/*
	histogram_quantile(
		0.99 ,
		sum(
			rate(
				api_request_duration_bucket {
				xxx = "xxx" }[xxx] offset xxx
			)
		) by(xxxx)
	)
*/
func CreateHistogramSumRateWithOffset(name string, labelMap map[string]string, duration int, unit byte, offsetNum int, offsetUnit byte, filterLabelName []string, quantile float64) *InstantVector {
	return CreateSumRate(name, labelMap, duration, unit, offsetNum, offsetUnit, filterLabelName).AddFunction(functions.HistogramQuantileFunction(quantile))
}
```


### 方法命名中所缩略词解释**
- 同比：SPLY (Same Period Last Year)
- 环比：LP (Last Period)


## 参考

[英文官方文档](https://prometheus.io/docs/prometheus/latest/querying/basics/)

[中文官方文档](https://prometheus.fuckcloudnative.io/di-san-zhang-prometheus/di-4-jie-cha-xun)