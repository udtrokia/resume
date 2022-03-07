# [clearloop](https://github.com/clearloop)

Unfamous rust developer, emacs user.

For a quick view, this guy can disassemable substrate ( see my substrate [samples](https://clearloop.github.io/clearloop/) ), modify WASM interpreter, handle async rust, drawing cross-platform UI with react and dart ( just for wonderful design ).

For a timeline of my working experience, see my [linkedin](https://www.linkedin.com/in/mercury-fletcher-2277191a3/).

Located in HangZhou, China, GMT + 8, contact me at mercuryfletch@gmail.com.



## Projects

### [elvis](https://github.com/clearloop/elvis)  - Your WASM UI Library

This is a cross-platform UI library written in rust, it uses virtual DOM, compiles to WASM, currently only support `web`, faster than vue for 3~4 times, react around eight times, see [benchmark](https://elvisjs.github.io/book/benchmark.html).

This project joined The Mozilla Open Lab in the summer of 2020.



### [leetcode-cli](https://github.com/clearloop/leetcode-cli) - Command Line Client for Leetcode Written in Rust

Just command tool with intereting commands, has over 3.2k downloads now.



### [sup](https://github.com/clearloop/sup) - Substrate Package Manager

Yet another command line tool of mine, got a grant from web3 foundation.



### [wildboats.xyz](https://wildboats.xyz) - NFT marketplace based on 0xprotocol V4 (_in developing_)

The term of this project is satirizing OpenSea, inspried by sudoswap.xyz but with a marketplace UI, trading NFTs with zero protocol fee for fun, not in production nor open to chinese users due to the policy.



---



## ChainSafe - Full Time, WFH | 2021.4 - now

Now I'm working at ChainSafe as a rust developer.

### [PINT](https://github.com/chainSafe/pint) - A DeFi project built with substrate

- write && maintain the CI and the docker of PINT
- construct the constituent pallet
- handling the parachain stuffs ( including writing e2e tests for rococo )
- built the javascript tool for PINT as well, very familiar with polkadot-js



### [Filecoindot](https://github.com/chainsafe/filecoindot) - Filecoin to substrate bridge

- built the offchain worker of filecoindot
- wrote the polkadot-js api for it



### ChainSafe/thegarii - The Graph Arweave Integration Implementation

- built the relayer client for arweave
- handling the protobuf block format
- PRing to the graph with Arweave integration



## Patract Labs - Full Time, on Site | 2020.11 - 2021.4

Worked at [Patract Labs](https://github.com/patractlabs) as a full stack engineer, write Rust, Typescript, Svelte, enjoyed the job here.

### [Ceres](https://github.com/patractlabs/ceres) - Run ink! contract anywhere

This is the best project I have created in Patract Labs ever, extracted the WASM executor from substrate, mocked the whole chain environment of `pallet-contracts`, can run and debug ink! contract with both wasmi and wasmtime supports without a substrate chain, already supported browser intergration.



### [Megaclite](https://github.com/patractlabs/megaclite) -  zero-knowledge proof tool with ink! contract APIs

To be honest, I'm not good at cryptography, but reading the source code of open source projects is quite easy to me, we had exported the `groth16 proof system` with `curve bn254 and bls12_371` from [zkcrypto](https://github.com/zkcrypto), [matter-labs](https://github.com/matter-labs) and [arkworks](https://github.com/arkworks-rs/utils) and finally use the library of arkworks, because it has more curves and supports compiling to WASM it self.

In this project, my jobs contains exporting libraries, writing testing ink! contract, integrating the solution into the `pallet-contract` of substrate, and do benchmarks.



### WASM Debugger

The main job of here is writing code around WASM, for example, get the WASM debug info of ink! Contract while it's executing in pallet-contract, if interested, you can view the posts I wrote below.

* [Enable WASM Backtrace](https://patractlabs.github.io/blog/enable-wasm-backtrace.html)
* [Support wasmtime in sandbox](https://patractlabs.github.io/blog/support-wasmtime-in-sandbox.html)



---



## Dawrinia - Full Time, WFH | 2020.3 - 2020.11

The last job of mine is in [Darwinia Network](https://github.com/darwinia-network), left because finally I realized I'm working on some projects centralized( I'm not saying that... ), 

I wrote tools for the bridge between Darwinia and Ethereum.



### [Shadow](https://github.com/darwinia-network/shadow) - Darwinia ZK proof server

>  Note: This ZK implemented with merkle mountain tree.

I had written the first version of shadow in golang, after it works on the testnet, I rewrote it in Rust, the code I had written has no big changing during these months but renaming directories, it's stable.



### [Bridger](https://github.com/darwinia-network/bridger) - Darwinia Relayer Client and Watchtower 

Had written the first version of bridger in [typescript](https://github.com/darwinia-network/dj), like the shadow project, I rewrote it in Rust, with async things.





