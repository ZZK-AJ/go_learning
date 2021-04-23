package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 1 {
		fmt.Println("Hello World", os.Args[1])  // os.Args获取命令行参数,因为在main函数中不支持传入参数
		os.Exit(0)
	}
	//os.Exit(123)
}

// 一定要使用 package main 但是对目录名没有什么要求

/* 
通过不断的实践，快速试错，大师错误的次数远远比菜鸟尝试的次数多
*/

// 在 main 函数里面不支持返回值，可以用 os.exit 来退出