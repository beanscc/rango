package sliceutil

import (
	"fmt"
	"strconv"
	"strings"
)

/*
join.go 文件提供切片的 Join 方法。包含以下实现方式：
- JoinFunc(l int, sep string, fn func(i int) string) string
*/

// JoinFunc 将长度为 l 的切片中的每个元素使用 sep 字符串连接起来并返回一个新的字符串
func JoinFunc(l int, sep string, fn func(i int) string) string {
	return joinFunc(l, sep, fn)
}

// joinFunc 将长度为 l 的切片中的每个元素使用 sep 字符串连接起来并返回一个新的字符串
// l 表示切片的长度，sep表示连接字符串，fn 方法用于将切片中每项转成 string
func joinFunc(l int, sep string, fn func(i int) string) string {
	if l == 0 {
		return ""
	}

	ss := make([]string, 0, l)
	for i := 0; i < l; i++ {
		ss = append(ss, fn(i))
	}

	return strings.Join(ss, sep)
}

// JoinInts 将 []int 中每一项，使用 sep 连接组成一个字符串并返回一个新的字符串
func JoinInts(xi []int, sep string) string {
	return joinFunc(len(xi), sep, func(i int) string {
		return strconv.Itoa(xi[i])
	})
}

// JoinInt32s 将 []int32 中每一项，使用 sep 连接组成一个字符串并返回
func JoinInt32s(xi []int32, sep string) string {
	return joinFunc(len(xi), sep, func(i int) string {
		return fmt.Sprint(xi[i])
	})
}

// JoinInt64s 将 []int64 中每一项，使用 sep 连接组成一个字符串并返回
func JoinInt64s(xi []int64, sep string) string {
	return joinFunc(len(xi), sep, func(i int) string {
		return fmt.Sprint(xi[i])
	})
}
