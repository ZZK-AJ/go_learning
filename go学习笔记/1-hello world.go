package main
// package main 定义了包名。
// 你必须在源文件中非注释的第一行指明这个文件属于哪个包
// 如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包

import "fmt"
/* 
import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素）
fmt 包实现了格式化 IO（输入/输出）的函数
*/

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello"+"World!")  // Go 语言的字符串可以通过 + 实现
}

// 命令行 go run hello.go 执行代码
// 使用 go build 命令来生成二进制文件
// $ go build hello.go 
// $ ls
// hello    hello.go
// $ ./hello 
// Hello, World!

// 注意的是 { 不能单独放在一行
// Go 程序中，一行代表一个语句结束。
// 每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾

/*
Go 语言中变量的声明必须使用空格隔开，如：
var age int;
*/