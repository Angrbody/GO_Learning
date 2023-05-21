# Go的异常处理 - 编写好的错误处理
1. 函数/断言通过多返回值来判断操作是否发生异常

## 与其他主要编程语言的差异
1. 没有异常机制
2. error类型实现了error接口
3. 可以通过errors.New来快速创建错误实例

## 尽早失败的思想
在程序的前面进行错误判断，即如果发生错误，就解决错误

1. 可以使程序的执行效率变高，避免嵌套结构
2. 可以使程序看起来更简洁清晰

## panic
- panic 用于不可以恢复的错误
- panic退出前会执行defer指定的内容

## OS.Exit
- 退出时不会调用defer指定的函数
- 退出时不会输出当前调用栈信息

## recover()
恢复错误
需要谨慎使用，可能会导致僵尸服务进程
重启 - "let it crash"可能是最好的方法