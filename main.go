package main

import (
    "fmt"
    "flag"
)

func main() {

    // flags
    start := flag.Int("start", 0, "Starting count")
    inc := flag.Int("inc", 1, "Increment number")
    limit := flag.Int("limit", 10000, "Count max value")
    safe := flag.Bool("safe", false, "Use thread-safe counting?")
    flag.Parse()
    
    // process
    fmt.Printf("Counter Start Value: %d\n", *start)
    fmt.Printf("Counter Increment: %d\n", *inc)
    fmt.Printf("Counter Max Value: %d\n", *limit)
    if !*safe {
        fmt.Println("Using not safe counters.") 
    }

}
