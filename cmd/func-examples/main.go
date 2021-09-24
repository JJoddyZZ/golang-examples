package main

import (
	"fmt"

	"github.com/JJoddyZZ/golang-examples/somefunc"
)

func main() {
	f := somefunc.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
