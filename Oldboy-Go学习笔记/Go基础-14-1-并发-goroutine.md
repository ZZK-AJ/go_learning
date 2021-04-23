[TOC]

并发是编程里面一个非常重要的概念，Go语言在语言层面天生支持并发，这也是Go语言流行的一个很重要的原因。

# Go语言中的并发编程

## 并发与并行

==并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）==。

==并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）==。

==Go语言的并发通过 `goroutine` 实现。`goroutine` 类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个 `goroutine` 并发工作。==

==`goroutine `是由Go语言的运行时调度完成，而线程是由操作系统调度完成。==

==Go语言还提供 `channel` 在多个 `goroutine` 间进行通信。`goroutine` 和 `channel `是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础。==



# goroutine

在 Java/C++ 中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务和然后自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？ Go 语言中的 goroutine 就是这样一种机制，goroutine 的概念类似于线程，但 ==goroutine 由 Go 程序运行时的调度和管理==。Go 程序会智能地将 goroutine 中的任务合理地分配给每个 CPU。

==Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制==。

==在 Go 语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能 goroutinue，当你需要让某个任务并发执行的时候，你只需要起一个 goroutinue 就可以了，就是这么简单粗暴。==



## 使用goroutine

Go 程序中使用 `go` 关键字为一个函数创建一个 `goroutine`。一个函数可以被创建多个  `goroutine`，一个 `goroutine`必定对应一个函数。

### 启动单个goroutine

==启动 goroutine 的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个 `go` 关键字==。

举个例子如下：

```go
func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	hello()
	fmt.Println("main goroutine done!")
}
```

这个示例中 hello 函数和下面的语句是串行的，执行的结果是打印完 `Hello Goroutine!` 后打印 `main goroutine done!` 。

接下来我们==在调用 hello 函数前面加上关键字 `go`==，也就是启动一个 goroutine 去执行 hello 这个函数。

```go
func main() {
	go hello() // 启动另外一个 goroutine 去执行 hello 函数
	fmt.Println("main goroutine done!")
}
```

这一次的执行结果只打印了 `main goroutine done!` ，并没有打印 `Hello Goroutine!` 。为什么呢？

==在程序启动时，Go程序就会为 `main()` 函数创建一个默认的 `goroutine` 。== 当 main() 函数返回的时候该 `goroutine `就结束了，所有在 `main()` 函数中启动的 `goroutine` 会一同结束，`main` 函数所在的 `goroutine` 就像是权利的游戏中的夜王，其他的 `goroutine` 都是异鬼，夜王一死它转化的那些异鬼也就全部 GG 了。

所以我们要想办法让 main 函数等一等hello函数，最简单粗暴的方式就是 Sleep 了。

```go
func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}
```

执行上面的代码你会发现，这一次先打印 `main goroutine done!` ，然后紧接着打印 `Hello Goroutine!` 。

==首先为什么会先打印 `main goroutine done!` 是因为我们在创建新的goroutine的时候需要花费一些时间，而此时 mian 函数所在的 `goroutine` 是继续执行的。==



### sync.WaitGroup

代码中生硬的使用 `time.Sleep` 肯定是不合适的，==Go语言中可以使用 `sync.WaitGroup` 来实现并发任务的同步。==

 `sync.WaitGroup` 有以下几个方法：

|             方法名              |        功能         |
| :-----------------------------: | :-----------------: |
| (wg * WaitGroup) Add(delta int) |    计数器+delta     |
|     (wg *WaitGroup) Done()      |      计数器-1       |
|     (wg *WaitGroup) Wait()      | 阻塞直到计数器变为0 |

`sync.WaitGroup ` 内部维护着一个计数器，计数器的值可以增加和减少。

例如当我们启动了 N 个并发任务时，就将计数器值增加 N。每个任务完成时通过调用 Done() 方法将计数器减1。通过调用 Wait() 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。

(==因此` sync.WaitGroup ` 就是通过内部维护的计数器来保证并发任务全部执行完 ==)

我们利用 `sync.WaitGroup` 将上面的代码优化一下：

```go
var wg sync.WaitGroup

func hello() {
	defer wg.Done()		// 任务完成时通过调用 Done() 方法将计数器减1 
  // Go 语言的 defer 会在当前函数返回前执行传入的函数（如这里的 wg.Done()），它会经常被用于关闭文件描述符、关闭数据库连接以及解锁资源，用于保证函数能够被执行。
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()	// 阻塞直到计数器变为0
}

//Hello Goroutine!
//main goroutine done!
```

==需要注意 `sync.WaitGroup` 是一个结构体，传递的时候要传递指针==。



### 启动多个goroutine

在Go语言中实现并发就是这样简单，我们还可以启动多个 `goroutine` 。让我们再来一个例子：

```go
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()	// 阻塞直到技术器为0
}

/*
Hello Goroutine! 2
Hello Goroutine! 9
Hello Goroutine! 5
Hello Goroutine! 6
Hello Goroutine! 7
Hello Goroutine! 8
Hello Goroutine! 0
Hello Goroutine! 3
Hello Goroutine! 1
Hello Goroutine! 4
 */
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个 `goroutine` 是并发执行的，而 `goroutine` 的调度是随机的。



## goroutine与线程

### 可增长的栈

OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个 `goroutine` 的栈在其生命周期开始时只有很小的栈（典型情况下2KB），`goroutine` 的栈不是固定的，他可以按需增大和缩小，`goroutine ` 的栈大小限制可以达到1GB，虽然极少会用到这么大。所以==在Go语言中一次创建十万左右的 `goroutine` 也是可以的。==



### goroutine调度

OS 线程是由 OS 内核来调度的，`goroutine` 则是由 ==Go 运行时（runtime）自己的调度器调度的==，这个调度器==使用一个称为 m:n 调度的技术==（复用/调度 m 个 goroutine 到 n 个 OS 线程）。goroutine 的调度不需要切换内核语境，所以==调用一个 goroutine 比调度一个线程成本低很多==。



### GOMAXPROCS

==Go运行时的调度器使用 `GOMAXPROCS` 参数来确定需要使用多少个 OS 线程来同时执行 Go 代码==。默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，调度器会把 Go 代码同时调度到 8 个 OS 线程上（ GOMAXPROCS 是m:n 调度中的 n ）。

Go 语言中可以通过 `runtime.GOMAXPROCS()` 函数设置当前程序并发时占用的 CPU 逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。==Go1.5版本之后，默认使用全部的 CPU 逻辑核心数。==

我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子：

```go
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

/*
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
A: 9
B: 1
B: 2
B: 3
B: 4
B: 5
B: 6
B: 7
B: 8
B: 9
 */
```

两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。 



将逻辑核心数设为2，此时两个任务并行执行，代码如下。

```go
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
```





==Go 语言中的操作系统线程和 goroutine 的关系==：

1. 一个操作系统线程对应用户态多个 goroutine。
2. go 程序可以同时使用多个操作系统线程。
3. goroutine 和 OS 线程是多对多的关系，即 m:n。




