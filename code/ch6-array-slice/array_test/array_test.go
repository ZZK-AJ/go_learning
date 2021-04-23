package array_test

import "testing"

// 连续存储空间，数组和切片
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[2] = 666

	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

// TestArrayTravel 数组的两种遍历方式
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	t.Log("===")
	// 注意这里
	for _, e := range arr3 {
		t.Log(e)
	}
}

// TestArraySection 数组元素的切片
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[1:]
	arr3_sec1 := arr3[1:len(arr3_sec)]
	t.Log(arr3_sec)
	// [2 3 4 5]
	t.Log(arr3_sec1)
	// [2 3 4]
}
