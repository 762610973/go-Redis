package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type k struct {
}

func main() {
	ss := []int{1, 2, 3, 4}
	for _, v := range ss {
		v += 10
	}
	fmt.Println(ss)
	x := k{}
	y := k{}
	fmt.Printf("%p\n", &x)
	fmt.Printf("%p\n", &y)
	fmt.Println(unsafe.Sizeof(x))
	m := map[string]struct{}{}
	m["hello"] = struct{}{}
	if t, ok := m["hello"]; ok {
		fmt.Println("True")
		fmt.Println(t)
		fmt.Println(reflect.TypeOf(t))
	}
	s := make([]int, 3, 44)
	s[0] = 1
	s[1] = 12
	s[2] = 13
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("123"))
	fmt.Println(len("qer"))
	fmt.Println(len([]rune("中文qwe")))

}
