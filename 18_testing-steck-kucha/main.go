package main

import "fmt"

func main() {
	// в стек добавлена переменн функция createPrt, потом в стек добавлена переменная a
	a := createPrt() // переменная а продолжить жить, потому что в Go есть механизм Escape analysis (переменная отправляется в кучу)

	// всё что было связанно с функцией в стеке удаляется

	fmt.Println(*a) // обращаемся к a котоаря лежит в куче

	*a = 10 // меняем значение переменной

	fmt.Println(*a) // выводим значение переменной
}

func createPrt() *int {
	a := 5
	return &a
}
