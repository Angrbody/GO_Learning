package emptyinterface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// 如果p可以被断言为int类型
	// if v, ok := p.(int); ok {
	// 	fmt.Println("Integer", v)
	// 	return
	// }

	// if v, ok := p.(string); ok {
	// 	fmt.Println("String", v)
	// 	return
	// }

	// fmt.Println("Unknown Type")

	// 简化上面的写法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("David Tao")
}
