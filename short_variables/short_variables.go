package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 //簡易的な定義、関数内でのみ定義できる
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
