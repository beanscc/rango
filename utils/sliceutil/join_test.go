package sliceutil

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_JoinFunc(t *testing.T) {
	type args struct {
		values interface{} // 必须是切片
		len    int
		sep    string
		fn     func(i int) string
	}

	testValues := []interface{}{
		[]int{1, 2, 3},
		[]float32{1.0001, 1.0002, 2.00},
		[]int32{1, 2, 3},
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"t1",
			args{
				values: testValues[0],
				len:    len(testValues[0].([]int)),
				sep:    ",",
				fn: func(i int) string {
					return strconv.Itoa(testValues[0].([]int)[i])
				},
			},
			"1,2,3",
		},
		{
			"t2",
			args{
				values: testValues[1],
				len:    len(testValues[1].([]float32)),
				sep:    ",",
				fn: func(i int) string {
					return fmt.Sprintf("%v", testValues[1].([]float32)[i])
				},
			},
			"1.0001,1.0002,2",
		},
		{
			"t3",
			args{
				values: testValues[2],
				len:    len(testValues[2].([]int32)),
				sep:    ",",
				fn: func(i int) string {
					return fmt.Sprintf("'%v'", testValues[2].([]int32)[i])
				},
			},
			"'1','2','3'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinFunc(tt.args.len, tt.args.sep, tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinFunc() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_JoinInts(t *testing.T) {
	type args struct {
		ints []int
		sep  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{ints: []int{1, 2, 3}, sep: ","}, "1,2,3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinInts(tt.args.ints, tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinInts() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_JoinInt32s(t *testing.T) {
	type args struct {
		ints []int32
		sep  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{ints: []int32{1, 2, 3}, sep: ","}, "1,2,3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinInt32s(tt.args.ints, tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinInt32s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_JoinInt64s(t *testing.T) {
	type args struct {
		ints []int64
		sep  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{ints: []int64{1, 2, 3}, sep: ","}, "1,2,3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinInt64s(tt.args.ints, tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinInt64s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

// go test -v -count=1 -bench=BenchmarkJoinInt64s -benchtime=3s -benchmem -run BenchmarkJoinInt64s
func BenchmarkJoinInt64s(b *testing.B) {
	tests := make([]int64, 0)
	for i := 0; i < 1000000; i++ {
		tests = append(tests, int64(i))
	}

	for i := 0; i < b.N; i++ {
		_ = JoinInt64s(tests, "@@@@@@@@@@@@@")
	}
}
