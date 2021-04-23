package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("Pet Speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("...", host)
}

// 匿名嵌套类型
type Dog struct {
	Pet
}

// 对 Speak 方法重载
func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()		// Wang!
	dog.SpeakTo("Chao")	// Pet Speak... Chao
}
