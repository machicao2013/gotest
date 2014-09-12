package main

import "fmt"

func main() {
    mySlice := make([]int, 2)
    fmt.Println("len: ", len(mySlice), ", cap: ", cap(mySlice))
    for i := 0; i < 8; i++ {
        mySlice = append(mySlice, i)
        fmt.Println("len: ", len(mySlice), ", cap: ", cap(mySlice))
    }
    fmt.Println(mySlice)
}
