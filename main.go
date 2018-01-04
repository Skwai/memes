package main

import (
	"fmt"
)

func main() {
	a := App{}
	a.Initialize()
	a.Run(":3001")

	m, err := getRandomMemes(4)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}
}