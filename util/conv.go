package util

import (
	"container/list"
	"fmt"
	"strings"
)

func ListToString(items *list.List) string {
	return ListToNoFormatString(items)
}

func ListToNoFormatString(items *list.List) string {
	buffer := &strings.Builder{}
	for i := items.Front(); i != nil; i = i.Next() {
		buffer.WriteString(fmt.Sprintf("%v", i.Value))
	}
	return buffer.String()
}

func ListToFormatString(items *list.List) string {
	buffer := &strings.Builder{}
	for i := items.Front(); i != nil; i = i.Next() {
		str := fmt.Sprintf("%v", i.Value)
		if str == "(" {
			buffer.WriteString(str + "\n")
		} else if str == ")" {
			buffer.WriteString("\n" + str)
		} else {
			buffer.WriteString(str)
		}
	}
	return buffer.String()
}
