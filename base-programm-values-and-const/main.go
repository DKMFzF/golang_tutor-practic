package main // определяет пакет в котором храниться файл (обязательно)

import (
	"fmt" // импорт библиотки
	"reflect"
)

// основаня функция точка входа в приложения (обязательно)
func main() {
	var message string = "asd"

	fmt.Println(reflect.TypeOf(message))
}
