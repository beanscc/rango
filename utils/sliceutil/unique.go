package sliceutil

import (
	"reflect"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method unique -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// uniqueReflect 通过反射达到切片元素去重的作用
func uniqueReflect(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	// 仅支持切片类型
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	l := rv.Len()
	// 小于 2 个元素，即本身
	if l < 2 {
		return arr
	}

	// new resp
	resp := reflect.MakeSlice(rv.Type(), 0, l)
	// new map 用来标记切片中值是否已经存在了
	sliceElemType := rv.Index(0).Type()
	m := reflect.MakeMap(reflect.MapOf(sliceElemType, reflect.ValueOf(true).Type()))
	for i := 0; i < l; i++ {
		if isValid := m.MapIndex(rv.Index(i)).IsValid(); !isValid { // index v 是否是有值
			resp = reflect.Append(resp, rv.Index(i))          // append
			m.SetMapIndex(rv.Index(i), reflect.ValueOf(true)) // set map
		}
	}

	return resp.Interface()
}
