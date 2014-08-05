
package main
import "sync"
import "fmt"

func main(){
    // go fmt.Printf("Printed in the bacground\n")
    m := make(map[int] string)
    m[2] = "First Value"
    var lock sync.Mutex
    go func() {
        lock.Lock()
        m[2] = "Second Value"
        lock.Unlock()
    }()
    lock.Lock()
    v := m[2]
    lock.Unlock()
    fmt.Printf("%s\n", v)
}
