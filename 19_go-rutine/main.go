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

	// использование анонимных функций в go-рутинах
	go func() {
		time.Sleep(time.Second)
		fmt.Println("2")
	}()

	time.Sleep(time.Second + time.Second) // блокировка main

	// после идёт выполнение друхи go-рутин

	fmt.Println("4") // в конце выводиться main go рутина
}
