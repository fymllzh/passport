package main

import "fmt"

func main() {
	s := make(map[int]string)
	s[1] = "name"
	d := map[int]string{}
	d[1] = "name"
	fmt.Println(s, d)
}
