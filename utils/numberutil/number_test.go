package numberutil

import (
	"fmt"
	"testing"
)

func ExampleEqualFloat64() {
	type args struct {
		x float64
		y float64
	}

	data := []args{
		{0.2, 0.2},
	}

	for _, v := range data {
		fmt.Printf("%-5t, %-5t, %-5t, ", v.x == v.y, EqualFloat64(v.x, v.y, -1), EqualFloat64(v.x, v.y, 0.000000000001))
		fmt.Println(v.x, v.y)
	}

	// 输出结果如下：
	// true  true  true  0.2 0.2
	// false false false 0.6 0.61
	// false false true  0.799 0.8
}

func TestEqualFloat64(t *testing.T) {
	type args struct {
		x float64
		y float64
	}

	data := []args{
		{0.2, 0.2},
		{0.6, 0.61},
		{0.799, 0.8},
	}

	for _, v := range data {
		fmt.Printf("%-5t %-5t %-5t ", v.x == v.y, EqualFloat64(v.x, v.y, -1), EqualFloat64(v.x, v.y, 0.01))
		fmt.Println(v.x, v.y)
	}
}

func TestEqualFloat64Prec(t *testing.T) {
	type args struct {
		x float64
		y float64
	}

	data := []args{
		{0.2, 0.2},
		{0.6, 0.61},
		{0.799, 0.8},
	}

	for _, v := range data {
		fmt.Printf("%-5t %-5t %-5t %-5t ", v.x == v.y, EqualFloat64(v.x, v.y, -1), EqualFloat64(v.x, v.y, 0.01), EqualFloat64Prec(v.x, v.y, 6))
		fmt.Println(v.x, v.y)
	}
}

func TestRound(t *testing.T) {
	testData := []struct {
		name  string
		input float64
		want  float64
	}{
		{"t1", 1.0, 1},
		{"t2", 1.00, 1},
		{"t3", 1.1, 1.1},
		{"t4", 1.0011, 1.001},
		{"t5", 1.0015, 1.002},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			r := Round(d.input, 3)
			if r != d.want {
				t.Errorf("t: %v, Round to %v failed. input=%v, want=%v, ret=%v", d.name, 3, d.input, d.want, r)
			}

			t.Logf("t: %v, Round to %v success. ret: %v, want: %v, ret=%v", d.name, 3, r, d.want, r)
		})
	}
}
