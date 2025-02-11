package main

import "fmt"

var msg string
var varIntNil *int
var number int

// init() - используется для инициализации
// глобальных переменных и дургих структур данных
func init() {
	msg = "скоро я тану ниньзя"
	number = 123
	varIntNil = &number
}

func main() {
	printMsg(&msg) // эта операция называется референсинг (ссылаемся)
	fmt.Println(msg)

	chacgeMsg(msg)
	fmt.Println(&msg)

	fmt.Println(&number)   // ссылаемся на область в памяти
	fmt.Println(varIntNil) // изначально nil но потом мы присваем указатель на области в масяти number

	*varIntNil = 1 // обращение к значению в участке памяти number
	fmt.Println(number)
}

func printMsg(msg *string) { // эта операция называется дерефересин (образаемся к облости памяти где храниться msg)
	*msg += "{ Дополнительый текст }" // тут у нас создаётся копия msg
	fmt.Println(*msg)
	// после того как функция закончила свою работу стек очищается
}

func chacgeMsg(mst string) {
	fmt.Println(&mst)
}
