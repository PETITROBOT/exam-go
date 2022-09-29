package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    for i, arg := range os.Args[1:] {
        if !regexp.MustCompile(`/d`).MatchString(arg) {
            fmt.Println("Argument " + strconv.Itoa(i+1) + " : " + arg)
        }
    }
}