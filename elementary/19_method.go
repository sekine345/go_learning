// メソッド
// ：構造体などのデータ型に紐づいた間数


package main

import "fmt"


// 構造体を定義
type user struct {
    name string
    score int
}

// 構造体に紐づいた関数を定義
// レシーバーと呼ばれる
func (u user) show() {
    fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

// メソッドへのデータ型の渡し方は、値渡しと参照渡しができる
// 値渡し、元のデータは変更されない
func (u user) hit1() {
    u.score++
}

// 参照渡し、元のデータが変更される
func (u *user) hit2() {
    u.score++
}

func main() {
    u := user{"taguchi", 50}
    u.show()  // name:taguchi, score:50

    u.hit1()
    u.show()  // name:taguchi, score:50, 元データは変更されていない

    u.hit2()
    u.show()  // name:taguchi, score:51, 元データが変更された

}


