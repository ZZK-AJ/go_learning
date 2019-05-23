package constant_test

import "testing"

/*
编写测试程序
源码文件以_test结尾，xxx_test.go
测试方法名以Test开头：
*/

const (
	Monday = 1 + iota  //对于连续常量的赋值
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota  //连续位常量的赋值
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
