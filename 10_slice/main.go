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

	fmt.Println("----------")

	const sizeMatrix = 3
	matrix := make([][]int, sizeMatrix)

	// обычный цикл for
	counter := 0
	for x := 0; x < sizeMatrix; x++ {
		matrix[x] = make([]int, sizeMatrix)
		for y := 0; y < sizeMatrix; y++ {
			counter++
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	}

	// цикл while типо
	count := 0
	for count < 10 {
		count++
	}
	fmt.Println(count)
}

func printMsg(msg []string) (string error) {
	if len(msg) == 0 {
		return errors.New("len zero slice")
	}

	msg[2] = "залупа"
	fmt.Println(msg)
	return nil
}
