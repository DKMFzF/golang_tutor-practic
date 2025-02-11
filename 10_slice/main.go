package main

import (
	"errors"
	"fmt"
)

func main() {
	// слайсы
	sclie := []string{"1", "2", "3"}
	printMsg(sclie)
	fmt.Println(sclie)

	// паника в слайсах
	// var panicSlice []string
	// тут будет runtime ошибка, т.к. Go язык с статичной типизацией
	// компилятор не может предвидеть такую ошикуб, так как в run time
	// уходит массив с фиксированной длинной
	// panicSlice[0] = "1"
	// fmt.Println(len(panicSlice)) // 0

	// что бы избежать паники
	noPanic := make([]string, 5)
	// noPanic = append(noPanic, "6")
	fmt.Println(noPanic)
	fmt.Println(len(noPanic))
	fmt.Println(cap(noPanic))
}

func printMsg(msg []string) (string error) {
	if len(msg) == 0 {
		return errors.New("len zero slice")
	}

	msg[2] = "залупа"
	fmt.Println(msg)
	return nil
}
