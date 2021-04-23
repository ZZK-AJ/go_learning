package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

// 传入的是一个 Programmer 的接口实现
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	// 初始化这个 GoProgrammer 结构体
	goProg := &GoProgrammer{}
	//goProg := GoProgrammer{}
	//上面一句会报错，接口类型 对应的应该是一个指针类型的实例
	// 还有 JavaProgrammer
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	// *polymorphism.GoProgrammer fmt.Println("Hello World!")
	// 可以看到输出的是 指针类型
	writeFirstProgram(javaProg)
	// *polymorphism.JavaProgrammer System.out.Println("Hello World!")
}
