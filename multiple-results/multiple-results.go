package main

import "fmt"

func swap(x, y string) (string, string) {
	// 関数は複数の戻り値を返すことができます。
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	/*
		a := swap("hello", "world")
		fmt.Println(a)
		// -> assignment mismatch: 1 variable but swap returns 2 values
	*/
}
