package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	hello(args[1])
}

func hello(s string) {
	fmt.Println("Hello " + s + "!")
}
