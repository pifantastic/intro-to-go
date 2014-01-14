package main

import "fmt"

func main() {
  ch := make(chan int, 2)

  ch <- 1
  ch <- 2
  ch <- 3

  fmt.Println("Wrote 2 integers to the channel")
}
