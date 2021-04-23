package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值函数测试
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// timeSpent 计算函数操作时长的函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 3)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	// 测试函数式编程，函数作为参数，装饰器功能
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 可变长参数 ...+参数类型，都会把参数转换为一个数组，然后通过遍历来完成
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
//Clear resources. 在 panic 之后也会执行，可以用来释放一些资源