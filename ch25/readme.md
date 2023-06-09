# sync.Pool
## sync.Pool是一种缓存池，最多只能缓存一个对象

## 对象获取结构
Processor
- 1. 私有对象（协程安全）
    如果私有对象中有所需的缓存，则直接从私有对象中获取
- 2. 共享池（协程不安全，需要加锁）
    私有对象不存在，尝试从当前Processor的共享池获取
- 3. 如果当前的Processor共享池也是空的，那么尝试去其他Processor的共享池获取
- 4. 如果所有子池都空，则由用户指定的New函数产生一个新的对象返回

## sync.Pool对象的生命周期（不能当作对象池用）
- GC会清除sync.Pool缓存的对象
- 对象的缓存有效期为下一次GC之前

## sync.Pool总结
- 适用于通过复用，降低复杂对象的创建和GC代价
- 协程安全，会有锁的开销（比较锁/创建对象的开销）
- 生命周期受GC影响，不适用于做连接池等，需要自己管理生命周期的资源的池化