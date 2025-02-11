package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := checkAge(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func checkAge(age int) (string, error) {
	if age > 10 {
		return "Тебе больше 10 лет и ты принят!", nil
	} else {
		return "Тебе меньше 10 лет и так не пойдёт\n", errors.New("you are too young")
	}
}
