package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// TestStringFn strings 字符串分割及合并
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

// TestConv 字符串和整形的转换
func TestConv(t *testing.T) {
	// i to a
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// a to i
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
