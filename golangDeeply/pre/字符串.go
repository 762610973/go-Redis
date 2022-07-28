package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("123"))
	fmt.Println(len("qer"))
	fmt.Println(len([]rune("中文qwe")))
	s := "慕课网"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sh.Len)
	x := 3
	fmt.Println(unsafe.Sizeof(x))
	sl := []int{1, 2, 3, 34, 5}
	fmt.Println(unsafe.Sizeof(sl))
	t := []int{1, 5, 7, 0}
	fmt.Println(t)
}
