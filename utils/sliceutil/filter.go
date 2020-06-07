package sliceutil

import (
	"reflect"
)

/*
filter.go 该文件提供切片的 Filter 方法。包含以下实现方式：
- FilterReflect(arr interface{}, fn func(i int) bool) interface{} : 通过反射完成过滤
- FilterFunc(l int, fn func(i int) bool, appender func(i int))
*/

// FilterReflect 通过反射完成切片过滤
// arr 必须是 slice 类型
// 只有当 fn(i) == true 时才保留索引项，返回一个新切片
func FilterReflect(arr interface{}, fn func(i int) bool) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	resp := reflect.MakeSlice(rv.Type(), 0, rv.Len())
	filterFunc(rv.Len(), fn, func(i int) {
		resp = reflect.Append(resp, rv.Index(i))
	})

	return resp.Interface()
}

// FilterFunc fn() == true 时，保留该项，否则丢弃改项，返回一个新的该类型的切片
func FilterFunc(l int, fn func(i int) bool, appender func(i int)) {
	filterFunc(l, fn, appender)
}

func filterFunc(l int, fn func(i int) bool, appender func(i int)) {
	for i := 0; i < l; i++ {
		if fn(i) {
			appender(i)
		}
	}
}

// FilterInts []int 切片过滤，只保留 fn(i) == true 时的索引项，返回一个新的切片
func FilterInts(xi []int, fn func(i int) bool) []int {
	l := len(xi)
	resp := make([]int, 0, l)
	appender := func(i int) {
		resp = append(resp, xi[i])
	}

	FilterFunc(l, fn, appender)

	return resp
}

// FilterInt32s []int32 切片过滤，只保留 fn(i) == true 时的索引项，返回一个新的切片
func FilterInt32s(xi []int32, fn func(i int) bool) []int32 {
	l := len(xi)
	resp := make([]int32, 0, l)
	appender := func(i int) {
		resp = append(resp, xi[i])
	}

	FilterFunc(l, fn, appender)

	return resp
}

// FilterInt64s []int64 切片过滤，只保留 fn(i) == true 时的索引项，返回一个新的切片
func FilterInt64s(xi []int64, fn func(i int) bool) []int64 {
	l := len(xi)
	resp := make([]int64, 0, l)
	appender := func(i int) {
		resp = append(resp, xi[i])
	}

	FilterFunc(l, fn, appender)

	return resp
}

// FilterStrings []string 切片过滤，只保留 fn(i) == true 时的索引项，返回一个新切片，
func FilterStrings(xs []string, fn func(i int) bool) []string {
	l := len(xs)
	resp := make([]string, 0, l)
	appender := func(i int) {
		resp = append(resp, xs[i])
	}

	FilterFunc(l, fn, appender)

	return resp
}
