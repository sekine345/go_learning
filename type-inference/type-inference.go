package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v (%v) is of type %T\n", v, v)

	s := "42"
	fmt.Printf("s (%v) is of type %T\n", s, s)
}

/*
明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、
変数の型は右側の変数から型推論される
*/
