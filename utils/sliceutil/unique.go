package sliceutil

import (
	"reflect"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method unique -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// UniqueReflect 通过反射达到切片元素去重的作用
// 若 keyFn == nil, 则按切片元素去重
// eg:
// type People struct {
//      Name string
//      Age int
// }
// slice := []People{
//      {Name: "t1", Age: 10},
//      {Name: "t2", Age: 10},
//      {Name: "t1", Age: 5},
// }
// - 按切片元素去重，则：u:=UniqueReflect(slice, nil).([]People) // 无重复，slice中的元素全部保留
// - 按 Name 属性去重，则：u:=UniqueReflect(slice, func(i int) interface{}{ return slice[i].Name}).([]People)
// - 按 Age 属性去重，则：u:=UniqueReflect(slice, func(i int) interface{}{ return slice[i].Age}).([]People)
func UniqueReflect(slice interface{}, keyFn func(i int) interface{}) interface{} {
	rv := reflect.ValueOf(slice)
	// 仅支持切片类型
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	l := rv.Len()
	// 小于 2 个元素，即本身
	if l < 2 {
		return slice
	}

	// new resp
	resp := reflect.MakeSlice(rv.Type(), 0, l)
	m := make(map[interface{}]bool, l)
	var kv interface{}
	for i := 0; i < l; i++ {
		if keyFn != nil {
			kv = keyFn(i)
		} else {
			kv = rv.Index(i).Interface()
		}
		if ok := m[kv]; !ok {
			resp = reflect.Append(resp, rv.Index(i))
			m[kv] = true
		}
	}

	return resp.Interface()
}

// // UniqueReflect 通过反射达到切片元素去重的作用
// func UniqueReflect(slice interface{}, keyFn func(i int) interface{}) interface{} {
// 	rv := reflect.ValueOf(slice)
// 	// 仅支持切片类型
// 	if rv.Kind() != reflect.Slice {
// 		panic("only support slice")
// 	}
//
// 	l := rv.Len()
// 	// 小于 2 个元素，即本身
// 	if l < 2 {
// 		return slice
// 	}
//
// 	// new resp
// 	resp := reflect.MakeSlice(rv.Type(), 0, l)
// 	// new map 用来标记切片中值是否已经存在了
// 	// sliceElemType := rv.Index(0).Type()
// 	var keyType reflect.Type
// 	if keyFn != nil {
// 		keyType = reflect.TypeOf(keyFn(0))
// 	} else {
// 		keyType = rv.Index(0).Type()
// 	}
// 	tv := reflect.ValueOf(true)
// 	m := reflect.MakeMap(reflect.MapOf(keyType, tv.Type()))
//
// 	var kv reflect.Value
// 	for i := 0; i < l; i++ {
// 		// if isValid := m.MapIndex(rv.Index(i)).IsValid(); !isValid { // index v 是否是有值
// 		// 	resp = reflect.Append(resp, rv.Index(i))          // append
// 		// 	m.SetMapIndex(rv.Index(i), reflect.ValueOf(true)) // set map
// 		// }
// 		if keyFn != nil {
// 			kv = reflect.ValueOf(keyFn(i))
// 		} else {
// 			kv = rv.Index(i)
// 		}
// 		if isValid := m.MapIndex(kv).IsValid(); !isValid {
// 			resp = reflect.Append(resp, rv.Index(i))
// 			m.SetMapIndex(kv, tv)
// 		}
// 	}
//
// 	return resp.Interface()
// }

// // uniqueReflect 通过反射达到切片元素去重的作用
// func uniqueReflect(slice interface{}) interface{} {
// 	rv := reflect.ValueOf(slice)
// 	// 仅支持切片类型
// 	if rv.Kind() != reflect.Slice {
// 		panic("only support slice")
// 	}
//
// 	l := rv.Len()
// 	// 小于 2 个元素，即本身
// 	if l < 2 {
// 		return slice
// 	}
//
// 	// new resp
// 	resp := reflect.MakeSlice(rv.Type(), 0, l)
// 	// new map 用来标记切片中值是否已经存在了
// 	sliceElemType := rv.Index(0).Type()
// 	m := reflect.MakeMap(reflect.MapOf(sliceElemType, reflect.ValueOf(true).Type()))
// 	for i := 0; i < l; i++ {
// 		if isValid := m.MapIndex(rv.Index(i)).IsValid(); !isValid { // index v 是否是有值
// 			resp = reflect.Append(resp, rv.Index(i))          // append
// 			m.SetMapIndex(rv.Index(i), reflect.ValueOf(true)) // set map
// 		}
// 	}
//
// 	return resp.Interface()
// }
