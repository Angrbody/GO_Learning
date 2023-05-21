# package
1. 基本服用模块单元
- 以首字母大写来表明可以被包外代码访问

2. 代码的package可以和所在的目录不一致

3. 同一目录里的Go代码的package要保持一致

## import报错
在gomodule=on模式下总是报错，找不到自定义的包
设置off后删除了所有的go.mod文件，debug成功
go mod后面再学

## 自定义包的注意事项
1. 函数首字母必须大写，否则不能被外部访问
2. 路径必须在go_path下（go mod开启就不需要这个）

## init方法
1. 在main被执行前，所有依赖的package的init方法都会被执行
2. 不同包的init函数按照包导入的依赖关系决定执行顺序（go编译器自动处理）
3. 每个包可以有多个init函数
4. 包的每个源文件也可以有多个init函数

## go get
1. 通过go get来获取远程依赖
- go get -u 强制从网络更新远程以来
- https://github.com/easierway/concurrent_map.git 需要去掉前面的https:// 和后面的 .git

2. 注意代码在GitHub上的组织形式，以适应go get
- 直接以代码路径开始，不要有src

## Go未解决的问题
1. 同一环境下，不同项目使用同一包的不同版本
2. 无法管理对包的特定版本的依赖
- 提出了vendor路径，细化了包的路径，很多第三方的依赖管理都是依托与vendor
- godep
- glide
    1. glide init
    2. glide install ：生成vendor文件
- dep