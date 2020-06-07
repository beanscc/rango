package sliceutil

// EqualStrings 判断 2 个字符串切片中是否相等(1. 切片长度必须相同 2. 切片中相同索引位置上的元素必须相等)
func EqualStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
