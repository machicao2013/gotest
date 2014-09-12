package bubblesort

func BubbleSort(values []int) {
    flag := true
    for i := 1; i < len(values); i++ {
        flag = true
        for j := 0; j < len(values) - i; j++ {
            if values[j] > values[j+1] {
                values[j], values[j+1] = values[j+1], values[j]
                flag = false
            }
        }
        if flag {
            break
        }
    }
}
