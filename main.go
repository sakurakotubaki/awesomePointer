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
	heapValue := createOnHeap()
	fmt.Println("Value on heap:", *heapValue)

	// new関数を使用してヒープ上に整数を割り当てる
	// Using the new function to allocate an integer on the heap
	anotherHeapValue := new(int)
	*anotherHeapValue = 100
	fmt.Println("Another heap value:", *anotherHeapValue)
}
