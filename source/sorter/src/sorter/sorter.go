package main

import "io"
import "os"
import "flag"
import "fmt"
import "bufio"
import "strconv"
import "time"

import "algorithm/bubblesort"
import "algorithm/qsort"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string)(values []int, err error) {
    file, err := os.Open(infile)// {{{
    if err != nil {
        fmt.Println("Failed to open file ", infile)
        return
    }
    defer file.Close()

    br := bufio.NewReader(file)
    values = make([]int, 0)
    for {
        line, prefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }
        if prefix {
            fmt.Println("A too long line, seems unexpected.")
            return
        }
        str := string(line)
        value, err1 := strconv.Atoi(str)
        if err1 != nil {
            err = err1
            fmt.Println("Invalid line ", str)
            return
        }
        values = append(values, value)
    }
    return// }}}
}

func writeValues(outfile string, values []int) error {
    file, err := os.Create(outfile)// {{{
    if err != nil {
        fmt.Println("Failed to create file ", outfile)
        return err
    }
    defer file.Close()

    bw := bufio.NewWriter(file)
    for _, value := range values {
        bw.WriteString(strconv.Itoa(value) + "\n")
    }
    bw.Flush()
    return nil// }}}
}

func main() {
    flag.Parse()

    if infile != nil {
        fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
    }

    values, err := readValues(*infile)

    if err == nil {
        t1 := time.Now()
        switch *algorithm {
            case "qsort":
                qsort.QuickSort(values)
                writeValues(*outfile, values)
            case "bubblesort":
                bubblesort.BubbleSort(values)
                writeValues(*outfile, values)
            default:
                fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
        }
        t2 := time.Now()
        fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
    } else {
        fmt.Println(err)
    }
}
