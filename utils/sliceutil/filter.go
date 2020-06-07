package sliceutil

import (
	"reflect"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method filter -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// filterReflect 通过反射完成切片过滤
// slice 必须是 切片类型
// 只有当 filter(i) == true 时才保留索引项，返回一个新切片
func filterReflect(slice interface{}, filter func(i int) bool) interface{} {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	l := rv.Len()
	resp := reflect.MakeSlice(rv.Type(), 0, l)
	for i := 0; i < l; i++ {
		if filter(i) {
			resp = reflect.Append(resp, rv.Index(i))
		}
	}

	return resp.Interface()
}
