package main // определяет пакет в котором храниться файл (обязательно)

import (
	"fmt" // импорт библиотки
)

// основаня функция точка входа в приложения (обязательно)
func main() {
	var varString string // путая строка (применяется с двумя ковычками)

	// все типы для int
	var number int
	var number8 int8
	var number16 int16
	var number32 int32
	var number64 int64

	// все типы для uint типа
	var numberUint uint // означает что это целое без знаковое число
	var numberUint8 uint8
	var numberUint16 uint16
	var numberUint32 uint32
	var numberUint64 uint64

	// float типы
	var float32 float32
	var float64 float64

	// boolean type
	var bool bool

	// byte/s type ASCII
	var byteVar byte
	var bytes = []byte("asd")

	// rune
	var varRune rune // синоним типа int32 и занимает 4 байта (применяется с одинарными ковычками)
	// var varRuneTwo rune = 's' // получаем на выходе бит

	fmt.Println("string type")
	fmt.Println(varString)

	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("nubmers type")
	fmt.Println(number)
	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)

	fmt.Println("")
	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("uint type")
	fmt.Println(numberUint)
	fmt.Println(numberUint8)
	fmt.Println(numberUint16)
	fmt.Println(numberUint32)
	fmt.Println(numberUint64)

	fmt.Println("")
	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("float type")
	fmt.Println(float32)
	fmt.Println(float64)

	fmt.Println("")
	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("bool type")
	fmt.Println(bool) // flase по дефолту

	fmt.Println("")
	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("byte/s type")
	fmt.Println(byteVar)
	fmt.Println(bytes)

	fmt.Println("")
	fmt.Println("-------------")
	fmt.Println("")

	fmt.Println("rune type")
	fmt.Println(varRune)
}
