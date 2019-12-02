# find-a-job

### javascript

+ 静态资源优化，webRTC，布局，
+ http 状态码


+ node / golang 并发区别

node 使用的是单线程异步处理模式，golang支持多线程处理

golang - 协程

又称微线程，在单线程上执行多个任务，用函数切换，开销极小。不通过操作系统调度，没有进程、线程的切换开销。genvent，monkey.patchall

多线程请求返回是无序的，那个线程有数据返回就处理那个线程，而协程返回的数据是有序的。

+ redux
状态机，对象存储，React Component -> Action -> Store -> (Component) -> Reducers.

+ 闭包
读取内部变量，回调，封装变量。

+ 原型链
所有函数都有 prototype，所有对象都有 \_\_proto\_\_，contrustor 指向 prototype，\_\_proto\_\_ 可以一直指到 null。

+ 并发

并发是轮流处理多个任务的能力 / 任务队列，并行是同时处理多个任务。this

js 是单线程的(一个主线程和任务队列，为了避免复杂性)，html5 web worker 允许 js 创建多个线程，但子线程受主线程控制，并不能操作 DOM，异步 -> 任务队列(event loop -> 回调函数)，

+ 进程

cpu 调度分派的基本单位，有单独的地址空间。一个进程至少有一个线程，线程没有独立地址空间，并发性高。



### react
+ redux 使用
+ 生命周期及原理
+ MVVC
+ MVC


### Rust
+ 线程锁
+ Rust Ownership
+ mut


### 数据库

+ SQL 多种查询办法


### 运维类

+ Docker 的使用
+ shell script

+ 指针地址
地址是字节编号，一串 16进制数字；指针是保存地址的变量。


### 网络层
+ tcp / udp 原理
1） TCP提供面向连接的传输，通信前要先建立连接（三次握手机制）； UDP提供无连接的传输，通信前不需要建立连接。
2） TCP提供可靠的传输（有序，无差错，不丢失，不重复）； UDP提供不可靠的传输。
3） TCP面向字节流的传输，因此它能将信息分割成组，并在接收端将其重组； UDP是面向数据报的传输，没有分组开销。
4） TCP提供拥塞控制和流量控制机制； UDP不提供拥塞控制和流量控制机制。
