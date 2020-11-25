package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

/*
defer ステートメントは、 defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるものです。
→関数が実行完了したら defer した関数が実行される。
*/
