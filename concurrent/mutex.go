package main
import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)
func main() {
    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    var ops int64 = 0

    
}
