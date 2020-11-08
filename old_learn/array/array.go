package main

import "fmt"

func main() {
	var a [2]string //a を文字列２つの配列と定義
	fmt.Println(a)
	a[0] = "Hello" // a の0番目の要素を定義
	fmt.Println(a) // からの配列が表示される
	fmt.Println(a[0])
	a[1] = "World"
	fmt.Println(a[0], a[1]) // 要素を1つずつ表示
	fmt.Println(a)          // 配列として表示する

	primes := [6]int{2, 3, 5, 7, 11, 13} // 6つの整数型で配列を定義、要素は{}で定義
	fmt.Println(primes)

}
