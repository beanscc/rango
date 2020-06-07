package sliceutil

import (
	"math"
	"reflect"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method chunk -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// chunkReflect 通过反射来分块
// arr 必须是 slice 类型，返回结果是 []slice 这样的结果，即 arr 支持多维切片类型
// size 必须是大于0的正整数
// 举个栗子： arr 是 []int， 则返回chunks的实际类型是[][]int
// 若 arr 是[][]int 这样的二维切片，则返回的是 [][][]int 三维切片
func chunkReflect(arr interface{}, size int) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	l := rv.Len()
	chunks := reflect.MakeSlice(reflect.SliceOf(rv.Type()), 0, chunkCap(l, size))
	for i := 0; i < l; i += size {
		end := i + size
		if end >= l {
			end = l
		}
		chunks = reflect.Append(chunks, rv.Slice(i, end))
	}

	return chunks.Interface()
}

// chunkCap 获取分块切片的容量
func chunkCap(l, size int) int {
	if size <= 0 {
		panic("size must > 0")
	}
	return int(math.Ceil(float64(l) / float64(size)))
}
