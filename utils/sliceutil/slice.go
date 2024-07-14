package sliceutil

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/beanscc/rango/utils/cast"
)

// Filter 过滤切片，仅返回 filter(i) == true 时的元素，返回一个新切片
func Filter[S ~[]E, E any](s S, filter func(i int, e E) bool) S {
	out := make(S, 0)
	for i, e := range s {
		if filter(i, e) {
			out = append(out, e)
		}
	}

	return out
}

// Unique 切换元素去重
// 去重后不保证原
func Unique[S ~[]E, E comparable](s S) S {
	l := len(s)
	if l < 2 {
		return s
	}

	m := make(map[E]struct{})
	out := make(S, 0)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			out = append(out, v)
			m[v] = struct{}{}
		}
	}

	return out
}

func Chunk[S ~[]E, E any](s S, length int) []S {
	l := len(s)
	chunks := make([]S, 0, Group(l, length))
	for i := 0; i < l; i += length {
		end := i + length
		if end >= l {
			end = l
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func Join[S ~[]E, E any](s S, sep string) string {
	return JoinFunc(len(s), sep, func(i int) string {
		return cast.ToString(s[i])
	})
}

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

// Group 按 length 大小对 total 进行分组，可分多少组
func Group(total, length int) int {
	if total <= length {
		// 6/7, 7
		return 1
	}

	d := total / length
	m := total % length
	if m == 0 {
		// 7,7
		return d
	}

	// 8, 7
	return d + 1
}

// ToAny 将其他类型的 slice 转成 []any
// 示例:
//
//	ss := []string{"a", "b"}
//	ToAny(ss...)  // []any{"a", "b"}
//	ToAny(ss)  // []any{[]string{"a", "b"}}
func ToAny[E any](s ...E) []any {
	out := make([]any, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}

// ToMap 将 s 按元素中信息转化成 map 结构
func ToMap[S ~[]E, E any, K comparable](s S, keyFn func(e E) K) map[K]E {
	out := make(map[K]E)
	for _, v := range s {
		out[keyFn(v)] = v
	}

	return out
}

// Convert 将 S1 转换为 S2 并保留转换中的错误
func Convert[S1 ~[]From, S2 ~[]To, From any, To any](s1 S1, fn func(From) (To, error)) (S2, error) {
	out := make(S2, 0, len(s1))
	es := make([]error, 0)
	for _, v := range s1 {
		t, err := fn(v)
		if err != nil {
			es = append(es, fmt.Errorf("convert %v, err: %w", v, err))
			continue
		}
		out = append(out, t)
	}
	return out, errors.Join(es...)
}

// Map 返回一个新的 S， 其中每个元素，都应用 mapping 方法对元素进行了修改
func Map[S ~[]E, E any](mapping func(E) E, s S) S {
	out := make(S, len(s))
	for i := range s {
		out[i] = mapping(s[i])
	}
	return out
}

// Walk 遍历 s, 对其中每个元素，应用 fn 方法
func Walk[S ~[]E, E any](s S, fn func(i int, e E)) {
	for i, e := range s {
		fn(i, e)
	}
}

// Shuffle 将切片顺序打乱
func Shuffle[S ~[]E, E any](s S) {
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
