## benchmark
- 用来测试哪种实现方法性能更好
- 第三方库的测试

## benchmark文件创建
- 后缀也是 _test.go
- 函数名为 func BenchmarkXXX

## benchmark测试执行
- go test -bench=.
- go test -bench=. benchmem （可以查看更细节的操作，如空间分配次数）