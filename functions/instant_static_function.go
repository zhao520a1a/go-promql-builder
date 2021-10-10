package functions

/*
无入参的函数集合
*/

//返回从 1970-01-01 到现在的秒数。注意：它不是直接返回当前时间，而是时间戳
func TimeFunction() *BaseFunction {
	return NewBaseFunction("time")
}

//将标量s转换为返回一个没有标签的的向量
func VectorFunction(value float64) *BaseFunction {
	return NewBaseFunction("vectors").AddFloatArgument(value)
}
