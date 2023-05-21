package encapsulationtest

import (
	"fmt"
	"testing"
	"unsafe"
)

type Person struct {
	id   string
	name string
	age  int
}

func TestCreatePerson(t *testing.T) {
	p1 := Person{"0", "liujinyang", 24}
	p2 := Person{id: "1", name: "linan"}
	p3 := new(Person) // p3是指针类型
	p2.age = 23
	p3.name = "xuchenlong"

	t.Log(p1)
	t.Log(p2)
	t.Log(p3)
	t.Logf("p1 is %T", &p1) // 加上&也会变成指针
	t.Logf("p3 is %T", p3)
}

// 值传递
func (p Person) String1() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&p.name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", p.id, p.name, p.age)
}

// 指针传递（没有对象产生，减少内存拷贝，推荐使用）
func (p *Person) String2() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&p.name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", p.id, p.name, p.age)
}

func TestStructOperations(t *testing.T) {
	p1 := Person{"0", "Bob", 20}
	p2 := &Person{"1", "Tom", 21} // 指针
	fmt.Printf("address is %x\n", unsafe.Pointer(&p1.name))
	fmt.Printf("address is %x\n", unsafe.Pointer(&p2.name))
	t.Log(p1.String1())
	t.Log(p2.String2())
}
