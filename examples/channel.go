package main

import "fmt"
import "time"

func main() {
  ch := make(chan int)

  go func() {
    time.Sleep(time.Second)
    fmt.Println("Done sleeping.")
    ch <- 1
  }()

  fmt.Println("Called goroutine.")

  // Blocks until value received.
  i := <-ch

  fmt.Printf("%d received\n", i)
}
