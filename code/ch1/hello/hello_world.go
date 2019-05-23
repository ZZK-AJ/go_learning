package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1], os.Args[2])  // os.Args获取命令行参数
	}
}

// 一定要使用 package main 但是对目录名没有什么要求