package stringutil

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// CharLength 返回字符个数
func CharLength(s string) int {
	return utf8.RuneCountInString(s)
	// 或 使用 []rune (比 utf8.RuneCountInString(s) 慢)
	// return len([]rune(s))
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome 判断一个单词是否回文单词
// 思路：第一个字符和倒数第一个字符比较是否相等，第二和倒数第二比较，依次比较
// 有一个不相同就不是回文，否则，则是
func IsPalindrome(word string) bool {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}

	return true
}

// // IsPalindromeRune 判断一个单词是否回文单词
// // 思路：使用utf8关于rune的方法，递归处理
// func IsPalindromeRune(word string) bool {
// 	if utf8.RuneCountInString(word) <= 1 {
// 		return true
// 	}
//
// 	first, sizeOfFirst := utf8.DecodeRuneInString(word)
// 	last, sizeofLast := utf8.DecodeLastRuneInString(word)
// 	if first != last {
// 		return false
// 	}
//
// 	return IsPalindromeRune(word[sizeOfFirst : len(word)-sizeofLast])
// }

// PadLeft 使用指定字符 pad 从左侧填充补齐字符串 s 到 width 指定的字符长度
// example: PadLeft("test", 7, '*')  = `***test`
func PadLeft(s string, width int, pad rune) string {
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}

	return s
}

// PadRight 使用指定字符 pad 从右侧填充补齐字符串 s 到 width 指定的字符长度
// example: PadRight("test", 7, '*')  = `test***`
func PadRight(s string, width int, pad rune) string {
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return s + strings.Repeat(string(pad), gap)
	}

	return s
}

// SimpleSimplifyWhitespace 去掉 s 首尾空白字符，且将 s 中间出现的空白字符（包括换行、tab键、多个空格）用一个空格替换
func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), "")
}

// SimplifyWhitespace 去掉 s 首尾空白字符，且将 s 中间出现的空白字符（包括换行、tab键、多个空格）用一个空格替换
// 比 SimpleSimplifyWhitespace() 更高效
func SimplifyWhitespace(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}

	s = buffer.String()

	// 去掉末尾的空格
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}

	return s
}

// SimplifyWhitespaceWithReg 使用正则替换 s 中间的空白字符为一个空格
func SimplifyWhitespaceWithReg(s string) string {
	regx := regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)
	return strings.TrimSpace(regx.ReplaceAllLiteralString(s, " "))
}

const (
	separatorSnake = '_'
)

// Snake2Camel 蛇形下划线格式转驼峰
// 若 title == true，转成大驼峰，即："under_score" -> "UnderScore"; 否则 -> "underScore"
func Snake2Camel(s string, title bool) string {
	var prev rune
	if title {
		prev = separatorSnake
	}
	return strings.Map(
		func(r rune) rune {
			if prev == separatorSnake {
				prev = r
				return unicode.ToTitle(r)
			}

			if r == separatorSnake {
				prev = r
				return -1
			}

			prev = r
			return r
		},
		s)
}

// 后面将 crypto 库迁移过来后，一并迁移过去
// // MD5 计算 md5 值
// func MD5(b []byte) string {
// 	h := md5.New()
// 	h.Write(b)
// 	return hex.EncodeToString(h.Sum(nil))
// }
//
// // MD5Str 计算多个字符拼接后的字符串的 md5 值
// func MD5Str(s ...string) string {
// 	return MD5([]byte(strings.Join(s, "")))
// }
//
// // UUID 生成 uuid
// func UUID() string {
// 	return uuid.NewV4().String()
// }

// GBK2UTF8 transform GBK bytes to UTF-8 bytes
func GBK2UTF8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// UTF82GBK transform UTF-8 bytes to GBK bytes
func UTF82GBK(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}
