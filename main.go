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
