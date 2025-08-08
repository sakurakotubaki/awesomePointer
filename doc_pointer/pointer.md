# ポインターとは？
変数に値を格納するときに、コンピュータはメモリの特定の位置（アドレス）
を指定して値を保持する。

ポインタでメモリ上のアドレスを参照する

```go
package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)
}
```

**実行する**

変数xの値を参照すると１０と表示される。

変数xに&をつけるとポインタ（メモリのアドレス）が表示される。

```shell
10
0x14000102020
```

メモリの中身を確認するときは、ポインター変数を定義する。

```go
package main

import "fmt"

func main() {
	// 通常の変数
	x := 10
	// 変数xのメモリのアドレスを表示
	address := &x
	// ポインタ変数
	var pointer *int = &x
	fmt.Println(x)
	fmt.Println(address)
	// メモリの中身を取得するときには、ポインタ変数うに" * "をつける。
	fmt.Println(*pointer)
}
```