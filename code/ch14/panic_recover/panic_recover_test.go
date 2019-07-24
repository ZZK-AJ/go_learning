package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

/*
panic 用于不可恢复的错误
panic 退出后，会执行 defer 指定的内容

os.Exit 退出不会调用 defer
*/

func TestPanicVxExit(t *testing.T) {

	// 可以看到可以打印出 defer 执行的内容
	defer func() {
		fmt.Println("Finally!")
	}()

	// 类似 try ... catch ...
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	// 要避免出错误之后，只是单纯的 recover 了，而没有去恢复，形成僵尸进程
	// 所以，let it crash，往往是我们恢复不确定性错误的最好方法
	// 就让程序崩溃，然后让守护进程重启服务来恢复

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}
