# Developer Advocate - Mercury

Hi, this is Mercury(or Tianyi Zhang in Chinese) to Parity.

Glad to see your advertizement in China, I'm a full stack developer with 3 years experience, familiar with `Rust`, `Dart`, and `Javascript`, used to be an intern at a blockchain company in SuZhou.

For Developer Advocate, I'm a self-taught programmer, have some empathy ability with non technical personnel in some way, besides, writing articles and editing pictures/videos are part of my life, this position might be suitable for me.



## Open Source Contributes

| id             | pull requests                                | status | repository            |
|----------------|----------------------------------------------|--------|-----------------------|
| [#1][#1]       | Add basic scripts, cita-secp256k1, v0.22.0.  | Merged | citahub/homebrew-cita |
| [#298][#298]   | Add help command                             | Merged | umijs/umi             |
| [#334][#334]   | Installation && Exectution Optimization.     | Merged | citahub/cita          |
| [#396][#396]   | Patch to absolute paths in starting scripts. | Merged | citahub/cita          |
| [#1986][#1986] | Add plain_object attribute                   | Open   | rustwasm/wasm-bindgen |
| [#1990][#1990] | Reflect optional struct fields in typescript | Open   | rustwasm/wasm-bindgen |

## Experience

### Freelancer

> January 2019 - Present

#### [elvis][elvis]

Author of `elvis`, this project abstracts UI structures in pure Rust. The core lib of elvis `elvis` depends on Rust std lib completly, and a derived lib named [elvis-web][elvis-web] uses `wasm-bindgen` porting the UI structures to `wasm`.

The aspiration of `elvis` is quite big that I can't keep developing it on my own after these days, it wants to simplify programming behaviors for green hands, and will involve into web from `wasm`, mobile from `flutter`, desktop from `electron`.

**[elvis.js][elvis.js]**

This is the browser version of `elvis`, contains [create-elvis-app][create-elvis-app], [elvis-cli][elvis-cli] and [calling-elvis][calling-elvis] for now.

+ `calling-elvis`: standardize elvis wasm exports for typescript and javascript.
+ `elvis-cli`: cli for elvis apps, actually a webpack plugin, depends on `webpack-dev-server`.
+ `create-elvis-app`: scaffold for creating elvis app.

The syntax of `elvis.js` is inspired by `flutter`, it implements OOP in frontend, but powered by `wasm` and `javascript`. Writing `elvis.js` requires no `html` knowledge nor `css` knowledge, web developing should be easy and efficient as it has never been.

#### [leetcode-cli][leetcode-cli]

Author of `leetcode-cli`, it saids that Rust developers love building wheels, I'm proud that I'm one of the lovers, `leetcode-cli` inspired by the `node` one, but **better**(don`t ask me why, because I wrote it in rust).

The difficult part of this project is handling browser cookies, for example, decrypting cookies from Chrome should across `pkcs5` and `aes-128-cbc`. Beside cryptography, terminal-ui kills me as well, keywords and commands designs took me really a long time, I had got this leetcode-cli idea till the beta version came out takes me one week, coding approach 20 hrs per day.

#### [authing.dart][authing.dart]

Author of `authing.dart`, a dart version of [authing][authing] GraphQL client, completed in 3 days.

####  [cdr.today][cdr.today]

Author of Ct, Ct is a social App which has been online in App Store, developed with `Dart` and `Go`.

+ use network link Conditioner testing bad network
+ use `rxdart` handling http stream
+ set up a smtp server to receive/send mails
+ modify  `zefyer`'s source code to construct Ct's RichText
+ embed `sqlite` to improve UE
+ build a color repo based Cupertino color style，adaptats both iOS and Android's `dark-mode`.
+ use redis and postgresql in backend
+ An auth system implemented with tweetnacl, [link][auth-system].

This project opens source for now, no docs...but a lot of [code][ct-mobile], it takes me almost 6 months thinking and coding at home day and night, the technique in this project might not complex, but the user experiences really bothered me, though it's available in App Store now, I'm not satisfied.

Ct is an experimental App for decentralize systems, it focus on the original type of network communicating, words, and the basic social relationships, small groups, the entry seems not clear for now, but I'll find it out one day.

#### [SpaceJam][spacejam]

Author of SpaceJam, SpaceJam is a micro-service system inspired by Blockchain and Smart Contracts, which is characterized by low coupling layer structure and cross-platform: 

+ __primitive__

  The basic layer in SpaceJam, implemented with the pure std library, which includes transaction data structure, I/O model, TCP APIs and macros. The network part is implemented with gossip protocol and proof-of-work.

+ __thruster__

  This is a  third-part plugin layer, which mainly implemented the transaction pools and peer-to-peer protocol. Technologies in this layer contents: database(sled), crypto(ed25519), and rust-libp2p.

+ __spaceboy__ 

  The advanced packaging layer which mainly contains client(clap) APIs.



#### LianShangBao

A weChat mini-program cooperated with PICC, which is about `blockchain` and `insurance`，using `taro`, `vue` and `node`，wrote EVM contracts and many APIs amoung `cita`, `leancloud` and `mongodb`.

At the same time, my friend and I built a doc-index tool(web app) `MedLinker` on Ethereum test network, using `webRTC` as a feature.



#### Others

| Name                         | Intro                                                 |
| ---------------------------  | ----------------------------------------------------- |
| [cjam][cjam]                 | A checker for cargo dependencies                      |
| [licer][licer]               | Yet another license generator                         |
| [radiancy][radiancy]         | Translated [blockchain_go][blockchain_go] into rust   |
| [sonata][sonata]             | Lisp style DSL                                        |
| [tweetnacl-rs][tweetnacl-rs] | Package [sodalite][sodalite]                          |
| [types][types]               | Rust type conditions using macro                      |



### Blockchain Developer at BlockShine, SuZhou, China

> January 2018 —— January 2019

Full-stack developer intern at BlockShine, developed blockchain wallet App in react-native and blockchain in rust, wrote a few docs.

As a partner of \<Blockchain Thirty-Six Stratagems\>, wrote down 36 articles about "what is Blockchain?" in Chinese for newbies in blockchain industry.



### Startup

> June 2017 —— January 2018

As a full-stack developer, wrote two WeChat mini programs and one React-native App, a few web pages.

I started my company when I was a sophomore, from 2017 till now, I've got involved into startups for 3 times, and started my own company for 2 times, the result is...I failed at any time, but still have plenty of fight left, Bob Dylan sings:

```
Crimson flames tied through my ears, rolling high and mighty traps. Pounced with fire on flaming roads, using ideas as my maps, "We'll meet on edges soon" said I, Proud 'neath heated brow"
```

That's what I'm thiking about now.





## Education 

#### B.A, AnQing Normal University

> September 2015 —— June 2019

I'm a folk who used to want to be a designer or a reporter, things had changed when I abstracted my first design for an campus App, no one can make it come true in my friendships, so I taught myself how to write programs and could not stop it...

+ **sport**: I'm the first guy who join the basketball team of the department of journalismin as a freshman in our department history, my friends and I beat seniors' team when we are sophomores, created history as well, absolutly the champions ; )

+ **music**: I'm a terrible guitar player as well, invovled into several shows in evening parties, no fans, no memories, but the truth is Rock N' Roll is my life, one of my big dream is becoming a Rock Star, driving my steam bicycle across the universe.

+ **editor**: Yet another terrible writer, I completed my first book when I was a sophomore, mainly about my cofusions about 42, and keep running wechat official accounts for 4 years, 300+ articles, 300+ followers, I write about my life, my thoughts about the meanning, about aching to be free as in freedom.

+ **designer**: A terrible designer, I taught myself how to edit pictures and videos when I was a freshman, the whole year, if I'm not in library nor classroom nor basketball court, I'm editing or recording pictures and videos. I'm quite familiar with Adobe family, but not a good designer in some way.

I taught myself most of the time in college, free from it when I was a junior(_broke the law of our school in some way..._), and finally got the diploma by luck. I'm not a good student for school, but might the better dude for human history.


## Contacts

| tel               | mail                  |
|-------------------|-----------------------|
| (+86) 18626153029 | cdr.today@foxmail.com |

[#1]: https://github.com/cryptape/homebrew-cita/pull/1
[#298]: https://github.com/umijs/umi/pull/298
[#334]: https://github.com/cryptape/cita/pull/334
[#396]: https://github.com/cryptape/cita/pull/396
[#1986]: https://github.com/rustwasm/wasm-bindgen/pull/1986
[#1990]: https://github.com/rustwasm/wasm-bindgen/pull/1990
[authing]: https://github.com/Authing/authing
[auth-system]: https://github.com/lark-in-today/mediumx-prototype
[authing.dart]: https://github.com/clearloop/authing.dart
[blockchain_go]: https://github.com/Jeiwan/blockchain_go
[calling-elvis]: https://github.com/clearloop/elvis.js/tree/master/packages/calling-elvis
[cdr.today]: https://cdr-today.github.io/intro/
[cjam]: https://crates.io/crates/cjam
[crates]: https://crates.io/users/clearloop
[create-elvis-app]: https://github.com/clearloop/elvis.js/tree/master/packages/create-elvis-app
[ct-mobile]: https://github.com/clearloop/ct-mobile.old
[elvis]: https://github.com/clearloop/elvis
[elvis.js]: https://github.com/clearloop/elvis.js
[elvis-cli]: https://github.com/clearloop/elvis.js/tree/master/packages/elvis-cli
[elvis-web]: https://github.com/clearloop/elvis/tree/master/web
[github]: https://github.com/clearloop
[leetcode-cli]: https://github.com/clearloop/leetcode-cli
[licer]: https://github.com/clearloop/licer
[radiancy]: https://github.com/udtrokia/Radiancy
[sodalite]: https://crates.io/crates/sodalite
[sonata]: https://crates.io/crates/sonata
[spacejam]: https://crates.io/crates/spacejam
[types]: https://crates.io/crates/typens
[tweetnacl-rs]: https://crates.io/crates/tweetnacl-rs
