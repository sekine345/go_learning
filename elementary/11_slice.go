// 配列: スライス
/*
配列は要素数固定で動的に変化できない
また、関数に配列を渡すときは丸ごと渡すことになるので、
メモリ的に非効率
配列をもっと便利に使えるデータ型：スライス
*/

package main


import "fmt"


func main() {
    a := [5]int{0, 1, 2, 3, 4}
    s := a[2:4]  // [2, 3]
    fmt.Println(a)  // [0 1 2 3 4]
    fmt.Println(s)  // [2 3]

    // スライスの参照先も書き換えるので、aも変わる
    s[1] = 30
    fmt.Println(a)  // [0 1 2 30 4]
    fmt.Println(s)  // [2 30]

    // 便利な命令
    fmt.Println(len(s))  // 2  // 配列の要素数
    fmt.Println(cap(s))  // 3  // 配列の最大容, 4 まで量
}
