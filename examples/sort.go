package main

import (
  "fmt"
  "sort"
)

type Sequence []int

func (s Sequence) Len() int {
  return len(s)
}

func (s Sequence) Less(i, j int) bool {
  return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func main() {
  s := Sequence{8, 6, 7, 5, 3, 0, 9}

  fmt.Printf("Unsorted: %+v\n", s)

  sort.Sort(s)

  fmt.Printf("Sorted: %+v\n", s)
}
