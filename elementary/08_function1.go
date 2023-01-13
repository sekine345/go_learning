// 間数

package main

import "fmt"


//func hi(name string) string {
//    //fmt.Println("hi!" + name)
//    msg := "hi!" + name
//    return msg
//}


func hi(name string) (msg string) {
    //fmt.Println("hi!" + name)
    msg = "hi!" + name  // 関数宣言時にmsgを定義している
    /*
    関数の定義の時にreturnする変数をあらかじめ宣言すると、
    returnで変数を描かなくてもいい
    関数が長くなった時に何を返しているかわかりずらそう
    */
    return
}


func main() {
    //hi("takeuchi")
    fmt.Println(hi("takeuchi"))
}
