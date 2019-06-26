// Package format reflect 练习.
// 12.2 节
// @date 2019/6/26
// @author Allen Lin
// @email xqlin@qq.com
package format

import (
	"reflect"
	"strconv"
)

// AnyStr 把任意值转换为字符串
func AnyStr(v interface{}) string {
	return formatAny(reflect.ValueOf(v))
}

// formatAny 格式化一个值，且不分析内部结构.
func formatAny(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String()) // 返回带双引号的字符串
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	// ... 忽略 float\complex
	case reflect.Slice, reflect.Map, reflect.Func, reflect.Chan, reflect.Ptr: // 引用类型返回地址.
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 64)
	default: // 聚合类型：reflect.Array reflect.Struct reflect.Interface
		return v.Type().String() + " value"
	}
}
