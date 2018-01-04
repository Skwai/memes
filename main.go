package main

import (
	"fmt"
)

func main() {
	a := App{}
	a.Initialize()
  a.Run(":3001")

  m := Meme{}
  err := m.getMeme("feqkVgjJpYtjy")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(m.Id, m.Url)
  }
}
