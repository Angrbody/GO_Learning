package interface_test

import "testing"

// 定义接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义实现对象类, duck type
type GoProgrammer struct {
}

// 保持方法的签名一样
// func后面跟着的括号里标识了该方法的实现者
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	// 定义一个接口变量，类型是programmer
	// var p Programmer
	// p = new(GoProgrammer)
	// t.Log(p.WriteHelloWorld())

	// 直接定义一个GoProgrammer的指针
	p1 := new(GoProgrammer)
	t.Log(p1.WriteHelloWorld())
}
