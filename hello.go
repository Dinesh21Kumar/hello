package main

import (
	"fmt"

	"rsc.io/quote"
)

func hello(s string) string {
	return "hello " + s
}

func getQuote() string {
	return quote.Glass()
}

func main() {
	fmt.Print(getQuote())
}
