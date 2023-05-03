package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 1. 对dog使用复合而非继承的方式
type Dog struct {
	p *Pet
}

// 重载父类的方法
func (d *Dog) Speak() {
	fmt.Print("wang!")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	myDog := new(Dog)
	myDog.Speak()
	myDog.SpeakTo("liu")

	// go不支持LSP类型转换
	// 不能使用基类指针指向子类对象，强制转换也不可以
	// var dog Pet = Pet(new(Dog))
}

// 2. 对cat通过内嵌对象采用匿名“继承”的方式
type Cat struct {
	Pet
}

// 即使没有函数重载，也可以使用父类方法
func TestCat(t *testing.T) {
	myCat := new(Cat)
	myCat.SpeakTo("jyliu")
}
