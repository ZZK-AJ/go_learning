package operator_test

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota  // 就是二进制移位
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)  长度不同不能比较
	t.Log(a == d)
}

// 按位置零 &^ 1 会置零，0 保持原来的
func TestBitClear(t *testing.T) {
	fmt.Println(Readable,Writable,Executable)
	a := 7 //0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

