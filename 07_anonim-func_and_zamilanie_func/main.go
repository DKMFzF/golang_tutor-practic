package main

import "fmt"

func main() {
	inc := incremetn()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func incremetn() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}
