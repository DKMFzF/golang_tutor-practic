package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// массив
	message := [3]string{"1", "2", "3"}
	fmt.Println(message)
	fmt.Println(message[0])
	fmt.Println(message[1])
	fmt.Println(message[2])

	// перезаписали
	message[2] = "228"
	fmt.Println(message)

	// errorSlice := [3]string{}
	printMsg(&message)
	fmt.Println(reflect.TypeOf(message))
}

// работа с массивами
func printMsg(msg *[3]string) (string error) {
	if len(msg) == 0 {
		return errors.New("len zero slice")
	}

	msg[2] = "залупа"
	fmt.Println(msg)
	return nil
}
