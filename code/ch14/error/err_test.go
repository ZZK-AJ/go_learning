package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// 在这个 package 里面定义两个预制的错误
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	// 快速失败，把错误的处理放在前面
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		// 这样通过预设的错误，当需要判断错误的时候，就比较方便了
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}

}

// 最佳实践，及早失败，避免嵌套的结构

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 提倡这种形式，err != nil 先判断错误及处理
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {

		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}
