package sliceutil

import (
	"reflect"
	"strings"
	"testing"
)

func BenchmarkFilterInts(b *testing.B) {
	type args struct {
		xi []int
		fn func(i int) bool
	}

	xi1 := []int{1, 2, 3, 4, 5, 6, 7}

	benchmarks := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{xi: xi1, fn: func(i int) bool {
			// 只保留偶数
			return xi1[i]%2 == 0
		}}, []int{2, 4, 6}},
		{"t2", args{xi: xi1, fn: func(i int) bool {
			// 只保留奇数
			return xi1[i]%2 == 1
		}}, []int{1, 3, 5, 7}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got := Filter(bm.args.xi, bm.args.fn)
				if !reflect.DeepEqual(got, bm.want) {
					b.Errorf("Filter() Ints failed. got=%v, want=%v", got, bm.want)
				}
			}
		})
	}
}

func Test_FilterInts(t *testing.T) {
	type args struct {
		xi []int
		fn func(i int) bool
	}

	xi1 := []int{1, 2, 3, 4, 5, 6, 7}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{xi: xi1, fn: func(i int) bool {
			// 只保留偶数
			return xi1[i]%2 == 0
		}}, []int{2, 4, 6}},
		{"t2", args{xi: xi1, fn: func(i int) bool {
			// 只保留奇数
			return xi1[i]%2 == 1
		}}, []int{1, 3, 5, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.args.xi, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() Ints failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_FilterStrings(t *testing.T) {
	type args struct {
		xs []string
		fn func(i int) bool
	}

	xs := [][]string{
		{"Go", "哈哈哈"},
		{"Go", "哈哈go哈", "go to school", "郝大的 go 讲义"},
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1", args{xs: xs[0], fn: func(i int) bool {
			// 只保留含 "go" 字符串的项
			return strings.Contains(xs[0][i], "go")
		}}, []string{}},
		{"t2", args{xs: xs[1], fn: func(i int) bool {
			// 只保留含 "go" 字符串的项
			return strings.Contains(xs[1][i], "go")
		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.args.xs, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() Strings failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
