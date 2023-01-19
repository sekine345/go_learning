// 配列


package main

import "fmt"


func main() {
    // 配列の宣言
    var a [5]int  // []は配列の要素数、intは型
    fmt.Println(a)  // [0 0 0 0 0]

    // 要素の上書き
    a[2] = 3
    a[4] = 10
    fmt.Println(a)  // [0 0 3 0 10]
    fmt.Println(a[2])  // 3

    // a[0] = "ra ra.."
    // fmt.Println(a[0])
    // ./10_arrangement.go:20:12: cannot use "ra ra.." (untyped string constant) as int value in assignment

    // 要素数が自明の時、宣言時の個数を省略できる
    b := [...]int{1, 3, 5}
    fmt.Println(b)  // [1 3 5]

    // これはできない
    // c := [2, 4, 6]
    // fmt.Println(c)
}
