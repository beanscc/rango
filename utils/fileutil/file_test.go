package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileExist(t *testing.T) {
	// 测试不存在问题
	filename := filepath.Join("/tmp", fmt.Sprintf("test_not_exist_%d", time.Now().UnixNano()))
	t.Logf("filename: %s", filename)
	exist, isDir, err := FileExist(filename)
	if exist || isDir || err != nil {
		t.Errorf("file:%s should not exist err: %v", filename, err)
	}

	// 创建
	fd, err := os.Create(filename)
	if err != nil {
		t.Errorf("create file err: %v", err)
		return
	}
	defer func() {
		fd.Close()
		os.Remove(filename)
	}()

	fd.WriteString("test")
	fd.Sync()

	exist, isDir, err = FileExist(filename)
	if !exist || isDir || err != nil {
		t.Errorf("file:%s should exist err: %v", filename, err)
	}

	// 目前测试
	testDir := filepath.Join("/tmp", "test_dir")
	defer os.Remove(testDir)
	// not exits dir
	exist, isDir, err = FileExist(testDir)
	if exist || isDir || err != nil {
		t.Errorf("dir: %s should not exist err: %v", testDir, err)
	}

	if err := os.MkdirAll(testDir, 0644); err != nil {
		t.Errorf("mkdir err: %v", err)
		return
	}

	// exist dir
	exist, isDir, err = FileExist(testDir)
	if !exist || !isDir || err != nil {
		t.Errorf("dir: %s should exist err: %v", testDir, err)
	}
}
