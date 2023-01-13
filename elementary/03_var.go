// 変数
/*
変数: 1文字目に注意
    1文字目が小文字なら、そのパッケージの中だけで見ることができるもの
    1文字目が大文字の場合、他のパッケージからも見ることができる
*/


package main

import "fmt"


func main(){
    //var msg string
    //msg = "hello world."
    //var msg string = "hello world."
    msg := "hello world."
    fmt.Println(msg)

    //var a, b int
    //a, b = 10, 15
    a, b := 10, 15
    fmt.Println(a, b)

    /*
    var(
        c int
        d string
    )
    c = 20
    d = "string"
    */
    c, d := 20, "string"
    fmt.Println(c, d)
}
