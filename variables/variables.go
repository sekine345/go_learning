package main

import "fmt"

var c, python, java bool

//関数外で宣言した変数はどこでも使える
func main() {
	var i int
	//関数内で宣言した変数は関数内でのみ使える
	fmt.Println(i, c, python, java)
	/*
		-> 0 false false false
		空のとき, int -> 0, bool -> fals
	*/

	i = 10
	fmt.Println(i)
	/*-> 10*/
}
