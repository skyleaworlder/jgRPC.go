# jgRPC.go

:cyclone: Toy RPC from 0 to 0.01

## Plan

![target](README.assets/target.png)

A Non-Standard RPC.

## Target

* [x] Self-Defined RPC Protocol; (Non-Standard)
* [x] Protocol Analysis, Serializable/Deserializable; (Simple)
* [x] Procedure Parameters Analysis, using Reflect; (Simple, Need polished)
* [x] Load Balance; (Consist Hash)
* [ ] Cache in Client; (Non-Standard)
* [ ] Service Register Center; (Simple)
* [ ] Service Discovery; (Simple)
* [ ] Health Probe; (Non-Standard)

## Now

![demo](https://skyleaworlder.github.io/2021/02/17/jgRPC/rpc.mp4)

The first window is `Client.go`, the second is `LdBlsServer.go`, while the third and the forth are `NodeServer.go`.

It's clear that `Load Balance Server` does work.

(From the second window, Node1 and Node2 seems out-off-balance for the number of Node Servers is still too small when it comes to using `CONSIST-HASH` as the algothrim of `Load Balance`.)
