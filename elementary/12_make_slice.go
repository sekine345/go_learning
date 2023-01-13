// スライス
// 配列を介さないでスライスを作る


package main

import "fmt"


func main() {
    // make: スライスを作る
    // s := make([]int, 3)  // make(型, len, cap); capは省略可能
    s := []int{1, 3, 5}  // 配列と同じように定義可能, []の中がからであることに注意
    fmt.Println(s)

    // append: スライスの末尾に要素を追加
    /*
    sの末尾に[2, 4, 6] を追加
    新しいスライスを作る
    sに渡せば、sが上書きされる
    ssに渡すと、sとは別の配列ができる
    ssの要素を書き換えても、sの要素は変わらない
    */
    ss := append(s, 2, 4, 6)
    ss[0] = 10
    fmt.Println(s)  // [1 3 5]
    fmt.Println(ss)  // [10 3 5 2 4 6]

    // copy: スライスのコピー
    /*
    要素数を指定したスライスを最初に定義し、後から要素をコピーする
    pythonのような書き方ではエラー: t := copy(s)
    copyは要素数を返す
    copyで生成した新しいスライスはコピー元とは別
    コピーしたスライスの要素を書き換えても、コピー元の要素は変わらない
    */
    t := make([]int, len(ss)) // ssと同じ要素数のスライス
    n := copy(t, ss)
    t[0] = 100
    fmt.Println(ss)  // [10 3 5 2 4 6]
    fmt.Println(t)  // [100 3 5 2 4 6]
    fmt.Println(n)

    // コピー先の要素数がコピー元より多い時、
    // コピー元の先頭に要素がコピーされる
    u := make([]int, 10)
    m := copy(u, s)
    fmt.Println(s) // [1 3 5]
    fmt.Println(u)  // [1 3 5 0 0 0 0 0 0 0]
    fmt.Println(m)  // 3; コピーした要素数なので
    fmt.Println(len(u))  // 10; uの要素数は10

    // コピー先の要素数がコピー元より少ない時、
    // コピー先の要素数だけコピーされ、入らない分は無視
    v := make([]int, 1)
    o := copy(v, s)
    fmt.Println(s)  // [1 3 5]
    fmt.Println(v)  // [1]
    fmt.Println(o)  // 1
}
