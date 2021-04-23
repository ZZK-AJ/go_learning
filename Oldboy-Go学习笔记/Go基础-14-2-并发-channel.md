[TOC]

不同的 `goroutine` 之间如何进行通信？

1. 全局变量+锁同步
2. `channel`



# channel

==单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。==

虽然可以使用共享内存进行数据交换，但是共享内存在不同的 `goroutine` 中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

==go 语言的并发模型是 CSP，提倡通过通信共享内存而不是通过共享内存而实现通信。==

如果说 `goroutine` 是 Go 程序并发的执行体，`channel` 就是它们之间的连接。`channel ` 是可以让一个 `goroutine` 发送特定值到另一个 `goroutine` 的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。==channel 像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序==。

==每一个通道都是一个具体类型的导管，也就是声明 channel 的时候需要为其指定元素类型。==



`Channel` 概念：

1. 类似 `unix` 的管道(`pipe`)
2. 先进先出
3. 线程安全，多个 `goroutine` 同时访问，不需要加锁
4. chan 是有类型的，整数的 `channel` 只能存放整数



## 声明channel

声明通道类型的格式如下：

```go
var 变量 chan 元素类型
```

举几个例子：

```go
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递 int 切片的通道
```

![image-20190616180313596](/Users/zzk/Library/Application Support/typora-user-images/image-20190616180313596.png)

![image-20190616180334287](/Users/zzk/Library/Application Support/typora-user-images/image-20190616180334287.png)





## 创建channel

通道是==引用类型==，通道类型的空值是 `nil`。

```go
var ch chan int
fmt.Println(ch) // <nil>
```

==声明的通道后需要使用 `make` 函数初始化之后才能使用==。 

创建 channel 的格式如下：

```go
make(chan 元素类型, [缓冲大小])  // 缓冲大小是可选的。
```

举几个例子：

```go
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
```



## channel操作

通道有发送（send）、接收 (receive）和关闭（close）三种操作。

发送和接收都使用 `<-` 符号。

现在我们先使用以下语句定义一个通道：

```go
ch := make(chan int)
```



如果希望 channel 仅发送数据，则必须在 channel 之后使用 `<-` 运算符。

如果希望 channel 接收数据，则必须在 channel 之前使用 `<-` 运算符。

### 发送

将一个值发送到通道中。

```go
ch <- 10 // 通过 channel ch 发送一个值 10
```

### 接收

从一个通道中接收值。

```go
x := <- ch // x 接收一个发送到 channel ch 的值

<-ch       // 从ch中接收值，忽略结果
```



![image-20190616180429343](/Users/zzk/Library/Application Support/typora-user-images/image-20190616180429343.png)

注意 1 里面的 = 一定要有，表示从 chan 里面读取数据，然后赋值



### 关闭

我们通过调用内置的 `close` 函数来关闭通道。

```go
close(ch)
```

关于关闭通道需要注意的事情是，只有在通知接收方 goroutine 所有的数据都发送完毕的时候才需要关闭通道。==通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的==。



关闭后的通道有以下特点：

1. 对一个关闭的通道再发送值就会导致 panic 。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致 panic 。



## 无缓冲的通道

无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：

```go
func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
```

上面这段代码能够通过编译，但是执行的时候会出现以下错误：

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../src/github.com/Q1mi/studygo/day06/channel02/main.go:8 +0x54
```

为什么会出现 `deadlock` 错误呢？

因为我们==使用 `ch := make(chan int)` 创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值==。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。

==上面的代码会阻塞在 `ch <- 10` 这一行代码形成死锁==，那如何解决这个问题呢？

一种方法是启用一个 `goroutine` 去接收值，例如：

```go
package main

import (
	"fmt"
)

func recv(c chan int) {		// 参数为 c 这个 chan int 类型的通道
	ret := <-c			// 从 channel 中读取数据
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)  // 直接使用 make 函数就可以初始化通道
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10	// 发送数据到 channel
	fmt.Println("发送成功")
}

/*
发送成功
接收成功 10
 */
```

无缓冲通道上的发送操作会阻塞，直到另一个 `goroutine` 在该通道上执行接收操作，这时值才能发送成功，两个 `goroutine` 将继续执行。相反，如果接收操作先执行，接收方的 goroutine 将阻塞，直到另一个 `goroutine` 在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为 `同步通道` 。



## 有缓冲的通道

解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在==使用 make 函数初始化通道的时候为其指定通道的容量==，例如：

```go
func main() {
	ch := make(chan int, 5) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println(len(ch)) // 通道内元素的数量为 3
	fmt.Println(cap(ch)) // 通道的容量为 5
	fmt.Println("发送成功")
}

/*
3
5
发送成功
 */
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

我们可以使用内置的 `len` 函数获取通道内元素的数量，使用 `cap` 函数获取通道的容量。

![image-20190616180640133](/Users/zzk/Library/Application Support/typora-user-images/image-20190616180640133.png)



## 如何优雅的从通道循环取值

当通过通道发送有限的数据时，我们可以通过  `close`  函数关闭通道来告知从该通道接收值的  `goroutine`  停止等待。==当通道被关闭时，往该通道发送值会引发 panic，从该通道里接收的值一直都是类型零值。==

那==如何判断一个通道是否被关闭==了呢？

我们来看下面这个例子：

```go
// channel 练习
func main() {
	ch1 := make(chan int)		// 无缓冲的通道，为什么后面的可以实现？
	ch2 := make(chan int)
  
	// 开启 goroutine 将 0~100 的数发送到 ch1 中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
  
	// 开启 goroutine 从 ch1 中接收值，并将该值的平方发送到 ch2 中
	go func() {
		for {
			i, ok := <- ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
  
	// 重点：在主 goroutine 中从 ch2 中接收值打印
	for i := range ch2 {
		fmt.Println(i)
	}
}

/*
0
1
4
9
16
25
36
49
64
81
 */
```

从上面的例子中我们看到==有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是 `for range` 的方式。==



## 单向通道

有的时候我们会将 `channel` 作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用`channel`都会对其进行限制，比如只能发送或只能接收。

Go 语言中提供了**单向通道**来处理这种情况。

![image-20190616180937550](/Users/zzk/Library/Application Support/typora-user-images/image-20190616180937550.png)

例如，我们把上面的例子改造如下：

```go
func counter(out chan<- int) {		// chan<- int 是一个只能发送的通道(<- 可以看到只发送到int)
	for i := 0; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {		// <-chan int 是一个只能接收的通道
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

/*
0
1
4
9
16
25
36
49
64
81
100
 */
```

其中，`out chan<- int` 是一个只能发送的通道，可以发送但是不能接收；`in <-chan int`是一个只能接收的通道，可以接收但是不能发送。两个的区别在于， `<-` 在 `chan` 前面就是只能接收，在后面就是只能发送。

==在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的==。



## 练习题

编写代码实现一个计算随机数的每个位置数字之和的程序，要求使用 `goroutine` 和 `channel` 构建生产者和消费者模式，可以指定启动的 goroutine 数量 –workerpool 模式。

在工作中我们通常会使用 `workerpool` 模式，控制 `goroutine` 的数量，防止 `goroutine` 泄漏和暴涨。



## select多路复用

在某些场景下我们==需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞==。你也许会写出如下代码使用遍历的方式来实现：

```go
for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，==Go 内置了 `select` 关键字，可以  `同时响应多个通道的操作 `== 。

 `select`  的使用类似于 ` switch ` 语句，它有一些列 `case`  分支和一个默认的分支。==每个 `case`  会对应一个通道的通信（接收或发送）过程。`select `会一直等待，直到某个 `case` 的通信操作完成时，就会执行 `case` 分支对应的语句==。

具体格式如下：

```go
select{
    case <-ch1:
        ...
    case data := <-ch2:
        ...
    case ch3<-data:
        ...
    default:
        默认操作
}
```

举个小例子来演示下 `select` 的使用：

```go
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

/*
0
2
4
6
8
这里的输出为什么是 0 2 4 6 8 ？
 */
```

使用 `select` 语句能提高代码的可读性。如果多个 `case` 同时满足， `select` 会随机选择一个。对于没有 `case` 的 `select{}` 会一直等待。


