package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	// 类型转换一定要显式
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)  // type MyInt int64 利用别名做类型转换
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1			//测试指针类型，但是指针类型不支持运算
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))

}
