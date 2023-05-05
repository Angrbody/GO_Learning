# 协程 Goroutine
## 线程和协程
1. 创建时默认的stack大小
- JDK5以后的java thread stack默认为1M
- Goroutinue的stack初始化大小为2K

2. 和KSE(Kernel Space Entity)的对应关系
- Java Thread为 1：1
- Goroutine是 M：N，多对多

提高并发能力
M:System Thread
P:Processor
G:Goroutine


