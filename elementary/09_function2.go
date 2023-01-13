// 関数
// 複数の返却、関数リテラル、即時関数


package main

import "fmt"


func swap(a, b int) (int, int) {
    return b, a
}


func main() {
    //fmt.Println(swap(1, 2))

    // 関数リテラル
    // 変数に関数を渡すことで、関数名を省略できる
    // 関数の中でしか使わない処理をまとめられる
    f := func(a, b int) (int, int) {
        return b, a
    }
    fmt.Println(f(1, 2))

    // 即時関数
    // 宣言と同時に実行
    // 使い方のイメージがつかない
    // 宣言と同時に実行するので、関数名もいらない
    func(name string) {
        fmt.Println("hi! " + name)
    }("takeuchi")
}
