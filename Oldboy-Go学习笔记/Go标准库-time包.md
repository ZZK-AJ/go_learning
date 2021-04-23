[TOC]

时间和日期是我们编程中经常会用到的，本文主要介绍了 Go 语言内置的 time 包的基本用法。

## Go语言中导入包

Go语言中使用 `import` 关键字导入包，包的名字使用双引号（”）包裹起来。

### 单行导入

```go
import "time"
import "fmt"
```

### 多行导入

导入多个包时可以使用圆括号，包名的顺序不影响导入效果，例如：

```go
import (
    "fmt"
    "time"
)
```

需要注意的是，Go语言中如果存在导入但没有使用的包，会发生编译错误。（也就是导入的包必须是在代码中用到的包。）



## time包

### 时间类型

`time.Time ` 类型表示时间。

```go
func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

/*
current time:2019-06-03 17:41:10.28374 +0800 CST m=+0.000213723
2019
June
3
17
2019-06-03 17:41:10
*/
```



### 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。

```go
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}
```

使用 `time.Unix()` 函数将时间戳转为时间格式。

```go
func main(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// current timestamp1:1559554944
// current timestamp2:1559554944248349000
```



## 定时器

使用 `time.Tick(时间间隔)` 来设置定时器。

```go
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}
```



## 时间间隔

Duration 类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年。 定义的时间间隔常量如下：

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

例如：`time.Duration `表示1纳秒，`time.Second `表示1秒。



### 时间加时间间隔

我们在日常的编码过程中可能会遇到要求 时间+时间间隔 的需求，Go 语言的时间对象有提供 `Add` 方法如下：

```go
func (t Time) Add(d Duration) Time
```

举个例子，求一个小时之后的时间：

```go
func main() {
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)
}
```



### 两个时间相减

求两个时间之间的差值：

```go
func (t Time) Sub(u Time) Duration
```

返回一个时间段 t-u。如果结果超出了 Duration 可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用 t.Add(-d)。



## 时间比较

Go语言的time包提供了时间对象之间比较的函数，如下：

### Equal

```go
func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

### Before

```go
func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真；否则返回假。

### After

```go
func (t Time) After(u Time) bool
```

如果t代表的时间点在u之后，返回真；否则返回假。



## 时间格式化

时间类型有一个自带的方法 `Format` 进行格式化，需要注意的是 Go 语言中格式化时间模板不是常见的 `Y-m-d H:M:S` 而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。

补充：如果想格式化为12小时方式，需指定 `PM`。

```go
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

/*
2019-06-03 17:58:56.083 Mon Jun
2019-06-03 05:58:56.083 PM Mon Jun
2019/06/03 17:58
17:58 2019/06/03
2019/06/03
*/
```

练习题：

1. 获取当前时间，格式化输出为2017/06/19 20:30:05`格式。
2. 编写程序统计一段代码的执行耗时时间，单位精确到微秒。

