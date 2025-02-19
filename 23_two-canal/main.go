package main

import (
	"fmt"
	"time"
)

func main() {
	message1 := make(chan string)
	message2 := make(chan string)

	// первая go-рутина
	go func() {
		for {
			message1 <- "Канал 1. Прошло 200 mc"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	// вторая go-рутина
	go func() {
		for {
			message2 <- "Канал 2. Прошло 1000 mc"
			time.Sleep(time.Second)
		}
	}()

	// выполняем чтение без блокировок
	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}
}
