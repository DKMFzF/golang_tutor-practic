package main

import (
	"fmt"
)

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	fmt.Println(sayHello("Кирилл", 19))
	fmt.Println("сумма числе ->", sum(a, b))
	fmt.Println(checkAge(2))
}

// функция которая говрит привет!
func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello %s! Тебе %d лет!", name, age) // функция генерирует строку
}

// функция сложения чисел
func sum(x int, y int) int {
	return x + y
}

// функция проверки возраста
func checkAge(age int) (string, bool) {
	if age > 10 {
		return "Ого тебе больше 10 лет!", true
	} else if 5 < age && age < 10 {
		return "Ого тебе больше 5 но меньше 10 лет!", true
	} else {
		return fmt.Sprintf("Увы тебе меньше 5 лет...тебе всего лишь %d", age), false
	}
}
