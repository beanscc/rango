package stringutil

import (
	"strings"
)

// 下划线格式转驼峰
// 当 title == false 时: under_score -> underScore
// 当 title == true 时：under_score -> UnderScore
func Snake2Camel(s string, title bool) string {
	ss := strings.Split(s, "_")
	if len(ss) < 2 {
		if title {
			return strings.Title(s)
		}
		return s
	}

	var su []string
	if title {
		su = ss
	} else {
		su = ss[1:]
	}

	for i := range su {
		su[i] = strings.Title(su[i])
	}

	return strings.Join(ss, "")
}
