// 应用程序入口
// 必须是main包 package main
// 必须是main方法：func main()
// 文件名不一定是main.go

package main // 包，表明代码所在的模块

import (
	"fmt" // 引入代码依赖
	"os"
)

// 功能实现
func main() {
	// go不能直接通过main(Args)输入参数
	// 需要借助 os.Args
	// fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("hello world !", os.Args[1])
	}
	fmt.Println("hello world !")
	// go不能通过return来实现返回值
	// 需要借助 os.Exit()
	os.Exit(0) // 0是正常，-1报err
}
