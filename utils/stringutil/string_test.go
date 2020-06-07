package stringutil

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

// go test -run TestCharLength -v
func TestCharLength(t *testing.T) {
	s := "Go编程"
	fmt.Println("s的字节长度：", len(s)) // 8 中文字符是用3个字节存的
	// 一个中文字符所占的字节长度
	fmt.Println("一个中文字符所占的字节长度：", len(string(rune('编')))) // 3
	// 中英混合字符串的字符长度（字符个数）
	fmt.Println("中英混合的字符串的字符长度，例如：`Go编程`字符长度是：", len([]rune(s))) // 4 以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。
	fmt.Println("中英混合的字符串的字符长度，例如：`Go编程`字符长度是：", utf8.RuneCountInString(s))

	l := CharLength(s) // 4
	t.Log("s len: ", l)
}

// go test -run TestReverse -v
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSnake2Camel(t *testing.T) {
	type args struct {
		s     string
		title bool
	}

	tests := []struct {
		name string
		args args
		want string
		note string
	}{
		{"t1", args{"underscore", true}, "Underscore", "不含下划线"},
		{"t2", args{"underscore", false}, "underscore", "不含下划线"},

		{"t3", args{"under_score", true}, "UnderScore", "标准下划线分隔"},
		{"t4", args{"under_score", false}, "underScore", "标准下划线分隔"},

		{"t5", args{"_under_score", true}, "_UnderScore", "下划线开始"},
		{"t6", args{"_under_score", false}, "UnderScore", "下划线开始"},

		{"t7", args{"under_score_", true}, "UnderScore", "下划线结束"},
		{"t8", args{"under_score_", false}, "underScore", "下划线结束"},

		{"t9", args{"_under_score_", true}, "_UnderScore", "下划线开始并结束"},
		{"t10", args{"_under_score_", false}, "UnderScore", "下划线开始并结束"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Snake2Camel(tt.args.s, tt.args.title)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snake2Camel() = %v, want=%v", got, tt.want)
			}
		})
	}
}
