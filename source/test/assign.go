package main

import "fmt"

type Integer int

func (a Integer) change() {
    a = 20
}

type Change interface {
    change()
}

func main() {
    var a Integer = 30
    var b Change = a
    b.change()
    fmt.Println(a)
    var c Change = &a
    c.change()
    fmt.Println(a)
}
