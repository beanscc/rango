package sliceutil

import (
	"reflect"
	"testing"
)

func Test_UniqueReflect(t *testing.T) {
	type args struct {
		value interface{}
		eType string
	}

	prtStructA := &struct_A{name: "test1"}
	prtStructA2 := &struct_A{name: "test2"}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"t1", args{[]int{1, 1, 2, 4, 6}, "int"}, []int{1, 2, 4, 6}},
		{"t2", args{[]int32{1, 3, 2, 3, 4, 4, 6}, "int32"}, []int32{1, 3, 2, 4, 6}},
		{"t3", args{[]string{"1", "3", "2", "3", "4", "4", "6", "中国", "中国", "中国人"}, "string"}, []string{"1", "3", "2", "4", "6", "中国", "中国人"}},
		{"t4", args{[]float32{1.0, 1.000, 2.34, 4, 56, 4.0}, "float32"}, []float32{1.0, 2.34, 4, 56}},
		{"t5", args{[]struct_A{{name: "test1"}, {name: "test1"}, {name: "test1", age: 2}, {name: "test2"}}, "struct_A"}, []struct_A{{name: "test1"}, {name: "test1", age: 2}, {name: "test2"}}}, // []struct
		{"t5", args{[]*struct_A{prtStructA, prtStructA, prtStructA2}, "*struct_A"}, []*struct_A{prtStructA, prtStructA2}},                                                                       // []*struct 切片指针，指针指向的地址不同即为不同项
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueReflect(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueReflect() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

type struct_A struct {
	name string
	age  int
}

type uniqueSlice_struct_A struct {
	origin []struct_A
	resp   []struct_A
}

func (s uniqueSlice_struct_A) Len() int {
	return len(s.origin)
}

func (s uniqueSlice_struct_A) Index(i int) interface{} {
	return s.origin[i]
}

func (s *uniqueSlice_struct_A) Push(i int) {
	s.resp = append(s.resp, s.Index(i).(struct_A))
}

func (s uniqueSlice_struct_A) Resp() interface{} {
	return s.resp
}

func Test_UniqueSlice(t *testing.T) {
	type args struct {
		origin interface{}
		eType  string
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"t1", args{[]int{1, 2, 3, 1, 3, 4}, "int"}, []int{1, 2, 3, 4}},                                                                                                                         // []int
		{"t2", args{[]string{"1", "2", "3", "1", "3", "4"}, "string"}, []string{"1", "2", "3", "4"}},                                                                                            // []string
		{"t3", args{[]struct_A{{name: "test1"}, {name: "test1"}, {name: "test2"}, {name: "test1", age: 2}}, "struct_A"}, []struct_A{{name: "test1"}, {name: "test2"}, {name: "test1", age: 2}}}, // []struct
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u Unique
			switch tt.args.eType {
			case "int":
				u = &IntUniqueSlice{origin: tt.args.origin.([]int)}
			case "string":
				u = &StringUniqueSlice{origin: tt.args.origin.([]string)}
			case "struct_A":
				u = &uniqueSlice_struct_A{origin: tt.args.origin.([]struct_A)}
			}

			got := UniqueSlice(u)
			t.Logf("UniqueSlice() = %+v", got)

			switch tt.args.eType {
			case "int":
				got = got.([]int)
				tt.want = tt.want.([]int)
			case "string":
				got = got.([]string)
				tt.want = tt.want.([]string)
			case "struct_A":
				got = got.([]struct_A)
				tt.want = tt.want.([]struct_A)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSlice() failed. got=%+v, want=%+v", got, tt.want)
			}
		})
	}
}

func Test_UniqueInts(t *testing.T) {
	type args struct {
		xi []int
	}

	tests := []struct {
		name string
		args args
		want []int
		note string
	}{
		{"t1", args{[]int{}}, []int{}, "空切片测试"},
		{"t2", args{[]int{1, 2, 2, 1, 2, 1, 1}}, []int{1, 2}, "多个重复向测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueInts(tt.args.xi)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInts() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_UniqueStrings(t *testing.T) {
	type args struct {
		xs []string
	}

	tests := []struct {
		name string
		args args
		want []string
		note string
	}{
		{"t1", args{[]string{}}, []string{}, "空切片测试"},
		{"t2", args{[]string{"1", "2", "2", "1", "2", "1", "1"}}, []string{"1", "2"}, "多个重复项测试"},
		{"t3", args{[]string{"哈哈", "go嘿哈", "皮卡皮卡", "gog", "go嘿哈", "哈哈", "1"}}, []string{"哈哈", "go嘿哈", "皮卡皮卡", "gog", "1"}, "中文多个重复项测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueStrings(tt.args.xs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrings() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_UniqueFloat64s(t *testing.T) {
	type args struct {
		xf []float64
	}

	tests := []struct {
		name string
		args args
		want []float64
		note string
	}{
		{"t1", args{[]float64{}}, []float64{}, "空切片测试"},
		{"t2", args{[]float64{1.01, 1.001, 1, 2.45, 2.4500, 1.0000, 3, 33}}, []float64{1.01, 1.001, 1, 2.45, 3, 33}, "多个重复项测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueFloat64s(tt.args.xf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueFloat64s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func Test_UniqueInt32s(t *testing.T) {
	type args struct {
		xi []int32
	}

	tests := []struct {
		name string
		args args
		want []int32
		note string
	}{
		{"t1", args{[]int32{}}, []int32{}, "空切片测试"},
		{"t2", args{[]int32{1, 2, 2, 1, 2, 1, 1}}, []int32{1, 2}, "多个重复向测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueInt32s(tt.args.xi)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInt32s() failed. got=%v, want=%v", got, tt.want)
			}
		})
	}
}
