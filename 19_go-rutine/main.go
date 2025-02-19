package main

import (
	"fmt"
	"time"
)

func main() {
	// добавление в стек
	go fmt.Println("1")
	go fmt.Println("1")
	go fmt.Println("1")

	time.Sleep(time.Second) // блокировка main

	// после идёт выполнение друхи go-рутин

	fmt.Println("4") // в конце выводиться main go рутина
}
