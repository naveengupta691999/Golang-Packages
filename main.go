package main

import (
	"fmt"
	"sample/addition"
	"sample/modulo"
)

func main() {
	a := 12
	b := 15
	fmt.Println(addition.GetSum(a, b))

	fmt.Println(modulo.GetMod(a, b))
}
