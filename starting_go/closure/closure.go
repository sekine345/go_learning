// closure
// 関数を関数から返す。
// 元の関数内からローカル変数への参照を検出されると、
// 関数終了後に破棄される他のローカル変数とは別のクロージャと結びついた変数として処理される


package main

import (
    "fmt"
)


func later() func(string) string {
    // laterという関数
    // 戻り値はstringを引数とする関数
    // その関数の戻り値はstring

    var store string

    return func(next string) string {
        // 一つ前に渡した文字列を返す関数
        s := store
        store = next
        return s
    }
}


func main() {
    f := later() // 関数定義、fは関数

    fmt.Printf("%v\n", f("A")) // ""
    fmt.Printf("%v\n", f("B")) // "A"
    fmt.Printf("%v\n", f("C")) // "B"
    fmt.Printf("%v\n", f("D")) // "C"
    fmt.Printf("%v\n", f("E")) // "D"
}
