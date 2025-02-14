package main

import (
	"fmt"
)

// Паника в Golang - это ошибка в runtime (любые ошибки при выполнении, которые нельзя обработать)

func main() {

	defer handlerPanic() // функция спасения проги от угрозы паники

	messages := []string{
		"msg1",
		"msg2",
		"msg3",
		"msg4",
		"msg5",
	}

	messages[5] = "msg6"

	printMessage(messages)
}

func printMessage(message []string) {
	fmt.Println(message)
}

// Функция которая обрабатывает панику
func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r) // выводит сообщение об ошибке
	}
	fmt.Println("handlerPanic() Выполнилась успешно")
}
