package main

import "fmt"

// Паника в Golang - это ошибка в runtime (любые ошибки при выполнении, которые нельзя обработать)

func main() {

	defer fmt.Println(printMessage())
	defer fmt.Println("Хуй")
	fmt.Println("main()")

	panic("ААААА ПАНИКА!!")
}

func printMessage() string {
	return "printMessage()"
}
