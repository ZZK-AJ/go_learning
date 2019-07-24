package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 初始化一个 slice 设置 len=3 cap=5
	// len 表示可访问的元素的个数 cap 表示容量
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// 这增加一个
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	// 0 0 0 1
	t.Log(len(s2), cap(s2))
	//	4 5
}

// 切片是如何实现可变长的
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
/* 可以看到 cap 的两倍增长的
   slice_test.go:26: 1 1
   slice_test.go:26: 2 2
   slice_test.go:26: 3 4
   slice_test.go:26: 4 4
   slice_test.go:26: 5 8
   slice_test.go:26: 6 8
   slice_test.go:26: 7 8
   slice_test.go:26: 8 8
   slice_test.go:26: 9 16
   slice_test.go:26: 10 16
*/

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}
/* 这里的 cap 就是指连续的存储空间到最后的长度
   slice_test.go:34: [Apr May Jun] 3 9
   slice_test.go:36: [Jun Jul Aug] 3 7
   slice_test.go:38: [Apr May Unknow]
   slice_test.go:39: [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
*/

// TestSliceComparing 切片只能和 nil 比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { // 切片只能和 nil 比较
	// 	t.Log("equal")
	// }
	t.Log(a, b)
}

/*
// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
    result := []int{}
    tempMap := map[int]byte{}  // 存放不重复主键
    for _, e := range slc{
        l := len(tempMap)
        tempMap[e] = 0
        if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
            result = append(result, e)
        }
    }
    return result
}
*/
// TestSliceRepetitionByMap 通过字典去除 slice 中重复的元素
func TestSliceRepetitionByMap(t *testing.T) {
	slc := []int{1,2,3,4,5,2,3,4,5}
	result := []int{}
	tempMap := map[int]byte{}  // 存放不重复主键
	for _, e := range slc{
		l := len(tempMap)
		fmt.Println(l)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
			fmt.Println(result)
		}
	}
	fmt.Println(result)
}

// 空struct的作用,slice 去重
// 总共初始化两个变量，一个长度为0的slice，一个空map。由于slice传参是按引用传递，没有创建额外的变量
// 利用了map的多返回值特性
// 空struct不占内存空间，可谓巧妙
func TestRemoveDuplicateElement(t *testing.T)  {
	s := []string{"hello", "world", "hello", "golang", "golang", "hello", "ruby", "python", "java","python"}
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	t.Log(result)
}

// 判断 slice 是否为空，cap 是 0,2,4,8 增长的
func TestEmptySlice(t *testing.T)  {
	var s []string
	//s := make([]string,3)
	t.Log(len(s),cap(s))
	s = append(s,"zzk")
	s = append(s,"zzk")
	t.Log(len(s),cap(s))
	s = append(s,"zzk")
	s = append(s,"zzk")
	s = append(s,"zzk")
	t.Log(len(s),cap(s))
	//slice[1] = ""
	if cap(s) == 0 {
		t.Log("s is empty!")
	} else {
		t.Log("s is not empty!")
	}
}