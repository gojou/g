package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		hello("World")
	} else {
		hello(args[1])
	}

}

func hello(s string) {
	fmt.Println("Hello " + s + "!")
}
