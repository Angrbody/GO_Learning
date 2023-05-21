go语言是不是面向对象的语言？
是，也不是
go语言不支持继承关系，通过复合来实现
go语言接口的实现采用duck-type

封装数据和行为：结构体struct
type Person struct {
    id string
    name string
    age int
}

初始化
p := Person{}
p := new(Person)    // 这里返回的是指针

Go语言的接口与依赖
1. 设计package的时候需要避免循环依赖
2. Duck Type式的接口实现

与其他主要编程语言的差异
1. 接口为非入侵性的，实现不依赖于接口定义
2. 所有接口的定义可以包含在接口使用者的包内

自定义类型
type IntConvertionFn func(n int) 