package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

// 接口类
type Programmer interface {
	WriteHelloWorld() Code
}

// 两种对象类
type GoProgrammer struct {
}

type JavaProgrammer struct {
}

// 两种对象对于方法的实现
func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"hello world\")"
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"hello world\")"
}

// interface传入的对象必须是指针类型
func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

// 测试多态
func TestPoly(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
