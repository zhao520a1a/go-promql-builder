package util

import (
	"fmt"
)

//参数无效
type errParamInvalid struct {
	Fun string
	Msg string
}

//包装错误
type errWarping struct {
	Fun string
	Msg string
	Err error
}

//没有找到
type errNotFound struct {
	Fun string
	Msg string
}

var _ error = (*errParamInvalid)(nil)
var _ error = (*errWarping)(nil)
var _ error = (*errNotFound)(nil)

func (e errParamInvalid) Error() string {
	return fmt.Sprintf("%s --> invalid parameter: %s ", e.Fun, e.Msg)
}

func (e errWarping) Error() string {
	return fmt.Sprintf("%s --> %s error:%s ", e.Fun, e.Msg, e.Err)
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("%s --> %s not found! ", e.Fun, e.Msg)
}

func GetStrContent(value ...interface{}) string {
	if len(value) == 0 {
		return ""
	}
	return fmt.Sprintln("params：", value)
}

func NewInvalidParamErr(fun, msg string) error {
	return &errParamInvalid{Fun: fun, Msg: msg}
}

func IsParamInvalidError(err error) bool {
	_, ok := err.(*errParamInvalid)
	return ok
}

func NewWrappingErr(fun, msg string, err error) error {
	return &errWarping{Fun: fun, Msg: msg, Err: err}
}

func NewWrappingError(fun string, err error) error {
	return &errWarping{Fun: fun, Err: err}
}

func IsWrappingError(err error) bool {
	_, ok := err.(*errWarping)
	return ok
}

func NewNotFoundErr(fun, msg string) error {
	return &errNotFound{Fun: fun, Msg: msg}
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*errNotFound)
	return ok
}
