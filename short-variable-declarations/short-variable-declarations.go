package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	/*
		関数の中では短い := の代入文を使い、暗黙的な型宣言ができる
		関数の外では、キーワードではじまる宣言( var, func, など)が必要
	*/

	fmt.Println(i, j, k, c, python, java)
}
