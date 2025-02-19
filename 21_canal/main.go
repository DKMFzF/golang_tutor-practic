package main

import (
	"fmt"
)

func main() {
	var msg chan string
	msg = make(chan string)

	fmt.Println(msg) // 0xc000098070

	// по этому используется анониманая функция
	go func() {
		msg <- "Hello" // после передаётся значение в канал (тоже произошел блок до момента пока не предасьбся данные)
	}()

	varString := <-msg // использование канала заблокировалось на этом этапе (пока даныне не передадуться)

	fmt.Println(varString)
}
