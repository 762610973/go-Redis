package main

import "fmt"

func main() {
	var a interface{}
	var b interface{}
	fmt.Println(a == b)
	var c *int
	a = c
	fmt.Println(a == nil)
	fmt.Println(c == nil)
	var d *int
	var e *int
	fmt.Println(d == e)
}
