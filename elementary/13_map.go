// map
/*
key: value のデータ型
pythonの辞書型みたいなもの
*/


package main

import "fmt"


func main() {
    // mapの定義
    /*
    makeで定義する
    make(map[<keyの型>]<valueの型>)
    */
    m := make(map[string]int)
    fmt.Println(m)  // map[]; 要素を定義していない場合
    m["a"] = 1
    m["b"] = 2
    fmt.Println(m)  // map[a:1 b:2]

    // 定義時に要素を設定する場合
    m2 := map[string]int{
        "a": 1,  // without "," -> syntax error: unexpected newline in composite literal; possibly missing comma or }
        "b": 2,  // without "," -> syntax error: unexpected newline in composite literal; possibly missing comma or }
    }  // 改行して要素の続きを加える時、末尾に","をつけないとエラー
    fmt.Println(m2)

    // delete: 要素の削除
    // da := delete(m, "a")  // deleteは何も要素を返さない -> delete(m, "a") (no value) used as value
    delete(m, "a")  // mから"a"を削除
    delete(m, "c")  // mから"c"を削除、指定したkeyがなければ、何もしない
    fmt.Println(m)  // map[b:2]

    // 存在チェック
    v_b, ok_b := m["b"]  // v_b: bのvalue, ok_b: map m に"b"というkeyがあるかどうかの論理値
    fmt.Println(v_b)  // 2
    fmt.Println(ok_b)  // true
    v_a, ok_a := m["a"]  // 存在しないkeyの場合
    fmt.Println(v_a)  // 0: mの定義時に、valueをintにしたため
    fmt.Println(ok_a)  // false

    bool_map := map[string]bool{"t":true}
    v_t, ok_t := bool_map["t"]
    v_f, ok_f := bool_map["f"]
    fmt.Println(v_t)  // true
    fmt.Println(ok_t)  // true
    fmt.Println(v_f)  // false: map bool_mapの定義時に、valueをboolとしたため、初期値がfalseに
    fmt.Println(ok_f)  // false
    fmt.Println(bool_map["t"])  // true, 存在チェックの表示はない
    t := bool_map["t"]
    fmt.Println(t)  // true, 存在チェックの表示はない
}
