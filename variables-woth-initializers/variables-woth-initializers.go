package main

import "fmt"

var i, j int = 1, 2

// i, j を、整数型として宣言、1, 2 と初期化する

func main() {
	var c, python, java = true, false, "no!"
	// 初期化子が与えられている場合、宣言時に型を明示しなくてもいい。
	// 型が複数混在している状態でも宣言できる
	fmt.Println(i, j, c, python, java)
}
