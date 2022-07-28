### 运行时runtime
- runtime:运行时、运行环境，runtime作为程序的一部分打包进二进制文件
- runtime与用户程序没有明显界限，直接通过函数调用
- ![runtime1](./assets/note-1657880710734.png)
- runtime能力
  - 内存管理能力
  - 垃圾回收能力（GC）
  - 超强的并发能力（Goroutine调度）
  - runtime有一定的屏蔽系统调用能力
  - 一些go的关键字是runtime下的函数
### 编译过程
  ![runtime2](./assets/note-1657880994563.png)
- `go build -n file`：查看编译过程
  ![runtime3](./assets/note-1657887863404.png)
- `$env:GOSSAFUNC="main" go build`：生成SSA中间代码
- `go build -gcflags -S main.go`：生成plan9汇编
- Go程序入口：`runtime/rt0_xxx.s`
  1. 初始化g0执行栈
    - g0是为了调度goroutine而产生的goroutine
    - g0是每个Go程序的第一个goroutine
  2. 运行时检测
      - 检查各种类型长度
      - 检查指针操作
      - 检查结构体字段的偏移量
      - 检查CAS操作
      - 检查atomic原子操作
      - 检查栈大小是否是2的幂次
  3. 参数初始化runtime.args
      - 对命令行中的参数进行处理
      - 参数数量赋值给argc int32
      - 参数值复制给argv **byte
  4. 调度器初始化runtime.schedinit
  ![runtime4](./assets/note-1657951948955.png)
  5. 创建主goroutine
      - 创建一个新的goroutine，执行runtime.main
      - 放入调度器等待调度
  6. 初始化M
     - 初始化一个M，用来调度主协程
  7. 主goroutine执行主函数
    ![runtime5](./assets/note-1657952522220.png)
- Go面向对象：“Yes and no”
  - struct的每个实例并不是“对象”，而是此类型的“值”
  - 组合中的匿名字段，通过语法糖达到类似继承的效果`b是a的字段，a直接调用b的方法其实是先调用b，再调用b的方法`
- gomod的作用：将Go包和Git项目关联起来。Go包的版本就是Git项目的Tag
  - `go env -w xxx=xxx`
  - 想用本地文件替代
    - go.mod文件追加：`replace xxx => xxx/xxx`
    - go vender 缓存到本地：`go build -mod vender`
### 空结构体
  - int大小跟随系统字长、指针的大小也是系统字长
  - 空结构体类型的变量有地址但是没有字长，地址都相同，并且初始都指向`zerobase`
  - 空结构体主要是为了节约内存(结合map实现hashset，配合channel当做纯信号)
### 字符串
- 字符串：本质是一个结构体，data指针指向底层byte数组，len表示utf编码长度之和，表示byte数组的长度（字节数）
 ```go
type stringStruct struct {
	str unsafe.Pointer
	len int
}
// 字符串切分
s = string([]rune(s)[:3])
```
- 字符串大小都是16个字节，切片大小都是32个字节
- 切片的本质是对数组的引用，通过字面量创建切片时是先申请空间，然后依次赋值
- 切片扩容是并发不安全的，并发扩容要加锁
- 字符串和切片都是对底层数组的引用
### map
- map原理![runtime6](./assets/note-1658998904264.png)
- ![runtime7](./assets/note-1658999700019.png)
- 根据哈希值的后B位来选择桶的编号，bmap存储八个键值对以及哈希值的高八位
- 装载因子或者溢出桶的增加都会触发map扩容，扩容可能并不是增加桶数，而是整理，map扩容采用渐进式，桶被操作时才会重新分配
- map的读写有并发问题
- goroutineA协程在桶中读数据时，goroutineB驱逐了这个桶，goroutineA可能读到错的数据或者读不到数据
- sync.Map删除
  - 相比于查询、修改、新增。删除更麻烦
  - 删除后可以分为正常删除、追加后删除
  - 提升后，被删除的key还需特殊处理
- sync.Map读写和追加分离
  - 不会引发扩容的操作（查，改）使用read map，性能好
  - 可能引发扩容的操作（新增）使用dirty map
### 接口
- 接口数据使用runtime.iface表示
- iface记录了数据的地址
- iface中也记录了接口类型信息和实现的方法
- 结构体实现方法，会自动实现一个结构体指针接受的方法
- 以结构体指针实现的方法，不会实现另一个
- 空接口：作为任意类型的函数入参，传入时新生成一个空接口然后再传
```go
type eface struct {
	data
	_type
}
```
### `nil,空接口，空结构体`
- nil是空，并不一定是空指针
  - Type must be a pointer, channel, func, interface, map, or slice type
  - nil不是空结构体的空值
- nil是6种类型的零值
- 每种类型的nil是不同的，无法比较
- 空结构体是Go中非常特殊的类型
  - 空结构体的值不是nil
  - 空结构体的指针也不是nil，但是都相同（zerobase
- 空接口不一定是nil接口
- 类型和值都是nil才是nil接口
- nil是多个类型的零值，或者空值
- 空结构体的指针和值都不是nil
- 空接口零值是nil，一旦有了类型信息就不是nil
### 内存对齐
- 非内存对齐：内存的原子性与效率收到影响
- 内存对齐：提高内存操作效率，有利于内存原子性
- 为了方便内存对齐，Go提供了对齐系数`unsafe.Alignof()`
- 对齐系数的含义是：变量的内存地址必须被对齐系数整除
- 如果系数为4，表示变量内存地址必须是4的倍数
- **结构体对齐**
  - 结构体对齐分为内部对齐和结构体之间对齐
  - 内部对齐：考虑成员大小和成员的对齐系数
  - 结构体长度填充：考虑自身对其系数和系统字长（string对齐系数是8）
  - 内部对齐
    - 指的是结构体内部成员的相对位置（偏移量）
    - 每个成员的偏移量是自身大小与其对齐系数较小值的倍数
  - 结构体长度填充
    - 指的是结构体通过增加长度，对齐系统字长
    - 结构体长度是最大成员长度与系统字长较小的整数倍
  - 结构体对齐系数是其成员最大对齐系数
  - 空结构体出现在结构体末尾时，需要补齐字长，空结构体单独出现时，地址为zerobase