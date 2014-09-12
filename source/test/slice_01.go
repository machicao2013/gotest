package main

import "fmt"

func main() {
    var myArray = [10]int{1,2,3,4,5,6,7,8,9,10}

    var mySlice = myArray[:5]

    mySlice[0] = 20

    for _, v := range myArray {
        fmt.Print(v, " ")
    }
    fmt.Println()

    for _, v := range mySlice {
        fmt.Print(v, " ")
    }
    fmt.Println()
}
