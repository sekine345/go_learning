// インターフェース
/*
メソッドの一覧を定義したデータ型
複数の構造体が同じメソッドを定義している時、それらの構造体をインターフェース型として扱える
複数の構造体に対する処理を簡単に書ける
*/


package main

import "fmt"


// インターフェース定義
type greeter interface {
    greet()  // 複数の構造体に共通で持たせるメソッド
}

// 構造体
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
    // 同じインターフェース型を満たしているので、スライスに入れることができる
    greeters := []greeter{japanese{}, american{}}

    // for文でメソッドを実行できる
    for _, greeter := range greeters {
        greeter.greet()
    }
}

