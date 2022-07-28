package main

import (
	"fmt"
	"unsafe"
)

type demo struct {
	a int64
	z struct{}
}

type s1 struct {
	a int32
	b int32
}

type s2 struct {
	a int16
	b int64
}
type user struct {
	a int32
	b []int32
	c string
	d bool
	e struct{}
}

func main() {
	x := []int32{}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(s1{}))
	fmt.Println(unsafe.Sizeof(s2{}))
	fmt.Println(unsafe.Sizeof(demo{}))
	fmt.Println(unsafe.Sizeof(user{}), unsafe.Alignof(user{}))
}
