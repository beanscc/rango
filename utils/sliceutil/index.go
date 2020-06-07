package sliceutil

import (
	"reflect"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method index -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// indexReflect 查找值 x 在切片 xi 中的索引位置
// xi 必须是切片类型，且 x 的类型必须和 xi 的元素类型一致
// 若 x 不在 xi 中，则 返回 -1; 若存在，则返回 x 在 xi 中的第一个索引位置
func indexReflect(slice interface{}, x interface{}) int {
	if rv := reflect.ValueOf(slice); rv.Kind() == reflect.Slice {
		l := rv.Len()
		for i := 0; i < l; i++ {
			if reflect.DeepEqual(rv.Index(i).Interface(), x) {
				return i
			}
		}
	}

	return -1
}
