package main

import "fmt"

func main() {
	// msg := make(chan string, 3)

	// без буфуризации
	// go func() {
	// 	msg <- "Hello"
	// 	msg <- "Привет мир!"
	// 	msg <- "Привет мир! 2"
	// }()

	// // каждая операция с каналом блокируется выполнение go-рутины
	// fmt.Println(<-msg)
	// fmt.Println(<-msg)
	// fmt.Println(<-msg)

	// c буферизацией
	// msg <- "Hello"
	// msg <- "Привет мир!"
	// msg <- "Привет мир! 2"

	// fmt.Println(<-msg)

	msg2 := make(chan string, 3)
	msg2 <- "Привет 1"
	msg2 <- "Привет 2"
	msg2 <- "Привет 3"

	close(msg2)

	// for {
	// 	if value, ok := <-msg2; !ok {
	// 		fmt.Println("chanel closed")
	// 		break
	// 	} else {
	// 		fmt.Println(value)
	// 	}
	// }

	// более удобный способ
	for m := range msg2 {
		fmt.Println(m)
	}
}
