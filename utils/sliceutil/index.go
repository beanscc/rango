package sliceutil

import (
	"reflect"
	"strings"
)

/*

index.go 文件提供查找 x 在 切片 xs 中第一次出现的索引位置的方法，若不在 xs 中，则返回 -1。包含以下几种不同的实现方式：
- IndexReflect(xs interface{}, x interface{}) int :使用反射在切片 xs 中，查找 x 是否存在，若存在，则返回 x 的索引位置，否则返回 -1
- IndexSlice(index Index, x interface{}) int : 查找 x 在切片中第一个出现的索引位置。目标切片实现 Index 接口定义的方法，调用该函数，即可求得x在切片中的索引位置
	包中已通过该方式提供了以下切片的实现
	- []int
	- []string
	- []float64
- IndexFunc(l int, fn func(i int) bool) int : 通过外部 fn(i) == true 时返回该索引，如何判定由调用者控制，比较灵活

*/

// IndexReflect 查找值 x 在切片 xi 中的索引位置
// xi 必须是切片类型，且 x 的类型必须和 xi 的元素类型一致
// 若 x 不在 xi 中，则 返回 -1; 若存在，则返回 x 在 xi 中的第一个索引位置
func IndexReflect(xi interface{}, x interface{}) int {
	if rv := reflect.ValueOf(xi); rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			if reflect.DeepEqual(rv.Index(i).Interface(), x) {
				return i
			}
		}
	}

	return -1
}

// IndexFunc 在长度为 l 的切片中，遍历切片中每一项，返回第一个 fn(i) == true 时的切片索引位置，否则返回 -1
func IndexFunc(l int, fn func(i int) bool) int {
	for i := 0; i < l; i++ {
		if fn(i) {
			return i
		}
	}

	return -1
}

// IndexInt32s 查找 x 在 xi 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt32s(xi []int32, x int32) int {
	return IndexFunc(len(xi), func(i int) bool {
		return x == xi[i]
	})
}

// IndexInt64s 查找 x 在 xi 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt64s(xi []int64, x int64) int {
	return IndexFunc(len(xi), func(i int) bool {
		return x == xi[i]
	})
}

// IndexStringsEqualFold 查找 x 在 xs 中第一次出现的索引位置（不区分大小写），若不存在，则返回 -1
func IndexStringsEqualFold(xs []string, x string) int {
	return IndexFunc(len(xs), func(i int) bool {
		return strings.EqualFold(xs[i], x)
	})
}

// Index 查找某项在一个切片中的第一个索引位置的接口方法定义
type Index interface {
	// 切片长度
	Len() int
	// 切片第 i 项的类型&值等于 x 的类型&值
	// x 的实际类型必须和切片的元素类型一致
	EqualTo(i int, x interface{}) bool
}

// IndexSlice 查找 x 在切片中第一次的索引位置。若不存在，则返回 -1
// 根据实现 Index 接口的数据结构（某个切片类型）的长度遍历循环每一项，查找 x 第一次出现的索引位置
// 若不存在，则返回 -1
func IndexSlice(index Index, x interface{}) int {
	for i := 0; i < index.Len(); i++ {
		if index.EqualTo(i, x) {
			return i
		}
	}

	return -1
}

// IndexInts 查找 x 在 xi 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInts(xi []int, x int) int {
	return IndexSlice(IntIndexSlice(xi), x)
}

// IndexStrings 查找 x 在 xs 中第一次出现的索引位置；若不存在，则返回 -1
// strict 是否严格区分大小写
func IndexStrings(xs []string, x string, strict bool) int {
	if strict {
		return IndexSlice(StringIndexSlice(xs), x)
	}

	return IndexStringsEqualFold(xs, x)
}

// IndexFloat64s 查找 x 在 xf 中第一次出现的索引位置；若不存在，则返回 -1
func IndexFloat64s(xf []float64, x float64) int {
	return IndexSlice(Float64IndexSlice(xf), x)
}

// IntIndexSlice []int 切片实现 Index 接口
type IntIndexSlice []int

func (s IntIndexSlice) Len() int { return len(s) }

func (s IntIndexSlice) EqualTo(i int, x interface{}) bool { return s[i] == x.(int) }

// StringIndexSlice []string 切片实现 Index 接口
type StringIndexSlice []string

func (s StringIndexSlice) Len() int { return len(s) }

func (s StringIndexSlice) EqualTo(i int, x interface{}) bool { return s[i] == x.(string) }

// Float64IndexSlice []float64 切片实现 Index 接口
type Float64IndexSlice []float64

func (s Float64IndexSlice) Len() int { return len(s) }

func (s Float64IndexSlice) EqualTo(i int, x interface{}) bool { return s[i] == x.(float64) }
