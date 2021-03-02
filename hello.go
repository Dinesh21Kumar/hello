package main

import "fmt"

func hello(s string) string {
	return "hello " + s
}

func main() {
	fmt.Print(hello("dinesh"))
}
