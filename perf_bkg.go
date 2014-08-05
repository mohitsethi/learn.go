package main

import "fmt"
import "sync"
import "os"

func main(){
    var w sync.WaitGroup
    for _, v := range os.Args {
        w.Add(1)
        go func(str string) {
            fmt.Printf("%s\n", str)
            w.Done()
        }(v)
    }
    w.Wait()
}
