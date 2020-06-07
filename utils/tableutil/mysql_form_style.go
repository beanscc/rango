package tableutil

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// runeDisplayLenMap utf8字符等宽显示所需的字节长度
var runeDisplayLenMap = map[int]int{
	1: 1, // 如 英文字母
	2: 1, // 如 埃塞俄比亚语
	3: 2, // 如 中文、日文
	4: 2, // 如 表情：😂
	5: 2,
	6: 2,
}

// mySQLFormStyleTable 形如 mysql 终端表格样式的终端表格
// mysql 终端表格，在等宽字体展示下，形如 excel 表格一样整体
type mySQLFormStyleTable struct{}

// NewTableLikeMySQLFormStyle return new table style like mysql form style
func NewTableLikeMySQLFormStyle() *mySQLFormStyleTable {
	return new(mySQLFormStyleTable)
}

// Show 打印显示
func (mt *mySQLFormStyleTable) Show(rows [][]string) {
	fmtRows := mt.FormatRows(rows)
	for _, row := range fmtRows {
		fmt.Println(row)
	}
}

// FormatRows 格式化每行数据
func (mt *mySQLFormStyleTable) FormatRows(rows [][]string) (fmtRows []string) {
	if len(rows) == 0 {
		return fmtRows
	}

	// 计算每列字符串显示所需的最大字节数  map[int]int
	columnsMaxDisplayLen := calColumnsMaxDisplayLen(rows)

	// 第一行数据默认为表头
	hrStr := formatHR(columnsMaxDisplayLen)
	// title 的首分割线
	fmtRows = append(fmtRows, hrStr)
	// title
	fmtRows = append(fmtRows, mt.formatRow(rows[0], columnsMaxDisplayLen))
	// title 的尾分割线
	fmtRows = append(fmtRows, hrStr)

	// 格式化除表头外的其他行数据
	for _, row := range rows[1:] {
		rowStr := mt.formatRow(row, columnsMaxDisplayLen)
		// 行数据
		fmtRows = append(fmtRows, rowStr)
		// 每行的分割线
		fmtRows = append(fmtRows, hrStr)
	}

	// 最后的行分割线
	// fmtRows = append(fmtRows, hrStr)

	return fmtRows
}

// formatRow 格式化行数据
func (mt *mySQLFormStyleTable) formatRow(row []string, columnsMaxDisplayLen map[int]int) string {
	columnData := []string{}
	for index, column := range row {
		charLen := utf8.RuneCountInString(column)
		displayLen := calStrDisplayLen(column)
		// 若该字符串等宽显示所需的字节数小于该列等宽显示所需的最大字节数，
		// 则给其 fmt 的时的字符长度 = 原字符长度 + （maxDisplayLen - displayLen）
		if displayLen < columnsMaxDisplayLen[index] {
			charLen += columnsMaxDisplayLen[index] - displayLen
		}

		columnData = append(columnData, fmt.Sprintf("%-*s", charLen, column))
	}

	return fmt.Sprintf("| %s |", strings.Join(columnData, " | "))
}

// formatHR 格式化行与行之间的分割线
func formatHR(columnsMaxDisplayLen map[int]int) string {
	var buf bytes.Buffer
	buf.WriteByte('+')
	for i := 0; i < len(columnsMaxDisplayLen); i++ {
		buf.WriteString(strings.Repeat("-", columnsMaxDisplayLen[i]+2))
		buf.WriteRune('+')
	}

	return buf.String()
}

// calColumnsMaxDisplayLen 计算多行数据每列数据显示的最大字节长度
func calColumnsMaxDisplayLen(rows [][]string) map[int]int {
	columnsMaxDisplayLen := make(map[int]int, 0)
	for _, row := range rows {
		for index, column := range row {
			cdLen := calStrDisplayLen(column)
			if _, ok := columnsMaxDisplayLen[index]; !ok {
				columnsMaxDisplayLen[index] = cdLen
			} else {
				if columnsMaxDisplayLen[index] < cdLen {
					columnsMaxDisplayLen[index] = cdLen
				}
			}
		}
	}

	return columnsMaxDisplayLen
}

// calStrDisplayLen 计算字符串等宽显示需要的字节长度
func calStrDisplayLen(s string) int {
	tmp, l := s, 0
	for {
		if len(tmp) == 0 {
			break
		}
		_, i := utf8.DecodeRuneInString(tmp)
		if charDisplayLen, ok := runeDisplayLenMap[i]; ok {
			l += charDisplayLen
		} else {
			l += i
		}

		tmp = tmp[i:]
	}

	return l
}
