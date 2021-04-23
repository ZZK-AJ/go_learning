package loop_test

import (
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	// 无限循环
	//for {
	//	fmt.Println(n)
	//	n++
	//}
}


