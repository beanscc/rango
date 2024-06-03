package fileutil

import (
	"fmt"
	"math"
	"os"
)

// FormatFileSize 以易于阅读的格式输出文件大小，如 1KB 234MB
func FormatFileSize(byteSize int64) string {
	if byteSize == 0 {
		return "0B"
	}

	// 单位
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	exp := math.Log(float64(byteSize)) / math.Log(1024)
	exp = math.Min(exp, float64(len(units)-1))

	filesize := float64(byteSize) / math.Pow(1024, exp)
	return fmt.Sprintf("%.2f%s", filesize, units[int(exp)])
}

// IsExist 判断 filename 是否路径/文件是否存在，以及 filename 是否文件夹
func IsExist(filename string) (exist bool, isDir bool, err error) {
	f, err := os.Stat(filename)
	if err == nil {
		return true, f.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, false, nil
	}

	// 其他错误
	return false, false, err
}
