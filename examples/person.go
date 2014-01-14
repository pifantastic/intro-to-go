package main

import "fmt"
import "time"

type Person struct {
  Name     string
  Birthday time.Time
}

func (person *Person) Age() int {
  return time.Now().Year() - person.Birthday.Year()
}

func main() {
  me := Person{
    Name:     "Aaron Forsander",
    Birthday: time.Date(1985, time.December, 13, 0, 0, 0, 0, time.UTC),
  }

  you := Person{}

  fmt.Printf("%s is %d years old\n", me.Name, me.Age())
  fmt.Printf("%s is %d years old\n", you.Name, you.Age())
}
