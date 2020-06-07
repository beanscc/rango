package sliceutil

import (
	"reflect"
	"strings"
	"testing"
)

func Test_FilterReflect(t *testing.T) {
	type args struct {
		arr interface{}
		fn  func(i int) bool
	}

	xs1 := []int{1, 2, 3, 4, 5, 6, 7}
	xs2 := []string{"Go", "哈哈go哈", "go to school", "郝大的 go 讲义"}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"t1", args{arr: xs1, fn: func(i int) bool {
			// 只保留偶数
			if xs1[i]%2 == 0 {
				return true
			}

			return false

		}}, []int{2, 4, 6}},
		{"t2", args{arr: xs2, fn: func(i int) bool {
			// 只保留含 "go" 字符串的项
			if strings.Contains(xs2[i], "go") {
				return true
			}

			return false
		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterReflect(tt.args.arr, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterReflect() failed. got=%#v, want=%#v", got, tt.want)
			}
		})
	}
}

func BenchmarkFilterReflect(b *testing.B) {
	type args struct {
		arr interface{}
		fn  func(i int) bool
	}

	xs1 := []int{1, 2, 3, 4, 5, 6, 7}
	xs2 := []string{"Go", "哈哈go哈", "go to school", "郝大的 go 讲义"}

	benchmarks := []struct {
		name string
		args args
		want interface{}
	}{
		{"t1", args{arr: xs1, fn: func(i int) bool {
			// 只保留偶数
			if xs1[i]%2 == 0 {
				return true
			}

			return false

		}}, []int{2, 4, 6}},
		{"t2", args{arr: xs2, fn: func(i int) bool {
			// 只保留含 "go" 字符串的项
			if strings.Contains(xs2[i], "go") {
				return true
			}

			return false
		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FilterReflect(bm.args.arr, bm.args.fn)
			}
		})
	}
}

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
			if xi1[i]%2 == 0 {
				return true
			}

			return false

		}}, []int{2, 4, 6}},
		{"t2", args{xi: xi1, fn: func(i int) bool {
			// 只保留奇数
			if xi1[i]%2 == 1 {
				return true
			}

			return false

		}}, []int{1, 3, 5, 7}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FilterInts(bm.args.xi, bm.args.fn)
				// got := FilterInts(bm.args.xi, bm.args.fn)
				// if !reflect.DeepEqual(got, bm.want) {
				// 	b.Errorf("FilterInts() failed. got=%v, want=%v", got, bm.want)
				// }
			}
		})
	}
}

// func Test_FilterFuncV2(t *testing.T) {
// 	type args struct {
// 		arr interface{}
// 		fn  func(i int) bool
// 	}

// 	xs1 := []int{1, 2, 3, 4, 5, 6, 7}
// 	xs2 := []string{"Go", "哈哈go哈", "go to school", "郝大的 go 讲义"}

// 	tests := []struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		{"t1", args{arr: xs1, fn: func(i int) bool {
// 			// 只保留偶数
// 			if xs1[i]%2 == 0 {
// 				return true
// 			}

// 			return false

// 		}}, []int{2, 4, 6}},
// 		{"t2", args{arr: xs2, fn: func(i int) bool {
// 			// 只保留含 "go" 字符串的项
// 			if strings.Contains(xs2[i], "go") {
// 				return true
// 			}

// 			return false
// 		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := FilterFuncV2(tt.args.arr, tt.args.fn)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("FilterFuncV2() failed. got=%#v, want=%#v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func BenchmarkFilterFuncV2(b *testing.B) {
// 	type args struct {
// 		arr interface{}
// 		fn  func(i int) bool
// 	}

// 	xs1 := []int{1, 2, 3, 4, 5, 6, 7}
// 	xs2 := []string{"Go", "哈哈go哈", "go to school", "郝大的 go 讲义"}

// 	benchmarks := []struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		{"t1", args{arr: xs1, fn: func(i int) bool {
// 			// 只保留偶数
// 			if xs1[i]%2 == 0 {
// 				return true
// 			}

// 			return false

// 		}}, []int{2, 4, 6}},
// 		{"t2", args{arr: xs2, fn: func(i int) bool {
// 			// 只保留含 "go" 字符串的项
// 			if strings.Contains(xs2[i], "go") {
// 				return true
// 			}

// 			return false
// 		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
// 	}

// 	for _, bm := range benchmarks {
// 		b.Run(bm.name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				FilterFuncV2(bm.args.arr, bm.args.fn)
// 			}
// 		})
// 	}
// }

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
			if xi1[i]%2 == 0 {
				return true
			}

			return false

		}}, []int{2, 4, 6}},
		{"t2", args{xi: xi1, fn: func(i int) bool {
			// 只保留奇数
			if xi1[i]%2 == 1 {
				return true
			}

			return false

		}}, []int{1, 3, 5, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterInts(tt.args.xi, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInts() failed. got=%v, want=%v", got, tt.want)
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
			if strings.Contains(xs[0][i], "go") {
				return true
			}

			return false
		}}, []string{}},
		{"t2", args{xs: xs[1], fn: func(i int) bool {
			// 只保留含 "go" 字符串的项
			if strings.Contains(xs[1][i], "go") {
				return true
			}

			return false
		}}, []string{"哈哈go哈", "go to school", "郝大的 go 讲义"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterStrings(tt.args.xs, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterStrings() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
