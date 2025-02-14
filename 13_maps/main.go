package main

import "fmt"

func main() {
	users := map[string]int{
		"Kirill":    1,
		"Maksim":    2,
		"Sergay":    3,
		"Alecsandr": 4,
	}

	index, exists := users["Sergay"]

	// Проверка на то существует ли в мапе занчение по ключу
	if exists {
		fmt.Println(index)
	} else {
		fmt.Println("Такого человека нет")
	}

	// Итерация по мате
	for kay, item := range users {
		fmt.Println(kay, item)
	}
}
