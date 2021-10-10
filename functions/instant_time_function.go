package functions

//返回被给定 UTC 时间所在月的第几天，返回值范围：1~31
func DayOfMonthFunction() *BaseFunction {
	return NewBaseFunction("day_of_month")
}

//返回被给定 UTC 时间所在周的第几天，返回值范围：0~6，0 表示星期天
func DayOfWeekFunction() *BaseFunction {
	return NewBaseFunction("day_of_week")
}

//返回当月一共有多少天，返回值范围：28~31
func DaysInMonthFunction() *BaseFunction {
	return NewBaseFunction("days_in_month")
}

//回被给定 UTC 时间的当前第几个小时，时间范围：0~23
func HourFunction() *BaseFunction {
	return NewBaseFunction("hour")
}

//返回给定 UTC 时间当前小时的第多少分钟，结果范围：0~59
func MinuteFunction() *BaseFunction {
	return NewBaseFunction("minute")
}

//回给定 UTC 时间当前属于第几个月，结果范围：0~12
func MonthFunction() *BaseFunction {
	return NewBaseFunction("month")
}

//返回被给定 UTC 时间的当前年份
func YearFunction() *BaseFunction {
	return NewBaseFunction("delta")
}
