package main

import (
	"fmt"
)

type te struct {
	url string
	name string
}

func (t te) print() {
	fmt.Println(t.name + " : " + t.url)
}

func main() {
	zu := te{"https://golang.org", "hallo go"}

	zu.print()
}