Go Concurrency Patterns
===

This repository contains a curated list of Go concurrency patterns,
summarized from Google I/O talks:

- [Google I/O 2012 - Go Concurrency Patterns (Rob Pike)](https://youtu.be/f6kdp27TYZs?si=vQGLf8rfv8H3zakk)
- [Google I/O 2013 - Advanced Go Concurrency Patterns (Sameer Ajmani)](https://youtu.be/QDDwwePbDtw?si=8Sxjl-KXrWA9pNQJ)

## Table of Contents

### Basic

|                                      Name                                      |                     Description                     |               Go Playground               |
|:------------------------------------------------------------------------------:|:---------------------------------------------------:|:-----------------------------------------:|
|                       [0-start](1-basic/0-start/main.go)                       |          The initial example of goroutine           | [Play](https://go.dev/play/p/xqFZZwFrbSG) |
|                      [1-boring](1-basic/1-boring/main.go)                      |             A hello world to goroutine              | [Play](https://go.dev/play/p/eFVs9yXMnTm) |
|                     [2-channel](1-basic/2-channel/main.go)                     |             A hello world to go channel             | [Play](https://go.dev/play/p/-zZc8FJQ1jT) |
|                   [3-generator](1-basic/3-generator/main.go)                   |          Generator pattern implementation           | [Play](https://go.dev/play/p/hfb7q4i98zu) |
|                       [4-fanin](1-basic/4-fanin/main.go)                       |            Fan-in pattern implementation            | [Play](https://go.dev/play/p/umshVDeYb-c) |
|            [5-restore-sequence](1-basic/5-restore-sequence/main.go)            |                  Restore sequence                   | [Play](https://go.dev/play/p/2kxDiKNls9R) |
|                      [6-select](1-basic/6-select/main.go)                      |               Select statement in Go                | [Play](https://go.dev/play/p/YGsGHZ1gjBI) |
|        [7-fanin-select-timeout](1-basic/7-fanin-select-timeout/main.go)        |         Fan-in pattern with timeout support         | [Play](https://go.dev/play/p/dq21kOSicvY) |
|            [8-select-timeout-1](1-basic/8-select-timeout-1/main.go)            |            Select statement with timeout            | [Play](https://go.dev/play/p/sCj8UBqBVQm) |
|            [9-select-timeout-2](1-basic/9-select-timeout-2/main.go)            |       Select statement with timeout - global        | [Play](https://go.dev/play/p/uus-QqTReuj) |
|               [10-quit-channel](1-basic/10-quit-channel/main.go)               |              Channel with quit signal               | [Play](https://go.dev/play/p/FJGPDjF8fEk) |
|       [11-quit-channel-receive](1-basic/11-quit-channel-receive/main.go)       |          Receive message from quit channel          | [Play](https://go.dev/play/p/ibLDze5bGa1) |
|              [12-daisy-chain-1](1-basic/12-daisy-chain-1/main.go)              |           Daisy chain concurrency pattern           | [Play](https://go.dev/play/p/Pm5sVOKv_hK) |
|              [13-daisy-chain-2](1-basic/13-daisy-chain-2/main.go)              |    An equivalent daisy chain pattern Alternative    | [Play](https://go.dev/play/p/cUQWZ0lawTQ) |
|          [14-google-search-1.0](1-basic/14-google-search-1.0/main.go)          | Build a concurrent google search from the ground up | [Play](https://go.dev/play/p/JKv1xveiSdZ) |
|          [15-google-search-2.0](1-basic/15-google-search-2.0/main.go)          | Build a concurrent google search from the ground up | [Play](https://go.dev/play/p/RXc39fI3ViR) |
|          [16-google-search-2.1](1-basic/16-google-search-2.1/main.go)          | Build a concurrent google search from the ground up | [Play](https://go.dev/play/p/wiOlDBX6NCO) |
|          [17-google-search-3.0](1-basic/17-google-search-3.0/main.go)          | Build a concurrent google search from the ground up | [Play](https://go.dev/play/p/DDgi4H71aO4) |
|                  [18-sieve](1-basic/18-others-sieve/main.go)                   |                   Go prime sieve                    | [Play](https://go.dev/play/p/M2n1LCd2Bef) |
|                                19-loadbalancer                                 |               Go load balancer (TBC)                |                     -                     |
| [_**20-chatroulette**_](https://github.com/naichuan-zhang/go-grows-with-grace) |                Go chat roulette toy                 |                     -                     |
|             [21-powerseries](https://go.dev/test/chan/powser1.go)              |        Concurrent power series (by McIlroy)         |                     -                     |

### Advanced

|                     Name                      |                   Description                    |               Go Playground               |
|:---------------------------------------------:|:------------------------------------------------:|:-----------------------------------------:|
| [1-ping-pong](2-advanced/1-ping-pong/main.go) | A sample ping-pong two players game in goroutine | [Play](https://go.dev/play/p/3vOEYlUPSTW) |

## Takeaway Points

### Don't overdo it

- They are fun to play with, but don't overuse these ideas.
- Goroutines and channels are big ideas. They're tools for program construction.
- But sometimes all you need is a reference counter.
- Go has "sync" and "sync/atomic" packages that provides mutexes, condition variables, etc. They provide tools for
  smaller problems.
- Often, these things will work together to solve a bigger problem.
- Always use the right tool for the job.

### Conclusions

- Goroutines and channels make it easy to express complex operations dealing with
    - multiple inputs
    - multiple outputs
    - timeouts
    - failure
- And they're fun to use.

## Useful Links

**Go Home Page:** <br/>
[go.dev](https://go.dev/)

**Go Tour (learn Go in your browser):** <br/>
[go.dev/tour](https://go.dev/tour/)

**Package documentation:** <br/>
[pkg.go.dev/std](https://pkg.go.dev/std)

**Articles galore:** <br/>
[go.dev/doc](https://go.dev/doc/)

**Concurrency is not parallelism:** <br/>
[go.dev/blog/waza-talk](https://go.dev/blog/waza-talk)

**Go: code that grows with grace** (with chat roulette implementation): <br/>
[naichuan-zhang/go-grows-with-grace](https://github.com/naichuan-zhang/go-grows-with-grace)
