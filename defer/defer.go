package main

import "fmt"

func main() {
	//もとの関数の処理が終わるまで（ここでは main, return が終わるまで）遅延する
	defer fmt.Println("world")

	fmt.Println("hello")
}
