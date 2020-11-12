package main

import "fmt"

/*
func add(x int, y int) int {
	// 引数x, y は整数型、戻り地は整数型
	return x + y
}
*/

func add(x, y int) int {
	// 引数x, y の方が同じなら、型宣言は上記のように省略できる
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
