package main

import "fmt"

func main() {
	/*m := make(map[int]int)
	// 持续读
	go func() {
		for {
			_ = m[1]
		}
	}()
	// 持续写
	go func() {
		for {
			m[2] = 2
		}
	}()
	//比for好
	select {}*/
	x := defer1()
	fmt.Println(x)
	x2 := defer2()
	fmt.Println(x2)
}
func defer1() int {
	x := 3
	defer func() {
		x++
	}()
	return x
}
func defer2() (x int) {
	x = 3
	defer func() {
		x++
	}()
	return x
}
