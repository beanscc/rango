package sliceutil

import "strconv"

// Itoa []int转换成 []string
func Itoa(xi []int) []string {
	if len(xi) == 0 {
		return nil
	}

	xs := make([]string, 0, len(xi))
	for i := 0; i < len(xi); i++ {
		xs = append(xs, strconv.Itoa(xi[i]))
	}

	return xs
}

// Atoi []string 转换成 []int
func Atoi(xs []string) ([]int, error) {
	if len(xs) == 0 {
		return nil, nil
	}

	xi := make([]int, 0, len(xs))
	for i := 0; i < len(xs); i++ {
		ti, err := strconv.Atoi(xs[i])
		if err != nil {
			return xi, err
		}
		xi = append(xi, ti)
	}

	return xi, nil
}

// ToAny 将其他类型的 slice 转成 []any
// 示例:
//
//	ss := []string{"a", "b"}
//	ToAny(ss...)  // []any{"a", "b"}
//	ToAny(ss)  // []any{[]string{"a", "b"}}
func ToAny[T any](a ...T) []any {
	out := make([]any, len(a))
	for i, v := range a {
		out[i] = v
	}
	return out
}
