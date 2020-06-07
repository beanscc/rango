package sliceutil

import (
	"reflect"
	"testing"
)

func Test_chunkCap(t *testing.T) {
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
			got := chunkCap(tt.args.l, tt.args.size)
			if got != tt.want {
				t.Errorf("chunkCap() failed. got=%v, want=%v", got, tt.want)
				return
			}
		})
	}
}

func Test_ChunkReflect(t *testing.T) {
	type args struct {
		xs    interface{}
		etype string
		size  int
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"t1", args{xs: []int{1, 2, 3, 4, 5}, size: 3, etype: "int"}, [][]int{{1, 2, 3}, {4, 5}}},
		{"t2", args{xs: []string{"1", "2", "3", "4", "5"}, size: 3, etype: "string"}, [][]string{{"1", "2", "3"}, {"4", "5"}}},
		{"t3", args{
			xs: [][]int{
				{1, 2, 3},
				{4},
				{5, 6},
				{7}},
			size: 2,
		},
			[][][]int{
				[][]int{
					[]int{1, 2, 3},
					[]int{4}},
				[][]int{
					[]int{5, 6},
					[]int{7},
				},
			},
		}, // 对二维切片，分块;结果是一个三维切片
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChunkReflect(tt.args.xs, tt.args.size)
			t.Logf("ChunkReflect() =%#v", got)

			switch tt.args.etype {
			case "int":
				got = got.([][]int)
			case "string":
				got = got.([][]string)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkReflect() failed. got=%#v, want=%#v", got, tt.want)
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
			got := ChunkInts(tt.args.xi, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkInts() failed. got=%v, want=%v", got, tt.want)
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
			got := ChunkInt32s(tt.args.xi, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkInt32s() failed. got=%v, want=%v", got, tt.want)
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
			got := ChunkStrings(tt.args.xs, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkStrings() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
