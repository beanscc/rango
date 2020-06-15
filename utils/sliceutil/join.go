package sliceutil

import (
	"strings"
)

//go:generate go run ../../cmd/genericgenerator/main.go -method join -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string

// JoinFunc 将长度为 l 的切片中的每个元素，使用 sep 连接组成一个字符串并返回
func JoinFunc(l int, sep string, f func(i int) string) string {
	switch l {
	case 0:
		return ""
	case 1:
		return f(0)
	}

	ss := make([]string, 0, l)
	for i := 0; i < l; i++ {
		ss = append(ss, f(i))
	}

	return strings.Join(ss, sep)
}
