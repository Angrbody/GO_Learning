// go赋值可以进行自动类型推断
// 在一个赋值语句中可以对多个变量进行赋值

package fib
import "testing"
func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	a := 1
	b := 1

	t.Log(a)
	for i := 0 i < 5; i++ {
		t.Log(" ", b)
		tmp := a 
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a:=1
	b:=1
	// 1.使用中间变量来交换
		// tmp:=a 
		// a=b
		// b=tmp
		
	// 2. 直接交换
	a,b = b,a
	t.Log(a,b)
}