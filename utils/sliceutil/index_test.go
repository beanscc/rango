package sliceutil

import (
	"reflect"
	"testing"
)

func Test_IndexReflect(t *testing.T) {
	type args struct {
		xs interface{}
		x  interface{}
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{xs: []int{1, 2, 3, 4, 4, 4}, x: 4}, 3},
		{"t2", args{xs: []int{2}, x: 1}, -1},
		{"t3", args{xs: []string{"众多", "人人车"}, x: "人人车"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := indexReflect(tt.args.xs, tt.args.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexReflect() = %v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_IndexInts(t *testing.T) {
	type args struct {
		xi []int
		x  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{xi: []int{1, 2, 3}, x: 11}, -1},
		{"t2", args{xi: []int{1, 2, 3, 3}, x: 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexInts(tt.args.xi, tt.args.x)
			if got != tt.want {
				t.Errorf("IndexInts() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_IndexInt32s(t *testing.T) {
	type args struct {
		xi []int32
		x  int32
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{xi: []int32{1, 2, 3}, x: 11}, -1},
		{"t2", args{xi: []int32{1, 2, 3, 3}, x: 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexInt32s(tt.args.xi, tt.args.x)
			if got != tt.want {
				t.Errorf("IndexInt32s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_IndexFloat64s(t *testing.T) {
	type args struct {
		xf []float64
		x  float64
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{xf: []float64{1.001, 1, 2.0, 2, 3}, x: 11}, -1},
		{"t2", args{xf: []float64{1.001, 1, 2.0, 2, 3}, x: 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexFloat64s(tt.args.xf, tt.args.x)
			if got != tt.want {
				t.Errorf("IndexFloat64s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_IndexStrings(t *testing.T) {
	type args struct {
		xs     []string
		x      string
		strict bool
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{xs: []string{"1", "2", "3"}, x: "3.0", strict: true}, -1},
		{"t2", args{xs: []string{"1", "Go", "go", "测试"}, x: "Go", strict: false}, 1},
		{"t3", args{xs: []string{"1", "Go", "go", "测试"}, x: "GO", strict: true}, -1},
		{"t4", args{xs: []string{"1", "Go", "go", "测试"}, x: "go", strict: true}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexStrings(tt.args.xs, tt.args.x)
			if got != tt.want {
				t.Errorf("IndexStrings() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
