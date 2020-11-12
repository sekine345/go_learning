package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split(27))
	fmt.Println(split(13))
}

/*
戻り値となる変数に名前をつける( named return value )ことができます。
戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。

名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。
これを "naked" return と呼びます。
*/
