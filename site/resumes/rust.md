# Rust工程师 - 张天一

一年经验的 Rust 开发者，[github](https://github.com/clearloop) && [crates](https://crates.io/users/clearloop)，开源贡献: [#334](https://github.com/cryptape/cita/pull/334)，[#396](https://github.com/cryptape/cita/pull/396)，[#1](https://github.com/cryptape/homebrew-cita/pull/1)。



## 工作经历

### 自由职业

> 2019.1 - 至今

#### [SpaceJam](https://github.com/clearloop)

一个受区块链和智能合约启发的分布式微服务框架，特色在于层状可插拔的架构及跨平台:

+ __`primitive`__

SpaceJam 的基本层，纯 std lib 实现，主要包括基本数据结构，IO 模型及 TCP 接口，网络部分实现了 gossip 协议与工作量证明。

+ __`thruster`__

插件层，主要使用 libp2p 实现了 p2p 网络，ed25519 对数据进行加密处理。

+ __`spaceboy`__

高级封装层，API 大杂烩，cli 入口，使用 clap 制作。




#### [cdr.today](https://cdr-today.github.io/intro/)

Ct 作者，前端 `dart` 后端 `go`，学习使用了 flutter 的层状布局结构，运用 bloc(类似 redux)解耦应用请求及组件状态，了解了 Android / iOS 的项目配置及编译发布流程，已上线 App Store。

+ 使用 network link Conditioner 排查极端网络环境下产生的 bug
+ 使用 rxdart 缓存网络请求 stream
+ 搭建 smtp 服务进行邮件收发
+ 修改 zefyr 源码自定义 RichText
+ 嵌入 sqlite 根据用户使用习惯优化用户体验
+ 制作基于 Cupertino 色系，同时适配 iOS / Android dark-mode 自动换色的颜色库 CtColor.
+ 后端使用 redis 与 postgresql
+ 迭代出一个基于 tweetnacl 的免注册登录系统, [链接](https://github.com/lark-in-today/mediumx-prototype)



#### 链上保(已上线)

与 PICC 合作制作的 "区块链 + 保险" 类微信小程序，前端 `taro`, 后台 `vue`, 后端 `node`，其中制作了用于 EVM 的存储类智能合约，在前端结合 leencloud 完成了 DSL 式的查询接口，服务端混合 mongodb 接口及 cita 接口等。

同时期在以太坊测试链上搭建了一个文档索引工具 MedLinker，辅佐以 webRTC 进行点对点传输。



#### [tachion](https://quirky-perlman-c30b6a.netlify.com)

Tachion(区块链) 的展示型 PWA，纯 javascript 制作整个项目(动态渲染无html)，结合浏览器存储进行状态管理，其中做了一些静态资源优化，如上传 cdn，按需加载等。



#### 其他

| 项目名                                                    | 简介                                                                        |
| --------------------------------------------------------- | ------------------------------------------------------------                |
| [Radiancy](https://github.com/udtrokia/Radiancy)          | 使用 Rust 翻译了 [blockchain_go](https://github.com/Jeiwan/blockchain_go)。 |
| __[Sonata](https://crates.io/crates/sonata)__             | __SpaceJam 的合约语言，Lisp 语法 DSL，nostd，仅完成 lexer 部分。__          |
| __[tweetnacl-rs](https://crates.io/crates/tweetnacl-rs)__ | __封装 [sodalite](https://crates.io/crates/sodalite)。__              |
| __[types](https://crates.io/crates/types)__               | __一个用宏判断 Rust 类型的简易库。__                                        |
| __[cjam](https://crates.io/crates/cjam)__                 | __用于吐槽前同事文件依赖混乱 cargo 依赖检查器。__                           |



### 区块链开发实习, BlockShine, 苏州

> 2018.1 - 2019.1

React / Rust 工程师，使用 React 完成了 18 年石墨烯大会官网，使用 React-native 完成了一个区块链钱包，参与区块链开发及文档撰写等。



### 校园创业

> 2017.6 - 2018.1

主要使用 React / node，制作了两个微信小程序与一个垂直电商 App，多个网页。



## 教育经历

#### 安庆师范大学新闻学
> 2015.9 - 2019.6



## 联系方式

| tel               | mail             |
|-------------------|------------------|
| (+86) 18626153029 | udtrokia@163.com |
