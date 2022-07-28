package main

import "fmt"

type testI interface {
	sayI()
}

type umbrella struct{}

func (u *umbrella) sayI() {
	fmt.Println("ok")
}
func main() {
	var u testI = &umbrella{}

	switch u.(type) {
	case *umbrella:
		fmt.Println("continue ok")
	}
	var t testI = &umbrella{}
	t.sayI()
}
