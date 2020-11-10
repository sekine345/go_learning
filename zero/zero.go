package main

import "fmt"

func main() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)
	// -> 0, 0, false, ""
	fmt.Printf("%v, %v, %v, %v\n", i, f, b, s)
	// -> , 0, false,
}

/*
- %q
    対応する文字列をクォートでくくる
*/

/*
ゼロ値
- 数値型
    0
- bool型
    bool
- 文字列型
    "": 空文字列
*/
