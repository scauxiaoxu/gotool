package util

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

// 获取当前平台 操作系统 win or linux
func TestTime(t *testing.T) {
	// 自己看控制台输出有无异常
	os := Os()
	fmt.Println(os)
}

// 判断当前 文件/文件夹是否存在
func TestFileExists(t *testing.T) {
	// 测试文件
	_, file, _, _ := runtime.Caller(0)
	re := FileExists(file)
	if re != true {
		t.FailNow()
	}
	// 测试文件夹
	re = FileExists(path.Dir(file))
	if re != true {
		t.FailNow()
	}
}