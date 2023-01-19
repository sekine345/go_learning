// app.go 用のtestファイル

package main

import (
    "testing"
)


func TestAppName(t *testing.T) {
    expect := "Zoo Application"
    result := AppName()

    if result != expect {
        t.Errorf("%s != %s", result, expect)
    }
}
