package condition_test

import "testing"

/*
if condition1 {
} else if c2 {
} else {
}
condition 表达式结果必须是 bool
支持变量赋值
if var declaration；condition {
}
*/

func TestSwitchMultiCase(t *testing.T) {
	// 因为函数支持多返回值，所以可以在 if 里面判断
	//if v,err := someFun(); err == nil {
	//	// do something
	//	t.Log("a:",v)
	//} else {
	//	t.Log(err)
	//}

	// 在一个 case 被命中之后，会被 break 掉
	for i := 0; i < 5; i++ {
		// 判断 i 的值
		switch i {
		// case 后可以支持多项
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 把 switch 当做 if else 来使用
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("Unknow")
		}
	}
}
