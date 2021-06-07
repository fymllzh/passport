package main

import "fmt"

func main() {
	s := a()
	fmt.Printf("%T\n", s)
}

func a() interface{} {
	return "aaaaa"
}
