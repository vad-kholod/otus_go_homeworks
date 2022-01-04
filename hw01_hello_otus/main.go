package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const greeting = "Hello, OTUS!"

func main() {
	fmt.Print(stringutil.Reverse(greeting))
}
