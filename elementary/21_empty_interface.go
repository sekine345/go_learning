// 空のインターフェース
/*
関数に空のインターフェースを渡すことにより、
渡した構造体の種類を判別、処理を場合分けできる
interface{}

いまいちどういう時に使うかわからん
*/


package main

import "fmt"


// 空のインターフェースを使った間数
func show(t interface{}) {
    // 処理の中の実装においてはそのデータ型が何かを知る必要がある
    // 判別の方法は２種類
    // 1. 型アサーション
    /*
    _, ok := t.(japanese)  // 渡したデータ型がjapaneseかどうか判別
    if ok {
        fmt.Println("i am japanese.")
    } else {
        fmt.Println("i am not japanese.")
    }
    */
    // 2. 型switch
    switch t.(type) {
        case japanese:
            fmt.Println("i am japanese")
        default:
            fmt.Println("i am not japanese")
    }
}

// インターフェース定義
type greeter interface {
    greet()
}

// 構造体定義
type japanese struct {}
type american struct {}

// 共通メソッド
func (j japanese) greet() {
    fmt.Println("Konnitiha")
}
func (a american) greet() {
    fmt.Println("Hello")
}


// 利用
func main() {
    greeters := []greeter{japanese{}, american{}}

    for _, greeter := range greeters {
        greeter.greet()
        show(greeter)
    }
}

