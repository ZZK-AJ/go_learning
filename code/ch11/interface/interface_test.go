package interface_test

import (
	"testing"
)

// 定义接口，里面是一个个方法
type Programmer interface {
	WriteHelloWorld() string
	ReadHello() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	// 使用 \ 转义
	return "fmt.Println(\"Hello World\")"
}

func (g *GoProgrammer) ReadHello() string  {

	return "fmt.Println(\"ReadHello...\")"
}

func TestClient(t *testing.T) {
	// 定义一个接口变量
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
