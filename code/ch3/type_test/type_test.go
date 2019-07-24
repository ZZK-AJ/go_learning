package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

// 类型转换一定要显式
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// 如果直接 a=b 就是隐式类型转换
	b = int64(a)
	var c MyInt
	c = MyInt(b)  // type MyInt int64 利用别名做类型转换
	t.Log(a, b, c)
}

// 不支持指针运算
func TestPoint(t *testing.T) {
	a := 1			//整形边量 a 测试指针类型，但是指针类型不支持运算
	aPtr := &a		// a 的指针
	//aPtr = aPtr + 1	// 指针的运算时不允许的，不能用指针的运算去获取连续的变量 如数组
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)	// 用 %T 获取边量类型
}

func TestString(t *testing.T) {
	var s string	// 默认被初始化为空字符串
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))
	if s == ""{
		fmt.Println(" empty ")
	}
}
