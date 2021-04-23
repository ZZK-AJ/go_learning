[TOC]

Go语言中提供的映射关系容器为 `map` ，其内部使用 `散列表（hash）` 实现。

# map

map 是一种==无序的==基于 `key-value` 的数据结构，Go语言中的 ==map 是引用类型，必须初始化才能使用==。

## map定义

Go语言中  `map` 的==定义语法==如下：

```go
map[KeyType]ValueType
```

其中，

- KeyType : 表示键的类型。
- ValueType : 表示键对应的值的类型。

map 类型的变量==默认初始值为 nil== ，需要使用 make() 函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中 cap 表示 map 的容量，该参数虽然不是必须的，但是我们应该在初始化 map 的时候就为其指定一个合适的容量。

## map基本使用

map中的数据都是成对出现的，map的基本使用示例代码如下：

```go
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of scoreMap: %T\n", scoreMap)
}
```

输出：

```bash
map[小明:100 张三:90]
100
type of scoreMap: map[string]int
```

map 也支持在声明的时候填充元素，例如：

```go
func main() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo) //
}
```



## 判断某个键是否存在

Go 语言中有个判断 map 中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

举个例子：

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)		// map[小明:100 张三:90]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三zk"]	// 90 true
	if ok {
		fmt.Println(v,ok)	// 0 false
	} else {
		fmt.Println(v,ok)	// 0
		fmt.Println("查无此人")	// 查无此人
	}
}
```



## map的遍历

Go语言中使用 `for range` 遍历map。

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
```

但我们只想遍历key的时候，可以按下面的写法：

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```

**注意：** ==遍历map时的元素顺序与添加键值对的顺序无关。==



## 使用delete()函数删除键值对

使用 `delete()` 内建函数从 map 中根据 key 值删除一组键值对，`delete()` 函数的格式如下：

```go
delete(map, key)
```

其中，

- map : 表示要删除键值对的 map
- key : 表示要删除的键值对的键

示例代码如下：

```go
func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小明")//将小明:100从map中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}
}

//张三 90
//娜扎 60
```

## 按照指定顺序遍历map

```go
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)
	fmt.Println(scoreMap)	// map[]

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu %02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value			// 给scoreMap赋值
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)  // 这里是初始化了一个 keys 切片，len 为0，cap为200
	for key := range scoreMap {
		keys = append(keys, key)	// 把 map 的 key 加入到 keys 中
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map，其实就是随机生成keys，然后排序，最后按照这个来输出map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

## 元素为 map 类型的切片

下面的代码演示了切片中的元素为map类型时的操作：

```go
func main() {
	var mapSlice = make([]map[string]string, 3)		// make一个切片；切片中的元素为 map
	for index, value := range mapSlice {
		fmt.Printf("index: %d value: %v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

/* 输出
index:0 value:map[]
index:1 value:map[]
index:2 value:map[]
after init
index:0 value:map[address:沙河 name:小王子 password:123456]
index:1 value:map[]
index:2 value:map[]
*/
```

## 值为切片类型的 map

下面的代码演示了 map 中值为切片类型的操作：

```go
func main() {
	var sliceMap = make(map[string][]string, 3)  // map中的值([]string)为 slice
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]  // 判断 map 中键是否存在；没有这个key，ok 为flase
	fmt.Println(value,ok)	// after init
	if !ok {
		fmt.Println("!ok")
		value = make([]string, 0, 2)	// 初始化这个slice
	}
	value = append(value, "北京", "上海","广州")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

/*
map[]
after init
[] false
!ok
map[中国:[北京 上海 广州]]
*/
```



# 练习题

1. 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。









