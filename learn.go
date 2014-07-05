package main

import (
    "fmt"           // Go standard library
//    "net/http"      // A web server
//    "strconv"       // String conversions
)

func main(){
    fmt.Println(multipleReturns(5, 2))
}

func multipleReturns(x, y int) (sum, prod int) {
    return x + y, x * y
}
