package sliceutil

import (
	"reflect"
	"testing"
)

// func TestChunk(t *testing.T) {
// 	type args struct {
// 		xs   []any
// 		size int
// 	}
//
// 	tests := []struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		{"t1", args{xs: []uint{1, 2, 3, 4, 5}, size: 3}, [][]uint{{1, 2, 3}, {4, 5}}},
// 		{"t2", args{xs: []string{"1", "2", "3", "4", "5"}, size: 3}, [][]string{{"1", "2", "3"}, {"4", "5"}}},
// 		{"t3", args{
// 			xs: [][]int{
// 				{1, 2, 3},
// 				{4},
// 				{5, 6},
// 				{7}},
// 			size: 2}, [][][]int{
// 			[][]int{
// 				[]int{1, 2, 3},
// 				[]int{4}},
// 			[][]int{
// 				[]int{5, 6},
// 				[]int{7},
// 			},
// 		},
// 		}, // 对二维切片，分块;结果是一个三维切片
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := Chunk(tt.args.xs, tt.args.size)
// 			t.Logf("Chunk() =%+v", got)
//
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TestChunk() failed. got=%#v, want=%#v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_Group(t *testing.T) {
	type args struct {
		l    int
		size int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{10, 2}, 5},
		{"t1", args{10, 3}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Group(tt.args.l, tt.args.size)
			if got != tt.want {
				t.Errorf("Group() failed. got=%v, want=%v", got, tt.want)
				return
			}
		})
	}
}

func Test_ChunkInts(t *testing.T) {
	type args struct {
		xi   []int
		size int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"t1", args{size: 5, xi: []int{1, 2, 3, 4, 6, 5}}, [][]int{{1, 2, 3, 4, 6}, {5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.args.xi, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() []int failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_ChunkInt32s(t *testing.T) {
	type args struct {
		xi   []int32
		size int
	}

	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		{"t1", args{size: 5, xi: []int32{1, 2, 3, 4, 6, 5}}, [][]int32{{1, 2, 3, 4, 6}, {5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.args.xi, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() Int32s failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_ChunkStrings(t *testing.T) {
	type args struct {
		xs   []string
		size int
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"t1", args{xs: []string{"1", "2", "3", "4", "5"}, size: 3}, [][]string{{"1", "2", "3"}, {"4", "5"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.args.xs, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() Strings failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
