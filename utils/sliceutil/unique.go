package sliceutil

import (
	"reflect"
)

/*

unique.go 文件提供切片去重方法，包含以下几种不同的实现方式：
- UniqueReflect(arr interface{}) interface{} : 使用反射对一个切片进行去重，返回一个去重后的新切片
- UniqueSlice(u Unique) interface{} : 切片去重，目标切片实现 Unique 接口方法，传递给该函数，即可返回一个新的去重后的切片
	已支持以下切片类型，其他切片结构实现 Unique 接口方法即可（TestCase 中有关于结构体的示例）
	- []int
	- []string
	- []float64
- UniqueFunc(l int, fn func(i int) interface{}, appender func(i int)) : 切片去重
	- []int32
*/

// UniqueReflect 通过反射达到切片元素去重的作用
func UniqueReflect(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	// 仅支持切片类型
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	// 小于 2 个元素，即本身
	if rv.Len() < 2 {
		return arr
	}

	// new resp
	resp := reflect.MakeSlice(rv.Type(), 0, rv.Len())

	// new map 用来标记切片中值是否已经存在了
	sliceElemType := rv.Index(0).Type()
	m := reflect.MakeMap(reflect.MapOf(sliceElemType, reflect.ValueOf(true).Type()))

	for i := 0; i < rv.Len(); i++ {
		if isValid := m.MapIndex(rv.Index(i)).IsValid(); !isValid { // index v 是否是有值
			// append
			resp = reflect.Append(resp, rv.Index(i))
			// set map
			m.SetMapIndex(rv.Index(i), reflect.ValueOf(true))
		}
	}

	return resp.Interface()
}

// UniqueFunc
func UniqueFunc(l int, index func(i int) interface{}, appender func(i int)) {
	m := make(map[interface{}]bool, l)
	for i := 0; i < l; i++ {
		if _, ok := m[index(i)]; !ok {
			appender(i)
			m[index(i)] = true
		}
	}
}

// UniqueInt32s []int32 切片去重
func UniqueInt32s(xi []int32) []int32 {
	l := len(xi)
	if l < 2 {
		return xi
	}

	resp := make([]int32, 0, l)
	appender := func(i int) {
		resp = append(resp, xi[i])
	}

	indexFunc := func(i int) interface{} {
		return xi[i]
	}
	UniqueFunc(l, indexFunc, appender)

	return resp
}

// Unique 切片去重接口
type Unique interface {
	// Len 输入切片的长度
	Len() int
	// Index 输入切片的索引
	Index(i int) interface{}
	// Push 将输入切片的第 i 的项，添加到新切片中
	Push(i int)
	// Resp 去重后的新切片
	// 调用处，得到结果后自行断言处理即可
	Resp() interface{}
}

// UniqueSlice 切片去重
// note: 指针切片（如 []*int）,由于指针地址不一样，因此，即使指针存储的值一样，在此处也被认为是不同的值
func UniqueSlice(u Unique) interface{} {
	m := make(map[interface{}]bool, u.Len())
	for i := 0; i < u.Len(); i++ {
		if _, ok := m[u.Index(i)]; !ok {
			u.Push(i)
			m[u.Index(i)] = true
		}
	}

	return u.Resp()
}

// UniqueInts takes an input []int and returns a new []int without duplicate values.
func UniqueInts(xi []int) []int {
	return UniqueSlice(&IntUniqueSlice{origin: xi, resp: make([]int, 0, len(xi))}).([]int)
}

// UniqueStrings takes an input []string and returns a new []string without duplicate values.
func UniqueStrings(xs []string) []string {
	return UniqueSlice(&StringUniqueSlice{origin: xs, resp: make([]string, 0, len(xs))}).([]string)
}

// UniqueFloat64s takes an input []float64 and returns a new []float64 without duplicate values.
func UniqueFloat64s(xf []float64) []float64 {
	return UniqueSlice(&Float64UniqueSlice{origin: xf, resp: make([]float64, 0, len(xf))}).([]float64)
}

// IntUniqueSlice attaches the methods of Unique to []int.
// Takes an input []int and returns a new []int without duplicate values.
type IntUniqueSlice struct {
	origin []int
	resp   []int
}

func (s IntUniqueSlice) Len() int { return len(s.origin) }

func (s IntUniqueSlice) Index(i int) interface{} { return s.origin[i] }

func (s *IntUniqueSlice) Push(i int) { s.resp = append(s.resp, s.Index(i).(int)) }

func (s IntUniqueSlice) Resp() interface{} { return s.resp }

// StringUniqueSlice attaches the methods of Unique to []string.
// Takes an input []string and returns a new []string without duplicate values.
type StringUniqueSlice struct {
	origin []string
	resp   []string
}

func (s StringUniqueSlice) Len() int { return len(s.origin) }

func (s StringUniqueSlice) Index(i int) interface{} { return s.origin[i] }

func (s *StringUniqueSlice) Push(i int) { s.resp = append(s.resp, s.Index(i).(string)) }

func (s StringUniqueSlice) Resp() interface{} { return s.resp }

// Float64UniqueSlice attaches the methods of Unique to []float64.
// Takes an input []float64 and returns a new []float64 without duplicate values.
type Float64UniqueSlice struct {
	origin []float64
	resp   []float64
}

func (s Float64UniqueSlice) Len() int { return len(s.origin) }

func (s Float64UniqueSlice) Index(i int) interface{} { return s.origin[i] }

func (s *Float64UniqueSlice) Push(i int) { s.resp = append(s.resp, s.Index(i).(float64)) }

func (s Float64UniqueSlice) Resp() interface{} { return s.resp }
