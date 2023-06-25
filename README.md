# ex_go_xorshift
Golang で xorshift の実験
モンテカルロ法で円周率を求めるのにかかる時間を測る

## 使い方

```sh
$ go run main.go [math/rand | xorshift]
```

例:
```sh
$ go run go/main.go mathrand
use math/rand
Calculated pi: 3.1413518
Diff: 0.00024085358979331062
Elapsed: 3.098946161s
```

```sh
$ go run go/main.go xorshift
use xorshift
Calculated pi: 3.1414734
Diff: 0.00011925358979292255
Elapsed: 907.902478ms
```
