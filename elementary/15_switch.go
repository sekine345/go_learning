// switch


package main

import "fmt"


func check_signal(signal string) {
    // 条件が増えると改修に苦労しそう
    // switch を使ったほうが綺麗になる時はどんな場合か
    switch signal {
        case "red":
            fmt.Println("Stop")
        case "yellow":
            fmt.Println("Caution")
        case "green", "blue":  // 条件は複数書ける
            fmt.Println("Go")
        default:  // defaultは省略可能
            fmt.Println("wrong signal")
    }
}


func check_score(score int) {
    // 前回のifもswitchで書ける
    // switchの後に変数を書かず、caseの後ろに条件文を記載
    switch {
        case score > 80:
            fmt.Println("Great!")
        case score > 60:
            fmt.Println("Good!")
        default:
            fmt.Println("so, so...")
    }

}

func main() {
    check_signal("red")
    check_signal("yellow")
    check_signal("green")
    check_signal("blue")
    check_signal("aaa")

    check_score(85)
    check_score(65)
    check_score(40)
}
