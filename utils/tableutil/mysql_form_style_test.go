package tableutil

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// go test -v -run Test_Show
func Test_Show(t *testing.T) {
	type args struct {
		data [][]string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"t1",
			args{
				[][]string{
					[]string{"第一列title", "サブセクター第二列title", "第三列title", "第四列title"},
					[]string{"Волим да тече М"},
					[]string{"c1c3fgrt", "c22c2cc2c2c2", "c333333第三列", "第四列title"},
					[]string{"👌第一列title"},
				},
			},
		},
		{
			"t2",
			args{
				// | id | mobile      | province                    | city                     | area_code |
				// +----+-------------+-----------------------------+--------------------------+-----------+
				// | 1  | 18710842353 | 陕西first                   | 汉中first                | 34        |
				// | 2  | 18192677330 | first                       | second                   | 35        |
				// | 3  | 18710842359 | Волим да тече М             | サブセクター第二         | 56        |
				// +----+-------------+-----------------------------+--------------------------+-----------+

				[][]string{
					[]string{"id", "mobile", "province", "city", "area_code"},
					[]string{"134", "13400000000", "province", "北京", "010"},
					[]string{"4", "17800000000", "province first", "东京", "2200"},
					[]string{"100001", "18900000000", "Волим да тече М", "サブセクター第二 city", "56"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewTableLikeMySQLFormStyle().Show(tt.args.data)
		})
	}
}

func Test_calStrDisplayLen(t *testing.T) {
	tests := []string{
		"first",
		"Birst",
		"第一列title",
		"👌第一列title",
		"Волим да тече М",
		"サブセクター第二列title",
	}

	for _, tt := range tests {
		l := calStrDisplayLen(tt)
		t.Logf("display len=%v", l)
	}
}

func Test_CalCharLen(t *testing.T) {
	tests := []string{
		"first",
		"Birst",
		"第一列title",
		"👌第一列title",
		"Волим да тече М",
		"サブセクター第二列title",
	}

	// utf8字符等宽显示所需的字节长度
	runDisplayLenMap := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 2,
	}

	for _, tt := range tests {
		tmp := tt

		l := 0
		for {
			if len(tmp) == 0 {
				break
			}
			d, i := utf8.DecodeRuneInString(tmp)
			if charDisplayLen, ok := runDisplayLenMap[i]; ok {
				l += charDisplayLen
			} else {
				l += i
			}

			t.Logf("i=%v, d=%c, len=%v", i, d, l)

			tmp = tmp[i:]
		}
	}
}

func Test_fmt(t *testing.T) {
	tests := []string{
		"Bst first",
		"中国中国中",
		"汉中first",
	}

	// | 汉中first   |
	for _, tt := range tests {
		// fmt.Printf("fmt:%-*s | %v\n", len(tt), tt, len(tt))
		charLen := len([]rune(tt))
		fmt.Printf("fmt:%-*s | char:%v, byte:%v, displayLen=%v\n", charLen, tt, charLen, len(tt), calStrDisplayLen(tt))
		// fmt.Printf("fmt:%-7s | %v\n", tt, len([]rune(tt)))
	}
}

// `
// Menlo, Monaco, 'Courier New', monospace
// monospace, Monaco, 'Courier New', Menlo
// +-----------------+-------------------------+---------------+-------------+
// | 第一列title     | サブセクター第二列title | 第三列title   | 第四列title |
// +-----------------+-------------------------+---------------+-------------+
// | Волим да тече М |
// +-----------------+-------------------------+---------------+-------------+
// | c1c3fgrt        | c22c2cc2c2c2            | c333333第三列 | 第四列title |
// +-----------------+-------------------------+---------------+-------------+
// | 👌😂            |
// +-----------------+-------------------------+---------------+-------------+

// +----+-------------+-----------------------------+--------------------------+-----------+
// | id | mobile      | province                    | city                     | area_code |
// +----+-------------+-----------------------------+--------------------------+-----------+
// | 1  | 18710842353 | 陕西first                   | 汉中first                | 34        |
// | 2  | 18192677330 | first                       | second                   | 35        |
// | 3  | 18710842359 | Волим да тече М             | サブセクター第二         | 56        |
// +----+-------------+-----------------------------+--------------------------+-----------+
// `
