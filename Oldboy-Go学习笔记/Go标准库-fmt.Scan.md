[TOC]

本文介绍了Go语言中 `fmt` 包中从标准输入获取数据的的 `Scan` 系列函数、从 `io.Reader` 中获取数据的 `Fscan` 系列函数以及从字符串中获取数据的 `Sscan` 系列函数的用法。

# Scan系列

Go语言 `fmt` 包下有 `fmt.Scan` 、 `fmt.Scanf` 、 `fmt.Scanln` 三个函数，可以在程序运行过程中从标准输入获取用户的输入。

## fmt.Scan

### 语法

```go
func Scan(a ...interface{}) (n int, err error)
```

- Scan 从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

### 代码示例

```go
func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

将上面的代码编译后在终端执行，在终端依次输入 `小王子` 、`28 `和 `false` 使用空格分隔。

```bash
$ ./scan_demo 
小王子 28 false
扫描结果 name:小王子 age:28 married:false 
```

`fmt.Scan `从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。



## fmt.Scanf

### 语法

```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

- Scanf 从标准输入扫描文本，根据 format 参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

### 代码示例

```go
func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

将上面的代码编译后在终端执行，在终端按照指定的格式依次输入`小王子`、`28` 和 `false`。

```bash
$ ./scan_demo 
1:小王子 2:28 3:false
扫描结果 name:小王子 age:28 married:false 
```

`fmt.Scanf` 不同于 `fmt.Scan` 简单的以空格作为输入数据的分隔符，`fmt.Scanf `为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

例如，我们还是按照上个示例中以空格分隔的方式输入，`fmt.Scanf` 就不能正确扫描到输入的数据。

```bash
$ ./scan_demo 
小王子 28 false
扫描结果 name: age:0 married:false 
```



## fmt.Scanln

### 语法

```go
func Scanln(a ...interface{}) (n int, err error)
```

- `Scanln` 类似 `Scan`，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

### 代码示例

```go
func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

将上面的代码编译后在终端执行，在终端依次输入`小王子`、`28`和`false`使用空格分隔。

```bash
$ ./scan_demo 
小王子 28 false
扫描结果 name:小王子 age:28 married:false 
```

==`fmt.Scanln` 遇到回车就结束扫描了，这个比较常用。==



# Fscan系列

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
```

这几个函数功能分别类似于 `fmt.Scan`、`fmt.Scanf`、`fmt.Scanln` 三个函数，只不过它们不是从标准输入中读取数据而是从 `io.Reader` 中读取数据。



# Sscan系列

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```

这几个函数功能分别类似于 `fmt.Scan`、`fmt.Scanf`、`fmt.Scanln` 三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。

