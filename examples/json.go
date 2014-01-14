package main

import (
  "encoding/json"
  "fmt"
)

type Message struct {
  Name string `json:"namelol"`
  Body string
  Time int64
}

func main() {
  m := Message{"Alice", "Hello", 1294706395881547000}

  bytes, err := json.Marshal(m)
  if err != nil {
    fmt.Errorf("Error decoding json: %s", err)
  } else {
    fmt.Printf("JSON string: %s\n", bytes)
  }

  var m2 Message
  err = json.Unmarshal([]byte(`{"FOO":"LOL","Name":"Aaron","Body":"Hi!","Time":1294706395881547000}`), &m2)
  if err != nil {
    fmt.Errorf("Error encoding json: %s", err)
  } else {
    fmt.Printf("JSON object: %+v\n", m2)
  }
}
