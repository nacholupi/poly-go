package main

import (
	"fmt"
	"polygon/polygon"
	"strings"
)

func main() {
	fmt.Println("This is a test")
	r := polygon.NewCSVReqReader(strings.NewReader("ID1, 1.1, 2.2, 1, 4\n ID2, 1.1, 2.2, 111, 4"))
	w := &polygon.ConsoleWriter{}
	err := polygon.FromRadiusIO(r, w)
	if err != nil {
		fmt.Println(err)
	}
}
