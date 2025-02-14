package main

import "fmt"

func main() {
	fmt.Println(sum(1, 1, 1, 1))
}

func sum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(numbers); i++ {
		count += numbers[i]
		fmt.Println(count, numbers[i])
	}
	return count
}
