package main

import "fmt"

// この関数はヒープ上に整数を割り当て、そのポインターを返します
// This function allocates an integer on the heap and returns a pointer to it
func createOnHeap() *int {
	value := 42
	return &value // ローカル変数のアドレスを返すと、その変数はヒープにエスケープします
	// Returning the address of a local variable causes it to escape to the heap
}

func main() {
	// 基本的なポインター操作
	// Basic pointer operations
	x := 10
	p := &x // p is a pointer to x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value at address p (*p):", *p)
	*p = 20 // Change value of x through pointer
	fmt.Println("New value of x:", x)

	fmt.Println() // 空行を追加

	// ヒープメモリの例
	// Heap memory examples
	heapValue := createOnHeap()
	fmt.Println("Value on heap:", *heapValue)

	// new関数を使用してヒープ上に整数を割り当てる
	// Using the new function to allocate an integer on the heap
	anotherHeapValue := new(int)
	*anotherHeapValue = 100
	fmt.Println("Another heap value:", *anotherHeapValue)
}
