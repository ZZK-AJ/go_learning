package fib

import (
	"testing"
)


// 测试方法，以大小的 Test 开始
func TestFibList(t *testing.T) {
	// 定义变量 a b
	// var a int = 1
	// var b int = 1

	// var (
	// 	a int = 1
	// 	b     = 1
	// )
	a := 1  // 更快的方式，使用类型推断
	// a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a  //在一个赋值语句中，可以对多个进行赋值
	t.Log(a, b)
	//fmt.Println(a,b)
}
