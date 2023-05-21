## 并发控制
1. 锁 lock mutex RWLock
2. WaitGroup，类似thread.join()
    只有当wait的所有东西都完成，才能继续向下执行

## CSP
- 使用通道来完成两个实体之间的通信
- 通过channel进行通讯（类似于中间人），使得通信双方更松耦合一些
- golang的channel有容量限制，并且独立于处理Goroutine

## channel
1. classic channel
    - 收发双方需要同时在channel中，否则就阻塞等待
2. buffer channel
    - 不需要收发双方同时在channel中，维护一个buffer