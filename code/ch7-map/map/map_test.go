package my_map

import (
	"fmt"
	"testing"
)

// 三种 map 的初始化
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	// 第二位是 cap
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

	sm := map[string]string{"zzk":"1","aj":"2"}
	t.Log(sm)
	fmt.Println(sm)
}
/*
   map_test.go:8: 4
   map_test.go:9: len m1=3
   map_test.go:13: len m2=1
   map_test.go:16: len m3=0
*/


func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// 可以看到，key 不存在和值为零值返回的都是 0，都是一样的，那么这两个要怎么区分
	// ok 为 true 表示值存在，为 flase 表示值不存在
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

// TestTravelMap 遍历 map
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
