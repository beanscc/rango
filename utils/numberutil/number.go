package numberutil

import (
	"fmt"
	"math"
)

// Uint8FromInt int 转 uint8
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}

	return 0, fmt.Errorf("%d is out of the unit8 range", x)
}

func IntFromFloat64(x float64) (int, error) {
	// note: go 语言规范（https://golang.org/doc/go_spec.html）中，说明里 int 型所占的位数于 uint 相同
	// 并且 uint 总是 32位或64位的，意味着 int 型，至少是 32 位
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole), nil
	}

	return 0, fmt.Errorf("%g is out of the int32 range", x)
}

// float 相等判断

// EqualFloat64 在给定精度范围内比较两个 float64 是否相等
// if limit < 0，则精度设置为机器所能达到的最大精度
// if x 和 y 的两个近似值的差小于limit，则判定x，y是相等的
func EqualFloat64(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}

	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

// // EqualFloat32 在给定精度范围内比较两个 float32 是否相等
// // if limit < 0，则精度设置为机器所能达到的最大精度
// // if x 和 y 的两个近似值的差小于limit，则判定x，y是相等的
// func EqualFloat32(x, y, limit float32) bool {
// 	if limit <= 0.0 {
// 		limit = math.SmallestNonzeroFloat32
// 	}

// 	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
// }

// EqualFloat64Prec 在给定精度范围内比较两个 float64 是否相等
// decimals 表示小数点后精度的位数
// note: 效率比 EqualFloat64 差
func EqualFloat64Prec(x, y float64, decimals int) bool {
	return equalFloatPrec(x, y, decimals)
}

// EqualFloat32Prec 在给定精度范围内比较两个 float32 是否相等
// decimals 表示小数点后精度的位数
// note: 效率比 EqualFloat64 差
func EqualFloat32Prec(x, y float32, decimals int) bool {
	return equalFloatPrec(x, y, decimals)
}

func equalFloatPrec(x, y interface{}, decimals int) bool {
	xt := fmt.Sprintf("%.*f", decimals, x)
	yt := fmt.Sprintf("%.*f", decimals, y)

	return len(xt) == len(yt) && xt == yt
}

// NumberStr 数字千分位表示。数字转换成字符串，每 3 位使用 sep 分隔
// NumberStr(10000, ',') 输出： 10,000
func NumberStr(n int64, sep rune) string {
	ns := fmt.Sprint(n)
	for i := len(ns) - 3; i > 0; i -= 3 {
		ns = ns[:i] + string(sep) + ns[i:]
	}

	return ns
}

// Round 对 x 四舍五入保留 n 为小数点
// note: 若 x 正好是整数或小数点后位数不够n位，不进行补齐；
// 栗子(n=3,保留 3 位小数点)：1.0 -> 1.0； 1 -> 1; 1.1116 -> 1.112
func Round(x float64, n int) float64 {
	pow10 := math.Pow10(n)
	return math.Trunc((x+0.5/pow10)*pow10) / pow10
}
