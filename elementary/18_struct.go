// 構造体 struct
// 複数の値をまとまりとして管理できる


package main

import "fmt"


// struct 定義
type user struct {
    name string  // フィールド
    score int
}


func main() {
    // struct の利用
    u := new(user)
    fmt.Println(u)  // &{ 0}, 空文字と0の構造体のアドレスが返る

    // 値の定義
    u.name = "taguchi"
    u.score = 20
    fmt.Println(u)  // &{taguchi 20}

    u.score = 30
    fmt.Println(u)  // &{taguchi 30}, 値が上書きされる

    // 初めから値を持った構造体を生成するとき
    u2 := user{"taguchi", 20}
    fmt.Println(u2)  // {taguchi 20}

    // 丁寧な定義
    u3 := user{name: "taguchi", score: 30}
    fmt.Println(u3)  // {taguchi 30}
    fmt.Println(u3.name)  // taguchi

    // error
    // u4 := user{"error?", "20"}
    /*
    ./18_struct.go:40:5: u4 declared but not used
    ./18_struct.go:40:26: cannot use "20" (untyped string constant) as int value in struct literal
    */
}

