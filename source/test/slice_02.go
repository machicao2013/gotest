package main

import "fmt"

func main() {
    // 元素个数为5，初始值为0
    mySlice := make([]int, 5)
    length := len(mySlice)
    for i := 0; i < length; i++ {
        fmt.Print(mySlice[i], " ")
    }
    fmt.Println(cap(mySlice))

    // len = 5, cap = 10
    mySlice = make([]int, 5, 10)
    for _, v := range mySlice {
        fmt.Print(v, " ")
    }
    fmt.Println(cap(mySlice))

    // 创建并初始化包含5个元素的切片
    mySlice = []int{1,2,3,4,5}
    for _, v := range mySlice {
        fmt.Print(v, " ")
    }
    fmt.Println(cap(mySlice))
}
