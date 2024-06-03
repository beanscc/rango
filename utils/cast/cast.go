package cast

import (
	"fmt"
	"strconv"
)

// ToString 将 a 转为为 string
func ToString(a any) string {
	var s string
	switch v := a.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	case []rune:
		s = string(v)
	case int:
		s = strconv.Itoa(v)
	case int8:
		s = strconv.Itoa(int(v))
	case int16:
		s = strconv.Itoa(int(v))
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		s = strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		s = strconv.FormatBool(v)
	default:
		s = fmt.Sprint(v)
	}

	return s
}
