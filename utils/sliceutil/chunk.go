package sliceutil

import (
	"math"
	"reflect"
)

/*
chunk.go 文件提供将切片按指定长度分块的方法，最后一个分块的切片长度可能小于指定长度。包含以下几种实现方式：
- ChunkReflect(arr interface{}, size int) interface{} : 使用反射来完成切片分块功能
- ChunkFunc(l, size int, appender func(start, end int)) : 通过自定义 appender 函数，灵活应对各种切片类型的分块
*/

// ChunkReflect 通过反射来分块
// arr 必须是 slice 类型，返回结果是 []slice 这样的结果，即 arr 支持多维切片类型
// size 必须是大于0的正整数
// 举个栗子： arr 是 []int， 则返回chunks的实际类型是[][]int
// 若 arr 是[][]int 这样的二维切片，则返回的是 [][][]int 三维切片
func ChunkReflect(arr interface{}, size int) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice {
		panic("only support slice")
	}

	chunks := reflect.MakeSlice(reflect.SliceOf(rv.Type()), 0, chunkCap(rv.Len(), size))
	chunkFunc(rv.Len(), size, func(i, e int) {
		chunks = reflect.Append(chunks, rv.Slice(i, e))
	})

	return chunks.Interface()
}

// ChunkFunc chunk func
// 调用者可自行对其切片数据进行扩展
func ChunkFunc(l, size int, appender func(start, end int)) {
	chunkFunc(l, size, appender)
}

// 切片分块 func
func chunkFunc(l, size int, appender func(start, end int)) {
	for i := 0; i < l; i += size {
		end := i + size
		if end >= l {
			end = l
		}

		appender(i, end)
	}
}

// chunkCap 获取分块切片的容量
func chunkCap(l, size int) int {
	if size <= 0 {
		panic("size must > 0")
	}
	return int(math.Ceil(float64(l) / float64(size)))
}

// ChunkInts []int 切片按 size 大小分块
func ChunkInts(xi []int, size int) [][]int {
	l := len(xi)
	chunks := make([][]int, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xi[s:e])
	})
	return chunks
}

// ChunkInt32s []int32 切片按 size 大小分块
func ChunkInt32s(xi []int32, size int) [][]int32 {
	l := len(xi)
	chunks := make([][]int32, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xi[s:e])
	})
	return chunks
}

// ChunkInt64s []int64 切片按 size 大小分块成 [][]int64
func ChunkInt64s(xi []int64, size int) [][]int64 {
	l := len(xi)
	chunks := make([][]int64, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xi[s:e])
	})

	return chunks
}

// ChunkFloat32s []float32 切片按 size 大小分块成 [][]float32
func ChunkFloat32s(xf []float32, size int) [][]float32 {
	l := len(xf)
	chunks := make([][]float32, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xf[s:e])
	})

	return chunks
}

// ChunkFloat64s []float64 切片按 size 大小分块成 [][]float64
func ChunkFloat64s(xf []float64, size int) [][]float64 {
	l := len(xf)
	chunks := make([][]float64, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xf[s:e])
	})

	return chunks
}

// ChunkStrings []string 切片分块
func ChunkStrings(xs []string, size int) [][]string {
	l := len(xs)
	chunks := make([][]string, 0, chunkCap(l, size))
	chunkFunc(l, size, func(s, e int) {
		chunks = append(chunks, xs[s:e])
	})

	return chunks
}
