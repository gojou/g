package main

import (
	"fmt"
)

func main() {
	hello("World")
}

func hello(s string) {
	fmt.Println("Hello " + s + "!")
}
