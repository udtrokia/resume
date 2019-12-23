## Ash

#### Q1 - js 单线程异步问题
+ for / setTimeout ~> 同步
+ setTimeout delay 插入任务队列
+ 打印 i 受作用域问题影响，故创建 temp 变量 x


#### Q2 - LeetCode 第 1 题
+ O(n): arr 越长效率越低
+ 确保相加为 target 的两个 element 同时存在于数组即可


#### Q3 - promise 执行顺序及 setTimeout 队列问题
+ @result: 3, 5, 7, 8, 4, 6, 2, 1
+ Promise 定义即执行, 并优先于 setTimeout
+ await 阻塞线程


#### Q4 - 移动端 touch 事件封装调试
+ 通过 chrome 调试 mobile web page.
+ 如若不即时判断, 可用 touchstart/touchend 制作


#### Q5 - 封装 Event 类
+ 携带 data 的 Event 类 CustomEvent
+ addListener / removeListener 需要引用同一 function





## 总结

没有计时做...不过回头来看，计时做是会超出时间上限的。

#### Q1
无障碍。

#### Q2
昨天刚好做了这道题的 rust 版。

#### Q3
Q3 直接运行，复习了下 promise 把概念疏通结束。

#### Q4 
Q4 一开始用 safari 调试，多余地制作了 mousestart/mouseover...等事件的 listener，后想起来 touch 可用 chrome 调试后解决此问题。

#### Q5
Q5 花了不少时间，没直接处理过 Event 类，查了 MDN 文档和 stackoverflow 后解决。
