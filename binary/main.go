
package main

import "fmt"

func getVal(str string) int {
    val := 0
    for _, c := range str {
        if c == '0' {
            val = val + 1
        } else {
            val = val + 2 
        }
    }
    return val
}

func main() {
    s := "00010110"
    fmt.Println(getVal(s))
}