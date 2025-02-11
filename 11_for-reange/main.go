package main

import (
	"fmt"
)

func main() {
	slice := []string{
		"mesg 1",
		"mesg 2",
		"mesg 3",
	}

	// цикл для работы со слайсами
	// с помощью ключевого слова range
	// что бы итерироваться по слайсам
	for index, scl := range slice {
		fmt.Println(index+1, scl)
	}
}
