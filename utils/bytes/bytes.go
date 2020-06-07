package bytes

import "bytes"

// CountPrefix 统计sep从首部开始连续出现的次数
func CountPrefix(s, sep []byte) int {
	sl := len(s)
	sepl := len(sep)

	n := 0
	if sl < sepl {
		return n
	}

	for i := 0; i < sl; i += sepl {
		if bytes.Equal(s[i:i+sepl], sep) {
			n++
		} else {
			break
		}
	}

	return n
}
